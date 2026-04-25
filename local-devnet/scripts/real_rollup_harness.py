#!/usr/bin/env python3
"""Shared helpers for deploying and testing the real rollup contracts locally."""

from __future__ import annotations

import json
import re
import shutil
import subprocess
from dataclasses import dataclass
from pathlib import Path
from typing import Iterable

import rlp
from eth_abi import encode
from eth_tester import EthereumTester, PyEVMBackend
from solcx import compile_standard, install_solc
from web3 import Web3
from web3.contract import Contract
from web3.providers.eth_tester import EthereumTesterProvider
from web3._utils.events import EventLogErrorFlags


ROOT = Path(__file__).resolve().parents[2]
ROLLUP_ROOT = ROOT / "Rollup-Offchain"
ROLLUP_CONTRACT_DIR = ROOT / "Rollup-Offchain" / "contract"
VENDOR_DIR = ROOT / "local-devnet" / "vendor"
SOLC_VERSION = "0.8.34"
TX_GAS_LIMIT = 25_000_000
EVM_VERSION = "london"


@dataclass
class RollupChain:
    name: str
    chain_id: int
    w3: Web3
    tester: EthereumTester
    deployer: str
    data_types: Contract
    merkle_utils: Contract
    mpt_verifier: Contract
    zk_verifier: Contract
    local_state_manager: Contract
    transaction_executor: Contract
    cross_rollup: Contract
    bcp_manager: Contract


def _apply_window_config(source_name: str, content: str, window_config: dict | None) -> str:
    if not window_config:
        return content

    replacements: list[tuple[str, str, str]] = []
    if source_name == "CrossRollup.sol" and "dispute_time" in window_config:
        replacements.append(
            (
                r"uint256 public constant DISPUTE_TIME = \d+;",
                f"uint256 public constant DISPUTE_TIME = {int(window_config['dispute_time'])};",
                "DISPUTE_TIME",
            )
        )
    if source_name in {"BCPManager.sol", "ChallengeManager.sol"}:
        if "wait_period" in window_config:
            replacements.append(
                (
                    r"uint256 constant WAIT_PERIOD = \d+;",
                    f"uint256 constant WAIT_PERIOD = {int(window_config['wait_period'])};",
                    "WAIT_PERIOD",
                )
            )
        if "confirm_period" in window_config:
            replacements.append(
                (
                    r"uint256 constant CONFIRM_PERIOD = \d+;",
                    f"uint256 constant CONFIRM_PERIOD = {int(window_config['confirm_period'])};",
                    "CONFIRM_PERIOD",
                )
            )

    updated = content
    for pattern, replacement, label in replacements:
        updated, count = re.subn(pattern, replacement, updated, count=1)
        if count != 1:
            raise ValueError(f"failed to apply {label} override in {source_name}")
    return updated


def _source_map(window_config: dict | None = None) -> dict[str, dict[str, str]]:
    files = [
        "DataTypes.sol",
        "HeaderLib.sol",
        "IZKVerifier.sol",
        "MerkleUtils.sol",
        "LocalStateManager.sol",
        "TransactionExecutor.sol",
        "CrossRollup.sol",
        "BCPManager.sol",
        "ChallengeManager.sol",
    ]
    sources = {
        f"Rollup-Offchain/contract/{filename}": {
            "content": _apply_window_config(
                filename,
                (ROLLUP_CONTRACT_DIR / filename).read_text(),
                window_config,
            )
        }
        for filename in files
    }
    sources["Rollup-Offchain/contract/MPTVerifier.sol"] = {
        "content": (ROLLUP_CONTRACT_DIR / "MPTVerifier.sol").read_text()
    }
    sources["Rollup-Offchain/contract/ZKVerifier.sol"] = {
        "content": (ROLLUP_CONTRACT_DIR / "ZKVerifier.sol").read_text()
    }
    sources["local-devnet/contracts/CommitmentStore.sol"] = {
        "content": (ROOT / "local-devnet" / "contracts" / "CommitmentStore.sol").read_text()
    }
    sources["@openzeppelin/contracts/utils/Strings.sol"] = {
        "content": (VENDOR_DIR / "@openzeppelin" / "contracts" / "utils" / "Strings.sol").read_text()
    }
    sources["solidity-rlp/contracts/RLPReader.sol"] = {
        "content": (VENDOR_DIR / "solidity-rlp" / "contracts" / "RLPReader.sol").read_text()
    }
    return sources


def ensure_real_zk_verifier() -> None:
    subprocess.run(
        ["go", "run", "./cmd/zktool", "export-verifier"],
        cwd=ROLLUP_ROOT,
        check=True,
        capture_output=True,
        text=True,
    )


def _run_zk_tool(*args: str) -> dict:
    result = subprocess.run(
        ["go", "run", "./cmd/zktool", *args],
        cwd=ROLLUP_ROOT,
        check=True,
        capture_output=True,
        text=True,
    )
    stdout = result.stdout.strip()
    json_start = stdout.find("{")
    json_end = stdout.rfind("}")
    if json_start < 0 or json_end < json_start:
        raise ValueError(f"zktool did not return JSON output: {stdout}")
    return json.loads(stdout[json_start : json_end + 1])


def _normalize_proof_bundle(proof_bundle: dict) -> dict:
    return {
        "a": [int(value) for value in proof_bundle["a"]],
        "b": [[int(value) for value in row] for row in proof_bundle["b"]],
        "c": [int(value) for value in proof_bundle["c"]],
        "input": [int(value) for value in proof_bundle["input"]],
    }


def generate_header_hash_proof(header_bytes: bytes) -> dict:
    proof_bundle = _run_zk_tool("prove-header-hash", "--hash", _canonical_header_hash(header_bytes).hex())
    return _normalize_proof_bundle(proof_bundle)


def compile_real_rollup_contracts(window_config: dict | None = None) -> dict:
    ensure_real_zk_verifier()

    solc_binary = shutil.which("solc")
    if solc_binary is None:
        install_solc(SOLC_VERSION)

    kwargs = {"solc_version": SOLC_VERSION}
    if solc_binary is not None:
        kwargs = {"solc_binary": solc_binary}

    all_sources = _source_map(window_config)

    def compile_subset(source_keys: list[str], via_ir: bool) -> dict:
        return compile_standard(
            {
                "language": "Solidity",
                "sources": {key: all_sources[key] for key in source_keys},
                "settings": {
                    "optimizer": {"enabled": True, "runs": 200},
                    "evmVersion": EVM_VERSION,
                    "viaIR": via_ir,
                    "outputSelection": {
                        "*": {
                            "*": ["abi", "evm.bytecode.object"],
                        }
                    },
                },
            },
            **kwargs,
        )

    core_sources = [
        "Rollup-Offchain/contract/DataTypes.sol",
        "Rollup-Offchain/contract/MerkleUtils.sol",
        "Rollup-Offchain/contract/LocalStateManager.sol",
        "Rollup-Offchain/contract/TransactionExecutor.sol",
        "Rollup-Offchain/contract/CrossRollup.sol",
        "Rollup-Offchain/contract/MPTVerifier.sol",
        "@openzeppelin/contracts/utils/Strings.sol",
        "solidity-rlp/contracts/RLPReader.sol",
    ]
    manager_sources = [
        "Rollup-Offchain/contract/DataTypes.sol",
        "Rollup-Offchain/contract/HeaderLib.sol",
        "Rollup-Offchain/contract/IZKVerifier.sol",
        "Rollup-Offchain/contract/BCPManager.sol",
        "Rollup-Offchain/contract/MPTVerifier.sol",
        "Rollup-Offchain/contract/ZKVerifier.sol",
        "@openzeppelin/contracts/utils/Strings.sol",
        "solidity-rlp/contracts/RLPReader.sol",
    ]
    support_sources = [
        "local-devnet/contracts/CommitmentStore.sol",
    ]

    compiled_core = compile_subset(core_sources, via_ir=True)
    compiled_managers = compile_subset(manager_sources, via_ir=True)
    compiled_support = compile_subset(support_sources, via_ir=False)

    merged_contracts = {}
    for compiled in (compiled_core, compiled_managers, compiled_support):
        for source_key, contracts in compiled.get("contracts", {}).items():
            merged_contracts.setdefault(source_key, {}).update(contracts)

    return {"contracts": merged_contracts}


def _artifact(compiled: dict, source_key: str, contract_name: str) -> tuple[list[dict], str]:
    artifact = compiled["contracts"][source_key][contract_name]
    return artifact["abi"], artifact["evm"]["bytecode"]["object"]


def _deploy(w3: Web3, abi: list[dict], bytecode: str, from_account: str, *args) -> Contract:
    factory = w3.eth.contract(abi=abi, bytecode=bytecode)
    tx_hash = factory.constructor(*args).transact({"from": from_account, "gas": TX_GAS_LIMIT})
    receipt = w3.eth.wait_for_transaction_receipt(tx_hash)
    return w3.eth.contract(address=receipt.contractAddress, abi=abi)


def _event_descs(events: list) -> list[str]:
    descriptions = []
    for event in events:
        args = getattr(event, "args", None)
        if args is not None and "desc" in args:
            descriptions.append(args["desc"])
    return descriptions


def _error_messages(events: list) -> list[str]:
    messages = []
    for event in events:
        args = getattr(event, "args", None)
        if args is not None and "error" in args:
            messages.append(args["error"])
    return messages


def _canonical_header_hash(header_bytes: bytes) -> bytes:
    if len(header_bytes) == 32:
        return header_bytes
    return bytes(Web3.keccak(header_bytes))


def _commitment_leaf_hash(commitment_hash: bytes) -> bytes:
    return bytes(Web3.keccak(commitment_hash))


def build_commitment_proof(chain: RollupChain, block_number: int) -> dict[str, object]:
    block_len = int(chain.cross_rollup.functions.BlockLen().call())
    commitment_hashes = [block_commitment_hash(chain, index) for index in range(block_len)]
    if block_number >= len(commitment_hashes):
        raise IndexError(f"block {block_number} out of range for commitment proof")

    level = [_commitment_leaf_hash(commitment_hash) for commitment_hash in commitment_hashes]
    index = block_number
    siblings: list[bytes] = []
    while len(level) > 1:
        if len(level) % 2 == 1:
            level.append(bytes(32))

        siblings.append(level[index ^ 1])
        level = [
            bytes(Web3.keccak(level[i] + level[i + 1]))
            for i in range(0, len(level), 2)
        ]
        index //= 2

    return {
        "commitment_hash": commitment_hashes[block_number],
        "path": block_number,
        "siblings": siblings,
        "root": level[0],
    }


def build_state_root_bundle(header_state_root: bytes, commitment_root: bytes) -> bytes:
    return header_state_root + commitment_root


def encode_state_proof(commitment_proof: dict[str, object]) -> list[bytes]:
    state_proof = [
        commitment_proof["commitment_hash"],
        int(commitment_proof["path"]).to_bytes(32, "big"),
    ]
    state_proof.extend(commitment_proof["siblings"])
    return state_proof


def build_chain(name: str, chain_id: int, compiled: dict, source_chain_id: int = 1, target_chain_id: int = 2) -> RollupChain:
    backend = PyEVMBackend(genesis_parameters={"gas_limit": 30_000_000})
    tester = EthereumTester(backend=backend)
    w3 = Web3(EthereumTesterProvider(tester))
    deployer = w3.eth.accounts[0]

    data_types = _deploy(
        w3,
        *_artifact(compiled, "Rollup-Offchain/contract/DataTypes.sol", "DataTypes"),
        deployer,
    )
    merkle_utils = _deploy(
        w3,
        *_artifact(compiled, "Rollup-Offchain/contract/MerkleUtils.sol", "MerkleUtils"),
        deployer,
    )
    mpt_verifier = _deploy(
        w3,
        *_artifact(compiled, "Rollup-Offchain/contract/MPTVerifier.sol", "MPTVerifier"),
        deployer,
    )
    local_state_manager = _deploy(
        w3,
        *_artifact(compiled, "Rollup-Offchain/contract/LocalStateManager.sol", "LocalStateManager"),
        deployer,
        chain_id,
        merkle_utils.address,
        data_types.address,
    )
    zk_verifier = _deploy(
        w3,
        *_artifact(compiled, "Rollup-Offchain/contract/ZKVerifier.sol", "Verifier"),
        deployer,
    )
    transaction_executor = _deploy(
        w3,
        *_artifact(compiled, "Rollup-Offchain/contract/TransactionExecutor.sol", "TransactionExecutor"),
        deployer,
        mpt_verifier.address,
        data_types.address,
        merkle_utils.address,
        local_state_manager.address,
        source_chain_id,
        target_chain_id,
    )
    cross_rollup = _deploy(
        w3,
        *_artifact(compiled, "Rollup-Offchain/contract/CrossRollup.sol", "CrossRollup"),
        deployer,
        chain_id,
        deployer,
        transaction_executor.address,
        merkle_utils.address,
        data_types.address,
        mpt_verifier.address,
        local_state_manager.address,
    )
    bcp_manager = _deploy(
        w3,
        *_artifact(compiled, "Rollup-Offchain/contract/BCPManager.sol", "BCPManager"),
        deployer,
        chain_id,
        data_types.address,
        zk_verifier.address,
        mpt_verifier.address,
    )
    local_state_manager.functions.setRollupAddress(cross_rollup.address).transact({"from": deployer})
    cross_rollup.functions.setArbiterAddress(bcp_manager.address).transact({"from": deployer})

    return RollupChain(
        name=name,
        chain_id=chain_id,
        w3=w3,
        tester=tester,
        deployer=deployer,
        data_types=data_types,
        merkle_utils=merkle_utils,
        mpt_verifier=mpt_verifier,
        zk_verifier=zk_verifier,
        local_state_manager=local_state_manager,
        transaction_executor=transaction_executor,
        cross_rollup=cross_rollup,
        bcp_manager=bcp_manager,
    )


def mine_blocks(chain: RollupChain, block_count: int) -> dict[str, int]:
    before = int(chain.w3.eth.block_number)
    if block_count > 0:
        chain.tester.mine_blocks(block_count)
    after = int(chain.w3.eth.block_number)
    return {"before": before, "after": after, "mined": max(block_count, 0)}


def seed_account(chain: RollupChain, account: str, amount: int) -> None:
    tx_hash = chain.local_state_manager.functions.depositCoin(account, amount).transact(
        {"from": chain.deployer, "gas": TX_GAS_LIMIT}
    )
    chain.w3.eth.wait_for_transaction_receipt(tx_hash)


def l2_tx(from_addr: str, to_addr: str, value: int, chain_id: int, style: int) -> tuple:
    return (from_addr, to_addr, value, chain_id, style)


def l2_transition(account: str, value: int, style: int, chain_id: int) -> tuple:
    return (account, value, style, chain_id)


def build_transition_list(txs: list[tuple]) -> list[tuple]:
    transitions = []
    for from_addr, to_addr, value, _chain_id, style in txs:
        if style == 1:
            transitions.append(l2_transition(from_addr, value, 1, 1))
            transitions.append(l2_transition(to_addr, value, 2, 2))
        else:
            transitions.append(l2_transition(from_addr, value, 1, 2))
            transitions.append(l2_transition(to_addr, value, 2, 1))
    return transitions


def tx_root(chain: RollupChain, txs: list[tuple]) -> bytes:
    return chain.data_types.functions.TransactionsListHash(txs).call()


def transition_root(chain: RollupChain, transitions: list[tuple]) -> bytes:
    return chain.data_types.functions.L2TransitionListHash(transitions).call()


def block_commitment_hash(chain: RollupChain, block_number: int) -> bytes:
    block = chain.cross_rollup.functions.L2Blocks(block_number).call()
    encoded = encode(
        ["uint256", "bytes32", "bytes32", "bytes32", "uint256"],
        [block[0], block[1], block[2], block[3], block[4]],
    )
    return Web3.keccak(encoded)


def get_block_header_bundle(chain: RollupChain, block_number: int) -> dict[str, bytes | int]:
    block = chain.tester.backend.chain.get_canonical_block_by_number(block_number)
    return {
        "block_number": int(block.number),
        "header_bytes": rlp.encode(block.header),
        "state_root": bytes(block.header.state_root),
        "hash": bytes(block.header.hash),
        "parent_hash": bytes(block.header.parent_hash),
    }


def verify_generated_proof(
    chain: RollupChain,
    proof_bundle: dict,
    public_inputs: list[int] | None = None,
) -> bool:
    inputs = public_inputs or proof_bundle["input"]
    return chain.zk_verifier.functions.verifyProof(
        proof_bundle["a"],
        proof_bundle["b"],
        proof_bundle["c"],
        inputs,
    ).call()


def commit_block(chain: RollupChain, block_number: int, txs: list[tuple], transitions: list[tuple]) -> dict:
    computed_tx_root = tx_root(chain, txs)
    computed_transition_root = transition_root(chain, transitions)
    tx_hash = chain.cross_rollup.functions.commitBlock(
        block_number,
        computed_tx_root,
        computed_transition_root,
        txs,
        transitions,
    ).transact({"from": chain.deployer, "gas": TX_GAS_LIMIT})
    receipt = chain.w3.eth.wait_for_transaction_receipt(tx_hash)
    events = chain.cross_rollup.events.RollupBlockNotify().process_receipt(
        receipt, errors=EventLogErrorFlags.Ignore
    )
    error_events = chain.cross_rollup.events.ErrorNotify().process_receipt(
        receipt, errors=EventLogErrorFlags.Ignore
    )
    return {
        "tx_hash": receipt["transactionHash"].hex(),
        "status": int(receipt["status"]),
        "gas_used": int(receipt["gasUsed"]),
        "events": _event_descs(events),
        "errors": _error_messages(error_events),
        "tx_root": computed_tx_root.hex(),
        "transition_root": computed_transition_root.hex(),
    }


def try_commit_block(chain: RollupChain, block_number: int, txs: list[tuple], transitions: list[tuple]) -> dict:
    computed_tx_root = tx_root(chain, txs)
    computed_transition_root = transition_root(chain, transitions)
    try:
        return commit_block(chain, block_number, txs, transitions)
    except Exception as exc:
        return {
            "tx_hash": None,
            "status": 0,
            "gas_used": 0,
            "events": [],
            "errors": [str(exc)],
            "tx_root": computed_tx_root.hex(),
            "transition_root": computed_transition_root.hex(),
            "reverted": True,
        }


def confirm_latest(chain: RollupChain) -> dict:
    tx_hash = chain.cross_rollup.functions.checkCommit().transact(
        {"from": chain.deployer, "gas": TX_GAS_LIMIT}
    )
    receipt = chain.w3.eth.wait_for_transaction_receipt(tx_hash)
    events = chain.cross_rollup.events.RollupBlockNotify().process_receipt(
        receipt, errors=EventLogErrorFlags.Ignore
    )
    return {
        "tx_hash": receipt["transactionHash"].hex(),
        "status": int(receipt["status"]),
        "gas_used": int(receipt["gasUsed"]),
        "events": _event_descs(events),
    }


def execution_challenge(chain: RollupChain, block_number: int, txs: list[tuple]) -> dict:
    tx_hash = chain.cross_rollup.functions.ExecutionChallenge(block_number, txs).transact(
        {"from": chain.deployer, "gas": TX_GAS_LIMIT}
    )
    receipt = chain.w3.eth.wait_for_transaction_receipt(tx_hash)
    events = chain.cross_rollup.events.RollupBlockNotify().process_receipt(
        receipt, errors=EventLogErrorFlags.Ignore
    )
    error_events = chain.cross_rollup.events.ErrorNotify().process_receipt(
        receipt, errors=EventLogErrorFlags.Ignore
    )
    return {
        "tx_hash": receipt["transactionHash"].hex(),
        "status": int(receipt["status"]),
        "gas_used": int(receipt["gasUsed"]),
        "events": _event_descs(events),
        "errors": _error_messages(error_events),
    }


def checkpoint_remote_chain(
    chain: RollupChain,
    remote_chain_id: int,
    checkpoint_index: int,
    remote_block_id: int,
    header_bytes: bytes,
    state_root: bytes,
    contract_address: str,
) -> dict:
    tx_hash = chain.bcp_manager.functions.CheckPoint(
        (
            checkpoint_index,
            remote_chain_id,
            1,
            remote_block_id,
            state_root,
            _canonical_header_hash(header_bytes),
            contract_address,
        )
    ).transact({"from": chain.deployer, "gas": TX_GAS_LIMIT})
    receipt = chain.w3.eth.wait_for_transaction_receipt(tx_hash)
    events = chain.bcp_manager.events.BlockInfoNotify().process_receipt(
        receipt, errors=EventLogErrorFlags.Ignore
    )
    return {
        "tx_hash": receipt["transactionHash"].hex(),
        "status": int(receipt["status"]),
        "gas_used": int(receipt["gasUsed"]),
        "events": _event_descs(events),
    }


def create_consistency_challenge(
    chain: RollupChain,
    index: str,
    remote_chain_id: int,
    local_block_id: int,
    remote_block_id: int,
    remote_header_bytes: bytes,
    remote_state_root: bytes,
    remote_state_proof: list[bytes],
    questioned_rollup: str,
    challenger_account: str | None = None,
) -> dict:
    sender = challenger_account or chain.deployer
    tx_hash = chain.bcp_manager.functions.ChallengeCreate(
        (
            remote_chain_id,
            local_block_id,
            remote_block_id,
            remote_state_root,
            _canonical_header_hash(remote_header_bytes),
            remote_state_proof,
            questioned_rollup,
        ),
        index,
        remote_header_bytes,
    ).transact({"from": sender, "gas": TX_GAS_LIMIT})
    receipt = chain.w3.eth.wait_for_transaction_receipt(tx_hash)
    events = chain.bcp_manager.events.ChallengeStateNotify().process_receipt(
        receipt, errors=EventLogErrorFlags.Ignore
    )
    return {
        "tx_hash": receipt["transactionHash"].hex(),
        "status": int(receipt["status"]),
        "gas_used": int(receipt["gasUsed"]),
        "events": _event_descs(events),
    }


def challenge_question(
    chain: RollupChain,
    index: str,
    ack: bool,
    sender: str,
) -> dict:
    tx_hash = chain.bcp_manager.functions.ChallengeQuestion(index, ack).transact(
        {"from": sender, "gas": TX_GAS_LIMIT}
    )
    receipt = chain.w3.eth.wait_for_transaction_receipt(tx_hash)
    events = chain.bcp_manager.events.ChallengeStateNotify().process_receipt(
        receipt, errors=EventLogErrorFlags.Ignore
    )
    return {
        "tx_hash": receipt["transactionHash"].hex(),
        "status": int(receipt["status"]),
        "gas_used": int(receipt["gasUsed"]),
        "events": _event_descs(events),
    }


def challenge_response(
    chain: RollupChain,
    index: str,
    remote_chain_id: int,
    remote_block_id: int,
    header_bytes: bytes,
    state_root: bytes,
    sender: str | None = None,
) -> dict:
    tx_hash = chain.bcp_manager.functions.ChallengeResponse(
        index,
        (
            remote_chain_id,
            remote_block_id,
            state_root,
            _canonical_header_hash(header_bytes),
        ),
    ).transact({"from": sender or chain.deployer, "gas": TX_GAS_LIMIT})
    receipt = chain.w3.eth.wait_for_transaction_receipt(tx_hash)
    events = chain.bcp_manager.events.ChallengeStateNotify().process_receipt(
        receipt, errors=EventLogErrorFlags.Ignore
    )
    return {
        "tx_hash": receipt["transactionHash"].hex(),
        "status": int(receipt["status"]),
        "events": _event_descs(events),
    }


def final_consistency_challenge(
    chain: RollupChain,
    index: str,
    begin_header_bytes: bytes,
    end_header_bytes: bytes,
    sender: str | None = None,
) -> dict:
    proof_bundle = generate_header_hash_proof(end_header_bytes)
    tx_hash = chain.bcp_manager.functions.FinalChallenge(
        index,
        (
            proof_bundle["a"],
            proof_bundle["b"],
            proof_bundle["c"],
        ),
        begin_header_bytes,
        end_header_bytes,
    ).transact({"from": sender or chain.deployer, "gas": TX_GAS_LIMIT})
    receipt = chain.w3.eth.wait_for_transaction_receipt(tx_hash)
    events = chain.bcp_manager.events.ChallengeStateNotify().process_receipt(
        receipt, errors=EventLogErrorFlags.Ignore
    )
    rollup_events = chain.cross_rollup.events.RollupBlockNotify().process_receipt(
        receipt, errors=EventLogErrorFlags.Ignore
    )
    return {
        "tx_hash": receipt["transactionHash"].hex(),
        "status": int(receipt["status"]),
        "gas_used": int(receipt["gasUsed"]),
        "events": _event_descs(events),
        "rollup_events": _event_descs(rollup_events),
        "proof_inputs": proof_bundle["input"],
    }


def finalize_consistency_timeout(
    chain: RollupChain,
    index: str,
    sender: str | None = None,
) -> dict:
    tx_hash = chain.bcp_manager.functions.FinalizeChallenge(index).transact(
        {"from": sender or chain.deployer, "gas": TX_GAS_LIMIT}
    )
    receipt = chain.w3.eth.wait_for_transaction_receipt(tx_hash)
    events = chain.bcp_manager.events.ChallengeStateNotify().process_receipt(
        receipt, errors=EventLogErrorFlags.Ignore
    )
    rollup_events = chain.cross_rollup.events.RollupBlockNotify().process_receipt(
        receipt, errors=EventLogErrorFlags.Ignore
    )
    return {
        "tx_hash": receipt["transactionHash"].hex(),
        "status": int(receipt["status"]),
        "gas_used": int(receipt["gasUsed"]),
        "events": _event_descs(events),
        "rollup_events": _event_descs(rollup_events),
    }


def get_local_state(chain: RollupChain, account: str) -> dict:
    local_state = chain.local_state_manager.functions.getLocalState(account).call()
    return {
        "account": local_state[0],
        "value": int(local_state[1]),
        "lock": int(local_state[2]),
    }


def commit_and_confirm(chain: RollupChain, txs: list[tuple], transitions: list[tuple]) -> dict:
    block_number = int(chain.cross_rollup.functions.BlockLen().call())
    commit_result = commit_block(chain, block_number, txs, transitions)
    confirm_result = confirm_latest(chain)
    return {"commit": commit_result, "confirm": confirm_result}


def generate_accounts(prefix: str, count: int) -> list[str]:
    return [f"{prefix}-{index:04d}" for index in range(count)]


def pairwise(values: Iterable[str]) -> list[tuple[str, str]]:
    items = list(values)
    half = len(items) // 2
    return list(zip(items[:half], items[half:]))

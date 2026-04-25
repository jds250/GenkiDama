#!/usr/bin/env python3
"""Run the full rollup integrity flow on anvil, including real eth_getProof state proofs."""

from __future__ import annotations

import json
import socket
import subprocess
import time
from contextlib import ExitStack, contextmanager
from dataclasses import dataclass
from pathlib import Path

import rlp
from eth_abi import encode
from web3 import Web3

from real_rollup_harness import (
    ROOT,
    TX_GAS_LIMIT,
    _artifact,
    _deploy,
    block_commitment_hash,
    build_transition_list,
    challenge_question,
    challenge_response,
    checkpoint_remote_chain,
    commit_and_confirm,
    commit_block,
    compile_real_rollup_contracts,
    create_consistency_challenge,
    execution_challenge,
    final_consistency_challenge,
    generate_header_hash_proof,
    get_local_state,
    l2_tx,
    seed_account,
    verify_generated_proof,
)


SCRIPT_DIR = Path(__file__).resolve().parent
ROLLUP_BLOCKS_SLOT = 6
ROLLUP_BLOCK_FIELD_COUNT = 6
ROLLUP_COMMITMENT_FIELD_COUNT = 5


@dataclass
class AnvilInstance:
    name: str
    chain_id: int
    port: int
    rpc_url: str
    process: subprocess.Popen[str]


def _wait_for_rpc(rpc_url: str, timeout: float = 20.0) -> Web3:
    deadline = time.time() + timeout
    w3 = Web3(Web3.HTTPProvider(rpc_url))
    while time.time() < deadline:
        try:
            if w3.is_connected():
                return w3
        except Exception:
            pass
        time.sleep(0.2)
    raise TimeoutError(f"RPC did not become ready: {rpc_url}")


def _port_is_open(port: int) -> bool:
    with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as sock:
        sock.settimeout(0.1)
        return sock.connect_ex(("127.0.0.1", port)) == 0


@contextmanager
def anvil_instance(name: str, port: int, chain_id: int):
    if _port_is_open(port):
        raise RuntimeError(f"port {port} is already in use")

    cmd = [
        "anvil",
        "--host",
        "127.0.0.1",
        "--port",
        str(port),
        "--chain-id",
        str(chain_id),
        "--accounts",
        "10",
        "--balance",
        "1000000",
        "--gas-limit",
        str(TX_GAS_LIMIT),
        "--silent",
    ]
    process = subprocess.Popen(
        cmd,
        stdout=subprocess.PIPE,
        stderr=subprocess.STDOUT,
        text=True,
    )
    rpc_url = f"http://127.0.0.1:{port}"
    try:
        _wait_for_rpc(rpc_url)
        yield AnvilInstance(name=name, chain_id=chain_id, port=port, rpc_url=rpc_url, process=process)
    finally:
        process.terminate()
        try:
            process.wait(timeout=5)
        except subprocess.TimeoutExpired:
            process.kill()
            process.wait(timeout=5)


def deploy_http_rollup_chain(name: str, chain_id: int, rpc_url: str, compiled: dict, source_chain_id: int = 1, target_chain_id: int = 2):
    w3 = _wait_for_rpc(rpc_url)
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

    return type(
        "HttpRollupChain",
        (),
        {
            "name": name,
            "chain_id": chain_id,
            "w3": w3,
            "tester": None,
            "deployer": deployer,
            "data_types": data_types,
            "merkle_utils": merkle_utils,
            "mpt_verifier": mpt_verifier,
            "zk_verifier": zk_verifier,
            "local_state_manager": local_state_manager,
            "transaction_executor": transaction_executor,
            "cross_rollup": cross_rollup,
            "bcp_manager": bcp_manager,
        },
    )()


def mine_empty_blocks(w3: Web3, count: int) -> None:
    for _ in range(count):
        response = w3.provider.make_request("evm_mine", [])
        if response.get("error"):
            raise RuntimeError(response["error"])


def get_http_header_bundle(rpc_url: str, block_number: int) -> dict[str, bytes | int]:
    w3 = _wait_for_rpc(rpc_url)
    response = w3.provider.make_request("eth_getBlockByNumber", [hex(block_number), False])
    if response.get("error"):
        raise RuntimeError(response["error"])

    payload = response["result"]
    header_fields: list[object] = [
        _hex_to_bytes(payload["parentHash"]),
        _hex_to_bytes(payload["sha3Uncles"]),
        _hex_to_bytes(payload["miner"]),
        _hex_to_bytes(payload["stateRoot"]),
        _hex_to_bytes(payload["transactionsRoot"]),
        _hex_to_bytes(payload["receiptsRoot"]),
        _hex_to_bytes(payload["logsBloom"]),
        _to_int(payload["difficulty"]),
        _to_int(payload["number"]),
        _to_int(payload["gasLimit"]),
        _to_int(payload["gasUsed"]),
        _to_int(payload["timestamp"]),
        _hex_to_bytes(payload["extraData"]),
        _hex_to_bytes(payload["mixHash"]),
        _hex_to_bytes(payload["nonce"]),
    ]

    optional_int_fields = ["baseFeePerGas", "blobGasUsed", "excessBlobGas"]
    optional_bytes_fields = ["withdrawalsRoot", "parentBeaconBlockRoot", "requestsHash"]

    if payload.get("baseFeePerGas") is not None:
        header_fields.append(_to_int(payload["baseFeePerGas"]))
    if payload.get("withdrawalsRoot") is not None:
        header_fields.append(_hex_to_bytes(payload["withdrawalsRoot"]))
    if payload.get("blobGasUsed") is not None:
        header_fields.append(_to_int(payload["blobGasUsed"]))
    if payload.get("excessBlobGas") is not None:
        header_fields.append(_to_int(payload["excessBlobGas"]))
    if payload.get("parentBeaconBlockRoot") is not None:
        header_fields.append(_hex_to_bytes(payload["parentBeaconBlockRoot"]))
    if payload.get("requestsHash") is not None:
        header_fields.append(_hex_to_bytes(payload["requestsHash"]))

    header_bytes = rlp.encode(header_fields)
    return {
        "block_number": _to_int(payload["number"]),
        "header_bytes": header_bytes,
        "state_root": _hex_to_bytes(payload["stateRoot"]),
        "hash": _hex_to_bytes(payload["hash"]),
        "parent_hash": _hex_to_bytes(payload["parentHash"]),
    }


def find_block_by_hash(rpc_url: str, target_hash: bytes, start_block: int) -> dict[str, bytes | int]:
    current = start_block
    while current >= 0:
        bundle = get_http_header_bundle(rpc_url, current)
        if bundle["hash"] == target_hash:
            return bundle
        current -= 1
    raise ValueError(f"could not find block for hash {target_hash.hex()}")


def _hex_to_bytes(value: str) -> bytes:
    if isinstance(value, bytes):
        return value
    if isinstance(value, int):
        hex_value = hex(value)[2:]
    else:
        hex_value = value[2:] if value.startswith("0x") else value
    if len(hex_value) % 2 == 1:
        hex_value = "0" + hex_value
    return bytes.fromhex(hex_value)


def _to_int(value) -> int:
    if isinstance(value, int):
        return value
    if isinstance(value, str):
        return int(value, 16) if value.startswith("0x") else int(value)
    raise TypeError(f"unsupported numeric value: {value!r}")


def _encode_storage_value_rlp(raw_word: bytes) -> bytes:
    trimmed = raw_word.lstrip(b"\x00")
    return rlp.encode(trimmed)


def fetch_account_storage_proof(
    w3: Web3,
    address: str,
    slot_indexes: list[int],
    block_number: int,
) -> dict:
    slot_bytes = [Web3.to_hex(slot_index.to_bytes(32, "big")) for slot_index in slot_indexes]
    response = w3.provider.make_request(
        "eth_getProof",
        [address, slot_bytes, hex(block_number)],
    )
    if response.get("error"):
        raise RuntimeError(response["error"])
    return response["result"]

def rollup_block_commitment_slots(block_index: int) -> list[int]:
    base_slot = int.from_bytes(Web3.keccak(ROLLUP_BLOCKS_SLOT.to_bytes(32, "big")), "big")
    base_slot += block_index * ROLLUP_BLOCK_FIELD_COUNT
    return [base_slot + offset for offset in range(ROLLUP_COMMITMENT_FIELD_COUNT)]


def encode_rollup_block_state_proof(
    commitment_hash: bytes,
    rollup_block_index: int,
    state_proof: dict,
    slot_indexes: list[int],
) -> list[bytes]:
    account_value = rlp.encode(
        [
            _to_int(state_proof["nonce"]),
            _to_int(state_proof["balance"]),
            _hex_to_bytes(state_proof["storageHash"]),
            _hex_to_bytes(state_proof["codeHash"]),
        ]
    )
    account_nodes = [_hex_to_bytes(node) for node in state_proof["accountProof"]]

    encoded = [
        commitment_hash,
        rollup_block_index.to_bytes(32, "big"),
        account_value,
        encode(["uint256"], [len(account_nodes)]),
        encode(["uint256"], [len(slot_indexes)]),
    ]
    encoded.extend(account_nodes)

    for slot_index, storage_entry in zip(slot_indexes, state_proof["storageProof"]):
        storage_key = Web3.keccak(slot_index.to_bytes(32, "big"))
        storage_value = _encode_storage_value_rlp(_hex_to_bytes(storage_entry["value"]))
        storage_nodes = [_hex_to_bytes(node) for node in storage_entry["proof"]]
        encoded.append(storage_key)
        encoded.append(storage_value)
        encoded.append(encode(["uint256"], [len(storage_nodes)]))
        encoded.extend(storage_nodes)

    return encoded


def main() -> None:
    compiled = compile_real_rollup_contracts()

    with ExitStack() as stack:
        chain_a_inst = stack.enter_context(anvil_instance("chain-a", 9545, 1))
        chain_b_inst = stack.enter_context(anvil_instance("chain-b", 9645, 2))
        consistency_inst = stack.enter_context(anvil_instance("consistency-chain", 9745, 1))
        remote_consistency_inst = stack.enter_context(anvil_instance("remote-consistency-chain", 9845, 2))

        chain_a = deploy_http_rollup_chain("chain-a", chain_a_inst.chain_id, chain_a_inst.rpc_url, compiled)
        chain_b = deploy_http_rollup_chain("chain-b", chain_b_inst.chain_id, chain_b_inst.rpc_url, compiled)
        consistency_chain = deploy_http_rollup_chain(
            "consistency-chain",
            consistency_inst.chain_id,
            consistency_inst.rpc_url,
            compiled,
        )
        remote_consistency_chain = deploy_http_rollup_chain(
            "remote-consistency-chain",
            remote_consistency_inst.chain_id,
            remote_consistency_inst.rpc_url,
            compiled,
        )

        alice = "alice@chain-a"
        bob = "bob@chain-b"
        seed_account(chain_a, alice, 100)
        seed_account(chain_b, bob, 1)

        before = {
            "chain_a_alice": get_local_state(chain_a, alice),
            "chain_b_bob": get_local_state(chain_b, bob),
        }

        mirrored_txs = [l2_tx(alice, bob, 25, chain_a.chain_id, 1)]
        mirrored_transitions = build_transition_list(mirrored_txs)
        chain_a_commit = commit_block(chain_a, 0, mirrored_txs, mirrored_transitions)
        legal_challenge = execution_challenge(chain_a, 0, mirrored_txs)
        chain_a_confirm = commit_and_confirm(chain_b, mirrored_txs, mirrored_transitions)
        chain_a_confirm_result = chain_a.cross_rollup.functions.checkCommit().transact(
            {"from": chain_a.deployer, "gas": TX_GAS_LIMIT}
        )
        chain_a.w3.eth.wait_for_transaction_receipt(chain_a_confirm_result)

        after = {
            "chain_a_alice": get_local_state(chain_a, alice),
            "chain_b_bob": get_local_state(chain_b, bob),
        }

        seed_account(remote_consistency_chain, alice, 100)
        divergent_txs = [l2_tx(alice, bob, 24, consistency_chain.chain_id, 1)]
        divergent_transitions = build_transition_list(divergent_txs)
        remote_consistency_commit = commit_block(
            remote_consistency_chain,
            0,
            divergent_txs,
            divergent_transitions,
        )
        remote_consistency_confirm_tx = remote_consistency_chain.cross_rollup.functions.checkCommit().transact(
            {"from": remote_consistency_chain.deployer, "gas": TX_GAS_LIMIT}
        )
        remote_consistency_chain.w3.eth.wait_for_transaction_receipt(remote_consistency_confirm_tx)
        mine_empty_blocks(remote_consistency_chain.w3, 2)

        remote_final_block = int(remote_consistency_chain.w3.eth.block_number)
        remote_final_header = get_http_header_bundle(remote_consistency_inst.rpc_url, remote_final_block)
        remote_middle_header = find_block_by_hash(
            remote_consistency_inst.rpc_url,
            remote_final_header["parent_hash"],
            remote_final_block - 1,
        )
        remote_checkpoint_header = find_block_by_hash(
            remote_consistency_inst.rpc_url,
            remote_middle_header["parent_hash"],
            remote_middle_header["block_number"] - 1,
        )
        assert bytes(Web3.keccak(remote_middle_header["header_bytes"])) == remote_middle_header["hash"], remote_middle_header
        assert bytes(Web3.keccak(remote_final_header["header_bytes"])) == remote_final_header["hash"], remote_final_header
        assert remote_final_header["parent_hash"] == remote_middle_header["hash"], {
            "middle_hash": remote_middle_header["hash"].hex(),
            "final_parent": remote_final_header["parent_hash"].hex(),
            "middle_block": remote_middle_header["block_number"],
            "final_block": remote_final_header["block_number"],
        }
        remote_header_proof = verify_generated_proof(
            consistency_chain,
            generate_header_hash_proof(remote_final_header["header_bytes"]),
        )

        remote_commitment_hash = block_commitment_hash(remote_consistency_chain, 0)
        remote_rollup_slots = rollup_block_commitment_slots(0)
        state_proof = fetch_account_storage_proof(
            remote_consistency_chain.w3,
            remote_consistency_chain.cross_rollup.address,
            remote_rollup_slots,
            remote_final_block,
        )
        remote_state_proof = encode_rollup_block_state_proof(
            remote_commitment_hash,
            0,
            state_proof,
            remote_rollup_slots,
        )

        seed_account(consistency_chain, alice, 100)
        checkpoint_result = checkpoint_remote_chain(
            consistency_chain,
            remote_chain_id=remote_consistency_chain.chain_id,
            checkpoint_index=1,
            remote_block_id=remote_checkpoint_header["block_number"],
            header_bytes=remote_checkpoint_header["header_bytes"],
            state_root=remote_checkpoint_header["state_root"],
            contract_address=remote_consistency_chain.cross_rollup.address,
        )
        consistency_commit = commit_block(consistency_chain, 0, mirrored_txs, mirrored_transitions)
        consistency_create = create_consistency_challenge(
            consistency_chain,
            index="consistency-0",
            remote_chain_id=remote_consistency_chain.chain_id,
            local_block_id=0,
            remote_block_id=remote_final_header["block_number"],
            remote_header_bytes=remote_final_header["header_bytes"],
            remote_state_root=remote_final_header["state_root"],
            remote_state_proof=remote_state_proof,
            questioned_rollup=consistency_chain.cross_rollup.address,
        )
        questioner = consistency_chain.w3.eth.accounts[1]
        consistency_question = challenge_question(
            consistency_chain,
            index="consistency-0",
            ack=False,
            sender=questioner,
        )
        consistency_response = challenge_response(
            consistency_chain,
            index="consistency-0",
            remote_chain_id=remote_consistency_chain.chain_id,
            remote_block_id=remote_middle_header["block_number"],
            header_bytes=remote_middle_header["header_bytes"],
            state_root=remote_middle_header["state_root"],
            sender=consistency_chain.deployer,
        )
        consistency_ack = challenge_question(
            consistency_chain,
            index="consistency-0",
            ack=True,
            sender=questioner,
        )
        consistency_local_commitment_before = block_commitment_hash(consistency_chain, 0).hex()
        consistency_final = final_consistency_challenge(
            consistency_chain,
            index="consistency-0",
            begin_header_bytes=remote_middle_header["header_bytes"],
            end_header_bytes=remote_final_header["header_bytes"],
        )
        consistency_after = {
            "alice": get_local_state(consistency_chain, alice),
            "block_len": int(consistency_chain.cross_rollup.functions.BlockLen().call()),
        }

        assert before["chain_a_alice"]["value"] == 100
        assert before["chain_b_bob"]["value"] == 1
        assert chain_a_commit["events"] == ["1-0-commit success"]
        assert legal_challenge["errors"] == []
        assert legal_challenge["events"] == ["4-1-The Challenged Block is legal"]
        assert chain_a_confirm["commit"]["events"] == ["1-0-commit success"]
        assert chain_a_confirm["confirm"]["events"] == ["2-0-commit success"]
        assert after["chain_a_alice"]["value"] == 75
        assert after["chain_a_alice"]["lock"] == 0
        assert after["chain_b_bob"]["value"] == 26
        assert after["chain_b_bob"]["lock"] == 0
        assert remote_consistency_commit["events"] == ["1-0-commit success"]
        assert checkpoint_result["events"] == ["new info as the chain checkpoint"]
        assert consistency_commit["events"] == ["1-0-commit success"]
        assert consistency_create["events"] == ["create success!"]
        assert consistency_question["events"] == ["success! the commit was questioned"]
        assert consistency_response["events"] == ["challenge response submitted"]
        assert consistency_ack["events"] == ["True! In the last 1/2."]
        assert consistency_final["events"] == ["final proof success"], consistency_final
        assert consistency_final["rollup_events"] == ["3-0-rollback success"], consistency_final
        assert consistency_after["alice"]["value"] == 100
        assert consistency_after["alice"]["lock"] == 0
        assert consistency_after["block_len"] == 0

        result = {
            "summary": "full integrity flow passed on anvil",
            "deployments": {
                "chain_a": {
                    "cross_rollup": chain_a.cross_rollup.address,
                    "bcp_manager": chain_a.bcp_manager.address,
                },
                "chain_b": {
                    "cross_rollup": chain_b.cross_rollup.address,
                    "bcp_manager": chain_b.bcp_manager.address,
                },
                "consistency_chain": {
                    "cross_rollup": consistency_chain.cross_rollup.address,
                    "bcp_manager": consistency_chain.bcp_manager.address,
                },
                "remote_consistency_chain": {
                    "cross_rollup": remote_consistency_chain.cross_rollup.address,
                    "bcp_manager": remote_consistency_chain.bcp_manager.address,
                },
            },
            "before": before,
            "after": after,
            "mirrored_txs": mirrored_txs,
            "legal_challenge": {
                "commit": chain_a_commit,
                "challenge": legal_challenge,
            },
            "consistency_challenge": {
                "checkpoint": checkpoint_result,
                "create": consistency_create,
                "question": consistency_question,
                "response": consistency_response,
                "ack": consistency_ack,
                "final": consistency_final,
                "after": consistency_after,
                "remote_commitment_hash": remote_commitment_hash.hex(),
                "local_commitment_hash_before": consistency_local_commitment_before,
                "header_proof_valid": remote_header_proof,
                "state_proof": {
                    "account_nodes": len(state_proof["accountProof"]),
                    "storage_nodes": [len(entry["proof"]) for entry in state_proof["storageProof"]],
                    "storage_hash": state_proof["storageHash"],
                    "slot_values": [entry["value"] for entry in state_proof["storageProof"]],
                    "rollup_slots": [hex(slot) for slot in remote_rollup_slots],
                },
            },
            "notes": [
                "This run uses anvil eth_getProof for real account/storage proofs rooted in the remote L1 header state root.",
                "The proof now anchors directly to the remote CrossRollup storage layout and reconstructs the divergent block commitment from L2Blocks storage slots.",
                "The full flow covers mirrored commit/confirm, legal execution challenge validation, consistency challenge create/question/response/final-proof, and rollback.",
            ],
        }
        print(json.dumps(result, indent=2))


if __name__ == "__main__":
    main()

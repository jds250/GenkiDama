#!/usr/bin/env python3
"""Run a minimal dual-chain local integration flow."""

from __future__ import annotations

import json
import shutil
from dataclasses import dataclass
from pathlib import Path

from eth_tester import EthereumTester, PyEVMBackend
from solcx import compile_standard, install_solc
from web3 import Web3
from web3.providers.eth_tester import EthereumTesterProvider


ROOT = Path(__file__).resolve().parents[2]
CONTRACT_PATH = ROOT / "local-devnet" / "contracts" / "LocalCrossChainBridge.sol"
SOLC_VERSION = "0.8.20"


@dataclass
class LocalChain:
    name: str
    chain_id: int
    tester: EthereumTester
    w3: Web3
    deployer: str
    user_a: str
    user_b: str
    bridge: object


def compile_bridge() -> tuple[list[dict], str]:
    solc_binary = shutil.which("solc")
    if solc_binary is None:
        install_solc(SOLC_VERSION)

    source = CONTRACT_PATH.read_text()
    compile_kwargs = {
        "solc_version": SOLC_VERSION,
    }
    if solc_binary is not None:
        compile_kwargs = {
            "solc_binary": solc_binary,
        }

    compiled = compile_standard(
        {
            "language": "Solidity",
            "sources": {
                CONTRACT_PATH.name: {
                    "content": source,
                }
            },
            "settings": {
                "outputSelection": {
                    "*": {
                        "*": ["abi", "evm.bytecode.object"],
                    }
                }
            },
        },
        **compile_kwargs,
    )

    artifact = compiled["contracts"][CONTRACT_PATH.name]["LocalCrossChainBridge"]
    return artifact["abi"], artifact["evm"]["bytecode"]["object"]


def build_chain(name: str, chain_id: int, abi: list[dict], bytecode: str) -> LocalChain:
    backend = PyEVMBackend(genesis_parameters={"gas_limit": 30_000_000}, genesis_state=None)
    tester = EthereumTester(backend=backend)
    provider = EthereumTesterProvider(tester)
    w3 = Web3(provider)
    accounts = w3.eth.accounts
    deployer = accounts[0]
    user_a = accounts[1]
    user_b = accounts[2]

    bridge_factory = w3.eth.contract(abi=abi, bytecode=bytecode)
    tx_hash = bridge_factory.constructor(chain_id).transact({"from": deployer})
    receipt = w3.eth.wait_for_transaction_receipt(tx_hash)
    bridge = w3.eth.contract(address=receipt.contractAddress, abi=abi)

    return LocalChain(
        name=name,
        chain_id=chain_id,
        tester=tester,
        w3=w3,
        deployer=deployer,
        user_a=user_a,
        user_b=user_b,
        bridge=bridge,
    )


def relay_message(
    src_chain: LocalChain,
    dst_chain: LocalChain,
    from_account: str,
    to_account: str,
    amount: int,
) -> dict:
    tx_hash = src_chain.bridge.functions.commitOutbound(
        dst_chain.chain_id,
        to_account,
        amount,
    ).transact({"from": from_account})
    receipt = src_chain.w3.eth.wait_for_transaction_receipt(tx_hash)
    outbound_events = src_chain.bridge.events.OutboundCommitted().process_receipt(receipt)
    event = outbound_events[0]["args"]

    apply_hash = dst_chain.bridge.functions.applyInbound(
        event["messageId"],
        src_chain.chain_id,
        from_account,
        to_account,
        amount,
    ).transact({"from": dst_chain.deployer})
    apply_receipt = dst_chain.w3.eth.wait_for_transaction_receipt(apply_hash)

    return {
        "source_tx": receipt["transactionHash"].hex(),
        "dest_tx": apply_receipt["transactionHash"].hex(),
        "message_id": event["messageId"].hex(),
    }


def balance(chain: LocalChain, account: str) -> int:
    return chain.bridge.functions.balances(account).call()


def run() -> dict:
    abi, bytecode = compile_bridge()
    chain_a = build_chain("chain-a", 31337, abi, bytecode)
    chain_b = build_chain("chain-b", 31338, abi, bytecode)

    chain_a.bridge.functions.deposit(chain_a.user_a, 100).transact({"from": chain_a.deployer})
    chain_b.bridge.functions.deposit(chain_b.user_b, 30).transact({"from": chain_b.deployer})

    before = {
        "chain_a_user_a": balance(chain_a, chain_a.user_a),
        "chain_a_user_b": balance(chain_a, chain_a.user_b),
        "chain_b_user_a": balance(chain_b, chain_b.user_a),
        "chain_b_user_b": balance(chain_b, chain_b.user_b),
    }

    forward = relay_message(chain_a, chain_b, chain_a.user_a, chain_b.user_b, 25)
    backward = relay_message(chain_b, chain_a, chain_b.user_b, chain_a.user_a, 10)

    after = {
        "chain_a_user_a": balance(chain_a, chain_a.user_a),
        "chain_a_user_b": balance(chain_a, chain_a.user_b),
        "chain_b_user_a": balance(chain_b, chain_b.user_a),
        "chain_b_user_b": balance(chain_b, chain_b.user_b),
    }

    assert before["chain_a_user_a"] == 100
    assert before["chain_b_user_b"] == 30
    assert after["chain_a_user_a"] == 85
    assert after["chain_b_user_b"] == 45
    assert after["chain_a_user_b"] == 0
    assert after["chain_b_user_a"] == 0

    return {
        "environment": {
            "python_venv": str(ROOT / ".venv"),
            "solidity_contract": str(CONTRACT_PATH),
            "chains": [
                {
                    "name": chain_a.name,
                    "chain_id": chain_a.chain_id,
                    "bridge": chain_a.bridge.address,
                    "deployer": chain_a.deployer,
                    "user_a": chain_a.user_a,
                    "user_b": chain_a.user_b,
                },
                {
                    "name": chain_b.name,
                    "chain_id": chain_b.chain_id,
                    "bridge": chain_b.bridge.address,
                    "deployer": chain_b.deployer,
                    "user_a": chain_b.user_a,
                    "user_b": chain_b.user_b,
                },
            ],
        },
        "before": before,
        "after": after,
        "forward_relay": forward,
        "backward_relay": backward,
        "summary": "dual-chain local flow passed",
    }


if __name__ == "__main__":
    result = run()
    print(json.dumps(result, indent=2))

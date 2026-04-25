#!/usr/bin/env python3
"""Deploy the real rollup stack onto two running geth RPC endpoints."""

from __future__ import annotations

import argparse
import json
from pathlib import Path

from web3 import Web3
from web3.middleware import ExtraDataToPOAMiddleware

from real_rollup_harness import _artifact, _deploy, compile_real_rollup_contracts


ROOT = Path(__file__).resolve().parents[2]
DEFAULT_OUT = ROOT / "local-devnet" / "multinode-geth" / "deployments" / "latest.json"


def wait_connected(rpc_url: str) -> Web3:
    w3 = Web3(Web3.HTTPProvider(rpc_url))
    w3.middleware_onion.inject(ExtraDataToPOAMiddleware, layer=0)
    if not w3.is_connected():
        raise RuntimeError(f"could not connect to {rpc_url}")
    return w3


def deploy_chain(compiled: dict, rpc_url: str, chain_id: int, source_chain_id: int, target_chain_id: int) -> dict:
    w3 = wait_connected(rpc_url)
    if not w3.eth.accounts:
        raise RuntimeError(f"no unlocked accounts exposed by {rpc_url}")

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
    tx_hash = local_state_manager.functions.setRollupAddress(cross_rollup.address).transact({"from": deployer})
    w3.eth.wait_for_transaction_receipt(tx_hash)
    tx_hash = cross_rollup.functions.setArbiterAddress(bcp_manager.address).transact({"from": deployer})
    w3.eth.wait_for_transaction_receipt(tx_hash)

    return {
        "rpc_url": rpc_url,
        "chain_id": chain_id,
        "deployer": deployer,
        "contracts": {
            "DataTypes": data_types.address,
            "MerkleUtils": merkle_utils.address,
            "MPTVerifier": mpt_verifier.address,
            "LocalStateManager": local_state_manager.address,
            "ZKVerifier": zk_verifier.address,
            "TransactionExecutor": transaction_executor.address,
            "CrossRollup": cross_rollup.address,
            "BCPManager": bcp_manager.address,
        },
    }


def main() -> None:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--chain-a-rpc", default="http://127.0.0.1:8545")
    parser.add_argument("--chain-b-rpc", default="http://127.0.0.1:8645")
    parser.add_argument("--chain-a-id", type=int, default=1001)
    parser.add_argument("--chain-b-id", type=int, default=1002)
    parser.add_argument("--dispute-time", type=int, default=8)
    parser.add_argument("--wait-period", type=int, default=12)
    parser.add_argument("--confirm-period", type=int, default=24)
    parser.add_argument("--output", default=str(DEFAULT_OUT))
    args = parser.parse_args()

    compiled = compile_real_rollup_contracts(
        {
            "dispute_time": args.dispute_time,
            "wait_period": args.wait_period,
            "confirm_period": args.confirm_period,
        }
    )
    result = {
        "summary": "deployed real rollup stack to multinode geth devnet",
        "window_config": {
            "dispute_time": args.dispute_time,
            "wait_period": args.wait_period,
            "confirm_period": args.confirm_period,
        },
        "chain_a": deploy_chain(compiled, args.chain_a_rpc, args.chain_a_id, args.chain_a_id, args.chain_b_id),
        "chain_b": deploy_chain(compiled, args.chain_b_rpc, args.chain_b_id, args.chain_a_id, args.chain_b_id),
    }

    output_path = Path(args.output)
    output_path.parent.mkdir(parents=True, exist_ok=True)
    output_path.write_text(json.dumps(result, indent=2) + "\n")
    print(json.dumps(result, indent=2))


if __name__ == "__main__":
    main()

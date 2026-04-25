#!/usr/bin/env python3
"""Deploy and validate the real rollup contract stack on two local chains."""

from __future__ import annotations

import json

from real_rollup_harness import (
    block_commitment_hash,
    build_commitment_proof,
    build_state_root_bundle,
    build_transition_list,
    build_chain,
    challenge_question,
    challenge_response,
    commit_and_confirm,
    checkpoint_remote_chain,
    compile_real_rollup_contracts,
    create_consistency_challenge,
    encode_state_proof,
    execution_challenge,
    final_consistency_challenge,
    generate_header_hash_proof,
    get_block_header_bundle,
    get_local_state,
    l2_tx,
    seed_account,
    commit_block,
    verify_generated_proof,
)


def main() -> None:
    compiled = compile_real_rollup_contracts()
    chain_a = build_chain("chain-a", 1, compiled)
    chain_b = build_chain("chain-b", 2, compiled)

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

    malformed_chain = build_chain("malformed-chain", 1, compiled)
    seed_account(malformed_chain, alice, 100)
    malformed_commit = commit_block(
        malformed_chain,
        0,
        mirrored_txs,
        mirrored_transitions[:1],
    )

    chain_a_flow = commit_and_confirm(chain_a, mirrored_txs, mirrored_transitions)
    chain_b_flow = commit_and_confirm(chain_b, mirrored_txs, mirrored_transitions)

    execution_chain = build_chain("execution-chain", 1, compiled)
    seed_account(execution_chain, alice, 100)
    execution_commit = commit_block(execution_chain, 0, mirrored_txs, mirrored_transitions)
    legal_challenge = execution_challenge(execution_chain, 0, mirrored_txs)

    remote_consistency_chain = build_chain("remote-consistency-chain", 2, compiled)
    seed_account(remote_consistency_chain, alice, 100)
    divergent_txs = [l2_tx(alice, bob, 24, chain_a.chain_id, 1)]
    divergent_transitions = build_transition_list(divergent_txs)
    remote_consistency_commit = commit_block(
        remote_consistency_chain,
        0,
        divergent_txs,
        divergent_transitions,
    )
    remote_consistency_confirm = remote_consistency_chain.cross_rollup.functions.checkCommit().transact(
        {"from": remote_consistency_chain.deployer, "gas": 25_000_000}
    )
    remote_consistency_chain.w3.eth.wait_for_transaction_receipt(remote_consistency_confirm)
    remote_checkpoint_header = get_block_header_bundle(remote_consistency_chain, 1)
    remote_middle_header = get_block_header_bundle(remote_consistency_chain, 2)
    remote_final_header = get_block_header_bundle(remote_consistency_chain, 3)
    remote_final_header_proof = generate_header_hash_proof(remote_final_header["header_bytes"])
    remote_commitment_proof = build_commitment_proof(remote_consistency_chain, 0)

    consistency_chain = build_chain("consistency-chain", 1, compiled)
    seed_account(consistency_chain, alice, 100)
    checkpoint_result = checkpoint_remote_chain(
        consistency_chain,
        remote_chain_id=2,
        checkpoint_index=1,
        remote_block_id=remote_checkpoint_header["block_number"],
        header_bytes=remote_checkpoint_header["header_bytes"],
        state_root=build_state_root_bundle(
            remote_checkpoint_header["state_root"],
            remote_commitment_proof["root"],
        ),
        contract_address=consistency_chain.cross_rollup.address,
    )
    consistency_commit = commit_block(consistency_chain, 0, mirrored_txs, mirrored_transitions)
    consistency_create = create_consistency_challenge(
        consistency_chain,
        index="consistency-0",
        remote_chain_id=2,
        local_block_id=0,
        remote_block_id=remote_final_header["block_number"],
        remote_header_bytes=remote_final_header["header_bytes"],
        remote_state_root=build_state_root_bundle(
            remote_final_header["state_root"],
            remote_commitment_proof["root"],
        ),
        remote_state_proof=encode_state_proof(remote_commitment_proof),
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
        remote_chain_id=2,
        remote_block_id=remote_middle_header["block_number"],
        header_bytes=remote_middle_header["header_bytes"],
        state_root=build_state_root_bundle(
            remote_middle_header["state_root"],
            remote_commitment_proof["root"],
        ),
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
    zk_direct_verify = verify_generated_proof(consistency_chain, remote_final_header_proof)
    zk_direct_verify_bad = verify_generated_proof(
        consistency_chain,
        remote_final_header_proof,
        [
            remote_final_header_proof["input"][0] + 1,
            remote_final_header_proof["input"][1],
            remote_final_header_proof["input"][2],
            remote_final_header_proof["input"][3],
        ],
    )

    after = {
        "chain_a_alice": get_local_state(chain_a, alice),
        "chain_b_bob": get_local_state(chain_b, bob),
    }

    assert before["chain_a_alice"]["value"] == 100
    assert before["chain_a_alice"]["lock"] == 0
    assert before["chain_b_bob"]["value"] == 1
    assert after["chain_a_alice"]["value"] == 75
    assert after["chain_a_alice"]["lock"] == 0
    assert after["chain_b_bob"]["value"] == 26
    assert after["chain_b_bob"]["lock"] == 0
    assert malformed_commit["events"] == []
    assert "Illegal transition list" in malformed_commit["errors"]
    assert execution_commit["events"] == ["1-0-commit success"]
    assert legal_challenge["errors"] == []
    assert legal_challenge["events"] == ["4-1-The Challenged Block is legal"]
    assert consistency_commit["events"] == ["1-0-commit success"]
    assert consistency_create["events"] == ["create success!"]
    assert consistency_question["events"] == ["success! the commit was questioned"]
    assert consistency_response["events"] == ["challenge response submitted"]
    assert consistency_ack["events"] == ["True! In the last 1/2."]
    assert consistency_final["events"] == ["final proof success"]
    assert consistency_final["rollup_events"] == ["3-0-rollback success"]
    assert zk_direct_verify is True
    assert zk_direct_verify_bad is False
    assert consistency_after["alice"]["value"] == 100
    assert consistency_after["alice"]["lock"] == 0
    assert consistency_after["block_len"] == 0

    result = {
        "summary": "real rollup contracts basic flow passed",
        "deployments": {
            "chain_a": {
                "cross_rollup": chain_a.cross_rollup.address,
                "local_state_manager": chain_a.local_state_manager.address,
                "transaction_executor": chain_a.transaction_executor.address,
                "data_types": chain_a.data_types.address,
            },
            "chain_b": {
                "cross_rollup": chain_b.cross_rollup.address,
                "local_state_manager": chain_b.local_state_manager.address,
                "transaction_executor": chain_b.transaction_executor.address,
                "data_types": chain_b.data_types.address,
            },
            "consistency_chain": {
                "cross_rollup": consistency_chain.cross_rollup.address,
                "bcp_manager": consistency_chain.bcp_manager.address,
                "zk_verifier": consistency_chain.zk_verifier.address,
            },
        },
        "before": before,
        "after": after,
        "mirrored_txs": mirrored_txs,
        "mirrored_transitions": mirrored_transitions,
        "malformed_commit": malformed_commit,
        "chain_a_flow": chain_a_flow,
        "chain_b_flow": chain_b_flow,
        "legal_challenge": {
            "commit": execution_commit,
            "challenge": legal_challenge,
        },
        "consistency_challenge": {
            "remote_commit": remote_consistency_commit,
            "checkpoint": checkpoint_result,
            "commit": consistency_commit,
            "create": consistency_create,
            "question": consistency_question,
            "response": consistency_response,
            "ack": consistency_ack,
            "final": consistency_final,
            "after": consistency_after,
            "remote_commitment_hash": remote_commitment_proof["commitment_hash"].hex(),
            "remote_commitment_root": remote_commitment_proof["root"].hex(),
            "local_commitment_hash_before": consistency_local_commitment_before,
            "zk_direct_verify": zk_direct_verify,
            "zk_direct_verify_bad": zk_direct_verify_bad,
        },
        "notes": [
            "This validates deploy, depositCoin, commitBlock, checkCommit, a legal execution challenge, and the full consistency challenge state machine.",
            "The real MPT verifier is now part of the deployed stack; this eth-tester flow simply does not exercise eth_getProof-backed state proofs.",
            "The flow now commits the same cross-chain transaction batch on both chains using the full mirrored transition list.",
            "A malformed commit with an incomplete transition list is rejected with Illegal transition list.",
            "The consistency challenge test uses real RLP headers, a real commitment inclusion proof, and a real Groth16 verifier/proof pair generated off-chain.",
        ],
    }
    print(json.dumps(result, indent=2))


if __name__ == "__main__":
    main()

#!/usr/bin/env python3
"""Adversarial experiments for the local dual-chain rollup harness."""

from __future__ import annotations

import json

from real_rollup_harness import (
    build_commitment_proof,
    build_state_root_bundle,
    build_transition_list,
    build_chain,
    challenge_question,
    challenge_response,
    commit_block,
    compile_real_rollup_contracts,
    confirm_latest,
    create_consistency_challenge,
    encode_state_proof,
    execution_challenge,
    final_consistency_challenge,
    finalize_consistency_timeout,
    get_block_header_bundle,
    get_local_state,
    l2_tx,
    mine_blocks,
    seed_account,
    try_commit_block,
)


def sum_gas(*records: dict) -> int:
    return sum(int(record.get("gas_used", 0)) for record in records if record)


def prepare_consistency_fixture(
    compiled: dict,
    remote_txs: list[tuple],
    remote_transitions: list[tuple],
    local_txs: list[tuple] | None = None,
    local_transitions: list[tuple] | None = None,
) -> dict:
    alice = "alice@chain-a"
    bob = "bob@chain-b"

    mirrored_txs = local_txs or [l2_tx(alice, bob, 25, 1, 1)]
    mirrored_transitions = local_transitions or build_transition_list(mirrored_txs)

    remote_chain = build_chain("remote-consistency-chain", 2, compiled)
    seed_account(remote_chain, alice, 100)
    seed_account(remote_chain, bob, 1)
    remote_commit = commit_block(remote_chain, 0, remote_txs, remote_transitions)
    remote_confirm = confirm_latest(remote_chain)

    checkpoint_header = get_block_header_bundle(remote_chain, 1)
    middle_header = get_block_header_bundle(remote_chain, 2)
    final_header = get_block_header_bundle(remote_chain, 3)
    commitment_proof = build_commitment_proof(remote_chain, 0)

    consistency_chain = build_chain("consistency-chain", 1, compiled)
    seed_account(consistency_chain, alice, 100)

    checkpoint = consistency_chain.bcp_manager.functions.CheckPoint(
        (
            1,
            2,
            1,
            checkpoint_header["block_number"],
            build_state_root_bundle(checkpoint_header["state_root"], commitment_proof["root"]),
            checkpoint_header["hash"],
            consistency_chain.cross_rollup.address,
        )
    ).transact({"from": consistency_chain.deployer, "gas": 25_000_000})
    consistency_chain.w3.eth.wait_for_transaction_receipt(checkpoint)

    local_commit = commit_block(consistency_chain, 0, mirrored_txs, mirrored_transitions)
    create_result = create_consistency_challenge(
        consistency_chain,
        index="consistency-0",
        remote_chain_id=2,
        local_block_id=0,
        remote_block_id=final_header["block_number"],
        remote_header_bytes=final_header["header_bytes"],
        remote_state_root=build_state_root_bundle(final_header["state_root"], commitment_proof["root"]),
        remote_state_proof=encode_state_proof(commitment_proof),
        questioned_rollup=consistency_chain.cross_rollup.address,
    )

    return {
        "alice": alice,
        "bob": bob,
        "questioner": consistency_chain.w3.eth.accounts[1],
        "remote_chain": remote_chain,
        "consistency_chain": consistency_chain,
        "remote_commit": remote_commit,
        "remote_confirm": remote_confirm,
        "local_commit": local_commit,
        "create": create_result,
        "checkpoint_header": checkpoint_header,
        "middle_header": middle_header,
        "final_header": final_header,
        "commitment_proof": commitment_proof,
        "mirrored_txs": mirrored_txs,
        "mirrored_transitions": mirrored_transitions,
    }


def advance_to_final_stage(fixture: dict) -> dict:
    consistency_chain = fixture["consistency_chain"]
    question = challenge_question(
        consistency_chain,
        index="consistency-0",
        ack=False,
        sender=fixture["questioner"],
    )
    response = challenge_response(
        consistency_chain,
        index="consistency-0",
        remote_chain_id=2,
        remote_block_id=fixture["middle_header"]["block_number"],
        header_bytes=fixture["middle_header"]["header_bytes"],
        state_root=build_state_root_bundle(
            fixture["middle_header"]["state_root"],
            fixture["commitment_proof"]["root"],
        ),
    )
    ack = challenge_question(
        consistency_chain,
        index="consistency-0",
        ack=True,
        sender=fixture["questioner"],
    )
    return {"question": question, "response": response, "ack": ack}


def run_censorship_attack_experiment() -> dict:
    config = {"dispute_time": 2, "wait_period": 2, "confirm_period": 2}
    compiled = compile_real_rollup_contracts(config)
    fixture = prepare_consistency_fixture(compiled, remote_txs=[], remote_transitions=[])
    stage = advance_to_final_stage(fixture)
    finalize = final_consistency_challenge(
        fixture["consistency_chain"],
        index="consistency-0",
        begin_header_bytes=fixture["middle_header"]["header_bytes"],
        end_header_bytes=fixture["final_header"]["header_bytes"],
    )
    after = get_local_state(fixture["consistency_chain"], fixture["alice"])
    remote_after = get_local_state(fixture["remote_chain"], fixture["bob"])
    block_len = int(fixture["consistency_chain"].cross_rollup.functions.BlockLen().call())

    assert fixture["create"]["events"] == ["create success!"]
    assert finalize["events"] == ["final proof success"]
    assert finalize["rollup_events"] == ["3-0-rollback success"]
    assert after["value"] == 100
    assert remote_after["value"] == 1
    assert block_len == 0

    return {
        "name": "censorship_attack",
        "window_config": config,
        "attack_shape": "source chain commits user transfer while remote chain commits an empty batch",
        "local_commit": fixture["local_commit"],
        "remote_commit": fixture["remote_commit"],
        "stage": stage,
        "create": fixture["create"],
        "finalize": finalize,
        "after": after,
        "remote_after": remote_after,
        "block_len": block_len,
        "cost": {
            "attacker_gas": sum_gas(fixture["local_commit"], fixture["remote_commit"]),
            "defender_gas": sum_gas(fixture["create"], stage["question"], stage["response"], stage["ack"], finalize),
            "proof_count": 1,
            "challenge_rounds": 3,
        },
        "notes": "Simulates sequencer-side censorship by omitting the mirrored destination-chain transfer and shows the inconsistency is rolled back.",
    }


def run_dos_network_partition_experiment() -> dict:
    config = {"dispute_time": 2, "wait_period": 2, "confirm_period": 2}
    compiled = compile_real_rollup_contracts(config)
    remote_txs = [l2_tx("alice@chain-a", "bob@chain-b", 24, 1, 1)]
    remote_transitions = build_transition_list(remote_txs)
    fixture = prepare_consistency_fixture(compiled, remote_txs=remote_txs, remote_transitions=remote_transitions)

    mining = mine_blocks(fixture["consistency_chain"], config["confirm_period"])
    finalize = finalize_consistency_timeout(fixture["consistency_chain"], "consistency-0")
    after = get_local_state(fixture["consistency_chain"], fixture["alice"])
    block_len = int(fixture["consistency_chain"].cross_rollup.functions.BlockLen().call())

    assert fixture["create"]["events"] == ["create success!"]
    assert finalize["events"] == ["challenge timeout success"]
    assert finalize["rollup_events"] == ["3-0-rollback success"]
    assert after["value"] == 100
    assert block_len == 0

    return {
        "name": "dos_network_partition",
        "window_config": config,
        "create": fixture["create"],
        "mining": mining,
        "finalize": finalize,
        "after": after,
        "block_len": block_len,
        "cost": {
            "attacker_gas": sum_gas(fixture["local_commit"], fixture["remote_commit"]),
            "defender_gas": sum_gas(fixture["create"], finalize),
            "sustained_delay_blocks": mining["mined"],
            "proof_count": 0,
        },
        "notes": "Simulates DoS or network partition by preventing the questioner side from responding before the confirm window expires.",
    }


def run_proof_delay_experiment() -> dict:
    config = {"dispute_time": 2, "wait_period": 2, "confirm_period": 2}
    compiled = compile_real_rollup_contracts(config)
    remote_txs = [l2_tx("alice@chain-a", "bob@chain-b", 24, 1, 1)]
    remote_transitions = build_transition_list(remote_txs)
    fixture = prepare_consistency_fixture(compiled, remote_txs=remote_txs, remote_transitions=remote_transitions)
    stage = advance_to_final_stage(fixture)

    mining = mine_blocks(fixture["consistency_chain"], config["wait_period"])
    final_attempt = final_consistency_challenge(
        fixture["consistency_chain"],
        index="consistency-0",
        begin_header_bytes=fixture["middle_header"]["header_bytes"],
        end_header_bytes=fixture["final_header"]["header_bytes"],
    )
    finalize = finalize_consistency_timeout(fixture["consistency_chain"], "consistency-0")
    after = get_local_state(fixture["consistency_chain"], fixture["alice"])
    block_len = int(fixture["consistency_chain"].cross_rollup.functions.BlockLen().call())

    assert stage["ack"]["events"] == ["True! In the last 1/2."]
    assert final_attempt["events"] == ["fail! challenge step expired"]
    assert finalize["events"] == ["challenge timeout failed"]
    assert finalize["rollup_events"] == []
    assert after["value"] == 75
    assert block_len == 1

    return {
        "name": "proof_delay",
        "window_config": config,
        "stage": stage,
        "mining": mining,
        "final_attempt": final_attempt,
        "finalize": finalize,
        "after": after,
        "block_len": block_len,
        "cost": {
            "attacker_gas": sum_gas(fixture["local_commit"], fixture["remote_commit"]),
            "defender_gas_before_timeout": sum_gas(
                fixture["create"], stage["question"], stage["response"], stage["ack"], final_attempt
            ),
            "sustained_delay_blocks": mining["mined"],
            "proof_count": 1,
        },
        "notes": "Simulates prover delay after the dispute narrows to the final proof step; the proof arrives too late and liveness degrades.",
    }


def run_double_spend_detection_experiment() -> dict:
    compiled = compile_real_rollup_contracts()
    chain = build_chain("double-spend-chain", 1, compiled)
    alice = "alice@chain-a"
    bob = "bob@chain-b"
    charlie = "charlie@chain-b"
    seed_account(chain, alice, 100)

    double_spend_txs = [
        l2_tx(alice, bob, 60, 1, 1),
        l2_tx(alice, charlie, 60, 1, 1),
    ]
    double_spend_transitions = build_transition_list(double_spend_txs)
    attempt = try_commit_block(chain, 0, double_spend_txs, double_spend_transitions)
    after = get_local_state(chain, alice)
    block_len = int(chain.cross_rollup.functions.BlockLen().call())

    assert "Illegal state verify" in attempt["errors"]
    assert after["value"] == 100
    assert after["lock"] == 0
    assert block_len == 0

    return {
        "name": "double_spend_detection",
        "attempt": attempt,
        "after": after,
        "block_len": block_len,
        "cost": {
            "attacker_gas": attempt["gas_used"],
            "defender_gas": 0,
            "proof_count": 0,
        },
        "notes": "Submits two outgoing transfers from the same account whose combined value exceeds the balance; the batch is rejected before state mutation.",
    }


def run_window_sensitivity_experiment() -> dict:
    cases = [
        {
            "name": "tight_window",
            "config": {"dispute_time": 1, "wait_period": 1, "confirm_period": 1},
            "execution_delay_blocks": 1,
            "proof_delay_blocks": 1,
            "expect_execution_closed": True,
            "expect_final_expired": True,
        },
        {
            "name": "relaxed_window",
            "config": {"dispute_time": 3, "wait_period": 3, "confirm_period": 2},
            "execution_delay_blocks": 1,
            "proof_delay_blocks": 1,
            "expect_execution_closed": False,
            "expect_final_expired": False,
        },
    ]

    results = []
    for case in cases:
        compiled = compile_real_rollup_contracts(case["config"])

        alice = "alice@chain-a"
        bob = "bob@chain-b"
        txs = [l2_tx(alice, bob, 25, 1, 1)]
        transitions = build_transition_list(txs)

        execution_chain = build_chain(f"execution-{case['name']}", 1, compiled)
        seed_account(execution_chain, alice, 100)
        commit = commit_block(execution_chain, 0, txs, transitions)
        execution_mining = mine_blocks(execution_chain, case["execution_delay_blocks"])
        execution_result = execution_challenge(execution_chain, 0, txs)

        remote_txs = [l2_tx(alice, bob, 24, 1, 1)]
        remote_transitions = build_transition_list(remote_txs)
        fixture = prepare_consistency_fixture(compiled, remote_txs=remote_txs, remote_transitions=remote_transitions)
        stage = advance_to_final_stage(fixture)
        proof_mining = mine_blocks(fixture["consistency_chain"], case["proof_delay_blocks"])
        final_result = final_consistency_challenge(
            fixture["consistency_chain"],
            index="consistency-0",
            begin_header_bytes=fixture["middle_header"]["header_bytes"],
            end_header_bytes=fixture["final_header"]["header_bytes"],
        )

        execution_closed = "challenge window closed" in execution_result["errors"]
        final_expired = final_result["events"] == ["fail! challenge step expired"]

        assert execution_closed is case["expect_execution_closed"]
        assert final_expired is case["expect_final_expired"]

        results.append(
            {
                "name": case["name"],
                "window_config": case["config"],
                "execution_commit": commit,
                "execution_mining": execution_mining,
                "execution_result": execution_result,
                "execution_closed": execution_closed,
                "consistency_stage": stage,
                "proof_mining": proof_mining,
                "final_result": final_result,
                "final_expired": final_expired,
                "cost": {
                    "execution_path_gas": sum_gas(commit, execution_result),
                    "consistency_path_gas": sum_gas(
                        fixture["create"], stage["question"], stage["response"], stage["ack"], final_result
                    ),
                    "sustained_delay_blocks": execution_mining["mined"] + proof_mining["mined"],
                    "proof_count": 1,
                },
            }
        )

    return {
        "name": "window_sensitivity",
        "cases": results,
        "notes": "Compares how tighter and looser dispute windows affect execution-challenge admissibility and final-proof availability.",
    }


def build_attack_cost_quantification(experiments: list[dict]) -> dict:
    summary = []
    for experiment in experiments:
        cost = experiment.get("cost")
        if cost is not None:
            summary.append({"name": experiment["name"], **cost})
            continue

        if "cases" in experiment:
            for case in experiment["cases"]:
                summary.append({"name": f"{experiment['name']}::{case['name']}", **case["cost"]})

    return {
        "name": "attack_cost_quantification",
        "cases": summary,
        "notes": "Uses gas consumption, proof count, and sustained delay blocks as local-devnet proxies for attacker effort and defender recovery cost.",
    }


def main() -> None:
    experiments = [
        run_censorship_attack_experiment(),
        run_dos_network_partition_experiment(),
        run_proof_delay_experiment(),
        run_double_spend_detection_experiment(),
        run_window_sensitivity_experiment(),
    ]
    experiments.append(build_attack_cost_quantification(experiments))

    print(
        json.dumps(
            {
                "summary": "adversarial experiments passed",
                "experiments": experiments,
            },
            indent=2,
        )
    )


if __name__ == "__main__":
    main()

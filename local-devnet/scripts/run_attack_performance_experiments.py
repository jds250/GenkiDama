#!/usr/bin/env python3
"""Measure throughput and latency degradation under attack conditions."""

from __future__ import annotations

import argparse
import json
import time

from real_rollup_harness import (
    build_commitment_proof,
    build_chain,
    build_state_root_bundle,
    build_transition_list,
    challenge_question,
    challenge_response,
    commit_and_confirm,
    commit_block,
    compile_real_rollup_contracts,
    confirm_latest,
    create_consistency_challenge,
    encode_state_proof,
    execution_challenge,
    finalize_consistency_timeout,
    final_consistency_challenge,
    generate_accounts,
    get_block_header_bundle,
    get_local_state,
    l2_tx,
    mine_blocks,
    seed_account,
)
from run_adversarial_experiments import advance_to_final_stage, prepare_consistency_fixture, sum_gas


def seed_users(chain, senders, receivers, initial_balance):
    for sender in senders:
        seed_account(chain, sender, initial_balance)
    for receiver in receivers:
        seed_account(chain, receiver, 1)


def make_batch(senders, receivers, value):
    txs = []
    for sender, receiver in zip(senders, receivers):
        txs.append(l2_tx(sender, receiver, value, 1, 1))
    return txs, build_transition_list(txs)


def relative_delta(base: float, other: float) -> float:
    if base == 0:
        return 0.0
    return (other - base) / base


def safe_ratio(numerator: float, denominator: float) -> float:
    if denominator == 0:
        return 0.0
    return numerator / denominator


def run_throughput_baseline(compiled: dict, num_batches: int, batch_size: int, transfer_value: int) -> dict:
    chain_a = build_chain("baseline-chain-a", 1, compiled)
    chain_b = build_chain("baseline-chain-b", 2, compiled)

    sender_count = num_batches * batch_size
    senders = generate_accounts("baseline-sender", sender_count)
    receivers = generate_accounts("baseline-receiver", sender_count)
    seed_users(chain_a, senders, receivers, 1000)
    seed_users(chain_b, receivers, senders, 1)

    start = time.perf_counter()
    cursor = 0
    gas_used = 0
    for _ in range(num_batches):
        batch_senders = senders[cursor : cursor + batch_size]
        batch_receivers = receivers[cursor : cursor + batch_size]
        cursor += batch_size

        txs, transitions = make_batch(batch_senders, batch_receivers, transfer_value)
        result_a = commit_and_confirm(chain_a, txs, transitions)
        result_b = commit_and_confirm(chain_b, txs, transitions)
        gas_used += sum_gas(result_a["commit"], result_a["confirm"], result_b["commit"], result_b["confirm"])

    elapsed = time.perf_counter() - start
    attempted_business = num_batches * batch_size
    finalized_business = attempted_business

    return {
        "name": "baseline",
        "elapsed_seconds": elapsed,
        "attempted_business_transactions": attempted_business,
        "finalized_business_transactions": finalized_business,
        "attempted_business_tps": safe_ratio(attempted_business, elapsed),
        "finalized_business_tps": safe_ratio(finalized_business, elapsed),
        "control_plane_batches_per_second": safe_ratio(num_batches, elapsed),
        "gas_used": gas_used,
        "notes": "Optimistic mirrored commit and confirm on both chains with no dispute activity.",
    }


def run_throughput_execution_challenge_load(
    compiled: dict, num_batches: int, batch_size: int, transfer_value: int
) -> dict:
    chain_a = build_chain("exec-challenge-chain-a", 1, compiled)
    chain_b = build_chain("exec-challenge-chain-b", 2, compiled)

    sender_count = num_batches * batch_size
    senders = generate_accounts("challenge-sender", sender_count)
    receivers = generate_accounts("challenge-receiver", sender_count)
    seed_users(chain_a, senders, receivers, 1000)
    seed_users(chain_b, receivers, senders, 1)

    start = time.perf_counter()
    cursor = 0
    gas_used = 0
    legal_challenges = 0
    for _ in range(num_batches):
        batch_senders = senders[cursor : cursor + batch_size]
        batch_receivers = receivers[cursor : cursor + batch_size]
        cursor += batch_size

        txs, transitions = make_batch(batch_senders, batch_receivers, transfer_value)
        commit_a = commit_block(chain_a, 0 if int(chain_a.cross_rollup.functions.BlockLen().call()) == 0 else int(chain_a.cross_rollup.functions.BlockLen().call()), txs, transitions)
        challenge_a = execution_challenge(chain_a, int(chain_a.cross_rollup.functions.BlockLen().call()) - 1, txs)
        confirm_a = confirm_latest(chain_a)

        commit_b = commit_block(chain_b, 0 if int(chain_b.cross_rollup.functions.BlockLen().call()) == 0 else int(chain_b.cross_rollup.functions.BlockLen().call()), txs, transitions)
        challenge_b = execution_challenge(chain_b, int(chain_b.cross_rollup.functions.BlockLen().call()) - 1, txs)
        confirm_b = confirm_latest(chain_b)

        if challenge_a["events"] == ["4-1-The Challenged Block is legal"]:
            legal_challenges += 1
        if challenge_b["events"] == ["4-1-The Challenged Block is legal"]:
            legal_challenges += 1

        gas_used += sum_gas(commit_a, challenge_a, confirm_a, commit_b, challenge_b, confirm_b)

    elapsed = time.perf_counter() - start
    attempted_business = num_batches * batch_size
    finalized_business = attempted_business

    return {
        "name": "frequent_execution_challenge",
        "elapsed_seconds": elapsed,
        "attempted_business_transactions": attempted_business,
        "finalized_business_transactions": finalized_business,
        "attempted_business_tps": safe_ratio(attempted_business, elapsed),
        "finalized_business_tps": safe_ratio(finalized_business, elapsed),
        "control_plane_batches_per_second": safe_ratio(num_batches, elapsed),
        "gas_used": gas_used,
        "legal_challenges_triggered": legal_challenges,
        "notes": "Every batch is subjected to a legal execution challenge before confirmation on both chains.",
    }


def run_throughput_censorship_recovery(compiled: dict, num_batches: int, batch_size: int, transfer_value: int) -> dict:
    config = {"dispute_time": 2, "wait_period": 2, "confirm_period": 2}
    compiled = compile_real_rollup_contracts(config)
    start = time.perf_counter()
    attempted_business = 0
    finalized_business = 0
    recovered_batches = 0
    gas_used = 0

    for batch_id in range(num_batches):
        local_txs = [l2_tx("alice@chain-a", "bob@chain-b", transfer_value, 1, 1) for _ in range(batch_size)]
        local_transitions = build_transition_list(local_txs)

        fixture = prepare_consistency_fixture(
            compiled,
            remote_txs=[],
            remote_transitions=[],
            local_txs=local_txs,
            local_transitions=local_transitions,
        )
        stage = advance_to_final_stage(fixture)
        final_result = final_consistency_challenge(
            fixture["consistency_chain"],
            index="consistency-0",
            begin_header_bytes=fixture["middle_header"]["header_bytes"],
            end_header_bytes=fixture["final_header"]["header_bytes"],
        )

        attempted_business += len(local_txs)
        if final_result["events"] == ["final proof success"] and final_result["rollup_events"] == ["3-0-rollback success"]:
            recovered_batches += 1
        gas_used += sum_gas(
            fixture["local_commit"],
            fixture["remote_commit"],
            fixture["create"],
            stage["question"],
            stage["response"],
            stage["ack"],
            final_result,
        )

        # A censorship scenario should not safely finalize any business transfer.
        assert get_local_state(fixture["consistency_chain"], fixture["alice"])["value"] == 100

    elapsed = time.perf_counter() - start

    return {
        "name": "censorship_recovery_load",
        "elapsed_seconds": elapsed,
        "attempted_business_transactions": attempted_business,
        "finalized_business_transactions": finalized_business,
        "attempted_business_tps": safe_ratio(attempted_business, elapsed),
        "finalized_business_tps": safe_ratio(finalized_business, elapsed),
        "control_plane_batches_per_second": safe_ratio(num_batches, elapsed),
        "gas_used": gas_used,
        "recovered_attack_batches": recovered_batches,
        "notes": "Each batch is attacked by omitting the mirrored destination commit, then recovered through a full consistency challenge.",
    }


def run_latency_baseline(compiled: dict, batch_size: int, transfer_value: int) -> dict:
    chain_a = build_chain("latency-baseline-a", 1, compiled)
    chain_b = build_chain("latency-baseline-b", 2, compiled)

    senders = generate_accounts("latency-base-sender", batch_size)
    receivers = generate_accounts("latency-base-receiver", batch_size)
    seed_users(chain_a, senders, receivers, 1000)
    seed_users(chain_b, receivers, senders, 1)

    txs, transitions = make_batch(senders, receivers, transfer_value)
    before_a = int(chain_a.w3.eth.block_number)
    before_b = int(chain_b.w3.eth.block_number)
    start = time.perf_counter()
    result_a = commit_and_confirm(chain_a, txs, transitions)
    result_b = commit_and_confirm(chain_b, txs, transitions)
    elapsed = time.perf_counter() - start

    return {
        "name": "baseline_commit_confirm",
        "elapsed_seconds": elapsed,
        "block_delta_chain_a": int(chain_a.w3.eth.block_number) - before_a,
        "block_delta_chain_b": int(chain_b.w3.eth.block_number) - before_b,
        "block_delta_total": (int(chain_a.w3.eth.block_number) - before_a)
        + (int(chain_b.w3.eth.block_number) - before_b),
        "gas_used": sum_gas(result_a["commit"], result_a["confirm"], result_b["commit"], result_b["confirm"]),
        "final_state": {
            "chain_a_sender_0": get_local_state(chain_a, senders[0]),
            "chain_b_receiver_0": get_local_state(chain_b, receivers[0]),
        },
    }


def run_latency_execution_challenge(compiled: dict, batch_size: int, transfer_value: int) -> dict:
    chain_a = build_chain("latency-exec-challenge-a", 1, compiled)
    chain_b = build_chain("latency-exec-challenge-b", 2, compiled)
    senders = generate_accounts("latency-exec-sender", batch_size)
    receivers = generate_accounts("latency-exec-receiver", batch_size)
    seed_users(chain_a, senders, receivers, 1000)
    seed_users(chain_b, receivers, senders, 1)

    txs, transitions = make_batch(senders, receivers, transfer_value)
    before_a = int(chain_a.w3.eth.block_number)
    before_b = int(chain_b.w3.eth.block_number)
    start = time.perf_counter()
    commit_a = commit_block(chain_a, 0, txs, transitions)
    challenge_a = execution_challenge(chain_a, 0, txs)
    confirm_a = confirm_latest(chain_a)
    commit_b = commit_block(chain_b, 0, txs, transitions)
    challenge_b = execution_challenge(chain_b, 0, txs)
    confirm_b = confirm_latest(chain_b)
    elapsed = time.perf_counter() - start

    return {
        "name": "execution_challenge_path",
        "elapsed_seconds": elapsed,
        "block_delta_total": (int(chain_a.w3.eth.block_number) - before_a)
        + (int(chain_b.w3.eth.block_number) - before_b),
        "gas_used": sum_gas(commit_a, challenge_a, confirm_a, commit_b, challenge_b, confirm_b),
        "challenge_events": {
            "chain_a": challenge_a["events"],
            "chain_b": challenge_b["events"],
        },
        "final_state": {
            "chain_a_sender_0": get_local_state(chain_a, senders[0]),
            "chain_b_receiver_0": get_local_state(chain_b, receivers[0]),
        },
    }


def run_latency_consistency_recovery(compiled: dict, batch_size: int, transfer_value: int) -> dict:
    config = {"dispute_time": 2, "wait_period": 2, "confirm_period": 2}
    compiled = compile_real_rollup_contracts(config)
    alice = "alice@chain-a"
    bob = "bob@chain-b"
    remote_chain = build_chain("latency-consistency-remote", 2, compiled)
    consistency_chain = build_chain("latency-consistency-local", 1, compiled)
    seed_account(remote_chain, alice, 100)
    seed_account(remote_chain, bob, 1)
    seed_account(consistency_chain, alice, 100)
    questioner = consistency_chain.w3.eth.accounts[1]

    local_txs = [l2_tx(alice, bob, transfer_value, 1, 1) for _ in range(batch_size)]
    local_transitions = build_transition_list(local_txs)
    before_total = int(remote_chain.w3.eth.block_number) + int(consistency_chain.w3.eth.block_number)
    start = time.perf_counter()
    remote_commit = commit_block(remote_chain, 0, [], [])
    remote_confirm = confirm_latest(remote_chain)
    checkpoint_header = get_block_header_bundle(remote_chain, 1)
    middle_header = get_block_header_bundle(remote_chain, 2)
    final_header = get_block_header_bundle(remote_chain, 3)
    commitment_proof = build_commitment_proof(remote_chain, 0)
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
    checkpoint_receipt = consistency_chain.w3.eth.wait_for_transaction_receipt(checkpoint)
    local_commit = commit_block(consistency_chain, 0, local_txs, local_transitions)
    create = create_consistency_challenge(
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
    question = challenge_question(consistency_chain, index="consistency-0", ack=False, sender=questioner)
    response = challenge_response(
        consistency_chain,
        index="consistency-0",
        remote_chain_id=2,
        remote_block_id=middle_header["block_number"],
        header_bytes=middle_header["header_bytes"],
        state_root=build_state_root_bundle(middle_header["state_root"], commitment_proof["root"]),
    )
    ack = challenge_question(consistency_chain, index="consistency-0", ack=True, sender=questioner)
    final_result = final_consistency_challenge(
        consistency_chain,
        index="consistency-0",
        begin_header_bytes=middle_header["header_bytes"],
        end_header_bytes=final_header["header_bytes"],
    )
    elapsed = time.perf_counter() - start

    return {
        "name": "consistency_recovery_path",
        "elapsed_seconds": elapsed,
        "block_delta_total": int(remote_chain.w3.eth.block_number) + int(consistency_chain.w3.eth.block_number) - before_total,
        "gas_used": sum_gas(
            remote_commit,
            remote_confirm,
            {"gas_used": int(checkpoint_receipt["gasUsed"])},
            local_commit,
            create,
            question,
            response,
            ack,
            final_result,
        ),
        "rollback_events": final_result["rollup_events"],
        "final_state": get_local_state(consistency_chain, alice),
    }


def run_latency_proof_delay_timeout(compiled: dict, batch_size: int, transfer_value: int) -> dict:
    config = {"dispute_time": 2, "wait_period": 2, "confirm_period": 2}
    compiled = compile_real_rollup_contracts(config)
    remote_txs = [l2_tx("alice@chain-a", "bob@chain-b", transfer_value, 1, 1)]
    remote_transitions = build_transition_list(remote_txs)
    alice = "alice@chain-a"
    bob = "bob@chain-b"
    local_txs = [l2_tx(alice, bob, transfer_value + 1, 1, 1) for _ in range(batch_size)]
    local_transitions = build_transition_list(local_txs)
    remote_chain = build_chain("latency-proof-remote", 2, compiled)
    consistency_chain = build_chain("latency-proof-local", 1, compiled)
    seed_account(remote_chain, alice, 100)
    seed_account(remote_chain, bob, 1)
    seed_account(consistency_chain, alice, 100)
    questioner = consistency_chain.w3.eth.accounts[1]
    before_total = int(remote_chain.w3.eth.block_number) + int(consistency_chain.w3.eth.block_number)
    start = time.perf_counter()
    remote_commit = commit_block(remote_chain, 0, remote_txs, remote_transitions)
    remote_confirm = confirm_latest(remote_chain)
    checkpoint_header = get_block_header_bundle(remote_chain, 1)
    middle_header = get_block_header_bundle(remote_chain, 2)
    final_header = get_block_header_bundle(remote_chain, 3)
    commitment_proof = build_commitment_proof(remote_chain, 0)
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
    checkpoint_receipt = consistency_chain.w3.eth.wait_for_transaction_receipt(checkpoint)
    local_commit = commit_block(consistency_chain, 0, local_txs, local_transitions)
    create = create_consistency_challenge(
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
    question = challenge_question(consistency_chain, index="consistency-0", ack=False, sender=questioner)
    response = challenge_response(
        consistency_chain,
        index="consistency-0",
        remote_chain_id=2,
        remote_block_id=middle_header["block_number"],
        header_bytes=middle_header["header_bytes"],
        state_root=build_state_root_bundle(middle_header["state_root"], commitment_proof["root"]),
    )
    ack = challenge_question(consistency_chain, index="consistency-0", ack=True, sender=questioner)
    mining = mine_blocks(consistency_chain, config["wait_period"])
    final_attempt = final_consistency_challenge(
        consistency_chain,
        index="consistency-0",
        begin_header_bytes=middle_header["header_bytes"],
        end_header_bytes=final_header["header_bytes"],
    )
    timeout_finalize = finalize_consistency_timeout(consistency_chain, "consistency-0")
    elapsed = time.perf_counter() - start

    return {
        "name": "proof_delay_timeout_path",
        "elapsed_seconds": elapsed,
        "block_delta_total": int(remote_chain.w3.eth.block_number) + int(consistency_chain.w3.eth.block_number) - before_total,
        "gas_used": sum_gas(
            remote_commit,
            remote_confirm,
            {"gas_used": int(checkpoint_receipt["gasUsed"])},
            local_commit,
            create,
            question,
            response,
            ack,
            final_attempt,
            timeout_finalize,
        ),
        "delay_blocks": mining["mined"],
        "final_attempt_events": final_attempt["events"],
        "timeout_finalize_status": timeout_finalize["status"],
        "timeout_finalize_events": timeout_finalize["events"],
        "timeout_rollup_events": timeout_finalize["rollup_events"],
        "final_state": get_local_state(consistency_chain, alice),
    }


def run_throughput_experiments(num_batches: int, batch_size: int, transfer_value: int) -> dict:
    compiled = compile_real_rollup_contracts()

    baseline = run_throughput_baseline(compiled, num_batches, batch_size, transfer_value)
    execution_load = run_throughput_execution_challenge_load(compiled, num_batches, batch_size, transfer_value)
    censorship_load = run_throughput_censorship_recovery(compiled, num_batches, batch_size, transfer_value)

    scenarios = [baseline, execution_load, censorship_load]
    for scenario in scenarios[1:]:
        scenario["degradation_vs_baseline"] = {
            "attempted_tps_delta_ratio": relative_delta(
                baseline["attempted_business_tps"], scenario["attempted_business_tps"]
            ),
            "finalized_tps_delta_ratio": relative_delta(
                baseline["finalized_business_tps"], scenario["finalized_business_tps"]
            ),
            "gas_delta_ratio": relative_delta(baseline["gas_used"], scenario["gas_used"]),
        }

    return {
        "summary": "attack throughput degradation experiments completed",
        "parameters": {
            "num_batches": num_batches,
            "batch_size": batch_size,
            "transfer_value": transfer_value,
        },
        "scenarios": scenarios,
    }


def run_latency_experiments(batch_size: int, transfer_value: int) -> dict:
    compiled = compile_real_rollup_contracts()

    baseline = run_latency_baseline(compiled, batch_size, transfer_value)
    execution_path = run_latency_execution_challenge(compiled, batch_size, transfer_value)
    consistency_path = run_latency_consistency_recovery(compiled, batch_size, transfer_value)
    proof_delay_path = run_latency_proof_delay_timeout(compiled, batch_size, transfer_value)

    scenarios = [baseline, execution_path, consistency_path, proof_delay_path]
    for scenario in scenarios[1:]:
        scenario["degradation_vs_baseline"] = {
            "latency_delta_ratio": relative_delta(baseline["elapsed_seconds"], scenario["elapsed_seconds"]),
            "gas_delta_ratio": relative_delta(baseline["gas_used"], scenario["gas_used"]),
            "block_delta_increase": scenario["block_delta_total"] - baseline["block_delta_total"],
        }

    return {
        "summary": "attack latency degradation experiments completed",
        "parameters": {
            "batch_size": batch_size,
            "transfer_value": transfer_value,
        },
        "scenarios": scenarios,
    }


def main() -> None:
    parser = argparse.ArgumentParser()
    parser.add_argument("--num-batches", type=int, default=3)
    parser.add_argument("--batch-size", type=int, default=10)
    parser.add_argument("--transfer-value", type=int, default=1)
    args = parser.parse_args()

    throughput = run_throughput_experiments(args.num_batches, args.batch_size, args.transfer_value)
    latency = run_latency_experiments(args.batch_size, args.transfer_value)

    print(
        json.dumps(
            {
                "summary": "attack performance degradation experiments completed",
                "throughput": throughput,
                "latency": latency,
            },
            indent=2,
        )
    )


if __name__ == "__main__":
    main()

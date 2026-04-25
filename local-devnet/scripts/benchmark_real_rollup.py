#!/usr/bin/env python3
"""Benchmark local throughput for the real rollup commit path."""

from __future__ import annotations

import json
import time
import argparse

from real_rollup_harness import (
    build_transition_list,
    build_chain,
    commit_and_confirm,
    compile_real_rollup_contracts,
    generate_accounts,
    l2_tx,
    seed_account,
)


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


def benchmark(num_batches: int = 20, batch_size: int = 20, transfer_value: int = 1) -> dict:
    compiled = compile_real_rollup_contracts()
    chain_a = build_chain("chain-a", 1, compiled)
    chain_b = build_chain("chain-b", 2, compiled)

    sender_count = num_batches * batch_size
    sender_accounts = generate_accounts("sender", sender_count)
    receiver_accounts = generate_accounts("receiver", sender_count)

    seed_users(chain_a, sender_accounts, receiver_accounts, 1000)
    seed_users(chain_b, receiver_accounts, sender_accounts, 1)

    start = time.perf_counter()

    cursor = 0
    for _ in range(num_batches):
        batch_senders = sender_accounts[cursor : cursor + batch_size]
        batch_receivers = receiver_accounts[cursor : cursor + batch_size]
        cursor += batch_size

        txs, transitions = make_batch(batch_senders, batch_receivers, transfer_value)
        commit_and_confirm(chain_a, txs, transitions)
        commit_and_confirm(chain_b, txs, transitions)

    elapsed = time.perf_counter() - start
    total_business_transactions = num_batches * batch_size
    total_chain_side_transactions = total_business_transactions * 2
    total_commit_transactions = num_batches * 4

    return {
        "summary": "real rollup throughput benchmark completed",
        "parameters": {
            "num_batches": num_batches,
            "batch_size": batch_size,
            "transfer_value": transfer_value,
        },
        "results": {
            "elapsed_seconds": elapsed,
            "total_business_transactions": total_business_transactions,
            "total_chain_side_transactions": total_chain_side_transactions,
            "total_commit_related_transactions": total_commit_transactions,
            "business_transactions_per_second": total_business_transactions / elapsed,
            "chain_side_transactions_per_second": total_chain_side_transactions / elapsed,
            "chain_transactions_per_second": total_commit_transactions / elapsed,
        },
        "notes": [
            "Each batch commits the same mirrored cross-chain transaction list on both local chains, then confirms on both chains.",
            "This is local devnet throughput, useful for regression comparison rather than production sizing.",
        ],
    }


if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    parser.add_argument("--num-batches", type=int, default=10)
    parser.add_argument("--batch-size", type=int, default=10)
    parser.add_argument("--transfer-value", type=int, default=1)
    args = parser.parse_args()
    print(
        json.dumps(
            benchmark(
                num_batches=args.num_batches,
                batch_size=args.batch_size,
                transfer_value=args.transfer_value,
            ),
            indent=2,
        )
    )

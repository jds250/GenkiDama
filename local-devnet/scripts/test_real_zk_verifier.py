#!/usr/bin/env python3
"""Sanity-check the real Groth16 verifier used by the local devnet."""

from __future__ import annotations

import json

from real_rollup_harness import (
    build_chain,
    compile_real_rollup_contracts,
    generate_header_hash_proof,
    get_block_header_bundle,
    verify_generated_proof,
)


def main() -> None:
    compiled = compile_real_rollup_contracts()
    chain = build_chain("zk-test", 1, compiled)

    latest_block_number = int(chain.w3.eth.block_number)
    header_bundle = get_block_header_bundle(chain, latest_block_number)
    proof_bundle = generate_header_hash_proof(header_bundle["header_bytes"])

    verified = verify_generated_proof(chain, proof_bundle)
    verified_with_bad_input = verify_generated_proof(
        chain,
        proof_bundle,
        [
            proof_bundle["input"][0] + 1,
            proof_bundle["input"][1],
            proof_bundle["input"][2],
            proof_bundle["input"][3],
        ],
    )

    assert verified is True
    assert verified_with_bad_input is False

    print(
        json.dumps(
            {
                "summary": "real zk verifier sanity check passed",
                "latest_block_number": latest_block_number,
                "header_hash": header_bundle["hash"].hex(),
                "verified": verified,
                "verified_with_bad_input": verified_with_bad_input,
            },
            indent=2,
        )
    )


if __name__ == "__main__":
    main()

# Rollup Federation

Rollup Federation is a prototype cross-chain optimistic rollup system. The
repository contains:

- Solidity contracts for rollup commitment, local state management, execution
  challenge, and block consistency proof (BCP)
- Go services for challenge orchestration, proof generation, benchmarking, and
  local-network experiments
- local development tooling for spinning up deterministic EVM test networks and
  reproducing end-to-end protocol flows

The main implementation lives under
`Rollup-Offchain/`, while `local-devnet/` provides reproducible local
environments and scripts for integration testing.

## Core Components

- `Rollup-Offchain/contract`
  Solidity contracts for the protocol core, including `CrossRollup`,
  `LocalStateManager`, `TransactionExecutor`, `BCPManager`, and verifier
  interfaces.
- `Rollup-Offchain/challenger`
  Go implementation of challenger / validator-side logic, including execution
  challenge and consistency challenge helpers.
- `Rollup-Offchain/zk`
  Groth16-related tooling and proof helpers used by the consistency challenge
  path.
- `Rollup-Offchain/cmd`
  Command-line entry points for local experiments and debugging:
  `localnetexp`, `localnetbench`, `localnetattackperf`, `debugmpt`,
  `headerdump`, and `zktool`.
- `local-devnet`
  Local chain harness, deployment scripts, and small helper contracts used for
  development and integration testing.
- `EventWatcher`, `MosWatcher`
  Legacy watcher prototypes kept for reference.

## Repository Layout

```text
.
├── Rollup-Offchain/         # Main protocol implementation
│   ├── challenger/          # Challenge / validator logic
│   ├── cmd/                 # Local experiment and debug entry points
│   ├── contract/            # Solidity contracts and Go bindings
│   ├── executor/            # Offchain execution helpers
│   ├── localnet/            # Local network config loading
│   └── zk/                  # ZK proof helpers
├── local-devnet/            # Local devnet generation and deployment tooling
│   ├── contracts/           # Auxiliary contracts for local testing
│   ├── multinode-geth/      # Deterministic 16-node geth devnet assets
│   └── scripts/             # Bootstrap, deployment, and experiment scripts
├── EventWatcher/            # Legacy Ethereum-side watcher prototype
└── MosWatcher/              # Legacy second-chain watcher prototype
```

## Prerequisites

Recommended local tooling:

- Go `1.20+`
- Python `3.9+`
- `solc`
- Docker (optional, for the multinode geth devnet)

Create the Python environment used by the local harness:

```bash
./local-devnet/scripts/bootstrap_local_devnet.sh
source .venv/bin/activate
```

## Quick Start

### 1. Run the basic local flow

```bash
source .venv/bin/activate
python local-devnet/scripts/run_real_rollup_flow.py
```

This validates the core mirrored commit / confirm flow together with the
challenge path on two local development chains.

### 2. Run the local adversarial suite

```bash
source .venv/bin/activate
python local-devnet/scripts/run_adversarial_experiments.py
```

### 3. Run the 16-node geth devnet

Generate the deterministic topology:

```bash
source .venv/bin/activate
python local-devnet/scripts/generate_multinode_geth_devnet.py --validators-per-chain 1
```

Start the local network:

```bash
cd local-devnet/multinode-geth
docker compose up -d
```

Deploy the protocol contracts:

```bash
source ../../.venv/bin/activate
python ../scripts/deploy_multinode_rollup_stack.py \
  --chain-a-rpc http://127.0.0.1:8545 \
  --chain-b-rpc http://127.0.0.1:8645
```

Then run the Go-based experiment entry points from `Rollup-Offchain/`, for
example:

```bash
cd ../../Rollup-Offchain
go run ./cmd/localnetexp
go run ./cmd/localnetbench --num-batches 2 --batch-size 20
go run ./cmd/localnetattackperf --scenario execution_challenge --num-batches 2 --batch-size 20
```

## Notes

- Generated experiment outputs under `local-devnet/results/` and local runtime
  state such as `local-pids.json` are intentionally not meant for version
  control.
- Several directories are retained for historical or comparative experiments.
  If you only want the main implementation, start from `Rollup-Offchain/` and
  `local-devnet/`.

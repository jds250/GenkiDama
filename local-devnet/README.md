# Local Devnet

This directory contains the local development and experiment harness for the
Rollup Federation prototype.

It supports two styles of local testing:

- lightweight local-chain scripts for quick end-to-end validation
- a deterministic multinode geth devnet for more realistic integration and
  performance experiments

## Contents

- `contracts/`
  Auxiliary contracts used only by local testing utilities.
- `scripts/`
  Bootstrap, deployment, benchmark, and experiment runners.
- `multinode-geth/`
  Generated assets and compose files for the 16-node private geth network.

## Bootstrap

From the repository root:

```bash
./local-devnet/scripts/bootstrap_local_devnet.sh
source .venv/bin/activate
```

The bootstrap script creates a local Python environment, installs the Python
dependencies declared in `local-devnet/requirements.txt`, and ensures a usable
`solc` is available.

## Common Workflows

### Quick end-to-end contract flow

```bash
source .venv/bin/activate
python local-devnet/scripts/run_real_rollup_flow.py
```

### Adversarial / security regression

```bash
source .venv/bin/activate
python local-devnet/scripts/run_adversarial_experiments.py
```

### Local throughput benchmark

```bash
source .venv/bin/activate
python local-devnet/scripts/benchmark_real_rollup.py --num-batches 10 --batch-size 10
```

## Multinode Geth Devnet

Generate the deterministic `2 chains x 8 nodes` topology:

```bash
source .venv/bin/activate
python local-devnet/scripts/generate_multinode_geth_devnet.py --validators-per-chain 1
```

Start the network:

```bash
cd local-devnet/multinode-geth
docker compose up -d
```

Deploy the rollup stack:

```bash
source ../../.venv/bin/activate
python ../scripts/deploy_multinode_rollup_stack.py \
  --chain-a-rpc http://127.0.0.1:8545 \
  --chain-b-rpc http://127.0.0.1:8645
```

Further details are documented in `local-devnet/multinode-geth/README.md`.

## Version Control Notes

The following are local-only outputs and should not be committed:

- generated experiment results under `local-devnet/results/`
- `multinode-geth/local-pids.json`
- generated deployment snapshots such as `multinode-geth/deployments/latest.json`
- runtime state under `multinode-geth/runtime/`

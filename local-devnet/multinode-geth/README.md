# Multinode Geth Devnet

This directory contains a deterministic private geth network used by the local
integration and performance experiments.

## Topology

- `chain-a`: network id `1001`, chain id `1001`
- `chain-b`: network id `1002`, chain id `1002`
- `8` nodes per chain
- default validator layout: `1` Clique validator per chain
- external RPC endpoints:
  - `chain-a-node-00`: `http://127.0.0.1:8545` / `ws://127.0.0.1:8546`
  - `chain-a-node-01`: `http://127.0.0.1:8555` / `ws://127.0.0.1:8556`
  - `chain-b-node-00`: `http://127.0.0.1:8645` / `ws://127.0.0.1:8646`
  - `chain-b-node-01`: `http://127.0.0.1:8655` / `ws://127.0.0.1:8656`

## Usage

Generate or refresh the network artifacts from the repository root:

```bash
source .venv/bin/activate
python local-devnet/scripts/generate_multinode_geth_devnet.py --validators-per-chain 1
```

Start the network:

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

After deployment, point the Go offchain services at the generated deployment
JSON and run the benchmark or challenge entry points in `Rollup-Offchain/cmd/`.

## Generated Files

This directory is intended to be generated locally. The following should remain
out of version control:

- `chain-a/`
- `chain-b/`
- `docker-compose.yml`
- `network-summary.json`
- `deployments/`
- `local-pids.json`
- `runtime/`

The stable, source-controlled parts are the README plus helper scripts under
`scripts/`.

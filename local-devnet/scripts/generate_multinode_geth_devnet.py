#!/usr/bin/env python3
"""Generate a reproducible 2-chain x 8-node private geth devnet layout."""

from __future__ import annotations

import argparse
import json
import shutil
from dataclasses import dataclass
from pathlib import Path

from eth_account import Account
from eth_keys import keys
from web3 import Web3


ROOT = Path(__file__).resolve().parents[2]
OUT_ROOT = ROOT / "local-devnet" / "multinode-geth"
START_SCRIPT_REL = "./scripts/start-node.sh"
NODE_COUNT = 8
DEFAULT_VALIDATOR_COUNT = 4
PASSWORD = "rollup-devnet"
GETH_IMAGE = "ethereum/client-go:v1.13.15"


@dataclass(frozen=True)
class ChainLayout:
    slug: str
    chain_id: int
    network_id: int
    rpc_ports: tuple[int, int]
    ws_ports: tuple[int, int]


@dataclass(frozen=True)
class NodeSpec:
    chain: ChainLayout
    index: int
    service_name: str
    node_dir: Path
    nodekey_hex: str
    account_key_hex: str
    account_address: str
    public_key_hex: str
    is_validator: bool


CHAINS = (
    ChainLayout("chain-a", 1001, 1001, (8545, 8555), (8546, 8556)),
    ChainLayout("chain-b", 1002, 1002, (8645, 8655), (8646, 8656)),
)


def deterministic_key(label: str) -> bytes:
    return bytes(Web3.keccak(text=f"rollup-federation::{label}"))


def build_node_spec(chain: ChainLayout, index: int, validator_count: int) -> NodeSpec:
    service_name = f"{chain.slug}-node-{index:02d}"
    nodekey = deterministic_key(f"{chain.slug}-nodekey-{index}")
    account_key = deterministic_key(f"{chain.slug}-account-{index}")
    account = Account.from_key(account_key)
    public_key_hex = keys.PrivateKey(nodekey).public_key.to_bytes().hex()
    return NodeSpec(
        chain=chain,
        index=index,
        service_name=service_name,
        node_dir=OUT_ROOT / chain.slug / "nodes" / f"node-{index:02d}",
        nodekey_hex=nodekey.hex(),
        account_key_hex=account_key.hex(),
        account_address=account.address,
        public_key_hex=public_key_hex,
        is_validator=index < validator_count,
    )


def clique_extradata(validators: list[str]) -> str:
    vanity = "00" * 32
    signer_bytes = "".join(address.removeprefix("0x").lower() for address in validators)
    signature = "00" * 65
    return "0x" + vanity + signer_bytes + signature


def build_genesis(chain: ChainLayout, nodes: list[NodeSpec], validator_count: int) -> dict:
    alloc = {
        node.account_address: {"balance": "0x3635C9ADC5DEA00000"}  # 1000 ETH
        for node in nodes
    }
    validators = [node.account_address for node in nodes[:validator_count]]
    return {
        "config": {
            "chainId": chain.chain_id,
            "homesteadBlock": 0,
            "eip150Block": 0,
            "eip155Block": 0,
            "eip158Block": 0,
            "byzantiumBlock": 0,
            "constantinopleBlock": 0,
            "petersburgBlock": 0,
            "istanbulBlock": 0,
            "muirGlacierBlock": 0,
            "berlinBlock": 0,
            "londonBlock": 0,
            "arrowGlacierBlock": 0,
            "grayGlacierBlock": 0,
            "clique": {
                "period": 1,
                "epoch": 30000,
            },
        },
        "difficulty": "0x1",
        "gasLimit": "0x1c9c380",
        "extradata": clique_extradata(validators),
        "alloc": alloc,
    }


def build_static_nodes(nodes: list[NodeSpec]) -> list[str]:
    return [f"enode://{node.public_key_hex}@{node.service_name}:30303" for node in nodes]


def write_node_artifacts(node: NodeSpec, static_nodes: list[str]) -> dict:
    node.node_dir.mkdir(parents=True, exist_ok=True)
    (node.node_dir / "nodekey").write_text(node.nodekey_hex + "\n")
    (node.node_dir / "account.key").write_text(node.account_key_hex + "\n")
    (node.node_dir / "password.txt").write_text(PASSWORD + "\n")
    (node.node_dir / "static-nodes.json").write_text(json.dumps(static_nodes, indent=2) + "\n")
    static_nodes_toml = ",\n".join(f'  "{enode}"' for enode in static_nodes if enode)
    config_toml = f"""[Node.P2P]
StaticNodes = [
{static_nodes_toml}
]
TrustedNodes = [
{static_nodes_toml}
]
NoDiscovery = true
"""
    (node.node_dir / "config.toml").write_text(config_toml)

    metadata = {
        "service_name": node.service_name,
        "chain": node.chain.slug,
        "node_index": node.index,
        "is_validator": node.is_validator,
        "account_address": node.account_address,
        "p2p_port": 30303,
        "http_port": 8545,
        "ws_port": 8546,
        "authrpc_port": 8551,
        "enode": f"enode://{node.public_key_hex}@{node.service_name}:30303",
    }
    (node.node_dir / "metadata.json").write_text(json.dumps(metadata, indent=2) + "\n")
    return metadata


def compose_service_block(node: NodeSpec, host_http_port: int | None, host_ws_port: int | None) -> str:
    ports = []
    if host_http_port is not None:
        ports.append(f'      - "{host_http_port}:8545"')
    if host_ws_port is not None:
        ports.append(f'      - "{host_ws_port}:8546"')
    ports_block = "    ports: []" if not ports else "    ports:\n" + "\n".join(ports)

    extra_args = "--syncmode full"
    return f"""  {node.service_name}:
    image: {GETH_IMAGE}
    restart: unless-stopped
    networks:
      - {node.chain.slug}
    hostname: {node.service_name}
    entrypoint: ["/bin/sh", "/scripts/start-node.sh"]
    volumes:
      - ./{node.chain.slug}/nodes/node-{node.index:02d}:/data
      - ./{node.chain.slug}/genesis.json:/config/genesis.json:ro
      - {START_SCRIPT_REL}:/scripts/start-node.sh:ro
    environment:
      DATADIR: /data
      GENESIS_PATH: /config/genesis.json
      NETWORK_ID: "{node.chain.network_id}"
      IS_VALIDATOR: "{str(node.is_validator).lower()}"
      UNLOCK_ADDRESS: "{node.account_address}"
      EXTRA_GETH_ARGS: "{extra_args}"
{ports_block}
"""


def build_compose(chain_nodes: dict[str, list[NodeSpec]]) -> str:
    lines = ["services:"]
    for chain in CHAINS:
        nodes = chain_nodes[chain.slug]
        for index, node in enumerate(nodes):
            host_http = chain.rpc_ports[index] if index < len(chain.rpc_ports) else None
            host_ws = chain.ws_ports[index] if index < len(chain.ws_ports) else None
            lines.append(compose_service_block(node, host_http, host_ws))

    lines.append("networks:")
    for chain in CHAINS:
        lines.append(f"  {chain.slug}:")
        lines.append("    driver: bridge")
    return "\n".join(lines).replace("\n\n", "\n")


def write_chain_layout(chain: ChainLayout, validator_count: int) -> dict:
    nodes = [build_node_spec(chain, index, validator_count) for index in range(NODE_COUNT)]
    static_nodes = build_static_nodes(nodes)
    chain_dir = OUT_ROOT / chain.slug
    chain_dir.mkdir(parents=True, exist_ok=True)
    (chain_dir / "genesis.json").write_text(
        json.dumps(build_genesis(chain, nodes, validator_count), indent=2) + "\n"
    )
    metadata = [write_node_artifacts(node, static_nodes) for node in nodes]

    return {
        "chain": chain,
        "nodes": nodes,
        "metadata": metadata,
    }


def write_top_level_readme(summary: dict, validator_count: int) -> None:
    readme = f"""# Multinode Geth Devnet

This directory contains a deterministic `2 chains x 8 geth nodes` private
network for the rollup-federation experiments.

## Topology

- `chain-a`: network id `{CHAINS[0].network_id}`, chain id `{CHAINS[0].chain_id}`
- `chain-b`: network id `{CHAINS[1].network_id}`, chain id `{CHAINS[1].chain_id}`
- `8` nodes per chain
- `{validator_count}` Clique validators per chain
- external RPC endpoints:
  - `chain-a-node-00`: `http://127.0.0.1:{CHAINS[0].rpc_ports[0]}` / `ws://127.0.0.1:{CHAINS[0].ws_ports[0]}`
  - `chain-a-node-01`: `http://127.0.0.1:{CHAINS[0].rpc_ports[1]}` / `ws://127.0.0.1:{CHAINS[0].ws_ports[1]}`
  - `chain-b-node-00`: `http://127.0.0.1:{CHAINS[1].rpc_ports[0]}` / `ws://127.0.0.1:{CHAINS[1].ws_ports[0]}`
  - `chain-b-node-01`: `http://127.0.0.1:{CHAINS[1].rpc_ports[1]}` / `ws://127.0.0.1:{CHAINS[1].ws_ports[1]}`

## Usage

1. Generate or refresh the devnet artifacts:

```bash
source .venv/bin/activate
python local-devnet/scripts/generate_multinode_geth_devnet.py --validators-per-chain {validator_count}
```

2. Start the network after Docker is installed:

```bash
cd local-devnet/multinode-geth
docker compose up -d
```

3. Deploy the real rollup stack to the two chains:

```bash
source ../../.venv/bin/activate
python ../scripts/deploy_multinode_rollup_stack.py \\
  --chain-a-rpc http://127.0.0.1:{CHAINS[0].rpc_ports[0]} \\
  --chain-b-rpc http://127.0.0.1:{CHAINS[1].rpc_ports[0]}
```

4. Point the Go offchain services at the resulting deployment JSON and run the
throughput / challenge experiments on top of this network.

## Generated Outputs

- `docker-compose.yml`
- `network-summary.json`
- chain-specific genesis and node key material under `chain-a/` and `chain-b/`

## Validation

This repository validates the generated topology and deployment tooling for the
selected validator layout.
"""
    (OUT_ROOT / "README.md").write_text(readme)


def main() -> None:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument(
        "--clean",
        action="store_true",
        help="remove previously generated chain directories before regenerating",
    )
    parser.add_argument(
        "--validators-per-chain",
        type=int,
        default=DEFAULT_VALIDATOR_COUNT,
        help="number of Clique validators per chain; use 1 for deterministic throughput experiments",
    )
    args = parser.parse_args()
    if args.validators_per_chain < 1 or args.validators_per_chain > NODE_COUNT:
        raise ValueError(
            f"--validators-per-chain must be between 1 and {NODE_COUNT}, got {args.validators_per_chain}"
        )

    OUT_ROOT.mkdir(parents=True, exist_ok=True)
    for chain in CHAINS:
        chain_dir = OUT_ROOT / chain.slug
        if args.clean and chain_dir.exists():
            shutil.rmtree(chain_dir)

    chain_results = [write_chain_layout(chain, args.validators_per_chain) for chain in CHAINS]
    compose = build_compose({item["chain"].slug: item["nodes"] for item in chain_results})
    (OUT_ROOT / "docker-compose.yml").write_text(compose + "\n")

    summary = {
        "summary": "generated 2-chain 16-node geth devnet",
        "topology": {
            "chains": len(CHAINS),
            "nodes_per_chain": NODE_COUNT,
            "validators_per_chain": args.validators_per_chain,
            "total_nodes": len(CHAINS) * NODE_COUNT,
        },
        "chains": [
            {
                "slug": item["chain"].slug,
                "chain_id": item["chain"].chain_id,
                "network_id": item["chain"].network_id,
                "rpc_endpoints": [
                    f"http://127.0.0.1:{port}" for port in item["chain"].rpc_ports
                ],
                "ws_endpoints": [
                    f"ws://127.0.0.1:{port}" for port in item["chain"].ws_ports
                ],
                "validators": [
                    node.account_address for node in item["nodes"][: args.validators_per_chain]
                ],
                "nodes": item["metadata"],
            }
            for item in chain_results
        ],
    }
    (OUT_ROOT / "network-summary.json").write_text(json.dumps(summary, indent=2) + "\n")
    write_top_level_readme(summary, args.validators_per_chain)
    print(json.dumps(summary, indent=2))


if __name__ == "__main__":
    main()

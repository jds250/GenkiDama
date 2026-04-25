#!/usr/bin/env python3
"""Launch the generated 2-chain x 8-node geth devnet directly on the host."""

from __future__ import annotations

import argparse
import json
import os
import shutil
import subprocess
import time
from dataclasses import dataclass
from pathlib import Path

from web3 import Web3


ROOT = Path(__file__).resolve().parents[2]
DEVNET_ROOT = ROOT / "local-devnet" / "multinode-geth"
SUMMARY_PATH = DEVNET_ROOT / "network-summary.json"
PID_PATH = DEVNET_ROOT / "local-pids.json"
DEFAULT_GETH_BIN = ROOT / ".local-bin" / "geth"


@dataclass
class RunningNode:
    service_name: str
    pid: int
    rpc_url: str | None
    ws_url: str | None
    log_path: str


def ensure_generated() -> dict:
    if not SUMMARY_PATH.exists():
        raise FileNotFoundError(
            f"missing {SUMMARY_PATH}; run generate_multinode_geth_devnet.py first"
        )
    return json.loads(SUMMARY_PATH.read_text())


def chain_slug_from_service(service_name: str) -> str:
    return "-".join(service_name.split("-")[:2])


def node_dir_for(service_name: str) -> Path:
    chain_slug = chain_slug_from_service(service_name)
    node_suffix = service_name.split("-")[-1]
    return DEVNET_ROOT / chain_slug / "nodes" / node_suffix.replace("node", "node")


def local_p2p_port(chain_slug: str, node_index: int) -> int:
    return 30303 + node_index + (0 if chain_slug == "chain-a" else 100)


def local_authrpc_port(chain_slug: str, node_index: int) -> int:
    return 18551 + node_index + (0 if chain_slug == "chain-a" else 100)


def wait_rpc(url: str, timeout: float = 20.0) -> None:
    deadline = time.time() + timeout
    w3 = Web3(Web3.HTTPProvider(url))
    while time.time() < deadline:
        try:
            if w3.is_connected():
                return
        except Exception:
            pass
        time.sleep(0.25)
    raise TimeoutError(f"RPC did not become ready: {url}")


def resolve_geth_bin(explicit: str | None) -> str:
    if explicit:
        return explicit
    if DEFAULT_GETH_BIN.exists():
        return str(DEFAULT_GETH_BIN)
    return shutil.which("geth") or "geth"


def write_local_config(node_dir: Path, chain: dict) -> Path:
    static_nodes = []
    for peer in chain["nodes"]:
        pubkey = peer["enode"].split("://", 1)[1].split("@", 1)[0]
        static_nodes.append(
            f"enode://{pubkey}@127.0.0.1:{local_p2p_port(chain['slug'], peer['node_index'])}"
        )

    rendered = ",\n".join(f'  "{enode}"' for enode in static_nodes)
    path = node_dir / "config.local.toml"
    path.write_text(
        "[Node.P2P]\n"
        "StaticNodes = [\n"
        f"{rendered}\n"
        "]\n"
        "TrustedNodes = [\n"
        f"{rendered}\n"
        "]\n"
        "NoDiscovery = true\n"
    )
    return path


def init_node(node_dir: Path, genesis_path: Path, geth_bin: str) -> None:
    if not (node_dir / ".initialized").exists():
        subprocess.run(
            [geth_bin, "init", "--datadir", str(node_dir), str(genesis_path)],
            check=True,
            capture_output=True,
            text=True,
        )
        static_nodes_src = node_dir / "static-nodes.json"
        static_nodes_dest = node_dir / "geth" / "static-nodes.json"
        static_nodes_dest.parent.mkdir(parents=True, exist_ok=True)
        shutil.copy2(static_nodes_src, static_nodes_dest)
        (node_dir / ".initialized").write_text("")

    keystore_dir = node_dir / "keystore"
    if not (node_dir / ".account-imported").exists():
        if not keystore_dir.exists() or not any(keystore_dir.iterdir()):
            subprocess.run(
                [
                    geth_bin,
                    "account",
                    "import",
                    "--datadir",
                    str(node_dir),
                    "--password",
                    str(node_dir / "password.txt"),
                    str(node_dir / "account.key"),
                ],
                check=True,
                capture_output=True,
                text=True,
            )
        (node_dir / ".account-imported").write_text("")


def start_node(node: dict, chain: dict, geth_bin: str, single_sealer: bool) -> RunningNode:
    service_name = node["service_name"]
    node_dir = DEVNET_ROOT / chain["slug"] / "nodes" / f"node-{node['node_index']:02d}"
    genesis_path = DEVNET_ROOT / chain["slug"] / "genesis.json"
    init_node(node_dir, genesis_path, geth_bin)
    local_config_path = write_local_config(node_dir, chain)

    log_path = node_dir / "geth.log"
    log_file = open(log_path, "a")

    rpc_url = None
    ws_url = None
    http_port = None
    ws_port = None
    if node["node_index"] == 0:
        http_port = chain["rpc_endpoints"][0].rsplit(":", 1)[1]
        ws_port = chain["ws_endpoints"][0].rsplit(":", 1)[1]
        rpc_url = chain["rpc_endpoints"][0]
        ws_url = chain["ws_endpoints"][0]
    elif node["node_index"] == 1:
        http_port = chain["rpc_endpoints"][1].rsplit(":", 1)[1]
        ws_port = chain["ws_endpoints"][1].rsplit(":", 1)[1]
        rpc_url = chain["rpc_endpoints"][1]
        ws_url = chain["ws_endpoints"][1]

    args = [
        geth_bin,
        "--config",
        str(local_config_path),
        "--datadir",
        str(node_dir),
        "--networkid",
        str(chain["network_id"]),
        "--syncmode",
        "full",
        "--gcmode",
        "archive",
        "--port",
        str(local_p2p_port(chain["slug"], node["node_index"])),
        "--nodekey",
        str(node_dir / "nodekey"),
        "--nat",
        "none",
        "--ipcdisable",
        "--authrpc.addr",
        "127.0.0.1",
        "--authrpc.port",
        str(local_authrpc_port(chain["slug"], node["node_index"])),
        "--authrpc.vhosts",
        "*",
        "--verbosity",
        "3",
    ]

    if http_port is not None and ws_port is not None:
        args += [
            "--http",
            "--http.addr",
            "127.0.0.1",
            "--http.port",
            str(http_port),
            "--http.api",
            "eth,net,web3,personal,txpool,admin,clique,debug",
            "--http.corsdomain",
            "*",
            "--http.vhosts",
            "*",
            "--ws",
            "--ws.addr",
            "127.0.0.1",
            "--ws.port",
            str(ws_port),
            "--ws.api",
            "eth,net,web3,personal,txpool,admin,clique,debug",
        ]

    should_mine = node["is_validator"] and (not single_sealer or node["node_index"] == 0)
    if should_mine:
        args += [
            "--mine",
            "--miner.etherbase",
            node["account_address"],
            "--unlock",
            node["account_address"],
            "--password",
            str(node_dir / "password.txt"),
            "--allow-insecure-unlock",
        ]

    proc = subprocess.Popen(
        args,
        stdout=log_file,
        stderr=subprocess.STDOUT,
        start_new_session=True,
        env={**os.environ},
    )

    return RunningNode(
        service_name=service_name,
        pid=proc.pid,
        rpc_url=rpc_url,
        ws_url=ws_url,
        log_path=str(log_path),
    )


def main() -> None:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument(
        "--clean-stale",
        action="store_true",
        help="remove stale pid file before starting",
    )
    parser.add_argument(
        "--geth-bin",
        default=None,
        help="path to a compatible geth binary; defaults to .local-bin/geth when available",
    )
    parser.add_argument(
        "--single-sealer",
        action="store_true",
        help="only enable mining on node-00 of each chain for more deterministic experiments",
    )
    args = parser.parse_args()

    if PID_PATH.exists():
        if not args.clean_stale:
            raise RuntimeError(
                f"{PID_PATH} already exists; stop the local devnet first or rerun with --clean-stale"
            )
        PID_PATH.unlink()

    summary = ensure_generated()
    geth_bin = resolve_geth_bin(args.geth_bin)
    running: list[RunningNode] = []
    try:
        for chain in summary["chains"]:
            for node in chain["nodes"]:
                running.append(start_node(node, chain, geth_bin, args.single_sealer))

        for item in running:
            if item.rpc_url:
                wait_rpc(item.rpc_url)

        payload = {
            "summary": "launched local multinode geth devnet",
            "nodes": [
                {
                    "service_name": item.service_name,
                    "pid": item.pid,
                    "rpc_url": item.rpc_url,
                    "ws_url": item.ws_url,
                    "log_path": item.log_path,
                }
                for item in running
            ],
        }
        PID_PATH.write_text(json.dumps(payload, indent=2) + "\n")
        print(json.dumps(payload, indent=2))
    except Exception:
        for item in running:
            try:
                os.kill(item.pid, 15)
            except OSError:
                pass
        raise


if __name__ == "__main__":
    main()

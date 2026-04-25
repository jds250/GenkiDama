#!/usr/bin/env python3
"""Stop locally launched multinode geth processes."""

from __future__ import annotations

import json
import os
import signal
import time
from pathlib import Path


ROOT = Path(__file__).resolve().parents[2]
PID_PATH = ROOT / "local-devnet" / "multinode-geth" / "local-pids.json"


def main() -> None:
    if not PID_PATH.exists():
        print('{"summary": "no local multinode geth processes recorded"}')
        return

    payload = json.loads(PID_PATH.read_text())
    for node in payload.get("nodes", []):
        pid = node["pid"]
        try:
            os.kill(pid, signal.SIGTERM)
        except OSError:
            continue

    time.sleep(1.0)
    for node in payload.get("nodes", []):
        pid = node["pid"]
        try:
            os.kill(pid, 0)
        except OSError:
            continue
        try:
            os.kill(pid, signal.SIGKILL)
        except OSError:
            pass

    PID_PATH.unlink(missing_ok=True)
    print('{"summary": "stopped local multinode geth devnet"}')


if __name__ == "__main__":
    main()

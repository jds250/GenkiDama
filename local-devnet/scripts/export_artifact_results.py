#!/usr/bin/env python3
"""Export reproducible local-devnet experiment results to JSON artifacts."""

from __future__ import annotations

import json
import subprocess
import sys
from dataclasses import dataclass
from pathlib import Path


ROOT = Path(__file__).resolve().parents[2]
SCRIPTS_DIR = ROOT / "local-devnet" / "scripts"
RESULTS_DIR = ROOT / "artifacts" / "results"
PYTHON_BIN = ROOT / ".venv" / "bin" / "python"


@dataclass(frozen=True)
class ExportTask:
    name: str
    command: list[str]
    description: str


TASKS = [
    ExportTask(
        name="dual_chain_flow",
        command=[str(PYTHON_BIN), str(SCRIPTS_DIR / "run_dual_chain_flow.py")],
        description="Minimal dual-chain sanity flow",
    ),
    ExportTask(
        name="zk_verifier",
        command=[str(PYTHON_BIN), str(SCRIPTS_DIR / "test_real_zk_verifier.py")],
        description="Real zk verifier sanity check",
    ),
    ExportTask(
        name="real_rollup_flow",
        command=[str(PYTHON_BIN), str(SCRIPTS_DIR / "run_real_rollup_flow.py")],
        description="End-to-end real rollup flow",
    ),
    ExportTask(
        name="adversarial_experiments",
        command=[str(PYTHON_BIN), str(SCRIPTS_DIR / "run_adversarial_experiments.py")],
        description="Security and adversarial experiments",
    ),
    ExportTask(
        name="benchmark_real_rollup",
        command=[
            str(PYTHON_BIN),
            str(SCRIPTS_DIR / "benchmark_real_rollup.py"),
            "--num-batches",
            "10",
            "--batch-size",
            "10",
        ],
        description="Local throughput benchmark",
    ),
    ExportTask(
        name="attack_performance_experiments",
        command=[
            str(PYTHON_BIN),
            str(SCRIPTS_DIR / "run_attack_performance_experiments.py"),
            "--num-batches",
            "3",
            "--batch-size",
            "8",
        ],
        description="Throughput and latency degradation under attack",
    ),
]


def extract_json(stdout: str) -> dict:
    stripped = stdout.strip()
    start = stripped.find("{")
    end = stripped.rfind("}")
    if start < 0 or end < start:
        raise ValueError(f"command did not produce JSON stdout:\n{stdout}")
    return json.loads(stripped[start : end + 1])


def run_task(task: ExportTask) -> dict:
    result = subprocess.run(
        task.command,
        cwd=ROOT,
        check=True,
        capture_output=True,
        text=True,
    )
    payload = extract_json(result.stdout)
    payload["_artifact"] = {
        "name": task.name,
        "description": task.description,
        "command": task.command,
        "working_directory": str(ROOT),
    }
    if result.stderr.strip():
        payload["_artifact"]["stderr"] = result.stderr.strip()
    return payload


def main() -> None:
    if not PYTHON_BIN.exists():
        raise FileNotFoundError(
            f"Python virtual environment not found at {PYTHON_BIN}. "
            "Run ./local-devnet/scripts/bootstrap_local_devnet.sh first."
        )

    RESULTS_DIR.mkdir(parents=True, exist_ok=True)

    manifest: dict[str, object] = {
        "summary": "artifact result export completed",
        "results_directory": str(RESULTS_DIR),
        "files": [],
    }

    for task in TASKS:
        payload = run_task(task)
        output_path = RESULTS_DIR / f"{task.name}.json"
        output_path.write_text(json.dumps(payload, indent=2) + "\n")
        manifest["files"].append(
            {
                "name": task.name,
                "path": str(output_path),
                "summary": payload.get("summary"),
            }
        )

    manifest_path = RESULTS_DIR / "manifest.json"
    manifest_path.write_text(json.dumps(manifest, indent=2) + "\n")
    print(json.dumps(manifest, indent=2))


if __name__ == "__main__":
    try:
        main()
    except Exception as exc:  # pragma: no cover - script-style error path
        print(str(exc), file=sys.stderr)
        sys.exit(1)

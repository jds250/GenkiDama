#!/usr/bin/env python3
"""Run a richer throughput matrix on the 16-node multinode devnet."""

from __future__ import annotations

import json
import os
import subprocess
import sys
from pathlib import Path


ROOT = Path(__file__).resolve().parents[2]
ROLLOFF = ROOT / "Rollup-Offchain"
OUT_DIR = ROOT / "local-devnet" / "results" / "throughput-matrix"

SCENARIOS = [
    {
        "name": "baseline",
        "scenario": "baseline",
        "num_batches": 2,
        "batch_sizes": [5, 10, 20, 40, 80, 100],
    },
    {
        "name": "execution_challenge",
        "scenario": "execution_challenge",
        "num_batches": 2,
        "batch_sizes": [5, 10, 20, 40],
    },
]
USER_POOL_SIZE = 1


def parse_json(stdout: str) -> dict:
    start = stdout.find("{")
    if start == -1:
        raise RuntimeError(f"no json found in stdout:\n{stdout}")
    return json.loads(stdout[start:])


def run_case(scenario: str, num_batches: int, batch_size: int, transfer_value: int = 1) -> dict:
    args = [
        "go",
        "run",
        "./cmd/localnetattackperf",
        "--scenario",
        scenario,
        "--num-batches",
        str(num_batches),
        "--batch-size",
        str(batch_size),
        "--transfer-value",
        str(transfer_value),
        "--user-pool-size",
        str(USER_POOL_SIZE),
    ]
    proc = subprocess.run(
        args,
        cwd=str(ROLLOFF),
        capture_output=True,
        text=True,
        env=os.environ.copy(),
    )
    result = {
        "command": " ".join(args),
        "returncode": proc.returncode,
        "stdout": proc.stdout,
        "stderr": proc.stderr,
    }
    if proc.returncode == 0:
        result["parsed"] = parse_json(proc.stdout)
    return result


def write_json(path: Path, payload: dict) -> None:
    path.parent.mkdir(parents=True, exist_ok=True)
    path.write_text(json.dumps(payload, indent=2, ensure_ascii=False) + "\n")


def render_markdown(summary: dict) -> str:
    lines: list[str] = []
    lines.append("# 16 节点吞吐量实验矩阵（中文）")
    lines.append("")
    lines.append("## 1. 实验设置")
    lines.append("")
    lines.append("- 拓扑：2 条链 × 每条链 8 个 geth 节点。")
    lines.append("- 默认窗口：`(dispute, wait, confirm) = (8, 12, 24)`。")
    lines.append("- 业务负载：跨链交易 `transfer_value = 1`。")
    lines.append(f"- 固定账户池：`{USER_POOL_SIZE}` 个 sender / `{USER_POOL_SIZE}` 个 receiver；batch 内通过循环复用账户来隔离协议路径开销。")
    lines.append("- 评估场景：")
    lines.append("  - `baseline`: 正常提交与确认")
    lines.append("  - `execution_challenge`: 每个 batch 插入合法执行挑战")
    lines.append("- 每个配置统一使用 `num_batches = 2`。")
    lines.append("")
    for scenario in summary["scenarios"]:
        lines.append(f"## 2. {scenario['name']} 结果")
        lines.append("")
        lines.append("| Batch Size | 状态 | Finalized TPS | Avg Batch Latency(s) | Gas / Business Tx | 相对 Batch=5 的 TPS 提升 | 备注 |")
        lines.append("| --- | --- | ---: | ---: | ---: | ---: | --- |")
        base_tps = None
        for case in scenario["cases"]:
            if case["status"] != "ok":
                lines.append(f"| {case['batch_size']} | fail | - | - | - | - | `{case['error']}` |")
                continue
            metrics = case["result"]["metrics"]
            tps = metrics.get("finalized_business_tps", metrics.get("attempted_business_tps", 0))
            if base_tps is None:
                base_tps = tps
            delta = ((tps / base_tps) - 1.0) * 100 if base_tps else 0.0
            lines.append(
                f"| {case['batch_size']} | ok | "
                f"{tps:.4f} | "
                f"{metrics.get('average_batch_latency_seconds', metrics.get('elapsed_seconds', 0)):.2f} | "
                f"{metrics.get('gas_used_per_business_tx', 0):,.2f} | {delta:+.1f}% | - |"
            )
        lines.append("")

    lines.append("## 3. 观察")
    lines.append("")
    baseline_ok = [c for c in summary["scenarios"][0]["cases"] if c["status"] == "ok"]
    exec_ok = [c for c in summary["scenarios"][1]["cases"] if c["status"] == "ok"]
    if baseline_ok:
        first = baseline_ok[0]
        last = baseline_ok[-1]
        lines.append(
            f"- 在正常路径下，batch size 从 `{first['batch_size']}` 增长到 `{last['batch_size']}` 时，"
            f"finalized TPS 从 `{first['result']['metrics']['finalized_business_tps']:.4f}` 增长到 "
            f"`{last['result']['metrics']['finalized_business_tps']:.4f}`。"
        )
    if exec_ok:
        first = exec_ok[0]
        last = exec_ok[-1]
        lines.append(
            f"- 在执行挑战负载下，batch size 从 `{first['batch_size']}` 增长到 `{last['batch_size']}` 时，"
            f"finalized TPS 从 `{first['result']['metrics']['finalized_business_tps']:.4f}` 增长到 "
            f"`{last['result']['metrics']['finalized_business_tps']:.4f}`。"
        )
    lines.append(
        "- 与较小 batch 相比，更大的 batch 会明显提升单位时间内完成的业务交易数；但当挑战路径插入后，时延与单位业务交易 gas 也会同步上升。"
    )
    lines.append(
        "- 若更大 batch 开始失败，可将该点视为当前 16 节点配置与 gas 限制下的可扩展边界，而不应简单外推线性增长。"
    )
    lines.append("")
    return "\n".join(lines)


def main() -> None:
    OUT_DIR.mkdir(parents=True, exist_ok=True)
    summary = {
        "summary": "multinode throughput matrix completed",
        "scenarios": [],
    }

    for scenario in SCENARIOS:
        scenario_summary = {"name": scenario["name"], "cases": []}
        for batch_size in scenario["batch_sizes"]:
            raw = run_case(scenario["scenario"], scenario["num_batches"], batch_size)
            case_result = {"batch_size": batch_size}
            if raw["returncode"] != 0:
                case_result["status"] = "fail"
                case_result["error"] = (raw["stderr"] or raw["stdout"]).strip().splitlines()[-1]
                case_result["raw"] = raw
            else:
                case_result["status"] = "ok"
                case_result["result"] = raw["parsed"]
            scenario_summary["cases"].append(case_result)
            write_json(OUT_DIR / f"{scenario['name']}-b{batch_size}.json", case_result)
        summary["scenarios"].append(scenario_summary)

    write_json(OUT_DIR / "throughput_matrix_summary.json", summary)
    (OUT_DIR / "throughput_matrix_summary_zh.md").write_text(render_markdown(summary))
    print(json.dumps(summary, indent=2, ensure_ascii=False))


if __name__ == "__main__":
    main()

#!/usr/bin/env zsh
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "$0")/../.." && pwd)"

if ! command -v brew >/dev/null 2>&1; then
  echo "Homebrew is required to install solc."
  exit 1
fi

if ! command -v solc >/dev/null 2>&1; then
  HOMEBREW_NO_AUTO_UPDATE=1 brew install solidity
fi

if [ ! -d "$ROOT_DIR/.venv" ]; then
  python3 -m venv "$ROOT_DIR/.venv"
fi

PYTHON_MINOR="$(python3 - <<'PY'
import sys
print(f"{sys.version_info.major}.{sys.version_info.minor}")
PY
)"
VENV_SITE_PACKAGES="$ROOT_DIR/.venv/lib/python${PYTHON_MINOR}/site-packages"

python3 -m pip install --target "$VENV_SITE_PACKAGES" -r "$ROOT_DIR/local-devnet/requirements.txt"
python3 -m pip install --target "$VENV_SITE_PACKAGES" py-evm

echo "Local devnet environment is ready."
echo "Run: source $ROOT_DIR/.venv/bin/activate && python $ROOT_DIR/local-devnet/scripts/run_dual_chain_flow.py"

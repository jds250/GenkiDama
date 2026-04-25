#!/bin/sh
set -eu

DATADIR="${DATADIR:-/data}"
GENESIS_PATH="${GENESIS_PATH:-/config/genesis.json}"
ACCOUNT_KEY_PATH="${ACCOUNT_KEY_PATH:-$DATADIR/account.key}"
PASSWORD_FILE="${PASSWORD_FILE:-$DATADIR/password.txt}"
NODEKEY_PATH="${NODEKEY_PATH:-$DATADIR/nodekey}"
CONFIG_FILE="${CONFIG_FILE:-$DATADIR/config.toml}"
HTTP_PORT="${HTTP_PORT:-8545}"
WS_PORT="${WS_PORT:-8546}"
AUTHRPC_PORT="${AUTHRPC_PORT:-8551}"
P2P_PORT="${P2P_PORT:-30303}"
NETWORK_ID="${NETWORK_ID:?missing NETWORK_ID}"
UNLOCK_ADDRESS="${UNLOCK_ADDRESS:-}"
IS_VALIDATOR="${IS_VALIDATOR:-false}"

mkdir -p "$DATADIR/geth"

if [ ! -f "$DATADIR/.initialized" ]; then
  geth init --datadir "$DATADIR" "$GENESIS_PATH"
  if [ -f "$DATADIR/static-nodes.json" ]; then
    cp "$DATADIR/static-nodes.json" "$DATADIR/geth/static-nodes.json"
  fi
  touch "$DATADIR/.initialized"
fi

if [ ! -f "$DATADIR/.account-imported" ]; then
  if [ ! -d "$DATADIR/keystore" ] || [ -z "$(ls -A "$DATADIR/keystore" 2>/dev/null)" ]; then
    geth account import --datadir "$DATADIR" --password "$PASSWORD_FILE" "$ACCOUNT_KEY_PATH"
  fi
  touch "$DATADIR/.account-imported"
fi

COMMON_ARGS="
  --config $CONFIG_FILE
  --datadir $DATADIR
  --networkid $NETWORK_ID
  --syncmode full
  --gcmode archive
  --http
  --http.addr 0.0.0.0
  --http.port $HTTP_PORT
  --http.api eth,net,web3,personal,txpool,admin,clique,debug
  --http.corsdomain *
  --http.vhosts *
  --ws
  --ws.addr 0.0.0.0
  --ws.port $WS_PORT
  --ws.api eth,net,web3,personal,txpool,admin,clique,debug
  --authrpc.addr 0.0.0.0
  --authrpc.port $AUTHRPC_PORT
  --authrpc.vhosts *
  --port $P2P_PORT
  --nodekey $NODEKEY_PATH
  --nat none
  --ipcdisable
  --verbosity 3
"

VALIDATOR_ARGS=""
if [ "$IS_VALIDATOR" = "true" ]; then
  VALIDATOR_ARGS="
    --mine
    --miner.etherbase $UNLOCK_ADDRESS
    --unlock $UNLOCK_ADDRESS
    --password $PASSWORD_FILE
    --allow-insecure-unlock
  "
fi

# shellcheck disable=SC2086
exec geth $COMMON_ARGS $VALIDATOR_ARGS ${EXTRA_GETH_ARGS:-}

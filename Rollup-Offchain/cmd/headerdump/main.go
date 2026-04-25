package main

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

type headerBundle struct {
	BlockNumber uint64 `json:"block_number"`
	Hash        string `json:"hash"`
	ParentHash  string `json:"parent_hash"`
	StateRoot   string `json:"state_root"`
	HeaderBytes string `json:"header_bytes"`
}

func main() {
	rpcURL := flag.String("rpc-url", "", "Ethereum RPC URL")
	blockNumber := flag.Uint64("block-number", 0, "Block number to fetch")
	flag.Parse()

	if *rpcURL == "" {
		log.Fatal("missing --rpc-url")
	}

	client, err := ethclient.Dial(*rpcURL)
	if err != nil {
		log.Fatal(err)
	}

	block, err := client.BlockByNumber(context.Background(), new(big.Int).SetUint64(*blockNumber))
	if err != nil {
		log.Fatal(err)
	}

	var buf bytes.Buffer
	if err := block.Header().EncodeRLP(&buf); err != nil {
		log.Fatal(err)
	}

	output := headerBundle{
		BlockNumber: block.NumberU64(),
		Hash:        block.Hash().Hex(),
		ParentHash:  block.ParentHash().Hex(),
		StateRoot:   block.Root().Hex(),
		HeaderBytes: "0x" + hex.EncodeToString(buf.Bytes()),
	}

	encoded, err := json.Marshal(output)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(encoded))
}

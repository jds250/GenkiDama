package main

import (
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"rollup-offchain/zk"
	"strings"
)

type proofJSON struct {
	A     [2]string    `json:"a"`
	B     [2][2]string `json:"b"`
	C     [2]string    `json:"c"`
	Input [4]string    `json:"input"`
}

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "export-verifier":
		if err := zk.Groth16Init(); err != nil {
			exitWithError(err)
		}
		writeJSON(map[string]string{
			"status":   "ok",
			"contract": zk.CONTRACT_PATH + string(os.PathSeparator) + "ZKVerifier.sol",
		})
	case "prove-header-hash":
		proveHeaderHash(os.Args[2:])
	default:
		usage()
		os.Exit(1)
	}
}

func usage() {
	fmt.Fprintln(os.Stderr, "usage:")
	fmt.Fprintln(os.Stderr, "  go run ./cmd/zktool export-verifier")
	fmt.Fprintln(os.Stderr, "  go run ./cmd/zktool prove-header-hash --hash 0x...")
}

func proveHeaderHash(args []string) {
	fs := flag.NewFlagSet("prove-header-hash", flag.ExitOnError)
	hashHex := fs.String("hash", "", "32-byte header hash hex string")
	_ = fs.Parse(args)

	if *hashHex == "" {
		exitWithError(fmt.Errorf("missing --hash"))
	}

	hashBytes, err := decodeHash(*hashHex)
	if err != nil {
		exitWithError(err)
	}

	if err := zk.Groth16Init(); err != nil {
		exitWithError(err)
	}

	proof, err := zk.GenerateHeaderHashProof(hashBytes)
	if err != nil {
		exitWithError(err)
	}

	publicInputs := zk.GetHeaderProofPublicInputs(hashBytes)
	writeJSON(proofJSON{
		A: [2]string{
			proof.A[0].String(),
			proof.A[1].String(),
		},
		B: [2][2]string{
			{
				proof.B[0][0].String(),
				proof.B[0][1].String(),
			},
			{
				proof.B[1][0].String(),
				proof.B[1][1].String(),
			},
		},
		C: [2]string{
			proof.C[0].String(),
			proof.C[1].String(),
		},
		Input: [4]string{
			publicInputs[0].String(),
			publicInputs[1].String(),
			publicInputs[2].String(),
			publicInputs[3].String(),
		},
	})
}

func decodeHash(value string) ([]byte, error) {
	trimmed := strings.TrimPrefix(value, "0x")
	bytesValue, err := hex.DecodeString(trimmed)
	if err != nil {
		return nil, err
	}
	if len(bytesValue) != 32 {
		return nil, fmt.Errorf("expected 32-byte hash, got %d bytes", len(bytesValue))
	}
	return bytesValue, nil
}

func writeJSON(value interface{}) {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(value); err != nil {
		exitWithError(err)
	}
}

func exitWithError(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

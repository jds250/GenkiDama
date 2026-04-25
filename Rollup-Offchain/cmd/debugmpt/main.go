package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"

	"rollup-offchain/challenger"
	"rollup-offchain/contract"
	"rollup-offchain/localnet"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/trie"
)

func main() {
	ctx := context.Background()

	net, err := localnet.Load("")
	if err != nil {
		log.Fatal(err)
	}

	chainB, err := net.MustChain("chain-b")
	if err != nil {
		log.Fatal(err)
	}
	client, err := ethclient.Dial(chainB.RPCURL)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	crossAddr, err := chainB.Contract("CrossRollup")
	if err != nil {
		log.Fatal(err)
	}
	cross, err := contract.NewCrossRollup(crossAddr, client)
	if err != nil {
		log.Fatal(err)
	}

	blockCount, err := cross.BlockLen(nil)
	if err != nil {
		log.Fatal(err)
	}
	if blockCount.Cmp(big.NewInt(2)) < 0 {
		log.Fatalf("need at least two rollup blocks, got %s", blockCount)
	}

	targetRollupBlock := big.NewInt(1)
	head, err := client.BlockByNumber(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	accountProof, slots, err := challenger.FetchRollupBlockStorageProof(
		ctx,
		client,
		crossAddr,
		targetRollupBlock,
		head.Number(),
	)
	if err != nil {
		log.Fatal(err)
	}

	accountValue, err := rlp.EncodeToBytes([]interface{}{
		uint64(accountProof.Nonce),
		(*big.Int)(accountProof.Balance),
		accountProof.StorageHash[:],
		accountProof.CodeHash[:],
	})
	if err != nil {
		log.Fatal(err)
	}

	accountKey := crypto.Keccak256(crossAddr.Bytes())
	ok, err := verifyProof(head.Root(), accountKey, accountValue, accountProof.AccountProof)
	if err != nil {
		log.Fatalf("account proof failed: %v", err)
	}
	fmt.Printf("account proof ok=%v storageRoot=%s head=%d stateRoot=%s\n", ok, accountProof.StorageHash.Hex(), head.NumberU64(), head.Root().Hex())

	for index, slot := range slots {
		entry := accountProof.StorageProof[index]
		rawValue := common.LeftPadBytes((*big.Int)(entry.Value).Bytes(), 32)
		trimmed := trimLeftZeroes(rawValue)
		storageValue := []byte{}
		if len(trimmed) != 0 {
			storageValue, err = rlp.EncodeToBytes(trimmed)
			if err != nil {
				log.Fatal(err)
			}
		}

		storageKey := crypto.Keccak256(common.LeftPadBytes(slot.Bytes(), 32))
		ok, err := verifyProof(accountProof.StorageHash, storageKey, storageValue, entry.Proof)
		if err != nil {
			log.Fatalf("field %d storage proof failed: %v", index, err)
		}
		fmt.Printf("field %d slot=%s trieKey=%s value=%s ok=%v proofNodes=%d\n", index, slot.String(), hex.EncodeToString(storageKey), hexutil.Encode(storageValue), ok, len(entry.Proof))
	}
}

func verifyProof(root common.Hash, key []byte, expectedValue []byte, proof []hexutil.Bytes) (bool, error) {
	proofDB := rawdb.NewMemoryDatabase()
	for _, node := range proof {
		hash := crypto.Keccak256(node)
		if err := proofDB.Put(hash, node); err != nil {
			return false, err
		}
	}

	value, err := trie.VerifyProof(root, key, proofDB)
	if err != nil {
		return false, err
	}

	return common.Bytes2Hex(value) == common.Bytes2Hex(expectedValue), nil
}

func trimLeftZeroes(data []byte) []byte {
	for len(data) > 0 && data[0] == 0 {
		data = data[1:]
	}
	return data
}

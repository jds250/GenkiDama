package contract

import (
	"math/big"
	"rollup-offchain/trie"
)

func TrieProof2SCProof(proof trie.Bag) *MPTVerifierMerkleProof {
	scProof := [][]byte{}
	for i := 0; i < len(proof.Proof); i++ {
		if i%2 == 1 {
			scProof = append(scProof, proof.Proof[i])
		}
	}
	scExpectedValue := proof.Value
	scProofIndex := big.NewInt(0)
	scKey := proof.Key
	scExpectedRoot := trie.Hash{}
	scExpectedRoot.SetBytes(proof.Proof[0])
	scMPTProof := MPTVerifierMerkleProof{
		ExpectedRoot:  scExpectedRoot,
		Key:           scKey,
		Proof:         scProof,
		KeyIndex:      scProofIndex,
		ProofIndex:    big.NewInt(0),
		ExpectedValue: scExpectedValue,
	}

	return &scMPTProof
}

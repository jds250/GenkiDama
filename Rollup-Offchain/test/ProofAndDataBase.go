package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/syndtr/goleveldb/leveldb"
	"log"
	"math/big"
	"os"
	"rollup-offchain/contract"
	"rollup-offchain/trie"
)

func DataBaseTest() {
	fmt.Println("************************** DATABASE TEST **************************")

	// 打开数据库
	testDB, err := leveldb.OpenFile("./DataBaseA", nil)
	if err != nil {
		fmt.Println("ERROR: database open fail!")
		return
	}

	// 获取最新状态根 [32]byte格式
	var testRoot trie.Hash
	testRootBytes, err := testDB.Get([]byte(MyStateRoot), nil)
	if err != nil {
		fmt.Println("ERROR: state root not found! Create a new Root!")
		testRootBytes = trie.HexToHash("56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421").Bytes()
	}
	testRoot.SetBytes(testRootBytes)

	// 设置一个假的状态树testRoot跟开始创建状态树testTrie
	testTrie, _ := trie.NewTrie(testRoot, testDB)

	//for i := 0; i < 100; i++ {
	//	tmp := fmt.Sprintf("test %d", i)
	//	tmpHash := trie.Keccak256Hash([]byte(tmp))
	//	testTrie.Update(tmpHash.Bytes(), []byte("just test2"))
	//}

	testTrie.Update([]byte("big int"), big.NewInt(16909060).Bytes())
	var getValue big.Int
	valueBytes, err := testTrie.TryGet([]byte("big int"))
	if err != nil {
		log.Fatal(err)
	}
	bigValue := getValue.SetBytes(valueBytes)
	fmt.Println("value =", bigValue.Uint64())
	fmt.Println("value =", valueBytes)

	//fmt.Println("************************** Get Proof  **************************")
	//// 读取数据
	//var getValue big.Int
	//valueBytes, err := testTrie.TryGet(trie.Keccak256Hash([]byte("test 1")).Bytes())
	//if err != nil {
	//	log.Fatal(err)
	//}
	//getValue.SetBytes(valueBytes)
	//fmt.Println("value =", string(valueBytes))
	//
	//// 查看证明格式
	//proof := testTrie.Getproof(trie.Keccak256Hash([]byte("test 1")).Bytes())
	//for i := 0; i < len(proof.Proof); i++ {
	//	fmt.Println(i, " = ", proof.Proof[i], "(", len(proof.Proof[i]), ")")
	//}
	//fmt.Println("hash proof[3] = ", crypto.Keccak256(proof.Proof[3]))
	//fmt.Println("data bytes = ", valueBytes)
	//fmt.Println("data hash = ", crypto.Keccak256(valueBytes))
	//
	//fmt.Println("************************** Test Update  **************************")
	//// decode old proof
	//var nodeHashs [][]byte
	//for i := 0; i < len(proof.Proof); i++ {
	//	if i%2 == 1 {
	//		nodeHashs = append(nodeHashs, proof.Proof[i])
	//	}
	//}

	// update value and get new proof

	// compare with new proof

	fmt.Println("************************** Close DataBase  **************************")
	// 更新到数据库
	root, err := testTrie.Commit(nil)
	if err != nil {
		fmt.Println("ERROR: Trie Root Update Fail!", err)
	}

	// 更新数据库根
	err = testDB.Put([]byte(MyStateRoot), root.Bytes(), nil)
	if err != nil {
		fmt.Println("ERROR: Trie Root Update Fail!", err)
	}
	fmt.Println("new root =", root.Hex())
	testDB.Close()
}

func ProofTest() {

	var err error

	fmt.Println("************************** DATABASE TEST **************************")

	// 打开数据库
	testDB, err := leveldb.OpenFile("./DataBaseA", nil)
	if err != nil {
		fmt.Println("ERROR: database open fail!")
		return
	}

	// 获取最新状态根 [32]byte格式
	var testRoot trie.Hash
	testRootBytes, err := testDB.Get([]byte(MyStateRoot), nil)
	if err != nil {
		fmt.Println("ERROR: state root not found! Create a new Root!")
		testRootBytes = trie.HexToHash("56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421").Bytes()
	}
	testRoot.SetBytes(testRootBytes)

	// 设置一个假的状态树testRoot跟开始创建状态树testTrie
	testTrie, _ := trie.NewTrie(testRoot, testDB)

	// 存入数据
	//testAddr1 := trie.HexToHash("56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b431")
	//testTrie.Update(testAddr1.Bytes(), []byte("just test"))
	//for i := 0; i < 100; i++ {
	//	tmp := fmt.Sprintf("test %d", i)
	//	tmpHash := trie.BytesToHash([]byte(tmp))
	//	testTrie.Update(tmpHash.Bytes(), []byte("just test"))
	//}

	// 读取数据
	var getValue big.Int
	valueBytes, err := testTrie.TryGet(trie.Keccak256Hash([]byte("test 1")).Bytes())
	if err != nil {
		log.Fatal(err)
	}
	getValue.SetBytes(valueBytes)
	fmt.Println("value =", string(valueBytes))

	// 查看证明格式
	proof := testTrie.Getproof(trie.Keccak256Hash([]byte("test 1")).Bytes())
	for i := 0; i < len(proof.Proof); i++ {
		fmt.Println(i, " = ", proof.Proof[i], "(", len(proof.Proof[i]), ")")
	}
	fmt.Println("hash proof[3] = ", crypto.Keccak256(proof.Proof[3]))
	fmt.Println("data bytes = ", valueBytes)
	fmt.Println("data hash = ", crypto.Keccak256(valueBytes))

	// RLP Decode
	//rlpProof := [][]byte{}

	fmt.Println("************************** Proof TEST **************************")
	elems, _, err := rlp.SplitList(proof.Proof[1])
	cnt, _ := rlp.CountValues(elems)
	fmt.Println("elems: ", elems, "(", cnt, ")")

	elems, _, err = rlp.SplitList(proof.Proof[3])
	cnt, _ = rlp.CountValues(elems)
	fmt.Println("elems: ", elems, "(", cnt, ")")

	//fmt.Println("proof = ", proof.Proof)
	//tryProof, err := testTrie.TryProve(testAddr1.Bytes())
	//fmt.Println("proof = ", tryProof)

	// 更新到数据库
	root, err := testTrie.Commit(nil)
	if err != nil {
		fmt.Println("ERROR: Trie Root Update Fail!", err)
	}

	// 更新数据库根
	err = testDB.Put([]byte(MyStateRoot), root.Bytes(), nil)
	if err != nil {
		fmt.Println("ERROR: Trie Root Update Fail!", err)
	}
	fmt.Println("new root =", root.Hex())
	testDB.Close()

	// 本地验证
	fmt.Println("************************** Local Verifier TEST **************************")

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
	scMPTProof := contract.MPTVerifierMerkleProof{
		ExpectedRoot:  scExpectedRoot,
		Key:           scKey,
		Proof:         scProof,
		KeyIndex:      scProofIndex,
		ProofIndex:    big.NewInt(0),
		ExpectedValue: scExpectedValue,
	}

	fmt.Println("expected root = ", scMPTProof.ExpectedRoot)
	fmt.Println("expected value = ", scMPTProof.ExpectedValue)

	fmt.Println("************************** Verifier Init **************************")
	client, err := ethclient.Dial("ws://172.18.166.229:8546")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(client, "successfully connected")

	// file := "../../data/keystore/UTC-USER2--2023-09-10T10-46-57.143272940Z--5e6aff8b8f168a52e2b8c2882442aa51482292f8"
	file := "./data/UTC-USER2"
	jsonBytes, err := os.ReadFile(file)
	// password, _ := os.ReadFile("../../data/password.txt")
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	chainID, err := client.ChainID(ctx)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println("chainID = ", chainID)
	// Key, _ := keystore.DecryptKey(jsonBytes, string(password))
	Key, _ := keystore.DecryptKey(jsonBytes, myPassword)
	height, _ := client.BlockNumber(ctx)
	//blockNum, _ := client.BlockNumber(ctx)
	fmt.Println(Key, chainID, height)

	// 调用Verifier
	fmt.Println("************************** Verifier TEST **************************")
	// set tx config
	//auth, err := bind.NewTransactorWithChainID(strings.NewReader(string(jsonBytes)), myPassword, chainID)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//auth.GasPrice = big.NewInt(1)
	//auth.GasLimit = uint64(2100000)
	caller := &bind.CallOpts{}

	// build the MPTVerifier instance
	MVAddr := common.HexToAddress(MPTVerifierAddr)
	MV, _ := contract.NewMPTVerifier(MVAddr, client)
	//
	//MPTProof := contract.MPTVerifierMerkleProof{
	//	ExpectedRoot:  root,
	//	Key:           testAddr1.Bytes(),
	//	Proof:         proof.Proof,
	//	KeyIndex:      big.NewInt(0),
	//	ProofIndex:    big.NewInt(0),
	//	ExpectedValue: valueBytes,
	//}
	//
	res, err := MV.VerifyProof(caller, scMPTProof)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("res =", res)

}

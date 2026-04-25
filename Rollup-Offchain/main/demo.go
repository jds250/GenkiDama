package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/syndtr/goleveldb/leveldb"
	"log"
	"math/big"
	"os"
	"rollup-offchain/contract"
	"rollup-offchain/executor"
	"rollup-offchain/trie"
	"strings"
)

// 链下状态树操作

func accountInit(targetTrie *trie.Trie) {
	// 让每个可能账户都包含100的coin
	for i := int64(0); i <= int64(executor.UserRange); i++ {
		addrInt := big.NewInt(i)
		addrHash := trie.BigToHash(addrInt)

		valueInt := big.NewInt(100)

		targetTrie.Update(addrHash.Bytes(), valueInt.Bytes())
	}
}

func showState(targetTrie *trie.Trie) {
	// 查看每个账户的余额
	for i := int64(0); i <= int64(executor.UserRange); i++ {
		addrInt := big.NewInt(i)
		addrHash := trie.BigToHash(addrInt)

		balanceByte := targetTrie.Get(addrHash.Bytes())
		var balanceInt big.Int
		balanceInt.SetBytes(balanceByte)

		fmt.Print(i, "-", balanceInt.Int64(), "\t")
	}
	fmt.Println("")
}

// Trie使用模板
func demo() {
	fmt.Println("************************** TEST **************************")

	var err error

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
	testAddr1 := trie.HexToHash("56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b431")
	// testAddr2 := trie.HexToHash("56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b432")
	// testIntValue := big.NewInt(10)

	//testValue := []byte("200")
	//testTrie.Update(testAddr1.Bytes(), testValue)
	//fmt.Println("value =", string(testTrie.Get(testAddr1.Bytes())))

	// 读取数据
	var getValue big.Int
	valueBytes, err := testTrie.TryGet(testAddr1.Bytes())
	if err != nil {
		fmt.Println("ERROR: Trie Vaule TruGet Fail!", err)
	}
	getValue.SetBytes(valueBytes)
	fmt.Println("value =", string(valueBytes))

	// 获取证明
	proof := testTrie.Getproof(testAddr1.Bytes())
	fmt.Println("proof = ", proof)
	tryProof, err := testTrie.TryProve(testAddr1.Bytes())
	fmt.Println("proof = ", tryProof)

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

func ExecutorDemo() {
	fmt.Println("************************** 初始化 **************************")
	var err error

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

	fmt.Println("old root =", testRoot.Hex())

	// 设置一个假的状态树testRoot跟开始创建状态树testTrie
	testTrie, _ := trie.NewTrie(testRoot, testDB)

	fmt.Println("************************** 执行交易 **************************")
	// 初始化
	// accountInit(testTrie)
	showState(testTrie)

	// 生成测试交易
	txSet, err := executor.TxSetGenerate(10, 1, 1)
	if err != nil {
		fmt.Println("Tx Set Generate Fail!", err.Error())
	}

	for i := 0; i < len(txSet); i++ {
		txJSON, _ := json.Marshal(*txSet[i])
		fmt.Println(i, " = ", txJSON)
	}
	fmt.Println("")

	// 执行测试交易
	result, err := executor.TxsExecute(testTrie, txSet)
	if err != nil {
		fmt.Println("Tx Set Execute Fail!", err.Error())
	}
	for i := 0; i < executor.UserRange; i++ {
		txJSON, _ := json.Marshal(*txSet[i])
		fmt.Println(i, " = ", string(txJSON))
	}
	fmt.Println("")

	// 转化result为json
	resultJSON, err := json.Marshal(*result)
	fmt.Println("resultJSON = ", string(resultJSON))

	// 查看结果
	showState(testTrie)

	fmt.Println("************************** 更新数据库 **************************")
	// 更新到数据库
	root, err := testTrie.Commit(nil)
	if err != nil {
		fmt.Println("ERROR: Trie Root Update Fail!", err)
	}

	// 更新状态树根
	err = testDB.Put([]byte(MyStateRoot), root.Bytes(), nil)
	if err != nil {
		fmt.Println("ERROR: Trie Root Update Fail!", err)
	}
	fmt.Println("new root =", root.Hex())
	testDB.Close()
}

// 合约操作
//func storageTest(client *ethclient.Client, ctx context.Context, auth *bind.TransactOpts) {
//	// storage test
//	scAddress := common.HexToAddress(storageAddr)
//	scStorage, _ := NewStorage(scAddress, client)
//
//	// set string
//	tx, err := scStorage.SetShowStr(auth, "this is my test 3")
//	if err != nil {
//		log.Fatal(err)
//	}
//	txStr, err := tx.MarshalJSON()
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println("txStr = ", string(txStr))
//
//	// get String
//	copt := &bind.CallOpts{}
//	res, err := scStorage.GetShowRes(copt)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println("res = ", res)
//
//	// get abi respones
//	abiRes, err := scStorage.GetInfo(copt)
//	fmt.Println("abi res = ", string(abiRes))
//}

//func abiTest() {
//	fmt.Println("***** connect the eth node *****")
//	client, err := ethclient.Dial("ws://172.18.166.229:8546")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Println(client, "successfully connected")
//
//	// file := "../../data/keystore/UTC-USER2--2023-09-10T10-46-57.143272940Z--5e6aff8b8f168a52e2b8c2882442aa51482292f8"
//	file := "./data/UTC-USER2"
//	jsonBytes, err := os.ReadFile(file)
//	// password, _ := os.ReadFile("../../data/password.txt")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	ctx := context.Background()
//	chainID, err := client.ChainID(ctx)
//	if err != nil {
//		log.Fatal(err)
//	}
//	// fmt.Println("chainID = ", chainID)
//	// Key, _ := keystore.DecryptKey(jsonBytes, string(password))
//	Key, _ := keystore.DecryptKey(jsonBytes, myPassword)
//	height, _ := client.BlockNumber(ctx)
//	//blockNum, _ := client.BlockNumber(ctx)
//	fmt.Println(Key, chainID, height)
//
//	// contract call init
//	// test the contract
//	fmt.Println("***** Contract Caller Init *****")
//	caller := &bind.CallOpts{}
//	auth, err := bind.NewTransactorWithChainID(strings.NewReader(string(jsonBytes)), myPassword, chainID)
//	if err != nil {
//		log.Fatal(err)
//	}
//	auth.GasPrice = big.NewInt(1)
//	auth.GasLimit = uint64(2100000)
//
//	// build the DataTypes instance
//	DTAddr := common.HexToAddress(DataTypesAddr)
//	DT, _ := contract.NewDataTypes(DTAddr, client)
//
//	// build the MerkleUtils instance
//	MUAddr := common.HexToAddress(MerkleUtilsAddr)
//	MU, _ := contract.NewMerkelUtils(MUAddr, client)
//
//	// build the LSM instance
//	LSMAddr := common.HexToAddress(LocalStateManagerAddr)
//	LSM, _ := contract.NewLocalStateManager(LSMAddr, client)
//
//	// test Merkel Tree function
//	fmt.Println("***** use transition verify the root *****")
//	var toTran []contract.DataTypesLocalTransition
//	var tran1 = contract.DataTypesLocalTransition{
//		ChainID:    big.NewInt(1),
//		BlockNum:   big.NewInt(1),
//		Account:    "test 2",
//		Value:      big.NewInt(10),
//		Style:      big.NewInt(2),
//		UnlockTime: big.NewInt(1),
//	}
//	var tran2 = contract.DataTypesLocalTransition{
//		ChainID:    big.NewInt(1),
//		BlockNum:   big.NewInt(1),
//		Account:    "test 1",
//		Value:      big.NewInt(10),
//		Style:      big.NewInt(2),
//		UnlockTime: big.NewInt(1),
//	}
//	toTran = append(toTran, tran1, tran2)
//
//	ltByte1, err := DT.LocalTransitionToBytes(caller, tran1)
//	fmt.Println("ltByte = ", ltByte1)
//	ltByte2, err := DT.LocalTransitionToBytes(caller, tran2)
//	fmt.Println("ltByte = ", ltByte2)
//	testRoot, err := MU.GetMerkleRoot(caller, [][]byte{ltByte1, ltByte2})
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	succ, res, err := LSM.CheckRoot(caller, toTran, testRoot)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println("succ = ", succ)
//	fmt.Println("res = ", res)
//
//}

func TestLocalStateManager(client *ethclient.Client, ctx context.Context) {
	// file := "../../data/keystore/UTC-USER2--2023-09-10T10-46-57.143272940Z--5e6aff8b8f168a52e2b8c2882442aa51482292f8"
	file := "./data/UTC-USER2"
	jsonBytes, err := os.ReadFile(file)
	chainID, err := client.ChainID(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("***** test the contract *****")
	tx1, pendRes, err := client.TransactionByHash(ctx, common.HexToHash("0x1411db158885d8aafc5db7d6d7ec1bdb7a877e3f49275c7fa4ed5605a2a5f367"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("is pending = ", pendRes)
	fmt.Println("gas = ", tx1.Gas())

	// get the block
	block, err := client.BlockByNumber(ctx, big.NewInt(428))
	resTx := block.Transaction(common.HexToHash("0x86ff2d6e65e19ca7e9437fd187b5a5ce0960e59e6fb651f6bfa0a3cf876ab871"))
	if resTx == nil {
		fmt.Println("res tx =  nothing")
	} else {
		fmt.Println("res tx = ", resTx.Hash())
	}

	// test the contract
	fmt.Println("***** test the contract *****")
	auth, err := bind.NewTransactorWithChainID(strings.NewReader(string(jsonBytes)), myPassword, chainID)
	if err != nil {
		log.Fatal(err)
	}
	auth.GasPrice = big.NewInt(1)
	auth.GasLimit = uint64(2100000)

	// build the LSM instance
	LSMAddr := common.HexToAddress(LocalStateManagerAddr)
	LSM, _ := contract.NewLocalStateManager(LSMAddr, client)

	// deposit
	fmt.Println("***** deposit to test 2 *****")
	tx, err := LSM.DepositCoin(auth, "test 2", big.NewInt(10))
	if err != nil {
		log.Fatal(err)
	}
	txStr, _ := tx.MarshalJSON()
	fmt.Println("DepositCoin tx = ", string(txStr))

	err = GetOneLocalStateNotify(client)
	if err != nil {
		log.Fatal(err)
	}

	//// check the value
	fmt.Println("***** check balance of test 2 *****")
	caller := &bind.CallOpts{}
	state1, err := LSM.GetLocalState(caller, "test 1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("local state = ", state1.Account, state1.Value, state1.Lock)

	// use transition
	//fmt.Println("***** verify and use transition *****")
	//var toTran []contract.DataTypesLocalTransition
	//var tran1 = contract.DataTypesLocalTransition{
	//	ChainID:    big.NewInt(1),
	//	BlockNum:   big.NewInt(1),
	//	Account:    "test 2",
	//	Value:      big.NewInt(10),
	//	Style:      big.NewInt(2),
	//	UnlockTime: big.NewInt(1),
	//}
	//var tran2 = contract.DataTypesLocalTransition{
	//	ChainID:    big.NewInt(1),
	//	BlockNum:   big.NewInt(1),
	//	Account:    "test 1",
	//	Value:      big.NewInt(10),
	//	Style:      big.NewInt(2),
	//	UnlockTime: big.NewInt(1),
	//}
	//toTran = append(toTran, tran1, tran2)
	//tx2, err := LSM.VerifyAndLockLocalState(auth, toTran)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//txStr2, _ := tx2.MarshalJSON()
	//fmt.Println("VerifyAndLockLocalState tx = ", string(txStr2))

	// confirm the transition
	//fmt.Println("***** confirm the transition *****")
	//
	//lockNum, err := LSM.LockTransitionNum(&bind.CallOpts{})
	//fmt.Println("lock state num = ", lockNum.Int64())
	//
	//transition, err := LSM.LockTransitionPool(&bind.CallOpts{}, big.NewInt(1))
	//fmt.Println("transition = ", transition)
	//
	//blockNum, _ := client.BlockNumber(ctx)
	//tx3, err := LSM.ConfirmLocalState(auth, big.NewInt(int64(blockNum)+10))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("blockNum = ", big.NewInt(int64(blockNum)+10).Int64())
	//txStr3, _ := tx3.MarshalJSON()
	//fmt.Println("ConfirmLocalState tx = ", string(txStr3))
}

// GetOneLocalStateNotify get a local state event
func GetOneLocalStateNotify(client *ethclient.Client) error {
	// listen the event
	lsmAddress := common.HexToAddress(LocalStateManagerAddr)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{lsmAddress},
	}
	logs := make(chan types.Log)

	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	// 获取事件哈希值
	LocalStateNotifyHash := crypto.Keccak256Hash([]byte("LocalStateNotify(string,string,uint256)"))
	ErrorNotifyHash := crypto.Keccak256Hash([]byte("ErrorNotify(uint256,string)"))
	fmt.Println("LocalStateNotifyHash = ", LocalStateNotifyHash)
	fmt.Println("ErrorNotifyHash = ", ErrorNotifyHash)

	// build filter  and iterator
	lsmFilter, err := contract.NewLocalStateManagerFilterer(lsmAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	// begin listening
	fmt.Println("LISTENING FOR ONE EVENT:")
	select {
	case err := <-sub.Err():
		log.Fatal(err)
	case vLog := <-logs:
		//fmt.Println("***** receive an event *****")
		fmt.Printf("Log Block Number: %d\n", vLog.BlockNumber)
		fmt.Printf("Log Index: %d\n", vLog.Index)
		fmt.Printf("Log Topic: %s\n", vLog.Topics[0].Hex())

		switch vLog.Topics[0].Hex() {
		case LocalStateNotifyHash.Hex():
			// process the local state notify event
			fmt.Printf("Log Name: LocalStateNotify{\n")

			lsmNotify, err := lsmFilter.ParseLocalStateNotify(vLog)
			if err != nil {
				log.Fatal(err)
			}

			//NoTokenOwner := common.HexToAddress(vLog.Topics[1].Hex())
			//Spender := common.HexToAddress(vLog.Topics[2].Hex())

			fmt.Printf("    Account: %s\n", lsmNotify.Account)
			fmt.Printf("    Operation Style: %s\n", lsmNotify.Style)
			fmt.Printf("    Value: %s\n", lsmNotify.Value)
			fmt.Println("}")

			return nil
		case ErrorNotifyHash.Hex():
			// process the error notify event
			fmt.Printf("Log Name: LocalStateNotify\n")

			errNotify, err := lsmFilter.ParseErrorNotify(vLog)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("   error State:", errNotify.State)
			fmt.Println("   error ErrorString", errNotify.Error)
			fmt.Println("}")
			return errors.New("get an error Event")
		default:
			return errors.New("Other type of event")
		}
	}

	return nil
}

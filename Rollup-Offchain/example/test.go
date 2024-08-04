package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/big"
	"os"
	"rollup-offchain/contract"
	"rollup-offchain/executor"
	"rollup-offchain/model"
	"rollup-offchain/trie"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var TestTxNum = 5

// use your own password
var myPassword = "123456789"

//
//var DataTypesAddr = " 0x1C4632eA7A9942cF70FAbBd6fD5346eCC6Ef44F2"
//var MerkleUtilsAddr = " 0x22B472414887C98634240e38dE86a19aA07f0887"
//var LocalStateManagerAddr = " 0x87a7f3455560A20b18fa4423E350e3212B4Dd7b4"
//var MPTVerifierAddr = " 0x59Fb245CD187265D254004022366a7Bc73C9451f"
//var TransactionExecutorAddr = " 0x7C9Acb5d7143A638ACd18be690Da8f71AB913DcC"
//var CrossRollupAddr = " 0xa73A3C322c4C81909D331a6d9e8733d95a5EFd4b"
//
//var MyStateRoot = "stateRoot"

func CommiteTestMain() {
	//// 初始化客户端(use your own IP address)
	fmt.Println("***** Client Init *****")
	client, err := ethclient.Dial("ws://xxx.xx.xxx.xxx:8545")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// 获取chainID
	ctx := context.Background()
	chainID, err := client.ChainID(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ChainID = ", chainID.Int64())

	// 读取密钥文件
	file := "../data/UTC-USER2"
	jsonBytes, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	// 创建执行上下文
	caller := &bind.CallOpts{}
	auth, err := bind.NewTransactorWithChainID(strings.NewReader(string(jsonBytes)), myPassword, chainID)
	if err != nil {
		log.Fatal(err)
	}
	auth.GasPrice = big.NewInt(1)
	auth.GasLimit = uint64(8000000)

	//// 初始化日志功能
	//fmt.Println("***** Log Init *****")
	//OnChainLogFile, err := os.OpenFile(OnChainLogPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	//if err != nil {
	//	log.Fatalln("Failed to open error log file: ", err)
	//}
	//OffChainLogFile, err := os.OpenFile(OffChainLogPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	//if err != nil {
	//	log.Fatalln("Failed to open error log file: ", err)
	//}
	//logInit(OnChainLogFile, OffChainLogFile)
	//costlog.CostLogInit(OnChain, OffChain)
	//defer OnChainLogFile.Close()
	//defer OffChainLogFile.Close()

	fmt.Println("************************** 执行器初始化 **************************")
	// init contract
	executor.RollupInit(client, common.HexToAddress(contract.USER1_Addr), big.NewInt(1))

	fmt.Println("************************** 初始化账户状态 **************************")
	// 初始化账户状态
	//L2AccountInit(client, auth, ctx)

	// 查询account
	addrInt := big.NewInt(1)
	addrHash := trie.BigToHash(addrInt)
	var tmpAddr [32]byte
	tmpAddr = addrHash
	ls, err := executor.LSMInstance.GetLocalState(caller, string(tmpAddr[:]))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("#### account = ", ls.Account, ", balance = ", ls.Value)

	fmt.Println("************************** commit 测试 **************************")
	//CommitRandTxs(16, client, auth, caller, ctx)

	fmt.Println("************************** exe challenge 测试 **************************")
	//ExeChallenge(big.NewInt(4), client, auth, caller, ctx)

	fmt.Println("************************** 交易查询 **************************")
	// 1-611410 5 - 2348460, 10 - 4467392
	// search the commit transaction
	txHash := "0x0833948cdc18b8d3c7cf04c950f8c403483a1216cd972de05fd6c9e2fbdc477c"
	receipt, err := client.TransactionReceipt(ctx, common.HexToHash(txHash))
	if err != nil {
		log.Fatal(err)
	}
	rJson, _ := receipt.MarshalJSON()
	fmt.Println("####\nreceipt str = ", string(rJson))

	// // search target block
	//blk, _ := client.BlockByNumber(ctx, big.NewInt(128))
	//blkTxs := blk.Transactions()
	//fmt.Println(blkTxs.Len())
	//
	//fmt.Println(blkTxs[1].Hash().Hex())

}

func ExeChallenge(index *big.Int, client *ethclient.Client, auth *bind.TransactOpts, caller *bind.CallOpts, ctx context.Context) {
	fileName := "Test-Tx-" + index.String()
	txSet := TestTxRead(fileName)
	fmt.Println("fileTx len = ", len(txSet))

	//transits := executor.ExeCrossTxSet(txSet)

	tmpTxs := model.Local2DtTxList(txSet)
	//tmpTss := model.Local2DtTsList(transits)

	tx, err := executor.CRInstance.ExecutionChallenge(auth, index, tmpTxs)
	if err != nil {
		return
	}
	fmt.Println("exe challenge tx hash  = ", tx.Hash())
}

func CommitRandTxs(txNum int, client *ethclient.Client, auth *bind.TransactOpts, caller *bind.CallOpts, ctx context.Context) {

	// 生成测试交易
	txSet, err := executor.CrossTxSetGenerate(txNum)
	if err != nil {
		fmt.Println("Tx Set Generate Fail!", err.Error())
	}

	// 打印交易信息
	//for i := 0; i < len(txSet); i++ {
	//	txJSON, _ := json.Marshal(txSet[i])
	//	fmt.Println(i, " = ", string(txJSON))
	//}
	//fmt.Println("")

	// 执行测试交易
	transits := executor.ExeCrossTxSet(txSet)
	if err != nil {
		fmt.Println("Tx Set Execute Fail!", err.Error())
	}
	// 打印执行结果
	for i := 0; i < len(transits); i++ {
		strJSON, _ := json.Marshal(transits[i])
		fmt.Println(i, " = ", string(strJSON))
	}
	fmt.Println("")

	// 提交执行结果上链
	chain_id, err := executor.LSMInstance.ChainID(caller)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("chain id = ", chain_id.Int64())

	// commit tx
	tmpTxs := model.Local2DtTxList(txSet)
	tmpTss := model.Local2DtTsList(transits)

	// filter the transition with id=1
	var tmpLocalTss []contract.DataTypesL2Transition
	for i := 0; i < len(tmpTss); i++ {
		if tmpTss[i].ChainID == big.NewInt(1) {
			tmpLocalTss = append(tmpLocalTss, tmpTss[i])
		}
	}

	fmt.Println("data len = ", len(tmpTxs), len(tmpTss))

	blkLen, err := executor.CRInstance.BlockLen(caller)
	txRoot, err := executor.DTInstance.TransactionsListHash(caller, tmpTxs)
	tsRoot, err := executor.DTInstance.L2TransitionListHash(caller, tmpTss)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("block data = ", blkLen, txRoot, tsRoot)

	// execute the commit
	//tx, err := executor.LSMInstance.VerifyAndLockLocalState(auth, tmpTss, big.NewInt(1), big.NewInt(2))
	//tx, err := executor.CRInstance.CommitBlock(auth, big.NewInt(blkLen.Int64()), txRoot, tsRoot, tmpTxs, tmpTss)
	tx, err := executor.CRInstance.CommitBlock(auth, big.NewInt(blkLen.Int64()), txRoot, tsRoot, tmpTxs, tmpLocalTss)
	if err != nil {
		return
	}
	fmt.Println("commit tx hash  = ", tx.Hash())

	// write transaction
	fileName := "Test-Tx-" + blkLen.String()
	fmt.Println("fileName = ", fileName)

	TestTxWrite(fileName, txSet)
}

// L2AccountInit 初始化链上账户状态（确保资金充足）
func L2AccountInit(client *ethclient.Client, auth *bind.TransactOpts, ctx context.Context) {
	// 让每个可能账户都包含100的coin
	fmt.Println("\n account list")
	for i := int64(0); i <= int64(executor.UserRange); i++ {
		addrInt := big.NewInt(i)
		addrHash := trie.BigToHash(addrInt)
		var tmpAddr [32]byte
		tmpAddr = addrHash

		valueInt := big.NewInt(1000)

		fmt.Println(" - addrHash = ", string(tmpAddr[:]), "; value", valueInt.String())

		resTx, err := executor.LSMInstance.DepositCoin(auth, string(tmpAddr[:]), valueInt)
		if err != nil {
			log.Fatal(err)
		}
		txStr, err := resTx.MarshalJSON()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("txStr = ", string(txStr))
	}

}

// TestTxWrite write L2 transaction
func TestTxWrite(fileName string, txs []model.L2Transaction) error {
	openFile, e := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 777)
	if e != nil {
		fmt.Println(e)
	}

	for i := 0; i < len(txs); i++ {
		jsonStr, err := json.Marshal(txs[i])
		if err != nil {

		}
		openFile.WriteString(string(jsonStr) + "\n")
	}

	openFile.Close()
	return nil
}

// TestTxRead read L2 transaction
func TestTxRead(fileName string) []model.L2Transaction {
	var res []model.L2Transaction

	openFile, e := os.OpenFile(fileName, os.O_RDWR, 777)
	if e != nil {
		fmt.Println(e)
	}

	//buf := make([]byte, 1024)
	//for {
	//	len, _ := openFile.Read(buf)
	//	if len == 0 {
	//		break
	//	}
	//	fmt.Println(string(buf))
	//}
	//openFile.Close()

	reader := bufio.NewReader(openFile)
	for {
		line, _, e := reader.ReadLine()
		if e == io.EOF {
			break
		}
		if e != nil {
			log.Fatal(e)
		}
		fmt.Println(string(line))

		var tmpTx model.L2Transaction
		err := json.Unmarshal(line[:], &tmpTx)
		if err != nil {
			log.Fatal(err)
		}

		res = append(res, tmpTx)
	}
	openFile.Close()

	return res
}

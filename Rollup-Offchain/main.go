package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	// "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"io"
	"log"
	"math/big"
	"os"
	"rollup-offchain/contract"
	"rollup-offchain/cosmos_executor"
	"rollup-offchain/costlog"
	"rollup-offchain/executor"
	"rollup-offchain/model"
	"strconv"
	"strings"
	"time"
)

var TestTxNum = 5

var myPassword = "123456789"

var DataTypesAddr = " 0x1C4632eA7A9942cF70FAbBd6fD5346eCC6Ef44F2"
var MerkleUtilsAddr = " 0x22B472414887C98634240e38dE86a19aA07f0887"
var LocalStateManagerAddr = " 0x87a7f3455560A20b18fa4423E350e3212B4Dd7b4"
var MPTVerifierAddr = " 0x59Fb245CD187265D254004022366a7Bc73C9451f"
var TransactionExecutorAddr = " 0x7C9Acb5d7143A638ACd18be690Da8f71AB913DcC"
var CrossRollupAddr = " 0xa73A3C322c4C81909D331a6d9e8733d95a5EFd4b"

var MyStateRoot = "stateRoot"

var OnChain *log.Logger
var OffChain *log.Logger
var OnChainLogPath = "./OnChainInfo-mos-4.txt"
var OffChainLogPath = "./OffChainInfo-mos-4.txt"
var TxLogPath = "./TxLogInfo.txt"

// 通信所需
var err error
var EthClient *ethclient.Client
var EthAuth *bind.TransactOpts
var MosClient *ethclient.Client
var MosAuth *bind.TransactOpts

var MosRpcUrl = "ws://172.18.166.229:8646"
var MosFullRollupAddr = "0x03D389c532e2b0Aa33Cd302f1e147843245B177a"
var MosFromAddr = "0x7cb61d4117ae31a12e393a1cfa3bac666481d02e"
var MosMnemonic = "gesture inject test cycle original hollow east ridge hen combine junk child bacon zero hope comfort vacuum milk pitch cage oppose unhappy lunar seat"
var MosPrivateKey = "e9b1d63e8acd7fe676acb43afb390d4b0202dab61abec9cf2a561e4becb147de"

var EthRpcURL = "ws://172.18.166.229:8546"
var EthFullRollupAddr = "0x8CfcA745E804Fa454cE841754BC0c4D1ae2dc35F"
var EthFromAddr = "0x0bc09f70C683B83E784704331D26F6a244274b93"
var EthMyPassword = "123456789"

func logInit(onChainFile *os.File, offChainFile *os.File) {
	OnChain = log.New(os.Stdout, "OnChain: ", log.Ldate|log.Ltime|log.Lshortfile)
	OnChain.SetOutput(onChainFile)

	OffChain = log.New(os.Stdout, "OffChain: ", log.Ldate|log.Ltime|log.Lshortfile)
	OffChain.SetOutput(offChainFile)
}

func ClientInit(chainType int) {
	ctx := context.Background()

	// if chainType == 0 {
		// Ethereum
		EthClient, err = ethclient.Dial(EthRpcURL)
		if err != nil {
			log.Fatal(err)
		}
		chainID, err := EthClient.ChainID(ctx)
		if err != nil {
			log.Fatal(err)
		}

		file := "./data/UTC-USER2"
		jsonBytes, err := os.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}

		EthAuth, err = bind.NewTransactorWithChainID(strings.NewReader(string(jsonBytes)), myPassword, chainID)
		if err != nil {
			log.Fatal(err)
		}
		EthAuth.GasPrice = big.NewInt(1)
		EthAuth.GasLimit = uint64(8000000)
	// } else {
	// 	//
	// 	// Cosmos
	// 	//
	// 	fmt.Println("***** Client Init Begin *****")
	// 	MosClient, err = ethclient.Dial(MosRpcUrl)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	fmt.Println("***** Client Init Success *****")

	// 	chainID, err := MosClient.ChainID(ctx)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	fmt.Println(chainID)

	// 	private_key, err := crypto.HexToECDSA(MosPrivateKey)
	// 	MosAuth, err = bind.NewKeyedTransactorWithChainID(private_key, chainID)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	MosAuth.GasLimit = uint64(8000000)
	// 	MosAuth.GasPrice = big.NewInt(10)
	// }

}

func testMos() {
	//// 初始化客户端
	fmt.Println("***** Client Init *****")
	ClientInit(1)

	defer MosClient.Close()
	//defer EthClient.Close()

	//EthClient, err = ethclient.Dial("ws://172.18.166.229:8546")
	//if err != nil {
	//	log.Fatal(err)
	//}

	// 必要的赋值
	client := MosClient
	ctx := context.Background()
	caller := &bind.CallOpts{}

	// 获取chainID

	chainID, err := client.ChainID(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ChainID = ", chainID.Int64())

	// 初始化日志功能
	fmt.Println("***** Log Init *****")
	OnChainLogFile, err := os.OpenFile(OnChainLogPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open error log file: ", err)
	}
	OffChainLogFile, err := os.OpenFile(OffChainLogPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open error log file: ", err)
	}
	logInit(OnChainLogFile, OffChainLogFile)
	costlog.CostLogInit(OnChain, OffChain)
	defer OnChainLogFile.Close()
	defer OffChainLogFile.Close()

	fmt.Println("************************** 执行器初始化 **************************")
	// init contract
	// executor.RollupInit(client, common.HexToAddress(contract.USER1_Addr), big.NewInt(1))
	cosmos_executor.RollupInit(client, common.HexToAddress(contract.USER1_Addr), big.NewInt(1))

	fmt.Println("************************** 初始化账户状态 **************************")
	// 初始化账户状态
	//L2AccountInit(client, auth, ctx)

	// 查询account
	addrInt := big.NewInt(1032)
	addr := common.BigToAddress(addrInt)
	start := time.Now()
	ls, err := cosmos_executor.FRInstance.GetLocalState(caller, addr)
	if err != nil {
		//fmt.Println("here?")
		fmt.Println(err)
		log.Fatal(err)
	}
	fmt.Println("#### cost time = ", time.Now().Sub(start))
	fmt.Println("#### account = ", ls.Account, ", balance = ", ls.Value)

	fmt.Println("************************** commit 测试 **************************")
	cnts := [9]int{1, 4, 8, 12, 16, 20, 24, 28, 32}
	for k := 0; k < 5; k ++ {
		for i := 0; i < 9; i++ {
			time.Sleep(4 * time.Second)
			CommitRandTxs(1, cnts[i], client, MosAuth, caller, ctx)
		}
	}

	fmt.Println("************************** exe challenge 测试 **************************")
	//ExeChallenge(big.NewInt(5), client, auth, caller, ctx)

	fmt.Println("************************** roll back 测试 **************************")
	//RollbackBlock(big.NewInt(5), client, auth, caller, ctx)

	fmt.Println("************************** 交易查询 **************************")
	// search the commit transaction
	txHash := "0xaec73eb7dfec2305a0c0f73242f93ab7f979d1f4bcb8e666f5075a8caf5c5656"
	receipt, err := client.TransactionReceipt(ctx, common.HexToHash(txHash))
	if err != nil {
		log.Fatal(err)
	}
	// rJson, _ := receipt.MarshalJSON()
	// fmt.Println("####\nreceipt str = ", rJson)
	fmt.Println("####\nreceipt gas used = ", receipt.GasUsed)

	// // search target block
	//blk, _ := client.BlockByNumber(ctx, big.NewInt(128))
	//blkTxs := blk.Transactions()
	//fmt.Println(blkTxs.Len())
	//
	//fmt.Println(blkTxs[1].Hash().Hex())

}

func main() {

	//// 初始化客户端
	fmt.Println("***** Client Init *****")
	ClientInit(0)

	//defer MosClient.Close()
	defer EthClient.Close()

	//EthClient, err = ethclient.Dial("ws://172.18.166.229:8546")
	//if err != nil {
	//	log.Fatal(err)
	//}

	// 必要的赋值
	client := EthClient
	ctx := context.Background()
	caller := &bind.CallOpts{}

	// 获取chainID

	chainID, err := client.ChainID(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ChainID = ", chainID.Int64())

	// 初始化日志功能
	fmt.Println("***** Log Init *****")
	OnChainLogFile, err := os.OpenFile(OnChainLogPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open error log file: ", err)
	}
	OffChainLogFile, err := os.OpenFile(OffChainLogPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open error log file: ", err)
	}
	logInit(OnChainLogFile, OffChainLogFile)
	costlog.CostLogInit(OnChain, OffChain)
	defer OnChainLogFile.Close()
	defer OffChainLogFile.Close()

	fmt.Println("************************** 执行器初始化 **************************")
	// init contract
	executor.RollupInit(client, common.HexToAddress(contract.USER1_Addr), big.NewInt(1))
	//cosmos_executor.RollupInit(client, common.HexToAddress(contract.USER1_Addr), big.NewInt(1))

	fmt.Println("************************** 初始化账户状态 **************************")
	// 初始化账户状态
	//L2AccountInit(client, auth, ctx)
	// L2AccountInit(0, client, EthAuth, ctx)
	// for i:= 0; i < 10; i ++:
	// 	time.Sleep(1*time.Second)
	// 	L2AccountInit(0, client, EthAuth, ctx)

	// 查询account
	addrInt := big.NewInt(1032)
	addr := common.BigToAddress(addrInt)
	start := time.Now()
	ls, err := executor.FRInstance.GetLocalState(caller, addr)
	if err != nil {
		fmt.Println("here?")
		fmt.Println(err)
		log.Fatal(err)
	}
	fmt.Println("#### cost time = ", time.Now().Sub(start))
	fmt.Println("#### account = ", ls.Account, ", balance = ", ls.Value)

	fmt.Println("************************** commit 测试 **************************")
	// cnts := [9]int{1, 4, 8, 12, 16, 20, 24, 28, 32}
	// for i := 0; i < 9; i++ {
	// 	time.Sleep(10 * time.Second)
	// 	CommitRandTxs(0, cnts[i], client, EthAuth, caller, ctx)
	// }
	// CommitRandTxs(0, 4, client, EthAuth, caller, ctx)

	fmt.Println("************************** exe challenge 测试 **************************")
	//ExeChallenge(big.NewInt(5), client, auth, caller, ctx)

	fmt.Println("************************** roll back 测试 **************************")
	//RollbackBlock(big.NewInt(5), client, auth, caller, ctx)

	fmt.Println("************************** 交易查询 **************************")
	// search the commit transaction
	txHash := "0xaec73eb7dfec2305a0c0f73242f93ab7f979d1f4bcb8e666f5075a8caf5c5656"
	receipt, err := client.TransactionReceipt(ctx, common.HexToHash(txHash))
	if err != nil {
		log.Fatal(err)
	}
	// rJson, _ := receipt.MarshalJSON()
	// fmt.Println("####\nreceipt str = ", rJson)
	fmt.Println("####\nreceipt gas used = ", receipt.GasUsed)

	// // search target block
	//blk, _ := client.BlockByNumber(ctx, big.NewInt(128))
	//blkTxs := blk.Transactions()
	//fmt.Println(blkTxs.Len())
	//
	//fmt.Println(blkTxs[1].Hash().Hex())

}

func CommitRandTxs(chainType int, txNum int, client *ethclient.Client, auth *bind.TransactOpts, caller *bind.CallOpts, ctx context.Context) {

	var FRInstance *contract.FullRollup
	var DTInstance *contract.DataTypes

	if chainType == 0 {
		FRInstance = executor.FRInstance
		DTInstance = executor.DTInstance
	} else {
		FRInstance = cosmos_executor.FRInstance
		DTInstance = cosmos_executor.DTInstance
	}

	// 生成测试交易
	txSet, err := executor.CrossTxSetGenerate(txNum)
	if err != nil {
		fmt.Println("Tx Set Generate Fail!", err.Error())
	}

	start := time.Now()

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

	fmt.Println("Tx set generate success")

	// 提交执行结果上链
	chain_id, err := FRInstance.ChainID(caller)
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
		if tmpTss[i].ChainID.Int64() == 1 {
			tmpLocalTss = append(tmpLocalTss, tmpTss[i])
		}
	}

	fmt.Println("data len = ", len(tmpTxs), len(tmpTss), len(tmpLocalTss))

	var txDatas [][]byte
	var finalTxsStr []byte
	for i := 0; i < len(tmpTxs); i++ {
		tmpStr, err := DTInstance.L2TransactionToBytes(caller, tmpTxs[i])
		if err != nil {
			log.Fatal(err)
		}
		txDatas = append(txDatas, tmpStr)
		finalTxsStr = append(finalTxsStr, tmpStr...)
	}

	blkLen, err := FRInstance.BlockLen(caller)
	txRoot, err := DTInstance.BytesList2Hash(caller, txDatas)
	tsRoot, err := DTInstance.L2TransitionListHash(caller, tmpTss)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("block data = ", blkLen, txRoot, tsRoot)

	// execute the commit
	//tx, err := executor.LSMInstance.VerifyAndLockLocalState(auth, tmpTss, big.NewInt(1), big.NewInt(2))
	//tx, err := executor.CRInstance.CommitBlock(auth, big.NewInt(blkLen.Int64()), txRoot, tsRoot, tmpTxs, tmpTss)
	// var tx *types.Transaction
	// if chainType == 0 {
	// 	tx, err = executor.FRInstance.CommitBlock(auth, big.NewInt(blkLen.Int64()), txRoot, tsRoot, finalTxsStr, tmpLocalTss)
	// 	if err != nil {
	// 		return
	// 	}
	// } else {

	// }
	tx, err := FRInstance.CommitBlock(auth, big.NewInt(blkLen.Int64()), txRoot, tsRoot, finalTxsStr, tmpLocalTss)
	if err != nil {
		return
	}

	fmt.Println("commit tx hash  = ", tx.Hash())

	// write off-chain cost
	cost := time.Now().Sub(start)
	costlog.OffChainLog.Println(strconv.Itoa(chainType)+"\tCommitRandTxs\t"+tx.Hash().String()+"-"+strconv.Itoa(txNum), cost, start.UnixNano()/1e6)

	// write transaction
	fileName := "Test-Tx-" + blkLen.String()
	fmt.Println("fileName = ", fileName)

	TestTxWrite(fileName, txSet)
}

/*
*
*  以下内容暂不需要
*
 */
func RollbackBlock(index *big.Int, client *ethclient.Client, auth *bind.TransactOpts, caller *bind.CallOpts, ctx context.Context) {
	start := time.Now()

	fileName := "Test-Tx-" + index.String()
	txSet := TestTxRead(fileName)
	fmt.Println("fileTx len = ", len(txSet))

	//transits := executor.ExeCrossTxSet(txSet)
	// 执行交易并验证结果
	transits := executor.ExeCrossTxSet(txSet)
	tmpTss := model.Local2DtTsList(transits)

	tx, err := executor.FRInstance.RemoveLocalStateWithBlockNum(auth, index, tmpTss)
	if err != nil {
		return
	}
	fmt.Println("exe challenge tx hash  = ", tx.Hash())

	cost := time.Now().Sub(start)
	costlog.OffChainLog.Println("RollbackBlock-"+tx.Hash().String()+"-"+strconv.Itoa(len(txSet)), cost)
}

func ExeChallenge(index *big.Int, client *ethclient.Client, auth *bind.TransactOpts, caller *bind.CallOpts, ctx context.Context) {
	start := time.Now()

	fileName := "Test-Tx-" + index.String()
	txSet := TestTxRead(fileName)
	fmt.Println("fileTx len = ", len(txSet))

	//transits := executor.ExeCrossTxSet(txSet)
	// 执行交易并验证结果
	transits := executor.ExeCrossTxSet(txSet)
	tmpTss := model.Local2DtTsList(transits)
	_, err := executor.DTInstance.L2TransitionListHash(caller, tmpTss)

	tmpTxs := model.Local2DtTxList(txSet)
	//tmpTss := model.Local2DtTsList(transits)

	tx, err := executor.FRInstance.ExecutionChallenge(auth, index, tmpTxs)
	if err != nil {
		return
	}
	fmt.Println("exe challenge tx hash  = ", tx.Hash())

	cost := time.Now().Sub(start)
	costlog.OffChainLog.Println("ExeChallenge-"+tx.Hash().String()+"-"+strconv.Itoa(len(tmpTxs)), cost)

}

// L2AccountInit 初始化链上账户状态（确保资金充足）
func L2AccountInit(chainType int, client *ethclient.Client, auth *bind.TransactOpts, ctx context.Context) {
	// 让每个可能账户都包含100的coin
	fmt.Println("\n account list")

	num := int64(executor.UserRange)
	// num := int64(1)
	for i := int64(0); i <= num; i++ {
		if i%10 == 0 {
			time.Sleep(1 * time.Second)
		}

		addrInt := big.NewInt(i + 1000)
		addr := common.BigToAddress(addrInt)

		valueInt := big.NewInt(10)

		// fmt.Println(" - addrHash = ", addr.Hex(), "; value", valueInt.String())

		var err error
		var resTx *types.Transaction
		if chainType == 0 {
			resTx, err = executor.FRInstance.DepositCoin(auth, addr, valueInt)
		} else {
			resTx, err = cosmos_executor.FRInstance.DepositCoin(auth, addr, valueInt)
		}
		if err != nil {
			log.Fatal(err)
		}

		// txStr, err := resTx.MarshalJSON()
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// fmt.Println("txStr = ", string(txStr))
		fmt.Println("deposit tx hash  = ", resTx.Hash())
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

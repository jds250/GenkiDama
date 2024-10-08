package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"rollup-offchain/challenger"
	"rollup-offchain/contract"
	"rollup-offchain/costlog"
	"rollup-offchain/zk"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

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
var OnChainLogPath = "./OnChainInfo-bak.txt"
var OffChainLogPath = "./OffChainInfo-bak.txt"
var TxLogPath = "./TxLogInfo.txt"

func logInit(onChainFile *os.File, offChainFile *os.File) {
	OnChain = log.New(os.Stdout, "OnChain: ", log.Ldate|log.Ltime|log.Lshortfile)
	OnChain.SetOutput(onChainFile)

	OffChain = log.New(os.Stdout, "OffChain: ", log.Ldate|log.Ltime|log.Lshortfile)
	OffChain.SetOutput(offChainFile)
}

func main() {
	// 分析参数
	fmt.Println("***** Argument Analysis *****")
	for idx, args := range os.Args {
		fmt.Println("参数"+strconv.Itoa(idx)+":", args)
	}

	// 初始化客户端
	fmt.Println("***** Client Init *****")
	client, err := ethclient.Dial("ws://172.18.166.229:8546")
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
	file := "./data/UTC-USER2"
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
	auth.GasLimit = uint64(2100000)

	// 初始化zk电路：先读，如果读不到就自己生成
	fmt.Println("***** ZK Init *****")
	if os.Args[1] == "challenger" || os.Args[1] == "questioner" || os.Args[1] == "cbegin" {
		err = zk.Groth16Init()
		if err != nil {
			log.Fatal(err)
		}
	}

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
	defer OnChainLogFile.Close()
	defer OffChainLogFile.Close()

	// 初始化挑战者
	fmt.Println("***** Challenge Init *****")
	challenger.ChallangeInit(client, common.HexToAddress(contract.USER1_Addr), big.NewInt(1))

	// 根据参数创建监听者进程 或者 直接开始执行挑战和质疑
	fmt.Println("***** Function Start *****")
	myIndex := "1-23-6"

	if len(os.Args) < 2 {
		fmt.Println("Argument is Too Less!")
	}

	if len(os.Args) > 2 {
		myIndex = os.Args[2]
	}
	switch os.Args[1] {
	case "challenger":
		challenger.ChallengerListenerStart(client, auth, caller, ctx)
		break
	case "questioner":
		challenger.QuestionListenerStart(client, auth, caller, ctx)
		break
	case "checkpoint":
		err := challenger.AddCheckPoint(client, auth, caller, ctx)
		// err := challenger.ChallengeUpdate(client, caller, myIndex)
		if err != nil {
			log.Fatal(err)
		}
		break
	case "cbegin":
		err := challengeBegin(client, auth, caller, ctx, myIndex)
		if err != nil {
			log.Fatal(err)
		}
		break
	case "qbegin":
		err := questionBegin(client, auth, caller, ctx, myIndex)
		if err != nil {
			log.Fatal(err)
		}
		break
	case "commit":
		err := challenger.CommitAllRecords(client, ctx)
		if err != nil {
			log.Fatal(err)
		}
		break
	case "log":
		costlog.GenerateTxCostLog(OnChainLogPath, TxLogPath, client, ctx)
		break
	case "user":
		err := userTmp(client, auth, caller, ctx, myIndex)
		if err != nil {
			log.Fatal(err)
		}
		break
	default:
		fmt.Println("please add argument: challenger questioner checkpoint cbegin qbegin")
		break

	}
}

// challengeBegin 创建一个挑战，然后等待质疑
func challengeBegin(client *ethclient.Client, auth *bind.TransactOpts, caller *bind.CallOpts, ctx context.Context, index string) error {
	fmt.Println("challengeBegin Func")

	err := challenger.ChallengeCreate(client, auth, index)
	if err != nil {
		return err
	}

	return nil
}

// questionBegin
func questionBegin(client *ethclient.Client, auth *bind.TransactOpts, caller *bind.CallOpts, ctx context.Context, index string) error {
	fmt.Println("questionBegin Func")

	err := challenger.QuestionCreate(client, auth, ctx, index)
	if err != nil {
		return err
	}

	return nil
}

// questionBegin
func userTmp(client *ethclient.Client, auth *bind.TransactOpts, caller *bind.CallOpts, ctx context.Context, index string) error {
	fmt.Println("userTmp Func")

	res, err := challenger.CMInstance.ChallengePool(caller, index)
	if err != nil {
		return err
	}

	var record challenger.ChallengeRecord
	record = res

	fmt.Println("Index", record.Index)
	fmt.Println("State", record.State.Int64())
	fmt.Println("Begin", record.Begin.Int64())
	fmt.Println("End", record.End.Int64())
	fmt.Println("Challenger", record.Challenger)
	fmt.Println("Questioner", record.Questioner)

	return nil
}

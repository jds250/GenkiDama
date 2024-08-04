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

func BCPTest() {
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
	if os.Args[1] == "challenger" || os.Args[1] == "questioner" {
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
	costlog.CostLogInit(OnChain, OffChain)
	defer OnChainLogFile.Close()
	defer OffChainLogFile.Close()

	// 初始化挑战者
	fmt.Println("***** Challenge Init *****")
	challenger.ChallangeInit(client, common.HexToAddress(contract.USER1_Addr), big.NewInt(1))

	// 根据参数创建监听者进程 或者 直接开始执行挑战和质疑
	fmt.Println("***** Function Start *****")
	myIndex := "1-2-7" // 每次标识符不同

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
		err := costlog.GenerateTxCostLog(OnChainLogPath, TxLogPath, client, ctx)
		if err != nil {
			log.Fatal(err)
		}
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

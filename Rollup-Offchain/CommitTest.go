package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"os"
	"rollup-offchain/contract"
	"rollup-offchain/executor"
	"rollup-offchain/trie"
	"strings"
)

func CommitTest() {
	//// 初始化客户端
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
	ls, err := executor.FRInstance.GetLocalState(caller, string(tmpAddr[:]))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("#### account = ", ls.Account, ", balance = ", ls.Value)

	fmt.Println("************************** commit 测试 **************************")
	//CommitRandTxs(4, client, auth, caller, ctx)

	fmt.Println("************************** exe challenge 测试 **************************")
	//ExeChallenge(big.NewInt(4), client, auth, caller, ctx)

	fmt.Println("************************** 交易查询 **************************")
	// 1-611410 5 - 2348460, 10 - 4467392
	// search the commit transaction
	txHash := "0x4e5ce96c1be7d8b200c3355023443bd9a0455f1266ad52fb7b14e50f49c32576"
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

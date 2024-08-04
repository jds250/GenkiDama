package main

import (
	"EventWatcher/contract"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func EventListenDemo() {
	fmt.Println("***** argument test *****")
	for idx, args := range os.Args {
		fmt.Println("参数"+strconv.Itoa(idx)+":", args)
	}

	fmt.Println("***** connect the eth node *****")
	client, err := ethclient.Dial("ws://172.18.166.229:8546")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

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
	fmt.Println(Key, chainID, height)

	fmt.Println("***** test the contract *****")
	auth, err := bind.NewTransactorWithChainID(strings.NewReader(string(jsonBytes)), myPassword, chainID)
	if err != nil {
		log.Fatal(err)
	}
	auth.GasPrice = big.NewInt(1)
	auth.GasLimit = uint64(8000000)

	lsmAddress := common.HexToAddress(LocalStateManagerAddr)
	crAddress := common.HexToAddress(CrossRollupAddr)
	// scStorage, _ := NewStorage(scAddress, client)

	// build the abi
	//lsmABI, err := abi.JSON(strings.NewReader(contract.LocalStateManagerMetaData.ABI))
	//if err != nil {
	//	log.Fatal(err)
	//}

	// listen the event
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
	ConfirmStateHash := crypto.Keccak256Hash([]byte("ConfirmState(uint256,uint256,string)"))
	RollupBlockNotifyHash := crypto.Keccak256Hash([]byte("RollupBlockNotify(uint256,string)"))
	fmt.Println("LocalStateNotifyHash = ", LocalStateNotifyHash)
	fmt.Println("ErrorNotifyHash = ", ErrorNotifyHash)
	fmt.Println("ConfirmStateHash = ", ConfirmStateHash)

	// build filter  and iterator
	lsmFilter, err := contract.NewLocalStateManagerFilterer(lsmAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	crFilter, err := contract.NewCrossRollupFilterer(crAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	//blockNum, _ := client.BlockNumber(ctx)
	//endBlock := blockNum + 100
	//filterOpt := bind.FilterOpts{Start: blockNum, End: &endBlock, Context: ctx}
	//LSEventIterator, err := lsmFilter.FilterLocalStateNotify(&filterOpt)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//topicLocalState := crypto.Keccak256Hash([]byte(""))
	//topicErrorNotify := crypto.Keccak256Hash([]byte("eventName"))

	// 获取数据

	// begin listening
	fmt.Println("LISTENING STORAGE EVENT:")
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println("***** receive an event *****")
			fmt.Printf("Log Block Number: %d\n", vLog.BlockNumber)
			fmt.Printf("Log Index: %d\n", vLog.Index)
			fmt.Printf("Log Topic: %s\n", vLog.Topics[0].Hex())
			fmt.Printf("Log Detail: \n{\n")

			switch vLog.Topics[0].Hex() {
			case RollupBlockNotifyHash.Hex():
				// process the local state notify event
				fmt.Printf("Log Name: RollupBlockNotifyHash\n")

				rbNotify, err := crFilter.ParseRollupBlockNotify(vLog)
				if err != nil {
					log.Fatal(err)
				}

				//NoTokenOwner := common.HexToAddress(vLog.Topics[1].Hex())
				//Spender := common.HexToAddress(vLog.Topics[2].Hex())

				fmt.Printf("    BlockID: %s\n", rbNotify.BlockNumber)
				fmt.Printf("    Desc: %s\n", rbNotify.Desc)
				fmt.Println("}")
				break
			case LocalStateNotifyHash.Hex():
				// process the local state notify event
				fmt.Printf("Log Name: LocalStateNotify\n")

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
				break
			case ConfirmStateHash.Hex():
				// process the local state notify event
				fmt.Printf("Log Name: ConfirmState\n")

				event, err := lsmFilter.ParseConfirmState(vLog)
				if err != nil {
					log.Fatal(err)
				}

				//NoTokenOwner := common.HexToAddress(vLog.Topics[1].Hex())
				//Spender := common.HexToAddress(vLog.Topics[2].Hex())

				fmt.Printf("    unlock: %s\n", event.Unlock)
				fmt.Printf("    now Style: %s\n", event.Now)
				fmt.Printf("    result: %s\n", event.Result)
				fmt.Println("}")
				break
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
				break
			default:
				break
			}

			fmt.Println("***** finish an event *****")
		}
	}
}

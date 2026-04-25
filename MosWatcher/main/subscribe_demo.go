package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"os"
	"strings"
)

func scmain() {
	fmt.Println("***** connect the eth node *****")
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
	fmt.Println(Key, chainID, height)

	fmt.Println("***** test the contract *****")
	auth, err := bind.NewTransactorWithChainID(strings.NewReader(string(jsonBytes)), myPassword, chainID)
	if err != nil {
		log.Fatal(err)
	}
	auth.GasPrice = big.NewInt(1)
	auth.GasLimit = uint64(2100000)

	scAddress := common.HexToAddress(storageAddr)
	// scStorage, _ := NewStorage(scAddress, client)

	// build the abi
	//storageAbi, err := abi.JSON(strings.NewReader(StorageMetaData.ABI))
	//if err != nil {
	//	log.Fatal(err)
	//}

	// listen the event
	query := ethereum.FilterQuery{
		Addresses: []common.Address{scAddress},
	}
	logs := make(chan types.Log)

	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	// build filter
	storageFilter, err := NewStorageFilterer(scAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	// begin listening
	fmt.Println("LISTENING STORAGE EVENT:")
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println("vlog = ", vLog) // pointer to event log

			// decode the event
			event, err := storageFilter.ParseABIEvent(vLog)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(string(event.ReturnData)) // foo
			fmt.Println("issuccess = ", event.Success)
		}
	}
}

// UseIterator 实际上是遍历器, 只用于遍历已有的
func UseIterator(client *ethclient.Client, ctx context.Context) {
	// build Filter
	scAddress := common.HexToAddress(storageAddr)
	storageFilter, err := NewStorageFilterer(scAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	// get Iteration
	blockNum, _ := client.BlockNumber(ctx)
	endBlock := blockNum + 100
	filterOpt := bind.FilterOpts{blockNum, &endBlock, ctx}
	storageIterator, err := storageFilter.FilterABIEvent(&filterOpt)
	if err != nil {
		log.Fatal(err)
	}

	// get one event
	success := storageIterator.Next()
	if !success {
		log.Fatal(storageIterator.Error())
	}

	fmt.Println(string(storageIterator.Event.ReturnData))
	fmt.Println("is success = ", storageIterator.Event.Success)
}

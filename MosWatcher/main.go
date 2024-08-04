package main

import (
	"EventWatcher/contract"
	"EventWatcher/costlog"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"os"
	"strconv"
	"time"
)

var myPassword = "123456789"
var MosPrivateKey = "e9b1d63e8acd7fe676acb43afb390d4b0202dab61abec9cf2a561e4becb147de"
var MosUrl = "ws://172.18.166.229:8646"

// var MosUrl = "ws://127.0.0.1:8646"

var OnChain *log.Logger
var OffChain *log.Logger
var OnChainLogPath = "./On-Mos-log-4.txt"
var OffChainLogPath = "./Off-Mos-log-4.txt"
var TxLogPath = "./TxLogInfo.txt"

func logInit(onChainFile *os.File, offChainFile *os.File) {
	OnChain = log.New(os.Stdout, "OnChain: ", log.Ldate|log.Ltime|log.Lshortfile)
	OnChain.SetOutput(onChainFile)

	OffChain = log.New(os.Stdout, "OffChain: ", log.Ldate|log.Ltime|log.Lshortfile)
	OffChain.SetOutput(offChainFile)
}

func test() {
	start := time.Now()

	privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	if err != nil {
		log.Fatal(err)
	}

	data := []byte("hello")
	hash := crypto.Keccak256Hash(data)
	fmt.Println(hash.Hex())

	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	cost := time.Now().Sub(start)
	fmt.Println("cost = ", cost.Microseconds())
	fmt.Println(hexutil.Encode(signature))
	fmt.Println("cost = ", cost)
	fmt.Println(hexutil.Encode(signature))
}

func main() {
	fmt.Println("***** argument test *****")
	for idx, args := range os.Args {
		fmt.Println("参数"+strconv.Itoa(idx)+":", args)
	}

	fmt.Println("***** Init logs *****")
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

	fmt.Println("***** connect the eth node *****")
	client, err := ethclient.Dial(MosUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// fmt.Println(client, "successfully connected")

	// file := "../../data/keystore/UTC-USER2--2023-09-10T10-46-57.143272940Z--5e6aff8b8f168a52e2b8c2882442aa51482292f8"
	//file := "./data/UTC-USER2"
	//jsonBytes, err := os.ReadFile(file)

	// password, _ := os.ReadFile("../../data/password.txt")

	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	chainID, err := client.ChainID(ctx)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println("chainID = ", chainID)
	//Key, _ := keystore.DecryptKey(jsonBytes, string(myPassword))
	////Key, _ = keystore.DecryptKey(jsonBytes, myPassword)
	//height, _ := client.BlockNumber(ctx)
	//fmt.Println(Key, chainID, height)

	fmt.Println("***** test the contract *****")
	private_key, err := crypto.HexToECDSA(MosPrivateKey)
	auth, err := bind.NewKeyedTransactorWithChainID(private_key, chainID)
	if err != nil {
		log.Fatal(err)
	}
	auth.GasPrice = big.NewInt(1)
	auth.GasLimit = uint64(8000000)

	//lsmAddress := common.HexToAddress(contract.LocalStateManagerAddr)
	//crAddress := common.HexToAddress(contract.CrossRollupAddr)
	frAddress := common.HexToAddress(contract.MosFullRollup)
	// scStorage, _ := NewStorage(scAddress, client)

	// build the abi
	//lsmABI, err := abi.JSON(strings.NewReader(contract.LocalStateManagerMetaData.ABI))
	//if err != nil {
	//	log.Fatal(err)
	//}

	// listen the event --target
	query := ethereum.FilterQuery{
		//Addresses: []common.Address{lsmAddress},
		//Addresses: []common.Address{crAddress},
		Addresses: []common.Address{frAddress},
	}
	logs := make(chan types.Log)

	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	// 获取事件哈希值
	LocalStateNotifyHash := crypto.Keccak256Hash([]byte("LocalStateNotify(address,string,uint256)"))
	ErrorNotifyHash := crypto.Keccak256Hash([]byte("ErrorNotify(uint256,string)"))
	ConfirmStateHash := crypto.Keccak256Hash([]byte("ConfirmState(uint256,uint256,string)"))
	RollupBlockNotifyHash := crypto.Keccak256Hash([]byte("RollupBlockNotify(uint256,string)"))
	fmt.Println("LocalStateNotifyHash = ", LocalStateNotifyHash)
	fmt.Println("ErrorNotifyHash = ", ErrorNotifyHash)
	fmt.Println("ConfirmStateHash = ", ConfirmStateHash)

	// build filter  and iterator
	//lsmFilter, err := contract.NewLocalStateManagerFilterer(lsmAddress, client)
	//if err != nil {
	//	log.Fatal(err)
	//}

	//crFilter, err := contract.NewCrossRollupFilterer(crAddress, client)
	//if err != nil {
	//	log.Fatal(err)
	//}

	frFilter, err := contract.NewFullRollupFilterer(frAddress, client)
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

				rbNotify, err := frFilter.ParseRollupBlockNotify(vLog)
				if err != nil {
					log.Fatal(err)
				}

				//NoTokenOwner := common.HexToAddress(vLog.Topics[1].Hex())
				//Spender := common.HexToAddress(vLog.Topics[2].Hex())

				fmt.Printf("    BlockID: %s\n", rbNotify.BlockNumber)
				fmt.Printf("    Desc: %s\n", rbNotify.Desc)
				fmt.Println("}")

				cur := time.Now()
				costlog.OnChainLog.Println("New Block-"+rbNotify.BlockNumber.String(), cur.UnixNano() / 1e6)

				break
			case LocalStateNotifyHash.Hex():
				// process the local state notify event
				fmt.Printf("Log Name: LocalStateNotify\n")

				lsmNotify, err := frFilter.ParseLocalStateNotify(vLog)
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

			case ErrorNotifyHash.Hex():
				// process the error notify event
				fmt.Printf("Log Name: LocalStateNotify\n")
				// -- target
				//errNotify, err := lsmFilter.ParseErrorNotify(vLog)
				errNotify, err := frFilter.ParseErrorNotify(vLog)
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

//>>>>>>> 88ac45768bfab7c997467cb83bec894e6c387bd7

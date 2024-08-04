package challenger

import (
	"context"
	"fmt"
	"log"
	"os"
	"rollup-offchain/contract"
	"strconv"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func ChallengerListenerStart(client *ethclient.Client, auth *bind.TransactOpts, caller *bind.CallOpts, ctx context.Context) {
	fmt.Println("********** Argument Analysis **********")
	for idx, args := range os.Args {
		fmt.Println("参数"+strconv.Itoa(idx)+":", args)
	}

	fmt.Println("********** Init the Event Filter **********")
	cmAddress := common.HexToAddress(contract.BCPManagerAddr)

	// filter config
	query := ethereum.FilterQuery{
		Addresses: []common.Address{cmAddress},
	}
	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	// build filter  and iterator
	cmFilter, err := contract.NewChallengeManagerFilterer(cmAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Event Filter is Created! ")

	fmt.Println("********** Fetch the Event Hash **********")
	// 获取事件哈希值
	ChallengeStateNotifyHash := crypto.Keccak256Hash([]byte("ChallengeStateNotify(string,uint256,string)"))
	BlockInfoNotifyHash := crypto.Keccak256Hash([]byte("BlockInfoNotify(string,uint256,string)"))
	fmt.Println("ChallengeStateNotifyHash = ", ChallengeStateNotifyHash)
	fmt.Println("BlockInfoNotifyHash = ", BlockInfoNotifyHash)

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
			fmt.Printf("Log Detail: \n")

			switch vLog.Topics[0].Hex() {
			case BlockInfoNotifyHash.Hex():
				// process the local state notify event
				fmt.Printf("Log Name: RollupBlockNotifyHash\n")

				eventDetail, err := cmFilter.ParseBlockInfoNotify(vLog)
				if err != nil {
					log.Fatal(err)
				}

				//NoTokenOwner := common.HexToAddress(vLog.Topics[1].Hex())
				//Spender := common.HexToAddress(vLog.Topics[2].Hex())

				fmt.Printf("    Index: %s\n", eventDetail.Index)
				fmt.Printf("    Style: %s\n", eventDetail.Style)
				fmt.Printf("    Desc: %s\n", eventDetail.Desc)
				fmt.Println("")
				break
			case ChallengeStateNotifyHash.Hex():
				// process the local state notify event
				fmt.Printf("Log Name: ChallengeStateNotify\n")

				eventDetail, err := cmFilter.ParseChallengeStateNotify(vLog)
				if err != nil {
					log.Fatal(err)
				}

				// 打印部分事件信息
				fmt.Printf("    Index: %s\n", eventDetail.Index)
				fmt.Printf("    CurrentState: %s\n", eventDetail.CurrentState)
				fmt.Printf("    Desc: %s\n", eventDetail.Desc)
				fmt.Println("")

				// 调用挑战者的监听功能
				ChallengerRouting(eventDetail, client, auth, caller, ctx)

				break

			default:
				break
			}

			fmt.Println("***** finish an event *****")
		}
	}
}

func QuestionListenerStart(client *ethclient.Client, auth *bind.TransactOpts, caller *bind.CallOpts, ctx context.Context) {
	fmt.Println("********** Argument Analysis **********")
	for idx, args := range os.Args {
		fmt.Println("参数"+strconv.Itoa(idx)+":", args)
	}

	fmt.Println("********** Init the Event Filter **********")
	cmAddress := common.HexToAddress(contract.BCPManagerAddr)

	// filter config
	query := ethereum.FilterQuery{
		Addresses: []common.Address{cmAddress},
	}
	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	// build filter  and iterator
	cmFilter, err := contract.NewChallengeManagerFilterer(cmAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Event Filter is Created! ")

	fmt.Println("********** Fetch the Event Hash **********")
	// 获取事件哈希值
	ChallengeStateNotifyHash := crypto.Keccak256Hash([]byte("ChallengeStateNotify(string,uint256,string)"))
	BlockInfoNotifyHash := crypto.Keccak256Hash([]byte("BlockInfoNotify(string,uint256,string)"))
	fmt.Println("ChallengeStateNotifyHash = ", ChallengeStateNotifyHash)
	fmt.Println("BlockInfoNotifyHash = ", BlockInfoNotifyHash)

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
			fmt.Printf("Log Detail: \n")

			switch vLog.Topics[0].Hex() {
			case BlockInfoNotifyHash.Hex():
				// process the local state notify event
				fmt.Printf("Log Name: RollupBlockNotifyHash\n")

				eventDetail, err := cmFilter.ParseBlockInfoNotify(vLog)
				if err != nil {
					log.Fatal(err)
				}

				//NoTokenOwner := common.HexToAddress(vLog.Topics[1].Hex())
				//Spender := common.HexToAddress(vLog.Topics[2].Hex())

				fmt.Printf("    Index: %s\n", eventDetail.Index)
				fmt.Printf("    Style: %s\n", eventDetail.Style)
				fmt.Printf("    Desc: %s\n", eventDetail.Desc)
				fmt.Println("")
				break
			case ChallengeStateNotifyHash.Hex():
				// process the local state notify event
				fmt.Printf("Log Name: ChallengeStateNotify\n")

				eventDetail, err := cmFilter.ParseChallengeStateNotify(vLog)
				if err != nil {
					log.Fatal(err)
				}

				// 打印部分事件信息
				fmt.Printf("    Index: %s\n", eventDetail.Index)
				fmt.Printf("    CurrentState: %s\n", eventDetail.CurrentState)
				fmt.Printf("    Desc: %s\n", eventDetail.Desc)
				fmt.Println("")

				// 调用挑战者的监听功能
				QuestionerRouting(eventDetail, client, auth, caller, ctx)

				break

			default:
				break
			}

			fmt.Println("***** finish an event *****")
		}
	}
}

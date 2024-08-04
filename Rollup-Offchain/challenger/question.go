package challenger

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log"
	"math/big"
	"rollup-offchain/contract"
	"rollup-offchain/costlog"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

// QuestionerRouting 挑战者的监听功能实现
func QuestionerRouting(eventDetail *contract.ChallengeManagerChallengeStateNotify, client *ethclient.Client, auth *bind.TransactOpts, caller *bind.CallOpts, ctx context.Context) {
	id := isInMyQuestionPool(eventDetail.Index)

	if id < 0 && eventDetail.CurrentState.Int64() != 3 {
		fmt.Println("It is not nedd to process")
		return
	}

	// 挑战者的监听功能
	if eventDetail.CurrentState.Int64() == 3 { // 如果挑战成立就创建一个记录
		err := QuestionConfirm(client, caller, eventDetail.Index)
		if err != nil {
			log.Fatal(err)
		}
	} else if eventDetail.CurrentState.Int64() == 2 { // 如果挑战被回复就进行下一级挑战
		err := QuestionArgu(client, auth, caller, ctx, eventDetail.Index)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err := QuestionUpdate(client, caller, eventDetail.Index)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// MyQuestionPool 管理本人发起质疑的挑战
var MyQuestionPool []ChallengeRecord

func isInMyQuestionPool(index string) int {
	for i, c := range MyQuestionPool {
		if c.Index == index {
			return i
		}
	}

	return -1
}

func GetQuestionFromMyQuestionPool(index string) (int, ChallengeRecord, error) {
	for i, c := range MyQuestionPool {
		if c.Index == index {
			return i, c, nil
		}
	}

	return 0, ChallengeRecord{}, errors.New("element no exist")
}

func GetQuestionFromChain(client *ethclient.Client, caller *bind.CallOpts, index string) (*ChallengeRecord, error) {
	fmt.Println("# GetQuestionFromChain Begin")
	record, err := BCPInstance.ChallengePool(caller, index)
	if err != nil {
		return nil, err
	}

	if record.Questioner == MyAddr {
		fmt.Println("Question Get Success")
	} else {
		fmt.Println("No this Question")
		return nil, errors.New("GetQuestionFromChain: not my Question")
	}

	var cr ChallengeRecord
	cr = record

	fmt.Println("# GetQuestionFromChain End")
	return &cr, nil
}

// QuestionUpdate
// QuestionUpdate: update the challenge record
func QuestionUpdate(client *ethclient.Client, caller *bind.CallOpts, index string) error {
	fmt.Println("# Question Update Begin")
	record, err := BCPInstance.ChallengePool(caller, index)
	if err != nil {
		return err
	}

	fmt.Println("Get a Question!")

	if record.Questioner == MyAddr {
		// search local record
		id := isInMyQuestionPool(index)

		// update local record
		if id < 0 {
			fmt.Println("Create New Question! Success!")
			MyQuestionPool = append(MyQuestionPool, record)
		} else {
			fmt.Println("Update Local Question! Success!")
			MyQuestionPool[id] = record
		}
	}

	fmt.Println("# Question Update End")
	return nil
}

// QuestionCreate to start an interactive question process
// QuestionCreate: question the target challenge by index
func QuestionCreate(client *ethclient.Client, auth *bind.TransactOpts, ctx context.Context, index string) error {

	fmt.Println("********** QuestionCreate Begin")
	start := time.Now()

	// 调用质疑接口
	tx, err := BCPInstance.ChallengeQuestion(auth, index, false)
	if err != nil {
		return err
	}
	txStr, _ := tx.MarshalJSON()
	fmt.Println("ChallengeCreate Tx = ", string(txStr))

	// 计算最终时间开销并计算
	cost := time.Now().Sub(start)
	costlog.OffChainLog.Println("QuestionCreate", cost)
	// 添加交易记录
	costlog.OnChainLog.Println("QuestionCreate", tx.Hash().Hex())
	fmt.Println("********** QuestionCreate End")
	return nil
}

// QuestionConfirm state = 3
func QuestionConfirm(client *ethclient.Client, caller *bind.CallOpts, index string) error {
	fmt.Println("********** Question Confirm Begin")
	record, err := BCPInstance.ChallengePool(caller, index)
	if err != nil {
		return err
	}

	fmt.Println("Get a Question!")

	if record.Questioner == MyAddr {
		fmt.Println("Question Create Success")
		MyQuestionPool = append(MyQuestionPool, record)
	}

	fmt.Println("********** Question Confirm End")
	return nil
}

// QuestionArgu state = 2
// QuestionArgu: verify the current block and then reply a new question online
func QuestionArgu(client *ethclient.Client, auth *bind.TransactOpts, caller *bind.CallOpts, ctx context.Context, index string) error {
	fmt.Println("********** Question Argu Begin")
	start := time.Now()

	// 获取最新状态
	id := isInMyQuestionPool(index)
	if id < 0 {
		return errors.New("no this challenge")
	}

	err := QuestionUpdate(client, caller, index)
	if err != nil {
		return err
	}

	// 获取正确的中间区块数据
	middle := big.NewInt((MyQuestionPool[id].Begin.Int64() + MyQuestionPool[id].End.Int64()) / 2)
	block, err := client.BlockByNumber(ctx, middle)
	if err != nil {
		return err
	}

	// 比较区块数据的正确性
	ack := bytes.Equal(MyQuestionPool[id].MiddleBlock.BlockData, block.Header().Hash().Bytes())

	// 发起挑战响应
	tx, err := BCPInstance.ChallengeQuestion(auth, index, ack)
	if err != nil {
		return err
	}
	txStr, _ := tx.MarshalJSON()
	fmt.Println("QuestionArgu Tx = ", string(txStr))

	// 计算最终时间开销并计算
	cost := time.Now().Sub(start)
	costlog.OffChainLog.Println("QuestionArgu", cost)
	// 添加交易记录
	costlog.OnChainLog.Println("QuestionArgu", tx.Hash().Hex())
	fmt.Println("********** Question Argu End")
	return nil
}

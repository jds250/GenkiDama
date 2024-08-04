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
	"rollup-offchain/zk"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// ChallengerRouting 挑战者的监听功能实现
func ChallengerRouting(eventDetail *contract.ChallengeManagerChallengeStateNotify, client *ethclient.Client, auth *bind.TransactOpts, caller *bind.CallOpts, ctx context.Context) {
	id := isInMyChallengePool(eventDetail.Index)

	if id < 0 && eventDetail.CurrentState.Int64() != 1 {
		fmt.Println("It is not nedd to process")
		return
	}

	// 挑战者的监听功能
	if eventDetail.CurrentState.Int64() == 1 {
		// 如果创建成功那就检查该状态是否合法
		err := ChallengeUpdate(client, caller, eventDetail.Index)
		if err != nil {
			log.Fatal(err)
		}
	} else if eventDetail.CurrentState.Int64() == 3 {
		// 如果处于等待回复的状态，对质疑进行回复
		err := ChallengeResponse(client, auth, caller, ctx, eventDetail.Index)
		if err != nil {
			log.Fatal(err)
		}
	} else if eventDetail.CurrentState.Int64() == 4 {
		// 如果处于最终证明阶段，直接对其进行处理
		err := ChallengeFinalProof(client, auth, caller, ctx, eventDetail.Index)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err := ChallengeUpdate(client, caller, eventDetail.Index)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// MyChallengePool 管理本人发起的挑战
var MyChallengePool []ChallengeRecord

func isInMyChallengePool(index string) int {
	for i, c := range MyChallengePool {
		if c.Index == index {
			return i
		}
	}

	return -1
}

func GetRecordFromMyChallengePool(index string) (int, ChallengeRecord, error) {
	for i, c := range MyChallengePool {
		if c.Index == index {
			return i, c, nil
		}
	}

	return 0, ChallengeRecord{}, errors.New("element no exist")
}

func GetRecordFromChain(client *ethclient.Client, caller *bind.CallOpts, index string) (*ChallengeRecord, error) {
	fmt.Println("# GetRecordFromChain Begin")
	record, err := BCPInstance.ChallengePool(caller, index)
	if err != nil {
		return nil, err
	}

	if record.Challenger == MyAddr {
		fmt.Println("Challenge Get Success")
	} else {
		fmt.Println("No this Challenge")
		return nil, errors.New("GetRecordFromChain: not my challenge")
	}

	var cr ChallengeRecord
	cr = record

	fmt.Println("# GetRecordFromChain End")
	return &cr, nil
}

// ChallengeUpdate
// ChallengeUpdate: update the challenge record
func ChallengeUpdate(client *ethclient.Client, caller *bind.CallOpts, index string) error {
	fmt.Println("# Challenge Update Begin")
	record, err := BCPInstance.ChallengePool(caller, index)
	if err != nil {
		return err
	}

	fmt.Println("Get a Challenge!")

	if record.Challenger == MyAddr {
		// search local record
		id := isInMyChallengePool(index)

		// update local record
		if id < 0 {
			fmt.Println("Create New Challenge! Success!")
			MyChallengePool = append(MyChallengePool, record)
		} else {
			fmt.Println("Update Local Challenge! Success!")
			MyChallengePool[id] = record
		}
	}

	fmt.Println("# Challenge Update End")
	return nil
}

// ChallengeCreate state=0
// ChallengeCreate: create a new challenge and wait for confirm
// only when the challenge is confirmed, the record will create
func ChallengeCreate(client *ethclient.Client, auth *bind.TransactOpts, index string, headerID *big.Int) error {

	fmt.Println("********** Challenge Create Begin **********")
	start := time.Now()

	ctx := context.Background()
	// 获取当前链信息：区块高度，区块数据，状态根（暂无）
	//block, proof, err := zk.GetEthBlockAndProof(client, ctx, nil)
	//if err != nil {
	//	return err
	//}

	// 获取区块header信息
	block, err := client.BlockByNumber(ctx, headerID)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	err = block.Header().EncodeRLP(buf)
	if err != nil {
		return err
	}
	headerStr := buf.Bytes()

	// 执行脸上验证
	//isLegal, err := zk.VerifyProofByContract(client, ctx, proof)
	//if err != nil {
	//	return err
	//}
	//if isLegal == false {
	//	return errors.New("illegal challenge proof")
	//}

	// 创建挑战信息
	// TODO: build a real cross-chain inconsistency proof
	// TODO: - the state proof of commitment /
	// TODO: the transaction proof of commitment transaction(or execution challenge transaction)
	tmp1 := []byte("123")

	cState := contract.DataTypesChallengeState{
		ChainID:     MyChainID,
		L2BlockID:   big.NewInt(1),
		L1BlockID:   block.Number(),
		L1StateRoot: []byte("test"),
		L1BlockData: block.Header().Hash().Bytes(),
		StateProof:  [][]byte{tmp1},
		Account:     common.HexToAddress(contract.USER1_Addr),
	}

	// 调用挑战接口
	tx, err := BCPInstance.ChallengeCreate(auth, cState, index, headerStr)
	if err != nil {
		return err
	}
	txStr, _ := tx.MarshalJSON()
	fmt.Println("ChallengeCreate Tx = ", string(txStr))

	// 计算最终时间开销并计算
	cost := time.Now().Sub(start)
	costlog.OffChainLog.Println("ChallengeCreate", cost)
	// 添加交易记录
	costlog.OnChainLog.Println("ChallengeCreate", tx.Hash().Hex())
	fmt.Println("********** Challenge Create End **********")
	return nil
}

// ChallengeResponse state = 3
// ChallengeResponse: give the blockData of the new middle block
func ChallengeResponse(client *ethclient.Client, auth *bind.TransactOpts, caller *bind.CallOpts, ctx context.Context, index string) error {
	fmt.Println("********** Challenge Response Begin **********")

	start := time.Now()

	// 获取挑战记录
	id := isInMyChallengePool(index)
	if id < 0 {
		return errors.New("no this challenge")
	}
	err := ChallengeUpdate(client, caller, index)
	if err != nil {
		return err
	}
	targetChallenge := MyChallengePool[id]

	// 获取需要的区块信息
	middle := big.NewInt((targetChallenge.Begin.Int64() + targetChallenge.End.Int64()) / 2)
	block, err := client.BlockByNumber(ctx, middle)
	if err != nil {
		return err
	}

	// 构造返回值信息
	stateRoot := []byte("test")
	middleBlock := contract.DataTypesL1BlockInfo{
		BlockID:   middle,
		ChainID:   MyChainID,
		BlockData: block.Header().Hash().Bytes(),
		StateRoot: stateRoot,
	}

	// 发起回复
	tx, err := BCPInstance.ChallengeResponse(auth, index, middleBlock)
	if err != nil {
		return err
	}
	txStr, _ := tx.MarshalJSON()
	fmt.Println("ChallengeResponse Tx = ", string(txStr))

	// 更新本地记录
	MyChallengePool[id].State = big.NewInt(3)

	// 计算最终时间开销并计算
	cost := time.Now().Sub(start)
	costlog.OffChainLog.Println("ChallengeResponse", cost)
	// 添加交易记录
	costlog.OnChainLog.Println("ChallengeResponse", tx.Hash().Hex())

	fmt.Println("********** Challenge Response End **********")
	return nil
}

// ChallengeFinalProof state = 4
// TODO: generate the final proof and upload
func ChallengeFinalProof(client *ethclient.Client, auth *bind.TransactOpts, caller *bind.CallOpts, ctx context.Context, index string) error {
	fmt.Println("********** Challenge FinalProof Begin **********")
	start := time.Now()

	// 获取挑战记录
	id := isInMyChallengePool(index)
	if id < 0 {
		return errors.New("no this challenge")
	}
	err := ChallengeUpdate(client, caller, index)
	if err != nil {
		return err
	}
	targetChallenge := MyChallengePool[id]

	// 获取目标区块 以及 证明
	_, proof, err := zk.GetEthBlockAndProof(client, ctx, targetChallenge.End)
	if err != nil {
		return err
	}

	// 获取header信息
	blk1, err := client.BlockByNumber(ctx, targetChallenge.Begin)
	blk2, err := client.BlockByNumber(ctx, targetChallenge.End)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	err = blk1.Header().EncodeRLP(buf)
	h1Str := buf.Bytes()
	err = blk2.Header().EncodeRLP(buf)
	h2Str := buf.Bytes()
	if err != nil {
		return err
	}

	// 调用合约发起证明
	tx, err := BCPInstance.FinalChallenge(auth, index, *proof, h1Str, h2Str)
	if err != nil {
		return err
	}
	txStr, _ := tx.MarshalJSON()
	fmt.Println("ChallengeResponse Tx = ", string(txStr))

	// 计算最终时间开销并计算
	cost := time.Now().Sub(start)
	costlog.OffChainLog.Println("ChallengeFinalProof", cost)

	// 添加交易记录
	costlog.OnChainLog.Println("ChallengeFinalProof", tx.Hash().Hex())

	fmt.Println("********** Challenge FinalProof End **********")
	return nil
}

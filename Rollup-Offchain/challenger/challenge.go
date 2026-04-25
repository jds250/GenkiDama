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
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// ChallengerRouting 挑战者的监听功能实现
func ChallengerRouting(eventDetail *contract.BCPManagerChallengeStateNotify, client *ethclient.Client, auth *bind.TransactOpts, caller *bind.CallOpts, ctx context.Context) {
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
	record, err := fetchChallengeRecord(client, caller, index)
	if err != nil {
		return nil, err
	}

	if record.Challenger == MyAddr {
		fmt.Println("Challenge Get Success")
	} else {
		fmt.Println("No this Challenge")
		return nil, errors.New("GetRecordFromChain: not my challenge")
	}

	fmt.Println("# GetRecordFromChain End")
	return record, nil
}

// ChallengeUpdate
// ChallengeUpdate: update the challenge record
func ChallengeUpdate(client *ethclient.Client, caller *bind.CallOpts, index string) error {
	fmt.Println("# Challenge Update Begin")
	record, err := fetchChallengeRecord(client, caller, index)
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
			MyChallengePool = append(MyChallengePool, *record)
		} else {
			fmt.Println("Update Local Challenge! Success!")
			MyChallengePool[id] = *record
		}
	}

	fmt.Println("# Challenge Update End")
	return nil
}

// ChallengeCreate state=0
// ChallengeCreate: create a new challenge and wait for confirm
// only when the challenge is confirmed, the record will create
func ChallengeCreate(client *ethclient.Client, auth *bind.TransactOpts, index string, headerID *big.Int) error {
	return errors.New("use ChallengeCreateFromRemoteCommitment with a verified remote commitment hash")
}

func ChallengeCreateFromRemoteCommitment(
	client *ethclient.Client,
	auth *bind.TransactOpts,
	ctx context.Context,
	index string,
	localRollupBlockID *big.Int,
	remoteHeaderID *big.Int,
	remoteCommitmentHash [32]byte,
) error {
	fmt.Println("********** Challenge Create Begin **********")
	start := time.Now()

	remoteBlock, err := client.BlockByNumber(ctx, remoteHeaderID)
	if err != nil {
		return err
	}

	cState := BuildConsistencyChallengeState(
		common.HexToAddress(contract.CrossRollupAddr),
		localRollupBlockID,
		remoteBlock,
		remoteCommitmentHash,
	)

	if err := SubmitConsistencyChallenge(client, auth, index, cState, remoteBlock); err != nil {
		return err
	}

	cost := time.Now().Sub(start)
	costlog.OffChainLog.Println("ChallengeCreate", cost)
	fmt.Println("********** Challenge Create End **********")
	return nil
}

func ChallengeCreateFromRemoteProof(
	client *ethclient.Client,
	auth *bind.TransactOpts,
	index string,
	localRollupBlockID *big.Int,
	remoteBlock *types.Block,
	remoteProof *ConsistencyMerkleProof,
) error {
	fmt.Println("********** Challenge Create With Proof Begin **********")
	start := time.Now()

	cState := BuildConsistencyChallengeStateWithProof(
		common.HexToAddress(contract.CrossRollupAddr),
		localRollupBlockID,
		remoteBlock,
		remoteProof,
	)

	if err := SubmitConsistencyChallenge(client, auth, index, cState, remoteBlock); err != nil {
		return err
	}

	cost := time.Now().Sub(start)
	costlog.OffChainLog.Println("ChallengeCreateWithProof", cost)
	fmt.Println("********** Challenge Create With Proof End **********")
	return nil
}

func ChallengeCreateFromRemoteStateProof(
	client *ethclient.Client,
	auth *bind.TransactOpts,
	index string,
	localRollupBlockID *big.Int,
	remoteBlock *types.Block,
	remoteProof *ConsistencyMerkleProof,
	accountProof *AccountProofRPCResult,
	slot *big.Int,
) error {
	fmt.Println("********** Challenge Create With State Proof Begin **********")
	start := time.Now()

	cState, err := BuildConsistencyChallengeStateWithMPTProof(
		common.HexToAddress(contract.CrossRollupAddr),
		localRollupBlockID,
		remoteBlock,
		remoteProof,
		accountProof,
		slot,
	)
	if err != nil {
		return err
	}

	if err := SubmitConsistencyChallenge(client, auth, index, cState, remoteBlock); err != nil {
		return err
	}

	cost := time.Now().Sub(start)
	costlog.OffChainLog.Println("ChallengeCreateWithStateProof", cost)
	fmt.Println("********** Challenge Create With State Proof End **********")
	return nil
}

func ChallengeCreateFromRemoteRollupProof(
	client *ethclient.Client,
	auth *bind.TransactOpts,
	index string,
	localRollupBlockID *big.Int,
	remoteHeaderBlock *types.Block,
	remoteCommitmentHash [32]byte,
	remoteRollupBlockID *big.Int,
	accountProof *AccountProofRPCResult,
	slots []*big.Int,
) error {
	fmt.Println("********** Challenge Create With Rollup Proof Begin **********")
	start := time.Now()

	cState, err := BuildConsistencyChallengeStateWithRollupProof(
		common.HexToAddress(contract.CrossRollupAddr),
		localRollupBlockID,
		remoteHeaderBlock,
		remoteCommitmentHash,
		remoteRollupBlockID,
		accountProof,
		slots,
	)
	if err != nil {
		return err
	}

	if err := SubmitConsistencyChallenge(client, auth, index, cState, remoteHeaderBlock); err != nil {
		return err
	}

	cost := time.Now().Sub(start)
	costlog.OffChainLog.Println("ChallengeCreateWithRollupProof", cost)
	fmt.Println("********** Challenge Create With Rollup Proof End **********")
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
	middleBlock := contract.DataTypesL1BlockInfo{
		BlockID:   middle,
		ChainID:   MyChainID,
		BlockData: block.Header().Hash().Bytes(),
		StateRoot: block.Root().Bytes(),
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

	buf1 := new(bytes.Buffer)
	err = blk1.Header().EncodeRLP(buf1)
	if err != nil {
		return err
	}
	h1Str := buf1.Bytes()

	buf2 := new(bytes.Buffer)
	err = blk2.Header().EncodeRLP(buf2)
	h2Str := buf2.Bytes()
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

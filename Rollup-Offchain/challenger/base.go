package challenger

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"rollup-offchain/contract"
	"rollup-offchain/costlog"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// 局部变量

var IsChallenger bool
var CMInstance *contract.ChallengeManager
var BCPInstance *contract.BCPManager
var MyAddr common.Address
var MyChainID *big.Int

// ChallengeRecord 挑战记录的数据结构（与链上一致）
type ChallengeRecord struct {
	Index       string
	Detail      contract.DataTypesChallengeState
	ConfirmTime *big.Int
	Challenger  common.Address
	Questioner  common.Address
	State       *big.Int
	Begin       *big.Int
	End         *big.Int
	BeginBlock  contract.DataTypesL1BlockInfo
	EndBlock    contract.DataTypesL1BlockInfo
	MiddleBlock contract.DataTypesL1BlockInfo
}

type ChallengeCheckPoint struct {
	Index        *big.Int
	ChainID      *big.Int
	ChainType    *big.Int
	BlockID      *big.Int
	StateRoot    []byte
	BlockData    []byte
	ContractAddr common.Address
}

func ChallangeInit(client *ethclient.Client, myAddr common.Address, chainID *big.Int) {
	// 初始化挑战管理合约
	MyChainID = chainID
	MyAddr = myAddr
	BCPAddr := common.HexToAddress(contract.BCPManagerAddr)
	BCP, err := contract.NewBCPManager(BCPAddr, client)
	if err != nil {
		log.Fatal(err)
	}

	BCPInstance = BCP
}

func AddCheckPoint(client *ethclient.Client, auth *bind.TransactOpts, caller *bind.CallOpts, ctx context.Context) error {
	start := time.Now()

	// get the latest block
	fmt.Println("get the latest block")
	current, err := client.BlockNumber(ctx)
	if err != nil {
		return err
	}
	targetNum := big.NewInt(int64(current))

	// 获取最新区块数据
	fmt.Println("获取最新区块数据")
	block, err := client.BlockByNumber(ctx, targetNum)
	if err != nil {
		return err
	}

	// 构造合约输入
	fmt.Println("构造合约输入")
	checkPoint := contract.DataTypesChainCheckPoint{
		Index:        big.NewInt(1),
		ChainID:      MyChainID,
		ChainType:    big.NewInt(1),
		BlockID:      targetNum,
		StateRoot:    []byte("Test"),
		BlockData:    block.Hash().Bytes(),
		ContractAddr: common.HexToAddress(contract.BCPManagerAddr),
	}

	// 调用合约
	fmt.Println("调用合约")
	tx, err := BCPInstance.CheckPoint(auth, checkPoint)
	if err != nil {
		return err
	}
	txStr, _ := tx.MarshalJSON()

	// 计算最终时间开销并计算
	cost := time.Now().Sub(start)
	costlog.OffChainLog.Println("AddCheckPoint", cost)
	// 添加交易记录
	costlog.OnChainLog.Println("AddCheckPoint", tx.Hash().Hex())
	fmt.Println("AddCheckPoint Tx = ", string(txStr))

	return nil
}

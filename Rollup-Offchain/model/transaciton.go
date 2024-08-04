package model

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type StateUnit struct {
	Addr    [32]byte `json:"addr"`
	Balance int64    `json:"balance"`
}

type ChainStateUnit struct {
	ChainID int64    `json:"chain_id"`
	Addr    [32]byte `json:"addr"`
	Balance int64    `json:"balance"`
}

type CoinUnit struct {
	Addr    [32]byte `json:"addr"`
	Balance int64    `json:"balance"`
}

// Transaction 0-普通offchain交易 1-链下到链上交易 2-链上到链下交易 3-测试交易
// ChainID 0-L2,1-chain1, ...
// 规定12两类交易只能是单输入单输出
type Transaction struct {
	FromCoins []CoinUnit `json:"from_coins"`
	ToCoins   []CoinUnit `json:"to_coins"`
	ChainID   int64      `json:"chain_id"`
	Type      int64      `json:"type"`
}

// L2Transaction 1-L1 to L2; 2-L2 to L1
// ChainID 0-L2,1-chain1, ...
// 规定所有交易只能是单输入单输出
type L2Transaction struct {
	FromAddr common.Address `json:"from_addr"`
	ToAddr   common.Address `json:"to_addr"`
	Value    int64          `json:"value"`
	ChainID  int64          `json:"chain_id"`
	Type     int64          `json:"type"`
}

// L2Transition the transition to calculate proof
type L2Transition struct {
	Addr    common.Address `json:"addr"`
	Value   int64          `json:"value"`
	Style   int64          `json:"style"`   // 1-ACCOUNT-ADD; 2-ACCOUNT-SUB
	ChainID int64          `json:"chainID"` // chain ID
}

// LocalTransition 本地
type LocalTransition struct {
	ChainID    *big.Int
	BlockNum   *big.Int
	Account    common.Address
	Value      *big.Int
	Style      *big.Int // 1-ACCOUNT-ADD; 2-ACCOUNT-SUB
	UnlockTime *big.Int
}
type TxsResult struct {
	InCoins  []CoinUnit `json:"in_coins"`
	OutCoins []CoinUnit `json:"out_coins"`
}

type CommitBlock struct {
	BlockNum       int64             `json:"block_num"`
	StateRoot      [32]byte          `json:"state_root"` // the L2 state root
	TxRoot         [32]byte          `json:"tx_root"`
	TransitionRoot [32]byte          `json:"transition_root"`
	TxBatch        []Transaction     `json:"tx_batch"`
	TransitionList []LocalTransition `json:"transition_list"`
}

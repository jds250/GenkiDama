package costlog

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"time"
)

/*************************
 *     On Chain Cost     *
 *************************/

type TxRecord struct {
	Name   string
	TxHash common.Hash
}

var TxRecordSet []TxRecord

func Tx2Record(name string, tx *types.Transaction) TxRecord {
	return TxRecord{
		Name:   name,
		TxHash: tx.Hash(),
	}
}

func AddTxRecord(name string, tx *types.Transaction) {
	TxRecordSet = append(TxRecordSet, Tx2Record(name, tx))
}

func CommitAllRecords(client *ethclient.Client, ctx context.Context) error {
	for _, r := range TxRecordSet {
		cost, err := GetOnChainCostByTxRecord(client, ctx, r)
		if err != nil {
			return err
		}

		// Add Log
		OnChainLog.Println(r.Name, cost)
	}

	TxRecordSet = []TxRecord{}

	return nil
}

func GetOnChainCostByTxHash(client *ethclient.Client, ctx context.Context, hash common.Hash) (uint64, error) {
	// 获取链上交易
	receipt, err := client.TransactionReceipt(ctx, hash)
	if err != nil {
		return 0, err
	}

	// 获取交易执行开销
	cost := receipt.GasUsed

	return cost, nil
}

func GetOnChainCostOfTx(client *ethclient.Client, ctx context.Context, tx *types.Transaction) (uint64, error) {
	// 获取链上交易
	receipt, err := client.TransactionReceipt(ctx, tx.Hash())
	if err != nil {
		return 0, err
	}

	// 获取交易执行开销
	cost := receipt.GasUsed

	return cost, nil
}

func GetCostOfTxSet(client *ethclient.Client, ctx context.Context, txs []*types.Transaction) (uint64, error) {
	var a uint64 = 0
	for _, tx := range txs {
		c, err := GetOnChainCostOfTx(client, ctx, tx)
		if err != nil {
			return 0, err
		}
		a += c
	}
	return a, nil
}

func GetOnChainCostByTxRecord(client *ethclient.Client, ctx context.Context, record TxRecord) (uint64, error) {
	// 获取链上交易
	receipt, err := client.TransactionReceipt(ctx, record.TxHash)
	if err != nil {
		return 0, err
	}

	// 获取交易执行开销
	cost := receipt.GasUsed

	return cost, nil
}

func GetCostOfTxRecordSet(client *ethclient.Client, ctx context.Context, records []TxRecord) (uint64, error) {
	var a uint64 = 0
	for _, r := range records {
		c, err := GetOnChainCostByTxRecord(client, ctx, r)
		if err != nil {
			return 0, err
		}
		a += c
	}
	return a, nil
}

func LogTxCostLog(client *ethclient.Client, ctx context.Context, tx *types.Transaction, name string) error {
	cost, err := GetOnChainCostOfTx(client, ctx, tx)
	if err != nil {
		return err
	}

	// Add Log
	OnChainLog.Println(name, cost)

	return nil
}

func LogTxRecordsCostLog(client *ethclient.Client, ctx context.Context, tx *types.Transaction) error {
	for _, r := range TxRecordSet {
		cost, err := GetOnChainCostByTxRecord(client, ctx, r)
		if err != nil {
			return err
		}

		// Add Log
		OnChainLog.Println(r.Name, cost)
	}

	return nil
}

/*************************
 *     Off Chain Cost    *
 *************************/

func LogOffChainCost(name string, duration time.Duration) {
	OffChainLog.Println(name, duration.String())
}

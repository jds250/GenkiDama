package cosmos_executor

import (
	"crypto/rand"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"rollup-offchain/model"
)

// ExeCrossTxSet 执行跨链交易
func ExeCrossTxSet(txs []model.L2Transaction) []model.L2Transition {
	var resTransition []model.L2Transition

	for i := 0; i < len(txs); i++ {
		tmpTransit := ExeOneCrossTx(txs[i])
		resTransition = append(resTransition, tmpTransit[0])
		resTransition = append(resTransition, tmpTransit[1])
	}

	return resTransition
}

// ExeOneCrossTx 执行一个跨链交易，生成transition
func ExeOneCrossTx(tx model.L2Transaction) []model.L2Transition {
	var resIn model.L2Transition
	var resOut model.L2Transition

	if tx.Type == 1 {
		resIn.ChainID = 1
		resOut.ChainID = 2
	} else {
		resIn.ChainID = 2
		resOut.ChainID = 1
	}

	resIn.Addr = tx.FromAddr
	resIn.Value = tx.Value
	resIn.Style = 1 // 1-sub, 2-add

	resOut.Addr = tx.ToAddr
	resOut.Value = tx.Value
	resOut.Style = 2 // 1-sub, 2-add

	resList := []model.L2Transition{resIn, resOut}

	return resList
}

// CrossTxSetGenerate 生成跨链交易
func CrossTxSetGenerate(num int) ([]model.L2Transaction, error) {
	var resTxSet []model.L2Transaction

	for i := 0; i < num; i++ {
		tmpTx, err := RandomCrossTx()
		if err != nil {
			return nil, err
		}
		resTxSet = append(resTxSet, *tmpTx)
	}

	return resTxSet, nil
}

// RandomCrossTx 生成随机跨链交易
func RandomCrossTx() (*model.L2Transaction, error) {
	randAddrInt1, err := rand.Int(rand.Reader, big.NewInt(int64(UserRange)))
	if err != nil {
		fmt.Println("rand fail!", err.Error())
		return nil, err
	}
	addrHash1 := common.BigToAddress(big.NewInt(1000 + randAddrInt1.Int64()))

	randAddrInt2, err := rand.Int(rand.Reader, big.NewInt(int64(UserRange)))
	if err != nil {
		fmt.Println("rand fail!", err.Error())
		return nil, err
	}
	addrHash2 := common.BigToAddress(big.NewInt(1000 + randAddrInt2.Int64()))

	randType, err := rand.Int(rand.Reader, big.NewInt(10))
	if err != nil {
		fmt.Println("rand fail!", err.Error())
		return nil, err
	}
	tmpType := randType.Int64() % 2

	// 设置为测试交易
	tx := model.L2Transaction{FromAddr: addrHash1, ToAddr: addrHash2, ChainID: 0, Type: tmpType, Value: 1}

	return &tx, nil
}

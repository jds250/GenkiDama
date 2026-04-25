package model

import (
	"math/big"
	"rollup-offchain/contract"
)

func Local2DTTransaction(tx L2Transaction) contract.DataTypesL2Transaction {
	var dt_tx contract.DataTypesL2Transaction
	dt_tx.Style = big.NewInt(tx.Type)
	dt_tx.FromAddr = tx.FromAddr
	dt_tx.ToAddr = tx.ToAddr
	dt_tx.Value = big.NewInt(tx.Value)
	dt_tx.ChainID = big.NewInt(tx.ChainID)

	return dt_tx
}

func Local2DtTxList(txs []L2Transaction) []contract.DataTypesL2Transaction {
	var resList []contract.DataTypesL2Transaction
	for i := 0; i < len(txs); i++ {
		tmp := Local2DTTransaction(txs[i])
		resList = append(resList, tmp)
	}
	return resList
}

func Local2DTTransition(ts L2Transition) contract.DataTypesL2Transition {
	var dt_ts contract.DataTypesL2Transition
	dt_ts.Style = big.NewInt(ts.Style)
	dt_ts.Account = ts.Addr
	dt_ts.Value = big.NewInt(ts.Value)
	dt_ts.ChainID = big.NewInt(ts.ChainID)

	return dt_ts
}

func Local2DtTsList(tss []L2Transition) []contract.DataTypesL2Transition {
	var resList []contract.DataTypesL2Transition
	for i := 0; i < len(tss); i++ {
		tmp := Local2DTTransition(tss[i])
		resList = append(resList, tmp)
	}
	return resList
}

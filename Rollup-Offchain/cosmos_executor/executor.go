package cosmos_executor

import (
	"crypto/rand"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"rollup-offchain/model"
	"rollup-offchain/trie"
)

// ExecuteL2TxSetAndProof 假定所有交易相关账户已创建
func ExecuteL2TxSetAndProof(currentTrie *trie.Trie, txs []model.L2Transaction) ([]model.L2Transition, []trie.Bag, error) {
	var resTransition []model.L2Transition
	var resBag []trie.Bag

	for i := 0; i < len(txs); i++ {
		transits, bags, err := ExecuteOneTxAndProof(currentTrie, txs[i])
		if err != nil {
			return nil, nil, err
		}
		resTransition = append(resTransition, transits...)
		resBag = append(resBag, bags...)
	}

	return resTransition, resBag, nil
}

func ExecuteOneTxAndProof(currentTrie *trie.Trie, tx model.L2Transaction) ([]model.L2Transition, []trie.Bag, error) {
	var resTransition []model.L2Transition
	var resBag []trie.Bag

	// 检查资金是否充足
	var getValue big.Int
	if tx.Type != 1 && tx.Type != 3 {
		valueBytes, err := currentTrie.TryGet(tx.FromAddr.Bytes())
		if err != nil {
			return nil, nil, errors.New("Account Value (" + tx.FromAddr.Hex() + ") TryGet Fail!" + err.Error())
		}
		getValue.SetBytes(valueBytes)

		if getValue.Int64() < tx.Value {
			return nil, nil, errors.New("Account Value (" + tx.FromAddr.Hex() + ") Not Enough!" + err.Error())
		}
	}

	// 更新输入资金状态
	if tx.Type != 2 {
		bag1 := currentTrie.Getproof(tx.FromAddr.Bytes())
		newValue1 := getValue.Int64() - tx.Value
		currentTrie.Update(tx.FromAddr.Bytes(), big.NewInt(newValue1).Bytes())
		bag2 := currentTrie.Getproof(tx.FromAddr.Bytes())

		transit1 := model.L2Transition{Addr: tx.FromAddr, Value: tx.Value, Style: 2}
		resTransition = append(resTransition, transit1)
		resBag = append(resBag, bag1, bag2)
	}

	// 更新输出资金状态
	var getValue2 big.Int
	if tx.Type != 1 {
		valueBytes2, err := currentTrie.TryGet(tx.ToAddr.Bytes())
		if err != nil {
			return nil, nil, errors.New("Account Value (" + tx.ToAddr.Hex() + ") TryGet Fail!" + err.Error())
		}
		getValue2.SetBytes(valueBytes2)

		bag1 := currentTrie.Getproof(tx.ToAddr.Bytes())
		newValue2 := getValue2.Int64() + tx.Value
		currentTrie.Update(tx.FromAddr.Bytes(), big.NewInt(newValue2).Bytes())
		bag2 := currentTrie.Getproof(tx.ToAddr.Bytes())

		transit2 := model.L2Transition{Addr: tx.ToAddr, Value: tx.Value, Style: 2}
		resTransition = append(resTransition, transit2)
		resBag = append(resBag, bag1, bag2)
	}

	return resTransition, resBag, nil
}

// L2TxSetGenerate 生成随机交易集合
func L2TxSetGenerate(num int) ([]*model.L2Transaction, error) {
	var txSet []*model.L2Transaction

	// 单个交易单个交易生成
	for i := 0; i < num; i++ {
		randTx, err := RandomL2Tx()
		if err != nil {
			return nil, err
		}
		txSet = append(txSet, randTx)
	}

	return txSet, nil
}

// RandomL2Tx 生成随机交易
func RandomL2Tx() (*model.L2Transaction, error) {
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

	// 设置为测试交易
	tx := model.L2Transaction{FromAddr: addrHash1, ToAddr: addrHash2, ChainID: 0, Type: 0, Value: 1}

	return &tx, nil
}

// GetLocalTransition 获取某批交易执行后的本地状态转换
func GetLocalTransition(txs []model.L2Transaction) []model.LocalTransition {
	var resLT []model.LocalTransition

	for i := 0; i < len(txs); i++ {
		if txs[i].Type == 1 {
			newLT := model.LocalTransition{
				Account: txs[i].ToAddr,
				Value:   big.NewInt(txs[i].Value),
				Style:   big.NewInt(2),
				ChainID: big.NewInt(txs[i].ChainID),
			}
			resLT = append(resLT, newLT)
		}

		if txs[i].Type == 2 {
			newLT := model.LocalTransition{
				Account: txs[i].FromAddr,
				Value:   big.NewInt(txs[i].Value),
				Style:   big.NewInt(1),
				ChainID: big.NewInt(txs[i].ChainID),
			}
			resLT = append(resLT, newLT)
		}
	}
	return resLT
}

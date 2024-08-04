package executor

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"rollup-offchain/model"
	"rollup-offchain/trie"
)

var UserRange = 100

// TxSetGenerate 生成随机交易集合
func TxSetGenerate(num int, fromNum int, toNum int) ([]*model.Transaction, error) {
	var txSet []*model.Transaction

	// 单个交易单个交易生成
	for i := 0; i < num; i++ {
		randTx, err := RandomTx(fromNum, toNum)
		if err != nil {
			return nil, err
		}
		txSet = append(txSet, randTx)
	}

	return txSet, nil
}

// RandomTx 生成随机交易
func RandomTx(fromNum int, toNum int) (*model.Transaction, error) {
	var fromCoins []model.CoinUnit
	for i := 0; i < fromNum; i++ {
		// 创建新的coin
		var newCoin model.CoinUnit

		// 生成随机地址
		randAddrInt, err := rand.Int(rand.Reader, big.NewInt(int64(UserRange)))
		if err != nil {
			fmt.Println("rand fail!", err.Error())
			return nil, err
		}
		addrHash := trie.BigToHash(randAddrInt)

		// 赋值
		copy(newCoin.Addr[:32], addrHash.Bytes())
		newCoin.Balance = 1
		// 更新coin列表
		fromCoins = append(fromCoins, newCoin)
	}

	var toCoins []model.CoinUnit
	for i := 0; i < toNum; i++ {
		// 创建新的coin
		var newCoin model.CoinUnit

		// 生成随机地址
		randAddrInt, err := rand.Int(rand.Reader, big.NewInt(int64(UserRange)))
		if err != nil {
			fmt.Println("rand fail!", err.Error())
			return nil, err
		}
		addrHash := trie.BigToHash(randAddrInt)

		// 赋值
		copy(newCoin.Addr[:32], addrHash.Bytes())
		newCoin.Balance = 1
		// 更新coin列表
		toCoins = append(toCoins, newCoin)
	}

	// 设置为测试交易
	tx := model.Transaction{fromCoins, toCoins, 0, 3}

	return &tx, nil
}

// TxsUpload 交易数据上链
func TxsUpload() {
	fmt.Println("TxsUpload")
}

// TxsExecute 执行交易
/*
	根据交易类型执行交易
	验证交易
	生成local状态集合，更新l2状态树

	// 不考虑执行失败的情况
*/
func TxsExecute(currentTrie *trie.Trie, txs []*model.Transaction) (*model.TxsResult, error) {
	fmt.Println("TxsExecute")
	var result model.TxsResult

	for i := 0; i < len(txs); i++ {
		if txs[i].Type == 0 { // 普通L2交易
			err := ExecuteOneTx(currentTrie, txs[i])
			if err != nil {
				idStr := string(rune(i))
				return nil, errors.New("tx" + idStr + " execute fail: " + err.Error())
			}

		} else if txs[i].Type == 1 { // L1-L2 交易
			err := ExecuteOneTx(currentTrie, txs[i])
			if err != nil {
				idStr := string(rune(i))
				return nil, errors.New("tx" + idStr + " execute fail!" + err.Error())
			}

			// add local coin to incoin
			result.InCoins = append(result.InCoins, txs[i].FromCoins[0])
		} else if txs[i].Type == 2 { // L2-L1 交易
			err := ExecuteOneTx(currentTrie, txs[i])
			if err != nil {
				idStr := string(rune(i))
				return nil, errors.New("tx" + idStr + " execute fail!" + err.Error())
			}

			// add local coin to out coin
			result.OutCoins = append(result.OutCoins, txs[i].ToCoins[0])

		} else if txs[i].Type == 3 { // 测试交易
			err := ExecuteOneTx(currentTrie, txs[i])
			if err != nil {
				idStr := string(rune(i))
				return nil, errors.New("tx" + idStr + " execute fail!" + err.Error())
			}
		} else {
			return nil, errors.New("illegal transaction type!")
		}
	}

	return &result, nil
}

// ResultUpload 执行结果上链
func ResultUpload() {
	fmt.Println("ResultUpload")

}

// ExecuteOneTx 执行一笔交易，如果资金不足则返回fail
func ExecuteOneTx(currentTrie *trie.Trie, tx *model.Transaction) error {
	// 检查交易合法性
	if tx.Type == 0 {
		sumFrom := int64(0)
		sumTo := int64(0)

		for i := 0; i < len(tx.FromCoins); i++ {
			sumFrom += tx.FromCoins[i].Balance
		}
		for i := 0; i < len(tx.ToCoins); i++ {
			sumTo += tx.ToCoins[i].Balance
		}

		if sumTo != sumFrom {
			return errors.New("input and output are inbalance")
		}
	}

	// 检查资金是否充足
	if tx.Type != 1 && tx.Type != 3 {
		for i := 0; i < len(tx.FromCoins); i++ {
			var addr []byte
			addr = tx.FromCoins[i].Addr[:32]

			var getValue big.Int
			valueBytes, err := currentTrie.TryGet(addr)
			if err != nil {
				return errors.New("Trie Value TruGet Fail!" + err.Error())
			}
			getValue.SetBytes(valueBytes)

			if getValue.Int64() < tx.FromCoins[i].Balance {
				return errors.New("balance not enough")
			}
		}
	}

	// 更新输入资金状态
	if tx.Type != 1 {
		for i := 0; i < len(tx.FromCoins); i++ {
			var addr []byte
			addr = tx.FromCoins[i].Addr[:32]

			// 获取该账户的值
			var getValue big.Int
			valueBytes, err := currentTrie.TryGet(addr)
			if err != nil {
				return errors.New("Trie Value TruGet Fail!" + err.Error())
			}
			getValue.SetBytes(valueBytes)

			// 计算并更新最新状态
			newBalance := big.NewInt(getValue.Int64() - tx.FromCoins[i].Balance)
			currentTrie.Update(addr, newBalance.Bytes())
		}
	}

	// 更新输出资金状态
	if tx.Type != 2 {
		for i := 0; i < len(tx.ToCoins); i++ {
			var addr []byte
			addr = tx.ToCoins[i].Addr[:32]

			// 获取该账户的值
			var getValue big.Int
			valueBytes, err := currentTrie.TryGet(addr)
			if err != nil {
				return errors.New("Trie Value TruGet Fail!" + err.Error())
			}
			getValue.SetBytes(valueBytes)

			// 计算并更新最新状态
			newBalance := big.NewInt(getValue.Int64() + tx.ToCoins[i].Balance)
			currentTrie.Update(addr, newBalance.Bytes())
		}
	}

	return nil
}

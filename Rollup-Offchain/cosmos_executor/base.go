package cosmos_executor

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"rollup-offchain/contract"
)

var MyAddr common.Address
var MyChainID *big.Int
var LSMInstance *contract.LocalStateManager
var CRInstance *contract.CrossRollup
var FRInstance *contract.FullRollup
var TEInstance *contract.TransactionExecutor
var DTInstance *contract.DataTypes

// RollupInit 初始化rollup状态
func RollupInit(client *ethclient.Client, myAddr common.Address, chainID *big.Int) {
	// 初始化挑战管理合约
	MyChainID = chainID
	MyAddr = myAddr

	DTAddr := common.HexToAddress(contract.MosDataTypes)
	DT, err := contract.NewDataTypes(DTAddr, client)
	if err != nil {
		log.Fatal(err)
	}
	DTInstance = DT

	//LSMAddr := common.HexToAddress(contract.LocalStateManagerAddr)
	//LSM, err := contract.NewLocalStateManager(LSMAddr, client)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//LSMInstance = LSM

	//TEAddr := common.HexToAddress(contract.TransactionExecutorAddr)
	//TE, err := contract.NewTransactionExecutor(TEAddr, client)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//TEInstance = TE

	//CRAddr := common.HexToAddress(contract.CrossRollupAddr)
	//CR, err := contract.NewCrossRollup(CRAddr, client)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//CRInstance = CR

	FRAddr := common.HexToAddress(contract.MosFullRollup)
	FR, err := contract.NewFullRollup(FRAddr, client)
	if err != nil {
		log.Fatal(err)
	}
	FRInstance = FR
}

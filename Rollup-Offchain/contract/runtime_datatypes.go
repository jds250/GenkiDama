package contract

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

const runtimeDataTypesABI = `[
  {
    "inputs": [
      {
        "components": [
          {"internalType":"string","name":"fromAddr","type":"string"},
          {"internalType":"string","name":"toAddr","type":"string"},
          {"internalType":"uint256","name":"value","type":"uint256"},
          {"internalType":"uint256","name":"chainID","type":"uint256"},
          {"internalType":"uint256","name":"style","type":"uint256"}
        ],
        "internalType":"structDataTypes.L2Transaction[]",
        "name":"_transactions",
        "type":"tuple[]"
      }
    ],
    "name":"TransactionsListHash",
    "outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],
    "stateMutability":"view",
    "type":"function"
  },
  {
    "inputs": [
      {
        "components": [
          {"internalType":"string","name":"account","type":"string"},
          {"internalType":"uint256","name":"value","type":"uint256"},
          {"internalType":"uint256","name":"style","type":"uint256"},
          {"internalType":"uint256","name":"chainID","type":"uint256"}
        ],
        "internalType":"structDataTypes.L2Transition[]",
        "name":"_transitions",
        "type":"tuple[]"
      }
    ],
    "name":"L2TransitionListHash",
    "outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],
    "stateMutability":"view",
    "type":"function"
  }
]`

func runtimeDataTypesContract(address common.Address, caller bind.ContractCaller) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(runtimeDataTypesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, nil, nil), nil
}

func HashTransactions(caller bind.ContractCaller, dataTypesAddress common.Address, transactions []DataTypesL2Transaction) ([32]byte, error) {
	bound, err := runtimeDataTypesContract(dataTypesAddress, caller)
	if err != nil {
		return [32]byte{}, err
	}

	var out []interface{}
	if err := bound.Call(&bind.CallOpts{}, &out, "TransactionsListHash", transactions); err != nil {
		return [32]byte{}, err
	}

	return *abi.ConvertType(out[0], new([32]byte)).(*[32]byte), nil
}

func HashTransitions(caller bind.ContractCaller, dataTypesAddress common.Address, transitions []DataTypesL2Transition) ([32]byte, error) {
	bound, err := runtimeDataTypesContract(dataTypesAddress, caller)
	if err != nil {
		return [32]byte{}, err
	}

	var out []interface{}
	if err := bound.Call(&bind.CallOpts{}, &out, "L2TransitionListHash", transitions); err != nil {
		return [32]byte{}, err
	}

	return *abi.ConvertType(out[0], new([32]byte)).(*[32]byte), nil
}

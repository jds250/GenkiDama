// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// DataTypesBlock is an auto generated low-level Go binding around an user-defined struct.
type DataTypesBlock struct {
	Height              *big.Int
	TxRoot              [32]byte
	LocalTransitionRoot [32]byte
	StateRoot           [32]byte
	BlockSize           *big.Int
	CommitTime          *big.Int
}

// TransactionExecutorMetaData contains all meta data concerning the TransactionExecutor contract.
var TransactionExecutorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_MPTVerifierAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_DataTypesAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_MerkleUtilsAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_LocalStateManagerAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"fromAddr\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"toAddr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"style\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.L2Transaction[]\",\"name\":\"_transactions\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"_tsRoot\",\"type\":\"bytes32\"}],\"name\":\"VerifyTransactions\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"fromAddr\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"toAddr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"style\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.L2Transaction[]\",\"name\":\"_transactions\",\"type\":\"tuple[]\"}],\"name\":\"GetL2Transitions\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"account\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"style\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.L2Transition[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"fromAddr\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"toAddr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"style\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.L2Transaction[]\",\"name\":\"_transactions\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"txRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"localTransitionRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"blockSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commitTime\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.Block\",\"name\":\"_targetBlock\",\"type\":\"tuple\"}],\"name\":\"GetLocalTransition\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockNum\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"account\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"style\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockTime\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.LocalTransition[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"b\",\"type\":\"bytes\"}],\"name\":\"bytesToUint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// TransactionExecutorABI is the input ABI used to generate the binding from.
// Deprecated: Use TransactionExecutorMetaData.ABI instead.
var TransactionExecutorABI = TransactionExecutorMetaData.ABI

// TransactionExecutor is an auto generated Go binding around an Ethereum contract.
type TransactionExecutor struct {
	TransactionExecutorCaller     // Read-only binding to the contract
	TransactionExecutorTransactor // Write-only binding to the contract
	TransactionExecutorFilterer   // Log filterer for contract events
}

// TransactionExecutorCaller is an auto generated read-only Go binding around an Ethereum contract.
type TransactionExecutorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransactionExecutorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TransactionExecutorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransactionExecutorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TransactionExecutorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransactionExecutorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TransactionExecutorSession struct {
	Contract     *TransactionExecutor // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// TransactionExecutorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TransactionExecutorCallerSession struct {
	Contract *TransactionExecutorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// TransactionExecutorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TransactionExecutorTransactorSession struct {
	Contract     *TransactionExecutorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// TransactionExecutorRaw is an auto generated low-level Go binding around an Ethereum contract.
type TransactionExecutorRaw struct {
	Contract *TransactionExecutor // Generic contract binding to access the raw methods on
}

// TransactionExecutorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TransactionExecutorCallerRaw struct {
	Contract *TransactionExecutorCaller // Generic read-only contract binding to access the raw methods on
}

// TransactionExecutorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TransactionExecutorTransactorRaw struct {
	Contract *TransactionExecutorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTransactionExecutor creates a new instance of TransactionExecutor, bound to a specific deployed contract.
func NewTransactionExecutor(address common.Address, backend bind.ContractBackend) (*TransactionExecutor, error) {
	contract, err := bindTransactionExecutor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TransactionExecutor{TransactionExecutorCaller: TransactionExecutorCaller{contract: contract}, TransactionExecutorTransactor: TransactionExecutorTransactor{contract: contract}, TransactionExecutorFilterer: TransactionExecutorFilterer{contract: contract}}, nil
}

// NewTransactionExecutorCaller creates a new read-only instance of TransactionExecutor, bound to a specific deployed contract.
func NewTransactionExecutorCaller(address common.Address, caller bind.ContractCaller) (*TransactionExecutorCaller, error) {
	contract, err := bindTransactionExecutor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TransactionExecutorCaller{contract: contract}, nil
}

// NewTransactionExecutorTransactor creates a new write-only instance of TransactionExecutor, bound to a specific deployed contract.
func NewTransactionExecutorTransactor(address common.Address, transactor bind.ContractTransactor) (*TransactionExecutorTransactor, error) {
	contract, err := bindTransactionExecutor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TransactionExecutorTransactor{contract: contract}, nil
}

// NewTransactionExecutorFilterer creates a new log filterer instance of TransactionExecutor, bound to a specific deployed contract.
func NewTransactionExecutorFilterer(address common.Address, filterer bind.ContractFilterer) (*TransactionExecutorFilterer, error) {
	contract, err := bindTransactionExecutor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TransactionExecutorFilterer{contract: contract}, nil
}

// bindTransactionExecutor binds a generic wrapper to an already deployed contract.
func bindTransactionExecutor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TransactionExecutorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TransactionExecutor *TransactionExecutorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TransactionExecutor.Contract.TransactionExecutorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TransactionExecutor *TransactionExecutorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TransactionExecutor.Contract.TransactionExecutorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TransactionExecutor *TransactionExecutorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TransactionExecutor.Contract.TransactionExecutorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TransactionExecutor *TransactionExecutorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TransactionExecutor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TransactionExecutor *TransactionExecutorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TransactionExecutor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TransactionExecutor *TransactionExecutorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TransactionExecutor.Contract.contract.Transact(opts, method, params...)
}

// GetL2Transitions is a free data retrieval call binding the contract method 0xbe4a1619.
//
// Solidity: function GetL2Transitions((string,string,uint256,uint256,uint256)[] _transactions) pure returns((string,uint256,uint256,uint256)[])
func (_TransactionExecutor *TransactionExecutorCaller) GetL2Transitions(opts *bind.CallOpts, _transactions []DataTypesL2Transaction) ([]DataTypesL2Transition, error) {
	var out []interface{}
	err := _TransactionExecutor.contract.Call(opts, &out, "GetL2Transitions", _transactions)

	if err != nil {
		return *new([]DataTypesL2Transition), err
	}

	out0 := *abi.ConvertType(out[0], new([]DataTypesL2Transition)).(*[]DataTypesL2Transition)

	return out0, err

}

// GetL2Transitions is a free data retrieval call binding the contract method 0xbe4a1619.
//
// Solidity: function GetL2Transitions((string,string,uint256,uint256,uint256)[] _transactions) pure returns((string,uint256,uint256,uint256)[])
func (_TransactionExecutor *TransactionExecutorSession) GetL2Transitions(_transactions []DataTypesL2Transaction) ([]DataTypesL2Transition, error) {
	return _TransactionExecutor.Contract.GetL2Transitions(&_TransactionExecutor.CallOpts, _transactions)
}

// GetL2Transitions is a free data retrieval call binding the contract method 0xbe4a1619.
//
// Solidity: function GetL2Transitions((string,string,uint256,uint256,uint256)[] _transactions) pure returns((string,uint256,uint256,uint256)[])
func (_TransactionExecutor *TransactionExecutorCallerSession) GetL2Transitions(_transactions []DataTypesL2Transaction) ([]DataTypesL2Transition, error) {
	return _TransactionExecutor.Contract.GetL2Transitions(&_TransactionExecutor.CallOpts, _transactions)
}

// GetLocalTransition is a free data retrieval call binding the contract method 0x084372a0.
//
// Solidity: function GetLocalTransition((string,string,uint256,uint256,uint256)[] _transactions, (uint256,bytes32,bytes32,bytes32,uint256,uint256) _targetBlock) pure returns((uint256,uint256,string,uint256,uint256,uint256)[])
func (_TransactionExecutor *TransactionExecutorCaller) GetLocalTransition(opts *bind.CallOpts, _transactions []DataTypesL2Transaction, _targetBlock DataTypesBlock) ([]DataTypesLocalTransition, error) {
	var out []interface{}
	err := _TransactionExecutor.contract.Call(opts, &out, "GetLocalTransition", _transactions, _targetBlock)

	if err != nil {
		return *new([]DataTypesLocalTransition), err
	}

	out0 := *abi.ConvertType(out[0], new([]DataTypesLocalTransition)).(*[]DataTypesLocalTransition)

	return out0, err

}

// GetLocalTransition is a free data retrieval call binding the contract method 0x084372a0.
//
// Solidity: function GetLocalTransition((string,string,uint256,uint256,uint256)[] _transactions, (uint256,bytes32,bytes32,bytes32,uint256,uint256) _targetBlock) pure returns((uint256,uint256,string,uint256,uint256,uint256)[])
func (_TransactionExecutor *TransactionExecutorSession) GetLocalTransition(_transactions []DataTypesL2Transaction, _targetBlock DataTypesBlock) ([]DataTypesLocalTransition, error) {
	return _TransactionExecutor.Contract.GetLocalTransition(&_TransactionExecutor.CallOpts, _transactions, _targetBlock)
}

// GetLocalTransition is a free data retrieval call binding the contract method 0x084372a0.
//
// Solidity: function GetLocalTransition((string,string,uint256,uint256,uint256)[] _transactions, (uint256,bytes32,bytes32,bytes32,uint256,uint256) _targetBlock) pure returns((uint256,uint256,string,uint256,uint256,uint256)[])
func (_TransactionExecutor *TransactionExecutorCallerSession) GetLocalTransition(_transactions []DataTypesL2Transaction, _targetBlock DataTypesBlock) ([]DataTypesLocalTransition, error) {
	return _TransactionExecutor.Contract.GetLocalTransition(&_TransactionExecutor.CallOpts, _transactions, _targetBlock)
}

// VerifyTransactions is a free data retrieval call binding the contract method 0xd7b53ba1.
//
// Solidity: function VerifyTransactions((string,string,uint256,uint256,uint256)[] _transactions, bytes32 _tsRoot) view returns(bool, string)
func (_TransactionExecutor *TransactionExecutorCaller) VerifyTransactions(opts *bind.CallOpts, _transactions []DataTypesL2Transaction, _tsRoot [32]byte) (bool, string, error) {
	var out []interface{}
	err := _TransactionExecutor.contract.Call(opts, &out, "VerifyTransactions", _transactions, _tsRoot)

	if err != nil {
		return *new(bool), *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)

	return out0, out1, err

}

// VerifyTransactions is a free data retrieval call binding the contract method 0xd7b53ba1.
//
// Solidity: function VerifyTransactions((string,string,uint256,uint256,uint256)[] _transactions, bytes32 _tsRoot) view returns(bool, string)
func (_TransactionExecutor *TransactionExecutorSession) VerifyTransactions(_transactions []DataTypesL2Transaction, _tsRoot [32]byte) (bool, string, error) {
	return _TransactionExecutor.Contract.VerifyTransactions(&_TransactionExecutor.CallOpts, _transactions, _tsRoot)
}

// VerifyTransactions is a free data retrieval call binding the contract method 0xd7b53ba1.
//
// Solidity: function VerifyTransactions((string,string,uint256,uint256,uint256)[] _transactions, bytes32 _tsRoot) view returns(bool, string)
func (_TransactionExecutor *TransactionExecutorCallerSession) VerifyTransactions(_transactions []DataTypesL2Transaction, _tsRoot [32]byte) (bool, string, error) {
	return _TransactionExecutor.Contract.VerifyTransactions(&_TransactionExecutor.CallOpts, _transactions, _tsRoot)
}

// BytesToUint is a free data retrieval call binding the contract method 0x02d06d05.
//
// Solidity: function bytesToUint(bytes b) pure returns(uint256)
func (_TransactionExecutor *TransactionExecutorCaller) BytesToUint(opts *bind.CallOpts, b []byte) (*big.Int, error) {
	var out []interface{}
	err := _TransactionExecutor.contract.Call(opts, &out, "bytesToUint", b)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BytesToUint is a free data retrieval call binding the contract method 0x02d06d05.
//
// Solidity: function bytesToUint(bytes b) pure returns(uint256)
func (_TransactionExecutor *TransactionExecutorSession) BytesToUint(b []byte) (*big.Int, error) {
	return _TransactionExecutor.Contract.BytesToUint(&_TransactionExecutor.CallOpts, b)
}

// BytesToUint is a free data retrieval call binding the contract method 0x02d06d05.
//
// Solidity: function bytesToUint(bytes b) pure returns(uint256)
func (_TransactionExecutor *TransactionExecutorCallerSession) BytesToUint(b []byte) (*big.Int, error) {
	return _TransactionExecutor.Contract.BytesToUint(&_TransactionExecutor.CallOpts, b)
}

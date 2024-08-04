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

// DataTypesL2Transaction is an auto generated low-level Go binding around an user-defined struct.
type DataTypesL2Transaction struct {
	FromAddr common.Address
	ToAddr   common.Address
	Value    *big.Int
	ChainID  *big.Int
	Style    *big.Int
}

// DataTypesL2Transition is an auto generated low-level Go binding around an user-defined struct.
type DataTypesL2Transition struct {
	Account common.Address
	Value   *big.Int
	Style   *big.Int
	ChainID *big.Int
}

// DataTypesLocalState is an auto generated low-level Go binding around an user-defined struct.
type DataTypesLocalState struct {
	Account    common.Address
	Value      *big.Int
	Lock       *big.Int
	UnLockTime *big.Int
}

// DataTypesLocalTransition is an auto generated low-level Go binding around an user-defined struct.
type DataTypesLocalTransition struct {
	ChainID    *big.Int
	BlockNum   *big.Int
	Account    common.Address
	Value      *big.Int
	Style      *big.Int
	UnlockTime *big.Int
}

// DataTypesMetaData contains all meta data concerning the DataTypes contract.
var DataTypesMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"style\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.L2Transition\",\"name\":\"_l2t\",\"type\":\"tuple\"}],\"name\":\"L2TransitionToBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_ltBytes\",\"type\":\"bytes\"}],\"name\":\"BytesToL2Transition\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"style\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.L2Transition\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockNum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"style\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockTime\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.LocalTransition\",\"name\":\"_lt\",\"type\":\"tuple\"}],\"name\":\"LocalTransitionToBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_ltBytes\",\"type\":\"bytes\"}],\"name\":\"BytesToLocalTransition\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockNum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"style\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockTime\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.LocalTransition\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unLockTime\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.LocalState\",\"name\":\"_ls\",\"type\":\"tuple\"}],\"name\":\"LocalStateToBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_ltBytes\",\"type\":\"bytes\"}],\"name\":\"BytesToLocalState\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unLockTime\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.LocalState\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"fromAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"style\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.L2Transaction\",\"name\":\"_tx\",\"type\":\"tuple\"}],\"name\":\"L2TransactionToBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_txBytes\",\"type\":\"bytes\"}],\"name\":\"BytesToL2Transaction\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"fromAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"style\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.L2Transaction\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_strs\",\"type\":\"bytes[]\"}],\"name\":\"BytesList2Hash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"style\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.L2Transition[]\",\"name\":\"_transitions\",\"type\":\"tuple[]\"}],\"name\":\"L2TransitionListHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"fromAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"style\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.L2Transaction[]\",\"name\":\"_transactions\",\"type\":\"tuple[]\"}],\"name\":\"TransactionsListHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockNum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"style\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockTime\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.LocalTransition[]\",\"name\":\"_transitions\",\"type\":\"tuple[]\"}],\"name\":\"L2TransitionsEncode\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockNum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"style\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockTime\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.LocalTransition[]\",\"name\":\"_transitions\",\"type\":\"tuple[]\"}],\"name\":\"LocalTransitionsEncode\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"fromAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"style\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.L2Transaction[]\",\"name\":\"_transactions\",\"type\":\"tuple[]\"}],\"name\":\"L2TransactionsEncode\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// DataTypesABI is the input ABI used to generate the binding from.
// Deprecated: Use DataTypesMetaData.ABI instead.
var DataTypesABI = DataTypesMetaData.ABI

// DataTypes is an auto generated Go binding around an Ethereum contract.
type DataTypes struct {
	DataTypesCaller     // Read-only binding to the contract
	DataTypesTransactor // Write-only binding to the contract
	DataTypesFilterer   // Log filterer for contract events
}

// DataTypesCaller is an auto generated read-only Go binding around an Ethereum contract.
type DataTypesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataTypesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DataTypesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataTypesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DataTypesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataTypesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DataTypesSession struct {
	Contract     *DataTypes        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DataTypesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DataTypesCallerSession struct {
	Contract *DataTypesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// DataTypesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DataTypesTransactorSession struct {
	Contract     *DataTypesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// DataTypesRaw is an auto generated low-level Go binding around an Ethereum contract.
type DataTypesRaw struct {
	Contract *DataTypes // Generic contract binding to access the raw methods on
}

// DataTypesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DataTypesCallerRaw struct {
	Contract *DataTypesCaller // Generic read-only contract binding to access the raw methods on
}

// DataTypesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DataTypesTransactorRaw struct {
	Contract *DataTypesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDataTypes creates a new instance of DataTypes, bound to a specific deployed contract.
func NewDataTypes(address common.Address, backend bind.ContractBackend) (*DataTypes, error) {
	contract, err := bindDataTypes(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DataTypes{DataTypesCaller: DataTypesCaller{contract: contract}, DataTypesTransactor: DataTypesTransactor{contract: contract}, DataTypesFilterer: DataTypesFilterer{contract: contract}}, nil
}

// NewDataTypesCaller creates a new read-only instance of DataTypes, bound to a specific deployed contract.
func NewDataTypesCaller(address common.Address, caller bind.ContractCaller) (*DataTypesCaller, error) {
	contract, err := bindDataTypes(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DataTypesCaller{contract: contract}, nil
}

// NewDataTypesTransactor creates a new write-only instance of DataTypes, bound to a specific deployed contract.
func NewDataTypesTransactor(address common.Address, transactor bind.ContractTransactor) (*DataTypesTransactor, error) {
	contract, err := bindDataTypes(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DataTypesTransactor{contract: contract}, nil
}

// NewDataTypesFilterer creates a new log filterer instance of DataTypes, bound to a specific deployed contract.
func NewDataTypesFilterer(address common.Address, filterer bind.ContractFilterer) (*DataTypesFilterer, error) {
	contract, err := bindDataTypes(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DataTypesFilterer{contract: contract}, nil
}

// bindDataTypes binds a generic wrapper to an already deployed contract.
func bindDataTypes(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DataTypesMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DataTypes *DataTypesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DataTypes.Contract.DataTypesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DataTypes *DataTypesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DataTypes.Contract.DataTypesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DataTypes *DataTypesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DataTypes.Contract.DataTypesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DataTypes *DataTypesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DataTypes.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DataTypes *DataTypesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DataTypes.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DataTypes *DataTypesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DataTypes.Contract.contract.Transact(opts, method, params...)
}

// BytesList2Hash is a free data retrieval call binding the contract method 0x7192ff8a.
//
// Solidity: function BytesList2Hash(bytes[] _strs) view returns(bytes32)
func (_DataTypes *DataTypesCaller) BytesList2Hash(opts *bind.CallOpts, _strs [][]byte) ([32]byte, error) {
	var out []interface{}
	err := _DataTypes.contract.Call(opts, &out, "BytesList2Hash", _strs)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BytesList2Hash is a free data retrieval call binding the contract method 0x7192ff8a.
//
// Solidity: function BytesList2Hash(bytes[] _strs) view returns(bytes32)
func (_DataTypes *DataTypesSession) BytesList2Hash(_strs [][]byte) ([32]byte, error) {
	return _DataTypes.Contract.BytesList2Hash(&_DataTypes.CallOpts, _strs)
}

// BytesList2Hash is a free data retrieval call binding the contract method 0x7192ff8a.
//
// Solidity: function BytesList2Hash(bytes[] _strs) view returns(bytes32)
func (_DataTypes *DataTypesCallerSession) BytesList2Hash(_strs [][]byte) ([32]byte, error) {
	return _DataTypes.Contract.BytesList2Hash(&_DataTypes.CallOpts, _strs)
}

// BytesToL2Transaction is a free data retrieval call binding the contract method 0x9d4db57f.
//
// Solidity: function BytesToL2Transaction(bytes _txBytes) pure returns((address,address,uint256,uint256,uint256))
func (_DataTypes *DataTypesCaller) BytesToL2Transaction(opts *bind.CallOpts, _txBytes []byte) (DataTypesL2Transaction, error) {
	var out []interface{}
	err := _DataTypes.contract.Call(opts, &out, "BytesToL2Transaction", _txBytes)

	if err != nil {
		return *new(DataTypesL2Transaction), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesL2Transaction)).(*DataTypesL2Transaction)

	return out0, err

}

// BytesToL2Transaction is a free data retrieval call binding the contract method 0x9d4db57f.
//
// Solidity: function BytesToL2Transaction(bytes _txBytes) pure returns((address,address,uint256,uint256,uint256))
func (_DataTypes *DataTypesSession) BytesToL2Transaction(_txBytes []byte) (DataTypesL2Transaction, error) {
	return _DataTypes.Contract.BytesToL2Transaction(&_DataTypes.CallOpts, _txBytes)
}

// BytesToL2Transaction is a free data retrieval call binding the contract method 0x9d4db57f.
//
// Solidity: function BytesToL2Transaction(bytes _txBytes) pure returns((address,address,uint256,uint256,uint256))
func (_DataTypes *DataTypesCallerSession) BytesToL2Transaction(_txBytes []byte) (DataTypesL2Transaction, error) {
	return _DataTypes.Contract.BytesToL2Transaction(&_DataTypes.CallOpts, _txBytes)
}

// BytesToL2Transition is a free data retrieval call binding the contract method 0x88f27efe.
//
// Solidity: function BytesToL2Transition(bytes _ltBytes) pure returns((address,uint256,uint256,uint256))
func (_DataTypes *DataTypesCaller) BytesToL2Transition(opts *bind.CallOpts, _ltBytes []byte) (DataTypesL2Transition, error) {
	var out []interface{}
	err := _DataTypes.contract.Call(opts, &out, "BytesToL2Transition", _ltBytes)

	if err != nil {
		return *new(DataTypesL2Transition), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesL2Transition)).(*DataTypesL2Transition)

	return out0, err

}

// BytesToL2Transition is a free data retrieval call binding the contract method 0x88f27efe.
//
// Solidity: function BytesToL2Transition(bytes _ltBytes) pure returns((address,uint256,uint256,uint256))
func (_DataTypes *DataTypesSession) BytesToL2Transition(_ltBytes []byte) (DataTypesL2Transition, error) {
	return _DataTypes.Contract.BytesToL2Transition(&_DataTypes.CallOpts, _ltBytes)
}

// BytesToL2Transition is a free data retrieval call binding the contract method 0x88f27efe.
//
// Solidity: function BytesToL2Transition(bytes _ltBytes) pure returns((address,uint256,uint256,uint256))
func (_DataTypes *DataTypesCallerSession) BytesToL2Transition(_ltBytes []byte) (DataTypesL2Transition, error) {
	return _DataTypes.Contract.BytesToL2Transition(&_DataTypes.CallOpts, _ltBytes)
}

// BytesToLocalState is a free data retrieval call binding the contract method 0x94e591d4.
//
// Solidity: function BytesToLocalState(bytes _ltBytes) pure returns((address,uint256,uint256,uint256))
func (_DataTypes *DataTypesCaller) BytesToLocalState(opts *bind.CallOpts, _ltBytes []byte) (DataTypesLocalState, error) {
	var out []interface{}
	err := _DataTypes.contract.Call(opts, &out, "BytesToLocalState", _ltBytes)

	if err != nil {
		return *new(DataTypesLocalState), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesLocalState)).(*DataTypesLocalState)

	return out0, err

}

// BytesToLocalState is a free data retrieval call binding the contract method 0x94e591d4.
//
// Solidity: function BytesToLocalState(bytes _ltBytes) pure returns((address,uint256,uint256,uint256))
func (_DataTypes *DataTypesSession) BytesToLocalState(_ltBytes []byte) (DataTypesLocalState, error) {
	return _DataTypes.Contract.BytesToLocalState(&_DataTypes.CallOpts, _ltBytes)
}

// BytesToLocalState is a free data retrieval call binding the contract method 0x94e591d4.
//
// Solidity: function BytesToLocalState(bytes _ltBytes) pure returns((address,uint256,uint256,uint256))
func (_DataTypes *DataTypesCallerSession) BytesToLocalState(_ltBytes []byte) (DataTypesLocalState, error) {
	return _DataTypes.Contract.BytesToLocalState(&_DataTypes.CallOpts, _ltBytes)
}

// BytesToLocalTransition is a free data retrieval call binding the contract method 0xc09628eb.
//
// Solidity: function BytesToLocalTransition(bytes _ltBytes) pure returns((uint256,uint256,address,uint256,uint256,uint256))
func (_DataTypes *DataTypesCaller) BytesToLocalTransition(opts *bind.CallOpts, _ltBytes []byte) (DataTypesLocalTransition, error) {
	var out []interface{}
	err := _DataTypes.contract.Call(opts, &out, "BytesToLocalTransition", _ltBytes)

	if err != nil {
		return *new(DataTypesLocalTransition), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesLocalTransition)).(*DataTypesLocalTransition)

	return out0, err

}

// BytesToLocalTransition is a free data retrieval call binding the contract method 0xc09628eb.
//
// Solidity: function BytesToLocalTransition(bytes _ltBytes) pure returns((uint256,uint256,address,uint256,uint256,uint256))
func (_DataTypes *DataTypesSession) BytesToLocalTransition(_ltBytes []byte) (DataTypesLocalTransition, error) {
	return _DataTypes.Contract.BytesToLocalTransition(&_DataTypes.CallOpts, _ltBytes)
}

// BytesToLocalTransition is a free data retrieval call binding the contract method 0xc09628eb.
//
// Solidity: function BytesToLocalTransition(bytes _ltBytes) pure returns((uint256,uint256,address,uint256,uint256,uint256))
func (_DataTypes *DataTypesCallerSession) BytesToLocalTransition(_ltBytes []byte) (DataTypesLocalTransition, error) {
	return _DataTypes.Contract.BytesToLocalTransition(&_DataTypes.CallOpts, _ltBytes)
}

// L2TransactionToBytes is a free data retrieval call binding the contract method 0x59dfd602.
//
// Solidity: function L2TransactionToBytes((address,address,uint256,uint256,uint256) _tx) pure returns(bytes)
func (_DataTypes *DataTypesCaller) L2TransactionToBytes(opts *bind.CallOpts, _tx DataTypesL2Transaction) ([]byte, error) {
	var out []interface{}
	err := _DataTypes.contract.Call(opts, &out, "L2TransactionToBytes", _tx)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// L2TransactionToBytes is a free data retrieval call binding the contract method 0x59dfd602.
//
// Solidity: function L2TransactionToBytes((address,address,uint256,uint256,uint256) _tx) pure returns(bytes)
func (_DataTypes *DataTypesSession) L2TransactionToBytes(_tx DataTypesL2Transaction) ([]byte, error) {
	return _DataTypes.Contract.L2TransactionToBytes(&_DataTypes.CallOpts, _tx)
}

// L2TransactionToBytes is a free data retrieval call binding the contract method 0x59dfd602.
//
// Solidity: function L2TransactionToBytes((address,address,uint256,uint256,uint256) _tx) pure returns(bytes)
func (_DataTypes *DataTypesCallerSession) L2TransactionToBytes(_tx DataTypesL2Transaction) ([]byte, error) {
	return _DataTypes.Contract.L2TransactionToBytes(&_DataTypes.CallOpts, _tx)
}

// L2TransactionsEncode is a free data retrieval call binding the contract method 0x0e32b31d.
//
// Solidity: function L2TransactionsEncode((address,address,uint256,uint256,uint256)[] _transactions) view returns(bytes[])
func (_DataTypes *DataTypesCaller) L2TransactionsEncode(opts *bind.CallOpts, _transactions []DataTypesL2Transaction) ([][]byte, error) {
	var out []interface{}
	err := _DataTypes.contract.Call(opts, &out, "L2TransactionsEncode", _transactions)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// L2TransactionsEncode is a free data retrieval call binding the contract method 0x0e32b31d.
//
// Solidity: function L2TransactionsEncode((address,address,uint256,uint256,uint256)[] _transactions) view returns(bytes[])
func (_DataTypes *DataTypesSession) L2TransactionsEncode(_transactions []DataTypesL2Transaction) ([][]byte, error) {
	return _DataTypes.Contract.L2TransactionsEncode(&_DataTypes.CallOpts, _transactions)
}

// L2TransactionsEncode is a free data retrieval call binding the contract method 0x0e32b31d.
//
// Solidity: function L2TransactionsEncode((address,address,uint256,uint256,uint256)[] _transactions) view returns(bytes[])
func (_DataTypes *DataTypesCallerSession) L2TransactionsEncode(_transactions []DataTypesL2Transaction) ([][]byte, error) {
	return _DataTypes.Contract.L2TransactionsEncode(&_DataTypes.CallOpts, _transactions)
}

// L2TransitionListHash is a free data retrieval call binding the contract method 0xcd792992.
//
// Solidity: function L2TransitionListHash((address,uint256,uint256,uint256)[] _transitions) view returns(bytes32)
func (_DataTypes *DataTypesCaller) L2TransitionListHash(opts *bind.CallOpts, _transitions []DataTypesL2Transition) ([32]byte, error) {
	var out []interface{}
	err := _DataTypes.contract.Call(opts, &out, "L2TransitionListHash", _transitions)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// L2TransitionListHash is a free data retrieval call binding the contract method 0xcd792992.
//
// Solidity: function L2TransitionListHash((address,uint256,uint256,uint256)[] _transitions) view returns(bytes32)
func (_DataTypes *DataTypesSession) L2TransitionListHash(_transitions []DataTypesL2Transition) ([32]byte, error) {
	return _DataTypes.Contract.L2TransitionListHash(&_DataTypes.CallOpts, _transitions)
}

// L2TransitionListHash is a free data retrieval call binding the contract method 0xcd792992.
//
// Solidity: function L2TransitionListHash((address,uint256,uint256,uint256)[] _transitions) view returns(bytes32)
func (_DataTypes *DataTypesCallerSession) L2TransitionListHash(_transitions []DataTypesL2Transition) ([32]byte, error) {
	return _DataTypes.Contract.L2TransitionListHash(&_DataTypes.CallOpts, _transitions)
}

// L2TransitionToBytes is a free data retrieval call binding the contract method 0xab4bed0a.
//
// Solidity: function L2TransitionToBytes((address,uint256,uint256,uint256) _l2t) pure returns(bytes)
func (_DataTypes *DataTypesCaller) L2TransitionToBytes(opts *bind.CallOpts, _l2t DataTypesL2Transition) ([]byte, error) {
	var out []interface{}
	err := _DataTypes.contract.Call(opts, &out, "L2TransitionToBytes", _l2t)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// L2TransitionToBytes is a free data retrieval call binding the contract method 0xab4bed0a.
//
// Solidity: function L2TransitionToBytes((address,uint256,uint256,uint256) _l2t) pure returns(bytes)
func (_DataTypes *DataTypesSession) L2TransitionToBytes(_l2t DataTypesL2Transition) ([]byte, error) {
	return _DataTypes.Contract.L2TransitionToBytes(&_DataTypes.CallOpts, _l2t)
}

// L2TransitionToBytes is a free data retrieval call binding the contract method 0xab4bed0a.
//
// Solidity: function L2TransitionToBytes((address,uint256,uint256,uint256) _l2t) pure returns(bytes)
func (_DataTypes *DataTypesCallerSession) L2TransitionToBytes(_l2t DataTypesL2Transition) ([]byte, error) {
	return _DataTypes.Contract.L2TransitionToBytes(&_DataTypes.CallOpts, _l2t)
}

// L2TransitionsEncode is a free data retrieval call binding the contract method 0x0e114f88.
//
// Solidity: function L2TransitionsEncode((uint256,uint256,address,uint256,uint256,uint256)[] _transitions) view returns(bytes[])
func (_DataTypes *DataTypesCaller) L2TransitionsEncode(opts *bind.CallOpts, _transitions []DataTypesLocalTransition) ([][]byte, error) {
	var out []interface{}
	err := _DataTypes.contract.Call(opts, &out, "L2TransitionsEncode", _transitions)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// L2TransitionsEncode is a free data retrieval call binding the contract method 0x0e114f88.
//
// Solidity: function L2TransitionsEncode((uint256,uint256,address,uint256,uint256,uint256)[] _transitions) view returns(bytes[])
func (_DataTypes *DataTypesSession) L2TransitionsEncode(_transitions []DataTypesLocalTransition) ([][]byte, error) {
	return _DataTypes.Contract.L2TransitionsEncode(&_DataTypes.CallOpts, _transitions)
}

// L2TransitionsEncode is a free data retrieval call binding the contract method 0x0e114f88.
//
// Solidity: function L2TransitionsEncode((uint256,uint256,address,uint256,uint256,uint256)[] _transitions) view returns(bytes[])
func (_DataTypes *DataTypesCallerSession) L2TransitionsEncode(_transitions []DataTypesLocalTransition) ([][]byte, error) {
	return _DataTypes.Contract.L2TransitionsEncode(&_DataTypes.CallOpts, _transitions)
}

// LocalStateToBytes is a free data retrieval call binding the contract method 0x3731f9c1.
//
// Solidity: function LocalStateToBytes((address,uint256,uint256,uint256) _ls) pure returns(bytes)
func (_DataTypes *DataTypesCaller) LocalStateToBytes(opts *bind.CallOpts, _ls DataTypesLocalState) ([]byte, error) {
	var out []interface{}
	err := _DataTypes.contract.Call(opts, &out, "LocalStateToBytes", _ls)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// LocalStateToBytes is a free data retrieval call binding the contract method 0x3731f9c1.
//
// Solidity: function LocalStateToBytes((address,uint256,uint256,uint256) _ls) pure returns(bytes)
func (_DataTypes *DataTypesSession) LocalStateToBytes(_ls DataTypesLocalState) ([]byte, error) {
	return _DataTypes.Contract.LocalStateToBytes(&_DataTypes.CallOpts, _ls)
}

// LocalStateToBytes is a free data retrieval call binding the contract method 0x3731f9c1.
//
// Solidity: function LocalStateToBytes((address,uint256,uint256,uint256) _ls) pure returns(bytes)
func (_DataTypes *DataTypesCallerSession) LocalStateToBytes(_ls DataTypesLocalState) ([]byte, error) {
	return _DataTypes.Contract.LocalStateToBytes(&_DataTypes.CallOpts, _ls)
}

// LocalTransitionToBytes is a free data retrieval call binding the contract method 0x1a8c1683.
//
// Solidity: function LocalTransitionToBytes((uint256,uint256,address,uint256,uint256,uint256) _lt) pure returns(bytes)
func (_DataTypes *DataTypesCaller) LocalTransitionToBytes(opts *bind.CallOpts, _lt DataTypesLocalTransition) ([]byte, error) {
	var out []interface{}
	err := _DataTypes.contract.Call(opts, &out, "LocalTransitionToBytes", _lt)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// LocalTransitionToBytes is a free data retrieval call binding the contract method 0x1a8c1683.
//
// Solidity: function LocalTransitionToBytes((uint256,uint256,address,uint256,uint256,uint256) _lt) pure returns(bytes)
func (_DataTypes *DataTypesSession) LocalTransitionToBytes(_lt DataTypesLocalTransition) ([]byte, error) {
	return _DataTypes.Contract.LocalTransitionToBytes(&_DataTypes.CallOpts, _lt)
}

// LocalTransitionToBytes is a free data retrieval call binding the contract method 0x1a8c1683.
//
// Solidity: function LocalTransitionToBytes((uint256,uint256,address,uint256,uint256,uint256) _lt) pure returns(bytes)
func (_DataTypes *DataTypesCallerSession) LocalTransitionToBytes(_lt DataTypesLocalTransition) ([]byte, error) {
	return _DataTypes.Contract.LocalTransitionToBytes(&_DataTypes.CallOpts, _lt)
}

// LocalTransitionsEncode is a free data retrieval call binding the contract method 0x85c6bf27.
//
// Solidity: function LocalTransitionsEncode((uint256,uint256,address,uint256,uint256,uint256)[] _transitions) view returns(bytes[])
func (_DataTypes *DataTypesCaller) LocalTransitionsEncode(opts *bind.CallOpts, _transitions []DataTypesLocalTransition) ([][]byte, error) {
	var out []interface{}
	err := _DataTypes.contract.Call(opts, &out, "LocalTransitionsEncode", _transitions)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// LocalTransitionsEncode is a free data retrieval call binding the contract method 0x85c6bf27.
//
// Solidity: function LocalTransitionsEncode((uint256,uint256,address,uint256,uint256,uint256)[] _transitions) view returns(bytes[])
func (_DataTypes *DataTypesSession) LocalTransitionsEncode(_transitions []DataTypesLocalTransition) ([][]byte, error) {
	return _DataTypes.Contract.LocalTransitionsEncode(&_DataTypes.CallOpts, _transitions)
}

// LocalTransitionsEncode is a free data retrieval call binding the contract method 0x85c6bf27.
//
// Solidity: function LocalTransitionsEncode((uint256,uint256,address,uint256,uint256,uint256)[] _transitions) view returns(bytes[])
func (_DataTypes *DataTypesCallerSession) LocalTransitionsEncode(_transitions []DataTypesLocalTransition) ([][]byte, error) {
	return _DataTypes.Contract.LocalTransitionsEncode(&_DataTypes.CallOpts, _transitions)
}

// TransactionsListHash is a free data retrieval call binding the contract method 0x8b62ed3e.
//
// Solidity: function TransactionsListHash((address,address,uint256,uint256,uint256)[] _transactions) view returns(bytes32)
func (_DataTypes *DataTypesCaller) TransactionsListHash(opts *bind.CallOpts, _transactions []DataTypesL2Transaction) ([32]byte, error) {
	var out []interface{}
	err := _DataTypes.contract.Call(opts, &out, "TransactionsListHash", _transactions)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TransactionsListHash is a free data retrieval call binding the contract method 0x8b62ed3e.
//
// Solidity: function TransactionsListHash((address,address,uint256,uint256,uint256)[] _transactions) view returns(bytes32)
func (_DataTypes *DataTypesSession) TransactionsListHash(_transactions []DataTypesL2Transaction) ([32]byte, error) {
	return _DataTypes.Contract.TransactionsListHash(&_DataTypes.CallOpts, _transactions)
}

// TransactionsListHash is a free data retrieval call binding the contract method 0x8b62ed3e.
//
// Solidity: function TransactionsListHash((address,address,uint256,uint256,uint256)[] _transactions) view returns(bytes32)
func (_DataTypes *DataTypesCallerSession) TransactionsListHash(_transactions []DataTypesL2Transaction) ([32]byte, error) {
	return _DataTypes.Contract.TransactionsListHash(&_DataTypes.CallOpts, _transactions)
}

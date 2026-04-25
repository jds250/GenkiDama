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

// FullRollupMetaData contains all meta data concerning the FullRollup contract.
var FullRollupMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_committerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_MerkleUtilsAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_DataTypesAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_MPTVerifierAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"state\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"error\",\"type\":\"string\"}],\"name\":\"ErrorNotify\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"style\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"LocalStateNotify\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"desc\",\"type\":\"string\"}],\"name\":\"RollupBlockNotify\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DISPUTE_TIME\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WITHDRAW_WAIT_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ZERO_BYTES32\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"challengePool\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l2BlockID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l1BlockID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"l1StateRoot\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"l1BlockData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"challengeSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"committerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"confirmedNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"localStatePool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lock\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockTransitionNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lockTransitionPool\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockNum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"style\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"superAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_committerAddress\",\"type\":\"address\"}],\"name\":\"setCommitterAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BlockLen\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_txRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_localTransitionRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_transactions\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"style\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.L2Transition[]\",\"name\":\"_transitionList\",\"type\":\"tuple[]\"}],\"name\":\"commitBlock\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkCommit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_targetBlockHeight\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"fromAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"style\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.L2Transaction[]\",\"name\":\"_transactions\",\"type\":\"tuple[]\"}],\"name\":\"ExecutionChallenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_targetBlockHeight\",\"type\":\"uint256\"}],\"name\":\"rollBackBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"depositCoin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"refundCoin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"getLocalState\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lock\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.LocalState\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"style\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.L2Transition[]\",\"name\":\"_transits\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"_blockNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_lockTime\",\"type\":\"uint256\"}],\"name\":\"verifyAndLockLocalState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"style\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.L2Transition[]\",\"name\":\"_transits\",\"type\":\"tuple[]\"}],\"name\":\"verifyLocalState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockNum\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"transitionDatas\",\"type\":\"bytes\"}],\"name\":\"confirmLocalState\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockNum\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"style\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.L2Transition[]\",\"name\":\"_transits\",\"type\":\"tuple[]\"}],\"name\":\"removeLocalStateWithBlockNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"fromAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"style\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.L2Transaction[]\",\"name\":\"_transactions\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"_tsRoot\",\"type\":\"bytes32\"}],\"name\":\"VerifyTransactions\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"fromAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"style\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.L2Transaction[]\",\"name\":\"_transactions\",\"type\":\"tuple[]\"}],\"name\":\"GetL2Transitions\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"style\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.L2Transition[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"fromAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"style\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.L2Transaction[]\",\"name\":\"_transactions\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"txRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"localTransitionRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"blockSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commitTime\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.Block\",\"name\":\"_targetBlock\",\"type\":\"tuple\"}],\"name\":\"GetLocalTransition\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockNum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"style\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockTime\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.LocalTransition[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// FullRollupABI is the input ABI used to generate the binding from.
// Deprecated: Use FullRollupMetaData.ABI instead.
var FullRollupABI = FullRollupMetaData.ABI

// FullRollup is an auto generated Go binding around an Ethereum contract.
type FullRollup struct {
	FullRollupCaller     // Read-only binding to the contract
	FullRollupTransactor // Write-only binding to the contract
	FullRollupFilterer   // Log filterer for contract events
}

// FullRollupCaller is an auto generated read-only Go binding around an Ethereum contract.
type FullRollupCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FullRollupTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FullRollupTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FullRollupFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FullRollupFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FullRollupSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FullRollupSession struct {
	Contract     *FullRollup       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FullRollupCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FullRollupCallerSession struct {
	Contract *FullRollupCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// FullRollupTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FullRollupTransactorSession struct {
	Contract     *FullRollupTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// FullRollupRaw is an auto generated low-level Go binding around an Ethereum contract.
type FullRollupRaw struct {
	Contract *FullRollup // Generic contract binding to access the raw methods on
}

// FullRollupCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FullRollupCallerRaw struct {
	Contract *FullRollupCaller // Generic read-only contract binding to access the raw methods on
}

// FullRollupTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FullRollupTransactorRaw struct {
	Contract *FullRollupTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFullRollup creates a new instance of FullRollup, bound to a specific deployed contract.
func NewFullRollup(address common.Address, backend bind.ContractBackend) (*FullRollup, error) {
	contract, err := bindFullRollup(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FullRollup{FullRollupCaller: FullRollupCaller{contract: contract}, FullRollupTransactor: FullRollupTransactor{contract: contract}, FullRollupFilterer: FullRollupFilterer{contract: contract}}, nil
}

// NewFullRollupCaller creates a new read-only instance of FullRollup, bound to a specific deployed contract.
func NewFullRollupCaller(address common.Address, caller bind.ContractCaller) (*FullRollupCaller, error) {
	contract, err := bindFullRollup(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FullRollupCaller{contract: contract}, nil
}

// NewFullRollupTransactor creates a new write-only instance of FullRollup, bound to a specific deployed contract.
func NewFullRollupTransactor(address common.Address, transactor bind.ContractTransactor) (*FullRollupTransactor, error) {
	contract, err := bindFullRollup(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FullRollupTransactor{contract: contract}, nil
}

// NewFullRollupFilterer creates a new log filterer instance of FullRollup, bound to a specific deployed contract.
func NewFullRollupFilterer(address common.Address, filterer bind.ContractFilterer) (*FullRollupFilterer, error) {
	contract, err := bindFullRollup(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FullRollupFilterer{contract: contract}, nil
}

// bindFullRollup binds a generic wrapper to an already deployed contract.
func bindFullRollup(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FullRollupMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FullRollup *FullRollupRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FullRollup.Contract.FullRollupCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FullRollup *FullRollupRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FullRollup.Contract.FullRollupTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FullRollup *FullRollupRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FullRollup.Contract.FullRollupTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FullRollup *FullRollupCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FullRollup.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FullRollup *FullRollupTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FullRollup.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FullRollup *FullRollupTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FullRollup.Contract.contract.Transact(opts, method, params...)
}

// BlockLen is a free data retrieval call binding the contract method 0x94277f84.
//
// Solidity: function BlockLen() view returns(uint256)
func (_FullRollup *FullRollupCaller) BlockLen(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FullRollup.contract.Call(opts, &out, "BlockLen")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BlockLen is a free data retrieval call binding the contract method 0x94277f84.
//
// Solidity: function BlockLen() view returns(uint256)
func (_FullRollup *FullRollupSession) BlockLen() (*big.Int, error) {
	return _FullRollup.Contract.BlockLen(&_FullRollup.CallOpts)
}

// BlockLen is a free data retrieval call binding the contract method 0x94277f84.
//
// Solidity: function BlockLen() view returns(uint256)
func (_FullRollup *FullRollupCallerSession) BlockLen() (*big.Int, error) {
	return _FullRollup.Contract.BlockLen(&_FullRollup.CallOpts)
}

// DISPUTETIME is a free data retrieval call binding the contract method 0x380e9c56.
//
// Solidity: function DISPUTE_TIME() view returns(uint256)
func (_FullRollup *FullRollupCaller) DISPUTETIME(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FullRollup.contract.Call(opts, &out, "DISPUTE_TIME")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DISPUTETIME is a free data retrieval call binding the contract method 0x380e9c56.
//
// Solidity: function DISPUTE_TIME() view returns(uint256)
func (_FullRollup *FullRollupSession) DISPUTETIME() (*big.Int, error) {
	return _FullRollup.Contract.DISPUTETIME(&_FullRollup.CallOpts)
}

// DISPUTETIME is a free data retrieval call binding the contract method 0x380e9c56.
//
// Solidity: function DISPUTE_TIME() view returns(uint256)
func (_FullRollup *FullRollupCallerSession) DISPUTETIME() (*big.Int, error) {
	return _FullRollup.Contract.DISPUTETIME(&_FullRollup.CallOpts)
}

// GetL2Transitions is a free data retrieval call binding the contract method 0xb386eeb9.
//
// Solidity: function GetL2Transitions((address,address,uint256,uint256,uint256)[] _transactions) pure returns((address,uint256,uint256,uint256)[])
func (_FullRollup *FullRollupCaller) GetL2Transitions(opts *bind.CallOpts, _transactions []DataTypesL2Transaction) ([]DataTypesL2Transition, error) {
	var out []interface{}
	err := _FullRollup.contract.Call(opts, &out, "GetL2Transitions", _transactions)

	if err != nil {
		return *new([]DataTypesL2Transition), err
	}

	out0 := *abi.ConvertType(out[0], new([]DataTypesL2Transition)).(*[]DataTypesL2Transition)

	return out0, err

}

// GetL2Transitions is a free data retrieval call binding the contract method 0xb386eeb9.
//
// Solidity: function GetL2Transitions((address,address,uint256,uint256,uint256)[] _transactions) pure returns((address,uint256,uint256,uint256)[])
func (_FullRollup *FullRollupSession) GetL2Transitions(_transactions []DataTypesL2Transaction) ([]DataTypesL2Transition, error) {
	return _FullRollup.Contract.GetL2Transitions(&_FullRollup.CallOpts, _transactions)
}

// GetL2Transitions is a free data retrieval call binding the contract method 0xb386eeb9.
//
// Solidity: function GetL2Transitions((address,address,uint256,uint256,uint256)[] _transactions) pure returns((address,uint256,uint256,uint256)[])
func (_FullRollup *FullRollupCallerSession) GetL2Transitions(_transactions []DataTypesL2Transaction) ([]DataTypesL2Transition, error) {
	return _FullRollup.Contract.GetL2Transitions(&_FullRollup.CallOpts, _transactions)
}

// GetLocalTransition is a free data retrieval call binding the contract method 0x2036c3e5.
//
// Solidity: function GetLocalTransition((address,address,uint256,uint256,uint256)[] _transactions, (uint256,bytes32,bytes32,bytes32,uint256,uint256) _targetBlock) pure returns((uint256,uint256,address,uint256,uint256,uint256)[])
func (_FullRollup *FullRollupCaller) GetLocalTransition(opts *bind.CallOpts, _transactions []DataTypesL2Transaction, _targetBlock DataTypesBlock) ([]DataTypesLocalTransition, error) {
	var out []interface{}
	err := _FullRollup.contract.Call(opts, &out, "GetLocalTransition", _transactions, _targetBlock)

	if err != nil {
		return *new([]DataTypesLocalTransition), err
	}

	out0 := *abi.ConvertType(out[0], new([]DataTypesLocalTransition)).(*[]DataTypesLocalTransition)

	return out0, err

}

// GetLocalTransition is a free data retrieval call binding the contract method 0x2036c3e5.
//
// Solidity: function GetLocalTransition((address,address,uint256,uint256,uint256)[] _transactions, (uint256,bytes32,bytes32,bytes32,uint256,uint256) _targetBlock) pure returns((uint256,uint256,address,uint256,uint256,uint256)[])
func (_FullRollup *FullRollupSession) GetLocalTransition(_transactions []DataTypesL2Transaction, _targetBlock DataTypesBlock) ([]DataTypesLocalTransition, error) {
	return _FullRollup.Contract.GetLocalTransition(&_FullRollup.CallOpts, _transactions, _targetBlock)
}

// GetLocalTransition is a free data retrieval call binding the contract method 0x2036c3e5.
//
// Solidity: function GetLocalTransition((address,address,uint256,uint256,uint256)[] _transactions, (uint256,bytes32,bytes32,bytes32,uint256,uint256) _targetBlock) pure returns((uint256,uint256,address,uint256,uint256,uint256)[])
func (_FullRollup *FullRollupCallerSession) GetLocalTransition(_transactions []DataTypesL2Transaction, _targetBlock DataTypesBlock) ([]DataTypesLocalTransition, error) {
	return _FullRollup.Contract.GetLocalTransition(&_FullRollup.CallOpts, _transactions, _targetBlock)
}

// VerifyTransactions is a free data retrieval call binding the contract method 0xb2d49ffb.
//
// Solidity: function VerifyTransactions((address,address,uint256,uint256,uint256)[] _transactions, bytes32 _tsRoot) view returns(bool, string)
func (_FullRollup *FullRollupCaller) VerifyTransactions(opts *bind.CallOpts, _transactions []DataTypesL2Transaction, _tsRoot [32]byte) (bool, string, error) {
	var out []interface{}
	err := _FullRollup.contract.Call(opts, &out, "VerifyTransactions", _transactions, _tsRoot)

	if err != nil {
		return *new(bool), *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)

	return out0, out1, err

}

// VerifyTransactions is a free data retrieval call binding the contract method 0xb2d49ffb.
//
// Solidity: function VerifyTransactions((address,address,uint256,uint256,uint256)[] _transactions, bytes32 _tsRoot) view returns(bool, string)
func (_FullRollup *FullRollupSession) VerifyTransactions(_transactions []DataTypesL2Transaction, _tsRoot [32]byte) (bool, string, error) {
	return _FullRollup.Contract.VerifyTransactions(&_FullRollup.CallOpts, _transactions, _tsRoot)
}

// VerifyTransactions is a free data retrieval call binding the contract method 0xb2d49ffb.
//
// Solidity: function VerifyTransactions((address,address,uint256,uint256,uint256)[] _transactions, bytes32 _tsRoot) view returns(bool, string)
func (_FullRollup *FullRollupCallerSession) VerifyTransactions(_transactions []DataTypesL2Transaction, _tsRoot [32]byte) (bool, string, error) {
	return _FullRollup.Contract.VerifyTransactions(&_FullRollup.CallOpts, _transactions, _tsRoot)
}

// WITHDRAWWAITPERIOD is a free data retrieval call binding the contract method 0x9381c226.
//
// Solidity: function WITHDRAW_WAIT_PERIOD() view returns(uint256)
func (_FullRollup *FullRollupCaller) WITHDRAWWAITPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FullRollup.contract.Call(opts, &out, "WITHDRAW_WAIT_PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WITHDRAWWAITPERIOD is a free data retrieval call binding the contract method 0x9381c226.
//
// Solidity: function WITHDRAW_WAIT_PERIOD() view returns(uint256)
func (_FullRollup *FullRollupSession) WITHDRAWWAITPERIOD() (*big.Int, error) {
	return _FullRollup.Contract.WITHDRAWWAITPERIOD(&_FullRollup.CallOpts)
}

// WITHDRAWWAITPERIOD is a free data retrieval call binding the contract method 0x9381c226.
//
// Solidity: function WITHDRAW_WAIT_PERIOD() view returns(uint256)
func (_FullRollup *FullRollupCallerSession) WITHDRAWWAITPERIOD() (*big.Int, error) {
	return _FullRollup.Contract.WITHDRAWWAITPERIOD(&_FullRollup.CallOpts)
}

// ZEROBYTES32 is a free data retrieval call binding the contract method 0x069321b0.
//
// Solidity: function ZERO_BYTES32() view returns(bytes32)
func (_FullRollup *FullRollupCaller) ZEROBYTES32(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FullRollup.contract.Call(opts, &out, "ZERO_BYTES32")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ZEROBYTES32 is a free data retrieval call binding the contract method 0x069321b0.
//
// Solidity: function ZERO_BYTES32() view returns(bytes32)
func (_FullRollup *FullRollupSession) ZEROBYTES32() ([32]byte, error) {
	return _FullRollup.Contract.ZEROBYTES32(&_FullRollup.CallOpts)
}

// ZEROBYTES32 is a free data retrieval call binding the contract method 0x069321b0.
//
// Solidity: function ZERO_BYTES32() view returns(bytes32)
func (_FullRollup *FullRollupCallerSession) ZEROBYTES32() ([32]byte, error) {
	return _FullRollup.Contract.ZEROBYTES32(&_FullRollup.CallOpts)
}

// ChainID is a free data retrieval call binding the contract method 0xadc879e9.
//
// Solidity: function chainID() view returns(uint256)
func (_FullRollup *FullRollupCaller) ChainID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FullRollup.contract.Call(opts, &out, "chainID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChainID is a free data retrieval call binding the contract method 0xadc879e9.
//
// Solidity: function chainID() view returns(uint256)
func (_FullRollup *FullRollupSession) ChainID() (*big.Int, error) {
	return _FullRollup.Contract.ChainID(&_FullRollup.CallOpts)
}

// ChainID is a free data retrieval call binding the contract method 0xadc879e9.
//
// Solidity: function chainID() view returns(uint256)
func (_FullRollup *FullRollupCallerSession) ChainID() (*big.Int, error) {
	return _FullRollup.Contract.ChainID(&_FullRollup.CallOpts)
}

// ChallengePool is a free data retrieval call binding the contract method 0xeae5561a.
//
// Solidity: function challengePool(uint256 ) view returns(uint256 chainID, uint256 l2BlockID, uint256 l1BlockID, bytes l1StateRoot, bytes l1BlockData, address account)
func (_FullRollup *FullRollupCaller) ChallengePool(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ChainID     *big.Int
	L2BlockID   *big.Int
	L1BlockID   *big.Int
	L1StateRoot []byte
	L1BlockData []byte
	Account     common.Address
}, error) {
	var out []interface{}
	err := _FullRollup.contract.Call(opts, &out, "challengePool", arg0)

	outstruct := new(struct {
		ChainID     *big.Int
		L2BlockID   *big.Int
		L1BlockID   *big.Int
		L1StateRoot []byte
		L1BlockData []byte
		Account     common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ChainID = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.L2BlockID = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.L1BlockID = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.L1StateRoot = *abi.ConvertType(out[3], new([]byte)).(*[]byte)
	outstruct.L1BlockData = *abi.ConvertType(out[4], new([]byte)).(*[]byte)
	outstruct.Account = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// ChallengePool is a free data retrieval call binding the contract method 0xeae5561a.
//
// Solidity: function challengePool(uint256 ) view returns(uint256 chainID, uint256 l2BlockID, uint256 l1BlockID, bytes l1StateRoot, bytes l1BlockData, address account)
func (_FullRollup *FullRollupSession) ChallengePool(arg0 *big.Int) (struct {
	ChainID     *big.Int
	L2BlockID   *big.Int
	L1BlockID   *big.Int
	L1StateRoot []byte
	L1BlockData []byte
	Account     common.Address
}, error) {
	return _FullRollup.Contract.ChallengePool(&_FullRollup.CallOpts, arg0)
}

// ChallengePool is a free data retrieval call binding the contract method 0xeae5561a.
//
// Solidity: function challengePool(uint256 ) view returns(uint256 chainID, uint256 l2BlockID, uint256 l1BlockID, bytes l1StateRoot, bytes l1BlockData, address account)
func (_FullRollup *FullRollupCallerSession) ChallengePool(arg0 *big.Int) (struct {
	ChainID     *big.Int
	L2BlockID   *big.Int
	L1BlockID   *big.Int
	L1StateRoot []byte
	L1BlockData []byte
	Account     common.Address
}, error) {
	return _FullRollup.Contract.ChallengePool(&_FullRollup.CallOpts, arg0)
}

// ChallengeSize is a free data retrieval call binding the contract method 0x6316bc17.
//
// Solidity: function challengeSize() view returns(uint256)
func (_FullRollup *FullRollupCaller) ChallengeSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FullRollup.contract.Call(opts, &out, "challengeSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChallengeSize is a free data retrieval call binding the contract method 0x6316bc17.
//
// Solidity: function challengeSize() view returns(uint256)
func (_FullRollup *FullRollupSession) ChallengeSize() (*big.Int, error) {
	return _FullRollup.Contract.ChallengeSize(&_FullRollup.CallOpts)
}

// ChallengeSize is a free data retrieval call binding the contract method 0x6316bc17.
//
// Solidity: function challengeSize() view returns(uint256)
func (_FullRollup *FullRollupCallerSession) ChallengeSize() (*big.Int, error) {
	return _FullRollup.Contract.ChallengeSize(&_FullRollup.CallOpts)
}

// CommitterAddress is a free data retrieval call binding the contract method 0x2b19749e.
//
// Solidity: function committerAddress() view returns(address)
func (_FullRollup *FullRollupCaller) CommitterAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FullRollup.contract.Call(opts, &out, "committerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CommitterAddress is a free data retrieval call binding the contract method 0x2b19749e.
//
// Solidity: function committerAddress() view returns(address)
func (_FullRollup *FullRollupSession) CommitterAddress() (common.Address, error) {
	return _FullRollup.Contract.CommitterAddress(&_FullRollup.CallOpts)
}

// CommitterAddress is a free data retrieval call binding the contract method 0x2b19749e.
//
// Solidity: function committerAddress() view returns(address)
func (_FullRollup *FullRollupCallerSession) CommitterAddress() (common.Address, error) {
	return _FullRollup.Contract.CommitterAddress(&_FullRollup.CallOpts)
}

// ConfirmedNum is a free data retrieval call binding the contract method 0x88d1e611.
//
// Solidity: function confirmedNum() view returns(uint256)
func (_FullRollup *FullRollupCaller) ConfirmedNum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FullRollup.contract.Call(opts, &out, "confirmedNum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConfirmedNum is a free data retrieval call binding the contract method 0x88d1e611.
//
// Solidity: function confirmedNum() view returns(uint256)
func (_FullRollup *FullRollupSession) ConfirmedNum() (*big.Int, error) {
	return _FullRollup.Contract.ConfirmedNum(&_FullRollup.CallOpts)
}

// ConfirmedNum is a free data retrieval call binding the contract method 0x88d1e611.
//
// Solidity: function confirmedNum() view returns(uint256)
func (_FullRollup *FullRollupCallerSession) ConfirmedNum() (*big.Int, error) {
	return _FullRollup.Contract.ConfirmedNum(&_FullRollup.CallOpts)
}

// GetCurrentBlockNumber is a free data retrieval call binding the contract method 0x6fd902e1.
//
// Solidity: function getCurrentBlockNumber() view returns(uint256)
func (_FullRollup *FullRollupCaller) GetCurrentBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FullRollup.contract.Call(opts, &out, "getCurrentBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentBlockNumber is a free data retrieval call binding the contract method 0x6fd902e1.
//
// Solidity: function getCurrentBlockNumber() view returns(uint256)
func (_FullRollup *FullRollupSession) GetCurrentBlockNumber() (*big.Int, error) {
	return _FullRollup.Contract.GetCurrentBlockNumber(&_FullRollup.CallOpts)
}

// GetCurrentBlockNumber is a free data retrieval call binding the contract method 0x6fd902e1.
//
// Solidity: function getCurrentBlockNumber() view returns(uint256)
func (_FullRollup *FullRollupCallerSession) GetCurrentBlockNumber() (*big.Int, error) {
	return _FullRollup.Contract.GetCurrentBlockNumber(&_FullRollup.CallOpts)
}

// GetLocalState is a free data retrieval call binding the contract method 0xac288b3f.
//
// Solidity: function getLocalState(address _account) view returns((address,uint256,uint256))
func (_FullRollup *FullRollupCaller) GetLocalState(opts *bind.CallOpts, _account common.Address) (DataTypesLocalState, error) {
	var out []interface{}
	err := _FullRollup.contract.Call(opts, &out, "getLocalState", _account)

	if err != nil {
		return *new(DataTypesLocalState), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesLocalState)).(*DataTypesLocalState)

	return out0, err

}

// GetLocalState is a free data retrieval call binding the contract method 0xac288b3f.
//
// Solidity: function getLocalState(address _account) view returns((address,uint256,uint256))
func (_FullRollup *FullRollupSession) GetLocalState(_account common.Address) (DataTypesLocalState, error) {
	return _FullRollup.Contract.GetLocalState(&_FullRollup.CallOpts, _account)
}

// GetLocalState is a free data retrieval call binding the contract method 0xac288b3f.
//
// Solidity: function getLocalState(address _account) view returns((address,uint256,uint256))
func (_FullRollup *FullRollupCallerSession) GetLocalState(_account common.Address) (DataTypesLocalState, error) {
	return _FullRollup.Contract.GetLocalState(&_FullRollup.CallOpts, _account)
}

// LocalStatePool is a free data retrieval call binding the contract method 0x60050a6b.
//
// Solidity: function localStatePool(address ) view returns(address account, uint256 value, uint256 lock)
func (_FullRollup *FullRollupCaller) LocalStatePool(opts *bind.CallOpts, arg0 common.Address) (struct {
	Account common.Address
	Value   *big.Int
	Lock    *big.Int
}, error) {
	var out []interface{}
	err := _FullRollup.contract.Call(opts, &out, "localStatePool", arg0)

	outstruct := new(struct {
		Account common.Address
		Value   *big.Int
		Lock    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Account = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Value = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Lock = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// LocalStatePool is a free data retrieval call binding the contract method 0x60050a6b.
//
// Solidity: function localStatePool(address ) view returns(address account, uint256 value, uint256 lock)
func (_FullRollup *FullRollupSession) LocalStatePool(arg0 common.Address) (struct {
	Account common.Address
	Value   *big.Int
	Lock    *big.Int
}, error) {
	return _FullRollup.Contract.LocalStatePool(&_FullRollup.CallOpts, arg0)
}

// LocalStatePool is a free data retrieval call binding the contract method 0x60050a6b.
//
// Solidity: function localStatePool(address ) view returns(address account, uint256 value, uint256 lock)
func (_FullRollup *FullRollupCallerSession) LocalStatePool(arg0 common.Address) (struct {
	Account common.Address
	Value   *big.Int
	Lock    *big.Int
}, error) {
	return _FullRollup.Contract.LocalStatePool(&_FullRollup.CallOpts, arg0)
}

// LockTransitionNum is a free data retrieval call binding the contract method 0xc2c36156.
//
// Solidity: function lockTransitionNum() view returns(uint256)
func (_FullRollup *FullRollupCaller) LockTransitionNum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FullRollup.contract.Call(opts, &out, "lockTransitionNum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockTransitionNum is a free data retrieval call binding the contract method 0xc2c36156.
//
// Solidity: function lockTransitionNum() view returns(uint256)
func (_FullRollup *FullRollupSession) LockTransitionNum() (*big.Int, error) {
	return _FullRollup.Contract.LockTransitionNum(&_FullRollup.CallOpts)
}

// LockTransitionNum is a free data retrieval call binding the contract method 0xc2c36156.
//
// Solidity: function lockTransitionNum() view returns(uint256)
func (_FullRollup *FullRollupCallerSession) LockTransitionNum() (*big.Int, error) {
	return _FullRollup.Contract.LockTransitionNum(&_FullRollup.CallOpts)
}

// LockTransitionPool is a free data retrieval call binding the contract method 0x1f754bf7.
//
// Solidity: function lockTransitionPool(uint256 ) view returns(uint256 chainID, uint256 blockNum, address account, uint256 value, uint256 style, uint256 unlockTime)
func (_FullRollup *FullRollupCaller) LockTransitionPool(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ChainID    *big.Int
	BlockNum   *big.Int
	Account    common.Address
	Value      *big.Int
	Style      *big.Int
	UnlockTime *big.Int
}, error) {
	var out []interface{}
	err := _FullRollup.contract.Call(opts, &out, "lockTransitionPool", arg0)

	outstruct := new(struct {
		ChainID    *big.Int
		BlockNum   *big.Int
		Account    common.Address
		Value      *big.Int
		Style      *big.Int
		UnlockTime *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ChainID = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BlockNum = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Account = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Value = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Style = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.UnlockTime = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// LockTransitionPool is a free data retrieval call binding the contract method 0x1f754bf7.
//
// Solidity: function lockTransitionPool(uint256 ) view returns(uint256 chainID, uint256 blockNum, address account, uint256 value, uint256 style, uint256 unlockTime)
func (_FullRollup *FullRollupSession) LockTransitionPool(arg0 *big.Int) (struct {
	ChainID    *big.Int
	BlockNum   *big.Int
	Account    common.Address
	Value      *big.Int
	Style      *big.Int
	UnlockTime *big.Int
}, error) {
	return _FullRollup.Contract.LockTransitionPool(&_FullRollup.CallOpts, arg0)
}

// LockTransitionPool is a free data retrieval call binding the contract method 0x1f754bf7.
//
// Solidity: function lockTransitionPool(uint256 ) view returns(uint256 chainID, uint256 blockNum, address account, uint256 value, uint256 style, uint256 unlockTime)
func (_FullRollup *FullRollupCallerSession) LockTransitionPool(arg0 *big.Int) (struct {
	ChainID    *big.Int
	BlockNum   *big.Int
	Account    common.Address
	Value      *big.Int
	Style      *big.Int
	UnlockTime *big.Int
}, error) {
	return _FullRollup.Contract.LockTransitionPool(&_FullRollup.CallOpts, arg0)
}

// SuperAddress is a free data retrieval call binding the contract method 0x45d3b8db.
//
// Solidity: function superAddress() view returns(address)
func (_FullRollup *FullRollupCaller) SuperAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FullRollup.contract.Call(opts, &out, "superAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SuperAddress is a free data retrieval call binding the contract method 0x45d3b8db.
//
// Solidity: function superAddress() view returns(address)
func (_FullRollup *FullRollupSession) SuperAddress() (common.Address, error) {
	return _FullRollup.Contract.SuperAddress(&_FullRollup.CallOpts)
}

// SuperAddress is a free data retrieval call binding the contract method 0x45d3b8db.
//
// Solidity: function superAddress() view returns(address)
func (_FullRollup *FullRollupCallerSession) SuperAddress() (common.Address, error) {
	return _FullRollup.Contract.SuperAddress(&_FullRollup.CallOpts)
}

// VerifyLocalState is a free data retrieval call binding the contract method 0xcee8e99c.
//
// Solidity: function verifyLocalState((address,uint256,uint256,uint256)[] _transits) view returns(bool)
func (_FullRollup *FullRollupCaller) VerifyLocalState(opts *bind.CallOpts, _transits []DataTypesL2Transition) (bool, error) {
	var out []interface{}
	err := _FullRollup.contract.Call(opts, &out, "verifyLocalState", _transits)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyLocalState is a free data retrieval call binding the contract method 0xcee8e99c.
//
// Solidity: function verifyLocalState((address,uint256,uint256,uint256)[] _transits) view returns(bool)
func (_FullRollup *FullRollupSession) VerifyLocalState(_transits []DataTypesL2Transition) (bool, error) {
	return _FullRollup.Contract.VerifyLocalState(&_FullRollup.CallOpts, _transits)
}

// VerifyLocalState is a free data retrieval call binding the contract method 0xcee8e99c.
//
// Solidity: function verifyLocalState((address,uint256,uint256,uint256)[] _transits) view returns(bool)
func (_FullRollup *FullRollupCallerSession) VerifyLocalState(_transits []DataTypesL2Transition) (bool, error) {
	return _FullRollup.Contract.VerifyLocalState(&_FullRollup.CallOpts, _transits)
}

// ExecutionChallenge is a paid mutator transaction binding the contract method 0x27b7ec20.
//
// Solidity: function ExecutionChallenge(uint256 _targetBlockHeight, (address,address,uint256,uint256,uint256)[] _transactions) returns()
func (_FullRollup *FullRollupTransactor) ExecutionChallenge(opts *bind.TransactOpts, _targetBlockHeight *big.Int, _transactions []DataTypesL2Transaction) (*types.Transaction, error) {
	return _FullRollup.contract.Transact(opts, "ExecutionChallenge", _targetBlockHeight, _transactions)
}

// ExecutionChallenge is a paid mutator transaction binding the contract method 0x27b7ec20.
//
// Solidity: function ExecutionChallenge(uint256 _targetBlockHeight, (address,address,uint256,uint256,uint256)[] _transactions) returns()
func (_FullRollup *FullRollupSession) ExecutionChallenge(_targetBlockHeight *big.Int, _transactions []DataTypesL2Transaction) (*types.Transaction, error) {
	return _FullRollup.Contract.ExecutionChallenge(&_FullRollup.TransactOpts, _targetBlockHeight, _transactions)
}

// ExecutionChallenge is a paid mutator transaction binding the contract method 0x27b7ec20.
//
// Solidity: function ExecutionChallenge(uint256 _targetBlockHeight, (address,address,uint256,uint256,uint256)[] _transactions) returns()
func (_FullRollup *FullRollupTransactorSession) ExecutionChallenge(_targetBlockHeight *big.Int, _transactions []DataTypesL2Transaction) (*types.Transaction, error) {
	return _FullRollup.Contract.ExecutionChallenge(&_FullRollup.TransactOpts, _targetBlockHeight, _transactions)
}

// CheckCommit is a paid mutator transaction binding the contract method 0x551ee37e.
//
// Solidity: function checkCommit() returns()
func (_FullRollup *FullRollupTransactor) CheckCommit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FullRollup.contract.Transact(opts, "checkCommit")
}

// CheckCommit is a paid mutator transaction binding the contract method 0x551ee37e.
//
// Solidity: function checkCommit() returns()
func (_FullRollup *FullRollupSession) CheckCommit() (*types.Transaction, error) {
	return _FullRollup.Contract.CheckCommit(&_FullRollup.TransactOpts)
}

// CheckCommit is a paid mutator transaction binding the contract method 0x551ee37e.
//
// Solidity: function checkCommit() returns()
func (_FullRollup *FullRollupTransactorSession) CheckCommit() (*types.Transaction, error) {
	return _FullRollup.Contract.CheckCommit(&_FullRollup.TransactOpts)
}

// CommitBlock is a paid mutator transaction binding the contract method 0xb1381751.
//
// Solidity: function commitBlock(uint256 _blockNumber, bytes32 _txRoot, bytes32 _localTransitionRoot, bytes _transactions, (address,uint256,uint256,uint256)[] _transitionList) returns(bool)
func (_FullRollup *FullRollupTransactor) CommitBlock(opts *bind.TransactOpts, _blockNumber *big.Int, _txRoot [32]byte, _localTransitionRoot [32]byte, _transactions []byte, _transitionList []DataTypesL2Transition) (*types.Transaction, error) {
	return _FullRollup.contract.Transact(opts, "commitBlock", _blockNumber, _txRoot, _localTransitionRoot, _transactions, _transitionList)
}

// CommitBlock is a paid mutator transaction binding the contract method 0xb1381751.
//
// Solidity: function commitBlock(uint256 _blockNumber, bytes32 _txRoot, bytes32 _localTransitionRoot, bytes _transactions, (address,uint256,uint256,uint256)[] _transitionList) returns(bool)
func (_FullRollup *FullRollupSession) CommitBlock(_blockNumber *big.Int, _txRoot [32]byte, _localTransitionRoot [32]byte, _transactions []byte, _transitionList []DataTypesL2Transition) (*types.Transaction, error) {
	return _FullRollup.Contract.CommitBlock(&_FullRollup.TransactOpts, _blockNumber, _txRoot, _localTransitionRoot, _transactions, _transitionList)
}

// CommitBlock is a paid mutator transaction binding the contract method 0xb1381751.
//
// Solidity: function commitBlock(uint256 _blockNumber, bytes32 _txRoot, bytes32 _localTransitionRoot, bytes _transactions, (address,uint256,uint256,uint256)[] _transitionList) returns(bool)
func (_FullRollup *FullRollupTransactorSession) CommitBlock(_blockNumber *big.Int, _txRoot [32]byte, _localTransitionRoot [32]byte, _transactions []byte, _transitionList []DataTypesL2Transition) (*types.Transaction, error) {
	return _FullRollup.Contract.CommitBlock(&_FullRollup.TransactOpts, _blockNumber, _txRoot, _localTransitionRoot, _transactions, _transitionList)
}

// ConfirmLocalState is a paid mutator transaction binding the contract method 0xd6ce438f.
//
// Solidity: function confirmLocalState(uint256 _blockNum, bytes transitionDatas) returns(uint256)
func (_FullRollup *FullRollupTransactor) ConfirmLocalState(opts *bind.TransactOpts, _blockNum *big.Int, transitionDatas []byte) (*types.Transaction, error) {
	return _FullRollup.contract.Transact(opts, "confirmLocalState", _blockNum, transitionDatas)
}

// ConfirmLocalState is a paid mutator transaction binding the contract method 0xd6ce438f.
//
// Solidity: function confirmLocalState(uint256 _blockNum, bytes transitionDatas) returns(uint256)
func (_FullRollup *FullRollupSession) ConfirmLocalState(_blockNum *big.Int, transitionDatas []byte) (*types.Transaction, error) {
	return _FullRollup.Contract.ConfirmLocalState(&_FullRollup.TransactOpts, _blockNum, transitionDatas)
}

// ConfirmLocalState is a paid mutator transaction binding the contract method 0xd6ce438f.
//
// Solidity: function confirmLocalState(uint256 _blockNum, bytes transitionDatas) returns(uint256)
func (_FullRollup *FullRollupTransactorSession) ConfirmLocalState(_blockNum *big.Int, transitionDatas []byte) (*types.Transaction, error) {
	return _FullRollup.Contract.ConfirmLocalState(&_FullRollup.TransactOpts, _blockNum, transitionDatas)
}

// DepositCoin is a paid mutator transaction binding the contract method 0xd8c67f01.
//
// Solidity: function depositCoin(address _account, uint256 _value) returns()
func (_FullRollup *FullRollupTransactor) DepositCoin(opts *bind.TransactOpts, _account common.Address, _value *big.Int) (*types.Transaction, error) {
	return _FullRollup.contract.Transact(opts, "depositCoin", _account, _value)
}

// DepositCoin is a paid mutator transaction binding the contract method 0xd8c67f01.
//
// Solidity: function depositCoin(address _account, uint256 _value) returns()
func (_FullRollup *FullRollupSession) DepositCoin(_account common.Address, _value *big.Int) (*types.Transaction, error) {
	return _FullRollup.Contract.DepositCoin(&_FullRollup.TransactOpts, _account, _value)
}

// DepositCoin is a paid mutator transaction binding the contract method 0xd8c67f01.
//
// Solidity: function depositCoin(address _account, uint256 _value) returns()
func (_FullRollup *FullRollupTransactorSession) DepositCoin(_account common.Address, _value *big.Int) (*types.Transaction, error) {
	return _FullRollup.Contract.DepositCoin(&_FullRollup.TransactOpts, _account, _value)
}

// RefundCoin is a paid mutator transaction binding the contract method 0x96ddc709.
//
// Solidity: function refundCoin(address _account, uint256 _value) returns()
func (_FullRollup *FullRollupTransactor) RefundCoin(opts *bind.TransactOpts, _account common.Address, _value *big.Int) (*types.Transaction, error) {
	return _FullRollup.contract.Transact(opts, "refundCoin", _account, _value)
}

// RefundCoin is a paid mutator transaction binding the contract method 0x96ddc709.
//
// Solidity: function refundCoin(address _account, uint256 _value) returns()
func (_FullRollup *FullRollupSession) RefundCoin(_account common.Address, _value *big.Int) (*types.Transaction, error) {
	return _FullRollup.Contract.RefundCoin(&_FullRollup.TransactOpts, _account, _value)
}

// RefundCoin is a paid mutator transaction binding the contract method 0x96ddc709.
//
// Solidity: function refundCoin(address _account, uint256 _value) returns()
func (_FullRollup *FullRollupTransactorSession) RefundCoin(_account common.Address, _value *big.Int) (*types.Transaction, error) {
	return _FullRollup.Contract.RefundCoin(&_FullRollup.TransactOpts, _account, _value)
}

// RemoveLocalStateWithBlockNum is a paid mutator transaction binding the contract method 0xf20c6e1e.
//
// Solidity: function removeLocalStateWithBlockNum(uint256 _blockNum, (address,uint256,uint256,uint256)[] _transits) returns(uint256)
func (_FullRollup *FullRollupTransactor) RemoveLocalStateWithBlockNum(opts *bind.TransactOpts, _blockNum *big.Int, _transits []DataTypesL2Transition) (*types.Transaction, error) {
	return _FullRollup.contract.Transact(opts, "removeLocalStateWithBlockNum", _blockNum, _transits)
}

// RemoveLocalStateWithBlockNum is a paid mutator transaction binding the contract method 0xf20c6e1e.
//
// Solidity: function removeLocalStateWithBlockNum(uint256 _blockNum, (address,uint256,uint256,uint256)[] _transits) returns(uint256)
func (_FullRollup *FullRollupSession) RemoveLocalStateWithBlockNum(_blockNum *big.Int, _transits []DataTypesL2Transition) (*types.Transaction, error) {
	return _FullRollup.Contract.RemoveLocalStateWithBlockNum(&_FullRollup.TransactOpts, _blockNum, _transits)
}

// RemoveLocalStateWithBlockNum is a paid mutator transaction binding the contract method 0xf20c6e1e.
//
// Solidity: function removeLocalStateWithBlockNum(uint256 _blockNum, (address,uint256,uint256,uint256)[] _transits) returns(uint256)
func (_FullRollup *FullRollupTransactorSession) RemoveLocalStateWithBlockNum(_blockNum *big.Int, _transits []DataTypesL2Transition) (*types.Transaction, error) {
	return _FullRollup.Contract.RemoveLocalStateWithBlockNum(&_FullRollup.TransactOpts, _blockNum, _transits)
}

// RollBackBlock is a paid mutator transaction binding the contract method 0x3e711294.
//
// Solidity: function rollBackBlock(uint256 _targetBlockHeight) returns()
func (_FullRollup *FullRollupTransactor) RollBackBlock(opts *bind.TransactOpts, _targetBlockHeight *big.Int) (*types.Transaction, error) {
	return _FullRollup.contract.Transact(opts, "rollBackBlock", _targetBlockHeight)
}

// RollBackBlock is a paid mutator transaction binding the contract method 0x3e711294.
//
// Solidity: function rollBackBlock(uint256 _targetBlockHeight) returns()
func (_FullRollup *FullRollupSession) RollBackBlock(_targetBlockHeight *big.Int) (*types.Transaction, error) {
	return _FullRollup.Contract.RollBackBlock(&_FullRollup.TransactOpts, _targetBlockHeight)
}

// RollBackBlock is a paid mutator transaction binding the contract method 0x3e711294.
//
// Solidity: function rollBackBlock(uint256 _targetBlockHeight) returns()
func (_FullRollup *FullRollupTransactorSession) RollBackBlock(_targetBlockHeight *big.Int) (*types.Transaction, error) {
	return _FullRollup.Contract.RollBackBlock(&_FullRollup.TransactOpts, _targetBlockHeight)
}

// SetCommitterAddress is a paid mutator transaction binding the contract method 0x13df8728.
//
// Solidity: function setCommitterAddress(address _committerAddress) returns()
func (_FullRollup *FullRollupTransactor) SetCommitterAddress(opts *bind.TransactOpts, _committerAddress common.Address) (*types.Transaction, error) {
	return _FullRollup.contract.Transact(opts, "setCommitterAddress", _committerAddress)
}

// SetCommitterAddress is a paid mutator transaction binding the contract method 0x13df8728.
//
// Solidity: function setCommitterAddress(address _committerAddress) returns()
func (_FullRollup *FullRollupSession) SetCommitterAddress(_committerAddress common.Address) (*types.Transaction, error) {
	return _FullRollup.Contract.SetCommitterAddress(&_FullRollup.TransactOpts, _committerAddress)
}

// SetCommitterAddress is a paid mutator transaction binding the contract method 0x13df8728.
//
// Solidity: function setCommitterAddress(address _committerAddress) returns()
func (_FullRollup *FullRollupTransactorSession) SetCommitterAddress(_committerAddress common.Address) (*types.Transaction, error) {
	return _FullRollup.Contract.SetCommitterAddress(&_FullRollup.TransactOpts, _committerAddress)
}

// VerifyAndLockLocalState is a paid mutator transaction binding the contract method 0x45046217.
//
// Solidity: function verifyAndLockLocalState((address,uint256,uint256,uint256)[] _transits, uint256 _blockNum, uint256 _lockTime) returns(bool)
func (_FullRollup *FullRollupTransactor) VerifyAndLockLocalState(opts *bind.TransactOpts, _transits []DataTypesL2Transition, _blockNum *big.Int, _lockTime *big.Int) (*types.Transaction, error) {
	return _FullRollup.contract.Transact(opts, "verifyAndLockLocalState", _transits, _blockNum, _lockTime)
}

// VerifyAndLockLocalState is a paid mutator transaction binding the contract method 0x45046217.
//
// Solidity: function verifyAndLockLocalState((address,uint256,uint256,uint256)[] _transits, uint256 _blockNum, uint256 _lockTime) returns(bool)
func (_FullRollup *FullRollupSession) VerifyAndLockLocalState(_transits []DataTypesL2Transition, _blockNum *big.Int, _lockTime *big.Int) (*types.Transaction, error) {
	return _FullRollup.Contract.VerifyAndLockLocalState(&_FullRollup.TransactOpts, _transits, _blockNum, _lockTime)
}

// VerifyAndLockLocalState is a paid mutator transaction binding the contract method 0x45046217.
//
// Solidity: function verifyAndLockLocalState((address,uint256,uint256,uint256)[] _transits, uint256 _blockNum, uint256 _lockTime) returns(bool)
func (_FullRollup *FullRollupTransactorSession) VerifyAndLockLocalState(_transits []DataTypesL2Transition, _blockNum *big.Int, _lockTime *big.Int) (*types.Transaction, error) {
	return _FullRollup.Contract.VerifyAndLockLocalState(&_FullRollup.TransactOpts, _transits, _blockNum, _lockTime)
}

// FullRollupErrorNotifyIterator is returned from FilterErrorNotify and is used to iterate over the raw logs and unpacked data for ErrorNotify events raised by the FullRollup contract.
type FullRollupErrorNotifyIterator struct {
	Event *FullRollupErrorNotify // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FullRollupErrorNotifyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FullRollupErrorNotify)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FullRollupErrorNotify)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FullRollupErrorNotifyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FullRollupErrorNotifyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FullRollupErrorNotify represents a ErrorNotify event raised by the FullRollup contract.
type FullRollupErrorNotify struct {
	State *big.Int
	Error string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterErrorNotify is a free log retrieval operation binding the contract event 0x6307687441ab8283c24801dd603e8b616bd7909ac5cfa1fd14cac706a07d9a99.
//
// Solidity: event ErrorNotify(uint256 state, string error)
func (_FullRollup *FullRollupFilterer) FilterErrorNotify(opts *bind.FilterOpts) (*FullRollupErrorNotifyIterator, error) {

	logs, sub, err := _FullRollup.contract.FilterLogs(opts, "ErrorNotify")
	if err != nil {
		return nil, err
	}
	return &FullRollupErrorNotifyIterator{contract: _FullRollup.contract, event: "ErrorNotify", logs: logs, sub: sub}, nil
}

// WatchErrorNotify is a free log subscription operation binding the contract event 0x6307687441ab8283c24801dd603e8b616bd7909ac5cfa1fd14cac706a07d9a99.
//
// Solidity: event ErrorNotify(uint256 state, string error)
func (_FullRollup *FullRollupFilterer) WatchErrorNotify(opts *bind.WatchOpts, sink chan<- *FullRollupErrorNotify) (event.Subscription, error) {

	logs, sub, err := _FullRollup.contract.WatchLogs(opts, "ErrorNotify")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FullRollupErrorNotify)
				if err := _FullRollup.contract.UnpackLog(event, "ErrorNotify", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseErrorNotify is a log parse operation binding the contract event 0x6307687441ab8283c24801dd603e8b616bd7909ac5cfa1fd14cac706a07d9a99.
//
// Solidity: event ErrorNotify(uint256 state, string error)
func (_FullRollup *FullRollupFilterer) ParseErrorNotify(log types.Log) (*FullRollupErrorNotify, error) {
	event := new(FullRollupErrorNotify)
	if err := _FullRollup.contract.UnpackLog(event, "ErrorNotify", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FullRollupLocalStateNotifyIterator is returned from FilterLocalStateNotify and is used to iterate over the raw logs and unpacked data for LocalStateNotify events raised by the FullRollup contract.
type FullRollupLocalStateNotifyIterator struct {
	Event *FullRollupLocalStateNotify // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FullRollupLocalStateNotifyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FullRollupLocalStateNotify)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FullRollupLocalStateNotify)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FullRollupLocalStateNotifyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FullRollupLocalStateNotifyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FullRollupLocalStateNotify represents a LocalStateNotify event raised by the FullRollup contract.
type FullRollupLocalStateNotify struct {
	Account common.Address
	Style   string
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLocalStateNotify is a free log retrieval operation binding the contract event 0xe1838f5415173c2497fa5559272813a687a79327934a5b294c2d81ed70d76b0a.
//
// Solidity: event LocalStateNotify(address account, string style, uint256 value)
func (_FullRollup *FullRollupFilterer) FilterLocalStateNotify(opts *bind.FilterOpts) (*FullRollupLocalStateNotifyIterator, error) {

	logs, sub, err := _FullRollup.contract.FilterLogs(opts, "LocalStateNotify")
	if err != nil {
		return nil, err
	}
	return &FullRollupLocalStateNotifyIterator{contract: _FullRollup.contract, event: "LocalStateNotify", logs: logs, sub: sub}, nil
}

// WatchLocalStateNotify is a free log subscription operation binding the contract event 0xe1838f5415173c2497fa5559272813a687a79327934a5b294c2d81ed70d76b0a.
//
// Solidity: event LocalStateNotify(address account, string style, uint256 value)
func (_FullRollup *FullRollupFilterer) WatchLocalStateNotify(opts *bind.WatchOpts, sink chan<- *FullRollupLocalStateNotify) (event.Subscription, error) {

	logs, sub, err := _FullRollup.contract.WatchLogs(opts, "LocalStateNotify")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FullRollupLocalStateNotify)
				if err := _FullRollup.contract.UnpackLog(event, "LocalStateNotify", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLocalStateNotify is a log parse operation binding the contract event 0xe1838f5415173c2497fa5559272813a687a79327934a5b294c2d81ed70d76b0a.
//
// Solidity: event LocalStateNotify(address account, string style, uint256 value)
func (_FullRollup *FullRollupFilterer) ParseLocalStateNotify(log types.Log) (*FullRollupLocalStateNotify, error) {
	event := new(FullRollupLocalStateNotify)
	if err := _FullRollup.contract.UnpackLog(event, "LocalStateNotify", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FullRollupRollupBlockNotifyIterator is returned from FilterRollupBlockNotify and is used to iterate over the raw logs and unpacked data for RollupBlockNotify events raised by the FullRollup contract.
type FullRollupRollupBlockNotifyIterator struct {
	Event *FullRollupRollupBlockNotify // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FullRollupRollupBlockNotifyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FullRollupRollupBlockNotify)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FullRollupRollupBlockNotify)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FullRollupRollupBlockNotifyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FullRollupRollupBlockNotifyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FullRollupRollupBlockNotify represents a RollupBlockNotify event raised by the FullRollup contract.
type FullRollupRollupBlockNotify struct {
	BlockNumber *big.Int
	Desc        string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRollupBlockNotify is a free log retrieval operation binding the contract event 0x5e19c79d1d9c171225be02ab9d3a12ed7c13ccf394c6919a8d45a6bd4e66f6b6.
//
// Solidity: event RollupBlockNotify(uint256 blockNumber, string desc)
func (_FullRollup *FullRollupFilterer) FilterRollupBlockNotify(opts *bind.FilterOpts) (*FullRollupRollupBlockNotifyIterator, error) {

	logs, sub, err := _FullRollup.contract.FilterLogs(opts, "RollupBlockNotify")
	if err != nil {
		return nil, err
	}
	return &FullRollupRollupBlockNotifyIterator{contract: _FullRollup.contract, event: "RollupBlockNotify", logs: logs, sub: sub}, nil
}

// WatchRollupBlockNotify is a free log subscription operation binding the contract event 0x5e19c79d1d9c171225be02ab9d3a12ed7c13ccf394c6919a8d45a6bd4e66f6b6.
//
// Solidity: event RollupBlockNotify(uint256 blockNumber, string desc)
func (_FullRollup *FullRollupFilterer) WatchRollupBlockNotify(opts *bind.WatchOpts, sink chan<- *FullRollupRollupBlockNotify) (event.Subscription, error) {

	logs, sub, err := _FullRollup.contract.WatchLogs(opts, "RollupBlockNotify")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FullRollupRollupBlockNotify)
				if err := _FullRollup.contract.UnpackLog(event, "RollupBlockNotify", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRollupBlockNotify is a log parse operation binding the contract event 0x5e19c79d1d9c171225be02ab9d3a12ed7c13ccf394c6919a8d45a6bd4e66f6b6.
//
// Solidity: event RollupBlockNotify(uint256 blockNumber, string desc)
func (_FullRollup *FullRollupFilterer) ParseRollupBlockNotify(log types.Log) (*FullRollupRollupBlockNotify, error) {
	event := new(FullRollupRollupBlockNotify)
	if err := _FullRollup.contract.UnpackLog(event, "RollupBlockNotify", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

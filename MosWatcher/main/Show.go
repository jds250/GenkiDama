// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// ShowMetaData contains all meta data concerning the Show contract.
var ShowMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[],\"name\":\"showStr\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newStr\",\"type\":\"string\"}],\"name\":\"setShowStr\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"show\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ShowABI is the input ABI used to generate the binding from.
// Deprecated: Use ShowMetaData.ABI instead.
var ShowABI = ShowMetaData.ABI

// Show is an auto generated Go binding around an Ethereum contract.
type Show struct {
	ShowCaller     // Read-only binding to the contract
	ShowTransactor // Write-only binding to the contract
	ShowFilterer   // Log filterer for contract events
}

// ShowCaller is an auto generated read-only Go binding around an Ethereum contract.
type ShowCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShowTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ShowTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShowFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ShowFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShowSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ShowSession struct {
	Contract     *Show             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ShowCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ShowCallerSession struct {
	Contract *ShowCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ShowTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ShowTransactorSession struct {
	Contract     *ShowTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ShowRaw is an auto generated low-level Go binding around an Ethereum contract.
type ShowRaw struct {
	Contract *Show // Generic contract binding to access the raw methods on
}

// ShowCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ShowCallerRaw struct {
	Contract *ShowCaller // Generic read-only contract binding to access the raw methods on
}

// ShowTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ShowTransactorRaw struct {
	Contract *ShowTransactor // Generic write-only contract binding to access the raw methods on
}

// NewShow creates a new instance of Show, bound to a specific deployed contract.
func NewShow(address common.Address, backend bind.ContractBackend) (*Show, error) {
	contract, err := bindShow(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Show{ShowCaller: ShowCaller{contract: contract}, ShowTransactor: ShowTransactor{contract: contract}, ShowFilterer: ShowFilterer{contract: contract}}, nil
}

// NewShowCaller creates a new read-only instance of Show, bound to a specific deployed contract.
func NewShowCaller(address common.Address, caller bind.ContractCaller) (*ShowCaller, error) {
	contract, err := bindShow(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ShowCaller{contract: contract}, nil
}

// NewShowTransactor creates a new write-only instance of Show, bound to a specific deployed contract.
func NewShowTransactor(address common.Address, transactor bind.ContractTransactor) (*ShowTransactor, error) {
	contract, err := bindShow(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ShowTransactor{contract: contract}, nil
}

// NewShowFilterer creates a new log filterer instance of Show, bound to a specific deployed contract.
func NewShowFilterer(address common.Address, filterer bind.ContractFilterer) (*ShowFilterer, error) {
	contract, err := bindShow(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ShowFilterer{contract: contract}, nil
}

// bindShow binds a generic wrapper to an already deployed contract.
func bindShow(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ShowMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Show *ShowRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Show.Contract.ShowCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Show *ShowRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Show.Contract.ShowTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Show *ShowRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Show.Contract.ShowTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Show *ShowCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Show.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Show *ShowTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Show.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Show *ShowTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Show.Contract.contract.Transact(opts, method, params...)
}

// Show is a free data retrieval call binding the contract method 0xcc80f6f3.
//
// Solidity: function show() view returns(string)
func (_Show *ShowCaller) Show(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Show.contract.Call(opts, &out, "show")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Show is a free data retrieval call binding the contract method 0xcc80f6f3.
//
// Solidity: function show() view returns(string)
func (_Show *ShowSession) Show() (string, error) {
	return _Show.Contract.Show(&_Show.CallOpts)
}

// Show is a free data retrieval call binding the contract method 0xcc80f6f3.
//
// Solidity: function show() view returns(string)
func (_Show *ShowCallerSession) Show() (string, error) {
	return _Show.Contract.Show(&_Show.CallOpts)
}

// ShowStr is a free data retrieval call binding the contract method 0xcef54f13.
//
// Solidity: function showStr() view returns(string)
func (_Show *ShowCaller) ShowStr(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Show.contract.Call(opts, &out, "showStr")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ShowStr is a free data retrieval call binding the contract method 0xcef54f13.
//
// Solidity: function showStr() view returns(string)
func (_Show *ShowSession) ShowStr() (string, error) {
	return _Show.Contract.ShowStr(&_Show.CallOpts)
}

// ShowStr is a free data retrieval call binding the contract method 0xcef54f13.
//
// Solidity: function showStr() view returns(string)
func (_Show *ShowCallerSession) ShowStr() (string, error) {
	return _Show.Contract.ShowStr(&_Show.CallOpts)
}

// SetShowStr is a paid mutator transaction binding the contract method 0xb10396c5.
//
// Solidity: function setShowStr(string _newStr) returns(bool, string)
func (_Show *ShowTransactor) SetShowStr(opts *bind.TransactOpts, _newStr string) (*types.Transaction, error) {
	return _Show.contract.Transact(opts, "setShowStr", _newStr)
}

// SetShowStr is a paid mutator transaction binding the contract method 0xb10396c5.
//
// Solidity: function setShowStr(string _newStr) returns(bool, string)
func (_Show *ShowSession) SetShowStr(_newStr string) (*types.Transaction, error) {
	return _Show.Contract.SetShowStr(&_Show.TransactOpts, _newStr)
}

// SetShowStr is a paid mutator transaction binding the contract method 0xb10396c5.
//
// Solidity: function setShowStr(string _newStr) returns(bool, string)
func (_Show *ShowTransactorSession) SetShowStr(_newStr string) (*types.Transaction, error) {
	return _Show.Contract.SetShowStr(&_Show.TransactOpts, _newStr)
}

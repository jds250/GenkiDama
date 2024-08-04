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

// StorageMetaData contains all meta data concerning the Storage contract.
var StorageMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[],\"name\":\"value\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"showAddr\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_showAddr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"returnData\",\"type\":\"bytes\"}],\"name\":\"ABIEvent\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"number\",\"type\":\"uint256\"}],\"name\":\"store\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"retrieve\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_str\",\"type\":\"string\"}],\"name\":\"setShowStr\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getShowAddr\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getShowRes\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getInfo\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// StorageABI is the input ABI used to generate the binding from.
// Deprecated: Use StorageMetaData.ABI instead.
var StorageABI = StorageMetaData.ABI

// Storage is an auto generated Go binding around an Ethereum contract.
type Storage struct {
	StorageCaller     // Read-only binding to the contract
	StorageTransactor // Write-only binding to the contract
	StorageFilterer   // Log filterer for contract events
}

// StorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type StorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StorageSession struct {
	Contract     *Storage          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StorageCallerSession struct {
	Contract *StorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StorageTransactorSession struct {
	Contract     *StorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type StorageRaw struct {
	Contract *Storage // Generic contract binding to access the raw methods on
}

// StorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StorageCallerRaw struct {
	Contract *StorageCaller // Generic read-only contract binding to access the raw methods on
}

// StorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StorageTransactorRaw struct {
	Contract *StorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStorage creates a new instance of Storage, bound to a specific deployed contract.
func NewStorage(address common.Address, backend bind.ContractBackend) (*Storage, error) {
	contract, err := bindStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Storage{StorageCaller: StorageCaller{contract: contract}, StorageTransactor: StorageTransactor{contract: contract}, StorageFilterer: StorageFilterer{contract: contract}}, nil
}

// NewStorageCaller creates a new read-only instance of Storage, bound to a specific deployed contract.
func NewStorageCaller(address common.Address, caller bind.ContractCaller) (*StorageCaller, error) {
	contract, err := bindStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StorageCaller{contract: contract}, nil
}

// NewStorageTransactor creates a new write-only instance of Storage, bound to a specific deployed contract.
func NewStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*StorageTransactor, error) {
	contract, err := bindStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StorageTransactor{contract: contract}, nil
}

// NewStorageFilterer creates a new log filterer instance of Storage, bound to a specific deployed contract.
func NewStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*StorageFilterer, error) {
	contract, err := bindStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StorageFilterer{contract: contract}, nil
}

// bindStorage binds a generic wrapper to an already deployed contract.
func bindStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Storage *StorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Storage.Contract.StorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Storage *StorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.Contract.StorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Storage *StorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Storage.Contract.StorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Storage *StorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Storage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Storage *StorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Storage *StorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Storage.Contract.contract.Transact(opts, method, params...)
}

// GetInfo is a free data retrieval call binding the contract method 0x5a9b0b89.
//
// Solidity: function getInfo() view returns(bytes)
func (_Storage *StorageCaller) GetInfo(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getInfo")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetInfo is a free data retrieval call binding the contract method 0x5a9b0b89.
//
// Solidity: function getInfo() view returns(bytes)
func (_Storage *StorageSession) GetInfo() ([]byte, error) {
	return _Storage.Contract.GetInfo(&_Storage.CallOpts)
}

// GetInfo is a free data retrieval call binding the contract method 0x5a9b0b89.
//
// Solidity: function getInfo() view returns(bytes)
func (_Storage *StorageCallerSession) GetInfo() ([]byte, error) {
	return _Storage.Contract.GetInfo(&_Storage.CallOpts)
}

// GetShowAddr is a free data retrieval call binding the contract method 0xb6397d96.
//
// Solidity: function getShowAddr() view returns(address)
func (_Storage *StorageCaller) GetShowAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getShowAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetShowAddr is a free data retrieval call binding the contract method 0xb6397d96.
//
// Solidity: function getShowAddr() view returns(address)
func (_Storage *StorageSession) GetShowAddr() (common.Address, error) {
	return _Storage.Contract.GetShowAddr(&_Storage.CallOpts)
}

// GetShowAddr is a free data retrieval call binding the contract method 0xb6397d96.
//
// Solidity: function getShowAddr() view returns(address)
func (_Storage *StorageCallerSession) GetShowAddr() (common.Address, error) {
	return _Storage.Contract.GetShowAddr(&_Storage.CallOpts)
}

// GetShowRes is a free data retrieval call binding the contract method 0xa1f64064.
//
// Solidity: function getShowRes() view returns(string)
func (_Storage *StorageCaller) GetShowRes(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getShowRes")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetShowRes is a free data retrieval call binding the contract method 0xa1f64064.
//
// Solidity: function getShowRes() view returns(string)
func (_Storage *StorageSession) GetShowRes() (string, error) {
	return _Storage.Contract.GetShowRes(&_Storage.CallOpts)
}

// GetShowRes is a free data retrieval call binding the contract method 0xa1f64064.
//
// Solidity: function getShowRes() view returns(string)
func (_Storage *StorageCallerSession) GetShowRes() (string, error) {
	return _Storage.Contract.GetShowRes(&_Storage.CallOpts)
}

// Retrieve is a free data retrieval call binding the contract method 0x2e64cec1.
//
// Solidity: function retrieve() view returns(uint256)
func (_Storage *StorageCaller) Retrieve(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "retrieve")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Retrieve is a free data retrieval call binding the contract method 0x2e64cec1.
//
// Solidity: function retrieve() view returns(uint256)
func (_Storage *StorageSession) Retrieve() (*big.Int, error) {
	return _Storage.Contract.Retrieve(&_Storage.CallOpts)
}

// Retrieve is a free data retrieval call binding the contract method 0x2e64cec1.
//
// Solidity: function retrieve() view returns(uint256)
func (_Storage *StorageCallerSession) Retrieve() (*big.Int, error) {
	return _Storage.Contract.Retrieve(&_Storage.CallOpts)
}

// ShowAddr is a free data retrieval call binding the contract method 0xa422bb1a.
//
// Solidity: function showAddr() view returns(address)
func (_Storage *StorageCaller) ShowAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "showAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ShowAddr is a free data retrieval call binding the contract method 0xa422bb1a.
//
// Solidity: function showAddr() view returns(address)
func (_Storage *StorageSession) ShowAddr() (common.Address, error) {
	return _Storage.Contract.ShowAddr(&_Storage.CallOpts)
}

// ShowAddr is a free data retrieval call binding the contract method 0xa422bb1a.
//
// Solidity: function showAddr() view returns(address)
func (_Storage *StorageCallerSession) ShowAddr() (common.Address, error) {
	return _Storage.Contract.ShowAddr(&_Storage.CallOpts)
}

// Value is a free data retrieval call binding the contract method 0x3fa4f245.
//
// Solidity: function value() view returns(uint256)
func (_Storage *StorageCaller) Value(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "value")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Value is a free data retrieval call binding the contract method 0x3fa4f245.
//
// Solidity: function value() view returns(uint256)
func (_Storage *StorageSession) Value() (*big.Int, error) {
	return _Storage.Contract.Value(&_Storage.CallOpts)
}

// Value is a free data retrieval call binding the contract method 0x3fa4f245.
//
// Solidity: function value() view returns(uint256)
func (_Storage *StorageCallerSession) Value() (*big.Int, error) {
	return _Storage.Contract.Value(&_Storage.CallOpts)
}

// SetShowStr is a paid mutator transaction binding the contract method 0xb10396c5.
//
// Solidity: function setShowStr(string _str) returns(bool, string)
func (_Storage *StorageTransactor) SetShowStr(opts *bind.TransactOpts, _str string) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "setShowStr", _str)
}

// SetShowStr is a paid mutator transaction binding the contract method 0xb10396c5.
//
// Solidity: function setShowStr(string _str) returns(bool, string)
func (_Storage *StorageSession) SetShowStr(_str string) (*types.Transaction, error) {
	return _Storage.Contract.SetShowStr(&_Storage.TransactOpts, _str)
}

// SetShowStr is a paid mutator transaction binding the contract method 0xb10396c5.
//
// Solidity: function setShowStr(string _str) returns(bool, string)
func (_Storage *StorageTransactorSession) SetShowStr(_str string) (*types.Transaction, error) {
	return _Storage.Contract.SetShowStr(&_Storage.TransactOpts, _str)
}

// Store is a paid mutator transaction binding the contract method 0x6057361d.
//
// Solidity: function store(uint256 number) returns()
func (_Storage *StorageTransactor) Store(opts *bind.TransactOpts, number *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "store", number)
}

// Store is a paid mutator transaction binding the contract method 0x6057361d.
//
// Solidity: function store(uint256 number) returns()
func (_Storage *StorageSession) Store(number *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.Store(&_Storage.TransactOpts, number)
}

// Store is a paid mutator transaction binding the contract method 0x6057361d.
//
// Solidity: function store(uint256 number) returns()
func (_Storage *StorageTransactorSession) Store(number *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.Store(&_Storage.TransactOpts, number)
}

// StorageABIEventIterator is returned from FilterABIEvent and is used to iterate over the raw logs and unpacked data for ABIEvent events raised by the Storage contract.
type StorageABIEventIterator struct {
	Event *StorageABIEvent // Event containing the contract specifics and raw log

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
func (it *StorageABIEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageABIEvent)
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
		it.Event = new(StorageABIEvent)
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
func (it *StorageABIEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageABIEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageABIEvent represents a ABIEvent event raised by the Storage contract.
type StorageABIEvent struct {
	Success    bool
	ReturnData []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterABIEvent is a free log retrieval operation binding the contract event 0x0460adb39e28c2e6aa9aa61a439e432525b0d3df20cef597c3697d5e498fb7ff.
//
// Solidity: event ABIEvent(bool success, bytes returnData)
func (_Storage *StorageFilterer) FilterABIEvent(opts *bind.FilterOpts) (*StorageABIEventIterator, error) {

	logs, sub, err := _Storage.contract.FilterLogs(opts, "ABIEvent")
	if err != nil {
		return nil, err
	}
	return &StorageABIEventIterator{contract: _Storage.contract, event: "ABIEvent", logs: logs, sub: sub}, nil
}

// WatchABIEvent is a free log subscription operation binding the contract event 0x0460adb39e28c2e6aa9aa61a439e432525b0d3df20cef597c3697d5e498fb7ff.
//
// Solidity: event ABIEvent(bool success, bytes returnData)
func (_Storage *StorageFilterer) WatchABIEvent(opts *bind.WatchOpts, sink chan<- *StorageABIEvent) (event.Subscription, error) {

	logs, sub, err := _Storage.contract.WatchLogs(opts, "ABIEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageABIEvent)
				if err := _Storage.contract.UnpackLog(event, "ABIEvent", log); err != nil {
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

// ParseABIEvent is a log parse operation binding the contract event 0x0460adb39e28c2e6aa9aa61a439e432525b0d3df20cef597c3697d5e498fb7ff.
//
// Solidity: event ABIEvent(bool success, bytes returnData)
func (_Storage *StorageFilterer) ParseABIEvent(log types.Log) (*StorageABIEvent, error) {
	event := new(StorageABIEvent)
	if err := _Storage.contract.UnpackLog(event, "ABIEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

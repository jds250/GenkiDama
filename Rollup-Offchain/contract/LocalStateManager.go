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

// LocalStateManagerMetaData contains all meta data concerning the LocalStateManager contract.
var LocalStateManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_merkleUtilsAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_dataTypesAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"now\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"result\",\"type\":\"string\"}],\"name\":\"ConfirmState\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"state\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"error\",\"type\":\"string\"}],\"name\":\"ErrorNotify\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"account\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"style\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"LocalStateNotify\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ZERO_BYTES32\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"committerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"localStatePool\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"account\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lock\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockTransitionNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lockTransitionPool\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockNum\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"account\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"style\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"superAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_account\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"depositCoin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_account\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"refundCoin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_account\",\"type\":\"string\"}],\"name\":\"getLocalState\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"account\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lock\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.LocalState\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"account\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"style\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.L2Transition[]\",\"name\":\"_transits\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"_blockNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_lockTime\",\"type\":\"uint256\"}],\"name\":\"verifyAndLockLocalState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"account\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"style\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.L2Transition[]\",\"name\":\"_transits\",\"type\":\"tuple[]\"}],\"name\":\"verifyLocalState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockNum\",\"type\":\"uint256\"}],\"name\":\"confirmLocalStateWithBlockNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockNum\",\"type\":\"uint256\"}],\"name\":\"removeLocalStateWithBlockNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockNum\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"account\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"style\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockTime\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.LocalTransition[]\",\"name\":\"_transitions\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"}],\"name\":\"checkRoot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// LocalStateManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use LocalStateManagerMetaData.ABI instead.
var LocalStateManagerABI = LocalStateManagerMetaData.ABI

// LocalStateManager is an auto generated Go binding around an Ethereum contract.
type LocalStateManager struct {
	LocalStateManagerCaller     // Read-only binding to the contract
	LocalStateManagerTransactor // Write-only binding to the contract
	LocalStateManagerFilterer   // Log filterer for contract events
}

// LocalStateManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type LocalStateManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LocalStateManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LocalStateManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LocalStateManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LocalStateManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LocalStateManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LocalStateManagerSession struct {
	Contract     *LocalStateManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// LocalStateManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LocalStateManagerCallerSession struct {
	Contract *LocalStateManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// LocalStateManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LocalStateManagerTransactorSession struct {
	Contract     *LocalStateManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// LocalStateManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type LocalStateManagerRaw struct {
	Contract *LocalStateManager // Generic contract binding to access the raw methods on
}

// LocalStateManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LocalStateManagerCallerRaw struct {
	Contract *LocalStateManagerCaller // Generic read-only contract binding to access the raw methods on
}

// LocalStateManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LocalStateManagerTransactorRaw struct {
	Contract *LocalStateManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLocalStateManager creates a new instance of LocalStateManager, bound to a specific deployed contract.
func NewLocalStateManager(address common.Address, backend bind.ContractBackend) (*LocalStateManager, error) {
	contract, err := bindLocalStateManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LocalStateManager{LocalStateManagerCaller: LocalStateManagerCaller{contract: contract}, LocalStateManagerTransactor: LocalStateManagerTransactor{contract: contract}, LocalStateManagerFilterer: LocalStateManagerFilterer{contract: contract}}, nil
}

// NewLocalStateManagerCaller creates a new read-only instance of LocalStateManager, bound to a specific deployed contract.
func NewLocalStateManagerCaller(address common.Address, caller bind.ContractCaller) (*LocalStateManagerCaller, error) {
	contract, err := bindLocalStateManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LocalStateManagerCaller{contract: contract}, nil
}

// NewLocalStateManagerTransactor creates a new write-only instance of LocalStateManager, bound to a specific deployed contract.
func NewLocalStateManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*LocalStateManagerTransactor, error) {
	contract, err := bindLocalStateManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LocalStateManagerTransactor{contract: contract}, nil
}

// NewLocalStateManagerFilterer creates a new log filterer instance of LocalStateManager, bound to a specific deployed contract.
func NewLocalStateManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*LocalStateManagerFilterer, error) {
	contract, err := bindLocalStateManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LocalStateManagerFilterer{contract: contract}, nil
}

// bindLocalStateManager binds a generic wrapper to an already deployed contract.
func bindLocalStateManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LocalStateManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LocalStateManager *LocalStateManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LocalStateManager.Contract.LocalStateManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LocalStateManager *LocalStateManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LocalStateManager.Contract.LocalStateManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LocalStateManager *LocalStateManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LocalStateManager.Contract.LocalStateManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LocalStateManager *LocalStateManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LocalStateManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LocalStateManager *LocalStateManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LocalStateManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LocalStateManager *LocalStateManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LocalStateManager.Contract.contract.Transact(opts, method, params...)
}

// ZEROBYTES32 is a free data retrieval call binding the contract method 0x069321b0.
//
// Solidity: function ZERO_BYTES32() view returns(bytes32)
func (_LocalStateManager *LocalStateManagerCaller) ZEROBYTES32(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LocalStateManager.contract.Call(opts, &out, "ZERO_BYTES32")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ZEROBYTES32 is a free data retrieval call binding the contract method 0x069321b0.
//
// Solidity: function ZERO_BYTES32() view returns(bytes32)
func (_LocalStateManager *LocalStateManagerSession) ZEROBYTES32() ([32]byte, error) {
	return _LocalStateManager.Contract.ZEROBYTES32(&_LocalStateManager.CallOpts)
}

// ZEROBYTES32 is a free data retrieval call binding the contract method 0x069321b0.
//
// Solidity: function ZERO_BYTES32() view returns(bytes32)
func (_LocalStateManager *LocalStateManagerCallerSession) ZEROBYTES32() ([32]byte, error) {
	return _LocalStateManager.Contract.ZEROBYTES32(&_LocalStateManager.CallOpts)
}

// ChainID is a free data retrieval call binding the contract method 0xadc879e9.
//
// Solidity: function chainID() view returns(uint256)
func (_LocalStateManager *LocalStateManagerCaller) ChainID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LocalStateManager.contract.Call(opts, &out, "chainID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChainID is a free data retrieval call binding the contract method 0xadc879e9.
//
// Solidity: function chainID() view returns(uint256)
func (_LocalStateManager *LocalStateManagerSession) ChainID() (*big.Int, error) {
	return _LocalStateManager.Contract.ChainID(&_LocalStateManager.CallOpts)
}

// ChainID is a free data retrieval call binding the contract method 0xadc879e9.
//
// Solidity: function chainID() view returns(uint256)
func (_LocalStateManager *LocalStateManagerCallerSession) ChainID() (*big.Int, error) {
	return _LocalStateManager.Contract.ChainID(&_LocalStateManager.CallOpts)
}

// CheckRoot is a free data retrieval call binding the contract method 0xaafeeb39.
//
// Solidity: function checkRoot((uint256,uint256,string,uint256,uint256,uint256)[] _transitions, bytes32 _root) view returns(bool, bytes32)
func (_LocalStateManager *LocalStateManagerCaller) CheckRoot(opts *bind.CallOpts, _transitions []DataTypesLocalTransition, _root [32]byte) (bool, [32]byte, error) {
	var out []interface{}
	err := _LocalStateManager.contract.Call(opts, &out, "checkRoot", _transitions, _root)

	if err != nil {
		return *new(bool), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return out0, out1, err

}

// CheckRoot is a free data retrieval call binding the contract method 0xaafeeb39.
//
// Solidity: function checkRoot((uint256,uint256,string,uint256,uint256,uint256)[] _transitions, bytes32 _root) view returns(bool, bytes32)
func (_LocalStateManager *LocalStateManagerSession) CheckRoot(_transitions []DataTypesLocalTransition, _root [32]byte) (bool, [32]byte, error) {
	return _LocalStateManager.Contract.CheckRoot(&_LocalStateManager.CallOpts, _transitions, _root)
}

// CheckRoot is a free data retrieval call binding the contract method 0xaafeeb39.
//
// Solidity: function checkRoot((uint256,uint256,string,uint256,uint256,uint256)[] _transitions, bytes32 _root) view returns(bool, bytes32)
func (_LocalStateManager *LocalStateManagerCallerSession) CheckRoot(_transitions []DataTypesLocalTransition, _root [32]byte) (bool, [32]byte, error) {
	return _LocalStateManager.Contract.CheckRoot(&_LocalStateManager.CallOpts, _transitions, _root)
}

// CommitterAddress is a free data retrieval call binding the contract method 0x2b19749e.
//
// Solidity: function committerAddress() view returns(address)
func (_LocalStateManager *LocalStateManagerCaller) CommitterAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LocalStateManager.contract.Call(opts, &out, "committerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CommitterAddress is a free data retrieval call binding the contract method 0x2b19749e.
//
// Solidity: function committerAddress() view returns(address)
func (_LocalStateManager *LocalStateManagerSession) CommitterAddress() (common.Address, error) {
	return _LocalStateManager.Contract.CommitterAddress(&_LocalStateManager.CallOpts)
}

// CommitterAddress is a free data retrieval call binding the contract method 0x2b19749e.
//
// Solidity: function committerAddress() view returns(address)
func (_LocalStateManager *LocalStateManagerCallerSession) CommitterAddress() (common.Address, error) {
	return _LocalStateManager.Contract.CommitterAddress(&_LocalStateManager.CallOpts)
}

// GetLocalState is a free data retrieval call binding the contract method 0xc01f70ac.
//
// Solidity: function getLocalState(string _account) view returns((string,uint256,uint256))
func (_LocalStateManager *LocalStateManagerCaller) GetLocalState(opts *bind.CallOpts, _account string) (DataTypesLocalState, error) {
	var out []interface{}
	err := _LocalStateManager.contract.Call(opts, &out, "getLocalState", _account)

	if err != nil {
		return *new(DataTypesLocalState), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesLocalState)).(*DataTypesLocalState)

	return out0, err

}

// GetLocalState is a free data retrieval call binding the contract method 0xc01f70ac.
//
// Solidity: function getLocalState(string _account) view returns((string,uint256,uint256))
func (_LocalStateManager *LocalStateManagerSession) GetLocalState(_account string) (DataTypesLocalState, error) {
	return _LocalStateManager.Contract.GetLocalState(&_LocalStateManager.CallOpts, _account)
}

// GetLocalState is a free data retrieval call binding the contract method 0xc01f70ac.
//
// Solidity: function getLocalState(string _account) view returns((string,uint256,uint256))
func (_LocalStateManager *LocalStateManagerCallerSession) GetLocalState(_account string) (DataTypesLocalState, error) {
	return _LocalStateManager.Contract.GetLocalState(&_LocalStateManager.CallOpts, _account)
}

// LocalStatePool is a free data retrieval call binding the contract method 0x74bdbc6b.
//
// Solidity: function localStatePool(string ) view returns(string account, uint256 value, uint256 lock)
func (_LocalStateManager *LocalStateManagerCaller) LocalStatePool(opts *bind.CallOpts, arg0 string) (struct {
	Account string
	Value   *big.Int
	Lock    *big.Int
}, error) {
	var out []interface{}
	err := _LocalStateManager.contract.Call(opts, &out, "localStatePool", arg0)

	outstruct := new(struct {
		Account string
		Value   *big.Int
		Lock    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Account = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Value = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Lock = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// LocalStatePool is a free data retrieval call binding the contract method 0x74bdbc6b.
//
// Solidity: function localStatePool(string ) view returns(string account, uint256 value, uint256 lock)
func (_LocalStateManager *LocalStateManagerSession) LocalStatePool(arg0 string) (struct {
	Account string
	Value   *big.Int
	Lock    *big.Int
}, error) {
	return _LocalStateManager.Contract.LocalStatePool(&_LocalStateManager.CallOpts, arg0)
}

// LocalStatePool is a free data retrieval call binding the contract method 0x74bdbc6b.
//
// Solidity: function localStatePool(string ) view returns(string account, uint256 value, uint256 lock)
func (_LocalStateManager *LocalStateManagerCallerSession) LocalStatePool(arg0 string) (struct {
	Account string
	Value   *big.Int
	Lock    *big.Int
}, error) {
	return _LocalStateManager.Contract.LocalStatePool(&_LocalStateManager.CallOpts, arg0)
}

// LockTransitionNum is a free data retrieval call binding the contract method 0xc2c36156.
//
// Solidity: function lockTransitionNum() view returns(uint256)
func (_LocalStateManager *LocalStateManagerCaller) LockTransitionNum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LocalStateManager.contract.Call(opts, &out, "lockTransitionNum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockTransitionNum is a free data retrieval call binding the contract method 0xc2c36156.
//
// Solidity: function lockTransitionNum() view returns(uint256)
func (_LocalStateManager *LocalStateManagerSession) LockTransitionNum() (*big.Int, error) {
	return _LocalStateManager.Contract.LockTransitionNum(&_LocalStateManager.CallOpts)
}

// LockTransitionNum is a free data retrieval call binding the contract method 0xc2c36156.
//
// Solidity: function lockTransitionNum() view returns(uint256)
func (_LocalStateManager *LocalStateManagerCallerSession) LockTransitionNum() (*big.Int, error) {
	return _LocalStateManager.Contract.LockTransitionNum(&_LocalStateManager.CallOpts)
}

// LockTransitionPool is a free data retrieval call binding the contract method 0x1f754bf7.
//
// Solidity: function lockTransitionPool(uint256 ) view returns(uint256 chainID, uint256 blockNum, string account, uint256 value, uint256 style, uint256 unlockTime)
func (_LocalStateManager *LocalStateManagerCaller) LockTransitionPool(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ChainID    *big.Int
	BlockNum   *big.Int
	Account    string
	Value      *big.Int
	Style      *big.Int
	UnlockTime *big.Int
}, error) {
	var out []interface{}
	err := _LocalStateManager.contract.Call(opts, &out, "lockTransitionPool", arg0)

	outstruct := new(struct {
		ChainID    *big.Int
		BlockNum   *big.Int
		Account    string
		Value      *big.Int
		Style      *big.Int
		UnlockTime *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ChainID = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BlockNum = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Account = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Value = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Style = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.UnlockTime = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// LockTransitionPool is a free data retrieval call binding the contract method 0x1f754bf7.
//
// Solidity: function lockTransitionPool(uint256 ) view returns(uint256 chainID, uint256 blockNum, string account, uint256 value, uint256 style, uint256 unlockTime)
func (_LocalStateManager *LocalStateManagerSession) LockTransitionPool(arg0 *big.Int) (struct {
	ChainID    *big.Int
	BlockNum   *big.Int
	Account    string
	Value      *big.Int
	Style      *big.Int
	UnlockTime *big.Int
}, error) {
	return _LocalStateManager.Contract.LockTransitionPool(&_LocalStateManager.CallOpts, arg0)
}

// LockTransitionPool is a free data retrieval call binding the contract method 0x1f754bf7.
//
// Solidity: function lockTransitionPool(uint256 ) view returns(uint256 chainID, uint256 blockNum, string account, uint256 value, uint256 style, uint256 unlockTime)
func (_LocalStateManager *LocalStateManagerCallerSession) LockTransitionPool(arg0 *big.Int) (struct {
	ChainID    *big.Int
	BlockNum   *big.Int
	Account    string
	Value      *big.Int
	Style      *big.Int
	UnlockTime *big.Int
}, error) {
	return _LocalStateManager.Contract.LockTransitionPool(&_LocalStateManager.CallOpts, arg0)
}

// SuperAddress is a free data retrieval call binding the contract method 0x45d3b8db.
//
// Solidity: function superAddress() view returns(address)
func (_LocalStateManager *LocalStateManagerCaller) SuperAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LocalStateManager.contract.Call(opts, &out, "superAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SuperAddress is a free data retrieval call binding the contract method 0x45d3b8db.
//
// Solidity: function superAddress() view returns(address)
func (_LocalStateManager *LocalStateManagerSession) SuperAddress() (common.Address, error) {
	return _LocalStateManager.Contract.SuperAddress(&_LocalStateManager.CallOpts)
}

// SuperAddress is a free data retrieval call binding the contract method 0x45d3b8db.
//
// Solidity: function superAddress() view returns(address)
func (_LocalStateManager *LocalStateManagerCallerSession) SuperAddress() (common.Address, error) {
	return _LocalStateManager.Contract.SuperAddress(&_LocalStateManager.CallOpts)
}

// VerifyLocalState is a free data retrieval call binding the contract method 0x9043ee9b.
//
// Solidity: function verifyLocalState((string,uint256,uint256,uint256)[] _transits) view returns(bool)
func (_LocalStateManager *LocalStateManagerCaller) VerifyLocalState(opts *bind.CallOpts, _transits []DataTypesL2Transition) (bool, error) {
	var out []interface{}
	err := _LocalStateManager.contract.Call(opts, &out, "verifyLocalState", _transits)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyLocalState is a free data retrieval call binding the contract method 0x9043ee9b.
//
// Solidity: function verifyLocalState((string,uint256,uint256,uint256)[] _transits) view returns(bool)
func (_LocalStateManager *LocalStateManagerSession) VerifyLocalState(_transits []DataTypesL2Transition) (bool, error) {
	return _LocalStateManager.Contract.VerifyLocalState(&_LocalStateManager.CallOpts, _transits)
}

// VerifyLocalState is a free data retrieval call binding the contract method 0x9043ee9b.
//
// Solidity: function verifyLocalState((string,uint256,uint256,uint256)[] _transits) view returns(bool)
func (_LocalStateManager *LocalStateManagerCallerSession) VerifyLocalState(_transits []DataTypesL2Transition) (bool, error) {
	return _LocalStateManager.Contract.VerifyLocalState(&_LocalStateManager.CallOpts, _transits)
}

// ConfirmLocalStateWithBlockNum is a paid mutator transaction binding the contract method 0x66955c9a.
//
// Solidity: function confirmLocalStateWithBlockNum(uint256 _blockNum) returns(uint256)
func (_LocalStateManager *LocalStateManagerTransactor) ConfirmLocalStateWithBlockNum(opts *bind.TransactOpts, _blockNum *big.Int) (*types.Transaction, error) {
	return _LocalStateManager.contract.Transact(opts, "confirmLocalStateWithBlockNum", _blockNum)
}

// ConfirmLocalStateWithBlockNum is a paid mutator transaction binding the contract method 0x66955c9a.
//
// Solidity: function confirmLocalStateWithBlockNum(uint256 _blockNum) returns(uint256)
func (_LocalStateManager *LocalStateManagerSession) ConfirmLocalStateWithBlockNum(_blockNum *big.Int) (*types.Transaction, error) {
	return _LocalStateManager.Contract.ConfirmLocalStateWithBlockNum(&_LocalStateManager.TransactOpts, _blockNum)
}

// ConfirmLocalStateWithBlockNum is a paid mutator transaction binding the contract method 0x66955c9a.
//
// Solidity: function confirmLocalStateWithBlockNum(uint256 _blockNum) returns(uint256)
func (_LocalStateManager *LocalStateManagerTransactorSession) ConfirmLocalStateWithBlockNum(_blockNum *big.Int) (*types.Transaction, error) {
	return _LocalStateManager.Contract.ConfirmLocalStateWithBlockNum(&_LocalStateManager.TransactOpts, _blockNum)
}

// DepositCoin is a paid mutator transaction binding the contract method 0x1a06032d.
//
// Solidity: function depositCoin(string _account, uint256 _value) returns()
func (_LocalStateManager *LocalStateManagerTransactor) DepositCoin(opts *bind.TransactOpts, _account string, _value *big.Int) (*types.Transaction, error) {
	return _LocalStateManager.contract.Transact(opts, "depositCoin", _account, _value)
}

// DepositCoin is a paid mutator transaction binding the contract method 0x1a06032d.
//
// Solidity: function depositCoin(string _account, uint256 _value) returns()
func (_LocalStateManager *LocalStateManagerSession) DepositCoin(_account string, _value *big.Int) (*types.Transaction, error) {
	return _LocalStateManager.Contract.DepositCoin(&_LocalStateManager.TransactOpts, _account, _value)
}

// DepositCoin is a paid mutator transaction binding the contract method 0x1a06032d.
//
// Solidity: function depositCoin(string _account, uint256 _value) returns()
func (_LocalStateManager *LocalStateManagerTransactorSession) DepositCoin(_account string, _value *big.Int) (*types.Transaction, error) {
	return _LocalStateManager.Contract.DepositCoin(&_LocalStateManager.TransactOpts, _account, _value)
}

// RefundCoin is a paid mutator transaction binding the contract method 0x9ed1ecc8.
//
// Solidity: function refundCoin(string _account, uint256 _value) returns()
func (_LocalStateManager *LocalStateManagerTransactor) RefundCoin(opts *bind.TransactOpts, _account string, _value *big.Int) (*types.Transaction, error) {
	return _LocalStateManager.contract.Transact(opts, "refundCoin", _account, _value)
}

// RefundCoin is a paid mutator transaction binding the contract method 0x9ed1ecc8.
//
// Solidity: function refundCoin(string _account, uint256 _value) returns()
func (_LocalStateManager *LocalStateManagerSession) RefundCoin(_account string, _value *big.Int) (*types.Transaction, error) {
	return _LocalStateManager.Contract.RefundCoin(&_LocalStateManager.TransactOpts, _account, _value)
}

// RefundCoin is a paid mutator transaction binding the contract method 0x9ed1ecc8.
//
// Solidity: function refundCoin(string _account, uint256 _value) returns()
func (_LocalStateManager *LocalStateManagerTransactorSession) RefundCoin(_account string, _value *big.Int) (*types.Transaction, error) {
	return _LocalStateManager.Contract.RefundCoin(&_LocalStateManager.TransactOpts, _account, _value)
}

// RemoveLocalStateWithBlockNum is a paid mutator transaction binding the contract method 0x74a19b5a.
//
// Solidity: function removeLocalStateWithBlockNum(uint256 _blockNum) returns(uint256)
func (_LocalStateManager *LocalStateManagerTransactor) RemoveLocalStateWithBlockNum(opts *bind.TransactOpts, _blockNum *big.Int) (*types.Transaction, error) {
	return _LocalStateManager.contract.Transact(opts, "removeLocalStateWithBlockNum", _blockNum)
}

// RemoveLocalStateWithBlockNum is a paid mutator transaction binding the contract method 0x74a19b5a.
//
// Solidity: function removeLocalStateWithBlockNum(uint256 _blockNum) returns(uint256)
func (_LocalStateManager *LocalStateManagerSession) RemoveLocalStateWithBlockNum(_blockNum *big.Int) (*types.Transaction, error) {
	return _LocalStateManager.Contract.RemoveLocalStateWithBlockNum(&_LocalStateManager.TransactOpts, _blockNum)
}

// RemoveLocalStateWithBlockNum is a paid mutator transaction binding the contract method 0x74a19b5a.
//
// Solidity: function removeLocalStateWithBlockNum(uint256 _blockNum) returns(uint256)
func (_LocalStateManager *LocalStateManagerTransactorSession) RemoveLocalStateWithBlockNum(_blockNum *big.Int) (*types.Transaction, error) {
	return _LocalStateManager.Contract.RemoveLocalStateWithBlockNum(&_LocalStateManager.TransactOpts, _blockNum)
}

// VerifyAndLockLocalState is a paid mutator transaction binding the contract method 0x9d328c37.
//
// Solidity: function verifyAndLockLocalState((string,uint256,uint256,uint256)[] _transits, uint256 _blockNum, uint256 _lockTime) returns(bool)
func (_LocalStateManager *LocalStateManagerTransactor) VerifyAndLockLocalState(opts *bind.TransactOpts, _transits []DataTypesL2Transition, _blockNum *big.Int, _lockTime *big.Int) (*types.Transaction, error) {
	return _LocalStateManager.contract.Transact(opts, "verifyAndLockLocalState", _transits, _blockNum, _lockTime)
}

// VerifyAndLockLocalState is a paid mutator transaction binding the contract method 0x9d328c37.
//
// Solidity: function verifyAndLockLocalState((string,uint256,uint256,uint256)[] _transits, uint256 _blockNum, uint256 _lockTime) returns(bool)
func (_LocalStateManager *LocalStateManagerSession) VerifyAndLockLocalState(_transits []DataTypesL2Transition, _blockNum *big.Int, _lockTime *big.Int) (*types.Transaction, error) {
	return _LocalStateManager.Contract.VerifyAndLockLocalState(&_LocalStateManager.TransactOpts, _transits, _blockNum, _lockTime)
}

// VerifyAndLockLocalState is a paid mutator transaction binding the contract method 0x9d328c37.
//
// Solidity: function verifyAndLockLocalState((string,uint256,uint256,uint256)[] _transits, uint256 _blockNum, uint256 _lockTime) returns(bool)
func (_LocalStateManager *LocalStateManagerTransactorSession) VerifyAndLockLocalState(_transits []DataTypesL2Transition, _blockNum *big.Int, _lockTime *big.Int) (*types.Transaction, error) {
	return _LocalStateManager.Contract.VerifyAndLockLocalState(&_LocalStateManager.TransactOpts, _transits, _blockNum, _lockTime)
}

// LocalStateManagerConfirmStateIterator is returned from FilterConfirmState and is used to iterate over the raw logs and unpacked data for ConfirmState events raised by the LocalStateManager contract.
type LocalStateManagerConfirmStateIterator struct {
	Event *LocalStateManagerConfirmState // Event containing the contract specifics and raw log

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
func (it *LocalStateManagerConfirmStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LocalStateManagerConfirmState)
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
		it.Event = new(LocalStateManagerConfirmState)
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
func (it *LocalStateManagerConfirmStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LocalStateManagerConfirmStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LocalStateManagerConfirmState represents a ConfirmState event raised by the LocalStateManager contract.
type LocalStateManagerConfirmState struct {
	Unlock *big.Int
	Now    *big.Int
	Result string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterConfirmState is a free log retrieval operation binding the contract event 0xd96e2b90fe7c8bd09d66e5d3f4c4bdbc8ff848eeaa060aee5f7ee1826f05c31a.
//
// Solidity: event ConfirmState(uint256 unlock, uint256 now, string result)
func (_LocalStateManager *LocalStateManagerFilterer) FilterConfirmState(opts *bind.FilterOpts) (*LocalStateManagerConfirmStateIterator, error) {

	logs, sub, err := _LocalStateManager.contract.FilterLogs(opts, "ConfirmState")
	if err != nil {
		return nil, err
	}
	return &LocalStateManagerConfirmStateIterator{contract: _LocalStateManager.contract, event: "ConfirmState", logs: logs, sub: sub}, nil
}

// WatchConfirmState is a free log subscription operation binding the contract event 0xd96e2b90fe7c8bd09d66e5d3f4c4bdbc8ff848eeaa060aee5f7ee1826f05c31a.
//
// Solidity: event ConfirmState(uint256 unlock, uint256 now, string result)
func (_LocalStateManager *LocalStateManagerFilterer) WatchConfirmState(opts *bind.WatchOpts, sink chan<- *LocalStateManagerConfirmState) (event.Subscription, error) {

	logs, sub, err := _LocalStateManager.contract.WatchLogs(opts, "ConfirmState")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LocalStateManagerConfirmState)
				if err := _LocalStateManager.contract.UnpackLog(event, "ConfirmState", log); err != nil {
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

// ParseConfirmState is a log parse operation binding the contract event 0xd96e2b90fe7c8bd09d66e5d3f4c4bdbc8ff848eeaa060aee5f7ee1826f05c31a.
//
// Solidity: event ConfirmState(uint256 unlock, uint256 now, string result)
func (_LocalStateManager *LocalStateManagerFilterer) ParseConfirmState(log types.Log) (*LocalStateManagerConfirmState, error) {
	event := new(LocalStateManagerConfirmState)
	if err := _LocalStateManager.contract.UnpackLog(event, "ConfirmState", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LocalStateManagerErrorNotifyIterator is returned from FilterErrorNotify and is used to iterate over the raw logs and unpacked data for ErrorNotify events raised by the LocalStateManager contract.
type LocalStateManagerErrorNotifyIterator struct {
	Event *LocalStateManagerErrorNotify // Event containing the contract specifics and raw log

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
func (it *LocalStateManagerErrorNotifyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LocalStateManagerErrorNotify)
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
		it.Event = new(LocalStateManagerErrorNotify)
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
func (it *LocalStateManagerErrorNotifyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LocalStateManagerErrorNotifyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LocalStateManagerErrorNotify represents a ErrorNotify event raised by the LocalStateManager contract.
type LocalStateManagerErrorNotify struct {
	State *big.Int
	Error string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterErrorNotify is a free log retrieval operation binding the contract event 0x6307687441ab8283c24801dd603e8b616bd7909ac5cfa1fd14cac706a07d9a99.
//
// Solidity: event ErrorNotify(uint256 state, string error)
func (_LocalStateManager *LocalStateManagerFilterer) FilterErrorNotify(opts *bind.FilterOpts) (*LocalStateManagerErrorNotifyIterator, error) {

	logs, sub, err := _LocalStateManager.contract.FilterLogs(opts, "ErrorNotify")
	if err != nil {
		return nil, err
	}
	return &LocalStateManagerErrorNotifyIterator{contract: _LocalStateManager.contract, event: "ErrorNotify", logs: logs, sub: sub}, nil
}

// WatchErrorNotify is a free log subscription operation binding the contract event 0x6307687441ab8283c24801dd603e8b616bd7909ac5cfa1fd14cac706a07d9a99.
//
// Solidity: event ErrorNotify(uint256 state, string error)
func (_LocalStateManager *LocalStateManagerFilterer) WatchErrorNotify(opts *bind.WatchOpts, sink chan<- *LocalStateManagerErrorNotify) (event.Subscription, error) {

	logs, sub, err := _LocalStateManager.contract.WatchLogs(opts, "ErrorNotify")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LocalStateManagerErrorNotify)
				if err := _LocalStateManager.contract.UnpackLog(event, "ErrorNotify", log); err != nil {
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
func (_LocalStateManager *LocalStateManagerFilterer) ParseErrorNotify(log types.Log) (*LocalStateManagerErrorNotify, error) {
	event := new(LocalStateManagerErrorNotify)
	if err := _LocalStateManager.contract.UnpackLog(event, "ErrorNotify", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LocalStateManagerLocalStateNotifyIterator is returned from FilterLocalStateNotify and is used to iterate over the raw logs and unpacked data for LocalStateNotify events raised by the LocalStateManager contract.
type LocalStateManagerLocalStateNotifyIterator struct {
	Event *LocalStateManagerLocalStateNotify // Event containing the contract specifics and raw log

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
func (it *LocalStateManagerLocalStateNotifyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LocalStateManagerLocalStateNotify)
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
		it.Event = new(LocalStateManagerLocalStateNotify)
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
func (it *LocalStateManagerLocalStateNotifyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LocalStateManagerLocalStateNotifyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LocalStateManagerLocalStateNotify represents a LocalStateNotify event raised by the LocalStateManager contract.
type LocalStateManagerLocalStateNotify struct {
	Account string
	Style   string
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLocalStateNotify is a free log retrieval operation binding the contract event 0xdacd236636c2698694bb8b2c8d2c382259ca60664685470bb58608a1383c5b4c.
//
// Solidity: event LocalStateNotify(string account, string style, uint256 value)
func (_LocalStateManager *LocalStateManagerFilterer) FilterLocalStateNotify(opts *bind.FilterOpts) (*LocalStateManagerLocalStateNotifyIterator, error) {

	logs, sub, err := _LocalStateManager.contract.FilterLogs(opts, "LocalStateNotify")
	if err != nil {
		return nil, err
	}
	return &LocalStateManagerLocalStateNotifyIterator{contract: _LocalStateManager.contract, event: "LocalStateNotify", logs: logs, sub: sub}, nil
}

// WatchLocalStateNotify is a free log subscription operation binding the contract event 0xdacd236636c2698694bb8b2c8d2c382259ca60664685470bb58608a1383c5b4c.
//
// Solidity: event LocalStateNotify(string account, string style, uint256 value)
func (_LocalStateManager *LocalStateManagerFilterer) WatchLocalStateNotify(opts *bind.WatchOpts, sink chan<- *LocalStateManagerLocalStateNotify) (event.Subscription, error) {

	logs, sub, err := _LocalStateManager.contract.WatchLogs(opts, "LocalStateNotify")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LocalStateManagerLocalStateNotify)
				if err := _LocalStateManager.contract.UnpackLog(event, "LocalStateNotify", log); err != nil {
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

// ParseLocalStateNotify is a log parse operation binding the contract event 0xdacd236636c2698694bb8b2c8d2c382259ca60664685470bb58608a1383c5b4c.
//
// Solidity: event LocalStateNotify(string account, string style, uint256 value)
func (_LocalStateManager *LocalStateManagerFilterer) ParseLocalStateNotify(log types.Log) (*LocalStateManagerLocalStateNotify, error) {
	event := new(LocalStateManagerLocalStateNotify)
	if err := _LocalStateManager.contract.UnpackLog(event, "LocalStateNotify", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

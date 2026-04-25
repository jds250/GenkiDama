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

// CrossRollupMetaData contains all meta data concerning the CrossRollup contract.
var CrossRollupMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_committerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_TransactionExecutorAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_MerkleUtilsAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_DataTypesAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_MPTVerifierAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_LocalStateManagerAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"state\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"error\",\"type\":\"string\"}],\"name\":\"ErrorNotify\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"desc\",\"type\":\"string\"}],\"name\":\"RollupBlockNotify\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DISPUTE_TIME\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"L2Blocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"txRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"localTransitionRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"blockSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commitTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WITHDRAW_WAIT_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ZERO_BYTES32\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"challengePool\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l2BlockID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l1BlockID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"l1StateRoot\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"l1BlockData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"challengeSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"committerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"confirmedNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"superAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_committerAddress\",\"type\":\"address\"}],\"name\":\"setCommitterAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BlockLen\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_txRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_localTransitionRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"fromAddr\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"toAddr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"style\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.L2Transaction[]\",\"name\":\"_transactions\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"account\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"style\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.L2Transition[]\",\"name\":\"_transitionList\",\"type\":\"tuple[]\"}],\"name\":\"commitBlock\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkCommit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_targetBlockHeight\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"fromAddr\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"toAddr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"style\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.L2Transaction[]\",\"name\":\"_transactions\",\"type\":\"tuple[]\"}],\"name\":\"ExecutionChallenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"challengeArgu\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verifyCrossBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_targetBlockHeight\",\"type\":\"uint256\"}],\"name\":\"rollBackBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// CrossRollupABI is the input ABI used to generate the binding from.
// Deprecated: Use CrossRollupMetaData.ABI instead.
var CrossRollupABI = CrossRollupMetaData.ABI

// CrossRollup is an auto generated Go binding around an Ethereum contract.
type CrossRollup struct {
	CrossRollupCaller     // Read-only binding to the contract
	CrossRollupTransactor // Write-only binding to the contract
	CrossRollupFilterer   // Log filterer for contract events
}

// CrossRollupCaller is an auto generated read-only Go binding around an Ethereum contract.
type CrossRollupCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossRollupTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CrossRollupTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossRollupFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CrossRollupFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossRollupSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CrossRollupSession struct {
	Contract     *CrossRollup      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CrossRollupCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CrossRollupCallerSession struct {
	Contract *CrossRollupCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// CrossRollupTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CrossRollupTransactorSession struct {
	Contract     *CrossRollupTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// CrossRollupRaw is an auto generated low-level Go binding around an Ethereum contract.
type CrossRollupRaw struct {
	Contract *CrossRollup // Generic contract binding to access the raw methods on
}

// CrossRollupCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CrossRollupCallerRaw struct {
	Contract *CrossRollupCaller // Generic read-only contract binding to access the raw methods on
}

// CrossRollupTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CrossRollupTransactorRaw struct {
	Contract *CrossRollupTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCrossRollup creates a new instance of CrossRollup, bound to a specific deployed contract.
func NewCrossRollup(address common.Address, backend bind.ContractBackend) (*CrossRollup, error) {
	contract, err := bindCrossRollup(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CrossRollup{CrossRollupCaller: CrossRollupCaller{contract: contract}, CrossRollupTransactor: CrossRollupTransactor{contract: contract}, CrossRollupFilterer: CrossRollupFilterer{contract: contract}}, nil
}

// NewCrossRollupCaller creates a new read-only instance of CrossRollup, bound to a specific deployed contract.
func NewCrossRollupCaller(address common.Address, caller bind.ContractCaller) (*CrossRollupCaller, error) {
	contract, err := bindCrossRollup(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CrossRollupCaller{contract: contract}, nil
}

// NewCrossRollupTransactor creates a new write-only instance of CrossRollup, bound to a specific deployed contract.
func NewCrossRollupTransactor(address common.Address, transactor bind.ContractTransactor) (*CrossRollupTransactor, error) {
	contract, err := bindCrossRollup(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CrossRollupTransactor{contract: contract}, nil
}

// NewCrossRollupFilterer creates a new log filterer instance of CrossRollup, bound to a specific deployed contract.
func NewCrossRollupFilterer(address common.Address, filterer bind.ContractFilterer) (*CrossRollupFilterer, error) {
	contract, err := bindCrossRollup(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CrossRollupFilterer{contract: contract}, nil
}

// bindCrossRollup binds a generic wrapper to an already deployed contract.
func bindCrossRollup(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CrossRollupMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrossRollup *CrossRollupRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrossRollup.Contract.CrossRollupCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrossRollup *CrossRollupRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossRollup.Contract.CrossRollupTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrossRollup *CrossRollupRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrossRollup.Contract.CrossRollupTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrossRollup *CrossRollupCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrossRollup.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrossRollup *CrossRollupTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossRollup.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrossRollup *CrossRollupTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrossRollup.Contract.contract.Transact(opts, method, params...)
}

// BlockLen is a free data retrieval call binding the contract method 0x94277f84.
//
// Solidity: function BlockLen() view returns(uint256)
func (_CrossRollup *CrossRollupCaller) BlockLen(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrossRollup.contract.Call(opts, &out, "BlockLen")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BlockLen is a free data retrieval call binding the contract method 0x94277f84.
//
// Solidity: function BlockLen() view returns(uint256)
func (_CrossRollup *CrossRollupSession) BlockLen() (*big.Int, error) {
	return _CrossRollup.Contract.BlockLen(&_CrossRollup.CallOpts)
}

// BlockLen is a free data retrieval call binding the contract method 0x94277f84.
//
// Solidity: function BlockLen() view returns(uint256)
func (_CrossRollup *CrossRollupCallerSession) BlockLen() (*big.Int, error) {
	return _CrossRollup.Contract.BlockLen(&_CrossRollup.CallOpts)
}

// DISPUTETIME is a free data retrieval call binding the contract method 0x380e9c56.
//
// Solidity: function DISPUTE_TIME() view returns(uint256)
func (_CrossRollup *CrossRollupCaller) DISPUTETIME(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrossRollup.contract.Call(opts, &out, "DISPUTE_TIME")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DISPUTETIME is a free data retrieval call binding the contract method 0x380e9c56.
//
// Solidity: function DISPUTE_TIME() view returns(uint256)
func (_CrossRollup *CrossRollupSession) DISPUTETIME() (*big.Int, error) {
	return _CrossRollup.Contract.DISPUTETIME(&_CrossRollup.CallOpts)
}

// DISPUTETIME is a free data retrieval call binding the contract method 0x380e9c56.
//
// Solidity: function DISPUTE_TIME() view returns(uint256)
func (_CrossRollup *CrossRollupCallerSession) DISPUTETIME() (*big.Int, error) {
	return _CrossRollup.Contract.DISPUTETIME(&_CrossRollup.CallOpts)
}

// L2Blocks is a free data retrieval call binding the contract method 0xddff9280.
//
// Solidity: function L2Blocks(uint256 ) view returns(uint256 height, bytes32 txRoot, bytes32 localTransitionRoot, bytes32 stateRoot, uint256 blockSize, uint256 commitTime)
func (_CrossRollup *CrossRollupCaller) L2Blocks(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Height              *big.Int
	TxRoot              [32]byte
	LocalTransitionRoot [32]byte
	StateRoot           [32]byte
	BlockSize           *big.Int
	CommitTime          *big.Int
}, error) {
	var out []interface{}
	err := _CrossRollup.contract.Call(opts, &out, "L2Blocks", arg0)

	outstruct := new(struct {
		Height              *big.Int
		TxRoot              [32]byte
		LocalTransitionRoot [32]byte
		StateRoot           [32]byte
		BlockSize           *big.Int
		CommitTime          *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Height = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TxRoot = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.LocalTransitionRoot = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	outstruct.StateRoot = *abi.ConvertType(out[3], new([32]byte)).(*[32]byte)
	outstruct.BlockSize = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.CommitTime = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// L2Blocks is a free data retrieval call binding the contract method 0xddff9280.
//
// Solidity: function L2Blocks(uint256 ) view returns(uint256 height, bytes32 txRoot, bytes32 localTransitionRoot, bytes32 stateRoot, uint256 blockSize, uint256 commitTime)
func (_CrossRollup *CrossRollupSession) L2Blocks(arg0 *big.Int) (struct {
	Height              *big.Int
	TxRoot              [32]byte
	LocalTransitionRoot [32]byte
	StateRoot           [32]byte
	BlockSize           *big.Int
	CommitTime          *big.Int
}, error) {
	return _CrossRollup.Contract.L2Blocks(&_CrossRollup.CallOpts, arg0)
}

// L2Blocks is a free data retrieval call binding the contract method 0xddff9280.
//
// Solidity: function L2Blocks(uint256 ) view returns(uint256 height, bytes32 txRoot, bytes32 localTransitionRoot, bytes32 stateRoot, uint256 blockSize, uint256 commitTime)
func (_CrossRollup *CrossRollupCallerSession) L2Blocks(arg0 *big.Int) (struct {
	Height              *big.Int
	TxRoot              [32]byte
	LocalTransitionRoot [32]byte
	StateRoot           [32]byte
	BlockSize           *big.Int
	CommitTime          *big.Int
}, error) {
	return _CrossRollup.Contract.L2Blocks(&_CrossRollup.CallOpts, arg0)
}

// WITHDRAWWAITPERIOD is a free data retrieval call binding the contract method 0x9381c226.
//
// Solidity: function WITHDRAW_WAIT_PERIOD() view returns(uint256)
func (_CrossRollup *CrossRollupCaller) WITHDRAWWAITPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrossRollup.contract.Call(opts, &out, "WITHDRAW_WAIT_PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WITHDRAWWAITPERIOD is a free data retrieval call binding the contract method 0x9381c226.
//
// Solidity: function WITHDRAW_WAIT_PERIOD() view returns(uint256)
func (_CrossRollup *CrossRollupSession) WITHDRAWWAITPERIOD() (*big.Int, error) {
	return _CrossRollup.Contract.WITHDRAWWAITPERIOD(&_CrossRollup.CallOpts)
}

// WITHDRAWWAITPERIOD is a free data retrieval call binding the contract method 0x9381c226.
//
// Solidity: function WITHDRAW_WAIT_PERIOD() view returns(uint256)
func (_CrossRollup *CrossRollupCallerSession) WITHDRAWWAITPERIOD() (*big.Int, error) {
	return _CrossRollup.Contract.WITHDRAWWAITPERIOD(&_CrossRollup.CallOpts)
}

// ZEROBYTES32 is a free data retrieval call binding the contract method 0x069321b0.
//
// Solidity: function ZERO_BYTES32() view returns(bytes32)
func (_CrossRollup *CrossRollupCaller) ZEROBYTES32(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CrossRollup.contract.Call(opts, &out, "ZERO_BYTES32")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ZEROBYTES32 is a free data retrieval call binding the contract method 0x069321b0.
//
// Solidity: function ZERO_BYTES32() view returns(bytes32)
func (_CrossRollup *CrossRollupSession) ZEROBYTES32() ([32]byte, error) {
	return _CrossRollup.Contract.ZEROBYTES32(&_CrossRollup.CallOpts)
}

// ZEROBYTES32 is a free data retrieval call binding the contract method 0x069321b0.
//
// Solidity: function ZERO_BYTES32() view returns(bytes32)
func (_CrossRollup *CrossRollupCallerSession) ZEROBYTES32() ([32]byte, error) {
	return _CrossRollup.Contract.ZEROBYTES32(&_CrossRollup.CallOpts)
}

// ChainID is a free data retrieval call binding the contract method 0xadc879e9.
//
// Solidity: function chainID() view returns(uint256)
func (_CrossRollup *CrossRollupCaller) ChainID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrossRollup.contract.Call(opts, &out, "chainID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChainID is a free data retrieval call binding the contract method 0xadc879e9.
//
// Solidity: function chainID() view returns(uint256)
func (_CrossRollup *CrossRollupSession) ChainID() (*big.Int, error) {
	return _CrossRollup.Contract.ChainID(&_CrossRollup.CallOpts)
}

// ChainID is a free data retrieval call binding the contract method 0xadc879e9.
//
// Solidity: function chainID() view returns(uint256)
func (_CrossRollup *CrossRollupCallerSession) ChainID() (*big.Int, error) {
	return _CrossRollup.Contract.ChainID(&_CrossRollup.CallOpts)
}

// ChallengePool is a free data retrieval call binding the contract method 0xeae5561a.
//
// Solidity: function challengePool(uint256 ) view returns(uint256 chainID, uint256 l2BlockID, uint256 l1BlockID, bytes l1StateRoot, bytes l1BlockData, address account)
func (_CrossRollup *CrossRollupCaller) ChallengePool(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ChainID     *big.Int
	L2BlockID   *big.Int
	L1BlockID   *big.Int
	L1StateRoot []byte
	L1BlockData []byte
	Account     common.Address
}, error) {
	var out []interface{}
	err := _CrossRollup.contract.Call(opts, &out, "challengePool", arg0)

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
func (_CrossRollup *CrossRollupSession) ChallengePool(arg0 *big.Int) (struct {
	ChainID     *big.Int
	L2BlockID   *big.Int
	L1BlockID   *big.Int
	L1StateRoot []byte
	L1BlockData []byte
	Account     common.Address
}, error) {
	return _CrossRollup.Contract.ChallengePool(&_CrossRollup.CallOpts, arg0)
}

// ChallengePool is a free data retrieval call binding the contract method 0xeae5561a.
//
// Solidity: function challengePool(uint256 ) view returns(uint256 chainID, uint256 l2BlockID, uint256 l1BlockID, bytes l1StateRoot, bytes l1BlockData, address account)
func (_CrossRollup *CrossRollupCallerSession) ChallengePool(arg0 *big.Int) (struct {
	ChainID     *big.Int
	L2BlockID   *big.Int
	L1BlockID   *big.Int
	L1StateRoot []byte
	L1BlockData []byte
	Account     common.Address
}, error) {
	return _CrossRollup.Contract.ChallengePool(&_CrossRollup.CallOpts, arg0)
}

// ChallengeSize is a free data retrieval call binding the contract method 0x6316bc17.
//
// Solidity: function challengeSize() view returns(uint256)
func (_CrossRollup *CrossRollupCaller) ChallengeSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrossRollup.contract.Call(opts, &out, "challengeSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChallengeSize is a free data retrieval call binding the contract method 0x6316bc17.
//
// Solidity: function challengeSize() view returns(uint256)
func (_CrossRollup *CrossRollupSession) ChallengeSize() (*big.Int, error) {
	return _CrossRollup.Contract.ChallengeSize(&_CrossRollup.CallOpts)
}

// ChallengeSize is a free data retrieval call binding the contract method 0x6316bc17.
//
// Solidity: function challengeSize() view returns(uint256)
func (_CrossRollup *CrossRollupCallerSession) ChallengeSize() (*big.Int, error) {
	return _CrossRollup.Contract.ChallengeSize(&_CrossRollup.CallOpts)
}

// CommitterAddress is a free data retrieval call binding the contract method 0x2b19749e.
//
// Solidity: function committerAddress() view returns(address)
func (_CrossRollup *CrossRollupCaller) CommitterAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossRollup.contract.Call(opts, &out, "committerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CommitterAddress is a free data retrieval call binding the contract method 0x2b19749e.
//
// Solidity: function committerAddress() view returns(address)
func (_CrossRollup *CrossRollupSession) CommitterAddress() (common.Address, error) {
	return _CrossRollup.Contract.CommitterAddress(&_CrossRollup.CallOpts)
}

// CommitterAddress is a free data retrieval call binding the contract method 0x2b19749e.
//
// Solidity: function committerAddress() view returns(address)
func (_CrossRollup *CrossRollupCallerSession) CommitterAddress() (common.Address, error) {
	return _CrossRollup.Contract.CommitterAddress(&_CrossRollup.CallOpts)
}

// ConfirmedNum is a free data retrieval call binding the contract method 0x88d1e611.
//
// Solidity: function confirmedNum() view returns(uint256)
func (_CrossRollup *CrossRollupCaller) ConfirmedNum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrossRollup.contract.Call(opts, &out, "confirmedNum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConfirmedNum is a free data retrieval call binding the contract method 0x88d1e611.
//
// Solidity: function confirmedNum() view returns(uint256)
func (_CrossRollup *CrossRollupSession) ConfirmedNum() (*big.Int, error) {
	return _CrossRollup.Contract.ConfirmedNum(&_CrossRollup.CallOpts)
}

// ConfirmedNum is a free data retrieval call binding the contract method 0x88d1e611.
//
// Solidity: function confirmedNum() view returns(uint256)
func (_CrossRollup *CrossRollupCallerSession) ConfirmedNum() (*big.Int, error) {
	return _CrossRollup.Contract.ConfirmedNum(&_CrossRollup.CallOpts)
}

// GetCurrentBlockNumber is a free data retrieval call binding the contract method 0x6fd902e1.
//
// Solidity: function getCurrentBlockNumber() view returns(uint256)
func (_CrossRollup *CrossRollupCaller) GetCurrentBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrossRollup.contract.Call(opts, &out, "getCurrentBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentBlockNumber is a free data retrieval call binding the contract method 0x6fd902e1.
//
// Solidity: function getCurrentBlockNumber() view returns(uint256)
func (_CrossRollup *CrossRollupSession) GetCurrentBlockNumber() (*big.Int, error) {
	return _CrossRollup.Contract.GetCurrentBlockNumber(&_CrossRollup.CallOpts)
}

// GetCurrentBlockNumber is a free data retrieval call binding the contract method 0x6fd902e1.
//
// Solidity: function getCurrentBlockNumber() view returns(uint256)
func (_CrossRollup *CrossRollupCallerSession) GetCurrentBlockNumber() (*big.Int, error) {
	return _CrossRollup.Contract.GetCurrentBlockNumber(&_CrossRollup.CallOpts)
}

// SuperAddress is a free data retrieval call binding the contract method 0x45d3b8db.
//
// Solidity: function superAddress() view returns(address)
func (_CrossRollup *CrossRollupCaller) SuperAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossRollup.contract.Call(opts, &out, "superAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SuperAddress is a free data retrieval call binding the contract method 0x45d3b8db.
//
// Solidity: function superAddress() view returns(address)
func (_CrossRollup *CrossRollupSession) SuperAddress() (common.Address, error) {
	return _CrossRollup.Contract.SuperAddress(&_CrossRollup.CallOpts)
}

// SuperAddress is a free data retrieval call binding the contract method 0x45d3b8db.
//
// Solidity: function superAddress() view returns(address)
func (_CrossRollup *CrossRollupCallerSession) SuperAddress() (common.Address, error) {
	return _CrossRollup.Contract.SuperAddress(&_CrossRollup.CallOpts)
}

// ExecutionChallenge is a paid mutator transaction binding the contract method 0x770ce0f0.
//
// Solidity: function ExecutionChallenge(uint256 _targetBlockHeight, (string,string,uint256,uint256,uint256)[] _transactions) returns()
func (_CrossRollup *CrossRollupTransactor) ExecutionChallenge(opts *bind.TransactOpts, _targetBlockHeight *big.Int, _transactions []DataTypesL2Transaction) (*types.Transaction, error) {
	return _CrossRollup.contract.Transact(opts, "ExecutionChallenge", _targetBlockHeight, _transactions)
}

// ExecutionChallenge is a paid mutator transaction binding the contract method 0x770ce0f0.
//
// Solidity: function ExecutionChallenge(uint256 _targetBlockHeight, (string,string,uint256,uint256,uint256)[] _transactions) returns()
func (_CrossRollup *CrossRollupSession) ExecutionChallenge(_targetBlockHeight *big.Int, _transactions []DataTypesL2Transaction) (*types.Transaction, error) {
	return _CrossRollup.Contract.ExecutionChallenge(&_CrossRollup.TransactOpts, _targetBlockHeight, _transactions)
}

// ExecutionChallenge is a paid mutator transaction binding the contract method 0x770ce0f0.
//
// Solidity: function ExecutionChallenge(uint256 _targetBlockHeight, (string,string,uint256,uint256,uint256)[] _transactions) returns()
func (_CrossRollup *CrossRollupTransactorSession) ExecutionChallenge(_targetBlockHeight *big.Int, _transactions []DataTypesL2Transaction) (*types.Transaction, error) {
	return _CrossRollup.Contract.ExecutionChallenge(&_CrossRollup.TransactOpts, _targetBlockHeight, _transactions)
}

// ChallengeArgu is a paid mutator transaction binding the contract method 0xfdac640f.
//
// Solidity: function challengeArgu() returns()
func (_CrossRollup *CrossRollupTransactor) ChallengeArgu(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossRollup.contract.Transact(opts, "challengeArgu")
}

// ChallengeArgu is a paid mutator transaction binding the contract method 0xfdac640f.
//
// Solidity: function challengeArgu() returns()
func (_CrossRollup *CrossRollupSession) ChallengeArgu() (*types.Transaction, error) {
	return _CrossRollup.Contract.ChallengeArgu(&_CrossRollup.TransactOpts)
}

// ChallengeArgu is a paid mutator transaction binding the contract method 0xfdac640f.
//
// Solidity: function challengeArgu() returns()
func (_CrossRollup *CrossRollupTransactorSession) ChallengeArgu() (*types.Transaction, error) {
	return _CrossRollup.Contract.ChallengeArgu(&_CrossRollup.TransactOpts)
}

// CheckCommit is a paid mutator transaction binding the contract method 0x551ee37e.
//
// Solidity: function checkCommit() returns()
func (_CrossRollup *CrossRollupTransactor) CheckCommit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossRollup.contract.Transact(opts, "checkCommit")
}

// CheckCommit is a paid mutator transaction binding the contract method 0x551ee37e.
//
// Solidity: function checkCommit() returns()
func (_CrossRollup *CrossRollupSession) CheckCommit() (*types.Transaction, error) {
	return _CrossRollup.Contract.CheckCommit(&_CrossRollup.TransactOpts)
}

// CheckCommit is a paid mutator transaction binding the contract method 0x551ee37e.
//
// Solidity: function checkCommit() returns()
func (_CrossRollup *CrossRollupTransactorSession) CheckCommit() (*types.Transaction, error) {
	return _CrossRollup.Contract.CheckCommit(&_CrossRollup.TransactOpts)
}

// CommitBlock is a paid mutator transaction binding the contract method 0x59c25f99.
//
// Solidity: function commitBlock(uint256 _blockNumber, bytes32 _txRoot, bytes32 _localTransitionRoot, (string,string,uint256,uint256,uint256)[] _transactions, (string,uint256,uint256,uint256)[] _transitionList) returns(bool)
func (_CrossRollup *CrossRollupTransactor) CommitBlock(opts *bind.TransactOpts, _blockNumber *big.Int, _txRoot [32]byte, _localTransitionRoot [32]byte, _transactions []DataTypesL2Transaction, _transitionList []DataTypesL2Transition) (*types.Transaction, error) {
	return _CrossRollup.contract.Transact(opts, "commitBlock", _blockNumber, _txRoot, _localTransitionRoot, _transactions, _transitionList)
}

// CommitBlock is a paid mutator transaction binding the contract method 0x59c25f99.
//
// Solidity: function commitBlock(uint256 _blockNumber, bytes32 _txRoot, bytes32 _localTransitionRoot, (string,string,uint256,uint256,uint256)[] _transactions, (string,uint256,uint256,uint256)[] _transitionList) returns(bool)
func (_CrossRollup *CrossRollupSession) CommitBlock(_blockNumber *big.Int, _txRoot [32]byte, _localTransitionRoot [32]byte, _transactions []DataTypesL2Transaction, _transitionList []DataTypesL2Transition) (*types.Transaction, error) {
	return _CrossRollup.Contract.CommitBlock(&_CrossRollup.TransactOpts, _blockNumber, _txRoot, _localTransitionRoot, _transactions, _transitionList)
}

// CommitBlock is a paid mutator transaction binding the contract method 0x59c25f99.
//
// Solidity: function commitBlock(uint256 _blockNumber, bytes32 _txRoot, bytes32 _localTransitionRoot, (string,string,uint256,uint256,uint256)[] _transactions, (string,uint256,uint256,uint256)[] _transitionList) returns(bool)
func (_CrossRollup *CrossRollupTransactorSession) CommitBlock(_blockNumber *big.Int, _txRoot [32]byte, _localTransitionRoot [32]byte, _transactions []DataTypesL2Transaction, _transitionList []DataTypesL2Transition) (*types.Transaction, error) {
	return _CrossRollup.Contract.CommitBlock(&_CrossRollup.TransactOpts, _blockNumber, _txRoot, _localTransitionRoot, _transactions, _transitionList)
}

// RollBackBlock is a paid mutator transaction binding the contract method 0x3e711294.
//
// Solidity: function rollBackBlock(uint256 _targetBlockHeight) returns()
func (_CrossRollup *CrossRollupTransactor) RollBackBlock(opts *bind.TransactOpts, _targetBlockHeight *big.Int) (*types.Transaction, error) {
	return _CrossRollup.contract.Transact(opts, "rollBackBlock", _targetBlockHeight)
}

// RollBackBlock is a paid mutator transaction binding the contract method 0x3e711294.
//
// Solidity: function rollBackBlock(uint256 _targetBlockHeight) returns()
func (_CrossRollup *CrossRollupSession) RollBackBlock(_targetBlockHeight *big.Int) (*types.Transaction, error) {
	return _CrossRollup.Contract.RollBackBlock(&_CrossRollup.TransactOpts, _targetBlockHeight)
}

// RollBackBlock is a paid mutator transaction binding the contract method 0x3e711294.
//
// Solidity: function rollBackBlock(uint256 _targetBlockHeight) returns()
func (_CrossRollup *CrossRollupTransactorSession) RollBackBlock(_targetBlockHeight *big.Int) (*types.Transaction, error) {
	return _CrossRollup.Contract.RollBackBlock(&_CrossRollup.TransactOpts, _targetBlockHeight)
}

// SetCommitterAddress is a paid mutator transaction binding the contract method 0x13df8728.
//
// Solidity: function setCommitterAddress(address _committerAddress) returns()
func (_CrossRollup *CrossRollupTransactor) SetCommitterAddress(opts *bind.TransactOpts, _committerAddress common.Address) (*types.Transaction, error) {
	return _CrossRollup.contract.Transact(opts, "setCommitterAddress", _committerAddress)
}

// SetCommitterAddress is a paid mutator transaction binding the contract method 0x13df8728.
//
// Solidity: function setCommitterAddress(address _committerAddress) returns()
func (_CrossRollup *CrossRollupSession) SetCommitterAddress(_committerAddress common.Address) (*types.Transaction, error) {
	return _CrossRollup.Contract.SetCommitterAddress(&_CrossRollup.TransactOpts, _committerAddress)
}

// SetCommitterAddress is a paid mutator transaction binding the contract method 0x13df8728.
//
// Solidity: function setCommitterAddress(address _committerAddress) returns()
func (_CrossRollup *CrossRollupTransactorSession) SetCommitterAddress(_committerAddress common.Address) (*types.Transaction, error) {
	return _CrossRollup.Contract.SetCommitterAddress(&_CrossRollup.TransactOpts, _committerAddress)
}

// VerifyCrossBlock is a paid mutator transaction binding the contract method 0x9b854167.
//
// Solidity: function verifyCrossBlock() returns()
func (_CrossRollup *CrossRollupTransactor) VerifyCrossBlock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossRollup.contract.Transact(opts, "verifyCrossBlock")
}

// VerifyCrossBlock is a paid mutator transaction binding the contract method 0x9b854167.
//
// Solidity: function verifyCrossBlock() returns()
func (_CrossRollup *CrossRollupSession) VerifyCrossBlock() (*types.Transaction, error) {
	return _CrossRollup.Contract.VerifyCrossBlock(&_CrossRollup.TransactOpts)
}

// VerifyCrossBlock is a paid mutator transaction binding the contract method 0x9b854167.
//
// Solidity: function verifyCrossBlock() returns()
func (_CrossRollup *CrossRollupTransactorSession) VerifyCrossBlock() (*types.Transaction, error) {
	return _CrossRollup.Contract.VerifyCrossBlock(&_CrossRollup.TransactOpts)
}

// CrossRollupErrorNotifyIterator is returned from FilterErrorNotify and is used to iterate over the raw logs and unpacked data for ErrorNotify events raised by the CrossRollup contract.
type CrossRollupErrorNotifyIterator struct {
	Event *CrossRollupErrorNotify // Event containing the contract specifics and raw log

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
func (it *CrossRollupErrorNotifyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossRollupErrorNotify)
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
		it.Event = new(CrossRollupErrorNotify)
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
func (it *CrossRollupErrorNotifyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossRollupErrorNotifyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossRollupErrorNotify represents a ErrorNotify event raised by the CrossRollup contract.
type CrossRollupErrorNotify struct {
	State *big.Int
	Error string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterErrorNotify is a free log retrieval operation binding the contract event 0x6307687441ab8283c24801dd603e8b616bd7909ac5cfa1fd14cac706a07d9a99.
//
// Solidity: event ErrorNotify(uint256 state, string error)
func (_CrossRollup *CrossRollupFilterer) FilterErrorNotify(opts *bind.FilterOpts) (*CrossRollupErrorNotifyIterator, error) {

	logs, sub, err := _CrossRollup.contract.FilterLogs(opts, "ErrorNotify")
	if err != nil {
		return nil, err
	}
	return &CrossRollupErrorNotifyIterator{contract: _CrossRollup.contract, event: "ErrorNotify", logs: logs, sub: sub}, nil
}

// WatchErrorNotify is a free log subscription operation binding the contract event 0x6307687441ab8283c24801dd603e8b616bd7909ac5cfa1fd14cac706a07d9a99.
//
// Solidity: event ErrorNotify(uint256 state, string error)
func (_CrossRollup *CrossRollupFilterer) WatchErrorNotify(opts *bind.WatchOpts, sink chan<- *CrossRollupErrorNotify) (event.Subscription, error) {

	logs, sub, err := _CrossRollup.contract.WatchLogs(opts, "ErrorNotify")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossRollupErrorNotify)
				if err := _CrossRollup.contract.UnpackLog(event, "ErrorNotify", log); err != nil {
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
func (_CrossRollup *CrossRollupFilterer) ParseErrorNotify(log types.Log) (*CrossRollupErrorNotify, error) {
	event := new(CrossRollupErrorNotify)
	if err := _CrossRollup.contract.UnpackLog(event, "ErrorNotify", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossRollupRollupBlockNotifyIterator is returned from FilterRollupBlockNotify and is used to iterate over the raw logs and unpacked data for RollupBlockNotify events raised by the CrossRollup contract.
type CrossRollupRollupBlockNotifyIterator struct {
	Event *CrossRollupRollupBlockNotify // Event containing the contract specifics and raw log

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
func (it *CrossRollupRollupBlockNotifyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossRollupRollupBlockNotify)
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
		it.Event = new(CrossRollupRollupBlockNotify)
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
func (it *CrossRollupRollupBlockNotifyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossRollupRollupBlockNotifyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossRollupRollupBlockNotify represents a RollupBlockNotify event raised by the CrossRollup contract.
type CrossRollupRollupBlockNotify struct {
	BlockNumber *big.Int
	Desc        string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRollupBlockNotify is a free log retrieval operation binding the contract event 0x5e19c79d1d9c171225be02ab9d3a12ed7c13ccf394c6919a8d45a6bd4e66f6b6.
//
// Solidity: event RollupBlockNotify(uint256 blockNumber, string desc)
func (_CrossRollup *CrossRollupFilterer) FilterRollupBlockNotify(opts *bind.FilterOpts) (*CrossRollupRollupBlockNotifyIterator, error) {

	logs, sub, err := _CrossRollup.contract.FilterLogs(opts, "RollupBlockNotify")
	if err != nil {
		return nil, err
	}
	return &CrossRollupRollupBlockNotifyIterator{contract: _CrossRollup.contract, event: "RollupBlockNotify", logs: logs, sub: sub}, nil
}

// WatchRollupBlockNotify is a free log subscription operation binding the contract event 0x5e19c79d1d9c171225be02ab9d3a12ed7c13ccf394c6919a8d45a6bd4e66f6b6.
//
// Solidity: event RollupBlockNotify(uint256 blockNumber, string desc)
func (_CrossRollup *CrossRollupFilterer) WatchRollupBlockNotify(opts *bind.WatchOpts, sink chan<- *CrossRollupRollupBlockNotify) (event.Subscription, error) {

	logs, sub, err := _CrossRollup.contract.WatchLogs(opts, "RollupBlockNotify")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossRollupRollupBlockNotify)
				if err := _CrossRollup.contract.UnpackLog(event, "RollupBlockNotify", log); err != nil {
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
func (_CrossRollup *CrossRollupFilterer) ParseRollupBlockNotify(log types.Log) (*CrossRollupRollupBlockNotify, error) {
	event := new(CrossRollupRollupBlockNotify)
	if err := _CrossRollup.contract.UnpackLog(event, "RollupBlockNotify", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

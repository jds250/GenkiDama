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

// MerkelUtilsMetaData contains all meta data concerning the MerkelUtils contract.
var MerkelUtilsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"defaultHashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"tree\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_dataBlocks\",\"type\":\"bytes[]\"}],\"name\":\"getMerkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_dataBlock\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_path\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_siblings\",\"type\":\"bytes32[]\"}],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_dataBlock\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_path\",\"type\":\"uint256\"}],\"name\":\"update\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_leaf\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_path\",\"type\":\"uint256\"}],\"name\":\"updateLeaf\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_dataBlock\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_path\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_siblings\",\"type\":\"bytes32[]\"}],\"name\":\"verifyAndStore\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_dataBlock\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_path\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_siblings\",\"type\":\"bytes32[]\"}],\"name\":\"store\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_leaf\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_path\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_siblings\",\"type\":\"bytes32[]\"}],\"name\":\"storeLeaf\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_path\",\"type\":\"uint256\"}],\"name\":\"getSiblings\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"getRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_height\",\"type\":\"uint256\"}],\"name\":\"setMerkleRootAndHeight\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_parent\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_leftChild\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_rightChild\",\"type\":\"bytes32\"}],\"name\":\"storeNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_intVal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getNthBitFromRight\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_parent\",\"type\":\"bytes32\"}],\"name\":\"getChildren\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_parent\",\"type\":\"bytes32\"}],\"name\":\"getLeftSiblingKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_parent\",\"type\":\"bytes32\"}],\"name\":\"getRightSiblingKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true}]",
}

// MerkelUtilsABI is the input ABI used to generate the binding from.
// Deprecated: Use MerkelUtilsMetaData.ABI instead.
var MerkelUtilsABI = MerkelUtilsMetaData.ABI

// MerkelUtils is an auto generated Go binding around an Ethereum contract.
type MerkelUtils struct {
	MerkelUtilsCaller     // Read-only binding to the contract
	MerkelUtilsTransactor // Write-only binding to the contract
	MerkelUtilsFilterer   // Log filterer for contract events
}

// MerkelUtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type MerkelUtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkelUtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MerkelUtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkelUtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MerkelUtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkelUtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MerkelUtilsSession struct {
	Contract     *MerkelUtils      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MerkelUtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MerkelUtilsCallerSession struct {
	Contract *MerkelUtilsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// MerkelUtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MerkelUtilsTransactorSession struct {
	Contract     *MerkelUtilsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MerkelUtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type MerkelUtilsRaw struct {
	Contract *MerkelUtils // Generic contract binding to access the raw methods on
}

// MerkelUtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MerkelUtilsCallerRaw struct {
	Contract *MerkelUtilsCaller // Generic read-only contract binding to access the raw methods on
}

// MerkelUtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MerkelUtilsTransactorRaw struct {
	Contract *MerkelUtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMerkelUtils creates a new instance of MerkelUtils, bound to a specific deployed contract.
func NewMerkelUtils(address common.Address, backend bind.ContractBackend) (*MerkelUtils, error) {
	contract, err := bindMerkelUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MerkelUtils{MerkelUtilsCaller: MerkelUtilsCaller{contract: contract}, MerkelUtilsTransactor: MerkelUtilsTransactor{contract: contract}, MerkelUtilsFilterer: MerkelUtilsFilterer{contract: contract}}, nil
}

// NewMerkelUtilsCaller creates a new read-only instance of MerkelUtils, bound to a specific deployed contract.
func NewMerkelUtilsCaller(address common.Address, caller bind.ContractCaller) (*MerkelUtilsCaller, error) {
	contract, err := bindMerkelUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MerkelUtilsCaller{contract: contract}, nil
}

// NewMerkelUtilsTransactor creates a new write-only instance of MerkelUtils, bound to a specific deployed contract.
func NewMerkelUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*MerkelUtilsTransactor, error) {
	contract, err := bindMerkelUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MerkelUtilsTransactor{contract: contract}, nil
}

// NewMerkelUtilsFilterer creates a new log filterer instance of MerkelUtils, bound to a specific deployed contract.
func NewMerkelUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*MerkelUtilsFilterer, error) {
	contract, err := bindMerkelUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MerkelUtilsFilterer{contract: contract}, nil
}

// bindMerkelUtils binds a generic wrapper to an already deployed contract.
func bindMerkelUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MerkelUtilsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerkelUtils *MerkelUtilsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MerkelUtils.Contract.MerkelUtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerkelUtils *MerkelUtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerkelUtils.Contract.MerkelUtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerkelUtils *MerkelUtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerkelUtils.Contract.MerkelUtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerkelUtils *MerkelUtilsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MerkelUtils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerkelUtils *MerkelUtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerkelUtils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerkelUtils *MerkelUtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerkelUtils.Contract.contract.Transact(opts, method, params...)
}

// DefaultHashes is a free data retrieval call binding the contract method 0x48419ad8.
//
// Solidity: function defaultHashes(uint256 ) view returns(bytes32)
func (_MerkelUtils *MerkelUtilsCaller) DefaultHashes(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _MerkelUtils.contract.Call(opts, &out, "defaultHashes", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DefaultHashes is a free data retrieval call binding the contract method 0x48419ad8.
//
// Solidity: function defaultHashes(uint256 ) view returns(bytes32)
func (_MerkelUtils *MerkelUtilsSession) DefaultHashes(arg0 *big.Int) ([32]byte, error) {
	return _MerkelUtils.Contract.DefaultHashes(&_MerkelUtils.CallOpts, arg0)
}

// DefaultHashes is a free data retrieval call binding the contract method 0x48419ad8.
//
// Solidity: function defaultHashes(uint256 ) view returns(bytes32)
func (_MerkelUtils *MerkelUtilsCallerSession) DefaultHashes(arg0 *big.Int) ([32]byte, error) {
	return _MerkelUtils.Contract.DefaultHashes(&_MerkelUtils.CallOpts, arg0)
}

// GetChildren is a free data retrieval call binding the contract method 0xd37684ff.
//
// Solidity: function getChildren(bytes32 _parent) view returns(bytes32, bytes32)
func (_MerkelUtils *MerkelUtilsCaller) GetChildren(opts *bind.CallOpts, _parent [32]byte) ([32]byte, [32]byte, error) {
	var out []interface{}
	err := _MerkelUtils.contract.Call(opts, &out, "getChildren", _parent)

	if err != nil {
		return *new([32]byte), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return out0, out1, err

}

// GetChildren is a free data retrieval call binding the contract method 0xd37684ff.
//
// Solidity: function getChildren(bytes32 _parent) view returns(bytes32, bytes32)
func (_MerkelUtils *MerkelUtilsSession) GetChildren(_parent [32]byte) ([32]byte, [32]byte, error) {
	return _MerkelUtils.Contract.GetChildren(&_MerkelUtils.CallOpts, _parent)
}

// GetChildren is a free data retrieval call binding the contract method 0xd37684ff.
//
// Solidity: function getChildren(bytes32 _parent) view returns(bytes32, bytes32)
func (_MerkelUtils *MerkelUtilsCallerSession) GetChildren(_parent [32]byte) ([32]byte, [32]byte, error) {
	return _MerkelUtils.Contract.GetChildren(&_MerkelUtils.CallOpts, _parent)
}

// GetLeftSiblingKey is a free data retrieval call binding the contract method 0xdf7c7263.
//
// Solidity: function getLeftSiblingKey(bytes32 _parent) pure returns(bytes32)
func (_MerkelUtils *MerkelUtilsCaller) GetLeftSiblingKey(opts *bind.CallOpts, _parent [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _MerkelUtils.contract.Call(opts, &out, "getLeftSiblingKey", _parent)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLeftSiblingKey is a free data retrieval call binding the contract method 0xdf7c7263.
//
// Solidity: function getLeftSiblingKey(bytes32 _parent) pure returns(bytes32)
func (_MerkelUtils *MerkelUtilsSession) GetLeftSiblingKey(_parent [32]byte) ([32]byte, error) {
	return _MerkelUtils.Contract.GetLeftSiblingKey(&_MerkelUtils.CallOpts, _parent)
}

// GetLeftSiblingKey is a free data retrieval call binding the contract method 0xdf7c7263.
//
// Solidity: function getLeftSiblingKey(bytes32 _parent) pure returns(bytes32)
func (_MerkelUtils *MerkelUtilsCallerSession) GetLeftSiblingKey(_parent [32]byte) ([32]byte, error) {
	return _MerkelUtils.Contract.GetLeftSiblingKey(&_MerkelUtils.CallOpts, _parent)
}

// GetMerkleRoot is a free data retrieval call binding the contract method 0x40ff34ef.
//
// Solidity: function getMerkleRoot(bytes[] _dataBlocks) view returns(bytes32)
func (_MerkelUtils *MerkelUtilsCaller) GetMerkleRoot(opts *bind.CallOpts, _dataBlocks [][]byte) ([32]byte, error) {
	var out []interface{}
	err := _MerkelUtils.contract.Call(opts, &out, "getMerkleRoot", _dataBlocks)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetMerkleRoot is a free data retrieval call binding the contract method 0x40ff34ef.
//
// Solidity: function getMerkleRoot(bytes[] _dataBlocks) view returns(bytes32)
func (_MerkelUtils *MerkelUtilsSession) GetMerkleRoot(_dataBlocks [][]byte) ([32]byte, error) {
	return _MerkelUtils.Contract.GetMerkleRoot(&_MerkelUtils.CallOpts, _dataBlocks)
}

// GetMerkleRoot is a free data retrieval call binding the contract method 0x40ff34ef.
//
// Solidity: function getMerkleRoot(bytes[] _dataBlocks) view returns(bytes32)
func (_MerkelUtils *MerkelUtilsCallerSession) GetMerkleRoot(_dataBlocks [][]byte) ([32]byte, error) {
	return _MerkelUtils.Contract.GetMerkleRoot(&_MerkelUtils.CallOpts, _dataBlocks)
}

// GetNthBitFromRight is a free data retrieval call binding the contract method 0xdb0787cb.
//
// Solidity: function getNthBitFromRight(uint256 _intVal, uint256 _index) pure returns(uint8)
func (_MerkelUtils *MerkelUtilsCaller) GetNthBitFromRight(opts *bind.CallOpts, _intVal *big.Int, _index *big.Int) (uint8, error) {
	var out []interface{}
	err := _MerkelUtils.contract.Call(opts, &out, "getNthBitFromRight", _intVal, _index)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetNthBitFromRight is a free data retrieval call binding the contract method 0xdb0787cb.
//
// Solidity: function getNthBitFromRight(uint256 _intVal, uint256 _index) pure returns(uint8)
func (_MerkelUtils *MerkelUtilsSession) GetNthBitFromRight(_intVal *big.Int, _index *big.Int) (uint8, error) {
	return _MerkelUtils.Contract.GetNthBitFromRight(&_MerkelUtils.CallOpts, _intVal, _index)
}

// GetNthBitFromRight is a free data retrieval call binding the contract method 0xdb0787cb.
//
// Solidity: function getNthBitFromRight(uint256 _intVal, uint256 _index) pure returns(uint8)
func (_MerkelUtils *MerkelUtilsCallerSession) GetNthBitFromRight(_intVal *big.Int, _index *big.Int) (uint8, error) {
	return _MerkelUtils.Contract.GetNthBitFromRight(&_MerkelUtils.CallOpts, _intVal, _index)
}

// GetRightSiblingKey is a free data retrieval call binding the contract method 0xe913e47f.
//
// Solidity: function getRightSiblingKey(bytes32 _parent) pure returns(bytes32)
func (_MerkelUtils *MerkelUtilsCaller) GetRightSiblingKey(opts *bind.CallOpts, _parent [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _MerkelUtils.contract.Call(opts, &out, "getRightSiblingKey", _parent)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRightSiblingKey is a free data retrieval call binding the contract method 0xe913e47f.
//
// Solidity: function getRightSiblingKey(bytes32 _parent) pure returns(bytes32)
func (_MerkelUtils *MerkelUtilsSession) GetRightSiblingKey(_parent [32]byte) ([32]byte, error) {
	return _MerkelUtils.Contract.GetRightSiblingKey(&_MerkelUtils.CallOpts, _parent)
}

// GetRightSiblingKey is a free data retrieval call binding the contract method 0xe913e47f.
//
// Solidity: function getRightSiblingKey(bytes32 _parent) pure returns(bytes32)
func (_MerkelUtils *MerkelUtilsCallerSession) GetRightSiblingKey(_parent [32]byte) ([32]byte, error) {
	return _MerkelUtils.Contract.GetRightSiblingKey(&_MerkelUtils.CallOpts, _parent)
}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_MerkelUtils *MerkelUtilsCaller) GetRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MerkelUtils.contract.Call(opts, &out, "getRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_MerkelUtils *MerkelUtilsSession) GetRoot() ([32]byte, error) {
	return _MerkelUtils.Contract.GetRoot(&_MerkelUtils.CallOpts)
}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_MerkelUtils *MerkelUtilsCallerSession) GetRoot() ([32]byte, error) {
	return _MerkelUtils.Contract.GetRoot(&_MerkelUtils.CallOpts)
}

// GetSiblings is a free data retrieval call binding the contract method 0x101b166c.
//
// Solidity: function getSiblings(uint256 _path) view returns(bytes32[])
func (_MerkelUtils *MerkelUtilsCaller) GetSiblings(opts *bind.CallOpts, _path *big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _MerkelUtils.contract.Call(opts, &out, "getSiblings", _path)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetSiblings is a free data retrieval call binding the contract method 0x101b166c.
//
// Solidity: function getSiblings(uint256 _path) view returns(bytes32[])
func (_MerkelUtils *MerkelUtilsSession) GetSiblings(_path *big.Int) ([][32]byte, error) {
	return _MerkelUtils.Contract.GetSiblings(&_MerkelUtils.CallOpts, _path)
}

// GetSiblings is a free data retrieval call binding the contract method 0x101b166c.
//
// Solidity: function getSiblings(uint256 _path) view returns(bytes32[])
func (_MerkelUtils *MerkelUtilsCallerSession) GetSiblings(_path *big.Int) ([][32]byte, error) {
	return _MerkelUtils.Contract.GetSiblings(&_MerkelUtils.CallOpts, _path)
}

// Tree is a free data retrieval call binding the contract method 0xfd54b228.
//
// Solidity: function tree() view returns(bytes32 root, uint256 height)
func (_MerkelUtils *MerkelUtilsCaller) Tree(opts *bind.CallOpts) (struct {
	Root   [32]byte
	Height *big.Int
}, error) {
	var out []interface{}
	err := _MerkelUtils.contract.Call(opts, &out, "tree")

	outstruct := new(struct {
		Root   [32]byte
		Height *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Root = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Height = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Tree is a free data retrieval call binding the contract method 0xfd54b228.
//
// Solidity: function tree() view returns(bytes32 root, uint256 height)
func (_MerkelUtils *MerkelUtilsSession) Tree() (struct {
	Root   [32]byte
	Height *big.Int
}, error) {
	return _MerkelUtils.Contract.Tree(&_MerkelUtils.CallOpts)
}

// Tree is a free data retrieval call binding the contract method 0xfd54b228.
//
// Solidity: function tree() view returns(bytes32 root, uint256 height)
func (_MerkelUtils *MerkelUtilsCallerSession) Tree() (struct {
	Root   [32]byte
	Height *big.Int
}, error) {
	return _MerkelUtils.Contract.Tree(&_MerkelUtils.CallOpts)
}

// Verify is a free data retrieval call binding the contract method 0x30d90a76.
//
// Solidity: function verify(bytes32 _root, bytes _dataBlock, uint256 _path, bytes32[] _siblings) pure returns(bool)
func (_MerkelUtils *MerkelUtilsCaller) Verify(opts *bind.CallOpts, _root [32]byte, _dataBlock []byte, _path *big.Int, _siblings [][32]byte) (bool, error) {
	var out []interface{}
	err := _MerkelUtils.contract.Call(opts, &out, "verify", _root, _dataBlock, _path, _siblings)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Verify is a free data retrieval call binding the contract method 0x30d90a76.
//
// Solidity: function verify(bytes32 _root, bytes _dataBlock, uint256 _path, bytes32[] _siblings) pure returns(bool)
func (_MerkelUtils *MerkelUtilsSession) Verify(_root [32]byte, _dataBlock []byte, _path *big.Int, _siblings [][32]byte) (bool, error) {
	return _MerkelUtils.Contract.Verify(&_MerkelUtils.CallOpts, _root, _dataBlock, _path, _siblings)
}

// Verify is a free data retrieval call binding the contract method 0x30d90a76.
//
// Solidity: function verify(bytes32 _root, bytes _dataBlock, uint256 _path, bytes32[] _siblings) pure returns(bool)
func (_MerkelUtils *MerkelUtilsCallerSession) Verify(_root [32]byte, _dataBlock []byte, _path *big.Int, _siblings [][32]byte) (bool, error) {
	return _MerkelUtils.Contract.Verify(&_MerkelUtils.CallOpts, _root, _dataBlock, _path, _siblings)
}

// SetMerkleRootAndHeight is a paid mutator transaction binding the contract method 0x5c22b6d9.
//
// Solidity: function setMerkleRootAndHeight(bytes32 _root, uint256 _height) returns()
func (_MerkelUtils *MerkelUtilsTransactor) SetMerkleRootAndHeight(opts *bind.TransactOpts, _root [32]byte, _height *big.Int) (*types.Transaction, error) {
	return _MerkelUtils.contract.Transact(opts, "setMerkleRootAndHeight", _root, _height)
}

// SetMerkleRootAndHeight is a paid mutator transaction binding the contract method 0x5c22b6d9.
//
// Solidity: function setMerkleRootAndHeight(bytes32 _root, uint256 _height) returns()
func (_MerkelUtils *MerkelUtilsSession) SetMerkleRootAndHeight(_root [32]byte, _height *big.Int) (*types.Transaction, error) {
	return _MerkelUtils.Contract.SetMerkleRootAndHeight(&_MerkelUtils.TransactOpts, _root, _height)
}

// SetMerkleRootAndHeight is a paid mutator transaction binding the contract method 0x5c22b6d9.
//
// Solidity: function setMerkleRootAndHeight(bytes32 _root, uint256 _height) returns()
func (_MerkelUtils *MerkelUtilsTransactorSession) SetMerkleRootAndHeight(_root [32]byte, _height *big.Int) (*types.Transaction, error) {
	return _MerkelUtils.Contract.SetMerkleRootAndHeight(&_MerkelUtils.TransactOpts, _root, _height)
}

// Store is a paid mutator transaction binding the contract method 0x158933ad.
//
// Solidity: function store(bytes _dataBlock, uint256 _path, bytes32[] _siblings) returns()
func (_MerkelUtils *MerkelUtilsTransactor) Store(opts *bind.TransactOpts, _dataBlock []byte, _path *big.Int, _siblings [][32]byte) (*types.Transaction, error) {
	return _MerkelUtils.contract.Transact(opts, "store", _dataBlock, _path, _siblings)
}

// Store is a paid mutator transaction binding the contract method 0x158933ad.
//
// Solidity: function store(bytes _dataBlock, uint256 _path, bytes32[] _siblings) returns()
func (_MerkelUtils *MerkelUtilsSession) Store(_dataBlock []byte, _path *big.Int, _siblings [][32]byte) (*types.Transaction, error) {
	return _MerkelUtils.Contract.Store(&_MerkelUtils.TransactOpts, _dataBlock, _path, _siblings)
}

// Store is a paid mutator transaction binding the contract method 0x158933ad.
//
// Solidity: function store(bytes _dataBlock, uint256 _path, bytes32[] _siblings) returns()
func (_MerkelUtils *MerkelUtilsTransactorSession) Store(_dataBlock []byte, _path *big.Int, _siblings [][32]byte) (*types.Transaction, error) {
	return _MerkelUtils.Contract.Store(&_MerkelUtils.TransactOpts, _dataBlock, _path, _siblings)
}

// StoreLeaf is a paid mutator transaction binding the contract method 0x9c0de520.
//
// Solidity: function storeLeaf(bytes32 _leaf, uint256 _path, bytes32[] _siblings) returns()
func (_MerkelUtils *MerkelUtilsTransactor) StoreLeaf(opts *bind.TransactOpts, _leaf [32]byte, _path *big.Int, _siblings [][32]byte) (*types.Transaction, error) {
	return _MerkelUtils.contract.Transact(opts, "storeLeaf", _leaf, _path, _siblings)
}

// StoreLeaf is a paid mutator transaction binding the contract method 0x9c0de520.
//
// Solidity: function storeLeaf(bytes32 _leaf, uint256 _path, bytes32[] _siblings) returns()
func (_MerkelUtils *MerkelUtilsSession) StoreLeaf(_leaf [32]byte, _path *big.Int, _siblings [][32]byte) (*types.Transaction, error) {
	return _MerkelUtils.Contract.StoreLeaf(&_MerkelUtils.TransactOpts, _leaf, _path, _siblings)
}

// StoreLeaf is a paid mutator transaction binding the contract method 0x9c0de520.
//
// Solidity: function storeLeaf(bytes32 _leaf, uint256 _path, bytes32[] _siblings) returns()
func (_MerkelUtils *MerkelUtilsTransactorSession) StoreLeaf(_leaf [32]byte, _path *big.Int, _siblings [][32]byte) (*types.Transaction, error) {
	return _MerkelUtils.Contract.StoreLeaf(&_MerkelUtils.TransactOpts, _leaf, _path, _siblings)
}

// StoreNode is a paid mutator transaction binding the contract method 0x63327f89.
//
// Solidity: function storeNode(bytes32 _parent, bytes32 _leftChild, bytes32 _rightChild) returns()
func (_MerkelUtils *MerkelUtilsTransactor) StoreNode(opts *bind.TransactOpts, _parent [32]byte, _leftChild [32]byte, _rightChild [32]byte) (*types.Transaction, error) {
	return _MerkelUtils.contract.Transact(opts, "storeNode", _parent, _leftChild, _rightChild)
}

// StoreNode is a paid mutator transaction binding the contract method 0x63327f89.
//
// Solidity: function storeNode(bytes32 _parent, bytes32 _leftChild, bytes32 _rightChild) returns()
func (_MerkelUtils *MerkelUtilsSession) StoreNode(_parent [32]byte, _leftChild [32]byte, _rightChild [32]byte) (*types.Transaction, error) {
	return _MerkelUtils.Contract.StoreNode(&_MerkelUtils.TransactOpts, _parent, _leftChild, _rightChild)
}

// StoreNode is a paid mutator transaction binding the contract method 0x63327f89.
//
// Solidity: function storeNode(bytes32 _parent, bytes32 _leftChild, bytes32 _rightChild) returns()
func (_MerkelUtils *MerkelUtilsTransactorSession) StoreNode(_parent [32]byte, _leftChild [32]byte, _rightChild [32]byte) (*types.Transaction, error) {
	return _MerkelUtils.Contract.StoreNode(&_MerkelUtils.TransactOpts, _parent, _leftChild, _rightChild)
}

// Update is a paid mutator transaction binding the contract method 0xc3b45234.
//
// Solidity: function update(bytes _dataBlock, uint256 _path) returns()
func (_MerkelUtils *MerkelUtilsTransactor) Update(opts *bind.TransactOpts, _dataBlock []byte, _path *big.Int) (*types.Transaction, error) {
	return _MerkelUtils.contract.Transact(opts, "update", _dataBlock, _path)
}

// Update is a paid mutator transaction binding the contract method 0xc3b45234.
//
// Solidity: function update(bytes _dataBlock, uint256 _path) returns()
func (_MerkelUtils *MerkelUtilsSession) Update(_dataBlock []byte, _path *big.Int) (*types.Transaction, error) {
	return _MerkelUtils.Contract.Update(&_MerkelUtils.TransactOpts, _dataBlock, _path)
}

// Update is a paid mutator transaction binding the contract method 0xc3b45234.
//
// Solidity: function update(bytes _dataBlock, uint256 _path) returns()
func (_MerkelUtils *MerkelUtilsTransactorSession) Update(_dataBlock []byte, _path *big.Int) (*types.Transaction, error) {
	return _MerkelUtils.Contract.Update(&_MerkelUtils.TransactOpts, _dataBlock, _path)
}

// UpdateLeaf is a paid mutator transaction binding the contract method 0x272684b5.
//
// Solidity: function updateLeaf(bytes32 _leaf, uint256 _path) returns()
func (_MerkelUtils *MerkelUtilsTransactor) UpdateLeaf(opts *bind.TransactOpts, _leaf [32]byte, _path *big.Int) (*types.Transaction, error) {
	return _MerkelUtils.contract.Transact(opts, "updateLeaf", _leaf, _path)
}

// UpdateLeaf is a paid mutator transaction binding the contract method 0x272684b5.
//
// Solidity: function updateLeaf(bytes32 _leaf, uint256 _path) returns()
func (_MerkelUtils *MerkelUtilsSession) UpdateLeaf(_leaf [32]byte, _path *big.Int) (*types.Transaction, error) {
	return _MerkelUtils.Contract.UpdateLeaf(&_MerkelUtils.TransactOpts, _leaf, _path)
}

// UpdateLeaf is a paid mutator transaction binding the contract method 0x272684b5.
//
// Solidity: function updateLeaf(bytes32 _leaf, uint256 _path) returns()
func (_MerkelUtils *MerkelUtilsTransactorSession) UpdateLeaf(_leaf [32]byte, _path *big.Int) (*types.Transaction, error) {
	return _MerkelUtils.Contract.UpdateLeaf(&_MerkelUtils.TransactOpts, _leaf, _path)
}

// VerifyAndStore is a paid mutator transaction binding the contract method 0x4359356d.
//
// Solidity: function verifyAndStore(bytes _dataBlock, uint256 _path, bytes32[] _siblings) returns()
func (_MerkelUtils *MerkelUtilsTransactor) VerifyAndStore(opts *bind.TransactOpts, _dataBlock []byte, _path *big.Int, _siblings [][32]byte) (*types.Transaction, error) {
	return _MerkelUtils.contract.Transact(opts, "verifyAndStore", _dataBlock, _path, _siblings)
}

// VerifyAndStore is a paid mutator transaction binding the contract method 0x4359356d.
//
// Solidity: function verifyAndStore(bytes _dataBlock, uint256 _path, bytes32[] _siblings) returns()
func (_MerkelUtils *MerkelUtilsSession) VerifyAndStore(_dataBlock []byte, _path *big.Int, _siblings [][32]byte) (*types.Transaction, error) {
	return _MerkelUtils.Contract.VerifyAndStore(&_MerkelUtils.TransactOpts, _dataBlock, _path, _siblings)
}

// VerifyAndStore is a paid mutator transaction binding the contract method 0x4359356d.
//
// Solidity: function verifyAndStore(bytes _dataBlock, uint256 _path, bytes32[] _siblings) returns()
func (_MerkelUtils *MerkelUtilsTransactorSession) VerifyAndStore(_dataBlock []byte, _path *big.Int, _siblings [][32]byte) (*types.Transaction, error) {
	return _MerkelUtils.Contract.VerifyAndStore(&_MerkelUtils.TransactOpts, _dataBlock, _path, _siblings)
}

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

// MerkleUtilsMetaData contains all meta data concerning the MerkleUtils contract.
var MerkleUtilsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"defaultHashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tree\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_dataBlocks\",\"type\":\"bytes[]\"}],\"name\":\"getMerkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_dataBlock\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_path\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_siblings\",\"type\":\"bytes32[]\"}],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_dataBlock\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_path\",\"type\":\"uint256\"}],\"name\":\"update\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_leaf\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_path\",\"type\":\"uint256\"}],\"name\":\"updateLeaf\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_dataBlock\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_path\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_siblings\",\"type\":\"bytes32[]\"}],\"name\":\"verifyAndStore\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_dataBlock\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_path\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_siblings\",\"type\":\"bytes32[]\"}],\"name\":\"store\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_leaf\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_path\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_siblings\",\"type\":\"bytes32[]\"}],\"name\":\"storeLeaf\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_path\",\"type\":\"uint256\"}],\"name\":\"getSiblings\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_height\",\"type\":\"uint256\"}],\"name\":\"setMerkleRootAndHeight\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_parent\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_leftChild\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_rightChild\",\"type\":\"bytes32\"}],\"name\":\"storeNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_intVal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getNthBitFromRight\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_parent\",\"type\":\"bytes32\"}],\"name\":\"getChildren\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_parent\",\"type\":\"bytes32\"}],\"name\":\"getLeftSiblingKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_parent\",\"type\":\"bytes32\"}],\"name\":\"getRightSiblingKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// MerkleUtilsABI is the input ABI used to generate the binding from.
// Deprecated: Use MerkleUtilsMetaData.ABI instead.
var MerkleUtilsABI = MerkleUtilsMetaData.ABI

// MerkleUtils is an auto generated Go binding around an Ethereum contract.
type MerkleUtils struct {
	MerkleUtilsCaller     // Read-only binding to the contract
	MerkleUtilsTransactor // Write-only binding to the contract
	MerkleUtilsFilterer   // Log filterer for contract events
}

// MerkleUtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type MerkleUtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleUtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MerkleUtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleUtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MerkleUtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleUtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MerkleUtilsSession struct {
	Contract     *MerkleUtils      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MerkleUtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MerkleUtilsCallerSession struct {
	Contract *MerkleUtilsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// MerkleUtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MerkleUtilsTransactorSession struct {
	Contract     *MerkleUtilsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MerkleUtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type MerkleUtilsRaw struct {
	Contract *MerkleUtils // Generic contract binding to access the raw methods on
}

// MerkleUtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MerkleUtilsCallerRaw struct {
	Contract *MerkleUtilsCaller // Generic read-only contract binding to access the raw methods on
}

// MerkleUtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MerkleUtilsTransactorRaw struct {
	Contract *MerkleUtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMerkleUtils creates a new instance of MerkleUtils, bound to a specific deployed contract.
func NewMerkleUtils(address common.Address, backend bind.ContractBackend) (*MerkleUtils, error) {
	contract, err := bindMerkleUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MerkleUtils{MerkleUtilsCaller: MerkleUtilsCaller{contract: contract}, MerkleUtilsTransactor: MerkleUtilsTransactor{contract: contract}, MerkleUtilsFilterer: MerkleUtilsFilterer{contract: contract}}, nil
}

// NewMerkleUtilsCaller creates a new read-only instance of MerkleUtils, bound to a specific deployed contract.
func NewMerkleUtilsCaller(address common.Address, caller bind.ContractCaller) (*MerkleUtilsCaller, error) {
	contract, err := bindMerkleUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleUtilsCaller{contract: contract}, nil
}

// NewMerkleUtilsTransactor creates a new write-only instance of MerkleUtils, bound to a specific deployed contract.
func NewMerkleUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*MerkleUtilsTransactor, error) {
	contract, err := bindMerkleUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleUtilsTransactor{contract: contract}, nil
}

// NewMerkleUtilsFilterer creates a new log filterer instance of MerkleUtils, bound to a specific deployed contract.
func NewMerkleUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*MerkleUtilsFilterer, error) {
	contract, err := bindMerkleUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MerkleUtilsFilterer{contract: contract}, nil
}

// bindMerkleUtils binds a generic wrapper to an already deployed contract.
func bindMerkleUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MerkleUtilsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerkleUtils *MerkleUtilsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MerkleUtils.Contract.MerkleUtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerkleUtils *MerkleUtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerkleUtils.Contract.MerkleUtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerkleUtils *MerkleUtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerkleUtils.Contract.MerkleUtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerkleUtils *MerkleUtilsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MerkleUtils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerkleUtils *MerkleUtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerkleUtils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerkleUtils *MerkleUtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerkleUtils.Contract.contract.Transact(opts, method, params...)
}

// DefaultHashes is a free data retrieval call binding the contract method 0x48419ad8.
//
// Solidity: function defaultHashes(uint256 ) view returns(bytes32)
func (_MerkleUtils *MerkleUtilsCaller) DefaultHashes(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _MerkleUtils.contract.Call(opts, &out, "defaultHashes", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DefaultHashes is a free data retrieval call binding the contract method 0x48419ad8.
//
// Solidity: function defaultHashes(uint256 ) view returns(bytes32)
func (_MerkleUtils *MerkleUtilsSession) DefaultHashes(arg0 *big.Int) ([32]byte, error) {
	return _MerkleUtils.Contract.DefaultHashes(&_MerkleUtils.CallOpts, arg0)
}

// DefaultHashes is a free data retrieval call binding the contract method 0x48419ad8.
//
// Solidity: function defaultHashes(uint256 ) view returns(bytes32)
func (_MerkleUtils *MerkleUtilsCallerSession) DefaultHashes(arg0 *big.Int) ([32]byte, error) {
	return _MerkleUtils.Contract.DefaultHashes(&_MerkleUtils.CallOpts, arg0)
}

// GetChildren is a free data retrieval call binding the contract method 0xd37684ff.
//
// Solidity: function getChildren(bytes32 _parent) view returns(bytes32, bytes32)
func (_MerkleUtils *MerkleUtilsCaller) GetChildren(opts *bind.CallOpts, _parent [32]byte) ([32]byte, [32]byte, error) {
	var out []interface{}
	err := _MerkleUtils.contract.Call(opts, &out, "getChildren", _parent)

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
func (_MerkleUtils *MerkleUtilsSession) GetChildren(_parent [32]byte) ([32]byte, [32]byte, error) {
	return _MerkleUtils.Contract.GetChildren(&_MerkleUtils.CallOpts, _parent)
}

// GetChildren is a free data retrieval call binding the contract method 0xd37684ff.
//
// Solidity: function getChildren(bytes32 _parent) view returns(bytes32, bytes32)
func (_MerkleUtils *MerkleUtilsCallerSession) GetChildren(_parent [32]byte) ([32]byte, [32]byte, error) {
	return _MerkleUtils.Contract.GetChildren(&_MerkleUtils.CallOpts, _parent)
}

// GetLeftSiblingKey is a free data retrieval call binding the contract method 0xdf7c7263.
//
// Solidity: function getLeftSiblingKey(bytes32 _parent) pure returns(bytes32)
func (_MerkleUtils *MerkleUtilsCaller) GetLeftSiblingKey(opts *bind.CallOpts, _parent [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _MerkleUtils.contract.Call(opts, &out, "getLeftSiblingKey", _parent)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLeftSiblingKey is a free data retrieval call binding the contract method 0xdf7c7263.
//
// Solidity: function getLeftSiblingKey(bytes32 _parent) pure returns(bytes32)
func (_MerkleUtils *MerkleUtilsSession) GetLeftSiblingKey(_parent [32]byte) ([32]byte, error) {
	return _MerkleUtils.Contract.GetLeftSiblingKey(&_MerkleUtils.CallOpts, _parent)
}

// GetLeftSiblingKey is a free data retrieval call binding the contract method 0xdf7c7263.
//
// Solidity: function getLeftSiblingKey(bytes32 _parent) pure returns(bytes32)
func (_MerkleUtils *MerkleUtilsCallerSession) GetLeftSiblingKey(_parent [32]byte) ([32]byte, error) {
	return _MerkleUtils.Contract.GetLeftSiblingKey(&_MerkleUtils.CallOpts, _parent)
}

// GetMerkleRoot is a free data retrieval call binding the contract method 0x40ff34ef.
//
// Solidity: function getMerkleRoot(bytes[] _dataBlocks) view returns(bytes32)
func (_MerkleUtils *MerkleUtilsCaller) GetMerkleRoot(opts *bind.CallOpts, _dataBlocks [][]byte) ([32]byte, error) {
	var out []interface{}
	err := _MerkleUtils.contract.Call(opts, &out, "getMerkleRoot", _dataBlocks)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetMerkleRoot is a free data retrieval call binding the contract method 0x40ff34ef.
//
// Solidity: function getMerkleRoot(bytes[] _dataBlocks) view returns(bytes32)
func (_MerkleUtils *MerkleUtilsSession) GetMerkleRoot(_dataBlocks [][]byte) ([32]byte, error) {
	return _MerkleUtils.Contract.GetMerkleRoot(&_MerkleUtils.CallOpts, _dataBlocks)
}

// GetMerkleRoot is a free data retrieval call binding the contract method 0x40ff34ef.
//
// Solidity: function getMerkleRoot(bytes[] _dataBlocks) view returns(bytes32)
func (_MerkleUtils *MerkleUtilsCallerSession) GetMerkleRoot(_dataBlocks [][]byte) ([32]byte, error) {
	return _MerkleUtils.Contract.GetMerkleRoot(&_MerkleUtils.CallOpts, _dataBlocks)
}

// GetNthBitFromRight is a free data retrieval call binding the contract method 0xdb0787cb.
//
// Solidity: function getNthBitFromRight(uint256 _intVal, uint256 _index) pure returns(uint8)
func (_MerkleUtils *MerkleUtilsCaller) GetNthBitFromRight(opts *bind.CallOpts, _intVal *big.Int, _index *big.Int) (uint8, error) {
	var out []interface{}
	err := _MerkleUtils.contract.Call(opts, &out, "getNthBitFromRight", _intVal, _index)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetNthBitFromRight is a free data retrieval call binding the contract method 0xdb0787cb.
//
// Solidity: function getNthBitFromRight(uint256 _intVal, uint256 _index) pure returns(uint8)
func (_MerkleUtils *MerkleUtilsSession) GetNthBitFromRight(_intVal *big.Int, _index *big.Int) (uint8, error) {
	return _MerkleUtils.Contract.GetNthBitFromRight(&_MerkleUtils.CallOpts, _intVal, _index)
}

// GetNthBitFromRight is a free data retrieval call binding the contract method 0xdb0787cb.
//
// Solidity: function getNthBitFromRight(uint256 _intVal, uint256 _index) pure returns(uint8)
func (_MerkleUtils *MerkleUtilsCallerSession) GetNthBitFromRight(_intVal *big.Int, _index *big.Int) (uint8, error) {
	return _MerkleUtils.Contract.GetNthBitFromRight(&_MerkleUtils.CallOpts, _intVal, _index)
}

// GetRightSiblingKey is a free data retrieval call binding the contract method 0xe913e47f.
//
// Solidity: function getRightSiblingKey(bytes32 _parent) pure returns(bytes32)
func (_MerkleUtils *MerkleUtilsCaller) GetRightSiblingKey(opts *bind.CallOpts, _parent [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _MerkleUtils.contract.Call(opts, &out, "getRightSiblingKey", _parent)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRightSiblingKey is a free data retrieval call binding the contract method 0xe913e47f.
//
// Solidity: function getRightSiblingKey(bytes32 _parent) pure returns(bytes32)
func (_MerkleUtils *MerkleUtilsSession) GetRightSiblingKey(_parent [32]byte) ([32]byte, error) {
	return _MerkleUtils.Contract.GetRightSiblingKey(&_MerkleUtils.CallOpts, _parent)
}

// GetRightSiblingKey is a free data retrieval call binding the contract method 0xe913e47f.
//
// Solidity: function getRightSiblingKey(bytes32 _parent) pure returns(bytes32)
func (_MerkleUtils *MerkleUtilsCallerSession) GetRightSiblingKey(_parent [32]byte) ([32]byte, error) {
	return _MerkleUtils.Contract.GetRightSiblingKey(&_MerkleUtils.CallOpts, _parent)
}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_MerkleUtils *MerkleUtilsCaller) GetRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MerkleUtils.contract.Call(opts, &out, "getRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_MerkleUtils *MerkleUtilsSession) GetRoot() ([32]byte, error) {
	return _MerkleUtils.Contract.GetRoot(&_MerkleUtils.CallOpts)
}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_MerkleUtils *MerkleUtilsCallerSession) GetRoot() ([32]byte, error) {
	return _MerkleUtils.Contract.GetRoot(&_MerkleUtils.CallOpts)
}

// GetSiblings is a free data retrieval call binding the contract method 0x101b166c.
//
// Solidity: function getSiblings(uint256 _path) view returns(bytes32[])
func (_MerkleUtils *MerkleUtilsCaller) GetSiblings(opts *bind.CallOpts, _path *big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _MerkleUtils.contract.Call(opts, &out, "getSiblings", _path)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetSiblings is a free data retrieval call binding the contract method 0x101b166c.
//
// Solidity: function getSiblings(uint256 _path) view returns(bytes32[])
func (_MerkleUtils *MerkleUtilsSession) GetSiblings(_path *big.Int) ([][32]byte, error) {
	return _MerkleUtils.Contract.GetSiblings(&_MerkleUtils.CallOpts, _path)
}

// GetSiblings is a free data retrieval call binding the contract method 0x101b166c.
//
// Solidity: function getSiblings(uint256 _path) view returns(bytes32[])
func (_MerkleUtils *MerkleUtilsCallerSession) GetSiblings(_path *big.Int) ([][32]byte, error) {
	return _MerkleUtils.Contract.GetSiblings(&_MerkleUtils.CallOpts, _path)
}

// Tree is a free data retrieval call binding the contract method 0xfd54b228.
//
// Solidity: function tree() view returns(bytes32 root, uint256 height)
func (_MerkleUtils *MerkleUtilsCaller) Tree(opts *bind.CallOpts) (struct {
	Root   [32]byte
	Height *big.Int
}, error) {
	var out []interface{}
	err := _MerkleUtils.contract.Call(opts, &out, "tree")

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
func (_MerkleUtils *MerkleUtilsSession) Tree() (struct {
	Root   [32]byte
	Height *big.Int
}, error) {
	return _MerkleUtils.Contract.Tree(&_MerkleUtils.CallOpts)
}

// Tree is a free data retrieval call binding the contract method 0xfd54b228.
//
// Solidity: function tree() view returns(bytes32 root, uint256 height)
func (_MerkleUtils *MerkleUtilsCallerSession) Tree() (struct {
	Root   [32]byte
	Height *big.Int
}, error) {
	return _MerkleUtils.Contract.Tree(&_MerkleUtils.CallOpts)
}

// Verify is a free data retrieval call binding the contract method 0x30d90a76.
//
// Solidity: function verify(bytes32 _root, bytes _dataBlock, uint256 _path, bytes32[] _siblings) pure returns(bool)
func (_MerkleUtils *MerkleUtilsCaller) Verify(opts *bind.CallOpts, _root [32]byte, _dataBlock []byte, _path *big.Int, _siblings [][32]byte) (bool, error) {
	var out []interface{}
	err := _MerkleUtils.contract.Call(opts, &out, "verify", _root, _dataBlock, _path, _siblings)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Verify is a free data retrieval call binding the contract method 0x30d90a76.
//
// Solidity: function verify(bytes32 _root, bytes _dataBlock, uint256 _path, bytes32[] _siblings) pure returns(bool)
func (_MerkleUtils *MerkleUtilsSession) Verify(_root [32]byte, _dataBlock []byte, _path *big.Int, _siblings [][32]byte) (bool, error) {
	return _MerkleUtils.Contract.Verify(&_MerkleUtils.CallOpts, _root, _dataBlock, _path, _siblings)
}

// Verify is a free data retrieval call binding the contract method 0x30d90a76.
//
// Solidity: function verify(bytes32 _root, bytes _dataBlock, uint256 _path, bytes32[] _siblings) pure returns(bool)
func (_MerkleUtils *MerkleUtilsCallerSession) Verify(_root [32]byte, _dataBlock []byte, _path *big.Int, _siblings [][32]byte) (bool, error) {
	return _MerkleUtils.Contract.Verify(&_MerkleUtils.CallOpts, _root, _dataBlock, _path, _siblings)
}

// SetMerkleRootAndHeight is a paid mutator transaction binding the contract method 0x5c22b6d9.
//
// Solidity: function setMerkleRootAndHeight(bytes32 _root, uint256 _height) returns()
func (_MerkleUtils *MerkleUtilsTransactor) SetMerkleRootAndHeight(opts *bind.TransactOpts, _root [32]byte, _height *big.Int) (*types.Transaction, error) {
	return _MerkleUtils.contract.Transact(opts, "setMerkleRootAndHeight", _root, _height)
}

// SetMerkleRootAndHeight is a paid mutator transaction binding the contract method 0x5c22b6d9.
//
// Solidity: function setMerkleRootAndHeight(bytes32 _root, uint256 _height) returns()
func (_MerkleUtils *MerkleUtilsSession) SetMerkleRootAndHeight(_root [32]byte, _height *big.Int) (*types.Transaction, error) {
	return _MerkleUtils.Contract.SetMerkleRootAndHeight(&_MerkleUtils.TransactOpts, _root, _height)
}

// SetMerkleRootAndHeight is a paid mutator transaction binding the contract method 0x5c22b6d9.
//
// Solidity: function setMerkleRootAndHeight(bytes32 _root, uint256 _height) returns()
func (_MerkleUtils *MerkleUtilsTransactorSession) SetMerkleRootAndHeight(_root [32]byte, _height *big.Int) (*types.Transaction, error) {
	return _MerkleUtils.Contract.SetMerkleRootAndHeight(&_MerkleUtils.TransactOpts, _root, _height)
}

// Store is a paid mutator transaction binding the contract method 0x158933ad.
//
// Solidity: function store(bytes _dataBlock, uint256 _path, bytes32[] _siblings) returns()
func (_MerkleUtils *MerkleUtilsTransactor) Store(opts *bind.TransactOpts, _dataBlock []byte, _path *big.Int, _siblings [][32]byte) (*types.Transaction, error) {
	return _MerkleUtils.contract.Transact(opts, "store", _dataBlock, _path, _siblings)
}

// Store is a paid mutator transaction binding the contract method 0x158933ad.
//
// Solidity: function store(bytes _dataBlock, uint256 _path, bytes32[] _siblings) returns()
func (_MerkleUtils *MerkleUtilsSession) Store(_dataBlock []byte, _path *big.Int, _siblings [][32]byte) (*types.Transaction, error) {
	return _MerkleUtils.Contract.Store(&_MerkleUtils.TransactOpts, _dataBlock, _path, _siblings)
}

// Store is a paid mutator transaction binding the contract method 0x158933ad.
//
// Solidity: function store(bytes _dataBlock, uint256 _path, bytes32[] _siblings) returns()
func (_MerkleUtils *MerkleUtilsTransactorSession) Store(_dataBlock []byte, _path *big.Int, _siblings [][32]byte) (*types.Transaction, error) {
	return _MerkleUtils.Contract.Store(&_MerkleUtils.TransactOpts, _dataBlock, _path, _siblings)
}

// StoreLeaf is a paid mutator transaction binding the contract method 0x9c0de520.
//
// Solidity: function storeLeaf(bytes32 _leaf, uint256 _path, bytes32[] _siblings) returns()
func (_MerkleUtils *MerkleUtilsTransactor) StoreLeaf(opts *bind.TransactOpts, _leaf [32]byte, _path *big.Int, _siblings [][32]byte) (*types.Transaction, error) {
	return _MerkleUtils.contract.Transact(opts, "storeLeaf", _leaf, _path, _siblings)
}

// StoreLeaf is a paid mutator transaction binding the contract method 0x9c0de520.
//
// Solidity: function storeLeaf(bytes32 _leaf, uint256 _path, bytes32[] _siblings) returns()
func (_MerkleUtils *MerkleUtilsSession) StoreLeaf(_leaf [32]byte, _path *big.Int, _siblings [][32]byte) (*types.Transaction, error) {
	return _MerkleUtils.Contract.StoreLeaf(&_MerkleUtils.TransactOpts, _leaf, _path, _siblings)
}

// StoreLeaf is a paid mutator transaction binding the contract method 0x9c0de520.
//
// Solidity: function storeLeaf(bytes32 _leaf, uint256 _path, bytes32[] _siblings) returns()
func (_MerkleUtils *MerkleUtilsTransactorSession) StoreLeaf(_leaf [32]byte, _path *big.Int, _siblings [][32]byte) (*types.Transaction, error) {
	return _MerkleUtils.Contract.StoreLeaf(&_MerkleUtils.TransactOpts, _leaf, _path, _siblings)
}

// StoreNode is a paid mutator transaction binding the contract method 0x63327f89.
//
// Solidity: function storeNode(bytes32 _parent, bytes32 _leftChild, bytes32 _rightChild) returns()
func (_MerkleUtils *MerkleUtilsTransactor) StoreNode(opts *bind.TransactOpts, _parent [32]byte, _leftChild [32]byte, _rightChild [32]byte) (*types.Transaction, error) {
	return _MerkleUtils.contract.Transact(opts, "storeNode", _parent, _leftChild, _rightChild)
}

// StoreNode is a paid mutator transaction binding the contract method 0x63327f89.
//
// Solidity: function storeNode(bytes32 _parent, bytes32 _leftChild, bytes32 _rightChild) returns()
func (_MerkleUtils *MerkleUtilsSession) StoreNode(_parent [32]byte, _leftChild [32]byte, _rightChild [32]byte) (*types.Transaction, error) {
	return _MerkleUtils.Contract.StoreNode(&_MerkleUtils.TransactOpts, _parent, _leftChild, _rightChild)
}

// StoreNode is a paid mutator transaction binding the contract method 0x63327f89.
//
// Solidity: function storeNode(bytes32 _parent, bytes32 _leftChild, bytes32 _rightChild) returns()
func (_MerkleUtils *MerkleUtilsTransactorSession) StoreNode(_parent [32]byte, _leftChild [32]byte, _rightChild [32]byte) (*types.Transaction, error) {
	return _MerkleUtils.Contract.StoreNode(&_MerkleUtils.TransactOpts, _parent, _leftChild, _rightChild)
}

// Update is a paid mutator transaction binding the contract method 0xc3b45234.
//
// Solidity: function update(bytes _dataBlock, uint256 _path) returns()
func (_MerkleUtils *MerkleUtilsTransactor) Update(opts *bind.TransactOpts, _dataBlock []byte, _path *big.Int) (*types.Transaction, error) {
	return _MerkleUtils.contract.Transact(opts, "update", _dataBlock, _path)
}

// Update is a paid mutator transaction binding the contract method 0xc3b45234.
//
// Solidity: function update(bytes _dataBlock, uint256 _path) returns()
func (_MerkleUtils *MerkleUtilsSession) Update(_dataBlock []byte, _path *big.Int) (*types.Transaction, error) {
	return _MerkleUtils.Contract.Update(&_MerkleUtils.TransactOpts, _dataBlock, _path)
}

// Update is a paid mutator transaction binding the contract method 0xc3b45234.
//
// Solidity: function update(bytes _dataBlock, uint256 _path) returns()
func (_MerkleUtils *MerkleUtilsTransactorSession) Update(_dataBlock []byte, _path *big.Int) (*types.Transaction, error) {
	return _MerkleUtils.Contract.Update(&_MerkleUtils.TransactOpts, _dataBlock, _path)
}

// UpdateLeaf is a paid mutator transaction binding the contract method 0x272684b5.
//
// Solidity: function updateLeaf(bytes32 _leaf, uint256 _path) returns()
func (_MerkleUtils *MerkleUtilsTransactor) UpdateLeaf(opts *bind.TransactOpts, _leaf [32]byte, _path *big.Int) (*types.Transaction, error) {
	return _MerkleUtils.contract.Transact(opts, "updateLeaf", _leaf, _path)
}

// UpdateLeaf is a paid mutator transaction binding the contract method 0x272684b5.
//
// Solidity: function updateLeaf(bytes32 _leaf, uint256 _path) returns()
func (_MerkleUtils *MerkleUtilsSession) UpdateLeaf(_leaf [32]byte, _path *big.Int) (*types.Transaction, error) {
	return _MerkleUtils.Contract.UpdateLeaf(&_MerkleUtils.TransactOpts, _leaf, _path)
}

// UpdateLeaf is a paid mutator transaction binding the contract method 0x272684b5.
//
// Solidity: function updateLeaf(bytes32 _leaf, uint256 _path) returns()
func (_MerkleUtils *MerkleUtilsTransactorSession) UpdateLeaf(_leaf [32]byte, _path *big.Int) (*types.Transaction, error) {
	return _MerkleUtils.Contract.UpdateLeaf(&_MerkleUtils.TransactOpts, _leaf, _path)
}

// VerifyAndStore is a paid mutator transaction binding the contract method 0x4359356d.
//
// Solidity: function verifyAndStore(bytes _dataBlock, uint256 _path, bytes32[] _siblings) returns()
func (_MerkleUtils *MerkleUtilsTransactor) VerifyAndStore(opts *bind.TransactOpts, _dataBlock []byte, _path *big.Int, _siblings [][32]byte) (*types.Transaction, error) {
	return _MerkleUtils.contract.Transact(opts, "verifyAndStore", _dataBlock, _path, _siblings)
}

// VerifyAndStore is a paid mutator transaction binding the contract method 0x4359356d.
//
// Solidity: function verifyAndStore(bytes _dataBlock, uint256 _path, bytes32[] _siblings) returns()
func (_MerkleUtils *MerkleUtilsSession) VerifyAndStore(_dataBlock []byte, _path *big.Int, _siblings [][32]byte) (*types.Transaction, error) {
	return _MerkleUtils.Contract.VerifyAndStore(&_MerkleUtils.TransactOpts, _dataBlock, _path, _siblings)
}

// VerifyAndStore is a paid mutator transaction binding the contract method 0x4359356d.
//
// Solidity: function verifyAndStore(bytes _dataBlock, uint256 _path, bytes32[] _siblings) returns()
func (_MerkleUtils *MerkleUtilsTransactorSession) VerifyAndStore(_dataBlock []byte, _path *big.Int, _siblings [][32]byte) (*types.Transaction, error) {
	return _MerkleUtils.Contract.VerifyAndStore(&_MerkleUtils.TransactOpts, _dataBlock, _path, _siblings)
}

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

// MMRVerifierMMRConfig is an auto generated low-level Go binding around an user-defined struct.
type MMRVerifierMMRConfig struct {
	LeafNum *big.Int
}

// MMRVerifierProofBlock is an auto generated low-level Go binding around an user-defined struct.
type MMRVerifierProofBlock struct {
	Number     uint64
	AggrWeight *big.Int
}

// MMRVerifierProofElement is an auto generated low-level Go binding around an user-defined struct.
type MMRVerifierProofElement struct {
	Cat     *big.Int
	Right   bool
	LeafNum uint64
	Res     MMRVerifierProofRes
}

// MMRVerifierProofInfo is an auto generated low-level Go binding around an user-defined struct.
type MMRVerifierProofInfo struct {
	RootHash       []byte
	RootDifficulty *big.Int
	LeafNumber     uint64
	Elems          []MMRVerifierProofElement
	Checked        []uint64
}

// MMRVerifierProofRes is an auto generated low-level Go binding around an user-defined struct.
type MMRVerifierProofRes struct {
	H  [32]byte
	Td *big.Int
}

// MMRVerifierVerifyElem is an auto generated low-level Go binding around an user-defined struct.
type MMRVerifierVerifyElem struct {
	Res        MMRVerifierProofRes
	Index      uint64
	LeafNumber uint64
}

// MMRVerifierMetaData contains all meta data concerning the MMRVerifier contract.
var MMRVerifierMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"MAX_UINT\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"superAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"leafNum\",\"type\":\"uint256\"}],\"internalType\":\"structMMRVerifier.MMR_config\",\"name\":\"_mmr\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"RootHash\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"RootDifficulty\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"LeafNumber\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"Cat\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"Right\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"LeafNum\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"h\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"td\",\"type\":\"uint256\"}],\"internalType\":\"structMMRVerifier.ProofRes\",\"name\":\"Res\",\"type\":\"tuple\"}],\"internalType\":\"structMMRVerifier.ProofElement[]\",\"name\":\"Elems\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64[]\",\"name\":\"Checked\",\"type\":\"uint64[]\"}],\"internalType\":\"structMMRVerifier.ProofInfo\",\"name\":\"_info\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Number\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"AggrWeight\",\"type\":\"uint256\"}],\"internalType\":\"structMMRVerifier.ProofBlock[]\",\"name\":\"_blocks\",\"type\":\"tuple[]\"}],\"name\":\"VerifyMMRProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"h\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"td\",\"type\":\"uint256\"}],\"internalType\":\"structMMRVerifier.ProofRes\",\"name\":\"Res\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"Index\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"LeafNumber\",\"type\":\"uint64\"}],\"internalType\":\"structMMRVerifier.VerifyElem[]\",\"name\":\"nodes\",\"type\":\"tuple[]\"}],\"name\":\"getRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_leafNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_cnt\",\"type\":\"uint256\"}],\"name\":\"RandomSample\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"a\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"b\",\"type\":\"bytes32\"}],\"name\":\"hashCom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"a\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"b\",\"type\":\"bytes\"}],\"name\":\"bytesCom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"h1\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"h2\",\"type\":\"bytes32\"}],\"name\":\"merge2\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// MMRVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use MMRVerifierMetaData.ABI instead.
var MMRVerifierABI = MMRVerifierMetaData.ABI

// MMRVerifier is an auto generated Go binding around an Ethereum contract.
type MMRVerifier struct {
	MMRVerifierCaller     // Read-only binding to the contract
	MMRVerifierTransactor // Write-only binding to the contract
	MMRVerifierFilterer   // Log filterer for contract events
}

// MMRVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type MMRVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MMRVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MMRVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MMRVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MMRVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MMRVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MMRVerifierSession struct {
	Contract     *MMRVerifier      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MMRVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MMRVerifierCallerSession struct {
	Contract *MMRVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// MMRVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MMRVerifierTransactorSession struct {
	Contract     *MMRVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MMRVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type MMRVerifierRaw struct {
	Contract *MMRVerifier // Generic contract binding to access the raw methods on
}

// MMRVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MMRVerifierCallerRaw struct {
	Contract *MMRVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// MMRVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MMRVerifierTransactorRaw struct {
	Contract *MMRVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMMRVerifier creates a new instance of MMRVerifier, bound to a specific deployed contract.
func NewMMRVerifier(address common.Address, backend bind.ContractBackend) (*MMRVerifier, error) {
	contract, err := bindMMRVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MMRVerifier{MMRVerifierCaller: MMRVerifierCaller{contract: contract}, MMRVerifierTransactor: MMRVerifierTransactor{contract: contract}, MMRVerifierFilterer: MMRVerifierFilterer{contract: contract}}, nil
}

// NewMMRVerifierCaller creates a new read-only instance of MMRVerifier, bound to a specific deployed contract.
func NewMMRVerifierCaller(address common.Address, caller bind.ContractCaller) (*MMRVerifierCaller, error) {
	contract, err := bindMMRVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MMRVerifierCaller{contract: contract}, nil
}

// NewMMRVerifierTransactor creates a new write-only instance of MMRVerifier, bound to a specific deployed contract.
func NewMMRVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*MMRVerifierTransactor, error) {
	contract, err := bindMMRVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MMRVerifierTransactor{contract: contract}, nil
}

// NewMMRVerifierFilterer creates a new log filterer instance of MMRVerifier, bound to a specific deployed contract.
func NewMMRVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*MMRVerifierFilterer, error) {
	contract, err := bindMMRVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MMRVerifierFilterer{contract: contract}, nil
}

// bindMMRVerifier binds a generic wrapper to an already deployed contract.
func bindMMRVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MMRVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MMRVerifier *MMRVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MMRVerifier.Contract.MMRVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MMRVerifier *MMRVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MMRVerifier.Contract.MMRVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MMRVerifier *MMRVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MMRVerifier.Contract.MMRVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MMRVerifier *MMRVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MMRVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MMRVerifier *MMRVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MMRVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MMRVerifier *MMRVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MMRVerifier.Contract.contract.Transact(opts, method, params...)
}

// MAXUINT is a free data retrieval call binding the contract method 0xe5b5019a.
//
// Solidity: function MAX_UINT() view returns(uint64)
func (_MMRVerifier *MMRVerifierCaller) MAXUINT(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _MMRVerifier.contract.Call(opts, &out, "MAX_UINT")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MAXUINT is a free data retrieval call binding the contract method 0xe5b5019a.
//
// Solidity: function MAX_UINT() view returns(uint64)
func (_MMRVerifier *MMRVerifierSession) MAXUINT() (uint64, error) {
	return _MMRVerifier.Contract.MAXUINT(&_MMRVerifier.CallOpts)
}

// MAXUINT is a free data retrieval call binding the contract method 0xe5b5019a.
//
// Solidity: function MAX_UINT() view returns(uint64)
func (_MMRVerifier *MMRVerifierCallerSession) MAXUINT() (uint64, error) {
	return _MMRVerifier.Contract.MAXUINT(&_MMRVerifier.CallOpts)
}

// RandomSample is a free data retrieval call binding the contract method 0xe0fa204f.
//
// Solidity: function RandomSample(uint256 _leafNumber, uint256 _cnt) pure returns(uint256[])
func (_MMRVerifier *MMRVerifierCaller) RandomSample(opts *bind.CallOpts, _leafNumber *big.Int, _cnt *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _MMRVerifier.contract.Call(opts, &out, "RandomSample", _leafNumber, _cnt)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// RandomSample is a free data retrieval call binding the contract method 0xe0fa204f.
//
// Solidity: function RandomSample(uint256 _leafNumber, uint256 _cnt) pure returns(uint256[])
func (_MMRVerifier *MMRVerifierSession) RandomSample(_leafNumber *big.Int, _cnt *big.Int) ([]*big.Int, error) {
	return _MMRVerifier.Contract.RandomSample(&_MMRVerifier.CallOpts, _leafNumber, _cnt)
}

// RandomSample is a free data retrieval call binding the contract method 0xe0fa204f.
//
// Solidity: function RandomSample(uint256 _leafNumber, uint256 _cnt) pure returns(uint256[])
func (_MMRVerifier *MMRVerifierCallerSession) RandomSample(_leafNumber *big.Int, _cnt *big.Int) ([]*big.Int, error) {
	return _MMRVerifier.Contract.RandomSample(&_MMRVerifier.CallOpts, _leafNumber, _cnt)
}

// VerifyMMRProof is a free data retrieval call binding the contract method 0x18a2c5a3.
//
// Solidity: function VerifyMMRProof((uint256) _mmr, (bytes,uint256,uint64,(uint256,bool,uint64,(bytes32,uint256))[],uint64[]) _info, (uint64,uint256)[] _blocks) pure returns(bool)
func (_MMRVerifier *MMRVerifierCaller) VerifyMMRProof(opts *bind.CallOpts, _mmr MMRVerifierMMRConfig, _info MMRVerifierProofInfo, _blocks []MMRVerifierProofBlock) (bool, error) {
	var out []interface{}
	err := _MMRVerifier.contract.Call(opts, &out, "VerifyMMRProof", _mmr, _info, _blocks)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyMMRProof is a free data retrieval call binding the contract method 0x18a2c5a3.
//
// Solidity: function VerifyMMRProof((uint256) _mmr, (bytes,uint256,uint64,(uint256,bool,uint64,(bytes32,uint256))[],uint64[]) _info, (uint64,uint256)[] _blocks) pure returns(bool)
func (_MMRVerifier *MMRVerifierSession) VerifyMMRProof(_mmr MMRVerifierMMRConfig, _info MMRVerifierProofInfo, _blocks []MMRVerifierProofBlock) (bool, error) {
	return _MMRVerifier.Contract.VerifyMMRProof(&_MMRVerifier.CallOpts, _mmr, _info, _blocks)
}

// VerifyMMRProof is a free data retrieval call binding the contract method 0x18a2c5a3.
//
// Solidity: function VerifyMMRProof((uint256) _mmr, (bytes,uint256,uint64,(uint256,bool,uint64,(bytes32,uint256))[],uint64[]) _info, (uint64,uint256)[] _blocks) pure returns(bool)
func (_MMRVerifier *MMRVerifierCallerSession) VerifyMMRProof(_mmr MMRVerifierMMRConfig, _info MMRVerifierProofInfo, _blocks []MMRVerifierProofBlock) (bool, error) {
	return _MMRVerifier.Contract.VerifyMMRProof(&_MMRVerifier.CallOpts, _mmr, _info, _blocks)
}

// BytesCom is a free data retrieval call binding the contract method 0xeaad451e.
//
// Solidity: function bytesCom(bytes a, bytes b) pure returns(bool)
func (_MMRVerifier *MMRVerifierCaller) BytesCom(opts *bind.CallOpts, a []byte, b []byte) (bool, error) {
	var out []interface{}
	err := _MMRVerifier.contract.Call(opts, &out, "bytesCom", a, b)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BytesCom is a free data retrieval call binding the contract method 0xeaad451e.
//
// Solidity: function bytesCom(bytes a, bytes b) pure returns(bool)
func (_MMRVerifier *MMRVerifierSession) BytesCom(a []byte, b []byte) (bool, error) {
	return _MMRVerifier.Contract.BytesCom(&_MMRVerifier.CallOpts, a, b)
}

// BytesCom is a free data retrieval call binding the contract method 0xeaad451e.
//
// Solidity: function bytesCom(bytes a, bytes b) pure returns(bool)
func (_MMRVerifier *MMRVerifierCallerSession) BytesCom(a []byte, b []byte) (bool, error) {
	return _MMRVerifier.Contract.BytesCom(&_MMRVerifier.CallOpts, a, b)
}

// GetRoot is a free data retrieval call binding the contract method 0x1e5c04d1.
//
// Solidity: function getRoot(((bytes32,uint256),uint64,uint64)[] nodes) pure returns(bytes32, uint256)
func (_MMRVerifier *MMRVerifierCaller) GetRoot(opts *bind.CallOpts, nodes []MMRVerifierVerifyElem) ([32]byte, *big.Int, error) {
	var out []interface{}
	err := _MMRVerifier.contract.Call(opts, &out, "getRoot", nodes)

	if err != nil {
		return *new([32]byte), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetRoot is a free data retrieval call binding the contract method 0x1e5c04d1.
//
// Solidity: function getRoot(((bytes32,uint256),uint64,uint64)[] nodes) pure returns(bytes32, uint256)
func (_MMRVerifier *MMRVerifierSession) GetRoot(nodes []MMRVerifierVerifyElem) ([32]byte, *big.Int, error) {
	return _MMRVerifier.Contract.GetRoot(&_MMRVerifier.CallOpts, nodes)
}

// GetRoot is a free data retrieval call binding the contract method 0x1e5c04d1.
//
// Solidity: function getRoot(((bytes32,uint256),uint64,uint64)[] nodes) pure returns(bytes32, uint256)
func (_MMRVerifier *MMRVerifierCallerSession) GetRoot(nodes []MMRVerifierVerifyElem) ([32]byte, *big.Int, error) {
	return _MMRVerifier.Contract.GetRoot(&_MMRVerifier.CallOpts, nodes)
}

// HashCom is a free data retrieval call binding the contract method 0x329a7a6f.
//
// Solidity: function hashCom(bytes32 a, bytes32 b) pure returns(bool)
func (_MMRVerifier *MMRVerifierCaller) HashCom(opts *bind.CallOpts, a [32]byte, b [32]byte) (bool, error) {
	var out []interface{}
	err := _MMRVerifier.contract.Call(opts, &out, "hashCom", a, b)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HashCom is a free data retrieval call binding the contract method 0x329a7a6f.
//
// Solidity: function hashCom(bytes32 a, bytes32 b) pure returns(bool)
func (_MMRVerifier *MMRVerifierSession) HashCom(a [32]byte, b [32]byte) (bool, error) {
	return _MMRVerifier.Contract.HashCom(&_MMRVerifier.CallOpts, a, b)
}

// HashCom is a free data retrieval call binding the contract method 0x329a7a6f.
//
// Solidity: function hashCom(bytes32 a, bytes32 b) pure returns(bool)
func (_MMRVerifier *MMRVerifierCallerSession) HashCom(a [32]byte, b [32]byte) (bool, error) {
	return _MMRVerifier.Contract.HashCom(&_MMRVerifier.CallOpts, a, b)
}

// Merge2 is a free data retrieval call binding the contract method 0x40d9f877.
//
// Solidity: function merge2(bytes32 h1, bytes32 h2) pure returns(bytes32)
func (_MMRVerifier *MMRVerifierCaller) Merge2(opts *bind.CallOpts, h1 [32]byte, h2 [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _MMRVerifier.contract.Call(opts, &out, "merge2", h1, h2)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Merge2 is a free data retrieval call binding the contract method 0x40d9f877.
//
// Solidity: function merge2(bytes32 h1, bytes32 h2) pure returns(bytes32)
func (_MMRVerifier *MMRVerifierSession) Merge2(h1 [32]byte, h2 [32]byte) ([32]byte, error) {
	return _MMRVerifier.Contract.Merge2(&_MMRVerifier.CallOpts, h1, h2)
}

// Merge2 is a free data retrieval call binding the contract method 0x40d9f877.
//
// Solidity: function merge2(bytes32 h1, bytes32 h2) pure returns(bytes32)
func (_MMRVerifier *MMRVerifierCallerSession) Merge2(h1 [32]byte, h2 [32]byte) ([32]byte, error) {
	return _MMRVerifier.Contract.Merge2(&_MMRVerifier.CallOpts, h1, h2)
}

// SuperAddress is a free data retrieval call binding the contract method 0x45d3b8db.
//
// Solidity: function superAddress() view returns(address)
func (_MMRVerifier *MMRVerifierCaller) SuperAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MMRVerifier.contract.Call(opts, &out, "superAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SuperAddress is a free data retrieval call binding the contract method 0x45d3b8db.
//
// Solidity: function superAddress() view returns(address)
func (_MMRVerifier *MMRVerifierSession) SuperAddress() (common.Address, error) {
	return _MMRVerifier.Contract.SuperAddress(&_MMRVerifier.CallOpts)
}

// SuperAddress is a free data retrieval call binding the contract method 0x45d3b8db.
//
// Solidity: function superAddress() view returns(address)
func (_MMRVerifier *MMRVerifierCallerSession) SuperAddress() (common.Address, error) {
	return _MMRVerifier.Contract.SuperAddress(&_MMRVerifier.CallOpts)
}

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

// ZKVerifierMetaData contains all meta data concerning the ZKVerifier contract.
var ZKVerifierMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"}],\"name\":\"verifyProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"r\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ZKVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use ZKVerifierMetaData.ABI instead.
var ZKVerifierABI = ZKVerifierMetaData.ABI

// ZKVerifier is an auto generated Go binding around an Ethereum contract.
type ZKVerifier struct {
	ZKVerifierCaller     // Read-only binding to the contract
	ZKVerifierTransactor // Write-only binding to the contract
	ZKVerifierFilterer   // Log filterer for contract events
}

// ZKVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZKVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZKVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZKVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZKVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZKVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZKVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZKVerifierSession struct {
	Contract     *ZKVerifier       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZKVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZKVerifierCallerSession struct {
	Contract *ZKVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ZKVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZKVerifierTransactorSession struct {
	Contract     *ZKVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ZKVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZKVerifierRaw struct {
	Contract *ZKVerifier // Generic contract binding to access the raw methods on
}

// ZKVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZKVerifierCallerRaw struct {
	Contract *ZKVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// ZKVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZKVerifierTransactorRaw struct {
	Contract *ZKVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZKVerifier creates a new instance of ZKVerifier, bound to a specific deployed contract.
func NewZKVerifier(address common.Address, backend bind.ContractBackend) (*ZKVerifier, error) {
	contract, err := bindZKVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZKVerifier{ZKVerifierCaller: ZKVerifierCaller{contract: contract}, ZKVerifierTransactor: ZKVerifierTransactor{contract: contract}, ZKVerifierFilterer: ZKVerifierFilterer{contract: contract}}, nil
}

// NewZKVerifierCaller creates a new read-only instance of ZKVerifier, bound to a specific deployed contract.
func NewZKVerifierCaller(address common.Address, caller bind.ContractCaller) (*ZKVerifierCaller, error) {
	contract, err := bindZKVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZKVerifierCaller{contract: contract}, nil
}

// NewZKVerifierTransactor creates a new write-only instance of ZKVerifier, bound to a specific deployed contract.
func NewZKVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*ZKVerifierTransactor, error) {
	contract, err := bindZKVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZKVerifierTransactor{contract: contract}, nil
}

// NewZKVerifierFilterer creates a new log filterer instance of ZKVerifier, bound to a specific deployed contract.
func NewZKVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*ZKVerifierFilterer, error) {
	contract, err := bindZKVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZKVerifierFilterer{contract: contract}, nil
}

// bindZKVerifier binds a generic wrapper to an already deployed contract.
func bindZKVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZKVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZKVerifier *ZKVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZKVerifier.Contract.ZKVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZKVerifier *ZKVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZKVerifier.Contract.ZKVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZKVerifier *ZKVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZKVerifier.Contract.ZKVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZKVerifier *ZKVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZKVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZKVerifier *ZKVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZKVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZKVerifier *ZKVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZKVerifier.Contract.contract.Transact(opts, method, params...)
}

// VerifyProof is a free data retrieval call binding the contract method 0x6668a9fa.
//
// Solidity: function verifyProof(uint256[2] a, uint256[2][2] b, uint256[2] c) view returns(bool r)
func (_ZKVerifier *ZKVerifierCaller) VerifyProof(opts *bind.CallOpts, a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int) (bool, error) {
	var out []interface{}
	err := _ZKVerifier.contract.Call(opts, &out, "verifyProof", a, b, c)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyProof is a free data retrieval call binding the contract method 0x6668a9fa.
//
// Solidity: function verifyProof(uint256[2] a, uint256[2][2] b, uint256[2] c) view returns(bool r)
func (_ZKVerifier *ZKVerifierSession) VerifyProof(a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int) (bool, error) {
	return _ZKVerifier.Contract.VerifyProof(&_ZKVerifier.CallOpts, a, b, c)
}

// VerifyProof is a free data retrieval call binding the contract method 0x6668a9fa.
//
// Solidity: function verifyProof(uint256[2] a, uint256[2][2] b, uint256[2] c) view returns(bool r)
func (_ZKVerifier *ZKVerifierCallerSession) VerifyProof(a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int) (bool, error) {
	return _ZKVerifier.Contract.VerifyProof(&_ZKVerifier.CallOpts, a, b, c)
}

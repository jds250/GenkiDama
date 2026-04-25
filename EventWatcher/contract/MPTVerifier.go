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

// MPTVerifierMerkleProof is an auto generated low-level Go binding around an user-defined struct.
type MPTVerifierMerkleProof struct {
	ExpectedRoot  [32]byte
	Key           []byte
	Proof         [][]byte
	KeyIndex      *big.Int
	ProofIndex    *big.Int
	ExpectedValue []byte
}

// MPTVerifierMetaData contains all meta data concerning the MPTVerifier contract.
var MPTVerifierMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"expectedRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"proof\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"keyIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"proofIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"expectedValue\",\"type\":\"bytes\"}],\"internalType\":\"structMPTVerifier.MerkleProof\",\"name\":\"data\",\"type\":\"tuple\"}],\"name\":\"verifyProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true}]",
}

// MPTVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use MPTVerifierMetaData.ABI instead.
var MPTVerifierABI = MPTVerifierMetaData.ABI

// MPTVerifier is an auto generated Go binding around an Ethereum contract.
type MPTVerifier struct {
	MPTVerifierCaller     // Read-only binding to the contract
	MPTVerifierTransactor // Write-only binding to the contract
	MPTVerifierFilterer   // Log filterer for contract events
}

// MPTVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type MPTVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MPTVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MPTVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MPTVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MPTVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MPTVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MPTVerifierSession struct {
	Contract     *MPTVerifier      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MPTVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MPTVerifierCallerSession struct {
	Contract *MPTVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// MPTVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MPTVerifierTransactorSession struct {
	Contract     *MPTVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MPTVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type MPTVerifierRaw struct {
	Contract *MPTVerifier // Generic contract binding to access the raw methods on
}

// MPTVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MPTVerifierCallerRaw struct {
	Contract *MPTVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// MPTVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MPTVerifierTransactorRaw struct {
	Contract *MPTVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMPTVerifier creates a new instance of MPTVerifier, bound to a specific deployed contract.
func NewMPTVerifier(address common.Address, backend bind.ContractBackend) (*MPTVerifier, error) {
	contract, err := bindMPTVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MPTVerifier{MPTVerifierCaller: MPTVerifierCaller{contract: contract}, MPTVerifierTransactor: MPTVerifierTransactor{contract: contract}, MPTVerifierFilterer: MPTVerifierFilterer{contract: contract}}, nil
}

// NewMPTVerifierCaller creates a new read-only instance of MPTVerifier, bound to a specific deployed contract.
func NewMPTVerifierCaller(address common.Address, caller bind.ContractCaller) (*MPTVerifierCaller, error) {
	contract, err := bindMPTVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MPTVerifierCaller{contract: contract}, nil
}

// NewMPTVerifierTransactor creates a new write-only instance of MPTVerifier, bound to a specific deployed contract.
func NewMPTVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*MPTVerifierTransactor, error) {
	contract, err := bindMPTVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MPTVerifierTransactor{contract: contract}, nil
}

// NewMPTVerifierFilterer creates a new log filterer instance of MPTVerifier, bound to a specific deployed contract.
func NewMPTVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*MPTVerifierFilterer, error) {
	contract, err := bindMPTVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MPTVerifierFilterer{contract: contract}, nil
}

// bindMPTVerifier binds a generic wrapper to an already deployed contract.
func bindMPTVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MPTVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MPTVerifier *MPTVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MPTVerifier.Contract.MPTVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MPTVerifier *MPTVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MPTVerifier.Contract.MPTVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MPTVerifier *MPTVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MPTVerifier.Contract.MPTVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MPTVerifier *MPTVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MPTVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MPTVerifier *MPTVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MPTVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MPTVerifier *MPTVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MPTVerifier.Contract.contract.Transact(opts, method, params...)
}

// VerifyProof is a free data retrieval call binding the contract method 0x51ca3580.
//
// Solidity: function verifyProof((bytes32,bytes,bytes[],uint256,uint256,bytes) data) view returns(bool)
func (_MPTVerifier *MPTVerifierCaller) VerifyProof(opts *bind.CallOpts, data MPTVerifierMerkleProof) (bool, error) {
	var out []interface{}
	err := _MPTVerifier.contract.Call(opts, &out, "verifyProof", data)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyProof is a free data retrieval call binding the contract method 0x51ca3580.
//
// Solidity: function verifyProof((bytes32,bytes,bytes[],uint256,uint256,bytes) data) view returns(bool)
func (_MPTVerifier *MPTVerifierSession) VerifyProof(data MPTVerifierMerkleProof) (bool, error) {
	return _MPTVerifier.Contract.VerifyProof(&_MPTVerifier.CallOpts, data)
}

// VerifyProof is a free data retrieval call binding the contract method 0x51ca3580.
//
// Solidity: function verifyProof((bytes32,bytes,bytes[],uint256,uint256,bytes) data) view returns(bool)
func (_MPTVerifier *MPTVerifierCallerSession) VerifyProof(data MPTVerifierMerkleProof) (bool, error) {
	return _MPTVerifier.Contract.VerifyProof(&_MPTVerifier.CallOpts, data)
}

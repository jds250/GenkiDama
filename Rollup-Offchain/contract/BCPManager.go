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

// DataTypesChainCheckPoint is an auto generated low-level Go binding around an user-defined struct.
type DataTypesChainCheckPoint struct {
	Index        *big.Int
	ChainID      *big.Int
	ChainType    *big.Int
	BlockID      *big.Int
	StateRoot    []byte
	BlockData    []byte
	ContractAddr common.Address
}

// DataTypesChallengeState is an auto generated low-level Go binding around an user-defined struct.
type DataTypesChallengeState struct {
	ChainID     *big.Int
	L2BlockID   *big.Int
	L1BlockID   *big.Int
	L1StateRoot []byte
	L1BlockData []byte
	StateProof  [][]byte
	Account     common.Address
}

// DataTypesL1BlockInfo is an auto generated low-level Go binding around an user-defined struct.
type DataTypesL1BlockInfo struct {
	ChainID   *big.Int
	BlockID   *big.Int
	StateRoot []byte
	BlockData []byte
}

// DataTypesZKProof is an auto generated low-level Go binding around an user-defined struct.
type DataTypesZKProof struct {
	A [2]*big.Int
	B [2][2]*big.Int
	C [2]*big.Int
}

// BCPManagerMetaData contains all meta data concerning the BCPManager contract.
var BCPManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_dataTypesAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_zkVerifierAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"index\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"style\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"desc\",\"type\":\"string\"}],\"name\":\"BlockInfoNotify\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"index\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"currentState\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"desc\",\"type\":\"string\"}],\"name\":\"ChallengeStateNotify\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ZERO_BYTES32\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"chainCheckPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"stateRoot\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blockData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"contractAddr\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"challengePool\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"index\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l2BlockID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l1BlockID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"l1StateRoot\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"l1BlockData\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"stateProof\",\"type\":\"bytes[]\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"internalType\":\"structDataTypes.ChallengeState\",\"name\":\"detail\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"confirmTime\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"questioner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"state\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"begin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"stateRoot\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blockData\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.L1BlockInfo\",\"name\":\"beginBlock\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"stateRoot\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blockData\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.L1BlockInfo\",\"name\":\"endBlock\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"stateRoot\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blockData\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.L1BlockInfo\",\"name\":\"middleBlock\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"checkPointPool\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"stateRoot\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blockData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"contractAddr\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"committerAddresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"superAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l2BlockID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l1BlockID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"l1StateRoot\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"l1BlockData\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"stateProof\",\"type\":\"bytes[]\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"internalType\":\"structDataTypes.ChallengeState\",\"name\":\"_challenge\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"_index\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"_header\",\"type\":\"bytes\"}],\"name\":\"ChallengeCreate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_index\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"_ack\",\"type\":\"bool\"}],\"name\":\"ChallengeQuestion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_index\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"stateRoot\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blockData\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.L1BlockInfo\",\"name\":\"_middleBlock\",\"type\":\"tuple\"}],\"name\":\"ChallengeResponse\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_index\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"}],\"internalType\":\"structDataTypes.ZKProof\",\"name\":\"_finalProof\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_header1\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_header2\",\"type\":\"bytes\"}],\"name\":\"FinalChallenge\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"stateRoot\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blockData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"contractAddr\",\"type\":\"address\"}],\"internalType\":\"structDataTypes.ChainCheckPoint\",\"name\":\"_checkPoint\",\"type\":\"tuple\"}],\"name\":\"CheckPoint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BCPManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use BCPManagerMetaData.ABI instead.
var BCPManagerABI = BCPManagerMetaData.ABI

// BCPManager is an auto generated Go binding around an Ethereum contract.
type BCPManager struct {
	BCPManagerCaller     // Read-only binding to the contract
	BCPManagerTransactor // Write-only binding to the contract
	BCPManagerFilterer   // Log filterer for contract events
}

// BCPManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type BCPManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BCPManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BCPManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BCPManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BCPManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BCPManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BCPManagerSession struct {
	Contract     *BCPManager       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BCPManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BCPManagerCallerSession struct {
	Contract *BCPManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BCPManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BCPManagerTransactorSession struct {
	Contract     *BCPManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BCPManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type BCPManagerRaw struct {
	Contract *BCPManager // Generic contract binding to access the raw methods on
}

// BCPManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BCPManagerCallerRaw struct {
	Contract *BCPManagerCaller // Generic read-only contract binding to access the raw methods on
}

// BCPManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BCPManagerTransactorRaw struct {
	Contract *BCPManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBCPManager creates a new instance of BCPManager, bound to a specific deployed contract.
func NewBCPManager(address common.Address, backend bind.ContractBackend) (*BCPManager, error) {
	contract, err := bindBCPManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BCPManager{BCPManagerCaller: BCPManagerCaller{contract: contract}, BCPManagerTransactor: BCPManagerTransactor{contract: contract}, BCPManagerFilterer: BCPManagerFilterer{contract: contract}}, nil
}

// NewBCPManagerCaller creates a new read-only instance of BCPManager, bound to a specific deployed contract.
func NewBCPManagerCaller(address common.Address, caller bind.ContractCaller) (*BCPManagerCaller, error) {
	contract, err := bindBCPManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BCPManagerCaller{contract: contract}, nil
}

// NewBCPManagerTransactor creates a new write-only instance of BCPManager, bound to a specific deployed contract.
func NewBCPManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*BCPManagerTransactor, error) {
	contract, err := bindBCPManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BCPManagerTransactor{contract: contract}, nil
}

// NewBCPManagerFilterer creates a new log filterer instance of BCPManager, bound to a specific deployed contract.
func NewBCPManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*BCPManagerFilterer, error) {
	contract, err := bindBCPManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BCPManagerFilterer{contract: contract}, nil
}

// bindBCPManager binds a generic wrapper to an already deployed contract.
func bindBCPManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BCPManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BCPManager *BCPManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BCPManager.Contract.BCPManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BCPManager *BCPManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BCPManager.Contract.BCPManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BCPManager *BCPManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BCPManager.Contract.BCPManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BCPManager *BCPManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BCPManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BCPManager *BCPManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BCPManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BCPManager *BCPManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BCPManager.Contract.contract.Transact(opts, method, params...)
}

// ZEROBYTES32 is a free data retrieval call binding the contract method 0x069321b0.
//
// Solidity: function ZERO_BYTES32() view returns(bytes32)
func (_BCPManager *BCPManagerCaller) ZEROBYTES32(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BCPManager.contract.Call(opts, &out, "ZERO_BYTES32")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ZEROBYTES32 is a free data retrieval call binding the contract method 0x069321b0.
//
// Solidity: function ZERO_BYTES32() view returns(bytes32)
func (_BCPManager *BCPManagerSession) ZEROBYTES32() ([32]byte, error) {
	return _BCPManager.Contract.ZEROBYTES32(&_BCPManager.CallOpts)
}

// ZEROBYTES32 is a free data retrieval call binding the contract method 0x069321b0.
//
// Solidity: function ZERO_BYTES32() view returns(bytes32)
func (_BCPManager *BCPManagerCallerSession) ZEROBYTES32() ([32]byte, error) {
	return _BCPManager.Contract.ZEROBYTES32(&_BCPManager.CallOpts)
}

// ChainCheckPoints is a free data retrieval call binding the contract method 0xc8650f90.
//
// Solidity: function chainCheckPoints(uint256 ) view returns(uint256 index, uint256 chainID, uint256 chainType, uint256 blockID, bytes stateRoot, bytes blockData, address contractAddr)
func (_BCPManager *BCPManagerCaller) ChainCheckPoints(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Index        *big.Int
	ChainID      *big.Int
	ChainType    *big.Int
	BlockID      *big.Int
	StateRoot    []byte
	BlockData    []byte
	ContractAddr common.Address
}, error) {
	var out []interface{}
	err := _BCPManager.contract.Call(opts, &out, "chainCheckPoints", arg0)

	outstruct := new(struct {
		Index        *big.Int
		ChainID      *big.Int
		ChainType    *big.Int
		BlockID      *big.Int
		StateRoot    []byte
		BlockData    []byte
		ContractAddr common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Index = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ChainID = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ChainType = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.BlockID = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.StateRoot = *abi.ConvertType(out[4], new([]byte)).(*[]byte)
	outstruct.BlockData = *abi.ConvertType(out[5], new([]byte)).(*[]byte)
	outstruct.ContractAddr = *abi.ConvertType(out[6], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// ChainCheckPoints is a free data retrieval call binding the contract method 0xc8650f90.
//
// Solidity: function chainCheckPoints(uint256 ) view returns(uint256 index, uint256 chainID, uint256 chainType, uint256 blockID, bytes stateRoot, bytes blockData, address contractAddr)
func (_BCPManager *BCPManagerSession) ChainCheckPoints(arg0 *big.Int) (struct {
	Index        *big.Int
	ChainID      *big.Int
	ChainType    *big.Int
	BlockID      *big.Int
	StateRoot    []byte
	BlockData    []byte
	ContractAddr common.Address
}, error) {
	return _BCPManager.Contract.ChainCheckPoints(&_BCPManager.CallOpts, arg0)
}

// ChainCheckPoints is a free data retrieval call binding the contract method 0xc8650f90.
//
// Solidity: function chainCheckPoints(uint256 ) view returns(uint256 index, uint256 chainID, uint256 chainType, uint256 blockID, bytes stateRoot, bytes blockData, address contractAddr)
func (_BCPManager *BCPManagerCallerSession) ChainCheckPoints(arg0 *big.Int) (struct {
	Index        *big.Int
	ChainID      *big.Int
	ChainType    *big.Int
	BlockID      *big.Int
	StateRoot    []byte
	BlockData    []byte
	ContractAddr common.Address
}, error) {
	return _BCPManager.Contract.ChainCheckPoints(&_BCPManager.CallOpts, arg0)
}

// ChainID is a free data retrieval call binding the contract method 0xadc879e9.
//
// Solidity: function chainID() view returns(uint256)
func (_BCPManager *BCPManagerCaller) ChainID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BCPManager.contract.Call(opts, &out, "chainID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChainID is a free data retrieval call binding the contract method 0xadc879e9.
//
// Solidity: function chainID() view returns(uint256)
func (_BCPManager *BCPManagerSession) ChainID() (*big.Int, error) {
	return _BCPManager.Contract.ChainID(&_BCPManager.CallOpts)
}

// ChainID is a free data retrieval call binding the contract method 0xadc879e9.
//
// Solidity: function chainID() view returns(uint256)
func (_BCPManager *BCPManagerCallerSession) ChainID() (*big.Int, error) {
	return _BCPManager.Contract.ChainID(&_BCPManager.CallOpts)
}

// ChallengePool is a free data retrieval call binding the contract method 0x419a0282.
//
// Solidity: function challengePool(string ) view returns(string index, (uint256,uint256,uint256,bytes,bytes,bytes[],address) detail, uint256 confirmTime, address challenger, address questioner, uint256 state, uint256 begin, uint256 end, (uint256,uint256,bytes,bytes) beginBlock, (uint256,uint256,bytes,bytes) endBlock, (uint256,uint256,bytes,bytes) middleBlock)
func (_BCPManager *BCPManagerCaller) ChallengePool(opts *bind.CallOpts, arg0 string) (struct {
	Index       string
	Detail      DataTypesChallengeState
	ConfirmTime *big.Int
	Challenger  common.Address
	Questioner  common.Address
	State       *big.Int
	Begin       *big.Int
	End         *big.Int
	BeginBlock  DataTypesL1BlockInfo
	EndBlock    DataTypesL1BlockInfo
	MiddleBlock DataTypesL1BlockInfo
}, error) {
	var out []interface{}
	err := _BCPManager.contract.Call(opts, &out, "challengePool", arg0)

	outstruct := new(struct {
		Index       string
		Detail      DataTypesChallengeState
		ConfirmTime *big.Int
		Challenger  common.Address
		Questioner  common.Address
		State       *big.Int
		Begin       *big.Int
		End         *big.Int
		BeginBlock  DataTypesL1BlockInfo
		EndBlock    DataTypesL1BlockInfo
		MiddleBlock DataTypesL1BlockInfo
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Index = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Detail = *abi.ConvertType(out[1], new(DataTypesChallengeState)).(*DataTypesChallengeState)
	outstruct.ConfirmTime = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Challenger = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Questioner = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.State = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Begin = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.End = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.BeginBlock = *abi.ConvertType(out[8], new(DataTypesL1BlockInfo)).(*DataTypesL1BlockInfo)
	outstruct.EndBlock = *abi.ConvertType(out[9], new(DataTypesL1BlockInfo)).(*DataTypesL1BlockInfo)
	outstruct.MiddleBlock = *abi.ConvertType(out[10], new(DataTypesL1BlockInfo)).(*DataTypesL1BlockInfo)

	return *outstruct, err

}

// ChallengePool is a free data retrieval call binding the contract method 0x419a0282.
//
// Solidity: function challengePool(string ) view returns(string index, (uint256,uint256,uint256,bytes,bytes,bytes[],address) detail, uint256 confirmTime, address challenger, address questioner, uint256 state, uint256 begin, uint256 end, (uint256,uint256,bytes,bytes) beginBlock, (uint256,uint256,bytes,bytes) endBlock, (uint256,uint256,bytes,bytes) middleBlock)
func (_BCPManager *BCPManagerSession) ChallengePool(arg0 string) (struct {
	Index       string
	Detail      DataTypesChallengeState
	ConfirmTime *big.Int
	Challenger  common.Address
	Questioner  common.Address
	State       *big.Int
	Begin       *big.Int
	End         *big.Int
	BeginBlock  DataTypesL1BlockInfo
	EndBlock    DataTypesL1BlockInfo
	MiddleBlock DataTypesL1BlockInfo
}, error) {
	return _BCPManager.Contract.ChallengePool(&_BCPManager.CallOpts, arg0)
}

// ChallengePool is a free data retrieval call binding the contract method 0x419a0282.
//
// Solidity: function challengePool(string ) view returns(string index, (uint256,uint256,uint256,bytes,bytes,bytes[],address) detail, uint256 confirmTime, address challenger, address questioner, uint256 state, uint256 begin, uint256 end, (uint256,uint256,bytes,bytes) beginBlock, (uint256,uint256,bytes,bytes) endBlock, (uint256,uint256,bytes,bytes) middleBlock)
func (_BCPManager *BCPManagerCallerSession) ChallengePool(arg0 string) (struct {
	Index       string
	Detail      DataTypesChallengeState
	ConfirmTime *big.Int
	Challenger  common.Address
	Questioner  common.Address
	State       *big.Int
	Begin       *big.Int
	End         *big.Int
	BeginBlock  DataTypesL1BlockInfo
	EndBlock    DataTypesL1BlockInfo
	MiddleBlock DataTypesL1BlockInfo
}, error) {
	return _BCPManager.Contract.ChallengePool(&_BCPManager.CallOpts, arg0)
}

// CheckPointPool is a free data retrieval call binding the contract method 0xb9783568.
//
// Solidity: function checkPointPool(string ) view returns(uint256 index, uint256 chainID, uint256 chainType, uint256 blockID, bytes stateRoot, bytes blockData, address contractAddr)
func (_BCPManager *BCPManagerCaller) CheckPointPool(opts *bind.CallOpts, arg0 string) (struct {
	Index        *big.Int
	ChainID      *big.Int
	ChainType    *big.Int
	BlockID      *big.Int
	StateRoot    []byte
	BlockData    []byte
	ContractAddr common.Address
}, error) {
	var out []interface{}
	err := _BCPManager.contract.Call(opts, &out, "checkPointPool", arg0)

	outstruct := new(struct {
		Index        *big.Int
		ChainID      *big.Int
		ChainType    *big.Int
		BlockID      *big.Int
		StateRoot    []byte
		BlockData    []byte
		ContractAddr common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Index = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ChainID = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ChainType = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.BlockID = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.StateRoot = *abi.ConvertType(out[4], new([]byte)).(*[]byte)
	outstruct.BlockData = *abi.ConvertType(out[5], new([]byte)).(*[]byte)
	outstruct.ContractAddr = *abi.ConvertType(out[6], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// CheckPointPool is a free data retrieval call binding the contract method 0xb9783568.
//
// Solidity: function checkPointPool(string ) view returns(uint256 index, uint256 chainID, uint256 chainType, uint256 blockID, bytes stateRoot, bytes blockData, address contractAddr)
func (_BCPManager *BCPManagerSession) CheckPointPool(arg0 string) (struct {
	Index        *big.Int
	ChainID      *big.Int
	ChainType    *big.Int
	BlockID      *big.Int
	StateRoot    []byte
	BlockData    []byte
	ContractAddr common.Address
}, error) {
	return _BCPManager.Contract.CheckPointPool(&_BCPManager.CallOpts, arg0)
}

// CheckPointPool is a free data retrieval call binding the contract method 0xb9783568.
//
// Solidity: function checkPointPool(string ) view returns(uint256 index, uint256 chainID, uint256 chainType, uint256 blockID, bytes stateRoot, bytes blockData, address contractAddr)
func (_BCPManager *BCPManagerCallerSession) CheckPointPool(arg0 string) (struct {
	Index        *big.Int
	ChainID      *big.Int
	ChainType    *big.Int
	BlockID      *big.Int
	StateRoot    []byte
	BlockData    []byte
	ContractAddr common.Address
}, error) {
	return _BCPManager.Contract.CheckPointPool(&_BCPManager.CallOpts, arg0)
}

// CommitterAddresses is a free data retrieval call binding the contract method 0xd63d1112.
//
// Solidity: function committerAddresses(uint256 ) view returns(address)
func (_BCPManager *BCPManagerCaller) CommitterAddresses(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _BCPManager.contract.Call(opts, &out, "committerAddresses", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CommitterAddresses is a free data retrieval call binding the contract method 0xd63d1112.
//
// Solidity: function committerAddresses(uint256 ) view returns(address)
func (_BCPManager *BCPManagerSession) CommitterAddresses(arg0 *big.Int) (common.Address, error) {
	return _BCPManager.Contract.CommitterAddresses(&_BCPManager.CallOpts, arg0)
}

// CommitterAddresses is a free data retrieval call binding the contract method 0xd63d1112.
//
// Solidity: function committerAddresses(uint256 ) view returns(address)
func (_BCPManager *BCPManagerCallerSession) CommitterAddresses(arg0 *big.Int) (common.Address, error) {
	return _BCPManager.Contract.CommitterAddresses(&_BCPManager.CallOpts, arg0)
}

// SuperAddress is a free data retrieval call binding the contract method 0x45d3b8db.
//
// Solidity: function superAddress() view returns(address)
func (_BCPManager *BCPManagerCaller) SuperAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BCPManager.contract.Call(opts, &out, "superAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SuperAddress is a free data retrieval call binding the contract method 0x45d3b8db.
//
// Solidity: function superAddress() view returns(address)
func (_BCPManager *BCPManagerSession) SuperAddress() (common.Address, error) {
	return _BCPManager.Contract.SuperAddress(&_BCPManager.CallOpts)
}

// SuperAddress is a free data retrieval call binding the contract method 0x45d3b8db.
//
// Solidity: function superAddress() view returns(address)
func (_BCPManager *BCPManagerCallerSession) SuperAddress() (common.Address, error) {
	return _BCPManager.Contract.SuperAddress(&_BCPManager.CallOpts)
}

// ChallengeCreate is a paid mutator transaction binding the contract method 0xcc3c68b7.
//
// Solidity: function ChallengeCreate((uint256,uint256,uint256,bytes,bytes,bytes[],address) _challenge, string _index, bytes _header) returns(uint256, string)
func (_BCPManager *BCPManagerTransactor) ChallengeCreate(opts *bind.TransactOpts, _challenge DataTypesChallengeState, _index string, _header []byte) (*types.Transaction, error) {
	return _BCPManager.contract.Transact(opts, "ChallengeCreate", _challenge, _index, _header)
}

// ChallengeCreate is a paid mutator transaction binding the contract method 0xcc3c68b7.
//
// Solidity: function ChallengeCreate((uint256,uint256,uint256,bytes,bytes,bytes[],address) _challenge, string _index, bytes _header) returns(uint256, string)
func (_BCPManager *BCPManagerSession) ChallengeCreate(_challenge DataTypesChallengeState, _index string, _header []byte) (*types.Transaction, error) {
	return _BCPManager.Contract.ChallengeCreate(&_BCPManager.TransactOpts, _challenge, _index, _header)
}

// ChallengeCreate is a paid mutator transaction binding the contract method 0xcc3c68b7.
//
// Solidity: function ChallengeCreate((uint256,uint256,uint256,bytes,bytes,bytes[],address) _challenge, string _index, bytes _header) returns(uint256, string)
func (_BCPManager *BCPManagerTransactorSession) ChallengeCreate(_challenge DataTypesChallengeState, _index string, _header []byte) (*types.Transaction, error) {
	return _BCPManager.Contract.ChallengeCreate(&_BCPManager.TransactOpts, _challenge, _index, _header)
}

// ChallengeQuestion is a paid mutator transaction binding the contract method 0xbbacac68.
//
// Solidity: function ChallengeQuestion(string _index, bool _ack) returns(uint256, string)
func (_BCPManager *BCPManagerTransactor) ChallengeQuestion(opts *bind.TransactOpts, _index string, _ack bool) (*types.Transaction, error) {
	return _BCPManager.contract.Transact(opts, "ChallengeQuestion", _index, _ack)
}

// ChallengeQuestion is a paid mutator transaction binding the contract method 0xbbacac68.
//
// Solidity: function ChallengeQuestion(string _index, bool _ack) returns(uint256, string)
func (_BCPManager *BCPManagerSession) ChallengeQuestion(_index string, _ack bool) (*types.Transaction, error) {
	return _BCPManager.Contract.ChallengeQuestion(&_BCPManager.TransactOpts, _index, _ack)
}

// ChallengeQuestion is a paid mutator transaction binding the contract method 0xbbacac68.
//
// Solidity: function ChallengeQuestion(string _index, bool _ack) returns(uint256, string)
func (_BCPManager *BCPManagerTransactorSession) ChallengeQuestion(_index string, _ack bool) (*types.Transaction, error) {
	return _BCPManager.Contract.ChallengeQuestion(&_BCPManager.TransactOpts, _index, _ack)
}

// ChallengeResponse is a paid mutator transaction binding the contract method 0x414283d4.
//
// Solidity: function ChallengeResponse(string _index, (uint256,uint256,bytes,bytes) _middleBlock) returns(uint256, string)
func (_BCPManager *BCPManagerTransactor) ChallengeResponse(opts *bind.TransactOpts, _index string, _middleBlock DataTypesL1BlockInfo) (*types.Transaction, error) {
	return _BCPManager.contract.Transact(opts, "ChallengeResponse", _index, _middleBlock)
}

// ChallengeResponse is a paid mutator transaction binding the contract method 0x414283d4.
//
// Solidity: function ChallengeResponse(string _index, (uint256,uint256,bytes,bytes) _middleBlock) returns(uint256, string)
func (_BCPManager *BCPManagerSession) ChallengeResponse(_index string, _middleBlock DataTypesL1BlockInfo) (*types.Transaction, error) {
	return _BCPManager.Contract.ChallengeResponse(&_BCPManager.TransactOpts, _index, _middleBlock)
}

// ChallengeResponse is a paid mutator transaction binding the contract method 0x414283d4.
//
// Solidity: function ChallengeResponse(string _index, (uint256,uint256,bytes,bytes) _middleBlock) returns(uint256, string)
func (_BCPManager *BCPManagerTransactorSession) ChallengeResponse(_index string, _middleBlock DataTypesL1BlockInfo) (*types.Transaction, error) {
	return _BCPManager.Contract.ChallengeResponse(&_BCPManager.TransactOpts, _index, _middleBlock)
}

// CheckPoint is a paid mutator transaction binding the contract method 0xbb3e46a4.
//
// Solidity: function CheckPoint((uint256,uint256,uint256,uint256,bytes,bytes,address) _checkPoint) returns(uint256, string)
func (_BCPManager *BCPManagerTransactor) CheckPoint(opts *bind.TransactOpts, _checkPoint DataTypesChainCheckPoint) (*types.Transaction, error) {
	return _BCPManager.contract.Transact(opts, "CheckPoint", _checkPoint)
}

// CheckPoint is a paid mutator transaction binding the contract method 0xbb3e46a4.
//
// Solidity: function CheckPoint((uint256,uint256,uint256,uint256,bytes,bytes,address) _checkPoint) returns(uint256, string)
func (_BCPManager *BCPManagerSession) CheckPoint(_checkPoint DataTypesChainCheckPoint) (*types.Transaction, error) {
	return _BCPManager.Contract.CheckPoint(&_BCPManager.TransactOpts, _checkPoint)
}

// CheckPoint is a paid mutator transaction binding the contract method 0xbb3e46a4.
//
// Solidity: function CheckPoint((uint256,uint256,uint256,uint256,bytes,bytes,address) _checkPoint) returns(uint256, string)
func (_BCPManager *BCPManagerTransactorSession) CheckPoint(_checkPoint DataTypesChainCheckPoint) (*types.Transaction, error) {
	return _BCPManager.Contract.CheckPoint(&_BCPManager.TransactOpts, _checkPoint)
}

// FinalChallenge is a paid mutator transaction binding the contract method 0x36f0084d.
//
// Solidity: function FinalChallenge(string _index, (uint256[2],uint256[2][2],uint256[2]) _finalProof, bytes _header1, bytes _header2) returns(uint256, string)
func (_BCPManager *BCPManagerTransactor) FinalChallenge(opts *bind.TransactOpts, _index string, _finalProof DataTypesZKProof, _header1 []byte, _header2 []byte) (*types.Transaction, error) {
	return _BCPManager.contract.Transact(opts, "FinalChallenge", _index, _finalProof, _header1, _header2)
}

// FinalChallenge is a paid mutator transaction binding the contract method 0x36f0084d.
//
// Solidity: function FinalChallenge(string _index, (uint256[2],uint256[2][2],uint256[2]) _finalProof, bytes _header1, bytes _header2) returns(uint256, string)
func (_BCPManager *BCPManagerSession) FinalChallenge(_index string, _finalProof DataTypesZKProof, _header1 []byte, _header2 []byte) (*types.Transaction, error) {
	return _BCPManager.Contract.FinalChallenge(&_BCPManager.TransactOpts, _index, _finalProof, _header1, _header2)
}

// FinalChallenge is a paid mutator transaction binding the contract method 0x36f0084d.
//
// Solidity: function FinalChallenge(string _index, (uint256[2],uint256[2][2],uint256[2]) _finalProof, bytes _header1, bytes _header2) returns(uint256, string)
func (_BCPManager *BCPManagerTransactorSession) FinalChallenge(_index string, _finalProof DataTypesZKProof, _header1 []byte, _header2 []byte) (*types.Transaction, error) {
	return _BCPManager.Contract.FinalChallenge(&_BCPManager.TransactOpts, _index, _finalProof, _header1, _header2)
}

// BCPManagerBlockInfoNotifyIterator is returned from FilterBlockInfoNotify and is used to iterate over the raw logs and unpacked data for BlockInfoNotify events raised by the BCPManager contract.
type BCPManagerBlockInfoNotifyIterator struct {
	Event *BCPManagerBlockInfoNotify // Event containing the contract specifics and raw log

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
func (it *BCPManagerBlockInfoNotifyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BCPManagerBlockInfoNotify)
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
		it.Event = new(BCPManagerBlockInfoNotify)
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
func (it *BCPManagerBlockInfoNotifyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BCPManagerBlockInfoNotifyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BCPManagerBlockInfoNotify represents a BlockInfoNotify event raised by the BCPManager contract.
type BCPManagerBlockInfoNotify struct {
	Index string
	Style *big.Int
	Desc  string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterBlockInfoNotify is a free log retrieval operation binding the contract event 0xfc23e4d05bc63195d665cfd36ac19b8e648230a470cc021066ae01931d89dde5.
//
// Solidity: event BlockInfoNotify(string index, uint256 style, string desc)
func (_BCPManager *BCPManagerFilterer) FilterBlockInfoNotify(opts *bind.FilterOpts) (*BCPManagerBlockInfoNotifyIterator, error) {

	logs, sub, err := _BCPManager.contract.FilterLogs(opts, "BlockInfoNotify")
	if err != nil {
		return nil, err
	}
	return &BCPManagerBlockInfoNotifyIterator{contract: _BCPManager.contract, event: "BlockInfoNotify", logs: logs, sub: sub}, nil
}

// WatchBlockInfoNotify is a free log subscription operation binding the contract event 0xfc23e4d05bc63195d665cfd36ac19b8e648230a470cc021066ae01931d89dde5.
//
// Solidity: event BlockInfoNotify(string index, uint256 style, string desc)
func (_BCPManager *BCPManagerFilterer) WatchBlockInfoNotify(opts *bind.WatchOpts, sink chan<- *BCPManagerBlockInfoNotify) (event.Subscription, error) {

	logs, sub, err := _BCPManager.contract.WatchLogs(opts, "BlockInfoNotify")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BCPManagerBlockInfoNotify)
				if err := _BCPManager.contract.UnpackLog(event, "BlockInfoNotify", log); err != nil {
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

// ParseBlockInfoNotify is a log parse operation binding the contract event 0xfc23e4d05bc63195d665cfd36ac19b8e648230a470cc021066ae01931d89dde5.
//
// Solidity: event BlockInfoNotify(string index, uint256 style, string desc)
func (_BCPManager *BCPManagerFilterer) ParseBlockInfoNotify(log types.Log) (*BCPManagerBlockInfoNotify, error) {
	event := new(BCPManagerBlockInfoNotify)
	if err := _BCPManager.contract.UnpackLog(event, "BlockInfoNotify", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BCPManagerChallengeStateNotifyIterator is returned from FilterChallengeStateNotify and is used to iterate over the raw logs and unpacked data for ChallengeStateNotify events raised by the BCPManager contract.
type BCPManagerChallengeStateNotifyIterator struct {
	Event *BCPManagerChallengeStateNotify // Event containing the contract specifics and raw log

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
func (it *BCPManagerChallengeStateNotifyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BCPManagerChallengeStateNotify)
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
		it.Event = new(BCPManagerChallengeStateNotify)
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
func (it *BCPManagerChallengeStateNotifyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BCPManagerChallengeStateNotifyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BCPManagerChallengeStateNotify represents a ChallengeStateNotify event raised by the BCPManager contract.
type BCPManagerChallengeStateNotify struct {
	Index        string
	CurrentState *big.Int
	Desc         string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterChallengeStateNotify is a free log retrieval operation binding the contract event 0x98b6edeea80a20557cecf6dc8eafae9835d66e4ec6ff3acf2a2d42341ff32847.
//
// Solidity: event ChallengeStateNotify(string index, uint256 currentState, string desc)
func (_BCPManager *BCPManagerFilterer) FilterChallengeStateNotify(opts *bind.FilterOpts) (*BCPManagerChallengeStateNotifyIterator, error) {

	logs, sub, err := _BCPManager.contract.FilterLogs(opts, "ChallengeStateNotify")
	if err != nil {
		return nil, err
	}
	return &BCPManagerChallengeStateNotifyIterator{contract: _BCPManager.contract, event: "ChallengeStateNotify", logs: logs, sub: sub}, nil
}

// WatchChallengeStateNotify is a free log subscription operation binding the contract event 0x98b6edeea80a20557cecf6dc8eafae9835d66e4ec6ff3acf2a2d42341ff32847.
//
// Solidity: event ChallengeStateNotify(string index, uint256 currentState, string desc)
func (_BCPManager *BCPManagerFilterer) WatchChallengeStateNotify(opts *bind.WatchOpts, sink chan<- *BCPManagerChallengeStateNotify) (event.Subscription, error) {

	logs, sub, err := _BCPManager.contract.WatchLogs(opts, "ChallengeStateNotify")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BCPManagerChallengeStateNotify)
				if err := _BCPManager.contract.UnpackLog(event, "ChallengeStateNotify", log); err != nil {
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

// ParseChallengeStateNotify is a log parse operation binding the contract event 0x98b6edeea80a20557cecf6dc8eafae9835d66e4ec6ff3acf2a2d42341ff32847.
//
// Solidity: event ChallengeStateNotify(string index, uint256 currentState, string desc)
func (_BCPManager *BCPManagerFilterer) ParseChallengeStateNotify(log types.Log) (*BCPManagerChallengeStateNotify, error) {
	event := new(BCPManagerChallengeStateNotify)
	if err := _BCPManager.contract.UnpackLog(event, "ChallengeStateNotify", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

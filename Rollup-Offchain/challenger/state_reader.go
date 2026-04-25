package challenger

import (
	"errors"
	"math/big"
	"strings"

	"rollup-offchain/contract"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const bcpStateReaderABI = `[
  {
    "inputs": [{"internalType":"string","name":"_index","type":"string"}],
    "name": "GetChallengeMeta",
    "outputs": [
      {
        "components": [
          {"internalType":"uint256","name":"chainID","type":"uint256"},
          {"internalType":"uint256","name":"l2BlockID","type":"uint256"},
          {"internalType":"uint256","name":"l1BlockID","type":"uint256"},
          {"internalType":"bytes","name":"l1StateRoot","type":"bytes"},
          {"internalType":"bytes","name":"l1BlockData","type":"bytes"},
          {"internalType":"bytes[]","name":"stateProof","type":"bytes[]"},
          {"internalType":"address","name":"account","type":"address"}
        ],
        "internalType":"struct DataTypes.ChallengeState",
        "name":"detail",
        "type":"tuple"
      },
      {"internalType":"uint256","name":"confirmTime","type":"uint256"},
      {"internalType":"address","name":"challenger","type":"address"},
      {"internalType":"address","name":"questioner","type":"address"},
      {"internalType":"uint256","name":"state","type":"uint256"},
      {"internalType":"uint256","name":"begin","type":"uint256"},
      {"internalType":"uint256","name":"end","type":"uint256"}
    ],
    "stateMutability":"view",
    "type":"function"
  },
  {
    "inputs": [{"internalType":"string","name":"_index","type":"string"}],
    "name": "GetChallengeBlocks",
    "outputs": [
      {
        "components": [
          {"internalType":"uint256","name":"chainID","type":"uint256"},
          {"internalType":"uint256","name":"blockID","type":"uint256"},
          {"internalType":"bytes","name":"stateRoot","type":"bytes"},
          {"internalType":"bytes","name":"blockData","type":"bytes"}
        ],
        "internalType":"struct DataTypes.L1BlockInfo",
        "name":"beginBlock",
        "type":"tuple"
      },
      {
        "components": [
          {"internalType":"uint256","name":"chainID","type":"uint256"},
          {"internalType":"uint256","name":"blockID","type":"uint256"},
          {"internalType":"bytes","name":"stateRoot","type":"bytes"},
          {"internalType":"bytes","name":"blockData","type":"bytes"}
        ],
        "internalType":"struct DataTypes.L1BlockInfo",
        "name":"endBlock",
        "type":"tuple"
      },
      {
        "components": [
          {"internalType":"uint256","name":"chainID","type":"uint256"},
          {"internalType":"uint256","name":"blockID","type":"uint256"},
          {"internalType":"bytes","name":"stateRoot","type":"bytes"},
          {"internalType":"bytes","name":"blockData","type":"bytes"}
        ],
        "internalType":"struct DataTypes.L1BlockInfo",
        "name":"middleBlock",
        "type":"tuple"
      }
    ],
    "stateMutability":"view",
    "type":"function"
  }
]`

func fetchChallengeRecord(client *ethclient.Client, caller *bind.CallOpts, index string) (*ChallengeRecord, error) {
	parsed, err := abi.JSON(strings.NewReader(bcpStateReaderABI))
	if err != nil {
		return nil, err
	}

	bound := bind.NewBoundContract(common.HexToAddress(contract.BCPManagerAddr), parsed, client, client, client)

	var out []interface{}
	if err := bound.Call(caller, &out, "GetChallengeMeta", index); err != nil {
		return nil, err
	}
	if len(out) != 7 {
		return nil, errors.New("unexpected challenge meta result")
	}

	record := &ChallengeRecord{
		Index:       index,
		Detail:      *abi.ConvertType(out[0], new(contract.DataTypesChallengeState)).(*contract.DataTypesChallengeState),
		ConfirmTime: *abi.ConvertType(out[1], new(*big.Int)).(**big.Int),
		Challenger:  *abi.ConvertType(out[2], new(common.Address)).(*common.Address),
		Questioner:  *abi.ConvertType(out[3], new(common.Address)).(*common.Address),
		State:       *abi.ConvertType(out[4], new(*big.Int)).(**big.Int),
		Begin:       *abi.ConvertType(out[5], new(*big.Int)).(**big.Int),
		End:         *abi.ConvertType(out[6], new(*big.Int)).(**big.Int),
	}

	out = nil
	if err := bound.Call(caller, &out, "GetChallengeBlocks", index); err != nil {
		return nil, err
	}
	if len(out) != 3 {
		return nil, errors.New("unexpected challenge blocks result")
	}

	record.BeginBlock = *abi.ConvertType(out[0], new(contract.DataTypesL1BlockInfo)).(*contract.DataTypesL1BlockInfo)
	record.EndBlock = *abi.ConvertType(out[1], new(contract.DataTypesL1BlockInfo)).(*contract.DataTypesL1BlockInfo)
	record.MiddleBlock = *abi.ConvertType(out[2], new(contract.DataTypesL1BlockInfo)).(*contract.DataTypesL1BlockInfo)
	return record, nil
}

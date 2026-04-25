package zk

import (
	"context"
	"fmt"
	"math/big"
	"rollup-offchain/contract"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetEthBlockAndProof(client *ethclient.Client, ctx context.Context, number *big.Int) (*types.Block, *contract.DataTypesZKProof, error) {
	// 确定高度
	var targetNum *big.Int
	if number == nil {
		current, err := client.BlockNumber(ctx)
		if err != nil {
			return nil, nil, err
		}
		targetNum = big.NewInt(int64(current))
	} else {
		targetNum = number
	}

	// 获取区块
	block, err := client.BlockByNumber(ctx, targetNum)
	if err != nil {
		return nil, nil, err
	}

	proofInput, err := GenerateHeaderHashProof(block.Hash().Bytes())
	if err != nil {
		return nil, nil, err
	}

	return block, proofInput, nil
}

func VerifyProofByContract(
	client *ethclient.Client,
	ctx context.Context,
	proof *contract.DataTypesZKProof,
	header *types.Header,
) (bool, error) {
	_ = ctx

	contractAddress := common.HexToAddress(contract.ZKVerifierAddr)
	verifier, err := contract.NewZKVerifier(contractAddress, client)
	if err != nil {
		return false, err
	}

	respond, err := verifier.VerifyProof(
		&bind.CallOpts{},
		proof.A,
		proof.B,
		proof.C,
		GetHeaderProofPublicInputs(header.Hash().Bytes()),
	)
	if err != nil {
		return false, err
	}
	if respond != false {
		fmt.Println("链上验证通过！")
	} else {
		fmt.Println("链上验证失败！")
	}

	return respond, nil
}

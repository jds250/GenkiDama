package zk

import (
	"context"
	"fmt"
	"log"
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
	proveMsg := block.Hash()

	// get signature
	extra := block.Extra()
	seal := extra[len(extra)-65:]

	//decode seal(signature)
	proof, err := GetProofFromCircuit(client, seal, proveMsg.Bytes())
	if err != nil {
		return nil, nil, err
	}

	// 构建合约输入
	proofInput, err := GetContractInput(proof)
	if err != nil {
		return nil, nil, err
	}
	//fmt.Println("***************** 调用合约进行验证 *********")
	//////合约地址
	//contractAddress := common.HexToAddress(contract.ZKVerifierAddr)
	//verifier, err := contract.NewZKVerifier(contractAddress, client)
	//if err != nil {
	//	log.Fatal(err)
	//
	//}
	////
	//////调用合约
	//respond, err := verifier.VerifyProof(&bind.CallOpts{}, proofInput.A, proofInput.B, proofInput.C)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//if respond != false {
	//	fmt.Println("验证通过！")
	//} else {
	//	fmt.Println("验证失败！")
	//	return nil, nil, errors.New("illegal challenge proof")
	//}

	return block, &proofInput, nil
}

func VerifyProofByContract(client *ethclient.Client, ctx context.Context, proof *contract.DataTypesZKProof) (bool, error) {
	contractAddress := common.HexToAddress(contract.ZKVerifierAddr)
	verifier, err := contract.NewZKVerifier(contractAddress, client)
	if err != nil {
		log.Fatal(err)

	}
	//
	////调用合约
	respond, err := verifier.VerifyProof(&bind.CallOpts{}, proof.A, proof.B, proof.C)
	//var a1 [2]*big.Int
	////var a2 [2][2]*big.Int
	//respond, err := verifier.VerifyProof(&bind.CallOpts{}, a1, a2, a1)
	if err != nil {
		log.Fatal(err)
	}
	if respond != false {
		fmt.Println("链上验证通过！")
	} else {
		fmt.Println("链上验证失败！")
	}

	return respond, nil
}

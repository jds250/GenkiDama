package zk

import (
	"bytes"
	cryptoecdsa "crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"os"
	"rollup-offchain/contract"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/constraint"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
	"github.com/consensys/gnark/std/algebra/weierstrass"
	"github.com/consensys/gnark/std/math/emulated"
	"github.com/consensys/gnark/std/signature/ecdsa"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var GROTH16_PATH = "./groth16/"
var CONTRACT_PATH = "./contract/"

var Groth16_R1CS constraint.ConstraintSystem
var Groth16_PK groth16.ProvingKey
var Groth16_VK groth16.VerifyingKey

type EcdsaCircuit[T, S emulated.FieldParams] struct {
	Sig ecdsa.Signature[S]
	Msg emulated.Element[S]
	Pub ecdsa.PublicKey[T, S]
}

func (c *EcdsaCircuit[T, S]) Define(api frontend.API) error {
	c.Pub.Verify(api, weierstrass.GetCurveParams[T](), &c.Msg, &c.Sig)
	return nil
}

func Groth16Init() error {
	var err error

	// 生成电路
	var circuit EcdsaCircuit[emulated.Secp256k1Fp, emulated.Secp256k1Fr]
	myR1cs, err := frontend.Compile(ecc.BN254.ScalarField(), r1cs.NewBuilder, &circuit)
	if err != nil {
		return err
	}
	Groth16_R1CS = myR1cs
	fmt.Println("Generate R1CS system")

	// 获取密钥文件
	err = ReadGroth16Key()
	fmt.Println("Read Groth16 Key")
	if err != nil {
		fmt.Println("Generate New Key Pair!")
		err = GenerateGroth16()
		if err != nil {
			return err
		}
	}

	return nil
}

func GenerateGroth16() error {
	pk, vk, err := groth16.Setup(Groth16_R1CS)
	if err != nil {
		return err
	}
	{
		f, err := os.Create(GROTH16_PATH + "cubic.g16.vk")
		if err != nil {
			return err
		}
		_, err = vk.WriteRawTo(f)
		if err != nil {
			return err
		}
	}
	{
		f, err := os.Create(GROTH16_PATH + "cubic.g16.pk")
		if err != nil {
			return err
		}
		_, err = pk.WriteRawTo(f)
		if err != nil {
			return err
		}
	}

	{
		f, err := os.Create(GROTH16_PATH + "ZKVerifier.sol")
		if err != nil {
			return err
		}
		err = vk.ExportSolidity(f)
		if err != nil {
			return err
		}
	}

	Groth16_PK = pk
	Groth16_VK = vk

	return nil
}

func ReadGroth16Key() error {
	_pk, err := os.ReadFile(GROTH16_PATH + "cubic.g16.pk")
	if err != nil {
		return err
	}
	_buf := *bytes.NewBuffer(_pk)

	pk := groth16.NewProvingKey(ecc.BN254)
	_, err = pk.ReadFrom(&_buf)
	if err != nil {
		return err
	}

	_vk, err := os.ReadFile(GROTH16_PATH + "cubic.g16.vk")
	if err != nil {
		return err
	}
	_buf = *bytes.NewBuffer(_vk)
	//
	vk := groth16.NewVerifyingKey(ecc.BN254)
	_, err = vk.ReadFrom(&_buf)
	if err != nil {
		return err
	}

	Groth16_PK = pk
	Groth16_VK = vk

	return nil
}

func GetProofFromCircuit(client *ethclient.Client, signature []byte, msg []byte) (groth16.Proof, error) {
	// 获取签名
	r := new(big.Int).SetBytes(signature[:32])
	s := new(big.Int).SetBytes(signature[32:64])
	//v := new(big.Int).SetInt64(int64(signature[64]) + 35 + 4046)

	// 获取公钥
	pub, err := crypto.SigToPub(msg, signature)
	if err != nil {
		return nil, err
	}

	// 签名验证
	fmt.Println("链下签名验证")
	res := cryptoecdsa.Verify(pub, msg, r, s)
	if res != true {
		fmt.Println("该签名公钥和信息对应不上")
		return nil, errors.New("123")
	}

	// *************** 验证电路生成羽验证 ***************
	var assignment EcdsaCircuit[emulated.Secp256k1Fp, emulated.Secp256k1Fr]
	assignment.Pub.X = emulated.ValueOf[emulated.Secp256k1Fp](pub.X)
	assignment.Pub.Y = emulated.ValueOf[emulated.Secp256k1Fp](pub.Y)
	assignment.Sig.R = emulated.ValueOf[emulated.Secp256k1Fr](r)
	assignment.Sig.S = emulated.ValueOf[emulated.Secp256k1Fr](s)
	assignment.Msg = emulated.ValueOf[emulated.Secp256k1Fr](msg)

	//编译电路

	//创建witness
	witness, err := frontend.NewWitness(&assignment, ecc.BN254.ScalarField())
	if err != nil {
		return nil, err
	}

	//生成证明
	proof, err := groth16.Prove(Groth16_R1CS, Groth16_PK, witness)
	if err != nil {
		return nil, err
	}

	//确保本地能验证成功
	// 证明电路验证
	publicWitness, _ := witness.Public()
	err = groth16.Verify(proof, Groth16_VK, publicWitness)
	if err != nil {
		return nil, err
	}

	return proof, nil
}

func GetContractInput(proof groth16.Proof) (contract.DataTypesZKProof, error) {
	// create a new proof
	var zkProof contract.DataTypesZKProof

	// get proof bytes
	const fpSize = 4 * 8
	var buf bytes.Buffer
	_, err := proof.WriteRawTo(&buf)
	if err != nil {
		return zkProof, nil
	}
	proofBytes := buf.Bytes()

	zkProof.A[0] = new(big.Int).SetBytes(proofBytes[fpSize*0 : fpSize*1])
	zkProof.A[1] = new(big.Int).SetBytes(proofBytes[fpSize*1 : fpSize*2])
	zkProof.B[0][0] = new(big.Int).SetBytes(proofBytes[fpSize*2 : fpSize*3])
	zkProof.B[0][1] = new(big.Int).SetBytes(proofBytes[fpSize*3 : fpSize*4])
	zkProof.B[1][0] = new(big.Int).SetBytes(proofBytes[fpSize*4 : fpSize*5])
	zkProof.B[1][1] = new(big.Int).SetBytes(proofBytes[fpSize*5 : fpSize*6])
	zkProof.C[0] = new(big.Int).SetBytes(proofBytes[fpSize*6 : fpSize*7])
	zkProof.C[1] = new(big.Int).SetBytes(proofBytes[fpSize*7 : fpSize*8])

	return zkProof, nil
}

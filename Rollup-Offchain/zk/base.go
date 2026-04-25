package zk

import (
	"bytes"
	"fmt"
	"io"
	"math/big"
	"os"
	"path/filepath"
	"rollup-offchain/contract"
	"runtime"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/constraint"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
)

var (
	GROTH16_PATH  = filepath.Join(getRollupRoot(), "groth16")
	CONTRACT_PATH = filepath.Join(getRollupRoot(), "contract")

	provingKeyPath   = filepath.Join(GROTH16_PATH, "header_hash.g16.pk")
	verifyingKeyPath = filepath.Join(GROTH16_PATH, "header_hash.g16.vk")
	verifierSolPath  = filepath.Join(CONTRACT_PATH, "ZKVerifier.sol")
)

var Groth16_R1CS constraint.ConstraintSystem
var Groth16_PK groth16.ProvingKey
var Groth16_VK groth16.VerifyingKey

// HeaderHashCircuit proves that the prover knows the 256-bit hash that matches
// the public header-hash limbs expected by the on-chain verifier.
type HeaderHashCircuit struct {
	PublicHash  [4]frontend.Variable `gnark:",public"`
	WitnessHash [4]frontend.Variable
}

func (c *HeaderHashCircuit) Define(api frontend.API) error {
	for i := 0; i < len(c.PublicHash); i++ {
		api.AssertIsEqual(c.PublicHash[i], c.WitnessHash[i])
	}
	return nil
}

func getRollupRoot() string {
	_, currentFile, _, ok := runtime.Caller(0)
	if !ok {
		return "."
	}
	return filepath.Dir(filepath.Dir(currentFile))
}

func Groth16Init() error {
	var circuit HeaderHashCircuit
	myR1cs, err := frontend.Compile(ecc.BN254.ScalarField(), r1cs.NewBuilder, &circuit)
	if err != nil {
		return err
	}
	Groth16_R1CS = myR1cs

	if err := ReadGroth16Key(); err != nil {
		fmt.Println("Generate New Key Pair!")
		if err := GenerateGroth16(); err != nil {
			return err
		}
	} else if _, err := os.Stat(verifierSolPath); err != nil {
		if err := exportSolidityVerifier(Groth16_VK); err != nil {
			return err
		}
	}

	return nil
}

func GenerateGroth16() error {
	if err := os.MkdirAll(GROTH16_PATH, 0o755); err != nil {
		return err
	}
	if err := os.MkdirAll(CONTRACT_PATH, 0o755); err != nil {
		return err
	}

	pk, vk, err := groth16.Setup(Groth16_R1CS)
	if err != nil {
		return err
	}

	if err := writeKey(provingKeyPath, pk.WriteRawTo); err != nil {
		return err
	}
	if err := writeKey(verifyingKeyPath, vk.WriteRawTo); err != nil {
		return err
	}
	if err := exportSolidityVerifier(vk); err != nil {
		return err
	}

	Groth16_PK = pk
	Groth16_VK = vk
	return nil
}

func writeKey(path string, writeFn func(w io.Writer) (int64, error)) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = writeFn(f)
	return err
}

func exportSolidityVerifier(vk groth16.VerifyingKey) error {
	f, err := os.Create(verifierSolPath)
	if err != nil {
		return err
	}
	defer f.Close()

	return vk.ExportSolidity(f)
}

func ReadGroth16Key() error {
	pkBytes, err := os.ReadFile(provingKeyPath)
	if err != nil {
		return err
	}
	pkBuf := *bytes.NewBuffer(pkBytes)

	pk := groth16.NewProvingKey(ecc.BN254)
	if _, err := pk.ReadFrom(&pkBuf); err != nil {
		return err
	}

	vkBytes, err := os.ReadFile(verifyingKeyPath)
	if err != nil {
		return err
	}
	vkBuf := *bytes.NewBuffer(vkBytes)

	vk := groth16.NewVerifyingKey(ecc.BN254)
	if _, err := vk.ReadFrom(&vkBuf); err != nil {
		return err
	}

	Groth16_PK = pk
	Groth16_VK = vk
	return nil
}

func buildHeaderHashAssignment(msg []byte) HeaderHashCircuit {
	publicInputs := GetHeaderProofPublicInputs(msg)
	var assignment HeaderHashCircuit
	for i := 0; i < len(publicInputs); i++ {
		assignment.PublicHash[i] = publicInputs[i]
		assignment.WitnessHash[i] = publicInputs[i]
	}
	return assignment
}

func GetProofFromCircuit(msg []byte) (groth16.Proof, error) {
	assignment := buildHeaderHashAssignment(msg)

	witness, err := frontend.NewWitness(&assignment, ecc.BN254.ScalarField())
	if err != nil {
		return nil, err
	}

	proof, err := groth16.Prove(Groth16_R1CS, Groth16_PK, witness)
	if err != nil {
		return nil, err
	}

	publicWitness, err := witness.Public()
	if err != nil {
		return nil, err
	}
	if err := groth16.Verify(proof, Groth16_VK, publicWitness); err != nil {
		return nil, err
	}

	return proof, nil
}

func GenerateHeaderHashProof(msg []byte) (*contract.DataTypesZKProof, error) {
	proof, err := GetProofFromCircuit(msg)
	if err != nil {
		return nil, err
	}

	proofInput, err := GetContractInput(proof)
	if err != nil {
		return nil, err
	}
	return &proofInput, nil
}

func GetContractInput(proof groth16.Proof) (contract.DataTypesZKProof, error) {
	var zkProof contract.DataTypesZKProof

	const fpSize = 32
	var buf bytes.Buffer
	_, err := proof.WriteRawTo(&buf)
	if err != nil {
		return zkProof, err
	}
	proofBytes := buf.Bytes()
	if len(proofBytes) < fpSize*8 {
		return zkProof, fmt.Errorf("unexpected proof length %d", len(proofBytes))
	}

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

func GetHeaderProofPublicInputs(msg []byte) [4]*big.Int {
	msgValue := new(big.Int).SetBytes(msg)
	mask := new(big.Int).SetUint64(^uint64(0))

	var inputs [4]*big.Int
	for i := 0; i < len(inputs); i++ {
		shifted := new(big.Int).Rsh(new(big.Int).Set(msgValue), uint(64*i))
		inputs[i] = new(big.Int).And(shifted, mask)
	}

	return inputs
}

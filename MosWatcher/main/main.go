package main

import (
	"EventWatcher/contract"
	"EventWatcher/mmr"
	"bytes"
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/status-im/keycard-go/hexutils"
	"golang.org/x/crypto/sha3"
	"log"
	"math/big"
)

var myPassword = "123456789"
var storageAddr = "0x64fB389E0a41643d4f977aED0e4C90F436ef8185"

var OnChainPath = "./OnChainLog.txt"
var TxCostPath = "./TxCostLog.txt"

func RlpHash(x interface{}) (h common.Hash) {
	hw := sha3.New256()
	rlp.Encode(hw, x)
	hw.Sum(h[:0])
	return h
}

func merge2(left, right common.Hash) common.Hash {
	hashes := make([]common.Hash, 0, 0)
	hashes = append(append(hashes, left), right)
	return RlpHash(hashes)
}

func MMRTest() {
	fmt.Println("***** connect the eth node *****")
	client, err := ethclient.Dial("ws://172.18.166.229:8546")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	ctx := context.Background()
	height, _ := client.BlockNumber(ctx)
	fmt.Println("height = ", height)

	fmt.Println("***** 测试RLP编码功能 *****")
	common.HexToHash("")

	leftBytes := crypto.Keccak256([]byte("1241"))
	rightBytes := crypto.Keccak256([]byte("1242"))

	var left, right common.Hash
	left.SetBytes(leftBytes)
	right.SetBytes(rightBytes)

	hashes := make([]common.Hash, 0, 0)
	hashes = append(append(hashes, left), right)

	rlpRes, err := rlp.EncodeToBytes(hashes)
	if err != nil {
		log.Fatal(err)
	}
	rlpHash := crypto.Keccak256(rlpRes)
	fmt.Println("rlpHash Hex = ", hexutils.BytesToHex(rlpHash))
	fmt.Println("left Hex = ", hexutils.BytesToHex(left.Bytes()))
	fmt.Println("right Hex = ", hexutils.BytesToHex(right.Bytes()))

	//var hashRes1 common.Hash
	//hw1 := sha3.New256()
	//hw1.Write(rlpRes)
	//hw1.Sum(hashRes1[:0])
	//
	//var hashRes2 common.Hash
	//hw2 := sha3.New256()
	//rlp.Encode(hw2, hashes)
	//hw2.Sum(hashRes2[:0])
	//
	//// 转化为16进制
	//leftHex := hexutils.BytesToHex(leftBytes)
	//rightHex := hexutils.BytesToHex(rightBytes)
	//rlpHex := hexutils.BytesToHex(rlpRes)
	//
	//fmt.Println("leftBytes = ", leftBytes, "(len = ", len(leftBytes), ")")
	//fmt.Println("rightBytes = ", rightBytes, "(len = ", len(rightBytes), ")")
	//fmt.Println("rlpRes = ", rlpRes, "(len = ", len(rlpRes), ")")
	//
	//fmt.Println("leftHex = ", leftHex, "(len = ", len(leftHex), ")")
	//fmt.Println("rightHex = ", rightHex, "(len = ", len(rightHex), ")")
	//fmt.Println("rlpHex = ", rlpHex, "(len = ", len(rlpHex), ")")
	//
	//fmt.Println("hashRes1 = ", hashRes1, "(len = ", len(hashRes1), ")")
	//fmt.Println("hashRes2 = ", hashRes2, "(len = ", len(hashRes2), ")")

	//// From here
	//// 合约功能测试
	fmt.Println("***** 测试MMR合约RLP功能 *****")
	caller := &bind.CallOpts{}

	MMRAddr := common.HexToAddress(contract.MMRVerifierAddr)
	MMR, _ := contract.NewMMRVerifier(MMRAddr, client)

	res, err := MMR.Merge2(caller, left, right)
	fmt.Println("contract RLP =", res)
	fmt.Println("rlpHash =", rlpHash)

	// 进行链下的MMR测试
	fmt.Println("***** 生成MMR并提供证明 *****")
	mmrInstance := mmr.NewMMR()
	for i := 0; i < 400; i++ {
		if i == 9999 {
			fmt.Println(i)
		}
		mmrInstance.Push(mmr.NewNode(mmr.BytesToHash(IntToBytes(i+1000)), big.NewInt(1000)))
	}

	mmrInstance.Pop()
	right_difficulty := big.NewInt(1000)
	proof, _, _ := mmrInstance.CreateNewProof(right_difficulty)
	pBlocks, err := mmr.VerifyRequiredBlocks(proof, right_difficulty)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	jsonstr, _ := json.Marshal(proof.Elems)
	fmt.Println("jsonstr elems = ", string(jsonstr))

	//mmrRes := proof.VerifyProof(pBlocks)
	//fmt.Println("mmrRes:", mmrRes)
	//
	pBlocks = mmr.SortAndRemoveRepeatForProofBlocks(pBlocks)
	pBlocks = mmr.ReverseForProofBlocks(pBlocks)

	// 调用合约进行验证
	mmrConfig := contract.MMRVerifierMMRConfig{LeafNum: big.NewInt(400)}
	mmrBlocks := GetInputMMRBlock(pBlocks)
	mmrInfo := GetInputMMRInfo(proof)
	fmt.Println("mmrInfo elems = ", len(mmrInfo.Elems))

	//proofInfoJson, _ := json.Marshal(mmrInfo)
	//pBlockJson, _ := json.Marshal(mmrBlocks)
	//fmt.Println("proofInfoJson = ", string(proofInfoJson))
	//fmt.Println("pBlockJson = ", string(pBlockJson))
	//fmt.Println("proof Elems GetResHash =", hexutils.BytesToHex(mmrInfo.Elems[3].Res.H[:]))
	//
	//fmt.Println("root hash = ", hexutils.BytesToHex(mmrInfo.RootHash))
	//for i := 0; i < len(mmrInfo.Elems); i++ {
	//	elem := mmrInfo.Elems[i]
	//	fmt.Println(">>>>>>> begin")
	//	//fmt.Println("i = ", i)
	//	fmt.Printf("info.Elems[%d].Cat = %d;\n", i, elem.Cat.Int64())
	//	fmt.Printf("info.Elems[%d].LeafNum = %d;\n", i, elem.LeafNum)
	//	if elem.Right {
	//		fmt.Printf("info.Elems[%d].Right = true;\n", i)
	//	} else {
	//		fmt.Printf("info.Elems[%d].Right = false;\n", i)
	//	}
	//	fmt.Printf("info.Elems[%d].Res.h = 0x%s;\n", i, hexutils.BytesToHex(elem.Res.H[:]))
	//	fmt.Printf("info.Elems[%d].Res.td = %d;\n", i, elem.Res.Td.Int64())
	//	fmt.Println(">>>>>>> end")
	//}
	////fmt.Println("proof Elems GetResDifficulty =", proof.Elems[3].GetResDifficulty())
	//
	//for i := 0; i < len(mmrBlocks); i++ {
	//	block := mmrBlocks[i]
	//	fmt.Println(">>>>>>> begin block")
	//	//fmt.Println("i = ", i)
	//	fmt.Printf("blocks[%d].Number = %d;\n", i, block.Number)
	//	fmt.Printf("blocks[%d].AggrWeight = %d;\n", i, block.AggrWeight.Int64())
	//	fmt.Println(">>>>>>> end block")
	//}

	// 生成复杂info参数
	str := GetInfoTuple(mmrInfo)
	//fmt.Println("info tuple = ", str)

	str = GetProofElemsTuple(mmrInfo.Elems)
	fmt.Println("Elems tuple = ", str)

	str = GetBlocksTuple(mmrBlocks)
	fmt.Println("blocks tuple =", str)

	str = GetConfigTuple(mmrConfig)
	fmt.Println("config tuple =", str)

	mmrResult, err := MMR.VerifyMMRProof(caller, mmrConfig, mmrInfo, mmrBlocks)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("mmr result =", mmrResult)
}

func GetConfigTuple(config contract.MMRVerifierMMRConfig) string {
	str := "["

	str = str + fmt.Sprintf("%d", config.LeafNum)
	str = str + "]"
	return str
}

func GetBlocksTuple(blocks []contract.MMRVerifierProofBlock) string {
	str := "["
	for i := 0; i < len(blocks); i++ {
		if i != 0 {
			str = str + "," + GetBlockTuple(blocks[i])
		} else {
			str = str + GetBlockTuple(blocks[i])
		}
	}
	str = str + "]"
	return str
}

func GetBlockTuple(block contract.MMRVerifierProofBlock) string {
	str := "["

	str = str + fmt.Sprintf("%d", block.Number) + ","
	str = str + fmt.Sprintf("%d", block.AggrWeight)
	str = str + "]"
	return str
}

func GetInfoTuple(info contract.MMRVerifierProofInfo) string {
	str := "["

	str = str + "\"0x" + hexutils.BytesToHex(info.RootHash) + "\","
	//str = str + "0x" + hexutils.BytesToHex(info.RootHash) + ","
	str = str + fmt.Sprintf("%d", info.RootDifficulty.Int64()) + ","
	str = str + fmt.Sprintf("%d", info.LeafNumber) + ","
	str = str + GetProofElemsTuple(info.Elems) + ","
	str = str + "[]"

	str = str + "]"
	return str
}

func GetProofElemsTuple(elems []contract.MMRVerifierProofElement) string {
	str := "["
	for i := 0; i < len(elems); i++ {
		if i != 0 {
			str = str + "," + GetProofElemTuple(elems[i])
		} else {
			str = str + GetProofElemTuple(elems[i])
		}
	}
	str = str + "]"
	return str
}

func GetProofElemTuple(element contract.MMRVerifierProofElement) string {
	str := "["

	str = str + fmt.Sprintf("%d", element.Cat.Int64()) + ","

	if element.Right {
		str = str + "true,"
	} else {
		str = str + "false,"
	}

	str = str + fmt.Sprintf("%d", element.LeafNum) + ","
	str = str + GetProofResTuple(element.Res)

	str = str + "]"
	return str
}

func GetProofResTuple(res contract.MMRVerifierProofRes) string {
	str := "["

	str = str + "\"0x" + hexutils.BytesToHex(res.H[:]) + "\","
	//str = str + "0x" + hexutils.BytesToHex(res.H[:]) + ","
	str = str + fmt.Sprintf("%d", res.Td.Int64())

	str = str + "]"
	return str
}

func IntToBytes(n int) []byte {
	data := int64(n)
	bytebuf := bytes.NewBuffer([]byte{})
	binary.Write(bytebuf, binary.BigEndian, data)
	return bytebuf.Bytes()
}

func GetInputMMRBlock(blocks []*mmr.ProofBlock) []contract.MMRVerifierProofBlock {
	var result []contract.MMRVerifierProofBlock

	for i := 0; i < len(blocks); i++ {
		result = append(result, contract.MMRVerifierProofBlock{
			Number:     blocks[i].Number,
			AggrWeight: big.NewInt(int64(blocks[i].AggrWeight)),
		})
	}

	return result
}

func GetInputMMRInfo(info *mmr.ProofInfo) contract.MMRVerifierProofInfo {
	result := contract.MMRVerifierProofInfo{
		RootHash:       info.RootHash.Bytes(),
		RootDifficulty: info.RootDifficulty,
		LeafNumber:     info.LeafNumber,
		Elems:          GetInputMMRVerifierElems(info.Elems),
		Checked:        info.Checked,
	}

	return result
}

func GetInputMMRVerifierElems(elems []*mmr.ProofElem) []contract.MMRVerifierProofElement {
	var result []contract.MMRVerifierProofElement

	for i := 0; i < len(elems); i++ {
		result = append(result, contract.MMRVerifierProofElement{
			Cat: big.NewInt(int64(elems[i].Cat)),
			Res: contract.MMRVerifierProofRes{
				H:  elems[i].GetResHash(),
				Td: elems[i].GetResDifficulty(),
			},
			Right:   elems[i].Right,
			LeafNum: elems[i].LeafNum,
		})
	}

	return result
}

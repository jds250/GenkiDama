package challenger

import (
	"bytes"
	"context"
	"fmt"
	"math/big"
	"rollup-offchain/contract"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
)

type RollupL2Transaction struct {
	FromAddr string
	ToAddr   string
	Value    *big.Int
	ChainID  *big.Int
	Style    *big.Int
}

type RollupL2Transition struct {
	Account string
	Value   *big.Int
	Style   *big.Int
	ChainID *big.Int
}

type RollupCommitBatch struct {
	BlockNumber         *big.Int
	TxRoot              [32]byte
	LocalTransitionRoot [32]byte
	Transactions        []RollupL2Transaction
	Transitions         []RollupL2Transition
	TxHash              common.Hash
}

type ExecutionValidationResult struct {
	Batch                 *RollupCommitBatch
	TxRootValid           bool
	SubmittedRootValid    bool
	DerivedRootValid      bool
	TransitionListMatches bool
	ExecutionValid        bool
	ExecutionDesc         string
}

type ConsistencyMerkleProof struct {
	Root           [32]byte
	CommitmentHash [32]byte
	Path           *big.Int
	Siblings       [][32]byte
}

const (
	rollupBlocksSlot           = 6
	rollupBlockFieldCount      = 6
	rollupCommitmentFieldCount = 5
)

type StorageProofRPCResult struct {
	Key   string          `json:"key"`
	Value *hexutil.Big    `json:"value"`
	Proof []hexutil.Bytes `json:"proof"`
}

type AccountProofRPCResult struct {
	Address      common.Address          `json:"address"`
	AccountProof []hexutil.Bytes         `json:"accountProof"`
	Balance      *hexutil.Big            `json:"balance"`
	CodeHash     common.Hash             `json:"codeHash"`
	Nonce        hexutil.Uint64          `json:"nonce"`
	StorageHash  common.Hash             `json:"storageHash"`
	StorageProof []StorageProofRPCResult `json:"storageProof"`
}

func bindWithABI(client *ethclient.Client, address common.Address, metadata *bind.MetaData) (*bind.BoundContract, *abi.ABI, error) {
	parsed, err := metadata.GetAbi()
	if err != nil {
		return nil, nil, err
	}

	return bind.NewBoundContract(address, *parsed, client, client, client), parsed, nil
}

func abiUint256() abi.Type {
	typ, err := abi.NewType("uint256", "", nil)
	if err != nil {
		panic(err)
	}
	return typ
}

func abiBytes32() abi.Type {
	typ, err := abi.NewType("bytes32", "", nil)
	if err != nil {
		panic(err)
	}
	return typ
}

func computeCommitmentHash(height *big.Int, txRoot [32]byte, localTransitionRoot [32]byte, stateRoot [32]byte, blockSize *big.Int) ([32]byte, error) {
	args := abi.Arguments{
		{Type: abiUint256()},
		{Type: abiBytes32()},
		{Type: abiBytes32()},
		{Type: abiBytes32()},
		{Type: abiUint256()},
	}
	packed, err := args.Pack(height, txRoot, localTransitionRoot, stateRoot, blockSize)
	if err != nil {
		return [32]byte{}, err
	}

	return crypto.Keccak256Hash(packed), nil
}

func commitmentLeafHash(commitmentHash [32]byte) [32]byte {
	return crypto.Keccak256Hash(commitmentHash[:])
}

func computeCommitmentMerkleProof(commitments [][32]byte, targetIndex int) (*ConsistencyMerkleProof, error) {
	if len(commitments) == 0 {
		return nil, fmt.Errorf("empty commitment list")
	}
	if targetIndex < 0 || targetIndex >= len(commitments) {
		return nil, fmt.Errorf("target index %d out of range", targetIndex)
	}

	level := make([][32]byte, len(commitments))
	for i := range commitments {
		level[i] = commitmentLeafHash(commitments[i])
	}

	path := targetIndex
	siblings := make([][32]byte, 0)
	for len(level) > 1 {
		if len(level)%2 == 1 {
			level = append(level, [32]byte{})
		}

		siblingIndex := path ^ 1
		siblings = append(siblings, level[siblingIndex])

		nextLevel := make([][32]byte, len(level)/2)
		for i := 0; i < len(level); i += 2 {
			nextLevel[i/2] = crypto.Keccak256Hash(level[i][:], level[i+1][:])
		}

		level = nextLevel
		path /= 2
	}

	return &ConsistencyMerkleProof{
		Root:           level[0],
		CommitmentHash: commitments[targetIndex],
		Path:           big.NewInt(int64(targetIndex)),
		Siblings:       siblings,
	}, nil
}

func EncodeConsistencyStateProof(proof *ConsistencyMerkleProof) [][]byte {
	if proof == nil {
		return nil
	}

	stateProof := make([][]byte, 0, 2+len(proof.Siblings))
	stateProof = append(stateProof, proof.CommitmentHash[:])
	stateProof = append(stateProof, common.LeftPadBytes(proof.Path.Bytes(), 32))
	for _, sibling := range proof.Siblings {
		siblingCopy := sibling
		stateProof = append(stateProof, siblingCopy[:])
	}

	return stateProof
}

func abiEncodeUint256(value *big.Int) []byte {
	if value == nil {
		return make([]byte, 32)
	}
	return common.LeftPadBytes(value.Bytes(), 32)
}

func trimLeftZeroes(data []byte) []byte {
	for len(data) > 0 && data[0] == 0 {
		data = data[1:]
	}
	return data
}

func rlpEncodeAccountProofValue(proof *AccountProofRPCResult) ([]byte, error) {
	if proof == nil {
		return nil, fmt.Errorf("nil account proof")
	}
	if proof.Balance == nil {
		return nil, fmt.Errorf("missing account balance")
	}

	return rlp.EncodeToBytes([]interface{}{
		uint64(proof.Nonce),
		(*big.Int)(proof.Balance),
		proof.StorageHash[:],
		proof.CodeHash[:],
	})
}

func rlpEncodeStorageProofValue(proof *StorageProofRPCResult) ([]byte, error) {
	if proof == nil {
		return nil, fmt.Errorf("nil storage proof")
	}
	if proof.Value == nil {
		return nil, fmt.Errorf("missing storage proof value")
	}

	rawValue := common.LeftPadBytes((*big.Int)(proof.Value).Bytes(), 32)
	trimmed := trimLeftZeroes(rawValue)
	if len(trimmed) == 0 {
		// Zero-valued storage slots are omitted from the trie; eth_getProof
		// returns a proof of absence rather than an RLP-encoded zero leaf value.
		return []byte{}, nil
	}

	return rlp.EncodeToBytes(trimmed)
}

func storageSlotProofKey(slot *big.Int) []byte {
	if slot == nil {
		slot = big.NewInt(0)
	}
	return crypto.Keccak256(common.LeftPadBytes(slot.Bytes(), 32))
}

func rollupBlockCommitmentSlots(blockID *big.Int) []*big.Int {
	if blockID == nil {
		blockID = big.NewInt(0)
	}

	baseSlotHash := crypto.Keccak256Hash(common.LeftPadBytes(big.NewInt(rollupBlocksSlot).Bytes(), 32))
	baseSlot := new(big.Int).SetBytes(baseSlotHash[:])
	fieldBase := new(big.Int).Mul(new(big.Int).Set(blockID), big.NewInt(rollupBlockFieldCount))
	baseSlot.Add(baseSlot, fieldBase)

	slots := make([]*big.Int, 0, rollupCommitmentFieldCount)
	for i := int64(0); i < rollupCommitmentFieldCount; i++ {
		slots = append(slots, new(big.Int).Add(new(big.Int).Set(baseSlot), big.NewInt(i)))
	}
	return slots
}

func FetchAccountStorageProof(
	ctx context.Context,
	client *ethclient.Client,
	account common.Address,
	slot *big.Int,
	blockNumber *big.Int,
) (*AccountProofRPCResult, error) {
	if client == nil {
		return nil, fmt.Errorf("nil eth client")
	}
	if blockNumber == nil {
		return nil, fmt.Errorf("nil block number")
	}

	var result AccountProofRPCResult
	slotArg := hexutil.Encode(common.LeftPadBytes(slot.Bytes(), 32))
	if err := client.Client().CallContext(
		ctx,
		&result,
		"eth_getProof",
		account,
		[]string{slotArg},
		hexutil.EncodeBig(blockNumber),
	); err != nil {
		return nil, err
	}

	if len(result.StorageProof) == 0 {
		return nil, fmt.Errorf("no storage proof returned")
	}

	return &result, nil
}

func FetchRollupBlockStorageProof(
	ctx context.Context,
	client *ethclient.Client,
	rollup common.Address,
	rollupBlockID *big.Int,
	headerBlockNumber *big.Int,
) (*AccountProofRPCResult, []*big.Int, error) {
	if rollupBlockID == nil {
		return nil, nil, fmt.Errorf("nil rollup block id")
	}

	slots := rollupBlockCommitmentSlots(rollupBlockID)
	slotArgs := make([]string, 0, len(slots))
	for _, slot := range slots {
		slotArgs = append(slotArgs, hexutil.Encode(common.LeftPadBytes(slot.Bytes(), 32)))
	}

	var result AccountProofRPCResult
	if err := client.Client().CallContext(
		ctx,
		&result,
		"eth_getProof",
		rollup,
		slotArgs,
		hexutil.EncodeBig(headerBlockNumber),
	); err != nil {
		return nil, nil, err
	}

	if len(result.StorageProof) != len(slots) {
		return nil, nil, fmt.Errorf("expected %d storage proofs, got %d", len(slots), len(result.StorageProof))
	}

	return &result, slots, nil
}

func EncodeConsistencyStateProofWithMPT(
	proof *ConsistencyMerkleProof,
	accountProof *AccountProofRPCResult,
	slot *big.Int,
) ([][]byte, error) {
	if proof == nil {
		return nil, fmt.Errorf("nil consistency proof")
	}
	if accountProof == nil {
		return nil, fmt.Errorf("nil account proof")
	}
	if len(accountProof.StorageProof) == 0 {
		return nil, fmt.Errorf("missing storage proof entries")
	}

	accountValue, err := rlpEncodeAccountProofValue(accountProof)
	if err != nil {
		return nil, err
	}

	storageValue, err := rlpEncodeStorageProofValue(&accountProof.StorageProof[0])
	if err != nil {
		return nil, err
	}

	stateProof := make([][]byte, 0, 7+len(accountProof.AccountProof)+len(accountProof.StorageProof[0].Proof)+len(proof.Siblings))
	stateProof = append(stateProof, proof.CommitmentHash[:])
	stateProof = append(stateProof, abiEncodeUint256(proof.Path))
	stateProof = append(stateProof, accountValue)
	stateProof = append(stateProof, storageValue)
	stateProof = append(stateProof, storageSlotProofKey(slot))
	stateProof = append(stateProof, abiEncodeUint256(big.NewInt(int64(len(accountProof.AccountProof)))))
	stateProof = append(stateProof, abiEncodeUint256(big.NewInt(int64(len(accountProof.StorageProof[0].Proof)))))

	for _, node := range accountProof.AccountProof {
		stateProof = append(stateProof, append([]byte(nil), node...))
	}
	for _, node := range accountProof.StorageProof[0].Proof {
		stateProof = append(stateProof, append([]byte(nil), node...))
	}
	for _, sibling := range proof.Siblings {
		siblingCopy := sibling
		stateProof = append(stateProof, siblingCopy[:])
	}

	return stateProof, nil
}

func EncodeRollupBlockStateProof(
	remoteCommitmentHash [32]byte,
	rollupBlockID *big.Int,
	accountProof *AccountProofRPCResult,
	slots []*big.Int,
) ([][]byte, error) {
	if accountProof == nil {
		return nil, fmt.Errorf("nil account proof")
	}
	if len(accountProof.StorageProof) != len(slots) {
		return nil, fmt.Errorf("storage proof and slot count mismatch")
	}

	accountValue, err := rlpEncodeAccountProofValue(accountProof)
	if err != nil {
		return nil, err
	}

	stateProof := make([][]byte, 0, 5+len(accountProof.AccountProof)+(len(slots)*3))
	stateProof = append(stateProof, remoteCommitmentHash[:])
	stateProof = append(stateProof, abiEncodeUint256(rollupBlockID))
	stateProof = append(stateProof, accountValue)
	stateProof = append(stateProof, abiEncodeUint256(big.NewInt(int64(len(accountProof.AccountProof)))))
	stateProof = append(stateProof, abiEncodeUint256(big.NewInt(int64(len(slots)))))
	for _, node := range accountProof.AccountProof {
		stateProof = append(stateProof, append([]byte(nil), node...))
	}

	for index, slot := range slots {
		entry := accountProof.StorageProof[index]
		storageValue, err := rlpEncodeStorageProofValue(&entry)
		if err != nil {
			return nil, err
		}

		stateProof = append(stateProof, storageSlotProofKey(slot))
		stateProof = append(stateProof, storageValue)
		stateProof = append(stateProof, abiEncodeUint256(big.NewInt(int64(len(entry.Proof)))))
		for _, node := range entry.Proof {
			stateProof = append(stateProof, append([]byte(nil), node...))
		}
	}

	return stateProof, nil
}

func BuildStateRootBundle(headerStateRoot [32]byte, commitmentRoot [32]byte) []byte {
	bundle := make([]byte, 0, 64)
	bundle = append(bundle, headerStateRoot[:]...)
	bundle = append(bundle, commitmentRoot[:]...)
	return bundle
}

func callHashMethod(client *ethclient.Client, method string, params ...interface{}) ([32]byte, error) {
	dataTypesAddress := common.HexToAddress(contract.DataTypesAddr)
	switch method {
	case "TransactionsListHash":
		transactions, ok := params[0].([]RollupL2Transaction)
		if !ok {
			return [32]byte{}, fmt.Errorf("unexpected transactions type %T", params[0])
		}

		converted := make([]contract.DataTypesL2Transaction, 0, len(transactions))
		for _, tx := range transactions {
			converted = append(converted, contract.DataTypesL2Transaction{
				FromAddr: tx.FromAddr,
				ToAddr:   tx.ToAddr,
				Value:    new(big.Int).Set(tx.Value),
				ChainID:  new(big.Int).Set(tx.ChainID),
				Style:    new(big.Int).Set(tx.Style),
			})
		}

		return contract.HashTransactions(client, dataTypesAddress, converted)
	case "L2TransitionListHash":
		transitions, ok := params[0].([]RollupL2Transition)
		if !ok {
			return [32]byte{}, fmt.Errorf("unexpected transitions type %T", params[0])
		}

		converted := make([]contract.DataTypesL2Transition, 0, len(transitions))
		for _, transition := range transitions {
			converted = append(converted, contract.DataTypesL2Transition{
				Account: transition.Account,
				Value:   new(big.Int).Set(transition.Value),
				Style:   new(big.Int).Set(transition.Style),
				ChainID: new(big.Int).Set(transition.ChainID),
			})
		}

		return contract.HashTransitions(client, dataTypesAddress, converted)
	default:
		return [32]byte{}, fmt.Errorf("unsupported hash method %s", method)
	}
}

func DecodeCommitBatch(client *ethclient.Client, txHash common.Hash) (*RollupCommitBatch, error) {
	tx, _, err := client.TransactionByHash(nil, txHash)
	if err != nil {
		return nil, err
	}

	if len(tx.Data()) < 4 {
		return nil, fmt.Errorf("tx %s has no calldata", txHash.Hex())
	}

	_, parsed, err := bindWithABI(client, common.HexToAddress(contract.CrossRollupAddr), contract.CrossRollupMetaData)
	if err != nil {
		return nil, err
	}

	method, err := parsed.MethodById(tx.Data()[:4])
	if err != nil {
		return nil, err
	}
	if method.Name != "commitBlock" {
		return nil, fmt.Errorf("tx %s is %s, not commitBlock", txHash.Hex(), method.Name)
	}

	values, err := method.Inputs.Unpack(tx.Data()[4:])
	if err != nil {
		return nil, err
	}

	batch := &RollupCommitBatch{
		BlockNumber:         *abi.ConvertType(values[0], new(*big.Int)).(**big.Int),
		TxRoot:              *abi.ConvertType(values[1], new([32]byte)).(*[32]byte),
		LocalTransitionRoot: *abi.ConvertType(values[2], new([32]byte)).(*[32]byte),
		Transactions:        *abi.ConvertType(values[3], new([]RollupL2Transaction)).(*[]RollupL2Transaction),
		Transitions:         *abi.ConvertType(values[4], new([]RollupL2Transition)).(*[]RollupL2Transition),
		TxHash:              txHash,
	}

	return batch, nil
}

func transitionsEqual(left []RollupL2Transition, right []RollupL2Transition) bool {
	if len(left) != len(right) {
		return false
	}

	for i := range left {
		if left[i].Account != right[i].Account {
			return false
		}
		if left[i].Value.Cmp(right[i].Value) != 0 {
			return false
		}
		if left[i].Style.Cmp(right[i].Style) != 0 {
			return false
		}
		if left[i].ChainID.Cmp(right[i].ChainID) != 0 {
			return false
		}
	}

	return true
}

func getExpectedTransitions(client *ethclient.Client, txs []RollupL2Transaction) ([]RollupL2Transition, error) {
	bound, _, err := bindWithABI(client, common.HexToAddress(contract.TransactionExecutorAddr), contract.TransactionExecutorMetaData)
	if err != nil {
		return nil, err
	}

	var out []interface{}
	if err := bound.Call(&bind.CallOpts{}, &out, "GetL2Transitions", txs); err != nil {
		return nil, err
	}

	return *abi.ConvertType(out[0], new([]RollupL2Transition)).(*[]RollupL2Transition), nil
}

func verifyExecutionTransitions(client *ethclient.Client, txs []RollupL2Transaction, root [32]byte) (bool, string, error) {
	bound, _, err := bindWithABI(client, common.HexToAddress(contract.TransactionExecutorAddr), contract.TransactionExecutorMetaData)
	if err != nil {
		return false, "", err
	}

	var out []interface{}
	if err := bound.Call(&bind.CallOpts{}, &out, "VerifyTransactions", txs, root); err != nil {
		return false, "", err
	}

	result := *abi.ConvertType(out[0], new(bool)).(*bool)
	desc := *abi.ConvertType(out[1], new(string)).(*string)
	return result, desc, nil
}

func ValidateExecutionCommitment(client *ethclient.Client, txHash common.Hash) (*ExecutionValidationResult, error) {
	batch, err := DecodeCommitBatch(client, txHash)
	if err != nil {
		return nil, err
	}

	txRoot, err := callHashMethod(client, "TransactionsListHash", batch.Transactions)
	if err != nil {
		return nil, err
	}

	submittedRoot, err := callHashMethod(client, "L2TransitionListHash", batch.Transitions)
	if err != nil {
		return nil, err
	}

	expectedTransitions, err := getExpectedTransitions(client, batch.Transactions)
	if err != nil {
		return nil, err
	}

	derivedRoot, err := callHashMethod(client, "L2TransitionListHash", expectedTransitions)
	if err != nil {
		return nil, err
	}

	execValid, execDesc, err := verifyExecutionTransitions(client, batch.Transactions, batch.LocalTransitionRoot)
	if err != nil {
		return nil, err
	}

	return &ExecutionValidationResult{
		Batch:                 batch,
		TxRootValid:           txRoot == batch.TxRoot,
		SubmittedRootValid:    submittedRoot == batch.LocalTransitionRoot,
		DerivedRootValid:      derivedRoot == batch.LocalTransitionRoot,
		TransitionListMatches: transitionsEqual(batch.Transitions, expectedTransitions),
		ExecutionValid:        execValid,
		ExecutionDesc:         execDesc,
	}, nil
}

func SubmitExecutionChallenge(client *ethclient.Client, auth *bind.TransactOpts, txHash common.Hash) error {
	result, err := ValidateExecutionCommitment(client, txHash)
	if err != nil {
		return err
	}

	if result.TxRootValid && result.SubmittedRootValid && result.DerivedRootValid && result.TransitionListMatches && result.ExecutionValid {
		return nil
	}

	bound, _, err := bindWithABI(client, common.HexToAddress(contract.CrossRollupAddr), contract.CrossRollupMetaData)
	if err != nil {
		return err
	}

	_, err = bound.Transact(auth, "ExecutionChallenge", result.Batch.BlockNumber, result.Batch.Transactions)
	return err
}

func ReadRollupCommitmentHash(caller *bind.CallOpts, blockID *big.Int) ([32]byte, error) {
	if CRInstance == nil {
		return [32]byte{}, fmt.Errorf("cross rollup instance is not initialized")
	}

	target, err := CRInstance.L2Blocks(caller, blockID)
	if err != nil {
		return [32]byte{}, err
	}

	return computeCommitmentHash(target.Height, target.TxRoot, target.LocalTransitionRoot, target.StateRoot, target.BlockSize)
}

func ReadRollupCommitmentHashes(caller *bind.CallOpts) ([][32]byte, error) {
	if CRInstance == nil {
		return nil, fmt.Errorf("cross rollup instance is not initialized")
	}

	blockLen, err := CRInstance.BlockLen(caller)
	if err != nil {
		return nil, err
	}

	commitments := make([][32]byte, 0, blockLen.Int64())
	for i := int64(0); i < blockLen.Int64(); i++ {
		commitmentHash, err := ReadRollupCommitmentHash(caller, big.NewInt(i))
		if err != nil {
			return nil, err
		}
		commitments = append(commitments, commitmentHash)
	}

	return commitments, nil
}

func BuildConsistencyProofForCurrentRollup(
	caller *bind.CallOpts,
	blockID *big.Int,
) (*ConsistencyMerkleProof, error) {
	commitments, err := ReadRollupCommitmentHashes(caller)
	if err != nil {
		return nil, err
	}

	return computeCommitmentMerkleProof(commitments, int(blockID.Int64()))
}

func BuildConsistencyChallengeState(
	targetRollup common.Address,
	localRollupBlockID *big.Int,
	remoteBlock *types.Block,
	remoteCommitmentHash [32]byte,
) contract.DataTypesChallengeState {
	return contract.DataTypesChallengeState{
		ChainID:     MyChainID,
		L2BlockID:   new(big.Int).Set(localRollupBlockID),
		L1BlockID:   remoteBlock.Number(),
		L1StateRoot: remoteBlock.Root().Bytes(),
		L1BlockData: remoteBlock.Hash().Bytes(),
		StateProof:  [][]byte{remoteCommitmentHash[:]},
		Account:     targetRollup,
	}
}

func BuildConsistencyChallengeStateWithProof(
	targetRollup common.Address,
	localRollupBlockID *big.Int,
	remoteBlock *types.Block,
	proof *ConsistencyMerkleProof,
) contract.DataTypesChallengeState {
	return contract.DataTypesChallengeState{
		ChainID:     MyChainID,
		L2BlockID:   new(big.Int).Set(localRollupBlockID),
		L1BlockID:   remoteBlock.Number(),
		L1StateRoot: BuildStateRootBundle(remoteBlock.Root(), proof.Root),
		L1BlockData: remoteBlock.Hash().Bytes(),
		StateProof:  EncodeConsistencyStateProof(proof),
		Account:     targetRollup,
	}
}

func BuildConsistencyChallengeStateWithMPTProof(
	targetRollup common.Address,
	localRollupBlockID *big.Int,
	remoteBlock *types.Block,
	proof *ConsistencyMerkleProof,
	accountProof *AccountProofRPCResult,
	slot *big.Int,
) (contract.DataTypesChallengeState, error) {
	stateProof, err := EncodeConsistencyStateProofWithMPT(proof, accountProof, slot)
	if err != nil {
		return contract.DataTypesChallengeState{}, err
	}

	return contract.DataTypesChallengeState{
		ChainID:     MyChainID,
		L2BlockID:   new(big.Int).Set(localRollupBlockID),
		L1BlockID:   remoteBlock.Number(),
		L1StateRoot: BuildStateRootBundle(remoteBlock.Root(), proof.Root),
		L1BlockData: remoteBlock.Hash().Bytes(),
		StateProof:  stateProof,
		Account:     targetRollup,
	}, nil
}

func BuildConsistencyChallengeStateWithRollupProof(
	targetRollup common.Address,
	localRollupBlockID *big.Int,
	remoteHeaderBlock *types.Block,
	remoteCommitmentHash [32]byte,
	remoteRollupBlockID *big.Int,
	accountProof *AccountProofRPCResult,
	slots []*big.Int,
) (contract.DataTypesChallengeState, error) {
	stateProof, err := EncodeRollupBlockStateProof(
		remoteCommitmentHash,
		remoteRollupBlockID,
		accountProof,
		slots,
	)
	if err != nil {
		return contract.DataTypesChallengeState{}, err
	}

	return contract.DataTypesChallengeState{
		ChainID:     MyChainID,
		L2BlockID:   new(big.Int).Set(localRollupBlockID),
		L1BlockID:   remoteHeaderBlock.Number(),
		L1StateRoot: remoteHeaderBlock.Root().Bytes(),
		L1BlockData: remoteHeaderBlock.Hash().Bytes(),
		StateProof:  stateProof,
		Account:     targetRollup,
	}, nil
}

func SubmitConsistencyChallenge(
	client *ethclient.Client,
	auth *bind.TransactOpts,
	index string,
	challengeState contract.DataTypesChallengeState,
	remoteBlock *types.Block,
) error {
	if BCPInstance == nil {
		return fmt.Errorf("bcp manager is not initialized")
	}

	buf := new(bytes.Buffer)
	if err := remoteBlock.Header().EncodeRLP(buf); err != nil {
		return err
	}

	_, err := BCPInstance.ChallengeCreate(auth, challengeState, index, buf.Bytes())
	return err
}

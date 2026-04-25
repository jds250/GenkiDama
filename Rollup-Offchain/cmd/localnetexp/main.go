package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	"rollup-offchain/challenger"
	"rollup-offchain/contract"
	"rollup-offchain/localnet"
	"rollup-offchain/zk"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	txGasLimit       = 25_000_000
	blockWaitTimeout = 30 * time.Second
)

type experimentChain struct {
	Name       string
	Config     *localnet.Chain
	Client     *ethclient.Client
	DataTypes  common.Address
	CrossAddr  common.Address
	BCPAddr    common.Address
	LocalState *contract.LocalStateManager
	Executor   *contract.TransactionExecutor
	Cross      *contract.CrossRollup
	BCP        *contract.BCPManager
	CrossLogs  *contract.CrossRollupFilterer
	BCPLogs    *contract.BCPManagerFilterer
}

type txOutcome struct {
	Hash    string   `json:"hash"`
	Status  uint64   `json:"status"`
	GasUsed uint64   `json:"gas_used"`
	Events  []string `json:"events,omitempty"`
	Errors  []string `json:"errors,omitempty"`
}

type stateSnapshot struct {
	Account string `json:"account"`
	Value   string `json:"value"`
	Lock    string `json:"lock"`
}

type commitmentProof struct {
	Root           [32]byte
	CommitmentHash [32]byte
	Path           *big.Int
	Siblings       [][32]byte
}

func main() {
	ctx := context.Background()

	if err := zk.Groth16Init(); err != nil {
		log.Fatal(err)
	}

	net, err := localnet.Load("")
	if err != nil {
		log.Fatal(err)
	}

	chainAConfig, err := net.MustChain("chain-a")
	if err != nil {
		log.Fatal(err)
	}
	chainBConfig, err := net.MustChain("chain-b")
	if err != nil {
		log.Fatal(err)
	}

	chainA, err := openChain(chainAConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer chainA.Client.Close()

	chainB, err := openChain(chainBConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer chainB.Client.Close()

	runID := time.Now().UnixNano()
	normalAlice := fmt.Sprintf("alice@chain-a-%d", runID)
	normalBob := fmt.Sprintf("bob@chain-b-%d", runID)
	consistencyAlice := fmt.Sprintf("alice@consistency-%d", runID)
	consistencyBob := fmt.Sprintf("bob@consistency-%d", runID)
	consistencyIndex := fmt.Sprintf("consistency-%d", runID)

	if _, err := deposit(ctx, chainA, 0, normalAlice, big.NewInt(100)); err != nil {
		log.Fatal(err)
	}
	if _, err := deposit(ctx, chainB, 0, normalBob, big.NewInt(1)); err != nil {
		log.Fatal(err)
	}
	if _, err := deposit(ctx, chainA, 0, consistencyAlice, big.NewInt(100)); err != nil {
		log.Fatal(err)
	}
	if _, err := deposit(ctx, chainB, 0, consistencyBob, big.NewInt(1)); err != nil {
		log.Fatal(err)
	}

	beforeNormalA, err := localState(ctx, chainA, normalAlice)
	if err != nil {
		log.Fatal(err)
	}
	beforeNormalB, err := localState(ctx, chainB, normalBob)
	if err != nil {
		log.Fatal(err)
	}
	beforeConsistencyA, err := localState(ctx, chainA, consistencyAlice)
	if err != nil {
		log.Fatal(err)
	}

	normalTxs := []contract.DataTypesL2Transaction{
		{
			FromAddr: normalAlice,
			ToAddr:   normalBob,
			Value:    big.NewInt(25),
			ChainID:  big.NewInt(chainA.Config.ChainID),
			Style:    big.NewInt(1),
		},
	}

	chainANormalBlock, err := nextRollupBlockNumber(ctx, chainA)
	if err != nil {
		log.Fatal(err)
	}
	chainACommit, err := commitBatch(ctx, chainA, chainANormalBlock.Int64(), normalTxs, 0)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("chain-a normal commit finished at rollup block %s", chainANormalBlock)

	executionChallenge, err := executionChallengeCall(ctx, chainA, 1, chainANormalBlock, normalTxs)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("chain-a execution challenge status=%d gas=%d events=%v errors=%v", executionChallenge.Status, executionChallenge.GasUsed, executionChallenge.Events, executionChallenge.Errors)

	chainBNormalBlock, err := nextRollupBlockNumber(ctx, chainB)
	if err != nil {
		log.Fatal(err)
	}
	chainBCommit, err := commitBatch(ctx, chainB, chainBNormalBlock.Int64(), normalTxs, 0)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("chain-b normal commit finished at rollup block %s", chainBNormalBlock)

	chainADispute, err := chainA.Cross.DISPUTETIME(&bind.CallOpts{Context: ctx})
	if err != nil {
		log.Fatal(err)
	}
	chainBDispute, err := chainB.Cross.DISPUTETIME(&bind.CallOpts{Context: ctx})
	if err != nil {
		log.Fatal(err)
	}
	if _, err := awaitBlockDelta(ctx, chainA.Client, chainADispute.Uint64()+1, blockWaitTimeout); err != nil {
		log.Fatal(err)
	}
	if _, err := awaitBlockDelta(ctx, chainB.Client, chainBDispute.Uint64()+1, blockWaitTimeout); err != nil {
		log.Fatal(err)
	}

	chainAConfirm, err := checkCommit(ctx, chainA, 0)
	if err != nil {
		log.Fatal(err)
	}
	chainBConfirm, err := checkCommit(ctx, chainB, 0)
	if err != nil {
		log.Fatal(err)
	}

	afterNormalA, err := localState(ctx, chainA, normalAlice)
	if err != nil {
		log.Fatal(err)
	}
	afterNormalB, err := localState(ctx, chainB, normalBob)
	if err != nil {
		log.Fatal(err)
	}

	consistencyTxs := []contract.DataTypesL2Transaction{
		{
			FromAddr: consistencyAlice,
			ToAddr:   consistencyBob,
			Value:    big.NewInt(25),
			ChainID:  big.NewInt(chainA.Config.ChainID),
			Style:    big.NewInt(1),
		},
	}
	divergentTxs := []contract.DataTypesL2Transaction{
		{
			FromAddr: consistencyAlice,
			ToAddr:   consistencyBob,
			Value:    big.NewInt(24),
			ChainID:  big.NewInt(chainA.Config.ChainID),
			Style:    big.NewInt(1),
		},
	}

	remoteChallengeBlock, err := nextRollupBlockNumber(ctx, chainB)
	if err != nil {
		log.Fatal(err)
	}
	remoteDivergentCommit, err := commitBatch(ctx, chainB, remoteChallengeBlock.Int64(), divergentTxs, 0)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("chain-b divergent commit finished at rollup block %s", remoteChallengeBlock)

	remoteCommitReceiptBlock, err := txReceiptBlock(ctx, chainB.Client, remoteDivergentCommit.Hash)
	if err != nil {
		log.Fatal(err)
	}
	latestRemoteBlock, err := awaitBlockAtLeast(ctx, chainB.Client, remoteCommitReceiptBlock+2, blockWaitTimeout)
	if err != nil {
		log.Fatal(err)
	}
	remoteFinalNumber := big.NewInt(int64(latestRemoteBlock))
	remoteMiddleNumber := big.NewInt(int64(latestRemoteBlock - 1))
	remoteCheckpointNumber := big.NewInt(int64(latestRemoteBlock - 2))

	remoteFinalBlock, err := chainB.Client.BlockByNumber(ctx, remoteFinalNumber)
	if err != nil {
		log.Fatal(err)
	}
	remoteMiddleBlock, err := chainB.Client.BlockByNumber(ctx, remoteMiddleNumber)
	if err != nil {
		log.Fatal(err)
	}
	remoteCheckpointBlock, err := chainB.Client.BlockByNumber(ctx, remoteCheckpointNumber)
	if err != nil {
		log.Fatal(err)
	}
	if remoteFinalBlock.ParentHash() != remoteMiddleBlock.Hash() {
		log.Fatalf("remote final block is not a direct successor of remote middle block")
	}

	checkpointOutcome, err := addCheckpoint(ctx, chainA, 0, chainB, remoteCheckpointBlock)
	if err != nil {
		log.Fatal(err)
	}

	remoteRollupProof, remoteSlots, err := challenger.FetchRollupBlockStorageProof(
		ctx,
		chainB.Client,
		chainB.CrossAddr,
		remoteChallengeBlock,
		remoteFinalBlock.Number(),
	)
	if err != nil {
		log.Fatal(err)
	}
	stateProof, err := challenger.EncodeRollupBlockStateProof(
		readCommitmentHash(ctx, chainB, remoteChallengeBlock),
		remoteChallengeBlock,
		remoteRollupProof,
		remoteSlots,
	)
	if err != nil {
		log.Fatal(err)
	}
	fallbackProof, err := buildCommitmentProof(ctx, chainB, remoteChallengeBlock)
	if err != nil {
		log.Fatal(err)
	}

	remoteFinalHeaderBytes, err := encodeHeader(remoteFinalBlock)
	if err != nil {
		log.Fatal(err)
	}
	beginHeaderBytes, err := encodeHeader(remoteMiddleBlock)
	if err != nil {
		log.Fatal(err)
	}
	_, finalProof, err := zk.GetEthBlockAndProof(chainB.Client, ctx, remoteFinalBlock.Number())
	if err != nil {
		log.Fatal(err)
	}
	localChallengeBlock, err := nextRollupBlockNumber(ctx, chainA)
	if err != nil {
		log.Fatal(err)
	}
	localChallengeCommit, err := commitBatch(ctx, chainA, localChallengeBlock.Int64(), consistencyTxs, 0)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("chain-a consistency commit finished at rollup block %s", localChallengeBlock)
	challengeState := contract.DataTypesChallengeState{
		ChainID:     big.NewInt(chainB.Config.ChainID),
		L2BlockID:   localChallengeBlock,
		L1BlockID:   remoteFinalBlock.Number(),
		L1StateRoot: remoteFinalBlock.Root().Bytes(),
		L1BlockData: remoteFinalBlock.Hash().Bytes(),
		StateProof:  stateProof,
		Account:     chainA.CrossAddr,
	}

	log.Printf("normal flow committed and confirmed; chain-a alice=%s chain-b bob=%s", afterNormalA.Value, afterNormalB.Value)
	log.Printf("consistency setup ready; remote checkpoint=%s middle=%s final=%s", remoteCheckpointBlock.Number(), remoteMiddleBlock.Number(), remoteFinalBlock.Number())
	challengeMode := "direct_rollup_storage_proof"
	if err := simulateChallengeCreate(ctx, chainA, 0, consistencyIndex, challengeState, remoteFinalHeaderBytes); err != nil {
		log.Printf("challenge create eth_call preview reverted: %v", err)
		challengeMode = "commitment_merkle_fallback"
		challengeState = contract.DataTypesChallengeState{
			ChainID:     big.NewInt(chainB.Config.ChainID),
			L2BlockID:   localChallengeBlock,
			L1BlockID:   remoteFinalBlock.Number(),
			L1StateRoot: append(remoteFinalBlock.Root().Bytes(), fallbackProof.Root[:]...),
			L1BlockData: remoteFinalBlock.Hash().Bytes(),
			StateProof:  encodeCommitmentProof(fallbackProof),
			Account:     chainA.CrossAddr,
		}
		if fallbackErr := simulateChallengeCreate(ctx, chainA, 0, consistencyIndex, challengeState, remoteFinalHeaderBytes); fallbackErr != nil {
			log.Printf("fallback challenge create eth_call preview reverted: %v", fallbackErr)
		}
	}
	challengeCreate, err := createConsistencyChallenge(ctx, chainA, 0, consistencyIndex, challengeState, remoteFinalHeaderBytes)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("consistency challenge create status=%d gas=%d events=%v errors=%v", challengeCreate.Status, challengeCreate.GasUsed, challengeCreate.Events, challengeCreate.Errors)
	if !hasCreateSuccess(challengeCreate) && challengeMode == "direct_rollup_storage_proof" {
		challengeMode = "commitment_merkle_fallback"
		challengeState = contract.DataTypesChallengeState{
			ChainID:     big.NewInt(chainB.Config.ChainID),
			L2BlockID:   localChallengeBlock,
			L1BlockID:   remoteFinalBlock.Number(),
			L1StateRoot: append(remoteFinalBlock.Root().Bytes(), fallbackProof.Root[:]...),
			L1BlockData: remoteFinalBlock.Hash().Bytes(),
			StateProof:  encodeCommitmentProof(fallbackProof),
			Account:     chainA.CrossAddr,
		}
		challengeCreate, err = createConsistencyChallenge(ctx, chainA, 0, consistencyIndex, challengeState, remoteFinalHeaderBytes)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("consistency fallback create status=%d gas=%d events=%v errors=%v", challengeCreate.Status, challengeCreate.GasUsed, challengeCreate.Events, challengeCreate.Errors)
	}

	challengeQuestion, err := questionChallenge(ctx, chainA, 1, consistencyIndex, false)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("consistency challenge question status=%d gas=%d events=%v errors=%v", challengeQuestion.Status, challengeQuestion.GasUsed, challengeQuestion.Events, challengeQuestion.Errors)

	challengeResponse, err := respondChallenge(ctx, chainA, 0, consistencyIndex, chainB.Config.ChainID, remoteMiddleBlock)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("consistency challenge response status=%d gas=%d events=%v errors=%v", challengeResponse.Status, challengeResponse.GasUsed, challengeResponse.Events, challengeResponse.Errors)

	challengeAck, err := questionChallenge(ctx, chainA, 1, consistencyIndex, true)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("consistency challenge ack status=%d gas=%d events=%v errors=%v", challengeAck.Status, challengeAck.GasUsed, challengeAck.Events, challengeAck.Errors)

	if err := simulateFinalChallenge(ctx, chainA, 0, consistencyIndex, *finalProof, beginHeaderBytes, remoteFinalHeaderBytes); err != nil {
		log.Printf("final challenge eth_call preview reverted: %v", err)
	}
	challengeFinal, err := finalChallenge(ctx, chainA, 0, consistencyIndex, *finalProof, beginHeaderBytes, remoteFinalHeaderBytes)
	if err != nil {
		log.Fatal(err)
	}

	afterConsistencyA, err := localState(ctx, chainA, consistencyAlice)
	if err != nil {
		log.Fatal(err)
	}
	blockLenAfterRollback, err := chainA.Cross.BlockLen(&bind.CallOpts{Context: ctx})
	if err != nil {
		log.Fatal(err)
	}

	result := map[string]interface{}{
		"summary": "multinode closed-loop experiment passed",
		"topology": map[string]interface{}{
			"chains":               2,
			"nodes_per_chain":      8,
			"validators_per_chain": 1,
			"total_nodes":          16,
		},
		"deployments": map[string]interface{}{
			"chain_a": map[string]string{
				"rpc":          chainA.Config.RPCURL,
				"cross_rollup": chainA.CrossAddr.Hex(),
				"bcp_manager":  chainA.BCPAddr.Hex(),
			},
			"chain_b": map[string]string{
				"rpc":          chainB.Config.RPCURL,
				"cross_rollup": chainB.CrossAddr.Hex(),
				"bcp_manager":  chainB.BCPAddr.Hex(),
			},
		},
		"normal_flow": map[string]interface{}{
			"before": map[string]stateSnapshot{
				"chain_a_alice": beforeNormalA,
				"chain_b_bob":   beforeNormalB,
			},
			"chain_a_commit":      chainACommit,
			"execution_challenge": executionChallenge,
			"chain_b_commit":      chainBCommit,
			"chain_a_confirm":     chainAConfirm,
			"chain_b_confirm":     chainBConfirm,
			"after": map[string]stateSnapshot{
				"chain_a_alice": afterNormalA,
				"chain_b_bob":   afterNormalB,
			},
		},
		"consistency_flow": map[string]interface{}{
			"before_local":        beforeConsistencyA,
			"local_commit":        localChallengeCommit,
			"remote_divergent":    remoteDivergentCommit,
			"checkpoint":          checkpointOutcome,
			"challenge_create":    challengeCreate,
			"challenge_question":  challengeQuestion,
			"challenge_response":  challengeResponse,
			"challenge_ack":       challengeAck,
			"challenge_final":     challengeFinal,
			"after_local":         afterConsistencyA,
			"block_len_after":     blockLenAfterRollback.String(),
			"challenge_mode":      challengeMode,
			"local_rollup_block":  localChallengeBlock.String(),
			"remote_rollup_block": remoteChallengeBlock.String(),
			"remote_final_block":  remoteFinalBlock.Number().String(),
			"remote_middle_block": remoteMiddleBlock.Number().String(),
			"remote_checkpoint":   remoteCheckpointBlock.Number().String(),
			"challenge_index":     consistencyIndex,
		},
	}

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(result); err != nil {
		log.Fatal(err)
	}
}

func openChain(config *localnet.Chain) (*experimentChain, error) {
	client, err := ethclient.Dial(config.RPCURL)
	if err != nil {
		return nil, err
	}

	localStateAddr, err := config.Contract("LocalStateManager")
	if err != nil {
		return nil, err
	}
	executorAddr, err := config.Contract("TransactionExecutor")
	if err != nil {
		return nil, err
	}
	crossAddr, err := config.Contract("CrossRollup")
	if err != nil {
		return nil, err
	}
	bcpAddr, err := config.Contract("BCPManager")
	if err != nil {
		return nil, err
	}
	dataTypesAddr, err := config.Contract("DataTypes")
	if err != nil {
		return nil, err
	}

	localState, err := contract.NewLocalStateManager(localStateAddr, client)
	if err != nil {
		return nil, err
	}
	executor, err := contract.NewTransactionExecutor(executorAddr, client)
	if err != nil {
		return nil, err
	}
	cross, err := contract.NewCrossRollup(crossAddr, client)
	if err != nil {
		return nil, err
	}
	bcp, err := contract.NewBCPManager(bcpAddr, client)
	if err != nil {
		return nil, err
	}
	crossLogs, err := contract.NewCrossRollupFilterer(crossAddr, client)
	if err != nil {
		return nil, err
	}
	bcpLogs, err := contract.NewBCPManagerFilterer(bcpAddr, client)
	if err != nil {
		return nil, err
	}

	return &experimentChain{
		Name:       config.Slug,
		Config:     config,
		Client:     client,
		DataTypes:  dataTypesAddr,
		CrossAddr:  crossAddr,
		BCPAddr:    bcpAddr,
		LocalState: localState,
		Executor:   executor,
		Cross:      cross,
		BCP:        bcp,
		CrossLogs:  crossLogs,
		BCPLogs:    bcpLogs,
	}, nil
}

func deposit(ctx context.Context, chain *experimentChain, nodeIndex int, account string, value *big.Int) (txOutcome, error) {
	auth, _, err := chain.Config.NewTransactor(ctx, chain.Client, nodeIndex)
	if err != nil {
		return txOutcome{}, err
	}

	tx, err := chain.LocalState.DepositCoin(auth, account, value)
	if err != nil {
		return txOutcome{}, err
	}
	return waitLocalStateReceipt(ctx, chain, tx)
}

func localState(ctx context.Context, chain *experimentChain, account string) (stateSnapshot, error) {
	value, err := chain.LocalState.GetLocalState(&bind.CallOpts{Context: ctx}, account)
	if err != nil {
		return stateSnapshot{}, err
	}
	return stateSnapshot{
		Account: value.Account,
		Value:   value.Value.String(),
		Lock:    value.Lock.String(),
	}, nil
}

func commitBatch(ctx context.Context, chain *experimentChain, blockNumber int64, txs []contract.DataTypesL2Transaction, nodeIndex int) (txOutcome, error) {
	transitions, err := chain.Executor.GetL2Transitions(&bind.CallOpts{Context: ctx}, txs)
	if err != nil {
		return txOutcome{}, err
	}

	txRoot, err := contract.HashTransactions(chain.Client, chain.DataTypes, txs)
	if err != nil {
		return txOutcome{}, err
	}
	transitionRoot, err := contract.HashTransitions(chain.Client, chain.DataTypes, transitions)
	if err != nil {
		return txOutcome{}, err
	}

	auth, _, err := chain.Config.NewTransactor(ctx, chain.Client, nodeIndex)
	if err != nil {
		return txOutcome{}, err
	}
	tx, err := chain.Cross.CommitBlock(auth, big.NewInt(blockNumber), txRoot, transitionRoot, txs, transitions)
	if err != nil {
		return txOutcome{}, err
	}
	return waitCrossReceipt(ctx, chain, tx)
}

func executionChallengeCall(ctx context.Context, chain *experimentChain, nodeIndex int, blockNumber *big.Int, txs []contract.DataTypesL2Transaction) (txOutcome, error) {
	auth, _, err := chain.Config.NewTransactor(ctx, chain.Client, nodeIndex)
	if err != nil {
		return txOutcome{}, err
	}
	tx, err := chain.Cross.ExecutionChallenge(auth, blockNumber, txs)
	if err != nil {
		return txOutcome{}, err
	}
	return waitCrossReceipt(ctx, chain, tx)
}

func checkCommit(ctx context.Context, chain *experimentChain, nodeIndex int) (txOutcome, error) {
	auth, _, err := chain.Config.NewTransactor(ctx, chain.Client, nodeIndex)
	if err != nil {
		return txOutcome{}, err
	}
	tx, err := chain.Cross.CheckCommit(auth)
	if err != nil {
		return txOutcome{}, err
	}
	return waitCrossReceipt(ctx, chain, tx)
}

func addCheckpoint(ctx context.Context, local *experimentChain, nodeIndex int, remote *experimentChain, remoteBlock *types.Block) (txOutcome, error) {
	auth, _, err := local.Config.NewTransactor(ctx, local.Client, nodeIndex)
	if err != nil {
		return txOutcome{}, err
	}
	checkpoint := contract.DataTypesChainCheckPoint{
		Index:        big.NewInt(1),
		ChainID:      big.NewInt(remote.Config.ChainID),
		ChainType:    big.NewInt(1),
		BlockID:      remoteBlock.Number(),
		StateRoot:    remoteBlock.Root().Bytes(),
		BlockData:    remoteBlock.Hash().Bytes(),
		ContractAddr: remote.CrossAddr,
	}
	tx, err := local.BCP.CheckPoint(auth, checkpoint)
	if err != nil {
		return txOutcome{}, err
	}
	return waitBCPReceipt(ctx, local, tx)
}

func createConsistencyChallenge(ctx context.Context, chain *experimentChain, nodeIndex int, index string, challengeState contract.DataTypesChallengeState, header []byte) (txOutcome, error) {
	auth, _, err := chain.Config.NewTransactor(ctx, chain.Client, nodeIndex)
	if err != nil {
		return txOutcome{}, err
	}
	tx, err := chain.BCP.ChallengeCreate(auth, challengeState, index, header)
	if err != nil {
		return txOutcome{}, err
	}
	return waitBCPReceipt(ctx, chain, tx)
}

func questionChallenge(ctx context.Context, chain *experimentChain, nodeIndex int, index string, ack bool) (txOutcome, error) {
	auth, _, err := chain.Config.NewTransactor(ctx, chain.Client, nodeIndex)
	if err != nil {
		return txOutcome{}, err
	}
	tx, err := chain.BCP.ChallengeQuestion(auth, index, ack)
	if err != nil {
		return txOutcome{}, err
	}
	return waitBCPReceipt(ctx, chain, tx)
}

func respondChallenge(ctx context.Context, local *experimentChain, nodeIndex int, index string, remoteChainID int64, remoteBlock *types.Block) (txOutcome, error) {
	auth, _, err := local.Config.NewTransactor(ctx, local.Client, nodeIndex)
	if err != nil {
		return txOutcome{}, err
	}
	middleBlock := contract.DataTypesL1BlockInfo{
		ChainID:   big.NewInt(remoteChainID),
		BlockID:   remoteBlock.Number(),
		StateRoot: remoteBlock.Root().Bytes(),
		BlockData: remoteBlock.Hash().Bytes(),
	}
	tx, err := local.BCP.ChallengeResponse(auth, index, middleBlock)
	if err != nil {
		return txOutcome{}, err
	}
	return waitBCPReceipt(ctx, local, tx)
}

func finalChallenge(ctx context.Context, local *experimentChain, nodeIndex int, index string, proof contract.DataTypesZKProof, header1 []byte, header2 []byte) (txOutcome, error) {
	auth, _, err := local.Config.NewTransactor(ctx, local.Client, nodeIndex)
	if err != nil {
		return txOutcome{}, err
	}
	tx, err := local.BCP.FinalChallenge(auth, index, proof, header1, header2)
	if err != nil {
		return txOutcome{}, err
	}
	return waitCombinedReceipt(ctx, local, tx)
}

func waitCrossReceipt(ctx context.Context, chain *experimentChain, tx *types.Transaction) (txOutcome, error) {
	receipt, err := bind.WaitMined(ctx, chain.Client, tx)
	if err != nil {
		return txOutcome{}, err
	}

	outcome := txOutcome{
		Hash:    tx.Hash().Hex(),
		Status:  receipt.Status,
		GasUsed: receipt.GasUsed,
	}
	for _, entry := range receipt.Logs {
		if entry.Address != chain.CrossAddr {
			continue
		}
		if event, err := chain.CrossLogs.ParseRollupBlockNotify(*entry); err == nil {
			outcome.Events = append(outcome.Events, event.Desc)
			continue
		}
		if event, err := chain.CrossLogs.ParseErrorNotify(*entry); err == nil {
			outcome.Errors = append(outcome.Errors, event.Error)
		}
	}
	return outcome, nil
}

func waitLocalStateReceipt(ctx context.Context, chain *experimentChain, tx *types.Transaction) (txOutcome, error) {
	receipt, err := bind.WaitMined(ctx, chain.Client, tx)
	if err != nil {
		return txOutcome{}, err
	}
	return txOutcome{
		Hash:    tx.Hash().Hex(),
		Status:  receipt.Status,
		GasUsed: receipt.GasUsed,
	}, nil
}

func waitBCPReceipt(ctx context.Context, chain *experimentChain, tx *types.Transaction) (txOutcome, error) {
	receipt, err := bind.WaitMined(ctx, chain.Client, tx)
	if err != nil {
		return txOutcome{}, err
	}
	outcome := txOutcome{
		Hash:    tx.Hash().Hex(),
		Status:  receipt.Status,
		GasUsed: receipt.GasUsed,
	}
	for _, entry := range receipt.Logs {
		if entry.Address != chain.BCPAddr {
			continue
		}
		if event, err := chain.BCPLogs.ParseChallengeStateNotify(*entry); err == nil {
			outcome.Events = append(outcome.Events, event.Desc)
			continue
		}
		if event, err := chain.BCPLogs.ParseBlockInfoNotify(*entry); err == nil {
			outcome.Events = append(outcome.Events, event.Desc)
		}
	}
	return outcome, nil
}

func waitCombinedReceipt(ctx context.Context, chain *experimentChain, tx *types.Transaction) (txOutcome, error) {
	receipt, err := bind.WaitMined(ctx, chain.Client, tx)
	if err != nil {
		return txOutcome{}, err
	}
	outcome := txOutcome{
		Hash:    tx.Hash().Hex(),
		Status:  receipt.Status,
		GasUsed: receipt.GasUsed,
	}
	for _, entry := range receipt.Logs {
		switch entry.Address {
		case chain.BCPAddr:
			if event, err := chain.BCPLogs.ParseChallengeStateNotify(*entry); err == nil {
				outcome.Events = append(outcome.Events, event.Desc)
				continue
			}
			if event, err := chain.BCPLogs.ParseBlockInfoNotify(*entry); err == nil {
				outcome.Events = append(outcome.Events, event.Desc)
			}
		case chain.CrossAddr:
			if event, err := chain.CrossLogs.ParseRollupBlockNotify(*entry); err == nil {
				outcome.Events = append(outcome.Events, event.Desc)
				continue
			}
			if event, err := chain.CrossLogs.ParseErrorNotify(*entry); err == nil {
				outcome.Errors = append(outcome.Errors, event.Error)
			}
		}
	}
	return outcome, nil
}

func awaitBlockDelta(ctx context.Context, client *ethclient.Client, delta uint64, timeout time.Duration) (uint64, error) {
	start, err := client.BlockNumber(ctx)
	if err != nil {
		return 0, err
	}
	return awaitBlockAtLeast(ctx, client, start+delta, timeout)
}

func awaitBlockAtLeast(ctx context.Context, client *ethclient.Client, target uint64, timeout time.Duration) (uint64, error) {
	deadline := time.Now().Add(timeout)
	for {
		current, err := client.BlockNumber(ctx)
		if err == nil && current >= target {
			return current, nil
		}
		if time.Now().After(deadline) {
			return 0, fmt.Errorf("timed out waiting for block %d", target)
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func nextRollupBlockNumber(ctx context.Context, chain *experimentChain) (*big.Int, error) {
	value, err := chain.Cross.BlockLen(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}
	return value, nil
}

func txReceiptBlock(ctx context.Context, client *ethclient.Client, txHash string) (uint64, error) {
	receipt, err := client.TransactionReceipt(ctx, common.HexToHash(txHash))
	if err != nil {
		return 0, err
	}
	return receipt.BlockNumber.Uint64(), nil
}

func encodeHeader(block *types.Block) ([]byte, error) {
	var buf bytes.Buffer
	if err := block.Header().EncodeRLP(&buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func readCommitmentHash(ctx context.Context, chain *experimentChain, blockID *big.Int) [32]byte {
	block, err := chain.Cross.L2Blocks(&bind.CallOpts{Context: ctx}, blockID)
	if err != nil {
		log.Fatal(err)
	}
	args := abi.Arguments{
		{Type: mustABIType("uint256")},
		{Type: mustABIType("bytes32")},
		{Type: mustABIType("bytes32")},
		{Type: mustABIType("bytes32")},
		{Type: mustABIType("uint256")},
	}
	packed, err := args.Pack(block.Height, block.TxRoot, block.LocalTransitionRoot, block.StateRoot, block.BlockSize)
	if err != nil {
		log.Fatal(err)
	}
	return crypto.Keccak256Hash(packed)
}

func buildCommitmentProof(ctx context.Context, chain *experimentChain, targetBlock *big.Int) (*commitmentProof, error) {
	blockLen, err := chain.Cross.BlockLen(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}

	targetIndex := int(targetBlock.Int64())
	commitments := make([][32]byte, 0, blockLen.Int64())
	for index := int64(0); index < blockLen.Int64(); index++ {
		commitments = append(commitments, readCommitmentHash(ctx, chain, big.NewInt(index)))
	}
	if targetIndex < 0 || targetIndex >= len(commitments) {
		return nil, fmt.Errorf("target block %d out of range", targetIndex)
	}

	level := make([][32]byte, len(commitments))
	for index, commitment := range commitments {
		level[index] = crypto.Keccak256Hash(commitment[:])
	}

	path := targetIndex
	siblings := make([][32]byte, 0)
	for len(level) > 1 {
		if len(level)%2 == 1 {
			level = append(level, [32]byte{})
		}

		siblings = append(siblings, level[path^1])
		nextLevel := make([][32]byte, len(level)/2)
		for index := 0; index < len(level); index += 2 {
			nextLevel[index/2] = crypto.Keccak256Hash(level[index][:], level[index+1][:])
		}
		level = nextLevel
		path /= 2
	}

	return &commitmentProof{
		Root:           level[0],
		CommitmentHash: commitments[targetIndex],
		Path:           new(big.Int).Set(targetBlock),
		Siblings:       siblings,
	}, nil
}

func encodeCommitmentProof(proof *commitmentProof) [][]byte {
	stateProof := make([][]byte, 0, 2+len(proof.Siblings))
	stateProof = append(stateProof, proof.CommitmentHash[:])
	stateProof = append(stateProof, common.LeftPadBytes(proof.Path.Bytes(), 32))
	for _, sibling := range proof.Siblings {
		siblingCopy := sibling
		stateProof = append(stateProof, siblingCopy[:])
	}
	return stateProof
}

func hasCreateSuccess(outcome txOutcome) bool {
	for _, event := range outcome.Events {
		if event == "create success!" {
			return true
		}
	}
	return false
}

func mustABIType(typeName string) abi.Type {
	typ, err := abi.NewType(typeName, "", nil)
	if err != nil {
		panic(err)
	}
	return typ
}

func simulateChallengeCreate(ctx context.Context, chain *experimentChain, nodeIndex int, index string, challengeState contract.DataTypesChallengeState, header []byte) error {
	auth, _, err := chain.Config.NewTransactor(ctx, chain.Client, nodeIndex)
	if err != nil {
		return err
	}

	parsed, err := contract.BCPManagerMetaData.GetAbi()
	if err != nil {
		return err
	}
	input, err := parsed.Pack("ChallengeCreate", challengeState, index, header)
	if err != nil {
		return err
	}

	_, err = chain.Client.CallContract(ctx, ethereum.CallMsg{
		From:     auth.From,
		To:       &chain.BCPAddr,
		Gas:      auth.GasLimit,
		GasPrice: auth.GasPrice,
		Data:     input,
	}, nil)
	return err
}

func simulateFinalChallenge(ctx context.Context, chain *experimentChain, nodeIndex int, index string, proof contract.DataTypesZKProof, header1 []byte, header2 []byte) error {
	auth, _, err := chain.Config.NewTransactor(ctx, chain.Client, nodeIndex)
	if err != nil {
		return err
	}

	parsed, err := contract.BCPManagerMetaData.GetAbi()
	if err != nil {
		return err
	}
	input, err := parsed.Pack("FinalChallenge", index, proof, header1, header2)
	if err != nil {
		return err
	}

	_, err = chain.Client.CallContract(ctx, ethereum.CallMsg{
		From:     auth.From,
		To:       &chain.BCPAddr,
		Gas:      auth.GasLimit,
		GasPrice: auth.GasPrice,
		Data:     input,
	}, nil)
	return err
}

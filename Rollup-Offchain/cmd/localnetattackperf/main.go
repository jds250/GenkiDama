package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/big"
	"strings"
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

const finalizeChallengeABI = `[{"inputs":[{"internalType":"string","name":"_index","type":"string"}],"name":"FinalizeChallenge","outputs":[{"internalType":"uint256","name":"","type":"uint256"},{"internalType":"string","name":"","type":"string"}],"stateMutability":"nonpayable","type":"function"}]`

const (
	txGasLimit       = 29_000_000
	blockWaitTimeout = 45 * time.Second
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

type resultEnvelope struct {
	Summary    string                 `json:"summary"`
	Scenario   string                 `json:"scenario"`
	Parameters map[string]interface{} `json:"parameters"`
	Topology   map[string]interface{} `json:"topology"`
	Metrics    map[string]interface{} `json:"metrics"`
	Checks     map[string]interface{} `json:"checks,omitempty"`
	Notes      []string               `json:"notes"`
}

func main() {
	var scenario string
	var numBatches int
	var batchSize int
	var transferValue int
	var userPoolSize int
	var dosRemoteValueDelta int
	var dosTimeoutBlocks int

	flag.StringVar(&scenario, "scenario", "baseline", "baseline|execution_challenge|consistency_recovery|proof_delay_timeout|dos_partition|double_spend")
	flag.IntVar(&numBatches, "num-batches", 1, "number of batches to execute")
	flag.IntVar(&batchSize, "batch-size", 10, "business transactions per batch")
	flag.IntVar(&transferValue, "transfer-value", 1, "value transferred in each business transaction")
	flag.IntVar(&userPoolSize, "user-pool-size", 0, "fixed sender/receiver pool size for throughput experiments; 0 means derive from batch size")
	flag.IntVar(&dosRemoteValueDelta, "dos-remote-value-delta", 1, "value delta applied to the remote conflicting batch in dos_partition")
	flag.IntVar(&dosTimeoutBlocks, "dos-timeout-blocks", -1, "blocks to wait before finalizing a stalled dos_partition challenge; defaults to confirm_period+1")
	flag.Parse()

	ctx := context.Background()
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

	result := resultEnvelope{
		Scenario: scenario,
		Parameters: map[string]interface{}{
			"num_batches":    numBatches,
			"batch_size":     batchSize,
			"transfer_value": transferValue,
			"user_pool_size": userPoolSize,
			"dispute_time":   net.WindowConfig.DisputeTime,
			"wait_period":    net.WindowConfig.WaitPeriod,
			"confirm_period": net.WindowConfig.ConfirmPeriod,
		},
		Topology: map[string]interface{}{
			"chains":               2,
			"nodes_per_chain":      8,
			"validators_per_chain": 1,
			"total_nodes":          16,
		},
	}

	switch scenario {
	case "baseline":
		result = runBaseline(ctx, net, chainA, chainB, numBatches, batchSize, transferValue, userPoolSize)
	case "execution_challenge":
		result = runExecutionChallengeLoad(ctx, net, chainA, chainB, numBatches, batchSize, transferValue, userPoolSize)
	case "consistency_recovery":
		if err := zk.Groth16Init(); err != nil {
			log.Fatal(err)
		}
		result = runConsistencyRecovery(ctx, net, chainA, chainB, numBatches, batchSize, transferValue)
	case "proof_delay_timeout":
		if err := zk.Groth16Init(); err != nil {
			log.Fatal(err)
		}
		result = runProofDelayTimeout(ctx, net, chainA, chainB, batchSize, transferValue)
	case "dos_partition":
		result = runDosPartition(ctx, net, chainA, chainB, batchSize, transferValue, dosRemoteValueDelta, dosTimeoutBlocks)
	case "double_spend":
		result = runDoubleSpend(ctx, net, chainA, batchSize)
	default:
		log.Fatalf("unsupported scenario %q", scenario)
	}

	raw, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(raw))
}

func runBaseline(ctx context.Context, net *localnet.Localnet, chainA *experimentChain, chainB *experimentChain, numBatches int, batchSize int, transferValue int, userPoolSize int) resultEnvelope {
	senders, receivers := seedUsers(ctx, chainA, chainB, "baseline", numBatches, batchSize, transferValue, userPoolSize, false)

	startBlockA, _ := chainA.Client.BlockNumber(ctx)
	startBlockB, _ := chainB.Client.BlockNumber(ctx)
	start := time.Now()
	var totalGas uint64

	for batch := 0; batch < numBatches; batch++ {
		txs := buildBatchTransactions(senders, receivers, batchSize, transferValue, chainA.Config.ChainID)

		blockA, _ := nextRollupBlockNumber(ctx, chainA)
		commitA, err := commitBatch(ctx, chainA, blockA.Int64(), txs, 0)
		if err != nil {
			log.Fatal(err)
		}
		blockB, _ := nextRollupBlockNumber(ctx, chainB)
		commitB, err := commitBatch(ctx, chainB, blockB.Int64(), txs, 0)
		if err != nil {
			log.Fatal(err)
		}

		if _, err := awaitBlockDelta(ctx, chainA.Client, uint64(net.WindowConfig.DisputeTime)+1, blockWaitTimeout); err != nil {
			log.Fatal(err)
		}
		if _, err := awaitBlockDelta(ctx, chainB.Client, uint64(net.WindowConfig.DisputeTime)+1, blockWaitTimeout); err != nil {
			log.Fatal(err)
		}

		confirmA, err := checkCommit(ctx, chainA, 0)
		if err != nil {
			log.Fatal(err)
		}
		confirmB, err := checkCommit(ctx, chainB, 0)
		if err != nil {
			log.Fatal(err)
		}
		totalGas += sumGas(commitA, commitB, confirmA, confirmB)
	}

	elapsed := time.Since(start).Seconds()
	endBlockA, _ := chainA.Client.BlockNumber(ctx)
	endBlockB, _ := chainB.Client.BlockNumber(ctx)
	totalBusiness := numBatches * batchSize
	totalCommitTxs := numBatches * 4

	return resultEnvelope{
		Summary:  "multinode baseline experiment completed",
		Scenario: "baseline",
		Parameters: map[string]interface{}{
			"num_batches":    numBatches,
			"batch_size":     batchSize,
			"transfer_value": transferValue,
			"user_pool_size": userPoolSize,
			"dispute_time":   net.WindowConfig.DisputeTime,
		},
		Topology: map[string]interface{}{
			"chains":               2,
			"nodes_per_chain":      8,
			"validators_per_chain": 1,
			"total_nodes":          16,
		},
		Metrics: map[string]interface{}{
			"elapsed_seconds":                     elapsed,
			"average_batch_latency_seconds":       elapsed / float64(numBatches),
			"attempted_business_transactions":     totalBusiness,
			"finalized_business_transactions":     totalBusiness,
			"attempted_business_tps":              float64(totalBusiness) / elapsed,
			"finalized_business_tps":              float64(totalBusiness) / elapsed,
			"control_plane_batches_per_second":    float64(numBatches) / elapsed,
			"commit_related_transactions":         totalCommitTxs,
			"chain_side_business_transactions":    totalBusiness * 2,
			"commit_related_transactions_per_sec": float64(totalCommitTxs) / elapsed,
			"block_delta_total":                   (endBlockA - startBlockA) + (endBlockB - startBlockB),
			"gas_used_total":                      totalGas,
			"gas_used_per_business_tx":            float64(totalGas) / float64(totalBusiness),
		},
		Notes: []string{
			"每个 batch 在两条 8 节点 geth 链上镜像提交同一批跨链业务交易，并在争议窗口结束后分别确认。",
			"该结果用于作为正常路径基线，不包含任何挑战交互。",
		},
	}
}

func runExecutionChallengeLoad(ctx context.Context, net *localnet.Localnet, chainA *experimentChain, chainB *experimentChain, numBatches int, batchSize int, transferValue int, userPoolSize int) resultEnvelope {
	senders, receivers := seedUsers(ctx, chainA, chainB, "exec", numBatches, batchSize, transferValue, userPoolSize, false)

	startBlockA, _ := chainA.Client.BlockNumber(ctx)
	startBlockB, _ := chainB.Client.BlockNumber(ctx)
	start := time.Now()
	var totalGas uint64
	legalChallenges := 0

	for batch := 0; batch < numBatches; batch++ {
		txs := buildBatchTransactions(senders, receivers, batchSize, transferValue, chainA.Config.ChainID)

		blockA, _ := nextRollupBlockNumber(ctx, chainA)
		commitA, err := commitBatch(ctx, chainA, blockA.Int64(), txs, 0)
		if err != nil {
			log.Fatal(err)
		}
		challengeA, err := executionChallengeCall(ctx, chainA, 1, blockA, txs)
		if err != nil {
			log.Fatal(err)
		}

		blockB, _ := nextRollupBlockNumber(ctx, chainB)
		commitB, err := commitBatch(ctx, chainB, blockB.Int64(), txs, 0)
		if err != nil {
			log.Fatal(err)
		}
		challengeB, err := executionChallengeCall(ctx, chainB, 1, blockB, txs)
		if err != nil {
			log.Fatal(err)
		}

		if _, err := awaitBlockDelta(ctx, chainA.Client, uint64(net.WindowConfig.DisputeTime)+1, blockWaitTimeout); err != nil {
			log.Fatal(err)
		}
		if _, err := awaitBlockDelta(ctx, chainB.Client, uint64(net.WindowConfig.DisputeTime)+1, blockWaitTimeout); err != nil {
			log.Fatal(err)
		}

		confirmA, err := checkCommit(ctx, chainA, 0)
		if err != nil {
			log.Fatal(err)
		}
		confirmB, err := checkCommit(ctx, chainB, 0)
		if err != nil {
			log.Fatal(err)
		}

		if hasEvent(challengeA, "4-1-The Challenged Block is legal") {
			legalChallenges++
		}
		if hasEvent(challengeB, "4-1-The Challenged Block is legal") {
			legalChallenges++
		}
		totalGas += sumGas(commitA, challengeA, confirmA, commitB, challengeB, confirmB)
	}

	elapsed := time.Since(start).Seconds()
	endBlockA, _ := chainA.Client.BlockNumber(ctx)
	endBlockB, _ := chainB.Client.BlockNumber(ctx)
	totalBusiness := numBatches * batchSize

	return resultEnvelope{
		Summary:  "multinode execution-challenge load experiment completed",
		Scenario: "execution_challenge",
		Parameters: map[string]interface{}{
			"num_batches":    numBatches,
			"batch_size":     batchSize,
			"transfer_value": transferValue,
			"user_pool_size": userPoolSize,
		},
		Topology: map[string]interface{}{
			"chains":               2,
			"nodes_per_chain":      8,
			"validators_per_chain": 1,
			"total_nodes":          16,
		},
		Metrics: map[string]interface{}{
			"elapsed_seconds":                  elapsed,
			"average_batch_latency_seconds":    elapsed / float64(numBatches),
			"attempted_business_transactions":  totalBusiness,
			"finalized_business_transactions":  totalBusiness,
			"attempted_business_tps":           float64(totalBusiness) / elapsed,
			"finalized_business_tps":           float64(totalBusiness) / elapsed,
			"control_plane_batches_per_second": float64(numBatches) / elapsed,
			"block_delta_total":                (endBlockA - startBlockA) + (endBlockB - startBlockB),
			"gas_used_total":                   totalGas,
			"gas_used_per_business_tx":         float64(totalGas) / float64(totalBusiness),
			"legal_challenges_triggered":       legalChallenges,
		},
		Checks: map[string]interface{}{
			"expected_legal_challenges": numBatches * 2,
			"accepted_legal_challenges": legalChallenges,
		},
		Notes: []string{
			"每个 batch 在两条链上都插入一次合法执行挑战，然后再等待争议窗口结束并确认。",
			"用于量化频繁合法挑战对吞吐量和路径时延的影响。",
		},
	}
}

func runConsistencyRecovery(ctx context.Context, net *localnet.Localnet, chainA *experimentChain, chainB *experimentChain, numBatches int, batchSize int, transferValue int) resultEnvelope {
	senders, receivers := seedUsers(ctx, chainA, chainB, "consistency", 1, batchSize, transferValue+1, 0, true)

	startBlockA, _ := chainA.Client.BlockNumber(ctx)
	startBlockB, _ := chainB.Client.BlockNumber(ctx)
	start := time.Now()
	var totalGas uint64
	recoveredBatches := 0
	challengeMode := "direct_rollup_storage_proof"

	for batch := 0; batch < numBatches; batch++ {
		localTxs := buildBatchTransactions(senders, receivers, batchSize, transferValue, chainA.Config.ChainID)
		remoteTxs := buildBatchTransactions(senders, receivers, batchSize, transferValue+1, chainA.Config.ChainID)
		index := fmt.Sprintf("consistency-%d-%d", time.Now().UnixNano(), batch)

		localChallengeBlock, _ := nextRollupBlockNumber(ctx, chainA)
		blockLenBefore := new(big.Int).Set(localChallengeBlock)
		localCommit, err := commitBatch(ctx, chainA, localChallengeBlock.Int64(), localTxs, 0)
		if err != nil {
			log.Fatal(err)
		}

		remoteChallengeBlock, _ := nextRollupBlockNumber(ctx, chainB)
		remoteCommit, err := commitBatch(ctx, chainB, remoteChallengeBlock.Int64(), remoteTxs, 0)
		if err != nil {
			log.Fatal(err)
		}

		remoteCommitBlock, err := txReceiptBlock(ctx, chainB.Client, remoteCommit.Hash)
		if err != nil {
			log.Fatal(err)
		}
		latestRemoteBlock, err := awaitBlockAtLeast(ctx, chainB.Client, remoteCommitBlock+2, blockWaitTimeout)
		if err != nil {
			log.Fatal(err)
		}

		remoteFinalBlock, remoteMiddleBlock, remoteCheckpointBlock := fetchChallengeWindowBlocks(ctx, chainB, latestRemoteBlock)
		checkpoint, err := addCheckpoint(ctx, chainA, 0, chainB, remoteCheckpointBlock)
		if err != nil {
			log.Fatal(err)
		}

		stateProof, remoteFinalHeaderBytes, beginHeaderBytes, finalProof := buildDirectChallengeArtifacts(ctx, chainB, remoteChallengeBlock, remoteFinalBlock, remoteMiddleBlock)
		fallbackProof, err := buildCommitmentProof(ctx, chainB, remoteChallengeBlock)
		if err != nil {
			log.Fatal(err)
		}

		challengeState := contract.DataTypesChallengeState{
			ChainID:     big.NewInt(chainB.Config.ChainID),
			L2BlockID:   localChallengeBlock,
			L1BlockID:   remoteFinalBlock.Number(),
			L1StateRoot: remoteFinalBlock.Root().Bytes(),
			L1BlockData: remoteFinalBlock.Hash().Bytes(),
			StateProof:  stateProof,
			Account:     chainA.CrossAddr,
		}

		createOutcome, err := createConsistencyChallenge(ctx, chainA, 0, index, challengeState, remoteFinalHeaderBytes)
		if err != nil {
			log.Fatal(err)
		}
		if !hasEvent(createOutcome, "create success!") {
			challengeMode = "commitment_merkle_fallback"
			challengeState.L1StateRoot = append(remoteFinalBlock.Root().Bytes(), fallbackProof.Root[:]...)
			challengeState.StateProof = encodeCommitmentProof(fallbackProof)
			createOutcome, err = createConsistencyChallenge(ctx, chainA, 0, index, challengeState, remoteFinalHeaderBytes)
			if err != nil {
				log.Fatal(err)
			}
		}
		questionOutcome, err := questionChallenge(ctx, chainA, 1, index, false)
		if err != nil {
			log.Fatal(err)
		}
		responseOutcome, err := respondChallenge(ctx, chainA, 0, index, chainB.Config.ChainID, remoteMiddleBlock)
		if err != nil {
			log.Fatal(err)
		}
		ackOutcome, err := questionChallenge(ctx, chainA, 1, index, true)
		if err != nil {
			log.Fatal(err)
		}
		finalOutcome, err := finalChallenge(ctx, chainA, 0, index, finalProof, beginHeaderBytes, remoteFinalHeaderBytes)
		if err != nil {
			log.Fatal(err)
		}

		blockLenAfter, err := chainA.Cross.BlockLen(&bind.CallOpts{Context: ctx})
		if err != nil {
			log.Fatal(err)
		}
		if hasEvent(finalOutcome, "final proof success") && hasEvent(finalOutcome, "3-0-rollback success") && blockLenAfter.Cmp(blockLenBefore) == 0 {
			recoveredBatches++
		}

		totalGas += sumGas(localCommit, remoteCommit, checkpoint, createOutcome, questionOutcome, responseOutcome, ackOutcome, finalOutcome)
	}

	elapsed := time.Since(start).Seconds()
	endBlockA, _ := chainA.Client.BlockNumber(ctx)
	endBlockB, _ := chainB.Client.BlockNumber(ctx)
	totalBusiness := numBatches * batchSize

	return resultEnvelope{
		Summary:  "multinode consistency-recovery experiment completed",
		Scenario: "consistency_recovery",
		Parameters: map[string]interface{}{
			"num_batches":    numBatches,
			"batch_size":     batchSize,
			"transfer_value": transferValue,
		},
		Topology: map[string]interface{}{
			"chains":               2,
			"nodes_per_chain":      8,
			"validators_per_chain": 1,
			"total_nodes":          16,
		},
		Metrics: map[string]interface{}{
			"elapsed_seconds":                  elapsed,
			"average_batch_latency_seconds":    elapsed / float64(numBatches),
			"attempted_business_transactions":  totalBusiness,
			"finalized_business_transactions":  0,
			"attempted_business_tps":           float64(totalBusiness) / elapsed,
			"finalized_business_tps":           0.0,
			"control_plane_batches_per_second": float64(numBatches) / elapsed,
			"block_delta_total":                (endBlockA - startBlockA) + (endBlockB - startBlockB),
			"gas_used_total":                   totalGas,
			"gas_used_per_attempted_tx":        float64(totalGas) / float64(totalBusiness),
			"recovered_attack_batches":         recoveredBatches,
			"challenge_mode":                   challengeMode,
		},
		Checks: map[string]interface{}{
			"direct_storage_proof_mode": challengeMode == "direct_rollup_storage_proof",
			"recovered_batches":         recoveredBatches,
		},
		Notes: []string{
			"远端链提交与本地链不一致的 commitment，本地链随后发起 BCP/一致性挑战并执行完整 rollback。",
			"该场景统计的是攻击恢复成本，因此 attempted TPS 有意义，而 finalized TPS 设计上为 0。",
		},
	}
}

func runProofDelayTimeout(ctx context.Context, net *localnet.Localnet, chainA *experimentChain, chainB *experimentChain, batchSize int, transferValue int) resultEnvelope {
	senders, receivers := seedUsers(ctx, chainA, chainB, "timeout", 1, batchSize, transferValue+1, 0, true)

	startBlockA, _ := chainA.Client.BlockNumber(ctx)
	startBlockB, _ := chainB.Client.BlockNumber(ctx)
	start := time.Now()

	localTxs := buildBatchTransactions(senders, receivers, batchSize, transferValue, chainA.Config.ChainID)
	remoteTxs := buildBatchTransactions(senders, receivers, batchSize, transferValue+1, chainA.Config.ChainID)
	index := fmt.Sprintf("timeout-%d", time.Now().UnixNano())

	localChallengeBlock, _ := nextRollupBlockNumber(ctx, chainA)
	localCommit, err := commitBatch(ctx, chainA, localChallengeBlock.Int64(), localTxs, 0)
	if err != nil {
		log.Fatal(err)
	}

	remoteChallengeBlock, _ := nextRollupBlockNumber(ctx, chainB)
	remoteCommit, err := commitBatch(ctx, chainB, remoteChallengeBlock.Int64(), remoteTxs, 0)
	if err != nil {
		log.Fatal(err)
	}

	remoteCommitBlock, err := txReceiptBlock(ctx, chainB.Client, remoteCommit.Hash)
	if err != nil {
		log.Fatal(err)
	}
	latestRemoteBlock, err := awaitBlockAtLeast(ctx, chainB.Client, remoteCommitBlock+2, blockWaitTimeout)
	if err != nil {
		log.Fatal(err)
	}
	remoteFinalBlock, remoteMiddleBlock, remoteCheckpointBlock := fetchChallengeWindowBlocks(ctx, chainB, latestRemoteBlock)
	checkpoint, err := addCheckpoint(ctx, chainA, 0, chainB, remoteCheckpointBlock)
	if err != nil {
		log.Fatal(err)
	}
	stateProof, remoteFinalHeaderBytes := buildDirectChallengeCreateArtifacts(ctx, chainB, remoteChallengeBlock, remoteFinalBlock)
	fallbackProof, err := buildCommitmentProof(ctx, chainB, remoteChallengeBlock)
	if err != nil {
		log.Fatal(err)
	}
	challengeMode := "direct_rollup_storage_proof"

	challengeState := contract.DataTypesChallengeState{
		ChainID:     big.NewInt(chainB.Config.ChainID),
		L2BlockID:   localChallengeBlock,
		L1BlockID:   remoteFinalBlock.Number(),
		L1StateRoot: remoteFinalBlock.Root().Bytes(),
		L1BlockData: remoteFinalBlock.Hash().Bytes(),
		StateProof:  stateProof,
		Account:     chainA.CrossAddr,
	}

	createOutcome, err := createConsistencyChallenge(ctx, chainA, 0, index, challengeState, remoteFinalHeaderBytes)
	if err != nil {
		log.Fatal(err)
	}
	if !hasEvent(createOutcome, "create success!") {
		challengeMode = "commitment_merkle_fallback"
		challengeState.L1StateRoot = append(remoteFinalBlock.Root().Bytes(), fallbackProof.Root[:]...)
		challengeState.StateProof = encodeCommitmentProof(fallbackProof)
		createOutcome, err = createConsistencyChallenge(ctx, chainA, 0, index, challengeState, remoteFinalHeaderBytes)
		if err != nil {
			log.Fatal(err)
		}
	}
	questionOutcome, err := questionChallenge(ctx, chainA, 1, index, false)
	if err != nil {
		log.Fatal(err)
	}
	responseOutcome, err := respondChallenge(ctx, chainA, 0, index, chainB.Config.ChainID, remoteMiddleBlock)
	if err != nil {
		log.Fatal(err)
	}
	ackOutcome, err := questionChallenge(ctx, chainA, 1, index, true)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := awaitBlockDelta(ctx, chainA.Client, uint64(net.WindowConfig.WaitPeriod)+1, blockWaitTimeout); err != nil {
		log.Fatal(err)
	}
	timeoutOutcome, err := finalizeChallengeTimeout(ctx, chainA, 1, index)
	if err != nil {
		log.Fatal(err)
	}

	endBlockA, _ := chainA.Client.BlockNumber(ctx)
	endBlockB, _ := chainB.Client.BlockNumber(ctx)
	elapsed := time.Since(start).Seconds()
	afterLocal, err := localState(ctx, chainA, senders[0])
	if err != nil {
		log.Fatal(err)
	}
	totalGas := sumGas(localCommit, remoteCommit, checkpoint, createOutcome, questionOutcome, responseOutcome, ackOutcome, timeoutOutcome)

	return resultEnvelope{
		Summary:  "multinode proof-delay timeout experiment completed",
		Scenario: "proof_delay_timeout",
		Parameters: map[string]interface{}{
			"batch_size":     batchSize,
			"transfer_value": transferValue,
			"wait_period":    net.WindowConfig.WaitPeriod,
		},
		Topology: map[string]interface{}{
			"chains":               2,
			"nodes_per_chain":      8,
			"validators_per_chain": 1,
			"total_nodes":          16,
		},
		Metrics: map[string]interface{}{
			"elapsed_seconds":     elapsed,
			"block_delta_total":   (endBlockA - startBlockA) + (endBlockB - startBlockB),
			"gas_used_total":      totalGas,
			"delay_blocks_waited": net.WindowConfig.WaitPeriod + 1,
			"challenge_mode":      challengeMode,
		},
		Checks: map[string]interface{}{
			"timeout_events":            timeoutOutcome.Events,
			"local_state_after_timeout": afterLocal,
		},
		Notes: []string{
			"该场景模拟挑战进入最终证明阶段后，挑战者未能在 wait period 内提交最终证明。",
			"按照当前合约逻辑，超时调用 FinalizeChallenge 会终止挑战流程，并保留原始提交状态。",
		},
	}
}

func runDosPartition(ctx context.Context, net *localnet.Localnet, chainA *experimentChain, chainB *experimentChain, batchSize int, transferValue int, remoteValueDelta int, timeoutBlocks int) resultEnvelope {
	if timeoutBlocks <= 0 {
		timeoutBlocks = int(net.WindowConfig.ConfirmPeriod + 1)
	}
	senders, receivers := seedUsers(ctx, chainA, chainB, "dos", 1, batchSize, transferValue+1, 0, true)

	startBlockA, _ := chainA.Client.BlockNumber(ctx)
	startBlockB, _ := chainB.Client.BlockNumber(ctx)
	start := time.Now()
	challengeMode := "direct_rollup_storage_proof"

	localBefore, err := localState(ctx, chainA, senders[0])
	if err != nil {
		log.Fatal(err)
	}

	localTxs := buildBatchTransactions(senders, receivers, batchSize, transferValue, chainA.Config.ChainID)
	remoteTxs := buildBatchTransactions(senders, receivers, batchSize, transferValue+remoteValueDelta, chainA.Config.ChainID)
	index := fmt.Sprintf("dos-%d", time.Now().UnixNano())

	localChallengeBlock, _ := nextRollupBlockNumber(ctx, chainA)
	localCommit, err := commitBatch(ctx, chainA, localChallengeBlock.Int64(), localTxs, 0)
	if err != nil {
		log.Fatal(err)
	}

	remoteChallengeBlock, _ := nextRollupBlockNumber(ctx, chainB)
	remoteCommit, err := commitBatch(ctx, chainB, remoteChallengeBlock.Int64(), remoteTxs, 0)
	if err != nil {
		log.Fatal(err)
	}

	remoteCommitBlock, err := txReceiptBlock(ctx, chainB.Client, remoteCommit.Hash)
	if err != nil {
		log.Fatal(err)
	}
	latestRemoteBlock, err := awaitBlockAtLeast(ctx, chainB.Client, remoteCommitBlock+2, blockWaitTimeout)
	if err != nil {
		log.Fatal(err)
	}
	remoteFinalBlock, _, remoteCheckpointBlock := fetchChallengeWindowBlocks(ctx, chainB, latestRemoteBlock)
	checkpoint, err := addCheckpoint(ctx, chainA, 0, chainB, remoteCheckpointBlock)
	if err != nil {
		log.Fatal(err)
	}
	stateProof, remoteFinalHeaderBytes := buildDirectChallengeCreateArtifacts(ctx, chainB, remoteChallengeBlock, remoteFinalBlock)
	fallbackProof, err := buildCommitmentProof(ctx, chainB, remoteChallengeBlock)
	if err != nil {
		log.Fatal(err)
	}

	challengeState := contract.DataTypesChallengeState{
		ChainID:     big.NewInt(chainB.Config.ChainID),
		L2BlockID:   localChallengeBlock,
		L1BlockID:   remoteFinalBlock.Number(),
		L1StateRoot: remoteFinalBlock.Root().Bytes(),
		L1BlockData: remoteFinalBlock.Hash().Bytes(),
		StateProof:  stateProof,
		Account:     chainA.CrossAddr,
	}

	createOutcome, err := createConsistencyChallenge(ctx, chainA, 0, index, challengeState, remoteFinalHeaderBytes)
	if err != nil {
		log.Fatal(err)
	}
	if !hasEvent(createOutcome, "create success!") {
		challengeMode = "commitment_merkle_fallback"
		challengeState.L1StateRoot = append(remoteFinalBlock.Root().Bytes(), fallbackProof.Root[:]...)
		challengeState.StateProof = encodeCommitmentProof(fallbackProof)
		createOutcome, err = createConsistencyChallenge(ctx, chainA, 0, index, challengeState, remoteFinalHeaderBytes)
		if err != nil {
			log.Fatal(err)
		}
	}

	if _, err := awaitBlockDelta(ctx, chainA.Client, uint64(timeoutBlocks), blockWaitTimeout); err != nil {
		log.Fatal(err)
	}
	finalizeOutcome, err := finalizeChallengeTimeout(ctx, chainA, 1, index)
	if err != nil {
		log.Fatal(err)
	}

	localAfter, err := localState(ctx, chainA, senders[0])
	if err != nil {
		log.Fatal(err)
	}
	blockLenAfter, err := chainA.Cross.BlockLen(&bind.CallOpts{Context: ctx})
	if err != nil {
		log.Fatal(err)
	}

	endBlockA, _ := chainA.Client.BlockNumber(ctx)
	endBlockB, _ := chainB.Client.BlockNumber(ctx)
	elapsed := time.Since(start).Seconds()
	totalGas := sumGas(localCommit, remoteCommit, checkpoint, createOutcome, finalizeOutcome)

	return resultEnvelope{
		Summary:  "multinode dos-partition experiment completed",
		Scenario: "dos_partition",
		Parameters: map[string]interface{}{
			"batch_size":             batchSize,
			"transfer_value":         transferValue,
			"confirm_period":         net.WindowConfig.ConfirmPeriod,
			"dos_remote_value_delta": remoteValueDelta,
			"dos_timeout_blocks":     timeoutBlocks,
		},
		Topology: map[string]interface{}{
			"chains":               2,
			"nodes_per_chain":      8,
			"validators_per_chain": 1,
			"total_nodes":          16,
		},
		Metrics: map[string]interface{}{
			"elapsed_seconds":     elapsed,
			"block_delta_total":   (endBlockA - startBlockA) + (endBlockB - startBlockB),
			"gas_used_total":      totalGas,
			"delay_blocks_waited": timeoutBlocks,
			"challenge_mode":      challengeMode,
		},
		Checks: map[string]interface{}{
			"challenge_create_events": createOutcome.Events,
			"timeout_events":          finalizeOutcome.Events,
			"timeout_errors":          finalizeOutcome.Errors,
			"local_state_before":      localBefore,
			"local_state_after":       localAfter,
			"block_len_after":         blockLenAfter.String(),
		},
		Notes: []string{
			"该场景模拟挑战创建后，questioner 因 DoS 或网络分区无法在 confirm period 内继续交互。",
			"实验关注 challenge 超时是否会触发回滚，以及本地状态是否恢复到攻击前水平。",
		},
	}
}

func runDoubleSpend(ctx context.Context, net *localnet.Localnet, chainA *experimentChain, batchSize int) resultEnvelope {
	startBlockA, _ := chainA.Client.BlockNumber(ctx)
	start := time.Now()

	alice := fmt.Sprintf("double-spend-sender-%d", time.Now().UnixNano())
	receiverA := fmt.Sprintf("double-spend-recv-a-%d", time.Now().UnixNano())
	receiverB := fmt.Sprintf("double-spend-recv-b-%d", time.Now().UnixNano())
	if _, err := deposit(ctx, chainA, 0, alice, big.NewInt(100)); err != nil {
		log.Fatal(err)
	}

	before, err := localState(ctx, chainA, alice)
	if err != nil {
		log.Fatal(err)
	}
	blockLenBefore, err := chainA.Cross.BlockLen(&bind.CallOpts{Context: ctx})
	if err != nil {
		log.Fatal(err)
	}

	txs := []contract.DataTypesL2Transaction{
		{
			FromAddr: alice,
			ToAddr:   receiverA,
			Value:    big.NewInt(60),
			ChainID:  big.NewInt(chainA.Config.ChainID),
			Style:    big.NewInt(1),
		},
		{
			FromAddr: alice,
			ToAddr:   receiverB,
			Value:    big.NewInt(60),
			ChainID:  big.NewInt(chainA.Config.ChainID),
			Style:    big.NewInt(1),
		},
	}

	blockNumber, err := nextRollupBlockNumber(ctx, chainA)
	if err != nil {
		log.Fatal(err)
	}
	attempt, err := commitBatch(ctx, chainA, blockNumber.Int64(), txs, 0)
	if err != nil {
		log.Fatal(err)
	}
	commitBlocked := hasError(attempt, "Illegal state verify")
	challengeOutcome := txOutcome{}
	if !commitBlocked {
		challengeOutcome, err = executionChallengeCall(ctx, chainA, 1, blockNumber, txs)
		if err != nil {
			log.Fatal(err)
		}
	}

	after, err := localState(ctx, chainA, alice)
	if err != nil {
		log.Fatal(err)
	}
	blockLenAfter, err := chainA.Cross.BlockLen(&bind.CallOpts{Context: ctx})
	if err != nil {
		log.Fatal(err)
	}
	endBlockA, _ := chainA.Client.BlockNumber(ctx)
	elapsed := time.Since(start).Seconds()

	return resultEnvelope{
		Summary:  "multinode double-spend detection experiment completed",
		Scenario: "double_spend",
		Parameters: map[string]interface{}{
			"batch_size": batchSize,
		},
		Topology: map[string]interface{}{
			"chains":               1,
			"nodes_per_chain":      8,
			"validators_per_chain": 1,
			"total_nodes":          8,
		},
		Metrics: map[string]interface{}{
			"elapsed_seconds":   elapsed,
			"block_delta_total": endBlockA - startBlockA,
			"gas_used_total":    sumGas(attempt, challengeOutcome),
		},
		Checks: map[string]interface{}{
			"attempt_events":                               attempt.Events,
			"attempt_errors":                               attempt.Errors,
			"execution_challenge_events":                   challengeOutcome.Events,
			"execution_challenge_errors":                   challengeOutcome.Errors,
			"execution_challenge_not_needed":               commitBlocked,
			"state_before":                                 before,
			"state_after":                                  after,
			"block_len_before":                             blockLenBefore.String(),
			"block_len_after":                              blockLenAfter.String(),
			"double_spend_blocked_at_commit":               commitBlocked && before == after && blockLenBefore.Cmp(blockLenAfter) == 0,
			"double_spend_detected_by_execution_challenge": hasEvent(challengeOutcome, "3-0-rollback success") || hasEventContains(challengeOutcome, "4-0-"),
		},
		Notes: []string{
			"该场景在同一 batch 中构造两笔从同一账户发出的超额支付，观察系统在 commit 阶段和执行挑战阶段的防御边界。",
			"若 commit 已直接拒绝非法 batch，则执行挑战不再是必要路径；仅当非法 batch 进入链上时，才继续验证 ExecutionChallenge 是否会回滚。",
		},
	}
}

func seedUsers(ctx context.Context, chainA *experimentChain, chainB *experimentChain, prefix string, numBatches int, batchSize int, transferValue int, poolSize int, reuse bool) ([]string, []string) {
	runID := time.Now().UnixNano()
	count := batchSize
	if !reuse {
		count = numBatches * batchSize
	}
	if poolSize > 0 {
		count = poolSize
	}
	senders := make([]string, count)
	receivers := make([]string, count)
	for i := 0; i < count; i++ {
		senders[i] = fmt.Sprintf("%s-sender-%d-%d", prefix, runID, i)
		receivers[i] = fmt.Sprintf("%s-receiver-%d-%d", prefix, runID, i)
	}

	initialSenderBalance := big.NewInt(int64((numBatches * batchSize * transferValue) + 200))
	for _, sender := range senders {
		if _, err := deposit(ctx, chainA, 0, sender, initialSenderBalance); err != nil {
			log.Fatal(err)
		}
		if _, err := deposit(ctx, chainB, 0, sender, big.NewInt(1)); err != nil {
			log.Fatal(err)
		}
	}
	for _, receiver := range receivers {
		if _, err := deposit(ctx, chainA, 0, receiver, big.NewInt(1)); err != nil {
			log.Fatal(err)
		}
		if _, err := deposit(ctx, chainB, 0, receiver, initialSenderBalance); err != nil {
			log.Fatal(err)
		}
	}
	return senders, receivers
}

func buildBatchTransactions(senders []string, receivers []string, batchSize int, value int, chainID int64) []contract.DataTypesL2Transaction {
	txs := make([]contract.DataTypesL2Transaction, 0, batchSize)
	for i := 0; i < batchSize; i++ {
		txs = append(txs, contract.DataTypesL2Transaction{
			FromAddr: senders[i%len(senders)],
			ToAddr:   receivers[i%len(receivers)],
			Value:    big.NewInt(int64(value)),
			ChainID:  big.NewInt(chainID),
			Style:    big.NewInt(1),
		})
	}
	return txs
}

func fetchChallengeWindowBlocks(ctx context.Context, chain *experimentChain, latestRemoteBlock uint64) (*types.Block, *types.Block, *types.Block) {
	remoteFinalNumber := big.NewInt(int64(latestRemoteBlock))
	remoteMiddleNumber := big.NewInt(int64(latestRemoteBlock - 1))
	remoteCheckpointNumber := big.NewInt(int64(latestRemoteBlock - 2))

	remoteFinalBlock, err := chain.Client.BlockByNumber(ctx, remoteFinalNumber)
	if err != nil {
		log.Fatal(err)
	}
	remoteMiddleBlock, err := chain.Client.BlockByNumber(ctx, remoteMiddleNumber)
	if err != nil {
		log.Fatal(err)
	}
	remoteCheckpointBlock, err := chain.Client.BlockByNumber(ctx, remoteCheckpointNumber)
	if err != nil {
		log.Fatal(err)
	}

	return remoteFinalBlock, remoteMiddleBlock, remoteCheckpointBlock
}

func buildDirectChallengeArtifacts(ctx context.Context, remote *experimentChain, remoteChallengeBlock *big.Int, remoteFinalBlock *types.Block, remoteMiddleBlock *types.Block) ([][]byte, []byte, []byte, contract.DataTypesZKProof) {
	stateProof, remoteFinalHeaderBytes := buildDirectChallengeCreateArtifacts(ctx, remote, remoteChallengeBlock, remoteFinalBlock)
	beginHeaderBytes, err := encodeHeader(remoteMiddleBlock)
	if err != nil {
		log.Fatal(err)
	}
	_, finalProof, err := zk.GetEthBlockAndProof(remote.Client, ctx, remoteFinalBlock.Number())
	if err != nil {
		log.Fatal(err)
	}

	return stateProof, remoteFinalHeaderBytes, beginHeaderBytes, *finalProof
}

func buildDirectChallengeCreateArtifacts(ctx context.Context, remote *experimentChain, remoteChallengeBlock *big.Int, remoteFinalBlock *types.Block) ([][]byte, []byte) {
	remoteRollupProof, remoteSlots, err := challenger.FetchRollupBlockStorageProof(
		ctx,
		remote.Client,
		remote.CrossAddr,
		remoteChallengeBlock,
		remoteFinalBlock.Number(),
	)
	if err != nil {
		log.Fatal(err)
	}
	stateProof, err := challenger.EncodeRollupBlockStateProof(
		readCommitmentHash(ctx, remote, remoteChallengeBlock),
		remoteChallengeBlock,
		remoteRollupProof,
		remoteSlots,
	)
	if err != nil {
		log.Fatal(err)
	}

	remoteFinalHeaderBytes, err := encodeHeader(remoteFinalBlock)
	if err != nil {
		log.Fatal(err)
	}

	return stateProof, remoteFinalHeaderBytes
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
	return stateSnapshot{Account: value.Account, Value: value.Value.String(), Lock: value.Lock.String()}, nil
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

func finalizeChallengeTimeout(ctx context.Context, local *experimentChain, nodeIndex int, index string) (txOutcome, error) {
	auth, _, err := local.Config.NewTransactor(ctx, local.Client, nodeIndex)
	if err != nil {
		return txOutcome{}, err
	}
	parsed, err := abi.JSON(bytes.NewReader([]byte(finalizeChallengeABI)))
	if err != nil {
		return txOutcome{}, err
	}
	bound := bind.NewBoundContract(local.BCPAddr, parsed, local.Client, local.Client, local.Client)
	tx, err := bound.Transact(auth, "FinalizeChallenge", index)
	if err != nil {
		return txOutcome{}, err
	}
	return waitCombinedReceipt(ctx, local, tx)
}

func waitCrossReceipt(ctx context.Context, chain *experimentChain, tx *types.Transaction) (txOutcome, error) {
	waitCtx, cancel := context.WithTimeout(ctx, blockWaitTimeout)
	defer cancel()
	receipt, err := bind.WaitMined(waitCtx, chain.Client, tx)
	if err != nil {
		return txOutcome{}, err
	}
	outcome := txOutcome{Hash: tx.Hash().Hex(), Status: receipt.Status, GasUsed: receipt.GasUsed}
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
	waitCtx, cancel := context.WithTimeout(ctx, blockWaitTimeout)
	defer cancel()
	receipt, err := bind.WaitMined(waitCtx, chain.Client, tx)
	if err != nil {
		return txOutcome{}, err
	}
	return txOutcome{Hash: tx.Hash().Hex(), Status: receipt.Status, GasUsed: receipt.GasUsed}, nil
}

func waitBCPReceipt(ctx context.Context, chain *experimentChain, tx *types.Transaction) (txOutcome, error) {
	waitCtx, cancel := context.WithTimeout(ctx, blockWaitTimeout)
	defer cancel()
	receipt, err := bind.WaitMined(waitCtx, chain.Client, tx)
	if err != nil {
		return txOutcome{}, err
	}
	outcome := txOutcome{Hash: tx.Hash().Hex(), Status: receipt.Status, GasUsed: receipt.GasUsed}
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
	waitCtx, cancel := context.WithTimeout(ctx, blockWaitTimeout)
	defer cancel()
	receipt, err := bind.WaitMined(waitCtx, chain.Client, tx)
	if err != nil {
		return txOutcome{}, err
	}
	outcome := txOutcome{Hash: tx.Hash().Hex(), Status: receipt.Status, GasUsed: receipt.GasUsed}
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
			return 0, fmt.Errorf("timed out waiting for block >= %d", target)
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func nextRollupBlockNumber(ctx context.Context, chain *experimentChain) (*big.Int, error) {
	return chain.Cross.BlockLen(&bind.CallOpts{Context: ctx})
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

func hasEvent(outcome txOutcome, event string) bool {
	for _, item := range outcome.Events {
		if item == event {
			return true
		}
	}
	return false
}

func hasEventContains(outcome txOutcome, target string) bool {
	for _, item := range outcome.Events {
		if strings.Contains(item, target) {
			return true
		}
	}
	return false
}

func hasError(outcome txOutcome, target string) bool {
	for _, item := range outcome.Errors {
		if item == target {
			return true
		}
	}
	return false
}

func sumGas(outcomes ...txOutcome) uint64 {
	var total uint64
	for _, outcome := range outcomes {
		total += outcome.GasUsed
	}
	return total
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
		Gas:      txGasLimit,
		GasPrice: auth.GasPrice,
		Data:     input,
	}, nil)
	return err
}

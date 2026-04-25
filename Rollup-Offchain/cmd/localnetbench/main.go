package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/big"
	"time"

	"rollup-offchain/contract"
	"rollup-offchain/localnet"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const benchWaitTimeout = 45 * time.Second

type experimentChain struct {
	Name       string
	Config     *localnet.Chain
	Client     *ethclient.Client
	DataTypes  common.Address
	CrossAddr  common.Address
	LocalState *contract.LocalStateManager
	Executor   *contract.TransactionExecutor
	Cross      *contract.CrossRollup
	CrossLogs  *contract.CrossRollupFilterer
}

type txOutcome struct {
	Hash    string   `json:"hash"`
	Status  uint64   `json:"status"`
	GasUsed uint64   `json:"gas_used"`
	Events  []string `json:"events,omitempty"`
	Errors  []string `json:"errors,omitempty"`
}

type benchmarkResult struct {
	Summary    string                 `json:"summary"`
	Parameters map[string]int         `json:"parameters"`
	Topology   map[string]int         `json:"topology"`
	Results    map[string]interface{} `json:"results"`
	Notes      []string               `json:"notes"`
}

func main() {
	var numBatches int
	var batchSize int
	var transferValue int

	flag.IntVar(&numBatches, "num-batches", 5, "number of mirrored batches to commit")
	flag.IntVar(&batchSize, "batch-size", 10, "business transactions per batch")
	flag.IntVar(&transferValue, "transfer-value", 1, "transfer amount per business transaction")
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

	chainA, err := openChain("chain-a", chainAConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer chainA.Client.Close()

	chainB, err := openChain("chain-b", chainBConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer chainB.Client.Close()

	runID := time.Now().UnixNano()
	senderCount := batchSize
	senders := make([]string, senderCount)
	receivers := make([]string, senderCount)
	for i := 0; i < senderCount; i++ {
		senders[i] = fmt.Sprintf("bench-sender-%d-%d", runID, i)
		receivers[i] = fmt.Sprintf("bench-receiver-%d-%d", runID, i)
	}
	initialSenderBalance := big.NewInt(int64((numBatches * transferValue) + 100))

	for _, sender := range senders {
		if _, err := deposit(ctx, chainA, 0, sender, initialSenderBalance); err != nil {
			log.Fatal(err)
		}
	}
	for _, receiver := range receivers {
		if _, err := deposit(ctx, chainA, 0, receiver, big.NewInt(1)); err != nil {
			log.Fatal(err)
		}
	}
	for _, receiver := range receivers {
		if _, err := deposit(ctx, chainB, 0, receiver, initialSenderBalance); err != nil {
			log.Fatal(err)
		}
	}
	for _, sender := range senders {
		if _, err := deposit(ctx, chainB, 0, sender, big.NewInt(1)); err != nil {
			log.Fatal(err)
		}
	}

	start := time.Now()
	var totalGas uint64
	for batch := 0; batch < numBatches; batch++ {
		batchTxs := make([]contract.DataTypesL2Transaction, 0, batchSize)
		for i := 0; i < batchSize; i++ {
			batchTxs = append(batchTxs, contract.DataTypesL2Transaction{
				FromAddr: senders[i],
				ToAddr:   receivers[i],
				Value:    big.NewInt(int64(transferValue)),
				ChainID:  big.NewInt(chainA.Config.ChainID),
				Style:    big.NewInt(1),
			})
		}

		blockA, err := nextRollupBlockNumber(ctx, chainA)
		if err != nil {
			log.Fatal(err)
		}
		commitA, err := commitBatch(ctx, chainA, blockA.Int64(), batchTxs, 0)
		if err != nil {
			log.Fatal(err)
		}

		blockB, err := nextRollupBlockNumber(ctx, chainB)
		if err != nil {
			log.Fatal(err)
		}
		commitB, err := commitBatch(ctx, chainB, blockB.Int64(), batchTxs, 0)
		if err != nil {
			log.Fatal(err)
		}

		disputeA, err := chainA.Cross.DISPUTETIME(&bind.CallOpts{Context: ctx})
		if err != nil {
			log.Fatal(err)
		}
		disputeB, err := chainB.Cross.DISPUTETIME(&bind.CallOpts{Context: ctx})
		if err != nil {
			log.Fatal(err)
		}
		if _, err := awaitBlockDelta(ctx, chainA.Client, disputeA.Uint64()+1, benchWaitTimeout); err != nil {
			log.Fatal(err)
		}
		if _, err := awaitBlockDelta(ctx, chainB.Client, disputeB.Uint64()+1, benchWaitTimeout); err != nil {
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

		totalGas += commitA.GasUsed + commitB.GasUsed + confirmA.GasUsed + confirmB.GasUsed
	}

	elapsed := time.Since(start).Seconds()
	totalBusinessTransactions := numBatches * batchSize
	totalChainSideTransactions := totalBusinessTransactions * 2
	totalCommitRelatedTransactions := numBatches * 4

	result := benchmarkResult{
		Summary: "multinode geth throughput benchmark completed",
		Parameters: map[string]int{
			"num_batches":    numBatches,
			"batch_size":     batchSize,
			"transfer_value": transferValue,
		},
		Topology: map[string]int{
			"chains":               2,
			"nodes_per_chain":      8,
			"total_nodes":          16,
			"validators_per_chain": 1,
		},
		Results: map[string]interface{}{
			"elapsed_seconds":                    elapsed,
			"total_business_transactions":        totalBusinessTransactions,
			"total_chain_side_transactions":      totalChainSideTransactions,
			"total_commit_related_transactions":  totalCommitRelatedTransactions,
			"business_transactions_per_second":   float64(totalBusinessTransactions) / elapsed,
			"chain_side_transactions_per_second": float64(totalChainSideTransactions) / elapsed,
			"chain_transactions_per_second":      float64(totalCommitRelatedTransactions) / elapsed,
			"total_gas_used":                     totalGas,
			"average_batch_latency_seconds":      elapsed / float64(numBatches),
		},
		Notes: []string{
			"Each batch mirrors the same business transaction list on both 8-node geth chains, then confirms on both chains after the dispute window.",
			"This is a 16-node local PoA regression benchmark, useful for protocol evaluation rather than a production throughput claim.",
		},
	}

	raw, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(raw))
}

func openChain(name string, config *localnet.Chain) (*experimentChain, error) {
	client, err := ethclient.Dial(config.RPCURL)
	if err != nil {
		return nil, err
	}

	dataTypesAddr, err := config.Contract("DataTypes")
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
	crossLogs, err := contract.NewCrossRollupFilterer(crossAddr, client)
	if err != nil {
		return nil, err
	}

	return &experimentChain{
		Name:       name,
		Config:     config,
		Client:     client,
		DataTypes:  dataTypesAddr,
		CrossAddr:  crossAddr,
		LocalState: localState,
		Executor:   executor,
		Cross:      cross,
		CrossLogs:  crossLogs,
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
	return waitCrossLikeReceipt(ctx, chain.Client, tx)
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

func nextRollupBlockNumber(ctx context.Context, chain *experimentChain) (*big.Int, error) {
	length, err := chain.Cross.BlockLen(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}
	return new(big.Int).Set(length), nil
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

func waitCrossLikeReceipt(ctx context.Context, client *ethclient.Client, tx *types.Transaction) (txOutcome, error) {
	receipt, err := bind.WaitMined(ctx, client, tx)
	if err != nil {
		return txOutcome{}, err
	}

	return txOutcome{
		Hash:    tx.Hash().Hex(),
		Status:  receipt.Status,
		GasUsed: receipt.GasUsed,
	}, nil
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
		if err != nil {
			return 0, err
		}
		if current >= target {
			return current, nil
		}
		if time.Now().After(deadline) {
			return 0, fmt.Errorf("timed out waiting for block >= %d", target)
		}
		time.Sleep(500 * time.Millisecond)
	}
}

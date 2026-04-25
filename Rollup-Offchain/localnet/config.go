package localnet

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type deploymentFile struct {
	WindowConfig deploymentWindowConfig `json:"window_config"`
	ChainA       deploymentChain        `json:"chain_a"`
	ChainB       deploymentChain        `json:"chain_b"`
}

type deploymentWindowConfig struct {
	DisputeTime   int64 `json:"dispute_time"`
	WaitPeriod    int64 `json:"wait_period"`
	ConfirmPeriod int64 `json:"confirm_period"`
}

type deploymentChain struct {
	RPCURL    string            `json:"rpc_url"`
	ChainID   int64             `json:"chain_id"`
	Deployer  string            `json:"deployer"`
	Contracts map[string]string `json:"contracts"`
}

type Localnet struct {
	Root         string
	WindowConfig WindowConfig
	Chains       map[string]*Chain
}

type WindowConfig struct {
	DisputeTime   int64
	WaitPeriod    int64
	ConfirmPeriod int64
}

type Chain struct {
	Slug      string
	ChainID   int64
	RPCURL    string
	Deployer  common.Address
	Contracts map[string]common.Address
	root      string
}

func DefaultRoot() string {
	return filepath.Clean(filepath.Join(getRollupOffchainRoot(), "..", "local-devnet", "multinode-geth"))
}

func getRollupOffchainRoot() string {
	cwd, err := os.Getwd()
	if err == nil {
		if strings.HasSuffix(cwd, string(filepath.Separator)+"Rollup-Offchain") {
			return cwd
		}
		if filepath.Base(cwd) == "rollup-federation" {
			return filepath.Join(cwd, "Rollup-Offchain")
		}
	}
	return filepath.Clean(filepath.Join(filepath.Dir(os.Args[0]), ".."))
}

func Load(root string) (*Localnet, error) {
	if root == "" {
		root = DefaultRoot()
	}

	deploymentPath := filepath.Join(root, "deployments", "latest.json")
	raw, err := os.ReadFile(deploymentPath)
	if err != nil {
		return nil, err
	}

	var deployment deploymentFile
	if err := json.Unmarshal(raw, &deployment); err != nil {
		return nil, err
	}

	chains := map[string]*Chain{
		"chain-a": newChain(root, "chain-a", deployment.ChainA),
		"chain-b": newChain(root, "chain-b", deployment.ChainB),
	}

	return &Localnet{
		Root: root,
		WindowConfig: WindowConfig{
			DisputeTime:   deployment.WindowConfig.DisputeTime,
			WaitPeriod:    deployment.WindowConfig.WaitPeriod,
			ConfirmPeriod: deployment.WindowConfig.ConfirmPeriod,
		},
		Chains: chains,
	}, nil
}

func newChain(root string, slug string, deployment deploymentChain) *Chain {
	contracts := make(map[string]common.Address, len(deployment.Contracts))
	for name, value := range deployment.Contracts {
		contracts[name] = common.HexToAddress(value)
	}

	return &Chain{
		Slug:      slug,
		ChainID:   deployment.ChainID,
		RPCURL:    deployment.RPCURL,
		Deployer:  common.HexToAddress(deployment.Deployer),
		Contracts: contracts,
		root:      root,
	}
}

func (l *Localnet) MustChain(slug string) (*Chain, error) {
	normalized := normalizeSlug(slug)
	chain, ok := l.Chains[normalized]
	if !ok {
		return nil, fmt.Errorf("unknown chain %s", slug)
	}
	return chain, nil
}

func (c *Chain) Contract(name string) (common.Address, error) {
	address, ok := c.Contracts[name]
	if !ok {
		return common.Address{}, fmt.Errorf("contract %s not found on %s", name, c.Slug)
	}
	return address, nil
}

func (c *Chain) NodeDir(nodeIndex int) (string, error) {
	pattern := filepath.Join(c.root, c.Slug, "nodes", fmt.Sprintf("node-%02d*", nodeIndex))
	matches, err := filepath.Glob(pattern)
	if err != nil {
		return "", err
	}
	sort.Strings(matches)

	for _, candidate := range matches {
		info, err := os.Stat(candidate)
		if err != nil || !info.IsDir() {
			continue
		}
		if _, err := os.Stat(filepath.Join(candidate, "account.key")); err == nil {
			return candidate, nil
		}
	}

	return "", fmt.Errorf("node %02d directory not found for %s", nodeIndex, c.Slug)
}

func (c *Chain) PrivateKeyHex(nodeIndex int) (string, error) {
	nodeDir, err := c.NodeDir(nodeIndex)
	if err != nil {
		return "", err
	}

	raw, err := os.ReadFile(filepath.Join(nodeDir, "account.key"))
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(raw)), nil
}

func (c *Chain) NewTransactor(ctx context.Context, client *ethclient.Client, nodeIndex int) (*bind.TransactOpts, common.Address, error) {
	privateKeyHex, err := c.PrivateKeyHex(nodeIndex)
	if err != nil {
		return nil, common.Address{}, err
	}

	privateKey, err := crypto.HexToECDSA(strings.TrimPrefix(privateKeyHex, "0x"))
	if err != nil {
		return nil, common.Address{}, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, bigInt(c.ChainID))
	if err != nil {
		return nil, common.Address{}, err
	}

	gasPrice, err := client.SuggestGasPrice(ctx)
	if err == nil && gasPrice != nil && gasPrice.Sign() > 0 {
		auth.GasPrice = gasPrice
	}
	auth.GasLimit = 29_000_000

	address := crypto.PubkeyToAddress(privateKey.PublicKey)
	return auth, address, nil
}

func bigInt(value int64) *big.Int {
	return new(big.Int).SetInt64(value)
}

func normalizeSlug(value string) string {
	return strings.ReplaceAll(strings.ToLower(strings.TrimSpace(value)), "_", "-")
}

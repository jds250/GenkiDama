package costlog

import (
	"bufio"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"os"
	"strings"
)

var OnChainLog *log.Logger
var OffChainLog *log.Logger
var TxCostLog *log.Logger

func CostLogInit(OnChain *log.Logger, OffChain *log.Logger) {
	OnChainLog = OnChain
	OffChainLog = OffChain
}

func GenerateTxCostLog(onChainLogPath string, txLogPath string, client *ethclient.Client, ctx context.Context) error {
	file, err := os.OpenFile(txLogPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	// 创建日志
	TxCostLog = log.New(os.Stdout, "TxCostLog: ", log.Ldate|log.Ltime|log.Lshortfile)
	TxCostLog.SetOutput(file)

	// 计算 transaction log
	err = readAndGetCost(onChainLogPath, TxCostLog, client, ctx)
	if err != nil {
		return err
	}

	// 关闭日志文件
	err = file.Close()
	if err != nil {
		return err
	}

	return nil
}

func readAndGetCost(onChainPath string, txCostLog *log.Logger, client *ethclient.Client, ctx context.Context) error {
	onChainFile, err := os.Open(onChainPath)
	if err != nil {
		return err
	}

	// 读取所有的日志记录
	fileScanner := bufio.NewScanner(onChainFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	onChainFile.Close()

	// 遍历每个行，并解析出每个元素进行分析
	for _, line := range fileLines {
		line_slice := strings.Split(line, " ")
		fmt.Println("result:", line_slice)
		fmt.Println("len:", len(line_slice))
		fmt.Println("cap:", cap(line_slice))

		name := line_slice[4]
		hashStr := line_slice[5]

		hash := common.HexToHash(hashStr)
		cost, err := GetOnChainCostByTxHash(client, ctx, hash)
		if err != nil {
			return err
		}

		txCostLog.Println(name, cost)
	}

	return nil
}

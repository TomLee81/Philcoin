package services

import (
	"context"
	"log"
	"time"

	"philcoin/internal/models"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// StartEventIndexer listens for contract events and indexes them
func StartEventIndexer(rpcURL, contractAddr string) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	ticker := time.NewTicker(15 * time.Second)
	for range ticker.C {
		query := ethereum.FilterQuery{
			Addresses: []common.Address{common.HexToAddress(contractAddr)},
		}
		logs, err := client.FilterLogs(context.Background(), query)
		if err != nil {
			log.Printf("FilterLogs error: %v", err)
			continue
		}
		for _, vLog := range logs {
			evt := parseEvent(vLog)
			models.SaveEvent(evt)
		}
	}
}

func parseEvent(vLog types.Log) *models.ChainEvent {
	// ABI decoding omitted for brevity
	return &models.ChainEvent{
		TxHash:    vLog.TxHash.Hex(),
		BlockNum:  vLog.BlockNumber,
		Data:      vLog.Data,
		EventName: "Unknown",
		Timestamp: time.Now(),
	}
}

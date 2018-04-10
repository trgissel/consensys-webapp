package controllers

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// TransactionDetails bits of info
type TransactionDetails struct {
	GasUsed int `json:"gasUsed"`
}

// GetTransactionDetails get transaction details
func GetTransactionDetails(url string, txHashString string) (TransactionDetails, error) {
	var txDetails TransactionDetails
	txHash := common.HexToHash(txHashString)
	client, err := ethclient.Dial(url)
	ctx, cancel := context.WithTimeout(context.Background(), 10000*time.Millisecond)
	defer cancel() // releases resources if slowOperation completes before timeout elapses
	tx, _, err := client.TransactionByHash(ctx, txHash)
	if err != nil {
		log.Fatalf("Failed to find transaction by hash: %v", err)
		return txDetails, errors.New("404")
	}
	txDetails.GasUsed = int(tx.Gas())
	return txDetails, nil
}

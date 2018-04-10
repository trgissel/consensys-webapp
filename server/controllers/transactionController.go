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
	d := time.Now().Add(1000 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()
	tx, _, err := client.TransactionByHash(ctx, txHash)
	if err != nil {
		log.Fatalf("Failed to find transaction by hash: %v", err)
		return txDetails, errors.New("404")
	}
	txDetails.GasUsed = int(tx.Gas())
	defer cancel()
	return txDetails, nil
}

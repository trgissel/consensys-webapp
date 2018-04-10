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

type key int

// GetTransactionDetails get transaction details
func GetTransactionDetails(url string, txHashString string) (TransactionDetails, error) {
	var txDetails TransactionDetails
	txHash := common.HexToHash(txHashString)
	client, err := ethclient.Dial(url)
	d := time.Now().Add(1000 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()
	const requestIDKey key = 0
	ctx2 := context.WithValue(ctx, requestIDKey, txHash)
	tx, _, err := client.TransactionByHash(ctx2, txHash)
	if err != nil {
		cancel()
		log.Fatalf("Failed to find transaction by hash: %v", err)
		return txDetails, errors.New("404")
	}
	txDetails.GasUsed = int(tx.Gas())
	defer cancel()
	return txDetails, nil
}

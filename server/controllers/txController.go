package controllers

import (
	"context"
	"log"

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
	var ctx context.Context
	tx, _, err := client.TransactionByHash(ctx, txHash)
	if err != nil {
		log.Fatalf("Failed to find transaction by hash: %v", err)
		return txDetails, err
	}
	txDetails.GasUsed = int(tx.Gas())
	return txDetails, nil
}

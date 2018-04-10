package controllers

import (
	"context"
	"errors"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// TransactionDetails bits of info
type TransactionDetails struct {
	GasUsed int `json:"gasUsed"`
}

type key int

// GetTransactionDetails get transaction details
func GetTransactionDetails(ctx context.Context, url string, txHashString string) (TransactionDetails, error) {
	var txDetails TransactionDetails
	txHash := common.HexToHash(txHashString)
	client, err := ethclient.Dial(url)
	var txHashes [4]string
	ctx2 := context.WithValue(ctx, txHashes, txHash)
	tx, _, err := client.TransactionByHash(ctx2, txHash)
	if err != nil {
		log.Fatalf("Failed to find transaction by hash: %v", err)
		return txDetails, errors.New("404")
	}
	txDetails.GasUsed = int(tx.Gas())
	return txDetails, nil
}

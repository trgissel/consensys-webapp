package controllers

import (
	"context"
	"errors"
	"log"

	"github.com/ethereum/go-ethereum/common"
)

// TransactionDetails bits of info
type TransactionDetails struct {
	GasUsed int `json:"gasUsed"`
}

type key int

// GetTransactionDetails get transaction details
func GetTransactionDetails(url string, keystorePath string, passphrase string, txHashString string) (TransactionDetails, error) {
	var txDetails TransactionDetails
	conn, _, err := EthereumClientConnect(url, keystorePath, passphrase)
	txHash := common.HexToHash(txHashString)
	tx, _, err := conn.TransactionByHash(context.TODO(), txHash)
	if err != nil {
		log.Printf("Failed to find transaction by hash: %v\n", err)
		return txDetails, errors.New("404")
	}
	txDetails.GasUsed = int(tx.Gas())
	return txDetails, nil
}

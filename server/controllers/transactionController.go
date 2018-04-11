package controllers

import (
	"context"
	"encoding/hex"
	"errors"
	"log"

	"github.com/ethereum/go-ethereum/common"
)

// TransactionDetails bits of info
type TransactionDetails struct {
	Recipient string `json:"recipient"`
	Price     string `json:"price"`
	GasUsed   int    `json:"gasUsed"`
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
	price := *(tx.GasPrice())
	txDetails.Price = price.String()
	txDetails.GasUsed = int(tx.Gas())
	addressP := tx.To()
	if addressP != nil {
		txDetails.Recipient = hex.EncodeToString((*addressP)[:])
	}
	return txDetails, nil
}

package controllers

import (
	"encoding/hex"
	"fmt"
	"log"
	"os"

	"../contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

// EthereumClientConnect url should be in the form of "http://172.13.0.3:8545"
func EthereumClientConnect(url string, keystorePath string, passphrase string) (*ethclient.Client, *bind.TransactOpts, error) {
	file, err := os.Open(keystorePath)
	if err != nil {
		fmt.Println("Failed to open keystore file: %v", err)
		return nil, nil, err
	}
	// Create an IPC based RPC connection to a remote node and an authorized transactor
	fmt.Println("attempting to connect to ", url)
	conn, err1 := ethclient.Dial(url)
	if err1 != nil {
		fmt.Println("Failed to connect to the Ethereum client: %v", err1)
		return nil, nil, err1
	}
	fmt.Println("Connected!")
	auth, err2 := bind.NewTransactor(file, passphrase)
	if err2 != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err2)
		return nil, nil, err2
	}
	return conn, auth, nil
}

// CreateCarContract return tx and address
func CreateCarContract(url string, keystorePath string, passphrase string) (string, string, error) {
	conn, auth, err := EthereumClientConnect(url, keystorePath, passphrase)

	var defaultAddress string
	var defaultHash string

	if err != nil {
		log.Fatalf("unable to connect to ethereum network: %v", err)
		return defaultAddress, defaultHash, err
	}
	// Deploy a new awesome contract for the binding demo
	s := "BMW"
	var manufacture [32]byte
	copy(manufacture[:], s)
	address, tx, token, err := contracts.DeployCar(auth, conn, manufacture)
	if err != nil {
		log.Fatalf("Failed to deploy new token contract: %v", err)
		return defaultAddress, defaultHash, err
	}
	fmt.Printf("Contract pending deploy: 0x%x\n", address)
	fmt.Printf("Transaction waiting to be mined: 0x%x\n\n", tx.Hash())
	if token == nil {
		log.Fatalf("Unable to retrieve token")
	}
	var addressHex = "0x" + hex.EncodeToString(address[:])
	return addressHex, tx.Hash().String(), nil
}

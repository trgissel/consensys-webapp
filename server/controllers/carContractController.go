package controllers

import (
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

// EthereumClientConnect url should be in the form of "http://172.13.0.3:8545"
func EthereumClientConnect(url string, keystorePath string, passcode string) {
	file, err := os.Open(keystorePath)
	// Create an IPC based RPC connection to a remote node and an authorized transactor
	fmt.Println("attempting to connect to ", url)
	_, err1 := ethclient.Dial(url)
	if err1 != nil {
		fmt.Println("Failed to connect to the Ethereum client: %v", err1)
	} else {
		fmt.Println("Connected!")
	}
	_, err2 := bind.NewTransactor(file, passcode)
	if err2 != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}
}

package controllers

import (
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
)

// EthereumClientConnect url should be in the form of "http://172.13.0.3:8545"
func EthereumClientConnect(url string) {
	// Create an IPC based RPC connection to a remote node and an authorized transactor
	fmt.Println("attempting to connect to ", url)
	_, err := ethclient.Dial(url)
	if err != nil {
		fmt.Println("Failed to connect to the Ethereum client: %v", err)
	} else {
		fmt.Println("Connected!")
	}
}

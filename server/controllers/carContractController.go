package controllers

import (
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
)

func EthereumClientConnect() {
	// Create an IPC based RPC connection to a remote node and an authorized transactor
	//TODO: pass in via config
	_, err := ethclient.Dial("http://172.13.0.3:8545")
	if err != nil {
		fmt.Println("Failed to connect to the Ethereum client: %v", err)
	}
}

package controllers

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

const key = `paste the contents of your *testnet* key json here`

func Connect() {
	// Create an IPC based RPC connection to a remote node and an authorized transactor
	_, err := ethclient.Dial("/home/karalabe/.ethereum/testnet/geth.ipc")
	if err != nil {
		log.Fatal("Failed to connect to the Ethereum client: %v", err)
	}
}

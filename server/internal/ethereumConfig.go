package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// EthereumConfig is external
type EthereumConfig struct {
	URL     string `json:"nodeURL"`
	Account string `json:"account"`
}

// LoadEthereumConfig return EthereumConfig
func LoadEthereumConfig(configPath string) EthereumConfig {
	raw, err := ioutil.ReadFile(configPath)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var c EthereumConfig
	json.Unmarshal(raw, &c)
	return c
}

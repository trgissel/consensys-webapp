package models

// DeployContractTransaction represents a tran
type DeployContractTransaction struct {
	Address       [20]byte `json:"address"`
	TransactionID [32]byte `json:"tx"`
}

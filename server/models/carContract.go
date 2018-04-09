package models

// CarContractTransaction represents a tran
type CarContractTransaction struct {
	Address       [20]byte `json:"address"`
	TransactionID [32]byte `json:"tx"`
}

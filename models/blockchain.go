package models


// Struct Block Transaction Count
type BlockTransactionCount struct {
	PreviousHash string `json:"previous_hash"`
	DataHash string `json:"data_hash"`
	NumberOfTransactions int `json:"numberOfTransactions"`
	BlockNumber int `json:"blockNumber"`

}


// Struct Block
type Block struct {
	PreviousBlockHash string `json:"previousBlockHash"`
	CurrentBlockHash string `json:"currentBlockHash"`
	BlockNumber int `json:"blockNumber"`
}
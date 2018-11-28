package models



// Struct Invoke Response
type InvokeResponse struct {
	TxId string `json:"txId"`
	Payload interface{} `json:"payload"`
}

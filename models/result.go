package models


// Struct Result for Pending Transaction
type Result struct {
	RequestType string `json:"requestType"`
	Payload map[string]interface{}  `json:"payload"`
}

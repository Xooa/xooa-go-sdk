package models

import "github.com/graarh/golang-socketio"

//Event message Struct
type Message struct {
	TxId string `json:"txId"`
	Status string `json:"status"`
	BlockNum string `json:"blockNum"`
	EventName string `json:"eventName"`
	Payload string `json:"payload"`
}


type CallbackFn func (Message)

type Authenticate struct {
	Token string `json:"token"`
}


type EventClient struct {
	Client *gosocketio.Client
	EventRegex string
}


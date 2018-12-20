/**
 * Go SDK for Xooa
 *
 * Copyright 2018 Xooa
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at:
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License is distributed
 * on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License
 * for the specific language governing permissions and limitations under the License.
 *
 * Author: Rahul Kamboj
 */

package models

import "time"


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

/*
{
    "txid": "7edd234937bcef5563851ba00f0e89e26c0c22dfbfe8fb73e4dd697b5ce1c5dc",
    "createdt": "2018-12-20T12:28:38.804Z",
    "smartcontract": "rahulxooaals1asbos",
    "creator_msp_id": "XooaMSP",
    "endorser_msp_id": [
        "XooaMSP"
    ],
    "type": "ENDORSER_TRANSACTION",
    "read_set": [
        {
            "chaincode": "lscc",
            "set": [
                {
                    "key": "rahulxooaals1asbos",
                    "version": {
                        "block_num": "37",
                        "tx_num": "0"
                    }
                }
            ]
        },
        {
            "chaincode": "rahulxooaals1asbos",
            "set": []
        }
    ],
    "write_set": [
        {
            "chaincode": "lscc",
            "set": []
        },
        {
            "chaincode": "rahulxooaals1asbos",
            "set": [
                {
                    "key": "abc",
                    "is_delete": false,
                    "value": "sdfsdf"
                }
            ]
        }
    ]
}*/


type Transaction struct {
	TxId string `json:"txid"`
	CreatedAt time.Time `json:"createdt"`
	SmartContract string `json:"smartcontract"`
	CreatorMspId string `json:"creator_msp_id"`
	EndorserMspId []string `json:"endorser_msp_id"`
	Type string `json:"type"`
	ReadSet []SetCol `json:"read_set"`
	WriteSet []SetCol `json:"write_set"`
}

type SetCol struct {
	ChainCode string `json:"chaincode"`
	Set []Set `json:"set"`
}

type Set struct {
	Key string `json:"key"`
	IsDelete bool `json:"is_delete"`
	Value string `json:"value"`

}
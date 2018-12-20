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

package test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
	"context"
	"time"
	"github.com/Xooa/xooa-go-sdk/lib"
	"github.com/Xooa/xooa-go-sdk"
)



type Payload struct {
	Args []string `json:"args"`
}

func Before() string {
	cfg, err := xooa.NewConfiguration()
	if err != nil {
		fmt.Println(err);
	}
	x := xooa_client.NewXooaClient(cfg)
	invokeResponse , _,_ := x.Invoke(context.TODO(),"set",map[string]interface{}{"timeout" : "10000"},Payload{Args:[]string{"arg1","arg2"}})
	return  invokeResponse.TxId;
}


func TestInvoke(t *testing.T){
	Init()
	assert := assert.New(t)

	cfg, err := xooa.NewConfiguration()
	if err != nil {
		fmt.Println(err);
		t.Fail()
	}
	x := xooa_client.NewXooaClient(cfg)
	invokeResponse , pendingTransaction,errI := x.Invoke(context.TODO(),"set",map[string]interface{}{"timeout" : "10000"},Payload{Args:[]string{"arg1","arg2"}})
	if errI != nil {
		fmt.Println(errI)
		t.Fail()
	}
	assert.Equal(invokeResponse.Payload,"arg2","It must be same as input Paylod")
	assert.NotEqual(invokeResponse.TxId,"","Must be a non empty string ")
	assert.Equal(pendingTransaction.ResultId,"","Must be an empty string ")
	assert.Equal(pendingTransaction.ResultURL,"","Must be an empty string ")
	PassFailPrint(*t)
}

func TestInvokeAsync(t *testing.T){
	Init()
	assert := assert.New(t)

	cfg, err := xooa.NewConfiguration()
	if err != nil {
		fmt.Println(err);
		t.Fail()
	}
	x := xooa_client.NewXooaClient(cfg)
	invokeResponse , pendingTransaction,errI := x.InvokeAsync(context.TODO(),"set",map[string]interface{}{},Payload{Args:[]string{"arg1","arg2"}})
	if errI != nil {
		fmt.Println(errI)
		t.Fail()
	}
	assert.NotEqual(invokeResponse.Payload,"arg2","Must be Empty in Async")
	assert.Equal(invokeResponse.TxId,"","Must be an empty string")
	assert.NotEqual(pendingTransaction.ResultId,"","Must be an non empty String")
	assert.NotEqual(pendingTransaction.ResultURL,"","Must be an non empty String")
	if pendingTransaction.ResultId != "" {
		time.Sleep(4 * time.Second)
		invokeR, error := x.GetResultForInvoke(context.TODO(), pendingTransaction.ResultId)
		if error != nil {
			t.Fail()
		}
		assert.Equal(invokeR.Payload,"arg2","It must be same as input Paylod")
		assert.NotEqual(invokeR.TxId,"","Must be a non empty string ")

	}
	PassFailPrint(*t)
}

func TestGetTransactionByTransactionId(t *testing.T) {
	Init()
	txId := Before()
	assert := assert.New(t)

	cfg, err := xooa.NewConfiguration()
	if err != nil {
		fmt.Println(err);
		t.Fail()
	}
	x := xooa_client.NewXooaClient(cfg)

	transactionData, pendingRes, errB := x.GetTransactionByTransactionId(context.TODO(), txId, map[string]interface{}{})
	assert.NotEqual(transactionData.TxId, "", "Tx Id Cannot be Empty")
	assert.NotEqual(len(transactionData.EndorserMspId), 0, " Endorser MspId Cannot be Zero")
	assert.NotEqual(transactionData.SmartContract, "", " SmartContract must be a Stirng")
	assert.Equal(pendingRes.ResultId, "", "Response Cannot be pending")
	assert.Equal(pendingRes.ResultURL, "", "Response Cannot be pending")
	if errB != nil {
		fmt.Println(err);
		t.Fail()
	}
	PassFailPrint(*t)
}

func TestGetTransactionByTransactionIdAsync(t *testing.T) {
	Init()
	assert := assert.New(t)
	txId := Before()
	cfg, err := xooa.NewConfiguration()
	if err != nil {
		fmt.Println(err);
		t.Fail()
	}
	x := xooa_client.NewXooaClient(cfg)

	transactionData, pendingRes, errorI := x.GetTransactionByTransactionIdAsync(context.TODO(), txId, map[string]interface{}{})
	assert.Equal(transactionData.TxId, "", "Must Empty")
	assert.Equal(transactionData.CreatorMspId, "", "Must Empty")
	assert.Equal(len(transactionData.EndorserMspId),0 , " Must Zero")
	assert.NotEqual(pendingRes.ResultId, "", "Result Id Must be a String ")
	assert.NotEqual(pendingRes.ResultURL, "", "Result URL Must be a String ")
	if errorI != nil {
		t.Fail()
	}
	if pendingRes.ResultId != "" {
		time.Sleep(4 * time.Second)
		transactionData, errorI := x.GetResultForTransactionByTransactionId(context.TODO(), pendingRes.ResultId)
		if errorI != nil {
			t.Fail()
		}
		assert.NotEqual(transactionData.TxId, "", "Tx Id Cannot be Empty")
		assert.NotEqual(transactionData.SmartContract, "", " SmartContract must be a Stirng")

	}
	PassFailPrint(*t)
}




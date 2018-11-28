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
	"go-client/lib"
	"fmt"
	"context"
	"time"
	"go-client"
)



func TestQuery(t *testing.T) {
	Init()
	assert := assert.New(t)

	cfg, err := xooa.NewConfiguration()
	if err != nil {
		fmt.Println(err);
		t.Fail()
	}
	x := xooa_client.NewXooaClient(cfg)
	invokeResponse, pendingTransaction, errI := x.Query(context.TODO(), "get", map[string]interface{}{"args": "arg1"})
	if errI != nil {
		fmt.Println(errI)
		t.Fail()
	}
	assert.Equal(invokeResponse.Payload, "arg2", "It must be same as input Paylod")
	assert.Equal(pendingTransaction.ResultId, "", "Must be an empty string ")
	assert.Equal(pendingTransaction.ResultURL, "", "Must be an empty string ")
	PassFailPrint(*t)
}

func TestQueryAsync(t *testing.T) {
	Init()
	assert := assert.New(t)

	cfg, err := xooa.NewConfiguration()
	if err != nil {
		fmt.Println(err);
		t.Fail()
	}
	x := xooa_client.NewXooaClient(cfg)
	queryResponse, pendingTransaction, errI := x.QueryAsync(context.TODO(), "get", map[string]interface{}{"args":"arg1"})
	if errI != nil {
		fmt.Println(errI)
		t.Fail()
	}
	assert.NotEqual(queryResponse.Payload, "arg2", "Must be Empty in Async")
	assert.NotEqual(pendingTransaction.ResultId, "", "Must be an non empty String")
	assert.NotEqual(pendingTransaction.ResultURL, "", "Must be an non empty String")
	if pendingTransaction.ResultId != "" {
		time.Sleep(4 * time.Second)
		queryR, error := x.GetResultForQuery(context.TODO(), pendingTransaction.ResultId)
		if error != nil {
			t.Fail()
		}
		assert.Equal(queryR.Payload, "arg2", "It must be same as input Paylod")
	}
	PassFailPrint(*t)
}




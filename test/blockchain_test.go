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
	"github.com/Xooa/xooa-go-sdk/lib"
	"fmt"
	"context"
	"testing"
	"github.com/stretchr/testify/assert"
	"time"
	"github.com/fatih/color"
	"runtime"
	"github.com/Xooa/xooa-go-sdk"
)
// MyCaller returns the caller of the function that called it :)
func MyCaller() string {

	// we get the callers as uintptrs - but we just need 1
	fpcs := make([]uintptr, 1)

	// skip 3 levels to get to the caller of whoever called Caller()
	n := runtime.Callers(3, fpcs)
	if n == 0 {
		return "n/a" // proper error her would be better
	}

	// get the info of the actual function that's in the pointer
	fun := runtime.FuncForPC(fpcs[0]-1)
	if fun == nil {
		return "n/a"
	}

	// return its name
	return fun.Name()
}

func PassFailPrint(t testing.T){

	if t.Failed() != true {
		color.Green(MyCaller() + " ✓ Pass")
	} else {
		color.Red(MyCaller() + " X Fail")
	}
}

func Init() {
	url := "https://api.xooa.com/api/v1"
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBcGlLZXkiOiJZQTYxV1RTLUZOVjRETU4tSFAzUlBXUy00RlJKODRQIiwiQXBpU2VjcmV0IjoiWElkVTJSd3U4YVVaOHJuIiwiUGFzc3BocmFzZSI6IjQyMzBiYzczMmYwMTY1OGZjNTQ0Njk0N2I2N2Q1NjY2IiwiaWF0IjoxNTQ1MzA3MTAxfQ.1tssnKC550rali-YtxTxAcXlA1krzI4jmVogk2jdWMM"
	logLevel := "info"
	xooa_client.SetUrl(url)
	xooa_client.SetToken(token)
	xooa_client.SetLogLevel(logLevel)

}

func TestGetBlockByNumber(t *testing.T) {
	Init()
	assert := assert.New(t)

	cfg, err := xooa.NewConfiguration()
	if err != nil {
		fmt.Println(err);
		t.Fail()
	}
	x := xooa_client.NewXooaClient(cfg)

	blockData, pendingRes, errB := x.GetBlockByNumber(context.TODO(), "1", map[string]interface{}{})
	assert.NotEqual(blockData.BlockNumber, 0, "Block Number Cannot be 0")
	assert.NotEqual(blockData.DataHash, "", "Data Hash Cannot be Empty")
	assert.NotEqual(blockData.PreviousHash, "", " Previous Hash Cannot be Empty")
	assert.NotEqual(blockData.NumberOfTransactions, 0, " NumberOfTransactions must be a positive Number")
	assert.Equal(pendingRes.ResultId, "", "Response Cannot be pending")
	assert.Equal(pendingRes.ResultURL, "", "Response Cannot be pending")
	if errB != nil {
		fmt.Println(err);
		t.Fail()
	}
	PassFailPrint(*t)
}

func TestGetBlockByNumberAsync(t *testing.T) {
	Init()
	assert := assert.New(t)

	cfg, err := xooa.NewConfiguration()
	if err != nil {
		fmt.Println(err);
		t.Fail()
	}
	x := xooa_client.NewXooaClient(cfg)

	blockData, pendingRes, error := x.GetBlockByNumberAsync(context.TODO(), "1", map[string]interface{}{})
	assert.Equal(blockData.BlockNumber, 0, "Must Empty")
	assert.Equal(blockData.DataHash, "", "Must Empty")
	assert.Equal(blockData.PreviousHash, "", " Must Empty")
	assert.NotEqual(pendingRes.ResultId, "", "Result Id Must be a String ")
	assert.NotEqual(pendingRes.ResultURL, "", "Result URL Must be a String ")
	if error != nil {
		t.Fail()
	}
	if pendingRes.ResultId != "" {
		time.Sleep(4 * time.Second)
		block, error := x.GetResultForBlockByNumber(context.TODO(), pendingRes.ResultId)
		if error != nil {
			t.Fail()
		}
		assert.NotEqual(block.BlockNumber, 0, "Block Number Cannot be 0")
		assert.NotEqual(block.DataHash, "", "Data Hash Cannot be Empty")
		assert.NotEqual(block.PreviousHash, "", " Previous Hash Cannot be Empty")
		assert.NotEqual(block.NumberOfTransactions, 0, " NumberOfTransactions must be a positive Number")

	}
	PassFailPrint(*t)
}

func TestGetCurrentBlock(t *testing.T) {
	assert := assert.New(t)

	cfg, err := xooa.NewConfiguration()
	if err != nil {
		fmt.Println(err);
		t.Fail()
	}
	x := xooa_client.NewXooaClient(cfg)

	blockData, pendingRes, error := x.GetCurrentBlock(context.TODO(), map[string]interface{}{})
	assert.NotEqual(blockData.CurrentBlockHash, "", "Current Block Hash Cannot be Empty")
	assert.NotEqual(blockData.PreviousBlockHash, "", "Previous Block Hash Cannot be Empty")
	assert.NotEqual(blockData.BlockNumber, 0, " Block Number Cannot be 0")
	assert.Equal(pendingRes.ResultId, "", "Response Cannot be pending")
	assert.Equal(pendingRes.ResultURL, "", "Response Cannot be pending")
	if error != nil {
		t.Fail()
	}
	PassFailPrint(*t)

}

func TestGetCurrentBlockAsync(t *testing.T) {
	assert := assert.New(t)

	cfg, err := xooa.NewConfiguration()
	if err != nil {
		fmt.Println(err);
		t.Fail()
	}
	x := xooa_client.NewXooaClient(cfg)

	blockData, pendingRes, error := x.GetCurrentBlockAsync(context.TODO(), map[string]interface{}{})
	assert.Equal(blockData.CurrentBlockHash, "", "Current Block Hash Must be Empty")
	assert.Equal(blockData.PreviousBlockHash, "", "Previous Block Hash Must be Empty")
	assert.Equal(blockData.BlockNumber, 0, " Block Number Must be 0")
	assert.NotEqual(pendingRes.ResultId, "", "Result Id Must be a String ")
	assert.NotEqual(pendingRes.ResultURL, "", "Result URL Must be a String ")
	if error != nil {
		t.Fail()
	}
	if pendingRes.ResultId != "" {
		time.Sleep(4 * time.Second)
		block, error := x.GetResultForCurrentBlock(context.TODO(), pendingRes.ResultId)
		if error != nil {
			t.Fail()
		}
		assert.NotEqual(block.BlockNumber, 0, "Block Number Cannot be Zero")
		assert.NotEqual(block.CurrentBlockHash, "", "Block Hash Cannot be Empty")
		assert.NotEqual(block.PreviousBlockHash, "", "Block Hash Cannot be Empty")

	}
	PassFailPrint(*t)

}

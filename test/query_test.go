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




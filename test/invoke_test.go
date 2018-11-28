package test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
	"context"
	"time"
	"go-client/lib"
	"go-client"
)



type Payload struct {
	Args []string `json:"args"`
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



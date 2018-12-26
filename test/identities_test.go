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
	"github.com/Xooa/xooa-go-sdk/lib"
	"fmt"
	"context"
	"github.com/Xooa/xooa-go-sdk/models"
	"time"
	"github.com/Xooa/xooa-go-sdk"
)

func TestAuthenticatedIdentityInformation(t *testing.T){
	assert := assert.New(t)

	cfg, err := xooa.NewConfiguration()
	if err != nil {
		fmt.Println(err);
		t.Fail()
	}
	x := xooa_client.NewXooaClient(cfg)
	identity, errI := x.AuthenticatedIdentityInformation(context.TODO())
	if errI != nil {
		fmt.Println(errI)
		t.Fail()
	}
	assert.NotEqual(identity.Access,"","Must be a non empty String")
	assert.NotEqual(identity.Id,"","Must be a non empty String")
	assert.NotEqual(identity.IdentityName,"","Must be a non empty String")
	assert.NotEqual(identity.CreatedAt,"0001-01-01 00:00:00 +0000 UTC","Must be a non empty time String ")
	assert.NotEqual(identity.Attrs[0].Name,"","Must be a non empty String")
	assert.NotEqual(identity.Attrs[0].Value,"","Must be a non empty String")
	PassFailPrint(*t)
}


func TestIdentitiesGetAllIdentities(t *testing.T){
	assert := assert.New(t)

	cfg, err := xooa.NewConfiguration()
	if err != nil {
		fmt.Println(err);
		t.Fail()
	}
	x := xooa_client.NewXooaClient(cfg)
	identity, errI := x.IdentitiesGetAllIdentities(context.TODO())
	if errI != nil {
		fmt.Println(errI)
		t.Fail()
	}
	assert.NotEqual(identity[0].Id,"","Must be a non empty String")
	assert.NotEqual(identity[0].IdentityName,"","Must be a non empty String")
	assert.NotEqual(identity[0].CreatedAt,"0001-01-01 00:00:00 +0000 UTC","Must be a non empty time String ")
	PassFailPrint(*t)
}

func TestIdentityInformation(t *testing.T) {
	assert := assert.New(t)

	cfg, err := xooa.NewConfiguration()
	if err != nil {
		fmt.Println(err);
		t.Fail()
	}
	x := xooa_client.NewXooaClient(cfg)
	//Finding Id To mock
	identityMe,errI := x.AuthenticatedIdentityInformation(context.TODO())

	identity, errI := x.IdentityInformation(context.TODO(),identityMe.Id)
	if errI != nil {
		fmt.Println(errI)
		t.Fail()
	}
	assert.NotEqual(identity.Access,"","Must be a non empty String")
	assert.NotEqual(identity.Id,"","Must be a non empty String")
	assert.NotEqual(identity.IdentityName,"","Must be a non empty String")
	assert.NotEqual(identity.CreatedAt,"0001-01-01 00:00:00 +0000 UTC","Must be a non empty time String ")
	assert.NotEqual(identity.Attrs[0].Name,"","Must be a non empty String")
	assert.NotEqual(identity.Attrs[0].Value,"","Must be a non empty String")
	PassFailPrint(*t)
}

func TestEnrollIdentity(t *testing.T) {
	assert := assert.New(t)

	cfg, err := xooa.NewConfiguration()
	if err != nil {
		fmt.Println(err);
		t.Fail()
	}
	x := xooa_client.NewXooaClient(cfg)
	var aAttr []models.Attrs
	aAttr = append(aAttr,models.Attrs{Name:"Test Name",Ecert: true,Value: "Test Value"})
	identityN := models.Identity{
		IdentityName: "Test Identity",
		Access: "r",
		Attrs: aAttr,
		CanManageIdentities: true,
	}

	identity,pendingTransaction,errI := x.EnrollIdentity(context.TODO(),map[string]interface{}{},identityN)
	if errI != nil {
		fmt.Println(errI)
		t.Fail()
	}
	assert.NotEqual(identity.Access,"","Must be a non empty String")
	assert.NotEqual(identity.ApiToken,"","Must be a non empty String")
	assert.NotEqual(identity.Id,"","Must be a non empty String")
	assert.NotEqual(identity.IdentityName,"","Must be a non empty String")
	assert.NotEqual(identity.CreatedAt,"0001-01-01 00:00:00 +0000 UTC","Must be a non empty time String ")
	assert.NotEqual(identity.Attrs[0].Name,"","Must be a non empty String")
	assert.NotEqual(identity.Attrs[0].Value,"","Must be a non empty String")
	assert.Equal(pendingTransaction.ResultId,"","Must be an empty string ")
	assert.Equal(pendingTransaction.ResultURL,"","Must be an empty string ")
	PassFailPrint(*t)
}


func TestEnrollIdentityAsync(t *testing.T) {
	assert := assert.New(t)

	cfg, err := xooa.NewConfiguration()
	if err != nil {
		fmt.Println(err);
		t.Fail()
	}
	x := xooa_client.NewXooaClient(cfg)
	var aAttr []models.Attrs
	aAttr = append(aAttr,models.Attrs{Name:"Test Name",Ecert: true,Value: "Test Value"})
	identityN := models.Identity{
		IdentityName: "Test Identity",
		Access: "r",
		Attrs: aAttr,
		CanManageIdentities: true,
	}

	identity,pendingTransaction,errI := x.EnrollIdentityAsync(context.TODO(),map[string]interface{}{},identityN)
	if errI != nil {
		fmt.Println(errI)
		t.Fail()
	}
	assert.Equal(identity.Access,"","Must be an empty String")
	assert.Equal(identity.ApiToken,"","Must be an empty String")
	assert.Equal(identity.Id,"","Must be an empty String")
	assert.Equal(identity.IdentityName,"","Must be an empty String")
	assert.NotEqual(pendingTransaction.ResultId,"","Must be a non empty string ")
	assert.NotEqual(pendingTransaction.ResultURL,"","Must be a non empty string ")
	if pendingTransaction.ResultId != "" {
		time.Sleep(4 * time.Second)
		identityD, error := x.GetResultForEnrollIdentity(context.TODO(), pendingTransaction.ResultId)
		if error != nil {
			t.Fail()
		}
		assert.NotEqual(identityD.Access,"","Must be a non empty String")
		assert.NotEqual(identityD.ApiToken,"","Must be a non empty String")
		assert.NotEqual(identityD.Id,"","Must be a non empty String")
		assert.NotEqual(identityD.IdentityName,"","Must be a non empty String")
		assert.NotEqual(identityD.CreatedAt,"0001-01-01 00:00:00 +0000 UTC","Must be a non empty time String ")
		assert.NotEqual(identityD.Attrs[0].Name,"","Must be a non empty String")
		assert.NotEqual(identityD.Attrs[0].Value,"","Must be a non empty String")
	}
	PassFailPrint(*t)
}


func TestDeleteIdentity(t *testing.T){
	assert := assert.New(t)

	cfg, err := xooa.NewConfiguration()
	if err != nil {
		fmt.Println(err);
		t.Fail()
	}
	x := xooa_client.NewXooaClient(cfg)
	var aAttr []models.Attrs
	aAttr = append(aAttr,models.Attrs{Name:"Test Name",Ecert: true,Value: "Test Value"})
	identityN := models.Identity{
		IdentityName: "Test Identity",
		Access: "r",
		Attrs: aAttr,
		CanManageIdentities: true,
	}

	identity,pendingTransaction,errI := x.EnrollIdentity(context.TODO(),map[string]interface{}{},identityN)
	if errI != nil {
		fmt.Println(errI)
		t.Fail()
	}
	assert.NotEqual(identity.Access,"","Must be a non empty String")
	assert.NotEqual(identity.ApiToken,"","Must be a non empty String")
	assert.NotEqual(identity.Id,"","Must be a non empty String")
	assert.NotEqual(identity.IdentityName,"","Must be a non empty String")
	assert.NotEqual(identity.CreatedAt,"0001-01-01 00:00:00 +0000 UTC","Must be a non empty time String ")
	assert.NotEqual(identity.Attrs[0].Name,"","Must be a non empty String")
	assert.NotEqual(identity.Attrs[0].Value,"","Must be a non empty String")
	assert.Equal(pendingTransaction.ResultId,"","Must be an empty string ")
	assert.Equal(pendingTransaction.ResultURL,"","Must be an empty string ")

	identityD,pendingTransaction ,errD := x.DeleteIdentity(context.TODO(),identity.Id,map[string]interface{}{})
	if errD != nil {
		fmt.Println(errD)
		t.Fail()
	}
	assert.Equal(identityD.Deleted,true,"identity must delete")
	assert.Equal(pendingTransaction.ResultId,"","Must be an empty string ")
	assert.Equal(pendingTransaction.ResultURL,"","Must be an empty string ")


	PassFailPrint(*t)
}


func TestDeleteIdentityAsync(t *testing.T){
	assert := assert.New(t)

	cfg, err := xooa.NewConfiguration()
	if err != nil {
		fmt.Println(err);
		t.Fail()
	}
	x := xooa_client.NewXooaClient(cfg)
	var aAttr []models.Attrs
	aAttr = append(aAttr,models.Attrs{Name:"Test Name",Ecert: true,Value: "Test Value"})
	identityN := models.Identity{
		IdentityName: "Test Identity",
		Access: "r",
		Attrs: aAttr,
		CanManageIdentities: true,
	}

	identity,pendingTransaction,errI := x.EnrollIdentity(context.TODO(),map[string]interface{}{},identityN)
	if errI != nil {
		fmt.Println(errI)
		t.Fail()
	}
	assert.NotEqual(identity.Access,"","Must be a non empty String")
	assert.NotEqual(identity.ApiToken,"","Must be a non empty String")
	assert.NotEqual(identity.Id,"","Must be a non empty String")
	assert.NotEqual(identity.IdentityName,"","Must be a non empty String")
	assert.NotEqual(identity.CreatedAt,"0001-01-01 00:00:00 +0000 UTC","Must be a non empty time String ")
	assert.NotEqual(identity.Attrs[0].Name,"","Must be a non empty String")
	assert.NotEqual(identity.Attrs[0].Value,"","Must be a non empty String")
	assert.Equal(pendingTransaction.ResultId,"","Must be an empty string ")
	assert.Equal(pendingTransaction.ResultURL,"","Must be an empty string ")

	identityD,pendingTransaction ,errD := x.DeleteIdentityAsync(context.TODO(),identity.Id,map[string]interface{}{})
	if errD != nil {
		fmt.Println(errD)
		t.Fail()
	}
	assert.Equal(identityD.Deleted,false,"Must Wait as Async")
	assert.NotEqual(pendingTransaction.ResultId,"","Must be a non empty string ")
	assert.NotEqual(pendingTransaction.ResultURL,"","Must be a non empty string ")

	if pendingTransaction.ResultId != "" {
		time.Sleep(4 * time.Second)
		identityD, errA := x.GetResultForDeleteIdentity(context.TODO(), pendingTransaction.ResultId)
		if errA != nil {
			fmt.Println(errA)
			t.Fail()
		}
		assert.Equal(identityD.Deleted,true,"identity must delete")
	}

	PassFailPrint(*t)
}


func TestRegenerateToken(t *testing.T){
	assert := assert.New(t)

	cfg, err := xooa.NewConfiguration()
	if err != nil {
		fmt.Println(err);
		t.Fail()
	}
	x := xooa_client.NewXooaClient(cfg)
	var aAttr []models.Attrs
	aAttr = append(aAttr,models.Attrs{Name:"Test Name",Ecert: true,Value: "Test Value"})
	identityN := models.Identity{
		IdentityName: "Test Identity",
		Access: "r",
		Attrs: aAttr,
		CanManageIdentities: true,
	}

	identity,pendingTransaction,errI := x.EnrollIdentity(context.TODO(),map[string]interface{}{},identityN)
	if errI != nil {
		fmt.Println(errI)
		t.Fail()
	}
	assert.NotEqual(identity.Access,"","Must be a non empty String")
	assert.NotEqual(identity.ApiToken,"","Must be a non empty String")
	assert.NotEqual(identity.Id,"","Must be a non empty String")
	assert.NotEqual(identity.IdentityName,"","Must be a non empty String")
	assert.NotEqual(identity.CreatedAt,"0001-01-01 00:00:00 +0000 UTC","Must be a non empty time String ")
	assert.NotEqual(identity.Attrs[0].Name,"","Must be a non empty String")
	assert.NotEqual(identity.Attrs[0].Value,"","Must be a non empty String")
	assert.Equal(pendingTransaction.ResultId,"","Must be an empty string ")
	assert.Equal(pendingTransaction.ResultURL,"","Must be an empty string ")


	identityR,pendingTransaction,errI := x.RegenerateToken(context.TODO(),identity.Id,map[string]interface{}{})
	if errI != nil {
		fmt.Println(errI)
		t.Fail()
	}
	assert.NotEqual(identityR.Access,"","Must be a non empty String")
	assert.NotEqual(identityR.ApiToken,identity.ApiToken,"Api Token must change after Regenerating")
	assert.NotEqual(identityR.Id,"","Must be a non empty String")
	assert.NotEqual(identityR.IdentityName,"","Must be a non empty String")
	assert.NotEqual(identityR.CreatedAt,"0001-01-01 00:00:00 +0000 UTC","Must be a non empty time String ")
	assert.NotEqual(identityR.Attrs[0].Name,"","Must be a non empty String")
	assert.NotEqual(identityR.Attrs[0].Value,"","Must be a non empty String")
	assert.Equal(pendingTransaction.ResultId,"","Must be an empty string ")
	assert.Equal(pendingTransaction.ResultURL,"","Must be an empty string ")
	PassFailPrint(*t)
}


func TestRegenerateTokenAsync(t *testing.T){
	assert := assert.New(t)

	cfg, err := xooa.NewConfiguration()
	if err != nil {
		fmt.Println(err);
		t.Fail()
	}
	x := xooa_client.NewXooaClient(cfg)
	var aAttr []models.Attrs
	aAttr = append(aAttr,models.Attrs{Name:"Test Name",Ecert: true,Value: "Test Value"})
	identityN := models.Identity{
		IdentityName: "Test Identity",
		Access: "r",
		Attrs: aAttr,
		CanManageIdentities: true,
	}

	identity,pendingTransaction,errI := x.EnrollIdentity(context.TODO(),map[string]interface{}{},identityN)
	if errI != nil {
		fmt.Println(errI)
		t.Fail()
	}
	assert.NotEqual(identity.Access,"","Must be a non empty String")
	assert.NotEqual(identity.ApiToken,"","Must be a non empty String")
	assert.NotEqual(identity.Id,"","Must be a non empty String")
	assert.NotEqual(identity.IdentityName,"","Must be a non empty String")
	assert.NotEqual(identity.CreatedAt,"0001-01-01 00:00:00 +0000 UTC","Must be a non empty time String ")
	assert.NotEqual(identity.Attrs[0].Name,"","Must be a non empty String")
	assert.NotEqual(identity.Attrs[0].Value,"","Must be a non empty String")
	assert.Equal(pendingTransaction.ResultId,"","Must be an empty string ")
	assert.Equal(pendingTransaction.ResultURL,"","Must be an empty string ")

	identityR,pendingTransaction,errI := x.RegenerateTokenAsync(context.TODO(),identity.Id,map[string]interface{}{})
	if errI != nil {
		fmt.Println(errI)
		t.Fail()
	}
	assert.Equal(identityR.Access,"","Must be a empty String")
	assert.NotEqual(pendingTransaction.ResultId,"","Must be a non empty String")
	assert.NotEqual(pendingTransaction.ResultURL,"","Must be a non empty String ")
	if pendingTransaction.ResultId != "" {
		time.Sleep(4 * time.Second)
		regenerateResponse, errorI := x.GetResultForRegenerateTokenAsync(context.TODO(), pendingTransaction.ResultId)
		if errorI != nil {
			t.Fail()
		}
		assert.NotEqual(regenerateResponse.Access,"","Must be a non empty String")
		assert.NotEqual(regenerateResponse.ApiToken,identity.ApiToken,"Api Token must change after Regenerating")
		assert.NotEqual(regenerateResponse.Id,"","Must be a non empty String")
		assert.NotEqual(regenerateResponse.IdentityName,"","Must be a non empty String")
		assert.NotEqual(regenerateResponse.CreatedAt,"0001-01-01 00:00:00 +0000 UTC","Must be a non empty time String ")
		assert.NotEqual(regenerateResponse.Attrs[0].Name,"","Must be a non empty String")
		assert.NotEqual(regenerateResponse.Attrs[0].Value,"","Must be a non empty String")
	}
	PassFailPrint(*t)
}

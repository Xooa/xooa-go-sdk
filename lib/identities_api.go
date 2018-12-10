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

package xooa

import (
	"io/ioutil"
	"net/url"
	"strings"
	"context"
	"fmt"
	"github.com/Xooa/xooa-go-sdk/models"
	"encoding/json"
	log "github.com/sirupsen/logrus"
)

// Linger please
var (
	_ context.Context
)

type IdentitiesApiService service


/* AuthenticatedIdentityInformation Authenticated Identity Information
 This End Point Returns Information about the Authenticated Identity
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @return */
func (a *IdentitiesApiService) AuthenticatedIdentityInformation(ctx context.Context) (models.Identity, *models.XooaApiException) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := BasePath + "/identities/me/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	identity := models.Identity{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	if err != nil {
		log.Error(err)
		return identity, &models.XooaApiException{
			ErrorCode:0,
			ErrorMessage: err.Error(),
		}
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		log.Error(err)
		return identity, &models.XooaApiException{
			ErrorCode:0,
			ErrorMessage: err.Error(),
		}
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		log.Error(localVarHttpResponse.StatusCode,localVarHttpResponse.Body)
		return identity, models.NewXooaApiException(localVarHttpResponse);
	}
	content, _ := ioutil.ReadAll(localVarHttpResponse.Body)

	json.Unmarshal(content, &identity)
	log.Info(identity)
	return identity, nil
}

/* IdentitiesGetAllIdentities Get All Identities
 Get all identities from the identity registry
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @return */
func (a *IdentitiesApiService) IdentitiesGetAllIdentities(ctx context.Context) ([]models.Identity, *models.XooaApiException) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := BasePath + "/identities/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	identity := []models.Identity{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	if err != nil {
		log.Error(err)
		return identity, &models.XooaApiException{
			ErrorCode:0,
			ErrorMessage: err.Error(),
		}
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		log.Error(err)
		return identity, &models.XooaApiException{
			ErrorCode:0,
			ErrorMessage: err.Error(),
		}
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		log.Error(localVarHttpResponse.StatusCode,localVarHttpResponse.Body)
		return identity, models.NewXooaApiException(localVarHttpResponse);
	}
	content, _ := ioutil.ReadAll(localVarHttpResponse.Body)
	json.Unmarshal(content, &identity)
	log.Info(identity)
	return identity, nil
}

/* IdentityInformation Identity Information
 Get the specified identity from the identity registry
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param id Identity Id
 @return */
func (a *IdentitiesApiService) IdentityInformation(ctx context.Context, id string) (models.Identity, *models.XooaApiException) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := BasePath + "/identities/{Id}"
	localVarPath = strings.Replace(localVarPath, "{" + "Id" + "}", fmt.Sprintf("%v", id), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	identity := models.Identity{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		log.Error(err)
		return identity, &models.XooaApiException{
			ErrorCode:0,
			ErrorMessage: err.Error(),
		}
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		log.Error(err)
		return identity, &models.XooaApiException{
			ErrorCode:0,
			ErrorMessage: err.Error(),
		}
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		log.Error(localVarHttpResponse.StatusCode ,localVarHttpResponse)
		return identity, models.NewXooaApiException(localVarHttpResponse);
	}
	content, _ := ioutil.ReadAll(localVarHttpResponse.Body)

	json.Unmarshal(content, &identity)
	log.Info(identity)
	return identity, nil
}

/* EnrollIdentity Enroll Identity
 The Enroll Identity end point is used to  enroll new identities for the Smart Contract app.  A success response includes the API Key generated for the identity. This API Key can be used to call API End points on behalf of the enrolled identity. This End Point provides equivalent functionality to adding new identity manually using Xooa console, and identities added using this end point will appear, and can be managed, using Xooa console under the identities tab of the Smart Contract app Required permission: manage identities (canManageIdentities&#x3D;true)
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @param "async" (string) Call this request asynchronously without waiting for response
 @param "timeout" (string) Request timeout
 @return */
func (a *IdentitiesApiService) EnrollIdentity(ctx context.Context, localVarOptionals map[string]interface{}, localVarPostBody interface{}) (models.Identity, models.PendingTransactionResponse, *models.XooaApiException) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarFileName string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := BasePath + "/identities/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	pendingTransaction := models.PendingTransactionResponse{}
	identity := models.Identity{}
	async := false

	if err := typeCheckParameter(localVarOptionals["async"], "string", "async"); err != nil {
		log.Error(err)
		return identity, pendingTransaction, &models.XooaApiException{
			ErrorCode: 0,
			ErrorMessage: err.Error(),
		}
	}
	if err := typeCheckParameter(localVarOptionals["timeout"], "string", "timeout"); err != nil {
		log.Error(err)
		return identity, pendingTransaction, &models.XooaApiException{
			ErrorCode: 0,
			ErrorMessage: err.Error(),
		}
	}

	if localVarTempParam, localVarOk := localVarOptionals["async"].(string); localVarOk {
		async = true
		localVarQueryParams.Add("async", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["timeout"].(string); localVarOk {
		localVarQueryParams.Add("timeout", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		log.Error(err)
		return identity, pendingTransaction, &models.XooaApiException{
			ErrorCode: 0,
			ErrorMessage: err.Error(),
		}
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		log.Error(err)
		return identity, pendingTransaction, &models.XooaApiException{
			ErrorCode: 0,
			ErrorMessage: err.Error(),
		}
	}

	content, _ := ioutil.ReadAll(localVarHttpResponse.Body)
	if localVarHttpResponse.StatusCode >= 300 {
		log.Error(localVarHttpResponse.StatusCode,localVarHttpResponse)
		return identity, pendingTransaction, &models.XooaApiException{
			ErrorCode: 0,
			ErrorMessage: string(content),
		}
	}
	if localVarHttpResponse.StatusCode == 202 {
		log.Warn(localVarHttpResponse.StatusCode,localVarHttpResponse)
		json.Unmarshal(content, &pendingTransaction)
		return identity, pendingTransaction, nil
	}
	if async == true {
		json.Unmarshal(content, &pendingTransaction)
		log.Info(pendingTransaction)
	} else {
		json.Unmarshal(content, &identity)
	}
	log.Info(identity)
	return identity, pendingTransaction, nil
}


/* EnrollIdentityAsync Enroll Identity Async
 The Enroll Identity end point is used to  enroll new identities for the Smart Contract app.  A success response includes the API Key generated for the identity. This API Key can be used to call API End points on behalf of the enrolled identity. This End Point provides equivalent functionality to adding new identity manually using Xooa console, and identities added using this end point will appear, and can be managed, using Xooa console under the identities tab of the Smart Contract app Required permission: manage identities (canManageIdentities&#x3D;true)
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @param "async" (string) Call this request asynchronously without waiting for response
 @param "timeout" (string) Request timeout
 @return */
func (a *IdentitiesApiService) EnrollIdentityAsync(ctx context.Context, localVarOptionals map[string]interface{}, localVarPostBody interface{}) (models.Identity, models.PendingTransactionResponse, *models.XooaApiException) {
	localVarOptionals["async"] = "true"
	return a.EnrollIdentity(ctx, localVarOptionals, localVarPostBody)
}


/* DeleteIdentity Delete Identity
 Delete Identity.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param id Identity Id
 @param optional (nil or map[string]interface{}) with one or more of:
 @param "async" (string) Call this request asynchronously without waiting for response
 @param "timeout" (string) Request timeout in millisecond
 @return */
func (a *IdentitiesApiService) DeleteIdentity(ctx context.Context, id string, localVarOptionals map[string]interface{}) (models.DeleteIdentityResponse, models.PendingTransactionResponse, *models.XooaApiException) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := BasePath + "/identities/{Id}"
	localVarPath = strings.Replace(localVarPath, "{" + "Id" + "}", fmt.Sprintf("%v", id), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	deleteIdentity := models.DeleteIdentityResponse{}
	pendingTransaction := models.PendingTransactionResponse{}
	async := false

	if err := typeCheckParameter(localVarOptionals["async"], "string", "async"); err != nil {
		log.Error(err)
		return deleteIdentity, pendingTransaction, &models.XooaApiException{
			ErrorCode: 0,
			ErrorMessage: err.Error(),
		}
	}
	if err := typeCheckParameter(localVarOptionals["timeout"], "string", "timeout"); err != nil {
		log.Error(err)
		return deleteIdentity, pendingTransaction, &models.XooaApiException{
			ErrorCode: 0,
			ErrorMessage: err.Error(),
		}
	}

	if localVarTempParam, localVarOk := localVarOptionals["async"].(string); localVarOk {
		async = true
		localVarQueryParams.Add("async", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["timeout"].(string); localVarOk {
		localVarQueryParams.Add("timeout", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		log.Error(err)
		return deleteIdentity, pendingTransaction, &models.XooaApiException{
			ErrorCode: 0,
			ErrorMessage: err.Error(),
		}
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		log.Error(err)
		return deleteIdentity,pendingTransaction,models.NewXooaApiException(localVarHttpResponse)
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		log.Error(localVarHttpResponse.StatusCode,localVarHttpResponse)
		return deleteIdentity,pendingTransaction,models.NewXooaApiException(localVarHttpResponse)
	}

	content, err := ioutil.ReadAll(localVarHttpResponse.Body)
	if localVarHttpResponse.StatusCode == 202 {
		log.Warn(localVarHttpResponse.StatusCode,localVarHttpResponse)
		json.Unmarshal(content, &pendingTransaction)
		return deleteIdentity,pendingTransaction, nil
	}
	if err != nil {
		log.Error(err)
		return deleteIdentity, pendingTransaction, &models.XooaApiException {
			ErrorCode: 0,
			ErrorMessage: err.Error(),
		}
	}
	if async == true {
		json.Unmarshal(content, &pendingTransaction)
		log.Info(pendingTransaction)
	} else {
		json.Unmarshal(content, &deleteIdentity)
	}
	log.Info(deleteIdentity)
	return deleteIdentity,pendingTransaction, nil
}

/* DeleteIdentityAsync Delete Identity Async
 Delete Identity.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param id Identity Id
 @param optional (nil or map[string]interface{}) with one or more of:
 @param "async" (string) Call this request asynchronously without waiting for response
 @param "timeout" (string) Request timeout in millisecond
 @return */

func (a *IdentitiesApiService) DeleteIdentityAsync(ctx context.Context, id string, localVarOptionals map[string]interface{}) (models.DeleteIdentityResponse, models.PendingTransactionResponse, *models.XooaApiException) {
	localVarOptionals["async"] = "true"
	return a.DeleteIdentity(ctx, id, localVarOptionals)

}



/* RegenerateToken Regenerate Identity API Token
 Regenerate Identity API Token
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param id Identity Id
 @return */
func (a *IdentitiesApiService) RegenerateToken(ctx context.Context, id string, localVarOptionals map[string]interface{}) (models.Identity, models.PendingTransactionResponse,*models.XooaApiException) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := BasePath + "/identities/{Id}/regeneratetoken"
	localVarPath = strings.Replace(localVarPath, "{" + "Id" + "}", fmt.Sprintf("%v", id), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	pendingResponse := models.PendingTransactionResponse{}
	identity := models.Identity{}
	async := false


	if err := typeCheckParameter(localVarOptionals["async"], "string", "async"); err != nil {
		log.Error(err)
		return identity, pendingResponse, &models.XooaApiException{
			ErrorCode: 0,
			ErrorMessage: err.Error(),
		}
	}
	if err := typeCheckParameter(localVarOptionals["timeout"], "string", "timeout"); err != nil {
		log.Error(err)
		return identity, pendingResponse, &models.XooaApiException{
			ErrorCode: 0,
			ErrorMessage: err.Error(),
		}
	}

	if localVarTempParam, localVarOk := localVarOptionals["async"].(string); localVarOk {
		async = true
		localVarQueryParams.Add("async", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["timeout"].(string); localVarOk {
		localVarQueryParams.Add("timeout", parameterToString(localVarTempParam, ""))
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		log.Error(err)
		return identity,pendingResponse,&models.XooaApiException{
			ErrorCode:0,
			ErrorMessage:err.Error(),
		}
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		log.Error(err)
		return identity,pendingResponse,&models.XooaApiException{
			ErrorCode:0,
			ErrorMessage:err.Error(),
		}
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		log.Error(localVarHttpResponse.StatusCode,localVarHttpResponse)
		return identity,pendingResponse,models.NewXooaApiException(localVarHttpResponse)
	}

	content, err := ioutil.ReadAll(localVarHttpResponse.Body)
	if localVarHttpResponse.StatusCode == 202 {
		log.Warn(localVarHttpResponse.StatusCode,localVarHttpResponse)
		json.Unmarshal(content, &pendingResponse)
		return identity,pendingResponse, nil
	}
	if err != nil {
		log.Error(err)
		return identity, pendingResponse, &models.XooaApiException {
			ErrorCode: 0,
			ErrorMessage: err.Error(),
		}
	}
	if async == true {
		json.Unmarshal(content, &pendingResponse)
		log.Info(pendingResponse)

	} else {
		json.Unmarshal(content, &identity)
	}
	log.Info(identity)
	return identity,pendingResponse, nil
}


/* RegenerateTokenAsync Regenerate Identity API Token Async
 Regenerate Identity API Token
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param id Identity Id
 @return */
func (a *IdentitiesApiService) RegenerateTokenAsync(ctx context.Context, id string, localVarOptionals map[string]interface{}) (models.Identity, models.PendingTransactionResponse,*models.XooaApiException) {
	localVarOptionals["async"] = "true"
	return a.RegenerateToken(ctx,id,localVarOptionals)
}

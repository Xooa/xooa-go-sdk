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
	"encoding/json"
	"github.com/Xooa/xooa-go-sdk/models"
	log "github.com/sirupsen/logrus"
)

var (
	_ context.Context
)

type BlockchainApiService service


/* GetBlockByNumber Get block data (block #)
 Get specific block information such as hash, # of transactions
 * @param ctx context.Context for logging, tracing, etc.
 @param blockNumber Block number to fetch data
 @param optional (nil or map[string]interface{}) with one or more of:
 @param "async" (string) Call this request asynchronously without waiting for response
 @param "timeout" (string) Request timeout in millisecond
 @return */

func (a *BlockchainApiService) GetBlockByNumber(ctx context.Context, blockNumber string, localVarOptionals map[string]interface{}) (models.BlockTransactionCount, models.PendingTransactionResponse, *models.XooaApiException) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := BasePath + "/block/{BlockNumber}"
	localVarPath = strings.Replace(localVarPath, "{" + "BlockNumber" + "}", fmt.Sprintf("%v", blockNumber), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	blockTransactionCount := models.BlockTransactionCount{}
	pendingTransactionResponse := models.PendingTransactionResponse{}
	async := false

	if err := typeCheckParameter(localVarOptionals["async"], "string", "async"); err != nil {
		log.Error(err)
		return blockTransactionCount, pendingTransactionResponse, &models.XooaApiException{
			ErrorCode: 0,
			ErrorMessage: err.Error(),
		}
	}
	if err := typeCheckParameter(localVarOptionals["timeout"], "string", "timeout"); err != nil {
		log.Error(err)
		return blockTransactionCount, pendingTransactionResponse, &models.XooaApiException{
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
		return blockTransactionCount, pendingTransactionResponse, &models.XooaApiException{
			ErrorCode: 0,
			ErrorMessage: err.Error(),
		}
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return blockTransactionCount, pendingTransactionResponse,  models.NewXooaApiException(localVarHttpResponse)
	}

	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		log.Error(localVarHttpResponse.StatusCode,localVarHttpResponse)
		return blockTransactionCount, pendingTransactionResponse, models.NewXooaApiException(localVarHttpResponse)
	}
	content, err := ioutil.ReadAll(localVarHttpResponse.Body)
	if localVarHttpResponse.StatusCode == 202 {
		log.Warn(localVarHttpResponse.StatusCode,localVarHttpResponse)
		json.Unmarshal(content, &pendingTransactionResponse)
		return blockTransactionCount, pendingTransactionResponse, nil
	}
	if err != nil {
		log.Error(err)
		return blockTransactionCount, pendingTransactionResponse, &models.XooaApiException{
			ErrorCode: 0,
			ErrorMessage: err.Error(),
		}
	}
	if async == true {
		json.Unmarshal(content, &pendingTransactionResponse)

	} else {
		json.Unmarshal(content, &blockTransactionCount)
	}
	log.Info(blockTransactionCount)
	return blockTransactionCount, pendingTransactionResponse, nil
}

/* GetBlockByNumberAsync Get block data async (block #)
 Get specific block information such as hash, # of transactions
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param blockNumber Block number to fetch data
 @param optional (nil or map[string]interface{}) with one or more of:
 @param "async" (string) Call this request asynchronously without waiting for response
 @param "timeout" (string) Request timeout in millisecond
 @return */

func (a *BlockchainApiService) GetBlockByNumberAsync(ctx context.Context, blockNumber string, localVarOptionals map[string]interface{}) (models.BlockTransactionCount, models.PendingTransactionResponse, *models.XooaApiException) {
	localVarOptionals["async"] = "true"
	return a.GetBlockByNumber(ctx,blockNumber, localVarOptionals)

}



/* GetCurrentBlock Get current blocks
 Get current blocks in the network
 * @param ctx context.Context for logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @param "async" (string) Call this request asynchronously without waiting for response
 @param "timeout" (string) Request timeout in millisecond
 @return */

func (a *BlockchainApiService) GetCurrentBlock(ctx context.Context, localVarOptionals map[string]interface{}) (models.Block,models.PendingTransactionResponse, *models.XooaApiException) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := BasePath + "/block/current"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	block := models.Block{}
	async := false
	pendingTransactionResponse := models.PendingTransactionResponse{}

	if err := typeCheckParameter(localVarOptionals["async"], "string", "async"); err != nil {
		log.Error(err)
		return block, pendingTransactionResponse,  &models.XooaApiException{
			ErrorCode: 0,
			ErrorMessage: err.Error(),
		}
	}
	if err := typeCheckParameter(localVarOptionals["timeout"], "string", "timeout"); err != nil {
		log.Error(err)
		return block, pendingTransactionResponse, &models.XooaApiException{
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
		return block, pendingTransactionResponse, &models.XooaApiException{
			ErrorCode: 0,
			ErrorMessage: err.Error(),
		}
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		log.Error(err)
		return block, pendingTransactionResponse, models.NewXooaApiException(localVarHttpResponse)
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		log.Error(localVarHttpResponse.StatusCode,localVarHttpResponse)
		return block, pendingTransactionResponse, models.NewXooaApiException(localVarHttpResponse)
	}
	content, err := ioutil.ReadAll(localVarHttpResponse.Body)
	if localVarHttpResponse.StatusCode == 202 {
		log.Warn(localVarHttpResponse.StatusCode,localVarHttpResponse)
		json.Unmarshal(content, &pendingTransactionResponse)
		return block, pendingTransactionResponse, nil
	}
	if err != nil {
		log.Error(err)
		return block, pendingTransactionResponse, &models.XooaApiException {
			ErrorCode: 0,
			ErrorMessage: err.Error(),
		}
	}
	if async == true {
		json.Unmarshal(content, &pendingTransactionResponse)
	} else {
		json.Unmarshal(content, &block)
	}
	log.Info(block)
	return block, pendingTransactionResponse, nil
}

/* GetCurrentBlockAsync Get current blocks async
Get current blocks in the network
* @param ctx context.Context for logging, tracing, etc.
@param optional (nil or map[string]interface{}) with one or more of:
@param "async" (string) Call this request asynchronously without waiting for response
@param "timeout" (string) Request timeout in millisecond
@return */


func (a *BlockchainApiService) GetCurrentBlockAsync(ctx context.Context, localVarOptionals map[string]interface{}) (models.Block,models.PendingTransactionResponse, *models.XooaApiException) {
	localVarOptionals["async"] = "true"
	return a.GetCurrentBlock(ctx,localVarOptionals)
}


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
	"go-client/models"
	"encoding/json"
	"github.com/mitchellh/mapstructure"
	log "github.com/sirupsen/logrus"
)

// Linger please
var (
	_ context.Context
)

type ResultApiService service


/* ResultApiService Fetch result of previously submitted transaction
 API Returns result of previously submitted transaction
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param resultId Returned in previous Query/Invoke/Participant Operation
 @return */
func (a *ResultApiService) Result(ctx context.Context, resultId string) ( models.Result, *models.XooaApiException) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := BasePath + "/results/{ResultId}"
	localVarPath = strings.Replace(localVarPath, "{"+"ResultId"+"}", fmt.Sprintf("%v", resultId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	result := models.Result{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

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
		return result, &models.XooaApiException{
			ErrorCode:0,
			ErrorMessage: err.Error(),
		}
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		log.Error(err)
		return result, &models.XooaApiException{
			ErrorCode:0,
			ErrorMessage: err.Error(),
		}
	}
	defer localVarHttpResponse.Body.Close()
	content, _ := ioutil.ReadAll(localVarHttpResponse.Body)
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		log.Error(localVarHttpResponse.StatusCode , string(bodyBytes))
		return result, &models.XooaApiException {
			ErrorCode: 0,
			ErrorMessage: string(bodyBytes),
		}
	}

	json.Unmarshal(content, &result)
	log.Info(result)
	return result, nil
}

/* GetResultForCurrentBlock Fetch result of previously submitted transaction
 API Returns result of previously submitted transaction
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param resultId Returned in previous CurrentBlock Operation
 @return */
func (a *ResultApiService) GetResultForCurrentBlock(ctx context.Context, resultId string) ( models.Block, *models.XooaApiException) {
	res,error := a.Result(ctx,resultId)
	block := models.Block{}
	if error != nil {
		log.Error(error)
		return block,error
	}
	mapstructure.Decode(res.Payload,&block)
	return block,error


}

/* GetResultForBlockByNumber Fetch result of previously submitted transaction
 API Returns result of previously submitted transaction
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param resultId Returned in previous BlockByNumber Operation
 @return */
func (a *ResultApiService) GetResultForBlockByNumber(ctx context.Context, resultId string) ( models.BlockTransactionCount, *models.XooaApiException) {
	res,error := a.Result(ctx,resultId)
	block := models.BlockTransactionCount{}
	if error != nil {
		log.Error(error)
		return block,error
	}
	res.Payload["previousHash"] = res.Payload["previous_hash"]
	res.Payload["dataHash"] = res.Payload["data_hash"]
	mapstructure.Decode(res.Payload,&block,)
	return block,error


}

/* Fetch result of previously submitted transaction
 API Returns result of previously submitted transaction
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param resultId Returned in previous Invoke Operation
 @return */
func (a *ResultApiService) GetResultForInvoke(ctx context.Context, resultId string) ( models.InvokeResponse, *models.XooaApiException) {
	res,error := a.Result(ctx,resultId)
	invoke := models.InvokeResponse{}
	if error != nil {
		log.Error(error)
		return invoke,error
	}

	mapstructure.Decode(res.Payload,&invoke)
	return invoke,error


}

/* Fetch result of previously submitted transaction
 API Returns result of previously submitted transaction
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param resultId Returned in previous Query Operation
 @return */
func (a *ResultApiService) GetResultForQuery(ctx context.Context, resultId string) ( models.Query, *models.XooaApiException) {
	res,error := a.Result(ctx,resultId)
	query := models.Query{}
	if error != nil {
		log.Error(error)
		return query,error
	}

	mapstructure.Decode(res.Payload,&query)
	return query,error
}

/* GetResultForEnrollIdentity Fetch result of previously submitted transaction
 API Returns result of previously submitted transaction
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param resultId Returned in previous Enroll Identity Operation
 @return */
func (a *ResultApiService) GetResultForEnrollIdentity(ctx context.Context, resultId string) ( models.Identity, *models.XooaApiException) {
	res,error := a.Result(ctx,resultId)
	identity := models.Identity{}
	if error != nil {
		log.Error(error)
		return identity,error
	}
	mapstructure.Decode(res.Payload,&identity)
	return identity,error
}

/* GetResultForDeleteIdentity Fetch result of previously submitted transaction
 API Returns result of previously submitted transaction
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param resultId Returned in previous DeleteIdentity Operation
 @return */
func (a *ResultApiService) GetResultForDeleteIdentity(ctx context.Context, resultId string) ( models.DeleteIdentityResponse, *models.XooaApiException) {
	res,error := a.Result(ctx,resultId)
	identity := models.DeleteIdentityResponse{}
	if error != nil {
		log.Error(error)
		return identity,error
	}
	mapstructure.Decode(res.Payload,&identity)
	return identity,error
}

/* GetResultForRegenerateToken Fetch result of previously submitted transaction
 API Returns result of previously submitted transaction
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param resultId Returned in previous RegenerateToken Operation
 @return */
func (a *ResultApiService) GetResultForRegenerateTokenAsync(ctx context.Context, resultId string) ( models.RegenerateIdentityResponse, *models.XooaApiException) {
	res,error := a.Result(ctx,resultId)
	regenerateResponse := models.RegenerateIdentityResponse{}
	if error != nil {
		log.Error(error)
		return regenerateResponse,error
	}
	mapstructure.Decode(res.Payload,&regenerateResponse)
	return regenerateResponse,error
}





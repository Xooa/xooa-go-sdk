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

package xooa_client

import (
	"go-client/lib"
	"go-client/models"
	"context"
)

type XooaClient struct {
	Xooa *xooa.XooaClient

}

// NewXooaClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewXooaClient(cfg *xooa.Configuration)  *XooaClient{
	return &XooaClient{
		Xooa: xooa.NewXooaClient(cfg),
	}
}

// SetUrl setting Base path from Outside the SDK
func SetUrl(url string){
	xooa.SetUrl(url)
}

//Token for auth in Api
func SetToken(token string){
	xooa.SetToken(token)

}

//Setting Log Level For SDK
func SetLogLevel(logLevel string){
	xooa.SetLogLevel(logLevel)

}
/* Get specific block information such as hash, # of transactions
 * @param ctx context.Context for logging, tracing, etc.
 @param blockNumber Block number to fetch data
 @param optional (nil or map[string]interface{}) with one or more of:
 @param "async" (string) Call this request asynchronously without waiting for response
 @param "timeout" (string) Request timeout in millisecond
 @return */
func (x *XooaClient) GetBlockByNumber(ctx context.Context, blockNumber string, localVarOptionals map[string]interface{}) (models.BlockTransactionCount, models.PendingTransactionResponse, *models.XooaApiException) {
	return x.Xooa.BlockchainApi.GetBlockByNumber(ctx, blockNumber, localVarOptionals)
}
/* Get specific block information such as hash, # of transactions
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param blockNumber Block number to fetch data
 @param optional (nil or map[string]interface{}) with one or more of:
 @param "async" (string) Call this request asynchronously without waiting for response
 @param "timeout" (string) Request timeout in millisecond
 @return */
func (x *XooaClient) GetBlockByNumberAsync(ctx context.Context, blockNumber string, localVarOptionals map[string]interface{}) (models.BlockTransactionCount, models.PendingTransactionResponse, *models.XooaApiException) {
	return x.Xooa.BlockchainApi.GetBlockByNumberAsync(ctx, blockNumber, localVarOptionals)
}
/* Get current blocks in the network
 * @param ctx context.Context for logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @param "async" (string) Call this request asynchronously without waiting for response
 @param "timeout" (string) Request timeout in millisecond
 @return */
func (x *XooaClient) GetCurrentBlock(ctx context.Context, localVarOptionals map[string]interface{}) (models.Block, models.PendingTransactionResponse, *models.XooaApiException) {
	return x.Xooa.BlockchainApi.GetCurrentBlock(ctx, localVarOptionals)
}

/* Get current blocks in the network
* @param ctx context.Context for logging, tracing, etc.
@param optional (nil or map[string]interface{}) with one or more of:
@param "async" (string) Call this request asynchronously without waiting for response
@param "timeout" (string) Request timeout in millisecond
@return */
func (x *XooaClient) GetCurrentBlockAsync(ctx context.Context, localVarOptionals map[string]interface{}) (models.Block, models.PendingTransactionResponse, *models.XooaApiException) {
	return x.Xooa.BlockchainApi.GetCurrentBlockAsync(ctx, localVarOptionals)
}

//Identity Api's

/* This End Point Returns Information about the Authenticated Identity
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @return */
func (x *XooaClient) AuthenticatedIdentityInformation(ctx context.Context) (models.Identity, *models.XooaApiException) {
	return x.Xooa.IdentitiesApi.AuthenticatedIdentityInformation(ctx)
}

/* Get all identities from the identity registry
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @return */
func (x *XooaClient) IdentitiesGetAllIdentities(ctx context.Context) ([]models.Identity, *models.XooaApiException) {
	return x.Xooa.IdentitiesApi.IdentitiesGetAllIdentities(ctx)
}

/* Get the specified identity from the identity registry
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param id Identity Id
 @return */
func (x *XooaClient) IdentityInformation(ctx context.Context, id string) (models.Identity, *models.XooaApiException) {
	return x.Xooa.IdentitiesApi.IdentityInformation(ctx, id)
}

/* The Enroll Identity end point is used to  enroll new identities for the Smart Contract app.  A success response includes the API Key generated for the identity. This API Key can be used to call API End points on behalf of the enrolled identity. This End Point provides equivalent functionality to adding new identity manually using Xooa console, and identities added using this end point will appear, and can be managed, using Xooa console under the identities tab of the Smart Contract app Required permission: manage identities (canManageIdentities&#x3D;true)
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @param "async" (string) Call this request asynchronously without waiting for response
 @param "timeout" (string) Request timeout
 @return */
func (x *XooaClient) EnrollIdentity(ctx context.Context, localVarOptionals map[string]interface{}, localVarPostBody interface{}) (models.Identity, models.PendingTransactionResponse, *models.XooaApiException) {
	return x.Xooa.IdentitiesApi.EnrollIdentity(ctx, localVarOptionals, localVarPostBody)
}

/* The Enroll Identity end point is used to  enroll new identities for the Smart Contract app.  A success response includes the API Key generated for the identity. This API Key can be used to call API End points on behalf of the enrolled identity. This End Point provides equivalent functionality to adding new identity manually using Xooa console, and identities added using this end point will appear, and can be managed, using Xooa console under the identities tab of the Smart Contract app Required permission: manage identities (canManageIdentities&#x3D;true)
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
 @param "async" (string) Call this request asynchronously without waiting for response
 @param "timeout" (string) Request timeout
 @return */
func (x *XooaClient) EnrollIdentityAsync(ctx context.Context, localVarOptionals map[string]interface{}, localVarPostBody interface{}) (models.Identity, models.PendingTransactionResponse, *models.XooaApiException) {
	return x.Xooa.IdentitiesApi.EnrollIdentityAsync(ctx, localVarOptionals, localVarPostBody)
}

/* Delete Identity.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param id Identity Id
 @param optional (nil or map[string]interface{}) with one or more of:
 @param "async" (string) Call this request asynchronously without waiting for response
 @param "timeout" (string) Request timeout in millisecond
 @return */

func (x *XooaClient) DeleteIdentity(ctx context.Context, id string, localVarOptionals map[string]interface{}) (models.DeleteIdentityResponse, models.PendingTransactionResponse, *models.XooaApiException) {
	return x.Xooa.IdentitiesApi.DeleteIdentity(ctx, id, localVarOptionals)
}

/* Delete Identity.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param id Identity Id
 @param optional (nil or map[string]interface{}) with one or more of:
 @param "async" (string) Call this request asynchronously without waiting for response
 @param "timeout" (string) Request timeout in millisecond
 @return */
func (x *XooaClient) DeleteIdentityAsync(ctx context.Context, id string, localVarOptionals map[string]interface{}) (models.DeleteIdentityResponse, models.PendingTransactionResponse, *models.XooaApiException) {
	return x.Xooa.IdentitiesApi.DeleteIdentityAsync(ctx, id, localVarOptionals)
}

/* Regenerate Identity API Token
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param id Identity Id
 @return */
func (x *XooaClient) RegenerateToken(ctx context.Context, id string, localVarOptionals map[string]interface{}) (models.Identity, models.PendingTransactionResponse, *models.XooaApiException) {
	return x.Xooa.IdentitiesApi.RegenerateToken(ctx, id, localVarOptionals)
}

/* Regenerate Identity API Token
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param id Identity Id
 @return */

func (x *XooaClient) RegenerateTokenAsync(ctx context.Context, id string, localVarOptionals map[string]interface{}) (models.Identity, models.PendingTransactionResponse, *models.XooaApiException) {
	return x.Xooa.IdentitiesApi.RegenerateTokenAsync(ctx, id, localVarOptionals)
}

// Invoke Api's
/* The Invoke API End Point is used for submitting transaction for processing by the blockchain Smart Contract app. The end point must call a function already defined in your Smart Contract app which will process the Invoke request. The function name is part of the endpoint URL, or can be entered as the fcn parameter  when testing using the Sandbox. For example, if testing the sample get-set Smart Contract app, use ‘set’ (without quotes)  as the value for fcn.   The function arguments (number of arguments and type) is determined by the Smart Contract. The Smart Contract is also responsible for arguments validation and exception management. The payload object of Invoke Transaction Response in case of Final Response is determined by the Smart Contract app.   A success response may be either 200 or 202. For more details refer to Synchronous vs Asynchronous Calls Required permission: write (\&quot;Access\&quot;:\&quot;rw\&quot;)
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param fcn The blockchain Smart Contract app function name to invoke
 @param optional (nil or map[string]interface{}) with one or more of:
 @param "async" (string) Call this request asynchronously without waiting for response
 @param "timeout" (string) Request timeout in millisecond
 @return */
func (x *XooaClient) Invoke(ctx context.Context, fcn string, localVarOptionals map[string]interface{}, localVarPostBody interface{}) (models.InvokeResponse, models.PendingTransactionResponse, *models.XooaApiException) {
	return x.Xooa.InvokeApi.Invoke(ctx, fcn, localVarOptionals,localVarPostBody)
}

/* The Invoke API End Point is used for submitting transaction for processing by the blockchain Smart Contract app. The end point must call a function already defined in your Smart Contract app which will process the Invoke request. The function name is part of the endpoint URL, or can be entered as the fcn parameter  when testing using the Sandbox. For example, if testing the sample get-set Smart Contract app, use ‘set’ (without quotes)  as the value for fcn.   The function arguments (number of arguments and type) is determined by the Smart Contract. The Smart Contract is also responsible for arguments validation and exception management. The payload object of Invoke Transaction Response in case of Final Response is determined by the Smart Contract app.   A success response may be either 200 or 202. For more details refer to Synchronous vs Asynchronous Calls Required permission: write (\&quot;Access\&quot;:\&quot;rw\&quot;)
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param fcn The blockchain Smart Contract app function name to invoke
 @param optional (nil or map[string]interface{}) with one or more of:
 @param "timeout" (string) Request timeout in millisecond
 @return */
func (x *XooaClient) InvokeAsync(ctx context.Context, fcn string, localVarOptionals map[string]interface{}, localVarPostBody interface{}) (models.InvokeResponse, models.PendingTransactionResponse, *models.XooaApiException) {
	return x.Xooa.InvokeApi.InvokeAsync(ctx, fcn, localVarOptionals,localVarPostBody)
}

//Query Api's
/* The query API End Point is used for querying blockchain state. The end point must call a function already defined in your Smart Contract app which will process the query request. The function name is part of the endpoint URL, or can be entered as the fcn parameter  when testing using the Sandbox. The function arguments (number of arguments and type) is determined by the Smart Contract. The Smart Contract is responsible for validation and exception management. For example, if testing the sample get-set Smart Contract app, enter ‘get’ (without quotes) as the value for fcn.   The response body is also determined by the Smart Contract app, and that’s also the reason why a consistent response sample is unavailable for this end point. A success response may be either 200 or 202. For more details refer to Synchronous vs Asynchronous Calls. In contrast to Invoke, the Query end point will often return fast even when called in Synchronous mode  Required permission: read (\&quot;Access\&quot;:\&quot;rw\&quot; or \&quot;Access\&quot;:\&quot;r\&quot;)
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param fcn The blockchain Smart Contract app function name to call
 @param optional (nil or map[string]interface{}) with one or more of:
 @param "args" (string) Argument(s) to pass to the blockchain Smart Contract app function. example - [\&quot;arg1\&quot;,\&quot;arg2\&quot;]
 @param "async" (string) Call this request asynchronously without waiting for response
 @param "timeout" (string) Request timeout in millisecond
 @return */
func (x *XooaClient) Query(ctx context.Context, fcn string, localVarOptionals map[string]interface{}) (models.Query, models.PendingTransactionResponse, *models.XooaApiException) {
	return x.Xooa.QueryApi.Query(ctx, fcn, localVarOptionals)
}

/* The query API End Point is used for querying blockchain state. The end point must call a function already defined in your Smart Contract app which will process the query request. The function name is part of the endpoint URL, or can be entered as the fcn parameter  when testing using the Sandbox. The function arguments (number of arguments and type) is determined by the Smart Contract. The Smart Contract is responsible for validation and exception management. For example, if testing the sample get-set Smart Contract app, enter ‘get’ (without quotes) as the value for fcn.   The response body is also determined by the Smart Contract app, and that’s also the reason why a consistent response sample is unavailable for this end point. A success response may be either 200 or 202. For more details refer to Synchronous vs Asynchronous Calls. In contrast to Invoke, the Query end point will often return fast even when called in Synchronous mode  Required permission: read (\&quot;Access\&quot;:\&quot;rw\&quot; or \&quot;Access\&quot;:\&quot;r\&quot;)
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param fcn The blockchain Smart Contract app function name to call
 @param optional (nil or map[string]interface{}) with one or more of:
 @param "args" (string) Argument(s) to pass to the blockchain Smart Contract app function. example - [\&quot;arg1\&quot;,\&quot;arg2\&quot;]
 @param "async" (string) Call this request asynchronously without waiting for response
 @param "timeout" (string) Request timeout in millisecond
 @return */
func (x *XooaClient) QueryAsync(ctx context.Context, fcn string, localVarOptionals map[string]interface{}) (models.Query, models.PendingTransactionResponse, *models.XooaApiException) {
	return x.Xooa.QueryApi.QueryAsync(ctx, fcn, localVarOptionals)
}


//Result Api's
/* API Returns result of previously submitted transaction
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param resultId Returned in previous Query/Invoke/Participant Operation
 @return */

func (x *XooaClient) Result(ctx context.Context, resultId string) (models.Result, *models.XooaApiException) {
	return x.Xooa.ResultApi.Result(ctx, resultId)
}

/* API Returns result of previously submitted transaction
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param resultId Returned in previous CurrentBlock Operation
 @return */
func (x *XooaClient) GetResultForCurrentBlock(ctx context.Context, resultId string) (models.Block, *models.XooaApiException) {
	return x.Xooa.ResultApi.GetResultForCurrentBlock(ctx, resultId)
}

/* API Returns result of previously submitted transaction
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param resultId Returned in previous BlockByNumber Operation
 @return */
func (x *XooaClient) GetResultForBlockByNumber(ctx context.Context, resultId string) (models.BlockTransactionCount, *models.XooaApiException) {
	return x.Xooa.ResultApi.GetResultForBlockByNumber(ctx, resultId)

}
/* API Returns result of previously submitted transaction
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param resultId Returned in previous Invoke Operation
 @return */
func (x *XooaClient) GetResultForInvoke(ctx context.Context, resultId string) (models.InvokeResponse, *models.XooaApiException) {
	return x.Xooa.ResultApi.GetResultForInvoke(ctx, resultId)
}
/* API Returns result of previously submitted transaction
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param resultId Returned in previous Query Operation
 @return */
func (x *XooaClient) GetResultForQuery(ctx context.Context, resultId string) (models.Query, *models.XooaApiException) {
	return x.Xooa.ResultApi.GetResultForQuery(ctx, resultId)
}

/* API Returns result of previously submitted transaction
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param resultId Returned in previous Enroll Identity Operation
 @return */
func (x *XooaClient) GetResultForEnrollIdentity(ctx context.Context, resultId string) (models.Identity, *models.XooaApiException) {
	return x.Xooa.ResultApi.GetResultForEnrollIdentity(ctx, resultId)
}

/* API Returns result of previously submitted transaction
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param resultId Returned in previous DeleteIdentity Operation
 @return */
func (x *XooaClient) GetResultForDeleteIdentity(ctx context.Context, resultId string) (models.DeleteIdentityResponse, *models.XooaApiException) {
	return x.Xooa.ResultApi.GetResultForDeleteIdentity(ctx, resultId)
}

/* API Returns result of previously submitted transaction
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param resultId Returned in previous RegenerateToken Operation
 @return */
func (x *XooaClient) GetResultForRegenerateTokenAsync(ctx context.Context, resultId string) (models.RegenerateIdentityResponse, *models.XooaApiException) {
	return x.Xooa.ResultApi.GetResultForRegenerateTokenAsync(ctx, resultId)

}


//Event Client Api's
/* Initiating connection to Socket io server
 @param callbackFn function to return the events
 @return error
*/
func (x *XooaClient) ConnectEventServer(callbackFn models.CallbackFn) (error) {
	return x.Xooa.EventClientService.ConnectEventServer(callbackFn)

}

//Putting a Regex `"a*"` to capture all the events
func (x *XooaClient) SubscribeAllEvents() (error) {
	return x.Xooa.EventClientService.SubscribeAllEvents()

}

// UnSubscribing all the eventa
func (x *XooaClient) UnSubscribe()  {
	x.Xooa.EventClientService.UnSubscribe()

}













/*
 * Xooa Blockchain API's
 *
 * List of Xooa Blockchain API's 
 *
 * API version: v1
 * Contact: support@xooa.com
 */

package xooa

import (
	"io/ioutil"
	"net/url"
	"strings"
	"golang.org/x/net/context"
	"fmt"
	"go-client/models"
	"encoding/json"
	log "github.com/sirupsen/logrus"
)

// Linger please
var (
	_ context.Context
)

type QueryApiService service


/* Query - Query Blockchain data
 The query API End Point is used for querying blockchain state. The end point must call a function already defined in your Smart Contract app which will process the query request. The function name is part of the endpoint URL, or can be entered as the fcn parameter  when testing using the Sandbox. The function arguments (number of arguments and type) is determined by the Smart Contract. The Smart Contract is responsible for validation and exception management. For example, if testing the sample get-set Smart Contract app, enter ‘get’ (without quotes) as the value for fcn.   The response body is also determined by the Smart Contract app, and that’s also the reason why a consistent response sample is unavailable for this end point. A success response may be either 200 or 202. For more details refer to Synchronous vs Asynchronous Calls. In contrast to Invoke, the Query end point will often return fast even when called in Synchronous mode  Required permission: read (\&quot;Access\&quot;:\&quot;rw\&quot; or \&quot;Access\&quot;:\&quot;r\&quot;)
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param fcn The blockchain Smart Contract app function name to call
 @param optional (nil or map[string]interface{}) with one or more of:
 @param "args" (string) Argument(s) to pass to the blockchain Smart Contract app function. example - [\&quot;arg1\&quot;,\&quot;arg2\&quot;]
 @param "async" (string) Call this request asynchronously without waiting for response
 @param "timeout" (string) Request timeout in millisecond
 @return */
func (a *QueryApiService) Query(ctx context.Context, fcn string, localVarOptionals map[string]interface{}) ( models.Query,models.PendingTransactionResponse, *models.XooaApiException) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := BasePath + "/query/{fcn}"
	localVarPath = strings.Replace(localVarPath, "{"+"fcn"+"}", fmt.Sprintf("%v", fcn), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	queryResponse := models.Query{}
	pendingTransactionResponse := models.PendingTransactionResponse{}
	async := false

	if err := typeCheckParameter(localVarOptionals["args"], "string", "args"); err != nil {
		log.Error(err)
		return queryResponse, pendingTransactionResponse, &models.XooaApiException{
			ErrorCode: 0,
			ErrorMessage: err.Error(),
		}
	}
	if err := typeCheckParameter(localVarOptionals["async"], "string", "async"); err != nil {
		log.Error(err)
		return queryResponse, pendingTransactionResponse, &models.XooaApiException{
			ErrorCode: 0,
			ErrorMessage: err.Error(),
		}
	}
	if err := typeCheckParameter(localVarOptionals["timeout"], "string", "timeout"); err != nil {
		log.Error(err)
		return queryResponse, pendingTransactionResponse, &models.XooaApiException{
			ErrorCode: 0,
			ErrorMessage: err.Error(),
		}
	}

	if localVarTempParam, localVarOk := localVarOptionals["args"].(string); localVarOk {
		localVarQueryParams.Add("args", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["async"].(string); localVarOk {
		async = true
		localVarQueryParams.Add("async", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["timeout"].(string); localVarOk {
		localVarQueryParams.Add("timeout", parameterToString(localVarTempParam, ""))
	}
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
		return queryResponse, pendingTransactionResponse, &models.XooaApiException{
			ErrorCode: 0,
			ErrorMessage: err.Error(),
		}
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return queryResponse, pendingTransactionResponse, models.NewXooaApiException(localVarHttpResponse)
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		log.Error(localVarHttpResponse.StatusCode,localVarHttpResponse.Body)
		return queryResponse, pendingTransactionResponse, models.NewXooaApiException(localVarHttpResponse)

	}
	content, err := ioutil.ReadAll(localVarHttpResponse.Body)
	if localVarHttpResponse.StatusCode == 202 {
		json.Unmarshal(content, &pendingTransactionResponse)
		log.Warn(localVarHttpResponse.StatusCode,pendingTransactionResponse)
		return queryResponse,pendingTransactionResponse, nil
	}
	if err != nil {
		log.Error(err)
		return queryResponse, pendingTransactionResponse, &models.XooaApiException {
			ErrorCode: 0,
			ErrorMessage: err.Error(),
		}
	}
	if async == true {
		json.Unmarshal(content, &pendingTransactionResponse)
		log.Info(pendingTransactionResponse)
	} else {
		json.Unmarshal(content, &queryResponse)
		log.Info(queryResponse)
	}

	return queryResponse,pendingTransactionResponse, nil
}


/* QueryAsync Query Blockchain data
 The query API End Point is used for querying blockchain state. The end point must call a function already defined in your Smart Contract app which will process the query request. The function name is part of the endpoint URL, or can be entered as the fcn parameter  when testing using the Sandbox. The function arguments (number of arguments and type) is determined by the Smart Contract. The Smart Contract is responsible for validation and exception management. For example, if testing the sample get-set Smart Contract app, enter ‘get’ (without quotes) as the value for fcn.   The response body is also determined by the Smart Contract app, and that’s also the reason why a consistent response sample is unavailable for this end point. A success response may be either 200 or 202. For more details refer to Synchronous vs Asynchronous Calls. In contrast to Invoke, the Query end point will often return fast even when called in Synchronous mode  Required permission: read (\&quot;Access\&quot;:\&quot;rw\&quot; or \&quot;Access\&quot;:\&quot;r\&quot;)
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param fcn The blockchain Smart Contract app function name to call
 @param optional (nil or map[string]interface{}) with one or more of:
 @param "args" (string) Argument(s) to pass to the blockchain Smart Contract app function. example - [\&quot;arg1\&quot;,\&quot;arg2\&quot;]
 @param "async" (string) Call this request asynchronously without waiting for response
 @param "timeout" (string) Request timeout in millisecond
 @return */
func (a *QueryApiService) QueryAsync(ctx context.Context, fcn string, localVarOptionals map[string]interface{}) ( models.Query,models.PendingTransactionResponse ,*models.XooaApiException) {
	localVarOptionals["async"] = "true"
	return a.Query(ctx,fcn,localVarOptionals)

}


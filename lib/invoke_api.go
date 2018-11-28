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

type InvokeApiService service


/* Invoke Submit blockchain transaction
 The Invoke API End Point is used for submitting transaction for processing by the blockchain Smart Contract app. The end point must call a function already defined in your Smart Contract app which will process the Invoke request. The function name is part of the endpoint URL, or can be entered as the fcn parameter  when testing using the Sandbox. For example, if testing the sample get-set Smart Contract app, use ‘set’ (without quotes)  as the value for fcn.   The function arguments (number of arguments and type) is determined by the Smart Contract. The Smart Contract is also responsible for arguments validation and exception management. The payload object of Invoke Transaction Response in case of Final Response is determined by the Smart Contract app.   A success response may be either 200 or 202. For more details refer to Synchronous vs Asynchronous Calls Required permission: write (\&quot;Access\&quot;:\&quot;rw\&quot;)
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param fcn The blockchain Smart Contract app function name to invoke
 @param optional (nil or map[string]interface{}) with one or more of:
 @param "async" (string) Call this request asynchronously without waiting for response
 @param "timeout" (string) Request timeout in millisecond
 @return */
func (a *InvokeApiService) Invoke(ctx context.Context, fcn string, localVarOptionals map[string]interface{},localVarPostBody interface{}) ( models.InvokeResponse,models.PendingTransactionResponse ,*models.XooaApiException) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarFileName string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := BasePath + "/invoke/{fcn}"
	localVarPath = strings.Replace(localVarPath, "{"+"fcn"+"}", fmt.Sprintf("%v", fcn), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	invokeResponse := models.InvokeResponse{}
	pendingTransactionResponse := models.PendingTransactionResponse{}
	async := false

	if err := typeCheckParameter(localVarOptionals["async"], "string", "async"); err != nil {
		log.Error(err)
		return invokeResponse, pendingTransactionResponse,  &models.XooaApiException{
			ErrorCode: 0,
			ErrorMessage: err.Error(),
		}
	}
	if err := typeCheckParameter(localVarOptionals["timeout"], "string", "timeout"); err != nil {
		log.Error(err)
		return invokeResponse, pendingTransactionResponse, &models.XooaApiException{
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
		return invokeResponse, pendingTransactionResponse, &models.XooaApiException{
			ErrorCode: 0,
			ErrorMessage: err.Error(),
		}
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return invokeResponse, pendingTransactionResponse, models.NewXooaApiException(localVarHttpResponse)
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		log.Error(localVarHttpResponse.StatusCode,localVarHttpResponse)
		return invokeResponse, pendingTransactionResponse, models.NewXooaApiException(localVarHttpResponse)
	}
	content, err := ioutil.ReadAll(localVarHttpResponse.Body)
	if localVarHttpResponse.StatusCode == 202 {
		json.Unmarshal(content, &pendingTransactionResponse)
		log.Warn(localVarHttpResponse.StatusCode,pendingTransactionResponse)
		return invokeResponse,pendingTransactionResponse, nil
	}
	if err != nil {
		log.Error(err)
		return invokeResponse, pendingTransactionResponse, &models.XooaApiException {
			ErrorCode: 0,
			ErrorMessage: err.Error(),
		}
	}
	if async == true {
		json.Unmarshal(content, &pendingTransactionResponse)
		log.Info(pendingTransactionResponse)

	} else {
		json.Unmarshal(content, &invokeResponse)
		log.Info(invokeResponse)
	}

	return invokeResponse,pendingTransactionResponse, nil
}


/* InvokeAsync Submit blockchain transaction
 The Invoke API End Point is used for submitting transaction for processing by the blockchain Smart Contract app. The end point must call a function already defined in your Smart Contract app which will process the Invoke request. The function name is part of the endpoint URL, or can be entered as the fcn parameter  when testing using the Sandbox. For example, if testing the sample get-set Smart Contract app, use ‘set’ (without quotes)  as the value for fcn.   The function arguments (number of arguments and type) is determined by the Smart Contract. The Smart Contract is also responsible for arguments validation and exception management. The payload object of Invoke Transaction Response in case of Final Response is determined by the Smart Contract app.   A success response may be either 200 or 202. For more details refer to Synchronous vs Asynchronous Calls Required permission: write (\&quot;Access\&quot;:\&quot;rw\&quot;)
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param fcn The blockchain Smart Contract app function name to invoke
 @param optional (nil or map[string]interface{}) with one or more of:
 @param "async" (string) Call this request asynchronously without waiting for response
 @param "timeout" (string) Request timeout in millisecond
 @return */
func (a *InvokeApiService) InvokeAsync(ctx context.Context, fcn string, localVarOptionals map[string]interface{},localVarPostBody interface{}) ( models.InvokeResponse,models.PendingTransactionResponse ,*models.XooaApiException) {
	localVarOptionals["async"] = "true"
	return a.Invoke(ctx,fcn,localVarOptionals,localVarPostBody )
}

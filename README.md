# Xooa Go SDK
## Installation

```bash
go get github.com/Xooa/xooa-go-sdk
```

#### Import
 ```go
 import "github.com/Xooa/xooa-go-sdk"
 ```
 
 
## Usage

#### func  SetLogLevel

```go
func SetLogLevel(logLevel string)
```
Setting Log Level For SDK

#### func  SetToken

```go
func SetToken(token string)
```
Token for auth in Api

#### func  SetUrl

```go
func SetUrl(url string)
```
SetUrl setting Base path from Outside the SDK

#### type XooaClient

```go
type XooaClient struct {
	Xooa *xooa.XooaClient
}
```


#### func  NewXooaClient

```go
func NewXooaClient(cfg *xooa.Configuration) *XooaClient
```
NewXooaClient creates a new API client. Requires a userAgent string describing
your application. optionally a custom http.Client to allow for advanced features
such as caching.

#### func (*XooaClient) AuthenticatedIdentityInformation

```go
func (x *XooaClient) AuthenticatedIdentityInformation(ctx context.Context) (models.Identity, *models.XooaApiException)
```
This End Point Returns Information about the Authenticated Identity 

#### func (*XooaClient) ConnectEventServer

```go
func (x *XooaClient) ConnectEventServer(callbackFn models.CallbackFn) error
```

Initiating connection to Socket io server


#### func (*XooaClient) DeleteIdentity

```go
func (x *XooaClient) DeleteIdentity(ctx context.Context, id string, localVarOptionals map[string]interface{}) (models.DeleteIdentityResponse, models.PendingTransactionResponse, *models.XooaApiException)
```

#### func (*XooaClient) DeleteIdentityAsync

```go
func (x *XooaClient) DeleteIdentityAsync(ctx context.Context, id string, localVarOptionals map[string]interface{}) (models.DeleteIdentityResponse, models.PendingTransactionResponse, *models.XooaApiException)
```
Delete Identity. 

#### func (*XooaClient) EnrollIdentity

```go
func (x *XooaClient) EnrollIdentity(ctx context.Context, localVarOptionals map[string]interface{}, localVarPostBody interface{}) (models.Identity, models.PendingTransactionResponse, *models.XooaApiException)
```
The Enroll Identity end point is used to enroll new identities for the Smart
Contract app. A success response includes the API Key generated for the
identity. This API Key can be used to call API End points on behalf of the
enrolled identity. This End Point provides equivalent functionality to adding
new identity manually using Xooa console, and identities added using this end
point will appear, and can be managed, using Xooa console under the identities
tab of the Smart Contract app Required permission: manage identities
(canManageIdentities&#x3D;true)

#### func (*XooaClient) EnrollIdentityAsync

```go
func (x *XooaClient) EnrollIdentityAsync(ctx context.Context, localVarOptionals map[string]interface{}, localVarPostBody interface{}) (models.Identity, models.PendingTransactionResponse, *models.XooaApiException)
```
The Enroll Identity end point is used to enroll new identities for the Smart
Contract app. A success response includes the API Key generated for the
identity. This API Key can be used to call API End points on behalf of the
enrolled identity. This End Point provides equivalent functionality to adding
new identity manually using Xooa console, and identities added using this end
point will appear, and can be managed, using Xooa console under the identities
tab of the Smart Contract app Required permission: manage identities
(canManageIdentities&#x3D;true) 

#### func (*XooaClient) GetBlockByNumber

```go
func (x *XooaClient) GetBlockByNumber(ctx context.Context, blockNumber string, localVarOptionals map[string]interface{}) (models.BlockTransactionCount, models.PendingTransactionResponse, *models.XooaApiException)
```
Get specific block information such as hash, # of transactions * @param ctx
context.Context for logging, tracing, etc.

#### func (*XooaClient) GetBlockByNumberAsync

```go
func (x *XooaClient) GetBlockByNumberAsync(ctx context.Context, blockNumber string, localVarOptionals map[string]interface{}) (models.BlockTransactionCount, models.PendingTransactionResponse, *models.XooaApiException)
```
Get specific block information such as hash, # of transactions 

#### func (*XooaClient) GetCurrentBlock

```go
func (x *XooaClient) GetCurrentBlock(ctx context.Context, localVarOptionals map[string]interface{}) (models.Block, models.PendingTransactionResponse, *models.XooaApiException)
```
Get current blocks in the network 

#### func (*XooaClient) GetCurrentBlockAsync

```go
func (x *XooaClient) GetCurrentBlockAsync(ctx context.Context, localVarOptionals map[string]interface{}) (models.Block, models.PendingTransactionResponse, *models.XooaApiException)
```

Get current blocks in the network

#### func (*XooaClient) GetResultForBlockByNumber

```go
func (x *XooaClient) GetResultForBlockByNumber(ctx context.Context, resultId string) (models.BlockTransactionCount, *models.XooaApiException)
```
API Returns result of previously submitted transaction 

#### func (*XooaClient) GetResultForCurrentBlock

```go
func (x *XooaClient) GetResultForCurrentBlock(ctx context.Context, resultId string) (models.Block, *models.XooaApiException)
```
API Returns result of previously submitted transaction
#### func (*XooaClient) GetResultForDeleteIdentity

```go
func (x *XooaClient) GetResultForDeleteIdentity(ctx context.Context, resultId string) (models.DeleteIdentityResponse, *models.XooaApiException)
```
API Returns result of previously submitted transaction 

#### func (*XooaClient) GetResultForEnrollIdentity

```go
func (x *XooaClient) GetResultForEnrollIdentity(ctx context.Context, resultId string) (models.Identity, *models.XooaApiException)
```
API Returns result of previously submitted transaction 

#### func (*XooaClient) GetResultForInvoke

```go
func (x *XooaClient) GetResultForInvoke(ctx context.Context, resultId string) (models.InvokeResponse, *models.XooaApiException)
```
API Returns result of previously submitted transaction 

#### func (*XooaClient) GetResultForQuery

```go
func (x *XooaClient) GetResultForQuery(ctx context.Context, resultId string) (models.Query, *models.XooaApiException)
```
API Returns result of previously submitted transaction 

#### func (*XooaClient) GetResultForRegenerateTokenAsync

```go
func (x *XooaClient) GetResultForRegenerateTokenAsync(ctx context.Context, resultId string) (models.RegenerateIdentityResponse, *models.XooaApiException)
```
API Returns result of previously submitted transaction 

#### func (*XooaClient) IdentitiesGetAllIdentities

```go
func (x *XooaClient) IdentitiesGetAllIdentities(ctx context.Context) ([]models.Identity, *models.XooaApiException)
```
Get all identities from the identity registry 

#### func (*XooaClient) IdentityInformation

```go
func (x *XooaClient) IdentityInformation(ctx context.Context, id string) (models.Identity, *models.XooaApiException)
```
Get the specified identity from the identity registry 

#### func (*XooaClient) Invoke

```go
func (x *XooaClient) Invoke(ctx context.Context, fcn string, localVarOptionals map[string]interface{}, localVarPostBody interface{}) (models.InvokeResponse, models.PendingTransactionResponse, *models.XooaApiException)
```
The Invoke API End Point is used for submitting transaction for processing by the blockchain Smart Contract app. The end point must call a function already defined in your Smart Contract app which will process the Invoke request. The function name is part of the endpoint URL, or can be entered as the fcn parameter  when testing using the Sandbox. For example, if testing the sample get-set Smart Contract app, use ‘set’ (without quotes)  as the value for fcn.   The function arguments (number of arguments and type) is determined by the Smart Contract. The Smart Contract is also responsible for arguments validation and exception management. The payload object of Invoke Transaction Response in case of Final Response is determined by the Smart Contract app.   A success response may be either 200 or 202. For more details refer to Synchronous vs Asynchronous Calls Required permission: write ("Access":"rw")

#### func (*XooaClient) InvokeAsync

```go
func (x *XooaClient) InvokeAsync(ctx context.Context, fcn string, localVarOptionals map[string]interface{}, localVarPostBody interface{}) (models.InvokeResponse, models.PendingTransactionResponse, *models.XooaApiException)
```
The Invoke API End Point is used for submitting transaction for processing by
the blockchain Smart Contract app. The end point must call a function already
defined in your Smart Contract app which will process the Invoke request. The
function name is part of the endpoint URL, or can be entered as the fcn
parameter when testing using the Sandbox. For example, if testing the sample
get-set Smart Contract app, use ‘set’ (without quotes) as the value for fcn. The
function arguments (number of arguments and type) is determined by the Smart
Contract. The Smart Contract is also responsible for arguments validation and
exception management. The payload object of Invoke Transaction Response in case
of Final Response is determined by the Smart Contract app. A success response
may be either 200 or 202. For more details refer to Synchronous vs Asynchronous
Calls Required permission: write ("Access":"rw") 

#### func (*XooaClient) Query

```go
func (x *XooaClient) Query(ctx context.Context, fcn string, localVarOptionals map[string]interface{}) (models.Query, models.PendingTransactionResponse, *models.XooaApiException)
```

The query API End Point is used for querying blockchain state. The end point must call a function already defined in your Smart Contract app which will process the query request. The function name is part of the endpoint URL, or can be entered as the fcn parameter  when testing using the Sandbox. The function arguments (number of arguments and type) is determined by the Smart Contract. The Smart Contract is responsible for validation and exception management. For example, if testing the sample get-set Smart Contract app, enter ‘get’ (without quotes) as the value for fcn.   The response body is also determined by the Smart Contract app, and that’s also the reason why a consistent response sample is unavailable for this end point. A success response may be either 200 or 202. For more details refer to Synchronous vs Asynchronous Calls. In contrast to Invoke, the Query end point will often return fast even when called in Synchronous mode  Required permission: read ("Access":"rw" or "Access":"r")


#### func (*XooaClient) QueryAsync

```go
func (x *XooaClient) QueryAsync(ctx context.Context, fcn string, localVarOptionals map[string]interface{}) (models.Query, models.PendingTransactionResponse, *models.XooaApiException)
```
The query API End Point is used for querying blockchain state. The end point
must call a function already defined in your Smart Contract app which will
process the query request. The function name is part of the endpoint URL, or can
be entered as the fcn parameter when testing using the Sandbox. The function
arguments (number of arguments and type) is determined by the Smart Contract.
The Smart Contract is responsible for validation and exception management. For
example, if testing the sample get-set Smart Contract app, enter ‘get’ (without
quotes) as the value for fcn. The response body is also determined by the Smart
Contract app, and that’s also the reason why a consistent response sample is
unavailable for this end point. A success response may be either 200 or 202. For
more details refer to Synchronous vs Asynchronous Calls. In contrast to Invoke,
the Query end point will often return fast even when called in Synchronous mode
Required permission: read ("Access":"rw" or
"Access":"r") 

#### func (*XooaClient) RegenerateToken

```go
func (x *XooaClient) RegenerateToken(ctx context.Context, id string, localVarOptionals map[string]interface{}) (models.Identity, models.PendingTransactionResponse, *models.XooaApiException)
```
Regenerate Identity API Token 

#### func (*XooaClient) RegenerateTokenAsync

```go
func (x *XooaClient) RegenerateTokenAsync(ctx context.Context, id string, localVarOptionals map[string]interface{}) (models.Identity, models.PendingTransactionResponse, *models.XooaApiException)
```

#### func (*XooaClient) Result

```go
func (x *XooaClient) Result(ctx context.Context, resultId string) (models.Result, *models.XooaApiException)
```

#### func (*XooaClient) SubscribeAllEvents

```go
func (x *XooaClient) SubscribeAllEvents() error
```

#### func (*XooaClient) UnSubscribe

```go
func (x *XooaClient) UnSubscribe()
```
UnSubscribing all the event

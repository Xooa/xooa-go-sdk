# xooa
--
    import "go-client/lib"


## Usage

```go
var BasePath string = "https://api.ci.xooa.io/api/v1"
```
Base Path of Api's we have used

```go
var C *gosocketio.Client
```
Maintaining Socket io client to support connection close

```go
var EventPool string
```
Checking for the last TxId to avoid Duplication

```go
var LogLevel string = "error"
```

```go
var RegExp = "a*"
```
Reg exp to for all events

```go
var RegExpI *regexp.Regexp
```
Regex instance for handling Event Subscription

```go
var (
	Token string
)
```

#### func  CacheExpires

```go
func CacheExpires(r *http.Response) time.Time
```
CacheExpires helper function to determine remaining time before repeating a
request.

#### func  GetLogLevel

```go
func GetLogLevel(level string) (log.Level, error)
```

#### func  SetBasePath

```go
func SetBasePath(urlI string)
```
SetBasePath setting Base path from Outside the SDK

#### func  SetLogLevel

```go
func SetLogLevel(logLevel string) error
```
Setting Log Level For SDK

#### func  SetToken

```go
func SetToken(token string) error
```
Token for auth in Api

#### type BlockchainApiService

```go
type BlockchainApiService service
```


#### func (*BlockchainApiService) GetBlockByNumber

```go
func (a *BlockchainApiService) GetBlockByNumber(ctx context.Context, blockNumber string, localVarOptionals map[string]interface{}) (models.BlockTransactionCount, models.PendingTransactionResponse, *models.XooaApiException)
```

#### func (*BlockchainApiService) GetBlockByNumberAsync

```go
func (a *BlockchainApiService) GetBlockByNumberAsync(ctx context.Context, blockNumber string, localVarOptionals map[string]interface{}) (models.BlockTransactionCount, models.PendingTransactionResponse, *models.XooaApiException)
```

#### func (*BlockchainApiService) GetCurrentBlock

```go
func (a *BlockchainApiService) GetCurrentBlock(ctx context.Context, localVarOptionals map[string]interface{}) (models.Block, models.PendingTransactionResponse, *models.XooaApiException)
```

#### func (*BlockchainApiService) GetCurrentBlockAsync

```go
func (a *BlockchainApiService) GetCurrentBlockAsync(ctx context.Context, localVarOptionals map[string]interface{}) (models.Block, models.PendingTransactionResponse, *models.XooaApiException)
```

#### type Configuration

```go
type Configuration struct {
	Host          string            `json:"host,omitempty"`
	Scheme        string            `json:"scheme,omitempty"`
	DefaultHeader map[string]string `json:"defaultHeader,omitempty"`
	UserAgent     string            `json:"userAgent,omitempty"`
	HTTPClient    *http.Client
}
```

Configuration for setting extra properties of the XooaClient Object

#### func  NewConfiguration

```go
func NewConfiguration() (*Configuration, error)
```
Creating new Configuration Object

#### func (*Configuration) AddDefaultHeader

```go
func (c *Configuration) AddDefaultHeader(key string, value string)
```
Adding default Headers

#### func (*Configuration) SettingUserAgent

```go
func (c *Configuration) SettingUserAgent(key string, value string)
```
SettingUserAgent Setting User agent

#### type EventClientService

```go
type EventClientService service
```


#### func (*EventClientService) ConnectEventServer

```go
func (a *EventClientService) ConnectEventServer(callbackFn models.CallbackFn) error
```
Initiating connection to Socket io server @param callbackFn function to return
the events @return error

#### func (*EventClientService) SubscribeAllEvents

```go
func (a *EventClientService) SubscribeAllEvents() error
```
Putting a Regex `"a*"` to capture all the events

#### func (*EventClientService) SubscribeEvents

```go
func (a *EventClientService) SubscribeEvents(regEx string) error
```
TO capture the specific events specified in the call

#### func (*EventClientService) UnSubscribe

```go
func (a *EventClientService) UnSubscribe()
```
UnSubscribing all the eventa

#### type IdentitiesApiService

```go
type IdentitiesApiService service
```


#### func (*IdentitiesApiService) AuthenticatedIdentityInformation

```go
func (a *IdentitiesApiService) AuthenticatedIdentityInformation(ctx context.Context) (models.Identity, *models.XooaApiException)
```
AuthenticatedIdentityInformation Authenticated Identity Information This End
Point Returns Information about the Authenticated Identity * @param ctx
context.Context for authentication, logging, tracing, etc. @return

#### func (*IdentitiesApiService) DeleteIdentity

```go
func (a *IdentitiesApiService) DeleteIdentity(ctx context.Context, id string, localVarOptionals map[string]interface{}) (models.DeleteIdentityResponse, models.PendingTransactionResponse, *models.XooaApiException)
```
DeleteIdentity Delete Identity Delete Identity. * @param ctx context.Context for
authentication, logging, tracing, etc. @param id Identity Id @param optional
(nil or map[string]interface{}) with one or more of: @param "async" (string)
Call this request asynchronously without waiting for response @param "timeout"
(string) Request timeout in millisecond @return

#### func (*IdentitiesApiService) DeleteIdentityAsync

```go
func (a *IdentitiesApiService) DeleteIdentityAsync(ctx context.Context, id string, localVarOptionals map[string]interface{}) (models.DeleteIdentityResponse, models.PendingTransactionResponse, *models.XooaApiException)
```

#### func (*IdentitiesApiService) EnrollIdentity

```go
func (a *IdentitiesApiService) EnrollIdentity(ctx context.Context, localVarOptionals map[string]interface{}, localVarPostBody interface{}) (models.Identity, models.PendingTransactionResponse, *models.XooaApiException)
```
EnrollIdentity Enroll Identity The Enroll Identity end point is used to enroll
new identities for the Smart Contract app. A success response includes the API
Key generated for the identity. This API Key can be used to call API End points
on behalf of the enrolled identity. This End Point provides equivalent
functionality to adding new identity manually using Xooa console, and identities
added using this end point will appear, and can be managed, using Xooa console
under the identities tab of the Smart Contract app Required permission: manage
identities (canManageIdentities&#x3D;true) * @param ctx context.Context for
authentication, logging, tracing, etc. @param optional (nil or
map[string]interface{}) with one or more of: @param "async" (string) Call this
request asynchronously without waiting for response @param "timeout" (string)
Request timeout @return

#### func (*IdentitiesApiService) EnrollIdentityAsync

```go
func (a *IdentitiesApiService) EnrollIdentityAsync(ctx context.Context, localVarOptionals map[string]interface{}, localVarPostBody interface{}) (models.Identity, models.PendingTransactionResponse, *models.XooaApiException)
```
EnrollIdentityAsync Enroll Identity Async The Enroll Identity end point is used
to enroll new identities for the Smart Contract app. A success response includes
the API Key generated for the identity. This API Key can be used to call API End
points on behalf of the enrolled identity. This End Point provides equivalent
functionality to adding new identity manually using Xooa console, and identities
added using this end point will appear, and can be managed, using Xooa console
under the identities tab of the Smart Contract app Required permission: manage
identities (canManageIdentities&#x3D;true) * @param ctx context.Context for
authentication, logging, tracing, etc. @param optional (nil or
map[string]interface{}) with one or more of: @param "async" (string) Call this
request asynchronously without waiting for response @param "timeout" (string)
Request timeout @return

#### func (*IdentitiesApiService) IdentitiesGetAllIdentities

```go
func (a *IdentitiesApiService) IdentitiesGetAllIdentities(ctx context.Context) ([]models.Identity, *models.XooaApiException)
```
IdentitiesGetAllIdentities Get All Identities Get all identities from the
identity registry * @param ctx context.Context for authentication, logging,
tracing, etc. @return

#### func (*IdentitiesApiService) IdentityInformation

```go
func (a *IdentitiesApiService) IdentityInformation(ctx context.Context, id string) (models.Identity, *models.XooaApiException)
```
IdentityInformation Identity Information Get the specified identity from the
identity registry * @param ctx context.Context for authentication, logging,
tracing, etc. @param id Identity Id @return

#### func (*IdentitiesApiService) RegenerateToken

```go
func (a *IdentitiesApiService) RegenerateToken(ctx context.Context, id string, localVarOptionals map[string]interface{}) (models.Identity, models.PendingTransactionResponse, *models.XooaApiException)
```
RegenerateToken Regenerate Identity API Token Regenerate Identity API Token *
@param ctx context.Context for authentication, logging, tracing, etc. @param id
Identity Id @return

#### func (*IdentitiesApiService) RegenerateTokenAsync

```go
func (a *IdentitiesApiService) RegenerateTokenAsync(ctx context.Context, id string, localVarOptionals map[string]interface{}) (models.Identity, models.PendingTransactionResponse, *models.XooaApiException)
```
RegenerateTokenAsync Regenerate Identity API Token Async Regenerate Identity API
Token * @param ctx context.Context for authentication, logging, tracing, etc.
@param id Identity Id @return

#### type InvokeApiService

```go
type InvokeApiService service
```


#### func (*InvokeApiService) Invoke

```go
func (a *InvokeApiService) Invoke(ctx context.Context, fcn string, localVarOptionals map[string]interface{}, localVarPostBody interface{}) (models.InvokeResponse, models.PendingTransactionResponse, *models.XooaApiException)
```
Invoke Submit blockchain transaction The Invoke API End Point is used for
submitting transaction for processing by the blockchain Smart Contract app. The
end point must call a function already defined in your Smart Contract app which
will process the Invoke request. The function name is part of the endpoint URL,
or can be entered as the fcn parameter when testing using the Sandbox. For
example, if testing the sample get-set Smart Contract app, use ‘set’ (without
quotes) as the value for fcn. The function arguments (number of arguments and
type) is determined by the Smart Contract. The Smart Contract is also
responsible for arguments validation and exception management. The payload
object of Invoke Transaction Response in case of Final Response is determined by
the Smart Contract app. A success response may be either 200 or 202. For more
details refer to Synchronous vs Asynchronous Calls Required permission: write
(\&quot;Access\&quot;:\&quot;rw\&quot;) * @param ctx context.Context for
authentication, logging, tracing, etc. @param fcn The blockchain Smart Contract
app function name to invoke @param optional (nil or map[string]interface{}) with
one or more of: @param "async" (string) Call this request asynchronously without
waiting for response @param "timeout" (string) Request timeout in millisecond
@return

#### func (*InvokeApiService) InvokeAsync

```go
func (a *InvokeApiService) InvokeAsync(ctx context.Context, fcn string, localVarOptionals map[string]interface{}, localVarPostBody interface{}) (models.InvokeResponse, models.PendingTransactionResponse, *models.XooaApiException)
```
InvokeAsync Submit blockchain transaction The Invoke API End Point is used for
submitting transaction for processing by the blockchain Smart Contract app. The
end point must call a function already defined in your Smart Contract app which
will process the Invoke request. The function name is part of the endpoint URL,
or can be entered as the fcn parameter when testing using the Sandbox. For
example, if testing the sample get-set Smart Contract app, use ‘set’ (without
quotes) as the value for fcn. The function arguments (number of arguments and
type) is determined by the Smart Contract. The Smart Contract is also
responsible for arguments validation and exception management. The payload
object of Invoke Transaction Response in case of Final Response is determined by
the Smart Contract app. A success response may be either 200 or 202. For more
details refer to Synchronous vs Asynchronous Calls Required permission: write
(\&quot;Access\&quot;:\&quot;rw\&quot;) * @param ctx context.Context for
authentication, logging, tracing, etc. @param fcn The blockchain Smart Contract
app function name to invoke @param optional (nil or map[string]interface{}) with
one or more of: @param "async" (string) Call this request asynchronously without
waiting for response @param "timeout" (string) Request timeout in millisecond
@return

#### type QueryApiService

```go
type QueryApiService service
```


#### func (*QueryApiService) Query

```go
func (a *QueryApiService) Query(ctx context.Context, fcn string, localVarOptionals map[string]interface{}) (models.Query, models.PendingTransactionResponse, *models.XooaApiException)
```
Query - Query Blockchain data The query API End Point is used for querying
blockchain state. The end point must call a function already defined in your
Smart Contract app which will process the query request. The function name is
part of the endpoint URL, or can be entered as the fcn parameter when testing
using the Sandbox. The function arguments (number of arguments and type) is
determined by the Smart Contract. The Smart Contract is responsible for
validation and exception management. For example, if testing the sample get-set
Smart Contract app, enter ‘get’ (without quotes) as the value for fcn. The
response body is also determined by the Smart Contract app, and that’s also the
reason why a consistent response sample is unavailable for this end point. A
success response may be either 200 or 202. For more details refer to Synchronous
vs Asynchronous Calls. In contrast to Invoke, the Query end point will often
return fast even when called in Synchronous mode Required permission: read
(\&quot;Access\&quot;:\&quot;rw\&quot; or \&quot;Access\&quot;:\&quot;r\&quot;)
* @param ctx context.Context for authentication, logging, tracing, etc. @param
fcn The blockchain Smart Contract app function name to call @param optional (nil
or map[string]interface{}) with one or more of: @param "args" (string)
Argument(s) to pass to the blockchain Smart Contract app function. example -
[\&quot;arg1\&quot;,\&quot;arg2\&quot;] @param "async" (string) Call this
request asynchronously without waiting for response @param "timeout" (string)
Request timeout in millisecond @return

#### func (*QueryApiService) QueryAsync

```go
func (a *QueryApiService) QueryAsync(ctx context.Context, fcn string, localVarOptionals map[string]interface{}) (models.Query, models.PendingTransactionResponse, *models.XooaApiException)
```
QueryAsync Query Blockchain data The query API End Point is used for querying
blockchain state. The end point must call a function already defined in your
Smart Contract app which will process the query request. The function name is
part of the endpoint URL, or can be entered as the fcn parameter when testing
using the Sandbox. The function arguments (number of arguments and type) is
determined by the Smart Contract. The Smart Contract is responsible for
validation and exception management. For example, if testing the sample get-set
Smart Contract app, enter ‘get’ (without quotes) as the value for fcn. The
response body is also determined by the Smart Contract app, and that’s also the
reason why a consistent response sample is unavailable for this end point. A
success response may be either 200 or 202. For more details refer to Synchronous
vs Asynchronous Calls. In contrast to Invoke, the Query end point will often
return fast even when called in Synchronous mode Required permission: read
(\&quot;Access\&quot;:\&quot;rw\&quot; or \&quot;Access\&quot;:\&quot;r\&quot;)
* @param ctx context.Context for authentication, logging, tracing, etc. @param
fcn The blockchain Smart Contract app function name to call @param optional (nil
or map[string]interface{}) with one or more of: @param "args" (string)
Argument(s) to pass to the blockchain Smart Contract app function. example -
[\&quot;arg1\&quot;,\&quot;arg2\&quot;] @param "async" (string) Call this
request asynchronously without waiting for response @param "timeout" (string)
Request timeout in millisecond @return

#### type ResultApiService

```go
type ResultApiService service
```


#### func (*ResultApiService) GetResultForBlockByNumber

```go
func (a *ResultApiService) GetResultForBlockByNumber(ctx context.Context, resultId string) (models.BlockTransactionCount, *models.XooaApiException)
```
GetResultForBlockByNumber Fetch result of previously submitted transaction API
Returns result of previously submitted transaction * @param ctx context.Context
for authentication, logging, tracing, etc. @param resultId Returned in previous
BlockByNumber Operation @return

#### func (*ResultApiService) GetResultForCurrentBlock

```go
func (a *ResultApiService) GetResultForCurrentBlock(ctx context.Context, resultId string) (models.Block, *models.XooaApiException)
```
GetResultForCurrentBlock Fetch result of previously submitted transaction API
Returns result of previously submitted transaction * @param ctx context.Context
for authentication, logging, tracing, etc. @param resultId Returned in previous
CurrentBlock Operation @return

#### func (*ResultApiService) GetResultForDeleteIdentity

```go
func (a *ResultApiService) GetResultForDeleteIdentity(ctx context.Context, resultId string) (models.DeleteIdentityResponse, *models.XooaApiException)
```
GetResultForDeleteIdentity Fetch result of previously submitted transaction API
Returns result of previously submitted transaction * @param ctx context.Context
for authentication, logging, tracing, etc. @param resultId Returned in previous
DeleteIdentity Operation @return

#### func (*ResultApiService) GetResultForEnrollIdentity

```go
func (a *ResultApiService) GetResultForEnrollIdentity(ctx context.Context, resultId string) (models.Identity, *models.XooaApiException)
```
GetResultForEnrollIdentity Fetch result of previously submitted transaction API
Returns result of previously submitted transaction * @param ctx context.Context
for authentication, logging, tracing, etc. @param resultId Returned in previous
Enroll Identity Operation @return

#### func (*ResultApiService) GetResultForInvoke

```go
func (a *ResultApiService) GetResultForInvoke(ctx context.Context, resultId string) (models.InvokeResponse, *models.XooaApiException)
```
GetResultForInvoke Fetch result of previously submitted transaction API Returns
result of previously submitted transaction * @param ctx context.Context for
authentication, logging, tracing, etc. @param resultId Returned in previous
Invoke Operation @return

#### func (*ResultApiService) GetResultForQuery

```go
func (a *ResultApiService) GetResultForQuery(ctx context.Context, resultId string) (models.Query, *models.XooaApiException)
```
GetResultForQuery Fetch result of previously submitted transaction API Returns
result of previously submitted transaction * @param ctx context.Context for
authentication, logging, tracing, etc. @param resultId Returned in previous
Query Operation @return

#### func (*ResultApiService) GetResultForRegenerateTokenAsync

```go
func (a *ResultApiService) GetResultForRegenerateTokenAsync(ctx context.Context, resultId string) (models.RegenerateIdentityResponse, *models.XooaApiException)
```
GetResultForRegenerateToken Fetch result of previously submitted transaction API
Returns result of previously submitted transaction * @param ctx context.Context
for authentication, logging, tracing, etc. @param resultId Returned in previous
RegenerateToken Operation @return

#### func (*ResultApiService) Result

```go
func (a *ResultApiService) Result(ctx context.Context, resultId string) (models.Result, *models.XooaApiException)
```
ResultApiService Fetch result of previously submitted transaction API Returns
result of previously submitted transaction * @param ctx context.Context for
authentication, logging, tracing, etc. @param resultId Returned in previous
Query/Invoke/Participant Operation @return

#### type XooaClient

```go
type XooaClient struct {
	AppId string
	Cfg   *Configuration

	LogLevel string
	// API Services
	BlockchainApi      *BlockchainApiService
	IdentitiesApi      *IdentitiesApiService
	InvokeApi          *InvokeApiService
	QueryApi           *QueryApiService
	ResultApi          *ResultApiService
	EventClientService *EventClientService
}
```

XooaClient manages communication with the Xooa Blockchain API&#39;s API vv1 In
most cases there should be only one, shared, XooaClient.

#### func  NewXooaClient

```go
func NewXooaClient(cfg *Configuration) *XooaClient
```
NewXooaClient creates a new API client. Requires a userAgent string describing
your application. optionally a custom http.Client to allow for advanced features
such as caching.

#### func (*XooaClient) ChangeBasePath

```go
func (c *XooaClient) ChangeBasePath(path string)
```
Change base path to allow switching to mocks

# models
--
    import "go-client/models"


## Usage

#### type Attrs

```go
type Attrs struct {
	Name  string `json:"name"`
	Ecert bool   `json:"ecert"`
	Value string `json:"value"`
}
```

Struct Identity Attribute Object

#### type Authenticate

```go
type Authenticate struct {
	Token string `json:"token"`
}
```


#### type Block

```go
type Block struct {
	PreviousBlockHash string `json:"previousBlockHash"`
	CurrentBlockHash  string `json:"currentBlockHash"`
	BlockNumber       int    `json:"blockNumber"`
}
```

Struct Block

#### type BlockTransactionCount

```go
type BlockTransactionCount struct {
	PreviousHash         string `json:"previous_hash"`
	DataHash             string `json:"data_hash"`
	NumberOfTransactions int    `json:"numberOfTransactions"`
	BlockNumber          int    `json:"blockNumber"`
}
```

Struct Block Transaction Count

#### type CallbackFn

```go
type CallbackFn func(Message)
```


#### type DeleteIdentityResponse

```go
type DeleteIdentityResponse struct {
	Deleted bool `json:"deleted"`
}
```

Struct Delete Identity Response

#### type EventClient

```go
type EventClient struct {
	Client     *gosocketio.Client
	EventRegex string
}
```


#### type Identity

```go
type Identity struct {
	CanManageIdentities bool      `json:"canManageIdentities"`
	CreatedAt           time.Time `json:"createdAt"`
	ApiToken            string    `json:"ApiToken"`
	IdentityName        string    `json:"IdentityName"`
	Access              string    `json:"Access"`
	Id                  string    `json:"Id"`
	Attrs               []Attrs   `json:"Attrs"`
}
```

Struct Identity

#### type InvokeResponse

```go
type InvokeResponse struct {
	TxId    string      `json:"txId"`
	Payload interface{} `json:"payload"`
}
```

Struct Invoke Response

#### type Message

```go
type Message struct {
	TxId      string `json:"txId"`
	Status    string `json:"status"`
	BlockNum  string `json:"blockNum"`
	EventName string `json:"eventName"`
	Payload   string `json:"payload"`
}
```

Event message Struct

#### type PendingTransactionResponse

```go
type PendingTransactionResponse struct {
	ResultURL string `json:"resultURL"`
	ResultId  string `json:"resultId"`
}
```

Struct Pending Transaction Response

#### type Query

```go
type Query struct {
	Payload interface{} `json:"payload"`
}
```

Struct Query Payload

#### type RegenerateIdentityResponse

```go
type RegenerateIdentityResponse struct {
	Success bool `json:"success"`
}
```

Struct Regenerate Identity Response

#### type Result

```go
type Result struct {
	RequestType string                 `json:"requestType"`
	Payload     map[string]interface{} `json:"payload"`
}
```

Struct Result for Pending Transaction

#### type XooaApiException

```go
type XooaApiException struct {
	ErrorCode    int    `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
}
```

Xooa Api Exception Handling

#### func  NewXooaApiException

```go
func NewXooaApiException(httpResponse *http.Response) *XooaApiException
```
Function to Create new Xooa Api Exception from Output Object

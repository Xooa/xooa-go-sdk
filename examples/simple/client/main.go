/*
An Example of Xooa Go SDK 

*/


package main

import (
	"go-client"
	"go-client/lib"
	"fmt"
	"context"
	"go-client/models"
)



func Init() {
	url := "https://api.xooa.com/api/v1"

	//TODO fill your token here
	token := "<YOUR_TOKEN_HERE>"
	// Set log level to info
	// [TraceLevel:"trace",DebugLevel:"debug",InfoLevel:"info",WarnLevel:"warning",ErrorLevel:"error",FatalLevel: "fatal" PanicLevel:"panic"]
	xooa_client.SetLogLevel("info")
	// Set URL for Xooa API server
	xooa_client.SetUrl(url)
	//Set Token for the App
	xooa_client.SetToken(token)
}


//Payload for Invoke
type Payload struct {
	Args []string `json:"args"`
}

func main() {
	Init()
	//AppFunctions()
	EventFunctions()

}

// AppFunctions contains all the Api Related function examples
// Output of the function is redirected to the standard output console

func AppFunctions() {
	/*creating new configuration object for xooa_client */
	cfg, err := xooa.NewConfiguration()
	if err != nil {
		fmt.Println(err);
	}
	// Creating new xooa client Object
	x := xooa_client.NewXooaClient(cfg)

	//########################################################################
	//########## BlockChain Module Api's #####################################
	//########################################################################
	//Get block hashes (block #1)
	blockData, _, _ := x.GetBlockByNumber(context.TODO(), "1", map[string]interface{}{})
	fmt.Println(blockData)
	//Get current block hashes
	blockDataN, _, _ := x.GetCurrentBlock(context.TODO(), map[string]interface{}{})
	fmt.Println(blockDataN)

	//########################################################################
	//########## Identity Module Api's #######################################
	//########################################################################

	//This endpoint returns authenticated identity
	identity, _ := x.AuthenticatedIdentityInformation(context.TODO())
	fmt.Println(identity)

	//Get all identities from the identity registry.
	//(Required permission: manage identities (canManageIdentities=true)).

	identityAll, _ := x.IdentitiesGetAllIdentities(context.TODO())
	fmt.Println(identityAll)
	//Get the specified identity from the identity registry
	identity, _ = x.IdentityInformation(context.TODO(), identity.Id)
	fmt.Println(identity)

	//The Enroll identity endpoint is used to enroll new identities for the smart contract app.
	//Required permission: manage identities (canManageIdentities=true).

	// New Attribute Object
	var aAttr []models.Attrs
	aAttr = append(aAttr, models.Attrs{Name:"Test Name", Ecert: true, Value: "Test Value"})
	//New Identity Object
	identityN := models.Identity{
		IdentityName: "Test Identity",
		Access: "r",
		Attrs: aAttr,
		CanManageIdentities: true,
	}
	identityEnrolled, _, _ := x.EnrollIdentity(context.TODO(), map[string]interface{}{}, identityN)
	fmt.Println(identityEnrolled)

	//Deletes an identity. Required permission: manage identities (canManageIdentities=true)
	identityD, _, _ := x.DeleteIdentity(context.TODO(), identityEnrolled.Id, map[string]interface{}{})
	fmt.Println(identityD)
	//Generates new identity API Token. Required permission: manage identities (canManageIdentities=true)
	identityEnrolled, _, _ = x.EnrollIdentity(context.TODO(), map[string]interface{}{}, identityN)
	identityR, _, _ := x.RegenerateToken(context.TODO(), identity.Id, map[string]interface{}{})
	fmt.Println(identityR)

	//########################################################################
	//########## Invoke Module Api's #########################################
	//########################################################################

	//The invoke API endpoint is used for submitting transaction for processing by the blockchain smart contract app
	//when the transaction payload need to be persisted into the Ledger (new block is mined).
	//Required permission: write ("Access":"rw")
	//"set" function is the part of the smartcontract provided in this example
	invokeResponse, _, _ := x.Invoke(context.TODO(), "set", map[string]interface{}{"timeout" : "10000"}, Payload{Args:[]string{"arg1", "arg2"}})
	fmt.Println(invokeResponse)
	//########################################################################
	//########## Query Module Api's #########################################
	//########################################################################

	//The query API endpoint is used for querying (reading) a blockchain ledger using smart contract function.
	//The endpoint must call a function already defined in your smart contract app which will process the query request.
	//Required permission: read ("Access":"rw" or "Access":"r")
	//"get" function is the part of the smartcontract provided in this example
	queryResponse, _, _ := x.Query(context.TODO(), "get", map[string]interface{}{"args": "arg1"})
	fmt.Println(queryResponse)
}



// EventFunctions contains the Event listening functionality
// Output of the function is redirected to the standard output console

func EventFunctions() {
	/*creating new configuration object for xooa_client */
	cfg, err := xooa.NewConfiguration()
	if err != nil {
		fmt.Println(err);
	}
	// Creating new xooa client Object
	x := xooa_client.NewXooaClient(cfg)

	//Activating Event server
	x.ConnectEventServer(callbackFun);

	//Subscribing to all events that will be generated by the smartcontract in the example.
	x.SubscribeAllEvents();

	//Subscribing to specific events this will listen to the specific event passed as a keyword in the function argument
	//x.SubscribeEvents("set");
}

// Callback function for event
// this example is working for the event thrown by smartcontract in the example
//stub.SetEvent("set",[]byte(args[1]))

func callbackFun(content models.Message) {
	fmt.Println("Event Recieved... ", content)

}
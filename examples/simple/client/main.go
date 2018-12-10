/**
 * An Example of Xooa Go SDK
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

package main

import (
	"github.com/Xooa/xooa-go-sdk"
	"github.com/Xooa/xooa-go-sdk/lib"
	"fmt"
	"context"
	"github.com/Xooa/xooa-go-sdk/models"
)

func Init() {
	url := "https://api.xooa.com/api/v1"

	//TODO fill your token here
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBcGlLZXkiOiJKSFAyR1BKLTdHVE03U0UtTllFN1lXOS1YVjZDUTkzIiwiQXBpU2VjcmV0IjoidDV4NnJEZDVNcHdrU3ZBIiwiUGFzc3BocmFzZSI6ImYyOWRmOGQ1M2Q1ODE2MWI3M2Q1ODEyYWVmN2JmMDgwIiwiaWF0IjoxNTQ0NDE1OTk1fQ.1QGfAIeXFz6q1N4KOPjTAmP7OcAcHD7r3r7zaUCphTs"
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
	AppFunctions()
	//EventFunctions()

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

	//Generates new identity API Token. Required permission: manage identities (canManageIdentities=true)
	identityR, _, _ := x.RegenerateToken(context.TODO(), identityEnrolled.Id, map[string]interface{}{})
	fmt.Println(identityR)

	//Deletes an identity. Required permission: manage identities (canManageIdentities=true)
	identityD, _, _ := x.DeleteIdentity(context.TODO(), identityEnrolled.Id, map[string]interface{}{})
	fmt.Println(identityD)

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

	//Subscribing to specific events this will listen to the specific event passed as a keyword in the function argument
	x.SubscribeAllEvents();
}

// Callback function for event
// this example is working for the event thrown by smartcontract in the example
//stub.SetEvent("set",[]byte(args[1]))

func callbackFun(content models.Message) {
	fmt.Println("Event Recieved... ", content)

}
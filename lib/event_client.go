/*
 * Xooa Blockchain API's
 *
 * List of Xooa Event Client API's
 *
 * API version: v1
 * Contact: support@xooa.com
 */

package xooa

import (
	"github.com/gsocket-io/golang-socketio"
	"github.com/gsocket-io/golang-socketio/transport"
	"sync"
	"time"
	"net/url"
	"regexp"
	log "github.com/sirupsen/logrus"
	"go-client/models"
)

type EventClientService service

//Checking for the last TxId to avoid Duplication
var EventPool string

// Reg exp to for all events
var RegExp  = "a*"

// Regex instance for handling Event Subscription
var RegExpI *regexp.Regexp

// Maintaining Socket io client to support connection close
var C *gosocketio.Client

// Auth using  `authenticate` event to server
func sendJoin(c *gosocketio.Channel) {
	c.Emit("authenticate", models.Authenticate{Token:Token})
}

/* Initiating connection to Socket io server
 @param callbackFn function to return the events
 @return error
*/
func (a *EventClientService)  ConnectEventServer(callbackFn models.CallbackFn) error {
	//Parsing URL for Hostname
	urlD,err := url.Parse(BasePath)
	if err != nil {
		log.Fatal(err);
	}

	//Dialing Socket io server
	path := urlD.Host
	c, err := gosocketio.Dial(
		gosocketio.GetUrl(path, 443, true),
		transport.GetDefaultWebsocketTransport())
	if err != nil {
		log.Fatal(err)
	}
	//copying socket instance to global variable
	C = c;

	//Standard `event` for capturing all the events in from `Smart Contract `
	err = c.On("event", func(h *gosocketio.Channel, args models.Message) {
		if checkForDuplicate(args.TxId) {
			EventPool = args.TxId;
			RegExpI = regexp.MustCompile(RegExp)
			if RegExpI.MatchString(args.EventName) {
				callbackFn(args)
			}
		}



	})
	if err != nil {
		log.Fatal(err)
	}

	// Handling Disconnection
	err = c.On(gosocketio.OnDisconnection, func(h *gosocketio.Channel) {
		log.Info("Refreshing Connection ... ")
		a.ConnectEventServer(callbackFn);
	})
	if err != nil {
		log.Fatal(err)
	}

	//Event on Connection
	err = c.On(gosocketio.OnConnection, func(h *gosocketio.Channel) {
		log.Info("Connected")
		go sendJoin(h)
	})
	if err != nil {
		log.Fatal(err)
	}

	defer c.Close()

	//Keeping the server alive
	keepAlive(c)
	return nil

}

//Putting a Regex `"a*"` to capture all the events
func(a *EventClientService) SubscribeAllEvents() error{
	RegExp = "a*"
	return nil
}


// TO capture the specific events specified in the call
func(a *EventClientService) SubscribeEvents(regEx string) error{
	if regEx == "" {
		log.Error("regEx Cannot be Empty")
		return nil
	}

	RegExp = regEx
	return nil
}

// UnSubscribing all the eventa
func(a *EventClientService) UnSubscribe() {
	C.Close()
}

// Keeping alive using WaitGroup
func keepAlive(c *gosocketio.Client) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for {
			time.Sleep(5 * time.Second)
			if c.IsAlive() != true {
				log.Info("Connection Not Alive")
				wg.Done()
				return
			}

		}

	}()
	wg.Wait()
}

// Checking for Duplicate Event
func checkForDuplicate(eventId string) bool{
	return EventPool != eventId
}
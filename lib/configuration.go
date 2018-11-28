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
	"net/http"
	"errors"
	log "github.com/sirupsen/logrus"
	"net/url"
)

var (
	Token string

)


//Base Path of Api's we have used
var BasePath string = "https://api.xooa.com/api/v1"

var LogLevel string = "error"

//Token for auth in Api
func SetToken(token string) error{
	if (token != "") {
		Token = token;
	} else {
		return  errors.New("token Cannot be Empty")
	}
	return nil
}


//Setting Log Level For SDK
func SetLogLevel(logLevel string) error{
	if logLevel != ""  {
		level, err := GetLogLevel(logLevel)
		if err != nil {
			log.Error(err)
		}
		log.SetLevel(level);
	} else {
		return  errors.New("logLevel Cannot be Empty")
	}
	return nil
}

// SetUrl setting Base path from Outside the SDK
func SetUrl(urlI string) {
	if (urlI != "") {
		_,err := url.Parse(urlI)
		if err != nil {
			log.Error(err)
			return
		}
		BasePath = urlI ;
	} else {
		log.Error("Must be a non empty string ")
	}
}

type contextKey string

func (c contextKey) String() string {
	return "Bearer " + string(c)
}

// Configuration for setting extra properties of the XooaClient Object
type Configuration struct {
	Host          string                `json:"host,omitempty"`
	Scheme        string                `json:"scheme,omitempty"`
	DefaultHeader map[string]string     `json:"defaultHeader,omitempty"`
	UserAgent     string                `json:"userAgent,omitempty"`
	HTTPClient    *http.Client
}


// Creating new Configuration Object
func NewConfiguration() (*Configuration, error) {
	cfg := &Configuration{
		DefaultHeader: make(map[string]string),
	}

	return cfg, nil
}



// Adding default Headers
func (c *Configuration) AddDefaultHeader(key string, value string) {
	c.DefaultHeader[key] = value
}

// SettingUserAgent Setting User agent
func (c *Configuration) SettingUserAgent(key string, value string) {
	c.DefaultHeader[key] = value
}
package models

import (
	"net/http"
	"io/ioutil"
)


// Xooa Api Exception Handling
type XooaApiException struct {
	ErrorCode  int `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
}


// Function to Create new Xooa Api Exception from Output Object
func NewXooaApiException(httpResponse *http.Response) *XooaApiException {
	bodyBytes, _ := ioutil.ReadAll(httpResponse.Body)
	return &XooaApiException{
		ErrorCode: httpResponse.StatusCode,
		ErrorMessage: string(bodyBytes),
	}
}
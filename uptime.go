package main

import (
	"io/ioutil"
	"net/http"
	"time"
)

type UptimeClient struct {
	// client http.Client
}

func NewUptimeClient() *UptimeClient {
	return &UptimeClient{
		//client: http.Client{},
	}
}

func (uptime *UptimeClient) Check(checkConfig UptimeCheckConfig) *Event {
	var err error
	var response *http.Response
	responseString := ""

	httpClient := http.Client{}
	request := http.Request{
		Method: checkConfig.method,
		URL:    checkConfig.url,
		Header: checkConfig.headers,
	}
	// if the client should not follow redirects
	if !checkConfig.followRedirects {
		httpClient.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}
	}
	// Define the event because it records the scan start
	event := NewEvent(err, UptimeCheck, request.RemoteAddr)
	started := time.Now().UTC().UnixNano()
	// make the real request
	response, err = httpClient.Do(&request)

	// close the response unless there was an error
	if err == nil {
		defer response.Body.Close()
		responseData, err := ioutil.ReadAll(response.Body)
		if err == nil {
			responseString = string(responseData)
		}

		// simple check of the response for now...
		if response.StatusCode >= 400 {
			// this is an error
		}

		// if the check against the body is set
		if checkConfig.expectedBody != "" {

			if checkConfig.expectedBody != responseString {
				// this is considered an error
			}
		}

	}

	event.ScanEnd = time.Now().UTC()
	ended := event.ScanEnd.UnixNano()
	difference := ended - started
	event.ResponseBody = responseString
	event.ResponseCode = response.StatusCode
	event.ResponseLatencyMS = difference
	return event
}

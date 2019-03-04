package main

// This contains the various types of checks and their configuration

import (
	"log"
	"net/http"
	"net/url"
	"time"
)

// CheckType the type of check
type CheckType int

const (
	// TLSCheck checks for certificate expiration
	TLSCheck CheckType = iota + 1

	// PortScanCheck checks for open ports
	PortScanCheck

	// UptimeCheck checks whether the service is up
	UptimeCheck
)

type UptimeCheckConfig struct {
	method              string      // GET, HEAD, POST etc...
	headers             http.Header // additional headers
	url                 *url.URL    // url to check against
	expectedStatusCodes []int       // the expected http status code
	expectedBody        string      // the expected http body
	followRedirects     bool        // whether the client should follow redirects
}

func NewUptimeCheckConfig(method, URL, expectedBody string, headers http.Header) UptimeCheckConfig {

	parsedURL, err := url.Parse(URL)
	if err != nil {
		log.Println("The url is invalid", URL, err)
	}
	// TODO: notify the user that the URL is not valid!

	return UptimeCheckConfig{
		url:             parsedURL,
		method:          method,
		headers:         headers,
		expectedBody:    expectedBody,
		followRedirects: true,
	}
}

// Check defines the type of check
type Check struct {
	Frequency   int       `json:"frequency"`
	LastChecked time.Time `json:"last_checked"`
	Type        CheckType `json:"check_type"`
}

func NewCheck(frequencyMS int, checkType CheckType) Check {
	return Check{
		Frequency: frequencyMS,
		Type:      checkType,
	}
}

// Domain defines the list of domains to be checked
type Domain struct {
	Name   string  `json:"name"`
	Checks []Check `json:"checks"`
	Hosts  []Host  `json:"hosts"`
}

// Host defines which hosts in the domain to be checked
type Host struct {
	Name     string              `json:"name"`
	Checks   map[CheckType]Check `json:"checks"`
	Disabled bool                `json:"disabled"`
	Events   Events              `json:"events"`
}

func (h *Host) ShouldScan(scanType CheckType) bool {
	_, ok := h.Checks[scanType]
	return ok
}

func NewHost(host string) *Host {
	return &Host{
		Name:   host,
		Checks: make(map[CheckType]Check),
		Events: []Event{},
	}
}

func (h *Host) SetCheck(check Check) {
	h.Checks[check.Type] = check
}

func (h *Host) RemoveCheck(check Check) {
	delete(h.Checks, check.Type)
}

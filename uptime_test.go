package main

import (
	"log"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUptimeEvent(t *testing.T) {
	checkConfig := NewUptimeCheckConfig("GET", "https://www.livescore.com/statusutc/", "", http.Header{})
	uptimeCheck := NewUptimeClient()
	event := uptimeCheck.Check(checkConfig)
	log.Printf("%+v\n", event)
	assert.NotNil(t, event)
	assert.Equal(t, UptimeCheck, event.Type)
	assert.Equal(t, 200, event.ResponseCode)
	assert.NotEmpty(t, 200, event.ResponseBody)
	assert.True(t, event.ResponseLatencyMS > 10)

}

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDNSClient(t *testing.T) {
	client := NewDNSClient()
	assert.NotNil(t, client)

	client.Hosts("livescore.com")
}

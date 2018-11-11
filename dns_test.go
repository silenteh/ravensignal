package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDNSClient(t *testing.T) {
	client := NewDNSClient()
	assert.NotNil(t, client)

	hosts, err := client.IPV4Hosts("livescore.com")
	assert.Nil(t, err)
	assert.NotEmpty(t, hosts)

	hosts, err = client.IPV6Hosts("livescore.com")
	assert.Nil(t, err)
	assert.NotEmpty(t, hosts)
}

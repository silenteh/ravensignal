package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccountJson(t *testing.T) {
	account := NewAccount("1234", "sec", kSilverPlan)

	assert.NotNil(t, account)
}

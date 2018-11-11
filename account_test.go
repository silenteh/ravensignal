package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccountJson(t *testing.T) {
	// Create an account
	account := NewAccount("sec", kSilverPlan)
	account.ID = "1234"
	assert.NotNil(t, account)

	// check the account fields
	assert.Equal(t, "1234", account.ID)
	assert.Equal(t, "sec", account.CompanyName)
	assert.Equal(t, kSilverPlan, account.Plan)
	assert.Empty(t, account.Users)
	assert.Empty(t, account.Hosts)

	user := NewUser("1", "security@sec51.com", true)
	assert.NotNil(t, user)
	assert.Equal(t, "1", user.ID)
	assert.Equal(t, "security@sec51.com", user.Email)
	assert.True(t, user.IsAdmin)

	account.SetUser(*user)
	assert.NotEmpty(t, account.Users)

	// add the host again and check that it gets replaced
	user.IsAdmin = false
	account.SetUser(*user)
	assert.Equal(t, 1, len(account.Users))
	assert.False(t, account.Users[0].IsAdmin)

	host := NewHost("livescore.com")
	assert.NotNil(t, host)

	check := NewCheck(30000, PortScanCheck)
	host.SetCheck(check)

	assert.NotEmpty(t, host.Checks)

	account.SetHost(*host)
	assert.NotEmpty(t, account.Hosts)

	// add the host again and check that it gets replaced
	account.SetHost(*host)
	assert.Equal(t, 1, len(account.Hosts))

	account.Save()

	// Now try to load them all

	accounts := LoadAllAccounts()
	assert.NotEmpty(t, accounts)
	assert.Equal(t, account.ID, accounts[0].ID)

}

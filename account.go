package main

// Account of the company
type Account struct {
	ID          string
	CompanyName string
	Plan        Plan
	Domains     []Domain
	Users       []User
}

// User belonging to the account
type User struct {
	ID      string
	Email   string
	IsAdmin bool
}

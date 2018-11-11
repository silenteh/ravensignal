package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// Account of the company
type Account struct {
	ID          string `json:"id"`
	CompanyName string `json:"company"`
	Plan        Plan   `json:"plan"`
	Hosts       []Host `json:"hosts"`
	Users       []User `json:"users"`
}

// User belonging to the account
type User struct {
	ID      string `json:"id"`
	Email   string `json:"email"`
	IsAdmin bool   `json:"is_admin"`
}

func NewUser(id, email string, isAdmin bool) *User {
	return &User{
		ID:      id,
		Email:   email,
		IsAdmin: isAdmin,
	}
}

func NewAccount(id, company string, plan Plan) *Account {
	return &Account{
		ID:          id,
		CompanyName: company,
		Plan:        plan,
		Hosts:       []Host{},
		Users:       []User{},
	}
}

func (a *Account) SetUser(user User) {
	for index, existingUser := range a.Users {
		if existingUser.ID == user.ID {
			a.Users[index] = user
			return
		}
	}
	a.Users = append(a.Users, user)
}

func (a *Account) SetDomain(host Host) {
	for index, existingHost := range a.Hosts {
		if existingHost.Name == host.Name {
			a.Hosts[index] = host
			return
		}
	}
	a.Hosts = append(a.Hosts, host)
}

func (a *Account) ToJson() []byte {
	var data []byte
	var err error
	data, err = json.Marshal(a)
	if err != nil {
		log.Println("Failed to serialize Account JSON")
	}
	return data
}

func AccountFromJSON(id string) *Account {
	file, err := os.Open(fmt.Sprintf("%s.json", id))
	if err != nil {
		log.Println("Failed to read Account JSON file")
		file.Close()
		return nil
	}

	var account Account
	fileBuffer := bufio.NewReader(file)
	err = json.NewDecoder(fileBuffer).Decode(&account)
	if err != nil {
		log.Println("Failed to parse Account JSON file")
		file.Close()
		return nil
	}
	file.Close()
	return &account

}

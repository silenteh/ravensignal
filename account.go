package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

// Account of the company
type Account struct {
	ID          string `json:"id"`
	CompanyName string `json:"company"`
	Plan        Plan   `json:"plan"`
	Hosts       []Host `json:"hosts"`
	Users       []User `json:"users"`
}

// Go - starts to execute all the checks
func (a *Account) Go(srcIP net.IP, srcInterface net.Interface, srcIpAddrs ipAddrs) {
	for _, host := range a.Hosts {
		if host.Disabled {
			log.Println(host.Name, "is disabled, skipping checks.")
			continue
		}
		for _, check := range host.Checks {
			switch check.Type {
			case TLSCheck:
				go ScanTLS(&host, check.Frequency)
				break
			case UptimeCheck:
				break
			case PortScanCheck:
				go ScanPorts(srcIP, srcInterface, srcIpAddrs, &host, check.Frequency)
				break
			}
		}
	}
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

func NewAccount(company string, plan Plan) *Account {
	id, err := uuid.NewRandom()
	if err != nil {
		log.Println("Failed to create an Account ID", err)
		return nil
	}

	return &Account{
		ID:          id.String(),
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

func (a *Account) SetHost(host Host) {
	for index, existingHost := range a.Hosts {
		if existingHost.Name == host.Name {
			a.Hosts[index] = host
			return
		}
	}
	a.Hosts = append(a.Hosts, host)
}

// RemoveHost removes one host
func (a *Account) RemoveHost(name string) {
	for index, existingHost := range a.Hosts {
		if existingHost.Name == name {
			a.Hosts = append(a.Hosts[:index], a.Hosts[index+1:]...)
		}
	}
}

// RemoveUser removes one user
func (a *Account) RemoveUser(email string) {
	for index, existingUser := range a.Users {
		if existingUser.Email == email {
			a.Users = append(a.Users[:index], a.Users[index+1:]...)
		}
	}
}

func (a *Account) Save() error {
	var data []byte
	var err error
	data, err = json.Marshal(a)
	if err != nil {
		log.Println("Failed to serialize Account JSON")
		return err
	}
	if len(data) > 0 {
		err = ioutil.WriteFile(filepath.Join(kBaseDataFolder, fmt.Sprintf("%s.json", a.ID)), data, 0740)
	}

	return err
}

func AccountFromJSON(fileName string) *Account {
	file, err := os.Open(fileName)
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

func LoadAllAccounts() []Account {
	var accounts []Account
	files, err := filePathWalkDir(kBaseDataFolder)
	if err != nil {
		return accounts
	}

	for _, fileName := range files {
		account := AccountFromJSON(fileName)
		if account != nil {
			accounts = append(accounts, *account)
		}
	}
	return accounts
}

func filePathWalkDir(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

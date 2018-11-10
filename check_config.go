package main

import "time"

// CheckType the type of check
type CheckType int

const (
	// TLSCheck checks for certificate expiration
	TLSCheck = iota + 1

	// PortScanCheck checks for open ports
	PortScanCheck
)

// Check defines the type of check
type Check struct {
	Frequency   int
	LastChecked time.Time
	Type        CheckType
}

// Domain defines the list of domains to be checked
type Domain struct {
	Name  string
	Check Check
	Hosts []Host
}

// Host defines which hosts in the domain to be checked
type Host struct {
	Name        string
	Check       Check
	LastChecked time.Time
}

package main

import "time"

// CheckType the type of check
type CheckType int

const (
	// TLSCheck checks for certificate expiration
	TLSCheck = iota + 1

	// PortScanCheck checks for open ports
	PortScanCheck

	// UptimeCheck checks whether the service is up
	UptimeCheck
)

// Check defines the type of check
type Check struct {
	ID          string    `json:"id"`
	Frequency   int       `json:"frequency"`
	LastChecked time.Time `json:"last_checked"`
	Type        CheckType `json:"type"`
}

// Domain defines the list of domains to be checked
type Domain struct {
	Name  string `json:"name"`
	Check Check  `json:"check"`
	Hosts []Host `json:"hosts"`
}

// Host defines which hosts in the domain to be checked
type Host struct {
	Name   string  `json:"name"`
	Checks []Check `json:"checks"`
}

func NewHost(host string) *Host {
	return &Host{
		Name:   host,
		Checks: []Check{},
	}
}

func (h *Host) SetCheck(check Check) {
	for _, existingCheck := range h.Checks {
		if existingCheck.ID == check.ID {
			return
		}
	}
	h.Checks = append(h.Checks, check)
}

func (h *Host) RemoveCheck(check Check) {
	for index, existingCheck := range h.Checks {
		if existingCheck.ID == check.ID {

			h.Checks = append(h.Checks[:index], h.Checks[index+1:]...)
		}
	}
}

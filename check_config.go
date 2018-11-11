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
	Frequency   int       `json:"frequency"`
	LastChecked time.Time `json:"last_checked"`
	Type        CheckType `json:"check_type"`
}

func NewCheck(frequencyMS int, checkType CheckType) Check {
	return Check{
		Frequency: frequencyMS,
		Type:      checkType,
	}
}

// Domain defines the list of domains to be checked
type Domain struct {
	Name   string  `json:"name"`
	Checks []Check `json:"checks"`
	Hosts  []Host  `json:"hosts"`
}

// Host defines which hosts in the domain to be checked
type Host struct {
	Name     string              `json:"name"`
	Checks   map[CheckType]Check `json:"checks"`
	Disabled bool                `json:"disabled"`
	Events   Events              `json:"events"`
}

func (h *Host) ShouldScan(scanType CheckType) bool {
	_, ok := h.Checks[scanType]
	return ok
}

func NewHost(host string) *Host {
	return &Host{
		Name:   host,
		Checks: make(map[CheckType]Check),
		Events: []Event{},
	}
}

func (h *Host) SetCheck(check Check) {
	h.Checks[check.Type] = check
}

func (h *Host) RemoveCheck(check Check) {
	delete(h.Checks, check.Type)
}

package main

import (
	"log"
	"time"

	"github.com/google/uuid"
)

// UptimeStatus defines whether a service is up
type UptimeStatus int
type Events []Event

const (
	// StatusUP the service is up
	StatusUP = iota + 1

	// StatusDown the service is down
	StatusDown
)

type ExpiringCert struct {
	CA            string    `json:"ca"`
	Expires       time.Time `json:"expires"`
	ExpiresInDays string    `json:"expires_in_days"`
}

func (es *Events) AddEvent(event *Event) {
	newEvents := append(*es, *event)
	es = &newEvents
}

// Event Base
type Event struct {
	ID        string    `json:"id"`
	Error     error     `json:"error"`
	Type      CheckType `json:"check_type"`
	IPAddress string    `json:"ip_address"`
	ScanStart time.Time `json:"scan_start"`
	ScanEnd   time.Time `json:"scan_end"`
}

func NewEvent(err error, checkType CheckType, IPAddress string) *Event {
	id, err := uuid.NewRandom()
	if err != nil {
		log.Println("Failed to create an Account ID", err)
		return nil
	}
	return &Event{
		ID:        id.String(),
		Error:     err,
		Type:      checkType,
		ScanStart: time.Now().UTC(),
	}
}

func (e *Event) ScanEnded() {
	e.ScanEnd = time.Now().UTC()
}

// PortScanEvent results of the port scan
type PortScanEvent struct {
	Event
	OpenPorts []int `json:"open_ports"`
}

// TLSScanEvent results of the TLS scan
type TLSScanEvent struct {
	ExpiringCerts []ExpiringCert `json:"expiring_certs"`
	TLSGrade      int            `json:"tls_grade"` // TODO
	Event
}

func NewTLSScanEvent(err error, checkType CheckType, IPAddress, ca, in string, expires time.Time) *TLSScanEvent {

	event := NewEvent(err, checkType, IPAddress)

	tlsEvent := &TLSScanEvent{
		ExpiringCerts: []ExpiringCert{ExpiringCert{
			CA:            ca,
			Expires:       expires,
			ExpiresInDays: in,
		}},
	}
	tlsEvent.Event = *event

	return tlsEvent
}

// UptimeScanEvent result of the uptime scan
type UptimeScanEvent struct {
	Event
	Status    UptimeStatus `json:"uptime_status"`
	LatencyMS int          `json:"uptime_latency_ms"`
}

// TODO: notification event

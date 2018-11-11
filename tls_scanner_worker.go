package main

import (
	"log"
	"time"
)

func ScanTLS(host *Host, frequency int) {

	for range time.Tick(time.Millisecond * time.Duration(frequency)) {
		// 1. Resolve the IP address
		dnsClient := NewDNSClient()
		ipV4Addresses, err := dnsClient.IPV4Hosts(host.Name)
		if err != nil {
			host.Events.AddEvent(NewEvent(err, TLSCheck, host.Name))
			continue
		}
		// dnsClient.Hosts
		// 2. Do the TLS scan
		for _, ip := range ipV4Addresses {
			tlsEvent := CheckCert(ip, host.Name, false)
			log.Printf("tls event: %+v\n", *tlsEvent)
		}

		ipV6Addresses, err := dnsClient.IPV6Hosts(host.Name)
		if err != nil {
			host.Events.AddEvent(NewEvent(err, TLSCheck, host.Name))
			continue
		}
		// dnsClient.Hosts
		// 2. Do the TLS scan
		for _, ip := range ipV6Addresses {
			tlsEvent := CheckCert(ip, host.Name, true)
			log.Printf("tls event: %+v\n", *tlsEvent)
		}

	}

	// 3. Add the results to the host pointer

}

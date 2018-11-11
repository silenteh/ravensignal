package main

import (
	"log"
	"net"
	"time"
)

func ScanPorts(srcIP net.IP, srcInterface net.Interface, srcIpAddrs ipAddrs, host *Host, frequency int) {

	var dstIP net.IP

	router := ManualRouter{
		gw:            kGatewayIP,
		interfaceName: srcInterface,
		srcIP:         srcIpAddrs,
	}

	for range time.Tick(time.Millisecond * time.Duration(frequency)) {
		// 1. Resolve the IP address
		dnsClient := NewDNSClient()
		ipV4Addresses, err := dnsClient.IPV4Hosts(host.Name)
		if err != nil {
			host.Events.AddEvent(NewEvent(err, TLSCheck, host.Name))
			continue
		}
		// dnsClient.Hosts
		// 2. Do the Port scan
		for _, ip := range ipV4Addresses {
			// parse the destination IP
			if dstIP = net.ParseIP(ip); dstIP == nil {
				log.Printf("Target interface parsing issue: %q", ip)
				continue
			}
			log.Printf("Parsed destination IP: %s\n", dstIP.String())

			s, err := newScanner(srcIP, dstIP, router)
			if err != nil {
				log.Printf("unable to create scanner for %v: %v", dstIP, err)
				continue
			}
			if err := s.scan(); err != nil {
				log.Printf("unable to scan %v: %v", dstIP, err)
			}
			s.close()
		}

		// ipV6Addresses, err := dnsClient.IPV6Hosts(host.Name)
		// if err != nil {
		// 	host.Events.AddEvent(NewEvent(err, TLSCheck, host.Name))
		// 	continue
		// }
		// dnsClient.Hosts
		// 2. Do the TLS scan
		// for _, ip := range ipV6Addresses {
		// 	tlsEvent := CheckCert(ip, host.Name, true)
		// 	log.Printf("tls event: %+v\n", *tlsEvent)
		// }

	}

	// 3. Add the results to the host pointer

}

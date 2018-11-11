package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

func CheckCert(host Host) {
	resp, err := http.Head(host.Name)
	if err != nil {
		log.Printf("Unable to get %q: %s\n", host.Name, err)
		return
	}
	resp.Body.Close()
	if resp.TLS == nil {
		log.Printf("%q is not HTTPS\n", host.Name)
		return
	}

	for _, cert := range resp.TLS.PeerCertificates {
		for _, name := range cert.DNSNames {
			if !strings.Contains(host.Name, name) {
				continue
			}
			issuer := strings.Join(cert.Issuer.Organization, ", ")
			dur := cert.NotAfter.Sub(time.Now())
			fmt.Printf("Certificate for %q from %q expires %s (%.0f days).\n", name, issuer, cert.NotAfter, dur.Hours()/24)
		}
	}
}

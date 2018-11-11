package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

var tr = &http.Transport{
	MaxIdleConns:       10,
	IdleConnTimeout:    30 * time.Second,
	DisableCompression: true,
	TLSClientConfig: &tls.Config{
		InsecureSkipVerify: true,
	},
}

var client = &http.Client{Transport: tr}

func CheckCert(ip, hostname string, isIPV6 bool) *TLSScanEvent {
	url := fmt.Sprintf("https://%s", ip)
	if isIPV6 {
		url = fmt.Sprintf("https://[%s]", ip)
	}
	req, err := http.NewRequest("HEAD", url, nil)

	req.Header.Set("Host", hostname)
	req.Host = hostname
	resp, err := client.Do(req)
	if err != nil {
		err = fmt.Errorf("Unable to get %s for hostname %s: %s", ip, hostname, err)
		log.Println(err)
		return NewTLSScanEvent(err, TLSCheck, ip, "", "", time.Now().UTC())
	}
	resp.Body.Close()
	if resp.TLS == nil {
		err = fmt.Errorf("%s behind the IP %s is not HTTPS", ip, hostname)
		log.Println(err)
		return NewTLSScanEvent(err, TLSCheck, ip, "", "", time.Now().UTC())
	}

	for _, cert := range resp.TLS.PeerCertificates {
		for _, name := range cert.DNSNames {
			if !strings.Contains(hostname, name) {
				continue
			}
			issuer := strings.Join(cert.Issuer.Organization, ", ")
			dur := cert.NotAfter.Sub(time.Now())

			fmt.Printf("Certificate for %q from %q expires %s (%.0f days).\n", name, issuer, cert.NotAfter, dur.Hours()/24)
			return NewTLSScanEvent(err, TLSCheck, ip, issuer, fmt.Sprintf("%.0f days", dur.Hours()/24), cert.NotAfter)
		}
	}
	err = fmt.Errorf("The server at %s does not have a valid certificate for the hostname %s", ip, hostname)
	return NewTLSScanEvent(err, TLSCheck, ip, "", "", time.Now().UTC())
}

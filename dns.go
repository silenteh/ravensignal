package main

import (
	"log"
	"net"
	"os"

	"github.com/miekg/dns"
)

type DNSClient struct {
	config *dns.ClientConfig
	client *dns.Client
}

func NewDNSClient() *DNSClient {
	config, _ := dns.ClientConfigFromFile("/etc/resolv.conf")
	c := new(dns.Client)
	c.Net = "udp"

	return &DNSClient{
		config: config,
		client: c,
	}
}

func (dnsClient *DNSClient) Hosts(domain string) {
	msg := new(dns.Msg)
	msg.Id = dns.Id()
	// msg.SetQuestion(dns.Fqdn(domain), dns.TypeA)
	msg.RecursionDesired = true
	msg.RecursionAvailable = true
	msg.Question = make([]dns.Question, 1)
	msg.Question[0] = dns.Question{
		Name:   dns.Fqdn(domain),
		Qtype:  dns.TypeSOA,
		Qclass: dns.ClassINET,
	}

	r, _, err := dnsClient.client.Exchange(msg, net.JoinHostPort(dnsClient.config.Servers[0], dnsClient.config.Port))
	if r == nil {
		log.Fatalf("*** error: %s\n", err.Error())
	}

	if r.Rcode != dns.RcodeSuccess {
		log.Fatalf(" *** invalid answer name %s after MX query for %s\n", os.Args[1], os.Args[1])
	}
	// Stuff must be in the answer section
	// for _, a := range r.Answer {
	// 	log.Printf("%v\n", a.)
	// }
	var As []*dns.A
	for _, v := range r.Answer {
		if a, ok := v.(*dns.A); ok {
			As = append(As, a)
		}
	}
	log.Printf("%v\n", As)
}

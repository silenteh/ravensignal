package main

import (
	"fmt"
	"log"
	"net"

	"github.com/miekg/dns"
)

type DNSClient struct {
	config *dns.ClientConfig
	client *dns.Client
}

func NewDNSClient() *DNSClient {
	config, _ := dns.ClientConfigFromFile("/etc/resolv.conf")
	c := new(dns.Client)
	c.Net = "tcp"

	return &DNSClient{
		config: config,
		client: c,
	}
}

func (dnsClient *DNSClient) IPV4Hosts(domain string) ([]string, error) {

	var ipAddresses []string

	msg := new(dns.Msg)
	msg.Id = dns.Id()
	msg.RecursionDesired = true
	msg.RecursionAvailable = true
	msg.Question = make([]dns.Question, 1)
	msg.Question[0] = dns.Question{
		Name:   dns.Fqdn(domain),
		Qtype:  dns.TypeA,
		Qclass: dns.ClassINET,
	}

	r, _, err := dnsClient.client.Exchange(msg, net.JoinHostPort(dnsClient.config.Servers[0], dnsClient.config.Port))
	if r == nil {
		log.Printf("*** error: %s\n", err.Error())
		return ipAddresses, err
	}

	if r.Rcode != dns.RcodeSuccess {
		err := fmt.Errorf("Invalid DNS answer for hostname: %s", domain)
		log.Println(err)
		// return ipAddresses, err
	}

	for _, v := range r.Answer {
		if a, ok := v.(*dns.A); ok {
			ipAddresses = append(ipAddresses, a.A.String())
		}

		if a, ok := v.(*dns.AAAA); ok {
			ipAddresses = append(ipAddresses, a.AAAA.String())
		}
	}

	return ipAddresses, nil
}

func (dnsClient *DNSClient) IPV6Hosts(domain string) ([]string, error) {

	var ipAddresses []string

	msg := new(dns.Msg)
	msg.Id = dns.Id()
	// msg.SetQuestion(dns.Fqdn(domain), dns.TypeA)
	msg.RecursionDesired = true
	msg.RecursionAvailable = true
	msg.Question = make([]dns.Question, 1)
	msg.Question[0] = dns.Question{
		Name:   dns.Fqdn(domain),
		Qtype:  dns.TypeAAAA,
		Qclass: dns.ClassINET,
	}

	r, _, err := dnsClient.client.Exchange(msg, net.JoinHostPort(dnsClient.config.Servers[0], dnsClient.config.Port))
	if r == nil {
		log.Printf("*** error: %s\n", err.Error())
		return ipAddresses, err
	}

	if r.Rcode != dns.RcodeSuccess {
		err := fmt.Errorf("Invalid DNS answer for hostname: %s", domain)
		log.Println(err)
		// return ipAddresses, err
	}

	for _, v := range r.Answer {
		if a, ok := v.(*dns.A); ok {
			ipAddresses = append(ipAddresses, a.A.String())
		}

		if a, ok := v.(*dns.AAAA); ok {
			ipAddresses = append(ipAddresses, a.AAAA.String())
		}
	}

	return ipAddresses, nil
}

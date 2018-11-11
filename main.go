package main

import (
	"log"
	"net"
	"strings"
)

func main() {

	// try to setup the router
	interfaces, err := net.Interfaces()
	if err != nil {
		log.Fatalf("Error getting interfaces: %s", err)
	}

	var parsedInterface net.Interface
	for _, intf := range interfaces {
		log.Println("Checking interface:", intf.Name)
		if intf.Name == kInterfaceName {
			parsedInterface = intf
			break
		}
	}

	if parsedInterface.Name == "" {
		log.Printf("Could not find interface: %s\n", kInterfaceName)
		return
	}

	var srcIP net.IP
	addresses, err := parsedInterface.Addrs()
	if err != nil {
		log.Panicln(err)
	}

	ipv4Address := ""
	for _, addr := range addresses {
		if strings.HasPrefix(addr.String(), kInterfaceIP) {
			ipv4Address = addr.String()
		}
	}

	var srcIpAddrs ipAddrs
	srcIP, srcIPNet, _ := net.ParseCIDR(ipv4Address)

	log.Printf("Detected addresses %s", srcIP)
	log.Printf("IPV4: %s", srcIPNet.IP.To4())

	if v4 := srcIPNet.IP.To4(); v4 != nil {
		if srcIpAddrs.v4 == nil {
			srcIpAddrs.v4 = srcIP
		}
	} else if srcIpAddrs.v6 == nil {
		srcIpAddrs.v6 = srcIPNet.IP
	}

	log.Printf("SRC IP ADDRESS: %s", srcIpAddrs.v4)

	if srcIP == nil {
		log.Printf("Interface address parsing issue: %q", ipv4Address)
		return
	}

	accounts := LoadAllAccounts()
	for _, account := range accounts {
		account.Go(srcIP, parsedInterface, srcIpAddrs)
	}

	api := NewApi("127.0.0.1", "8080")
	api.Start()

}

package main

import (
	"log"
	"net"

	"github.com/google/gopacket/routing"
)

var defaultRoute = [4]byte{0, 0, 0, 0}

type BsdRouter struct {
	routing.Router
}

type ManualRouter struct {
	routing.Router
	interfaceName net.Interface
	gw            string
	srcIP         ipAddrs
}

// (iface *net.Interface, gateway, preferredSrc net.IP
func (m ManualRouter) Route(dst net.IP) (*net.Interface, net.IP, net.IP, error) {
	gwIP := net.ParseIP(m.gw)
	log.Println("SRC ip address", m.srcIP.v4)
	return &m.interfaceName, gwIP, m.srcIP.v4, nil
	// // try to setup the router
	// interfaces, err := net.Interfaces()
	// if err != nil {
	// 	log.Fatalf("Error getting interfaces: %s", err)
	// }

	// log.Printf("Detected interfaces: %+s", interfaces)

	// var parsedInterface net.Interface
	// for _, intf := range interfaces {
	// 	log.Println("Checking interface:", intf.Name)
	// 	if intf.Name == m.interfaceName {
	// 		parsedInterface = intf
	// 		break
	// 	}
	// }
	// if parsedInterface.Name == "" {
	// 	log.Printf("Could not find interface: %s\n", m.interfaceName)
	// }
	// var gw net.IP

	// // var dstIP net.IP
	// // addresses, err := parsedInterface.Addrs()
	// // if err != nil {
	// // 	log.Panicln(err)
	// // }

	// if gw = net.ParseIP(m.gw); gw == nil {
	// 	log.Printf("could not parse gw", m.gw)
	// }
	// // if dstIP = net.ParseIP(dst); dstIP == nil {
	// // 	log.Printf("non-ipv4 target: %q", dst)
	// // }
	// return &parsedInterface, gw, m.srcIP, nil
}

// func (br BsdRouter) Route(dst net.IP) (*net.Interface, net.IP, net.IP, error) {
// 	rib, _ := route.FetchRIB(0, route.RIBTypeRoute, 0)
// 	messages, err := route.ParseRIB(route.RIBTypeRoute, rib)

// 	if err != nil {
// 		return nil, nil, nil, fmt.Errorf("Syscall err: %s", err)
// 	}

// 	var netInt net.Interface
// 	var gw net.IP
// 	var preferredSrc net.IP

// 	for _, message := range messages {
// 		routeMessage := message.(*route.RouteMessage)
// 		addresses := routeMessage.Addrs

// 		var destination, gateway *route.Inet4Addr
// 		ok := false

// 		if destination, ok = addresses[0].(*route.Inet4Addr); !ok {
// 			continue
// 		}

// 		if gateway, ok = addresses[1].(*route.Inet4Addr); !ok {
// 			continue
// 		}

// 		if destination == nil || gateway == nil {
// 			continue
// 		}

// 		if destination.IP == defaultRoute {
// 			fmt.Println(gateway.IP)
// 		}
// 	}

// 	return &netInt, gw, preferredSrc, nil
// }

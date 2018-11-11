package main

import "github.com/sec51/goconf"

var (
	kInterfaceName = goconf.AppConf.DefaultString("interface.name", "lo")
	kInterfaceIP   = goconf.AppConf.DefaultString("interface.ip", "127.0.0.1")
	kGatewayIP     = goconf.AppConf.DefaultString("gw.ip", "127.0.0.1")
	kGatewayMAC    = goconf.AppConf.DefaultString("gw.mac", "127.0.0.1")
)

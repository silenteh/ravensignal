package main

import "github.com/sec51/goconf"

var (
	kBaseDataFolder = goconf.AppConf.DefaultString("data.folder", "data")
	kInterfaceName  = goconf.AppConf.DefaultString("interface.name", "en5")
	kInterfaceIP    = goconf.AppConf.DefaultString("interface.ip", "10.192.180.241")
	kGatewayIP      = goconf.AppConf.DefaultString("gw.ip", "10.192.180.254")
	kGatewayMAC     = goconf.AppConf.DefaultString("gw.mac", "f4:03:43:03:1d:dc")
)

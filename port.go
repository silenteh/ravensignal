package main

import (
	"fmt"
	"net"
)

// GetFreePort asks the kernel for a free open port that is ready to use.
func GetFreePort(address string) (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s", address))
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	l.Close()
	return l.Addr().(*net.TCPAddr).Port, nil
}

package network

import (
	"errors"
	"net"
)

type IPType string

const (
	IPv4 IPType = "IPv4"
	IPv6 IPType = "IPv6"
)

func getIPType(ipStr string) (IPType, error) {

	ip := net.ParseIP(ipStr)
	if ip == nil {
		return "", errors.New("Invalid IP address!")
	}

	if ip.To4() != nil {
		return IPv4, nil
	}
	return IPv6, nil
}

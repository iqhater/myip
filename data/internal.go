package data

import (
	"errors"
	"net"
)

// InternalData struct is a local adapter and local ip data.
type InternalData struct {
	AdapterName   string
	IntIP         string
	IPAddressType IPType
}

// GetInternalIP get internal local ip addess.
func (i *InternalData) GetInternalIP() {

	conn, err := net.Dial("udp", "8.8.8.8:80")

	defer func() {
		if r := recover(); r != nil {
			checkErr(errors.Join(err, r.(error)))
		}
	}()
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	i.IntIP = localAddr.IP.String()
}

// GetAdapterName get the current local adapter name.
func (i *InternalData) GetAdapterName() {

	// get all network interfaces
	networkInterfaces, err := net.Interfaces()
	checkErr(err)

	// loop through all network interfaces
	for _, networkInterface := range networkInterfaces {

		// get interface info
		interfaceInfo, err := net.InterfaceByName(networkInterface.Name)
		checkErr(err)

		addresses, err := interfaceInfo.Addrs()
		checkErr(err)

		// loop through all addresses of the current interface
		for _, address := range addresses {

			// parse ip address
			ipAddr, _, err := net.ParseCIDR(address.String())
			checkErr(err)

			// get type of ip address (IPv4 or IPv6)
			ipType, err := getIPType(ipAddr.String())
			checkErr(err)

			// check if ip address match local ip address
			if ipAddr.String() == i.IntIP {
				i.AdapterName = networkInterface.Name
				i.IPAddressType = ipType
				return
			}
		}
	}
}

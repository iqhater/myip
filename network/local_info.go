package network

import (
	"context"
	"errors"
	"log"
	"net"
	"time"
)

// LocalInfo struct is a local adapter and local ip data.
type LocalInfo struct {
	AdapterName   string
	LocalIP       string
	IPAddressType IPType
}

// GetLocalIP get local ip addess.
func (l *LocalInfo) GetLocalIP() {

	defer func() {
		if r := recover(); r != nil {
			log.Fatal(r.(error).Error())
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	retry := func() {
		log.Println("Retry to get local ip address after 3 seconds...")
		time.Sleep(3 * time.Second)
		l.GetLocalIP()
	}

	var d net.Dialer
	conn, err := d.DialContext(ctx, "udp", "8.8.8.8:80")
	if err != nil {
		checkErr(err)
		retry()
		return
	}
	defer conn.Close()

	localAddr, ok := conn.LocalAddr().(*net.UDPAddr)
	if !ok {
		checkErr(errors.New("failed to get local address"))
		retry()
		return
	}

	l.LocalIP = localAddr.IP.String()
}

// GetAdapterName get the current local adapter name.
func (l *LocalInfo) GetAdapterName() {

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
			if ipAddr.String() == l.LocalIP {
				l.AdapterName = networkInterface.Name
				l.IPAddressType = ipType
				return
			}
		}
	}
}

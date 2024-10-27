package data

import (
	"errors"
	"net"
	"strings"
)

// InternalData struct is a local adapter and local ip data.
type InternalData struct {
	AdapterName string
	IntIP       string
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

	l, err := net.Interfaces()
	checkErr(err)

	for _, f := range l {

		byNameInterface, err := net.InterfaceByName(f.Name)
		checkErr(err)

		addr, err := byNameInterface.Addrs()
		checkErr(err)

		for _, v := range addr {

			if v.String()[:strings.Index(v.String(), "/")] == i.IntIP {
				i.AdapterName = f.Name
			}
		}
	}
}

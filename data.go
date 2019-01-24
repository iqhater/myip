package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"
	"time"
)

// IPinfo data struct
type IPinfo struct {
	AdapterName string
	IntIP       string
	ExtIP       string `json:"ip"`
	Country     string `json:"country"`
	CountryCode string `json:"cc"`
}

func (i *IPinfo) getExternalIP(url string) *IPinfo {

	timeout := time.Duration(10 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	resp, err := client.Get(url)
	if err != nil {
		log.Println(err)
		return &IPinfo{ExtIP: "Can't get the remote IP. Bad response from host!"}
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	checkErr(err)

	err = json.Unmarshal(body, &i)
	checkErr(err)

	r := &IPinfo{
		ExtIP:       i.ExtIP,
		Country:     i.Country,
		CountryCode: i.CountryCode,
	}
	return r
}

func (i *IPinfo) getInternalIP() {

	conn, err := net.Dial("udp", "8.8.8.8:80")
	checkErr(err)

	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	i.IntIP = localAddr.IP.String()
}

func (i *IPinfo) getAdapterName() {

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

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

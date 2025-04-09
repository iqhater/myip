package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/iqhater/myip/network"
	v "github.com/iqhater/myip/view"
)

func main() {

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "'myip' - A utility that displays your IP addresses\n\n"+
			"Description:\n"+
			"  - Shows both internal and external IP addresses\n"+
			"  - Displays the network adapter name for your local IP\n"+
			"  - Retrieves country and region information\n\n"+
			"Usage: myip [no arguments required]\n")
	}
	flag.Parse()

	n := network.NewInfo()

	n.GetLocalIP()
	n.GetAdapterName()
	v.PrintLocalInfo(n)

	sources := []string{
		"https://api.myip.com/",
		"https://api.miip.my",
		"https://ip.seeip.org/geoip",
		"https://api64.ipify.org?format=json",
		"https://api.myip.la/en?json",
		"https://ipapi.co/ip/",
		"https://api.seeip.org/jsonip",
		"https://myip.arens.online/json",
	}

	timeout, err := time.ParseDuration("15s")
	if err != nil {
		log.Println("Can't parse timeout duration: ", err)
	}

	done := make(chan struct{})
	defer close(done)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	for _, url := range sources {

		go func() {

			for {
				select {
				case <-ctx.Done():
					done <- struct{}{}
					return
				default:
					err := n.GetPublicIP(url, timeout)
					if err == nil {
						cancel()
						return
					}
					// debug errors
					// log.Println(err)
				}
			}
		}()
	}
	<-done

	v.PrintPublicInfo(n)
}

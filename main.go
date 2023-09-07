package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/iqhater/myip/data"
	v "github.com/iqhater/myip/view"
)

func main() {

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "'myip' is a tiny utility that shows the internal and external IP address.\nAlso 'myip' shows the adapter name where local IP is located.\nUsage: myip [no arguments are required].\n")
	}
	flag.Parse()

	d := data.NewIPData()

	d.GetInternalIP()
	d.GetAdapterName()
	v.PrintInternal(d)

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

	timeout, err := time.ParseDuration("10s")
	if err != nil {
		log.Println("Can't parse timeout duration: ", err)
	}

	done := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())

	for _, url := range sources {

		url := url

		go func(ctx context.Context) {

			for {
				select {
				case <-ctx.Done():
					done <- struct{}{}
					return
				default:
					err := d.GetExternalIP(url, timeout)
					if err == nil {
						cancel()
						return
					}
					// debug errors
					// log.Println(err)
				}
			}
		}(ctx)
	}
	<-done
	defer close(done)

	v.PrintExternal(d)
}

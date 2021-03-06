package main

import (
	"flag"
	"fmt"
	"os"

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

	done := make(chan struct{})
	go func() {
		d.GetExternalIP("https://api.myip.com/", 20)
		close(done)
	}()
	<-done

	v.PrintExternal(d)
}

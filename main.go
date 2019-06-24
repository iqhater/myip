// Compile binary for Raspberry Pi (linux)
// env GOOS=darwin GOARCH=amd64 go build

// Remove Symbol and Debug info at compile
// go build -ldflags "-s -w"

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

	d.GetExternalIP("https://api.myip.com/")
	d.GetInternalIP()
	d.GetAdapterName()
	v.PrintResults(d)
}

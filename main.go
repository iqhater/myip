// Compile binary for Raspberry Pi (linux)
// env GOOS=darwin GOARCH=amd64 go build

// Remove Symbol and Debug info at compile
// go build -ldflags "-s -w"

package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "'myip' is a tiny utility which print the internal and external IP address.\nAlso myip show the adapter name where local IP is located.\nUsage: myip [no arguments are required].\n")
	}
	flag.Parse()

	i := &IPinfo{}

	i.getExternalIP("https://api.myip.com/")
	i.getInternalIP()
	i.getAdapterName()
	i.printResults()
}

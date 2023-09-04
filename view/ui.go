package view

import (
	"fmt"
	"io"
	"os"

	"github.com/iqhater/myip/data"
)

// out writer for test check output something...
var out io.Writer = os.Stdout

// PrintInternal print all internal ip data
func PrintInternal(d *data.IPData) {
	fmt.Fprintln(out, "Local interface name:", d.AdapterName)
	fmt.Fprintln(out, "Internal IP:", d.IntIP)
}

// PrintExternal print all external ip data
func PrintExternal(d *data.IPData) {

	if d.ExtIP != "" {
		fmt.Fprintln(out, "External IP:", d.ExtIP)
	}

	if d.Country != "" {
		fmt.Fprintln(out, "Country:", d.Country)
	}

	if d.CountryCode != "" {
		fmt.Fprintln(out, "Country Code:", d.CountryCode)
	}

	if d.Region != "" {
		fmt.Fprintln(out, "Region:", d.Region)
	}
}

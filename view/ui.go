package view

import (
	"fmt"
	"io"
	"os"

	"github.com/fatih/color"
	"github.com/iqhater/myip/data"
)

// out writer for test check output something...
var out io.Writer = os.Stdout

// PrintInternal print all internal ip data
func PrintInternal(d *data.IPData) {
	color.New(color.FgMagenta).Fprintln(out, "Local interface name:", d.AdapterName)
	fmt.Fprintf(out, "  Internal IP (%s): %s \n", d.InternalData.IPAddressType, d.IntIP)
}

// PrintExternal print all external ip data
func PrintExternal(d *data.IPData) {

	if d.ExtIP != "" {
		color.New(color.FgGreen).Fprintf(out, "  External IP (%s): %s \n", d.ExternalData.IPAddressType, d.ExtIP)
	}

	if d.Country != "" {
		fmt.Fprintln(out, "	     Country:", d.Country)
	}

	if d.CountryCode != "" {
		fmt.Fprintln(out, "	Country Code:", d.CountryCode)
	}

	if d.Region != "" {
		fmt.Fprintln(out, "	      Region:", d.Region)
	}
}

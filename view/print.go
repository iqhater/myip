package view

import (
	"fmt"
	"io"
	"os"

	"github.com/fatih/color"
	"github.com/iqhater/myip/network"
)

// out writer for test check output something...
var out io.Writer = os.Stdout

// PrintLocalInfo print all internal (local) ip data
func PrintLocalInfo(d *network.Info) {
	color.New(color.FgMagenta).Fprintln(out, "Local interface name:", d.AdapterName)
	fmt.Fprintf(out, "     Local IP (%s): %s \n", d.LocalInfo.IPAddressType, d.LocalIP)
}

// PrintPublicInfo print all external (public) ip data
func PrintPublicInfo(d *network.Info) {

	if d.PublicIP != "" {
		color.New(color.FgGreen).Fprintf(out, "    Public IP (%s): %s \n", d.PublicInfo.IPAddressType, d.PublicIP)
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

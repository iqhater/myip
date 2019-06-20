package data

import (
	"fmt"
	"io"
	"os"
)

// out wrtire for test check output something...
var out io.Writer = os.Stdout

// PrintResults print all recived ip data
func (d *IPData) PrintResults() {

	fmt.Fprintln(out, "Local interface name:", d.AdapterName)
	fmt.Fprintln(out, "Internal IP:", d.IntIP)
	fmt.Fprintln(out, "External IP:", d.ExtIP)
	fmt.Fprintln(out, "Country:", d.Country)
	fmt.Fprintln(out, "Country Code:", d.CountryCode)
}

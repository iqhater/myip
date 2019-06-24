package view

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/iqhater/myip/data"
)

func TestPrintResults(t *testing.T) {

	// arrange
	d := data.NewIPData()
	d.GetExternalIP("https://api.myip.com/")
	d.GetInternalIP()
	d.GetAdapterName()

	buf := &bytes.Buffer{}
	out = buf
	defer buf.Reset()

	// act
	PrintResults(d)
	n, err := fmt.Fprintln(os.Stdout, buf.String())

	// assert
	if n <= 1 || err != nil {
		t.Errorf("Bad PrintResults! %d, %v\n", n, err)
	}
}

package data

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func TestPrintResults(t *testing.T) {

	// arrange
	d := NewIPData()
	d.GetExternalIP("https://api.myip.com/")
	d.GetInternalIP()
	d.GetAdapterName()

	buf := &bytes.Buffer{}
	out = buf
	defer buf.Reset()

	// act
	d.PrintResults()
	n, err := fmt.Fprintln(os.Stdout, buf.String())

	// assert
	if n <= 1 || err != nil {
		t.Errorf("Bad PrintResults! %d, %v\n", n, err)
	}
}

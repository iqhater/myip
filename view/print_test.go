package view

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/iqhater/myip/network"
)

func TestPrintLocalInfo(t *testing.T) {

	// arrange
	d := network.NewInfo()
	d.AdapterName = "test_adapter_name"
	d.LocalIP = "192.168.1.44"

	buf := &bytes.Buffer{}
	out = buf
	defer buf.Reset()

	// act
	PrintLocalInfo(d)
	n, err := fmt.Fprintln(os.Stdout, buf.String())

	// assert
	if n <= 1 || err != nil {
		t.Errorf("Bad PrintLocalInfo! %d, %v\n", n, err)
	}
}

func TestPrintPublicInfo(t *testing.T) {

	// arrange
	d := network.NewInfo()
	d.PublicIP = "78.67.56.89"
	d.Country = "Russian Federation"
	d.CountryCode = "RU"
	d.Region = "Moscow"

	buf := &bytes.Buffer{}
	out = buf
	defer buf.Reset()

	// act
	PrintPublicInfo(d)
	n, err := fmt.Fprintln(os.Stdout, buf.String())

	// assert
	if n <= 1 || err != nil {
		t.Errorf("Bad PrintlLocalInfo! %d, %v\n", n, err)
	}
}

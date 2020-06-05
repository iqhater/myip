package view

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/iqhater/myip/data"
)

func TestPrintInternal(t *testing.T) {

	// arrange
	d := data.NewIPData()
	d.AdapterName = "test_adapter_name"
	d.IntIP = "192.168.1.44"

	buf := &bytes.Buffer{}
	out = buf
	defer buf.Reset()

	// act
	PrintInternal(d)
	n, err := fmt.Fprintln(os.Stdout, buf.String())

	// assert
	if n <= 1 || err != nil {
		t.Errorf("Bad PrintInternal! %d, %v\n", n, err)
	}
}

func TestPrintExternal(t *testing.T) {

	// arrange
	d := data.NewIPData()
	d.ExtIP = "78.67.56.89"
	d.Country = "Russian Federation"
	d.CountryCode = "RU"

	buf := &bytes.Buffer{}
	out = buf
	defer buf.Reset()

	// act
	PrintExternal(d)
	n, err := fmt.Fprintln(os.Stdout, buf.String())

	// assert
	if n <= 1 || err != nil {
		t.Errorf("Bad PrintInternal! %d, %v\n", n, err)
	}
}

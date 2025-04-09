package network

import (
	"bytes"
	"errors"
	"log"
	"os"
	"testing"

	"go.uber.org/goleak"
)

func TestMain(m *testing.M) {

	// clear all connections before checking for leaks
	if testTransport != nil {
		cleanupTransport()
	}

	// check goroutines leaks
	goleak.VerifyTestMain(m)

	// run tests
	os.Exit(m.Run())
}

func TestNewInfoNotNil(t *testing.T) {

	// act
	result := NewInfo()

	// assert
	if result == nil {
		t.Errorf("Info is nil! %s\n", result)
	}
}

func TestCheckErr(t *testing.T) {

	// arrange
	msg := "New test ERROR!"
	err := errors.New(msg)

	buf := &bytes.Buffer{}
	log.SetOutput(buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()
	defer buf.Reset()

	// act
	checkErr(err)

	// assert
	if len(buf.String()) < len(msg) {
		t.Errorf("log.Println does not fired! Error message: %v Buffer: %v\n", msg, buf.String())
	}
}

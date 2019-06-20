package data

import (
	"bytes"
	"errors"
	"log"
	"os"
	"testing"
)

func TestNewIPDataNotNil(t *testing.T) {

	// act
	result := NewIPData()

	// assert
	if result == nil {
		t.Errorf("IPData is nil! %s\n", result)
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

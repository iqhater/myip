package data

import (
	"log"
	"testing"
	"time"
)

func TestGetExternalIPNotEmpty(t *testing.T) {

	// arrange
	e := ExternalData{}
	url := "https://api.myip.com/"

	// act
	timeout, err := time.ParseDuration("5s")
	if err != nil {
		log.Println("Can't parse timeout duration: ", err)
	}
	err = e.GetExternalIP(url, timeout)
	if err != nil {
		t.Error(err)
	}

	// assert
	if e.ExtIP == "" {
		t.Errorf("Result ip field must not be empty! ExtIP:%s\n", e.ExtIP)
	}
}

func TestGetExternalIPBadURL(t *testing.T) {

	// arrange
	e := ExternalData{}
	url := "https://badurl.com"

	// act
	timeout, _ := time.ParseDuration("1s")
	err := e.GetExternalIP(url, timeout)

	// assert
	if err == nil {
		t.Errorf("Bad url address! Enter the correct URL. %s\n", err)
	}
}

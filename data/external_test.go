package data

import (
	"testing"
)

func TestGetExternalIPNotEmpty(t *testing.T) {

	// arrange
	e := ExternalData{}
	url := "https://api.myip.com/"

	// act
	result := e.GetExternalIP(url, 1)

	// assert
	if result.ExtIP == "" || result.Country == "" || result.CountryCode == "" {
		t.Errorf("Result struct must not be empty! ExtIP:%s Country:%s CountryCode: %s\n", result.ExtIP, result.Country, result.CountryCode)
	}
}

func TestGetExternalIPBadURL(t *testing.T) {

	// arrange
	e := ExternalData{}
	url := "https://badurl.com"

	// act
	result := e.GetExternalIP(url, 1)

	// assert
	if result.ExtIP != "Can't get the remote IP. Bad response from host!" {
		t.Errorf("Bad url address! Enter the correct URL. %s\n", result.ExtIP)
	}
}

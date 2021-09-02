package data

import (
	"testing"
)

func TestGetExternalIPNotEmpty(t *testing.T) {

	// arrange
	e := ExternalData{}
	url := "https://ip.seeip.org/geoip"

	// act
	result := e.GetExternalIP(url, 5)

	// assert
	if result.ExtIP == "" || result.Country == "" || result.CountryCode == "" || result.Region == "" {
		t.Errorf("Result struct must not be empty! ExtIP:%s Country:%s CountryCode: %s Region: %s\n", result.ExtIP, result.Country, result.CountryCode, result.Region)
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

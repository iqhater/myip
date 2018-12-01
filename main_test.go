package main

import (
	"testing"
)

func TestGetExternalIPNotEmpty(t *testing.T) {

	i := IPinfo{}
	url := "https://api.myip.com/"

	result := i.getExternalIP(url)

	if result.ExtIP == "" || result.Country == "" || result.CountryCode == "" {
		t.Error("Result struct must not be empty!")
	}
}

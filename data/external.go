package data

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
)

// ExternalData struct is a external any api resource where data comes from
type ExternalData struct {
	ExtIP         string `json:"ip"`
	Country       string `json:"country"`
	CountryCode   string `json:"cc"`
	Region        string `json:"region"`
	IPAddressType IPType
}

// GetExternalIP method get response from url and return new ExternalData struct what is for???
func (e *ExternalData) GetExternalIP(url string, timeout time.Duration) error {

	client := http.Client{
		Timeout: time.Duration(timeout),
	}

	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("Bad response (not 200 status) from host!")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &e)
	if err != nil {
		return err
	}

	// add ip address type
	ipType, err := getIPType(e.ExtIP)
	checkErr(err)

	e.IPAddressType = ipType

	return nil
}

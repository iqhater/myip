package network

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
)

// PublicInfo struct is a public any api resource where data comes from
type PublicInfo struct {
	PublicIP      string `json:"ip"`
	Country       string `json:"country"`
	CountryCode   string `json:"cc"`
	Region        string `json:"region"`
	IPAddressType IPType
}

// GetPublicIPWithTransport method get response from url and return error if exists
func (p *PublicInfo) GetPublicIPWithTransport(url string, timeout time.Duration, transport *http.Transport) error {

	client := &http.Client{
		Timeout:   timeout,
		Transport: transport,
	}

	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("Status code is not 200!")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &p)
	if err != nil {
		return err
	}

	// add ip address type
	ipType, err := getIPType(p.PublicIP)
	checkErr(err)

	p.IPAddressType = ipType

	return nil
}

// GetPublicIP wrapper for GetPublicIPWithTransport
func (p *PublicInfo) GetPublicIP(url string, timeout time.Duration) error {

	transport := &http.Transport{
		IdleConnTimeout:   5 * time.Second,
		DisableKeepAlives: true,
	}
	return p.GetPublicIPWithTransport(url, timeout, transport)
}

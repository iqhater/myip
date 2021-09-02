package data

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// ExternalData struct is a external any api resource where data comes from
type ExternalData struct {
	ExtIP       string `json:"ip"`
	Country     string `json:"country"`
	CountryCode string `json:"country_code"`
	Region      string `json:"region"`
}

// GetExternalIP method get response from url and return new ExternalData struct what is for???
func (e *ExternalData) GetExternalIP(url string, timeout time.Duration) *ExternalData {

	t := time.Duration(timeout * time.Second)
	client := http.Client{
		Timeout: t,
	}

	resp, err := client.Get(url)
	if err != nil {
		log.Println(err)
		return &ExternalData{ExtIP: "Can't get the remote IP. Bad response from host!"}
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	checkErr(err)

	err = json.Unmarshal(body, &e)
	checkErr(err)

	return &ExternalData{
		ExtIP:       e.ExtIP,
		Country:     e.Country,
		CountryCode: e.CountryCode,
		Region:      e.Region,
	}
}

package network

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// global transport for tests
var testTransport = &http.Transport{
	DisableKeepAlives: true,
}

// helper function to close idle connections after each test
func cleanupTransport() {
	testTransport.CloseIdleConnections()
}

func TestGetPublicIPNotEmpty(t *testing.T) {

	// arrange
	p := PublicInfo{}
	url := "https://api.seeip.org/jsonip"

	timeout, err := time.ParseDuration("5s")
	if err != nil {
		log.Println("Can't parse timeout duration: ", err)
	}

	// act
	err = p.GetPublicIPWithTransport(url, timeout, testTransport)
	if err != nil {
		t.Error(err)
	}

	// assert
	if p.PublicIP == "" {
		t.Errorf("Result ip field must not be empty! PublicIP:%s\n", p.PublicIP)
	}

	// cleanup
	cleanupTransport()
}

func TestGetPublicIPBadURL(t *testing.T) {

	// arrange
	p := PublicInfo{}
	url := "https://badurl.com"

	// act
	timeout, _ := time.ParseDuration("1s")
	err := p.GetPublicIPWithTransport(url, timeout, testTransport)

	// assert
	if err == nil {
		t.Errorf("Bad url address! Enter the correct URL. %s\n", err)
	}

	// cleanup
	cleanupTransport()
}

func TestGetPublicIPWithTransport(t *testing.T) {

	// arrange
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"ip":"192.168.1.1","country":"Test Country","cc":"TC","region":"Test Region"}`))
	}))
	defer server.Close()

	tests := []struct {
		name          string
		serverHandler func(w http.ResponseWriter, r *http.Request)
		wantErr       bool
		expectedIP    string
		expectedType  IPType
	}{
		{
			name: "successful response",
			serverHandler: func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"ip":"192.168.1.1","country":"Test Country","cc":"TC","region":"Test Region"}`))
			},
			wantErr:      false,
			expectedIP:   "192.168.1.1",
			expectedType: IPv4,
		},
		{
			name: "non-200 response",
			serverHandler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusInternalServerError)
			},
			wantErr: true,
		},
		{
			name: "invalid json response",
			serverHandler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`invalid json`))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			// arrange
			testServer := httptest.NewServer(http.HandlerFunc(tt.serverHandler))
			defer testServer.Close()

			testTransport := &http.Transport{
				DisableKeepAlives: true,
			}

			p := &PublicInfo{}
			timeout := time.Second

			// act
			err := p.GetPublicIPWithTransport(testServer.URL, timeout, testTransport)

			// assert
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPublicIPWithTransport() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// If no error expected, check the results
			if !tt.wantErr {
				if p.PublicIP != tt.expectedIP {
					t.Errorf("Expected IP %s, got %s", tt.expectedIP, p.PublicIP)
				}
				if p.IPAddressType != tt.expectedType {
					t.Errorf("Expected IP type %s, got %s", tt.expectedType, p.IPAddressType)
				}
			}
		})
	}
}

func TestGetPublicIPOK(t *testing.T) {

	// arrange
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"ip":"192.168.1.1","country":"Test Country","cc":"TC"}`))
	}))
	defer server.Close()

	p := &PublicInfo{}
	timeout := time.Second

	// act
	err := p.GetPublicIP(server.URL, timeout)

	// assert
	if err != nil {
		t.Errorf("GetPublicIP() returned error: %v", err)
	}
	if p.PublicIP != "192.168.1.1" {
		t.Errorf("Expected IP 192.168.1.1, got %s", p.PublicIP)
	}
	if p.Country != "Test Country" {
		t.Errorf("Expected country 'Test Country', got %s", p.Country)
	}
	if p.CountryCode != "TC" {
		t.Errorf("Expected country code 'TC', got %s", p.CountryCode)
	}
	if p.IPAddressType != IPv4 {
		t.Errorf("Expected IP type IPv4, got %s", p.IPAddressType)
	}
}

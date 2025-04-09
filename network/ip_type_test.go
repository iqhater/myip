package network

import (
	"errors"
	"testing"
)

func TestGetIPType(t *testing.T) {

	// arrange
	tests := []struct {
		name     string
		ipStr    string
		expected IPType
		err      error
	}{
		{"Valid IPv4", "192.168.0.1", IPv4, nil},
		{"Valid IPv6", "2001:0db8:85a3:0000:0000:8a2e:0370:7334", IPv6, nil},
		{"Invalid IP", "invalid", "", errors.New("Invalid IP address!")},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			// act
			ipType, err := getIPType(tt.ipStr)

			// assert
			if ipType != tt.expected {
				t.Errorf("Expected IP type %s, but got %s", tt.expected, ipType)
			}

			if tt.err != nil {
				if err == nil || err.Error() != tt.err.Error() {
					t.Errorf("Expected error %s, but got %s", tt.err, err)
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error, but got %s", err)
				}
			}
		})
	}
}

package network

import "testing"

func TestGetLocalIPOK(t *testing.T) {

	// arrange
	l := LocalInfo{}

	// act
	l.GetLocalIP()
	result := l.LocalIP

	// assert
	if result == "" {
		t.Errorf("Wrong internal IP address! Shoud be format like this XX.XX.XX.XX %s\n", result)
	}
}

func TestGetAdapterNameOK(t *testing.T) {

	// arrange
	l := LocalInfo{}
	l.GetLocalIP()

	// act
	l.GetAdapterName()
	result := l.AdapterName

	// assert
	if result == "" {
		t.Errorf("Empty adapter name! %s\n", result)
	}
}

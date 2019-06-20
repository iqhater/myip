package data

import "testing"

func TestGetInternalIP(t *testing.T) {

	// arrange
	i := InternalData{}

	// act
	i.GetInternalIP()
	result := i.IntIP

	// assert
	if result == "" {
		t.Errorf("Wrong internal IP address! Shoud be fromat like this XX.XX.XX.XX %s\n", result)
	}
}

func TestGetAdapterName(t *testing.T) {

	// arrange
	i := InternalData{}
	i.GetInternalIP()

	// act
	i.GetAdapterName()
	result := i.AdapterName

	// assert
	if result == "" {
		t.Errorf("Empty adapter name! %s\n", result)
	}
}

package main

import "fmt"

func (i *IPinfo) printResults() {

	fmt.Println("Interface Name:", i.AdapterName)
	fmt.Println("Internal IP:", i.IntIP)
	fmt.Println("External IP:", i.ExtIP)
	fmt.Println("Country:", i.Country)
	fmt.Println("Country Code:", i.CountryCode)
}

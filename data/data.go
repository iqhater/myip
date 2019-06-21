package data

import (
	"log"
)

// IPData union both internal and external structs
type IPData struct {
	InternalData
	ExternalData
}

// NewIPData constructor init new *IPData struct
func NewIPData() *IPData {
	return &IPData{}
}

func checkErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

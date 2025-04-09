package network

import (
	"log"
)

// Info combines both internal and external network information
type Info struct {
	LocalInfo
	PublicInfo
}

// NewInfo constructor init new *Info struct
func NewInfo() *Info {
	return &Info{}
}

func checkErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

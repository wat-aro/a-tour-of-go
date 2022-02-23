package main

import (
	"fmt"
)

type IPAddr [4]byte

func (ip_addr IPAddr) String() string {
	result := ""
	for index, value := range ip_addr {
		if index == 0 {
			result = fmt.Sprintf("%d", value)
		} else {
			result = result + "." + fmt.Sprintf("%d", value)
		}
	}
	return result
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

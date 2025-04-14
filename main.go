// ip_enum.go
// Author: Sam Caldwell
// Description: Enumerate all IP addresses in a given CIDR block.
// Usage: go run ip_enum.go 192.168.1.0/30

package main

import (
	"fmt"
	"log"
	"math/big"
	"net"
	"os"
)

// expandCIDR returns a slice of IP addresses within the given CIDR block.
func expandCIDR(cidr string) ([]string, error) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}

	var ips []string
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); incIP(ip) {
		ips = append(ips, ip.String())
	}

	return ips[1 : len(ips)-1], nil
}

// incIP increments the given IP address by 1.
func incIP(ip net.IP) {
	ipv := big.NewInt(0).SetBytes(ip.To16())
	ipv.Add(ipv, big.NewInt(1))
	copy(ip, ipv.Bytes()[len(ipv.Bytes())-len(ip):])
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s <CIDR>\n", os.Args[0])
	}

	cidr := os.Args[1]
	ips, err := expandCIDR(cidr)
	if err != nil {
		log.Fatalf("Invalid CIDR: %v\n", err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

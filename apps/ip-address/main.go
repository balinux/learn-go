package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	interfaces, err := net.Interfaces()
	if err != nil {
		log.Fatal(err)
	}

	for _, iface := range interfaces {
		// mendapatkan  semua alamat ang terasosiasi dengan interfaces
		addrs, err := iface.Addrs()
		if err != nil {
			log.Printf("error getin address for interface %s:%v", iface.Name, err)
			continue
		}

		for _, addr := range addrs {
			// memproses alamat ip
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			// cek apakah ipv4 atay ipv6
			if ip != nil && (ip.To4() != nil || ip.To16() != nil) {
				fmt.Printf("interface %s, ip address: %s\n", iface.Name, ip.String())
			}

		}

	}
}

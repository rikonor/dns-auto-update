package main

import (
	"fmt"
	"time"

	"github.com/rikonor/dns-auto-update/dnsproviders"
	"github.com/rikonor/dns-auto-update/publicip"
)

func main() {
	p := dnsproviders.NewGoDaddy()

	for {
		publicIP, err := publicip.GetPublicIP()
		if err != nil {
			fmt.Println("Failed to get public IP:", err)
			continue
		}

		dnsIP, err := p.GetCurrentDNSIP()
		if err != nil {
			fmt.Println("Failed to get DNS IP:", err)
			continue
		}

		if publicIP != dnsIP {
			p.SetCurrentDNSIP(publicIP)
		}

		time.Sleep(1 * time.Minute)
	}
}

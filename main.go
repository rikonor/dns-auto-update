package main

import (
	"flag"
	"log"
	"time"

	"github.com/rikonor/dns-auto-update/dnsproviders"
	"github.com/rikonor/dns-auto-update/publicip"
)

func main() {
	var (
		domainName   = flag.String("d", "", "Domain name to track")
		providerName = flag.String("p", "godaddy", "DNS Provider to use")
	)
	flag.Parse()

	if *domainName == "" {
		log.Fatalln("Please provide a domain name")
	}
	if *providerName != "godaddy" {
		log.Fatalln("Only GoDaddy is supported at the moment")
	}

	p := dnsproviders.NewGoDaddy(*domainName)

	for {
		time.Sleep(5 * time.Minute)

		publicIP, err := publicip.GetPublicIP()
		if err != nil {
			log.Println("Failed to get public IP:", err)
			continue
		}

		dnsIP, err := p.GetCurrentDNSIP()
		if err != nil {
			log.Println("Failed to get DNS IP:", err)
			continue
		}

		if publicIP != dnsIP {
			err = p.SetCurrentDNSIP(publicIP)
			if err != nil {
				log.Println("Failed to update DNS IP:", err)
				continue
			}
		}
	}
}

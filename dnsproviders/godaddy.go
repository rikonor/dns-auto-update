package dnsproviders

import "fmt"

type GoDaddy struct {
}

func NewGoDaddy() *GoDaddy {
	return &GoDaddy{}
}

func (p *GoDaddy) GetCurrentDNSIP() (string, error) {
	return "1.2.3.4", nil
}

func (p *GoDaddy) SetCurrentDNSIP(ip string) error {
	fmt.Println("Setting DNS IP:", ip)
	return nil
}

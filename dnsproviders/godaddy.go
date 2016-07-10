package dnsproviders

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

const (
	baseAPIURL = "https://api.godaddy.com/v1"
)

type GoDaddy struct {
	APIKey    string
	APISecret string
}

type Record struct {
	Type string `json:"type"`
	Name string `json:"name"`
	Data string `json:"data"`
	TTL  uint   `json:"ttl"`
}

func NewGoDaddy() *GoDaddy {
	apiKey, apiSecret := ParseEnv()
	return &GoDaddy{apiKey, apiSecret}
}

func ParseEnv() (string, string) {
	// Try and get GODADDY_API_KEY and GODADDY_API_SECRET from environment
	// If not found - panic
	apiKey := os.Getenv("GODADDY_API_KEY")
	apiSecret := os.Getenv("GODADDY_API_SECRET")
	if apiKey == "" || apiSecret == "" {
		panic("Please provide both GODADDY_API_KEY and GODADDY_API_SECRET")
	}
	return apiKey, apiSecret
}

func (p *GoDaddy) SetAuthHeaders(req *http.Request) {
	authHeader := fmt.Sprintf("sso-key %s:%s", p.APIKey, p.APISecret)
	req.Header.Add("Authorization", authHeader)
}

func (p *GoDaddy) GetCurrentDNSIP() (string, error) {
	url := baseAPIURL + "/domains/knickknacklabs.com/records/A"

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	p.SetAuthHeaders(req)

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	ts := []Record{}
	err = json.NewDecoder(res.Body).Decode(&ts)
	if err != nil {
		return "", err
	}

	// Handle only the case of one record for now
	if len(ts) == 0 {
		return "", errors.New("no results found")
	}
	if len(ts) > 1 {
		return "", errors.New("don't know how to handle more then one result")
	}

	t := ts[0]

	return t.Data, nil
}

func (p *GoDaddy) SetCurrentDNSIP(ip string) error {
	fmt.Println("Setting DNS IP:", ip)
	return nil
}

package dnsproviders

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	baseAPIURL         = "https://api.godaddy.com/v1"
	defaultARecordName = "@"
	defaultARecordTTL  = 600
)

type GoDaddy struct {
	APIKey    string
	APISecret string

	// The domain to track
	Domain string

	// A record name, defaults to "@"
	ARecordName string
}

type Record struct {
	Type string `json:"type,omitempty"`
	Name string `json:"name,omitempty"`
	Data string `json:"data,omitempty"`
	TTL  uint   `json:"ttl,omitempty"`
}

func NewGoDaddy(domain string) *GoDaddy {
	apiKey, apiSecret := ParseEnv()
	return &GoDaddy{
		APIKey:      apiKey,
		APISecret:   apiSecret,
		Domain:      domain,
		ARecordName: defaultARecordName,
	}
}

// ParseEnv tries to get the needed env vars
// to properly instantiate a GoDaddy dnsprovider
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

// SetAuthHeaders will set the correct authorization headers
// so requests to the GoDaddy API pass authentication
func (p *GoDaddy) SetAuthHeaders(req *http.Request) {
	authHeader := fmt.Sprintf("sso-key %s:%s", p.APIKey, p.APISecret)
	req.Header.Add("Authorization", authHeader)
}

// GetCurrentDNSIP will make a request to the GoDaddy API
// to retrieve the relevant A DNS record
func (p *GoDaddy) GetCurrentDNSIP() (string, error) {
	url := baseAPIURL + "/domains/" + p.Domain + "/records/A"

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

// SetCurrentDNSIP will update the DNS settings with a new A record
func (p *GoDaddy) SetCurrentDNSIP(ip string) error {
	log.Println("Setting DNS IP:", ip)

	url := baseAPIURL + "/domains/" + p.Domain + "/records/A"

	// Prep request body
	pl := []Record{
		Record{
			Name: defaultARecordName,
			Data: ip,
			TTL:  defaultARecordTTL,
		},
	}
	bd, err := json.Marshal(&pl)
	if err != nil {
		return err
	}

	client := &http.Client{}
	req, err := http.NewRequest("PUT", url, bytes.NewReader(bd))
	if err != nil {
		return err
	}

	p.SetAuthHeaders(req)
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprint("Failed to set DNS IP, request code:", res.StatusCode))
	}

	return nil
}

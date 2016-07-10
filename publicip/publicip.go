package publicip

import (
	"io/ioutil"
	"net/http"
)

const (
	publicIPEndPoint = "http://ifconfig.co"
)

// GetPublicIP gets the running host's public ip from the ifconfig.co public service
func GetPublicIP() (string, error) {
	res, err := http.Get(publicIPEndPoint)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	// Need to trim newline
	if b[len(b)-1] == byte('\n') {
		b = b[:len(b)-1]
	}

	return string(b), nil
}

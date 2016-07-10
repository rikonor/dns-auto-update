package publicip

import (
	"io/ioutil"
	"net/http"
)

const (
	publicIPEndPoint = "http://ifconfig.co"
)

func GetPublicIP() (string, error) {
	res, err := http.Get(publicIPEndPoint)
	if err != nil {
		return "", err
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

package dozens

import (
	"net/http"
	"strings"
)

func SendRequest(method string, uri string, body string, auth *auth) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, uri, strings.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Auth-Token", auth.AuthToken)
	res, err := client.Do(req)
	return res, err
}

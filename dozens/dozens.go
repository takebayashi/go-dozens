package dozens

import (
	"net/http"
	"strings"
)

type Client struct {
	httpClient *http.Client
	auth       *auth
}

func NewClient(client *http.Client, user string, key string) (*Client, error) {
	auth, err := GetToken(user, key)
	if err != nil {
		return nil, err
	}
	c := &Client{httpClient: client, auth: auth}
	return c, nil
}

func (c *Client) newRequest(method string, uri string, body string) (*http.Request, error) {
	req, err := http.NewRequest(method, uri, strings.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Auth-Token", c.auth.AuthToken)
	return req, nil
}

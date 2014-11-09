package dozens

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type Domain struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func (c *Client) ListDomains() ([]Domain, error) {
	res, err := c.sendRequest("GET", "http://dozens.jp/api/zone.json", "")
	if err != nil {
		return nil, err
	}
	return parseDomainListResponse(res)
}

func (c *Client) AddDomain(name string, mail string) ([]Domain, error) {
	reqBody, err := json.Marshal(map[string]string{"name": name, "mailaddress": mail})
	if err != nil {
		return nil, err
	}
	res, err := c.sendRequest("POST", "http://dozens.jp/api/zone/create.json", string(reqBody))
	if err != nil {
		return nil, err
	}
	return parseDomainListResponse(res)
}

func (c *Client) DeleteDomain(zone Domain) ([]Domain, error) {
	res, err := c.sendRequest("DELETE", "http://dozens.jp/api/zone/delete/"+zone.Id+".json", "")
	if err != nil {
		return nil, err
	}
	return parseDomainListResponse(res)
}

func parseDomainListResponse(res *http.Response) ([]Domain, error) {
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, errors.New(string(resBody))
	}
	var zones map[string][]Domain
	json.Unmarshal(resBody, &zones)
	return zones["domain"], nil
}

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

func (c *Client) ListDomains() ([]*Domain, error) {
	req, err := c.newRequest("GET", apiRoot+"/zone.json", "")
	if err != nil {
		return nil, err
	}
	return c.fetchDomainList(req)
}

func (c *Client) GetDomain(name string) (*Domain, error) {
	list, err := c.ListDomains()
	if err != nil {
		return nil, err
	}
	for _, d := range list {
		if d.Name == name {
			return d, nil
		}
	}
	return nil, errors.New("Not Found")
}

func (c *Client) AddDomain(domain *Domain, mail string) ([]*Domain, error) {
	reqBody, err := json.Marshal(map[string]string{"name": domain.Name, "mailaddress": mail})
	if err != nil {
		return nil, err
	}
	req, err := c.newRequest("POST", apiRoot+"/zone/create.json", string(reqBody))
	if err != nil {
		return nil, err
	}
	return c.fetchDomainList(req)
}

func (c *Client) DeleteDomain(domain *Domain) ([]*Domain, error) {
	req, err := c.newRequest("DELETE", apiRoot+"/zone/delete/"+domain.Id+".json", "")
	if err != nil {
		return nil, err
	}
	return c.fetchDomainList(req)
}

func (c *Client) fetchDomainList(req *http.Request) ([]*Domain, error) {
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, errors.New(string(resBody))
	}
	var result map[string][]*Domain
	json.Unmarshal(resBody, &result)
	return result["domain"], nil
}

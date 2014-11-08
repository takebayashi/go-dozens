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

func GetDomains(auth *auth) ([]Domain, error) {
	res, err := SendRequest("GET", "http://dozens.jp/api/zone.json", "", auth)
	if err != nil {
		return nil, err
	}
	return parseDomainListResponse(res)
}

func AddDomain(auth *auth, name string, mail string) ([]Domain, error) {
	reqBody, err := json.Marshal(map[string]string{"name": name, "mailaddress": mail})
	if err != nil {
		return nil, err
	}
	res, err := SendRequest("POST", "http://dozens.jp/api/zone/create.json", string(reqBody), auth)
	if err != nil {
		return nil, err
	}
	return parseDomainListResponse(res)
}

func DeleteDomain(auth *auth, zone Domain) ([]Domain, error) {
	res, err := SendRequest("DELETE", "http://dozens.jp/api/zone/delete/"+zone.Id+".json", "", auth)
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

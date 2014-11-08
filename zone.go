package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type Zone struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func GetZones(auth *auth) ([]Zone, error) {
	res, err := SendRequest("GET", "http://dozens.jp/api/zone.json", "", auth)
	if err != nil {
		return nil, err
	}
	return parseZoneListResponse(res)
}

func AddZone(auth *auth, name string, mail string) ([]Zone, error) {
	reqBody, err := json.Marshal(map[string]string{"name": name, "mailaddress": mail})
	if err != nil {
		return nil, err
	}
	res, err := SendRequest("POST", "http://dozens.jp/api/zone/create.json", string(reqBody), auth)
	if err != nil {
		return nil, err
	}
	return parseZoneListResponse(res)
}

func DeleteZone(auth *auth, zone Zone) ([]Zone, error) {
	res, err := SendRequest("DELETE", "http://dozens.jp/api/zone/delete/"+zone.Id+".json", "", auth)
	if err != nil {
		return nil, err
	}
	return parseZoneListResponse(res)
}

func parseZoneListResponse(res *http.Response) ([]Zone, error) {
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, errors.New(string(resBody))
	}
	var zones map[string][]Zone
	json.Unmarshal(resBody, &zones)
	return zones["domain"], nil
}

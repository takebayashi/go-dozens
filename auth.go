package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type auth struct {
	AuthToken string `json:"auth_token"`
}

func GetToken(user string, key string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://dozens.jp/api/authorize.json", nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Auth-User", user)
	req.Header.Set("X-Auth-Key", key)
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	if res.StatusCode != http.StatusOK {
		return "", errors.New(string(resBody))
	}
	var auth auth
	json.Unmarshal(resBody, &auth)
	return auth.AuthToken, nil
}

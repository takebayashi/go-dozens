package dozens

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type auth struct {
	AuthToken string `json:"auth_token"`
}

func GetToken(user string, key string) (*auth, error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://dozens.jp/api/authorize.json", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Auth-User", user)
	req.Header.Set("X-Auth-Key", key)
	res, err := client.Do(req)
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
	var auth auth
	json.Unmarshal(resBody, &auth)
	return &auth, nil
}

package dozens

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type Record struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	Prio    string `json:"prio"`
	Content string `json:"content"`
	Ttl     string `json:"ttl"`
}

func GetRecords(auth *auth, zone Zone) ([]Record, error) {
	res, err := SendRequest("GET", "http://dozens.jp/api/record/"+zone.Name+".json", "", auth)
	if err != nil {
		return nil, err
	}
	return parseRecordListResponse(res)
}

func AddRecord(auth *auth, zone Zone, name string, typ string, prio string, content string, ttl string) ([]Record, error) {
	req, err := json.Marshal(map[string]string{"domain": zone.Name, "name": name, "type": typ, "prio": prio, "content": content, "ttl": ttl})
	if err != nil {
		return nil, err
	}
	res, err := SendRequest("POST", "http://dozens.jp/api/record/create.json", string(req), auth)
	if err != nil {
		return nil, err
	}
	return parseRecordListResponse(res)
}

func DeleteRecord(auth *auth, record Record) ([]Record, error) {
	res, err := SendRequest("DELETE", "http://dozens.jp/api/record/delete/"+record.Id+".json", "", auth)
	if err != nil {
		return nil, err
	}
	return parseRecordListResponse(res)
}

func parseRecordListResponse(res *http.Response) ([]Record, error) {
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, errors.New(string(resBody))
	}
	var records map[string][]Record
	json.Unmarshal(resBody, &records)
	return records["record"], nil
}

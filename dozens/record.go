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

func (c *Client) ListRecords(zone Domain) ([]Record, error) {
	res, err := c.sendRequest("GET", "http://dozens.jp/api/record/"+zone.Name+".json", "")
	if err != nil {
		return nil, err
	}
	return parseRecordListResponse(res)
}

func (c *Client) AddRecord(auth *auth, zone Domain, name string, typ string, prio string, content string, ttl string) ([]Record, error) {
	req, err := json.Marshal(map[string]string{"domain": zone.Name, "name": name, "type": typ, "prio": prio, "content": content, "ttl": ttl})
	if err != nil {
		return nil, err
	}
	res, err := c.sendRequest("POST", "http://dozens.jp/api/record/create.json", string(req))
	if err != nil {
		return nil, err
	}
	return parseRecordListResponse(res)
}

func (c *Client) DeleteRecord(record Record) ([]Record, error) {
	res, err := c.sendRequest("DELETE", "http://dozens.jp/api/record/delete/"+record.Id+".json", "")
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

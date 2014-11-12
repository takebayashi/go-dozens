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

func (c *Client) ListRecords(zone *Domain) ([]*Record, error) {
	req, err := c.newRequest("GET", apiRoot+"/record/"+zone.Name+".json", "")
	if err != nil {
		return nil, err
	}
	return c.fetchRecordList(req)
}

func (c *Client) AddRecord(zone *Domain, record *Record) ([]*Record, error) {
	reqBody, err := json.Marshal(map[string]string{"domain": zone.Name, "name": record.Name, "type": record.Type, "prio": record.Prio, "content": record.Content, "ttl": record.Ttl})
	if err != nil {
		return nil, err
	}
	req, err := c.newRequest("POST", apiRoot+"/record/create.json", string(reqBody))
	if err != nil {
		return nil, err
	}
	return c.fetchRecordList(req)
}

func (c *Client) DeleteRecord(record *Record) ([]*Record, error) {
	req, err := c.newRequest("DELETE", apiRoot+"/record/delete/"+record.Id+".json", "")
	if err != nil {
		return nil, err
	}
	return c.fetchRecordList(req)
}

func (c *Client) fetchRecordList(req *http.Request) ([]*Record, error) {
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
	var records map[string][]*Record
	json.Unmarshal(resBody, &records)
	return records["record"], nil
}

package hcloud

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type CreateZoneRequest struct {
	domain string `json:"domain"`
}

type CreateZoneResponse struct {
}

type DeleteZoneResponse struct {
}

type BlankResponse struct {
}

type Record struct {
	DNSRecordID int    `json:"DNSRecordsID"`
	RecordType  string `json:"RecordType"`
	TTL         int    `json:"TTL"`
	RecordValue string `json:"RecordValue"`
	OrgID       string `json:"OrgID"`
	RecordName  string `json:"RecordName"`
	ZoneName    string `json:"ZoneName"`
	Weight      int    `json:"Weight"`
	Priority    int    `json:"Priority"`
}

type ListRecords struct {
	Records []Record
}

type CreateRecordRequest struct {
	Priority    string `json:"priority"`
	RecordName  string `json:"record_name"`
	RecordType  string `json:"record_type"`
	RecordValue string `json:"record_value"`
	TTL         string `json:"ttl"`
}

// TODO: Get zone
// TODO: List zones
// TODO: Update zone
// TODO: List Records
// TODO: Get Record
// TODO: Create Record
// TODO: Update Record
// TODO: Delete Record

func (c *Client) CreateZone(domain string) (string, error) {

	requestData := []byte(`{
		"domain": "` + domain + `"
	}`)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/node-api/dns/zones", c.BaseURL), bytes.NewBuffer(requestData))
	if err != nil {
		return "error", err
	}

	res := CreateZoneResponse{}
	if err := c.sendRequest(req, &res); err != nil {
		return "error", err
	}
	return "success", nil
}

func (c *Client) DeleteZone(domain string) (string, error) {

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/node-api/dns/zones/%s", c.BaseURL, domain), nil)
	if err != nil {
		return "error", err
	}

	res := DeleteZoneResponse{}
	if err := c.sendRequest(req, &res); err != nil {
		return "error", err
	}
	return "success", nil
}

func (c *Client) CreateRecord(domain string, record *CreateRecordRequest) (string, error) {

	record = &CreateRecordRequest{
		Priority:    record.Priority,
		RecordName:  record.RecordName,
		RecordType:  record.RecordType,
		RecordValue: record.RecordValue,
		TTL:         record.TTL,
	}

	body, _ := json.Marshal(record)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/node-api/dns/zones/%s/records", c.BaseURL, domain), bytes.NewBuffer(body))
	if err != nil {
		return "error", err
	}

	res := BlankResponse{}
	if err := c.sendRequest(req, &res); err != nil {
		return "error", err
	}
	return "success", nil
}

func (c *Client) GetRecord(domain string, recordId string) (*Record, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/node-api/dns/zones/%s/records/%s", c.BaseURL, domain, recordId), nil)
	if err != nil {
		return nil, err
	}

	res := Record{}
	if err := c.sendRequest(req, &res); err != nil {
		return &res, err
	}
	return &res, nil
}

func (c *Client) DeleteRecord(domain string, recordId string) (string, error) {

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/node-api/dns/zones/%s/records/%s", c.BaseURL, domain, recordId), nil)
	if err != nil {
		return "error", err
	}

	res := BlankResponse{}
	if err := c.sendRequest(req, &res); err != nil {
		return "error", err
	}
	return "success", nil
}

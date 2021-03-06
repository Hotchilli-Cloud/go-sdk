package hcloud

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

const (
	BaseURLV1 = "https://api.hotchilli.com" //"http://localhost:3500"
)

type Client struct {
	BaseURL    string
	apiKey     string
	HTTPClient *http.Client
}

type SuccessResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Object  string      `json:"object"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func NewClient(apiKey string) *Client {
	return &Client{
		BaseURL: BaseURLV1,
		apiKey:  apiKey,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}
func (c *Client) sendRequest(req *http.Request, v interface{}) error {
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		var errRes ErrorResponse
		if err = json.NewDecoder(res.Body).Decode(&errRes); err == nil {
			return errors.New(errRes.Message)
		}
		return fmt.Errorf("Unknown errro, status code : %d", res.StatusCode)
	}
	fullResponse := SuccessResponse{
		Data: v,
	}
	if err = json.NewDecoder(res.Body).Decode(&fullResponse); err != nil {
		return err
	}
	return nil
}

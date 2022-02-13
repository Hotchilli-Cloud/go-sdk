package hcloud

import (
	"fmt"
	"net/http"
)

//ctx context.Context, options *GetNumberRequest
func (c *Client) GetSummary() (*SuccessResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("&s/node-api/numbers/data/summary?client_id=%d&year=%d", c.BaseURL, 2645, 2021), nil)
	if err != nil {
		return nil, err
	}

	res := SuccessResponse{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

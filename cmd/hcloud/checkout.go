package hcloud

import (
	"fmt"
	"net/http"
)

type SessionType struct {
	ID                string      `json:"id"`
	Object            string      `json:"object"`
	CreatedTime       int         `json:"created_time"`
	PromotionCode     string      `json:"promotion_code"`
	PaymentIntent     string      `json:"payment_intent"`
	PaymentStatus     string      `json:"payment_status"`
	TaxRate           float32     `json:"tax_rate"`
	AmountSubtotal    float32     `json:"amount_subtotal"`
	AmountTotal       float32     `json:"amount_total"`
	AmountTax         float32     `json:"amount_tax"`
	Currency          string      `json:"currency"`
	CustomerID        string      `json:"customer_id"`
	CustomerPaymentID string      `json:"customer_payment_id"`
	Livemode          bool        `json:"livemode"`
	metadata          interface{} `json:"metadata"`
	LineItems         []string    `json:"line_items"`
}

func (c *Client) GetCheckoutSession(sessionId string) (*SessionType, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/node-api/checkout/session/%s?rt=sdk", c.BaseURL, sessionId), nil)
	if err != nil {
		return nil, err
	}

	res := SessionType{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *Client) CreateCheckoutSession() (*SessionType, error) {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/node-api/checkout/session?rt=sdk", c.BaseURL), nil)
	if err != nil {
		return nil, err
	}

	res := SessionType{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

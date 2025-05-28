package martianpay

import (
	"fmt"

	"github.com/MartianPay/martianpay-go-sample/pkg/developer"
)

type PaymentIntentCreateRequest struct {
	developer.PaymentIntentParams
}

type PaymentIntentCreateResp struct {
	developer.PaymentIntent
}

func (c *Client) CreatePaymentIntent(req PaymentIntentCreateRequest) (*PaymentIntentCreateResp, error) {
	var response PaymentIntentCreateResp
	err := c.sendRequest("POST", "/v1/payment_intents", req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

type PaymentIntentUpdateRequest struct {
	ID                string
	PaymentLinkId     *string                         `json:"payment_link_id"`
	PaymentMethodType *developer.PaymentMethodType    `json:"payment_method_type"` // crypto, visa, mastercard, apple pay, google pay, etc.
	PaymentMethodData *developer.PaymentMethodOptions `json:"payment_method_options"`
}

type PaymentIntentUpdateResp struct {
	developer.PaymentIntent
}

func (c *Client) UpdatePaymentIntent(req PaymentIntentUpdateRequest) (*PaymentIntentUpdateResp, error) {
	var response PaymentIntentUpdateResp
	err := c.sendRequest("POST", fmt.Sprintf("/v1/payment_intents/%s", req.ID), req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

type PaymentIntentGetReq struct {
	ID string
}

type PaymentIntentGetResp struct {
	developer.PaymentIntent
}

// GetPaymentIntent retrieves a specific payment intent by ID
func (c *Client) GetPaymentIntent(req PaymentIntentGetReq) (*PaymentIntentGetResp, error) {
	var response PaymentIntentGetResp
	err := c.sendRequest("GET", fmt.Sprintf("/v1/payment_intents/%s", req.ID), nil, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

type PaymentIntentListReq struct {
	// Pagination
	Page     int32 `json:"page" binding:"min=0"`               // Page number, starting from 0
	PageSize int32 `json:"page_size" binding:"required,min=1"` // Items per page

	// Filters
	Customer      *string `json:"customer,omitempty"`       // Filter by customer
	CustomerEmail *string `json:"customer_email,omitempty"` // Filter by customer email
}

type PaymentIntentListResp struct {
	PaymentIntents []*developer.PaymentIntent `json:"payment_intents"` // List of payment intents
	Total          int64                      `json:"total"`           // Total number of records matching the filters
	Page           int32                      `json:"page"`            // Current page number
	PageSize       int32                      `json:"page_size"`       // Items per page
}

// ListPaymentIntents retrieves a list of payment intents based on the provided parameters
func (c *Client) ListPaymentIntents(req PaymentIntentListReq) (*PaymentIntentListResp, error) {
	var response PaymentIntentListResp
	err := c.sendRequest("POST", "/v1/payment_intents/list", req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

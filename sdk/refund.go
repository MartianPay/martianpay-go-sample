package martianpay

import (
	"fmt"

	"github.com/MartianPay/martianpay-go-sample/pkg/developer"
)

type RefundCreateRequest struct {
	developer.RefundParams
}

type RefundCreateResp struct {
	Refunds []developer.Refund `json:"refunds"`
}

// CreateRefund creates a new refund
func (c *Client) CreateRefund(req RefundCreateRequest) (*RefundCreateResp, error) {
	var response RefundCreateResp
	err := c.sendRequest("POST", "/v1/refunds", req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

type RefundGetRequest struct {
	ID string
}

type RefundGetResp struct {
	developer.Refund
}

// GetRefund retrieves a specific refund by ID
func (c *Client) GetRefund(req RefundGetRequest) (*RefundGetResp, error) {
	var response RefundGetResp
	err := c.sendRequest("GET", fmt.Sprintf("/v1/refunds/%s", req.ID), nil, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

type RefundListRequest struct {
	// Pagination
	Page     int32 `json:"page" binding:"min=0"`               // Page number, starting from 0
	PageSize int32 `json:"page_size" binding:"required,min=1"` // Items per page

	// Filters
	PaymentIntent *string `json:"payment_intent,omitempty"` // Filter by payment intent ID
}

type RefundListResp struct {
	Refunds  []developer.Refund `json:"refunds"`   // List of refunds
	Total    int64              `json:"total"`     // Total number of records matching the filters
	Page     int32              `json:"page"`      // Current page number
	PageSize int32              `json:"page_size"` // Items per page
}

// ListRefunds retrieves a list of refunds based on the provided parameters
func (c *Client) ListRefunds(req RefundListRequest) (*RefundListResp, error) {
	var response RefundListResp
	err := c.sendRequest("POST", "/v1/refunds/list", req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

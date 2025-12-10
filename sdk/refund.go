package martianpay

import (
	"fmt"

	"github.com/MartianPay/martianpay-go-sample/pkg/developer"
)

// CreateRefund creates a refund for a payment
func (c *Client) CreateRefund(req *developer.RefundCreateRequest) (*developer.RefundCreateResp, error) {
	var response developer.RefundCreateResp
	err := c.sendRequest("POST", "/v1/refunds", req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// GetRefund retrieves a specific refund by ID
func (c *Client) GetRefund(refundID string) (*developer.RefundGetResp, error) {
	var response developer.RefundGetResp
	err := c.sendRequest("GET", fmt.Sprintf("/v1/refunds/%s", refundID), nil, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// ListRefunds lists all refunds with optional filters
func (c *Client) ListRefunds(req *developer.RefundListRequest) (*developer.RefundListResp, error) {
	var response developer.RefundListResp
	err := c.sendRequestWithQuery("GET", "/v1/refunds", req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

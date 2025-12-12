package martianpay

import (
	"fmt"

	"github.com/MartianPay/martianpay-go-sample/pkg/developer"
)

// ListPaymentLinks lists payment links with pagination
func (c *Client) ListPaymentLinks(params *developer.PaymentLinkListRequest) (*developer.PaymentLinkListResponse, error) {
	var resp developer.PaymentLinkListResponse
	err := c.sendRequestWithQuery("GET", "/v1/payment_links", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreatePaymentLink creates a payment link
func (c *Client) CreatePaymentLink(params *developer.PaymentLinkCreateRequest) (*developer.PaymentLink, error) {
	var resp developer.PaymentLink
	err := c.sendRequest("POST", "/v1/payment_links", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPaymentLink retrieves payment link details
func (c *Client) GetPaymentLink(linkID string) (*developer.PaymentLink, error) {
	path := fmt.Sprintf("/v1/payment_links/%s", linkID)
	var resp developer.PaymentLink
	err := c.sendRequest("GET", path, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdatePaymentLink updates payment link active status
func (c *Client) UpdatePaymentLink(linkID string, params *developer.PaymentLinkUpdateRequest) (*developer.PaymentLink, error) {
	path := fmt.Sprintf("/v1/payment_links/%s", linkID)
	var resp developer.PaymentLink
	err := c.sendRequest("POST", path, params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeletePaymentLink deletes inactive payment link
func (c *Client) DeletePaymentLink(linkID string) error {
	path := fmt.Sprintf("/v1/payment_links/%s", linkID)
	return c.sendRequest("DELETE", path, nil, nil)
}

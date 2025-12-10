package martianpay

import (
	"fmt"

	"github.com/MartianPay/martianpay-go-sample/pkg/developer"
)

// CreatePaymentIntent creates a new payment intent
func (c *Client) CreatePaymentIntent(req *developer.PaymentIntentCreateRequest) (*developer.PaymentIntentCreateResp, error) {
	var response developer.PaymentIntentCreateResp
	err := c.sendRequest("POST", "/v1/payment_intents", req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// UpdatePaymentIntent updates a payment intent
func (c *Client) UpdatePaymentIntent(id string, req *developer.PaymentIntentUpdateRequest) (*developer.PaymentIntentUpdateResp, error) {
	var response developer.PaymentIntentUpdateResp
	err := c.sendRequest("POST", fmt.Sprintf("/v1/payment_intents/%s", id), req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// GetPaymentIntent retrieves a specific payment intent by ID
func (c *Client) GetPaymentIntent(id string) (*developer.PaymentIntentGetResp, error) {
	var response developer.PaymentIntentGetResp
	err := c.sendRequest("GET", fmt.Sprintf("/v1/payment_intents/%s", id), nil, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// ListPaymentIntents retrieves a list of payment intents based on the provided parameters
func (c *Client) ListPaymentIntents(req *developer.PaymentIntentListRequest) (*developer.PaymentIntentListResp, error) {
	var response developer.PaymentIntentListResp
	err := c.sendRequestWithQuery("GET", "/v1/payment_intents", req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// CancelPaymentIntent cancels a payment intent
func (c *Client) CancelPaymentIntent(id string, req *developer.PaymentIntentCancelRequest) (*developer.PaymentIntentUpdateResp, error) {
	var response developer.PaymentIntentUpdateResp
	err := c.sendRequest("POST", fmt.Sprintf("/v1/payment_intents/%s/cancel", id), req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// CreatePaymentIntentLink creates a payment intent with payment link
func (c *Client) CreatePaymentIntentLink(req *developer.PaymentIntentLinkCreateRequest) (*developer.PaymentIntentLinkCreateResp, error) {
	var response developer.PaymentIntentLinkCreateResp
	err := c.sendRequest("POST", "/v1/payment_intents/link", req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// UpdatePaymentIntentLink updates a payment intent link
func (c *Client) UpdatePaymentIntentLink(id string, req *developer.PaymentIntentLinkUpdateRequest) (*developer.PaymentIntentUpdateResp, error) {
	var response developer.PaymentIntentUpdateResp
	err := c.sendRequest("POST", fmt.Sprintf("/v1/payment_intents/%s/link", id), req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// CreatePaymentIntentInvoice creates a payment intent with invoice
func (c *Client) CreatePaymentIntentInvoice(req *developer.PaymentIntentInvoiceCreateRequest) (*developer.PaymentIntentInvoiceCreateResponse, error) {
	var response developer.PaymentIntentInvoiceCreateResponse
	err := c.sendRequest("POST", "/v1/payment_intents/invoice", req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

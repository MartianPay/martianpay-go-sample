// Package martianpay provides SDK methods for managing payment intents.
// Payment intents represent a payment process initiated by the merchant for collecting payment from customers.
package martianpay

import (
	"fmt"

	"github.com/MartianPay/martianpay-go-sample/pkg/developer"
)

// CreatePaymentIntent creates a new payment intent.
// A payment intent tracks the entire payment lifecycle from creation to completion or failure.
//
// Parameters:
//   - req: Request containing amount, currency, customer info, and payment method details
//
// Returns:
//   - *developer.PaymentIntentCreateResp: The created payment intent with client secret for payment confirmation
//   - error: nil on success, error on failure
func (c *Client) CreatePaymentIntent(req *developer.PaymentIntentCreateRequest) (*developer.PaymentIntentCreateResp, error) {
	var response developer.PaymentIntentCreateResp
	err := c.sendRequest("POST", "/v1/payment_intents", req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// UpdatePaymentIntent updates an existing payment intent.
// Can modify amount, currency, customer information, and metadata before payment is confirmed.
//
// Parameters:
//   - id: The unique identifier of the payment intent to update
//   - req: Updated payment intent fields
//
// Returns:
//   - *developer.PaymentIntentUpdateResp: The updated payment intent details
//   - error: nil on success, error on failure (e.g., payment already confirmed)
func (c *Client) UpdatePaymentIntent(id string, req *developer.PaymentIntentUpdateRequest) (*developer.PaymentIntentUpdateResp, error) {
	var response developer.PaymentIntentUpdateResp
	err := c.sendRequest("POST", fmt.Sprintf("/v1/payment_intents/%s", id), req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// GetPaymentIntent retrieves detailed information about a specific payment intent.
// Includes payment status, amount, customer details, and transaction history.
//
// Parameters:
//   - id: The unique identifier of the payment intent
//
// Returns:
//   - *developer.PaymentIntentGetResp: Complete payment intent details
//   - error: nil on success, error on failure (e.g., payment intent not found)
func (c *Client) GetPaymentIntent(id string) (*developer.PaymentIntentGetResp, error) {
	var response developer.PaymentIntentGetResp
	err := c.sendRequest("GET", fmt.Sprintf("/v1/payment_intents/%s", id), nil, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// ListPaymentIntents retrieves a paginated list of payment intents.
// Can be filtered by status, customer, date range, and other criteria.
//
// Parameters:
//   - req: Query parameters including pagination, filters by status, customer, date range, etc.
//
// Returns:
//   - *developer.PaymentIntentListResp: List of payment intents with pagination metadata
//   - error: nil on success, error on failure
func (c *Client) ListPaymentIntents(req *developer.PaymentIntentListRequest) (*developer.PaymentIntentListResp, error) {
	var response developer.PaymentIntentListResp
	err := c.sendRequestWithQuery("GET", "/v1/payment_intents", req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// CancelPaymentIntent cancels a payment intent that has not yet been confirmed.
// Once cancelled, the payment intent cannot be resumed and a new one must be created.
//
// Parameters:
//   - id: The unique identifier of the payment intent to cancel
//   - req: Cancellation request containing optional cancellation reason
//
// Returns:
//   - *developer.PaymentIntentUpdateResp: The cancelled payment intent with updated status
//   - error: nil on success, error on failure (e.g., payment already confirmed)
func (c *Client) CancelPaymentIntent(id string, req *developer.PaymentIntentCancelRequest) (*developer.PaymentIntentUpdateResp, error) {
	var response developer.PaymentIntentUpdateResp
	err := c.sendRequest("POST", fmt.Sprintf("/v1/payment_intents/%s/cancel", id), req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// CreatePaymentIntentLink creates a payment intent with an associated payment link.
// The payment link allows customers to complete payment via a hosted payment page.
//
// Parameters:
//   - req: Request containing payment intent details and link configuration (expiry, redirect URLs)
//
// Returns:
//   - *developer.PaymentIntentLinkCreateResp: The created payment intent with payment link URL
//   - error: nil on success, error on failure
func (c *Client) CreatePaymentIntentLink(req *developer.PaymentIntentLinkCreateRequest) (*developer.PaymentIntentLinkCreateResp, error) {
	var response developer.PaymentIntentLinkCreateResp
	err := c.sendRequest("POST", "/v1/payment_intents/link", req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// UpdatePaymentIntentLink updates the payment link associated with a payment intent.
// Can modify link expiry time and redirect URLs.
//
// Parameters:
//   - id: The unique identifier of the payment intent
//   - req: Updated link configuration including expiry and redirect URLs
//
// Returns:
//   - *developer.PaymentIntentUpdateResp: The updated payment intent with modified link
//   - error: nil on success, error on failure
func (c *Client) UpdatePaymentIntentLink(id string, req *developer.PaymentIntentLinkUpdateRequest) (*developer.PaymentIntentUpdateResp, error) {
	var response developer.PaymentIntentUpdateResp
	err := c.sendRequest("POST", fmt.Sprintf("/v1/payment_intents/%s/link", id), req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// CreatePaymentIntentInvoice creates a payment intent with an associated invoice.
// The invoice contains itemized billing information and can be sent to customers.
//
// Parameters:
//   - req: Request containing payment intent details and invoice line items
//
// Returns:
//   - *developer.PaymentIntentInvoiceCreateResponse: The created payment intent with invoice details
//   - error: nil on success, error on failure
func (c *Client) CreatePaymentIntentInvoice(req *developer.PaymentIntentInvoiceCreateRequest) (*developer.PaymentIntentInvoiceCreateResponse, error) {
	var response developer.PaymentIntentInvoiceCreateResponse
	err := c.sendRequest("POST", "/v1/payment_intents/invoice", req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

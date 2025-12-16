// Package martianpay provides SDK methods for managing payment links.
// Payment links are shareable URLs that allow customers to make payments via a hosted checkout page.
package martianpay

import (
	"fmt"

	"github.com/MartianPay/martianpay-go-sample/pkg/developer"
)

// ListPaymentLinks retrieves a paginated list of payment links.
// Can be filtered by active status and other criteria.
//
// Parameters:
//   - params: Query parameters including pagination and filters (active status, creation date, etc.)
//
// Returns:
//   - *developer.PaymentLinkListResponse: List of payment links with pagination metadata
//   - error: nil on success, error on failure
func (c *Client) ListPaymentLinks(params *developer.PaymentLinkListRequest) (*developer.PaymentLinkListResponse, error) {
	var resp developer.PaymentLinkListResponse
	err := c.sendRequestWithQuery("GET", "/v1/payment_links", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreatePaymentLink creates a new payment link.
// Payment links can be shared with customers via email, SMS, or social media for easy payment collection.
//
// Parameters:
//   - params: Request containing link configuration (amount, currency, description, expiry, redirect URLs)
//
// Returns:
//   - *developer.PaymentLink: The created payment link with shareable URL
//   - error: nil on success, error on failure
func (c *Client) CreatePaymentLink(params *developer.PaymentLinkCreateRequest) (*developer.PaymentLink, error) {
	var resp developer.PaymentLink
	err := c.sendRequest("POST", "/v1/payment_links", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPaymentLink retrieves detailed information about a specific payment link.
// Includes link URL, configuration, status, and payment statistics.
//
// Parameters:
//   - linkID: The unique identifier of the payment link
//
// Returns:
//   - *developer.PaymentLink: Complete payment link details
//   - error: nil on success, error on failure (e.g., link not found)
func (c *Client) GetPaymentLink(linkID string) (*developer.PaymentLink, error) {
	path := fmt.Sprintf("/v1/payment_links/%s", linkID)
	var resp developer.PaymentLink
	err := c.sendRequest("GET", path, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdatePaymentLink updates an existing payment link's configuration.
// Can modify active status, redirect URLs, and metadata.
//
// Parameters:
//   - linkID: The unique identifier of the payment link to update
//   - params: Updated link fields (active status, redirect URLs, metadata)
//
// Returns:
//   - *developer.PaymentLink: The updated payment link details
//   - error: nil on success, error on failure
func (c *Client) UpdatePaymentLink(linkID string, params *developer.PaymentLinkUpdateRequest) (*developer.PaymentLink, error) {
	path := fmt.Sprintf("/v1/payment_links/%s", linkID)
	var resp developer.PaymentLink
	err := c.sendRequest("POST", path, params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeletePaymentLink permanently deletes an inactive payment link.
// Only links with active=false can be deleted. Active links must be deactivated first.
//
// Parameters:
//   - linkID: The unique identifier of the payment link to delete
//
// Returns:
//   - error: nil on success, error on failure (e.g., link is active or has pending payments)
func (c *Client) DeletePaymentLink(linkID string) error {
	path := fmt.Sprintf("/v1/payment_links/%s", linkID)
	return c.sendRequest("DELETE", path, nil, nil)
}

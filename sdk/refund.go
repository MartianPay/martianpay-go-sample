// Package martianpay provides SDK methods for managing refunds.
// Refunds allow merchants to return payment to customers for returned goods or cancelled services.
package martianpay

import (
	"fmt"

	"github.com/MartianPay/martianpay-go-sample/pkg/developer"
)

// CreateRefund creates a new refund for a payment.
// Refunds can be full or partial, and may take several days to process depending on the payment method.
//
// Parameters:
//   - req: Request containing payment intent ID, amount to refund, reason, and optional metadata
//
// Returns:
//   - *developer.RefundCreateResp: The created refund with ID and status
//   - error: nil on success, error on failure (e.g., insufficient funds or payment not captured)
func (c *Client) CreateRefund(req *developer.RefundCreateRequest) (*developer.RefundCreateResp, error) {
	var response developer.RefundCreateResp
	err := c.sendRequest("POST", "/v1/refunds", req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// GetRefund retrieves detailed information about a specific refund.
// Includes refund status, amount, reason, and associated payment intent.
//
// Parameters:
//   - refundID: The unique identifier of the refund
//
// Returns:
//   - *developer.RefundGetResp: Complete refund details
//   - error: nil on success, error on failure (e.g., refund not found)
func (c *Client) GetRefund(refundID string) (*developer.RefundGetResp, error) {
	var response developer.RefundGetResp
	err := c.sendRequest("GET", fmt.Sprintf("/v1/refunds/%s", refundID), nil, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// ListRefunds retrieves a paginated list of refunds.
// Can be filtered by payment intent, status, date range, and other criteria.
//
// Parameters:
//   - req: Query parameters including pagination and filters (payment intent, status, date range, etc.)
//
// Returns:
//   - *developer.RefundListResp: List of refunds with pagination metadata
//   - error: nil on success, error on failure
func (c *Client) ListRefunds(req *developer.RefundListRequest) (*developer.RefundListResp, error) {
	var response developer.RefundListResp
	err := c.sendRequestWithQuery("GET", "/v1/refunds", req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

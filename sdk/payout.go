// Package martianpay provides SDK methods for managing payouts.
// Payouts allow merchants to transfer funds from their MartianPay balance to their bank account.
package martianpay

import (
	"fmt"

	"github.com/MartianPay/martianpay-go-sample/pkg/developer"
)

// PreviewPayout calculates the fees and final amount for a payout before creation.
// Useful for showing merchants the expected payout details before confirming the transfer.
//
// Parameters:
//   - req: Request containing payout amount, currency, and destination details
//
// Returns:
//   - *developer.PayoutPreviewResp: Preview details including fees, net amount, and arrival time
//   - error: nil on success, error on failure (e.g., insufficient balance)
func (c *Client) PreviewPayout(req *developer.PayoutPreviewRequest) (*developer.PayoutPreviewResp, error) {
	var response developer.PayoutPreviewResp
	err := c.sendRequest("POST", "/v1/payouts/preview", req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// CreatePayout initiates a new payout to transfer funds to a bank account.
// Payouts may require approval depending on merchant settings and may take several business days.
//
// Parameters:
//   - req: Request containing payout amount, currency, destination account, and optional metadata
//
// Returns:
//   - *developer.PayoutCreateResp: The created payout with ID, status, and estimated arrival date
//   - error: nil on success, error on failure (e.g., insufficient balance or invalid account)
func (c *Client) CreatePayout(req *developer.PayoutCreateRequest) (*developer.PayoutCreateResp, error) {
	var response developer.PayoutCreateResp
	err := c.sendRequest("POST", "/v1/payouts", req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// GetPayout retrieves detailed information about a specific payout.
// Includes payout status, amount, fees, destination account, and tracking information.
//
// Parameters:
//   - payoutID: The unique identifier of the payout
//
// Returns:
//   - *developer.PayoutGetResp: Complete payout details
//   - error: nil on success, error on failure (e.g., payout not found)
func (c *Client) GetPayout(payoutID string) (*developer.PayoutGetResp, error) {
	var response developer.PayoutGetResp
	err := c.sendRequest("GET", fmt.Sprintf("/v1/payouts/%s", payoutID), nil, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// ListPayouts retrieves a paginated list of payouts.
// Can be filtered by status, date range, destination, and other criteria.
//
// Parameters:
//   - req: Query parameters including pagination and filters (status, date range, destination, etc.)
//
// Returns:
//   - *developer.PayoutListResp: List of payouts with pagination metadata
//   - error: nil on success, error on failure
func (c *Client) ListPayouts(req *developer.PayoutListRequest) (*developer.PayoutListResp, error) {
	var response developer.PayoutListResp
	err := c.sendRequestWithQuery("GET", "/v1/payouts", req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// CancelPayout cancels a pending payout before it is processed.
// Only payouts in pending status can be cancelled. Once in transit, payouts cannot be cancelled.
//
// Parameters:
//   - payoutID: The unique identifier of the payout to cancel
//
// Returns:
//   - *developer.Payout: The cancelled payout with updated status
//   - error: nil on success, error on failure (e.g., payout already in transit)
func (c *Client) CancelPayout(payoutID string) (*developer.Payout, error) {
	var response developer.Payout
	err := c.sendRequest("POST", fmt.Sprintf("/v1/payouts/%s/cancel", payoutID), nil, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// GetApprovalInstance retrieves approval workflow details for a resource.
// Used to check the approval status and history for payouts or other resources requiring approval.
//
// Parameters:
//   - resourceID: The unique identifier of the resource (e.g., payout ID)
//
// Returns:
//   - *developer.ApprovalInstance: Approval instance details including status, approvers, and comments
//   - error: nil on success, error on failure
func (c *Client) GetApprovalInstance(resourceID string) (*developer.ApprovalInstance, error) {
	params := map[string]string{"resource_id": resourceID}
	var response developer.ApprovalInstance
	err := c.sendRequestWithQuery("GET", "/v1/approval/detail", params, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// ApprovePayout approves a payout that is pending approval.
// Once approved, the payout will proceed to processing and fund transfer.
//
// Parameters:
//   - approvalID: The unique identifier of the approval instance
//   - comment: Optional comment explaining the approval decision
//
// Returns:
//   - error: nil on success, error on failure (e.g., insufficient permissions or already approved)
func (c *Client) ApprovePayout(approvalID string, comment string) error {
	requestBody := map[string]string{"comment": comment}
	return c.sendRequest("POST", fmt.Sprintf("/v1/approval/%s/approve", approvalID), requestBody, nil)
}

// RejectPayout rejects a payout that is pending approval.
// Rejected payouts will be cancelled and funds will remain in the merchant balance.
//
// Parameters:
//   - approvalID: The unique identifier of the approval instance
//   - reason: Required comment explaining the rejection reason
//
// Returns:
//   - error: nil on success, error on failure (e.g., insufficient permissions or already processed)
func (c *Client) RejectPayout(approvalID string, reason string) error {
	requestBody := map[string]string{"comment": reason}
	return c.sendRequest("POST", fmt.Sprintf("/v1/approval/%s/reject", approvalID), requestBody, nil)
}

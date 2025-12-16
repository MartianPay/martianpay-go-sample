// Package martianpay provides SDK methods for managing approval workflows.
// Approvals allow merchants to implement multi-level authorization for sensitive operations like payouts.
package martianpay

import (
	"fmt"

	"github.com/MartianPay/martianpay-go-sample/pkg/developer"
)

// GetApprovalDetail retrieves detailed information about an approval instance.
// Used to check approval status, approvers, and approval history for a resource.
//
// Parameters:
//   - params: Request containing resource ID or approval ID to query
//
// Returns:
//   - *developer.ApprovalInstance: Approval details including status, approvers, comments, and timeline
//   - error: nil on success, error on failure
func (c *Client) GetApprovalDetail(params *developer.ApprovalGetRequest) (*developer.ApprovalInstance, error) {
	var resp developer.ApprovalInstance
	err := c.sendRequestWithQuery("GET", "/v1/approval/detail", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ApproveApproval approves a pending approval request.
// Once approved, the associated resource (e.g., payout) will proceed to the next stage or execution.
//
// Parameters:
//   - approvalID: The unique identifier of the approval instance to approve
//
// Returns:
//   - *developer.ApprovalInstance: The updated approval instance with approved status
//   - error: nil on success, error on failure (e.g., insufficient permissions or already processed)
func (c *Client) ApproveApproval(approvalID string) (*developer.ApprovalInstance, error) {
	path := fmt.Sprintf("/v1/approval/%s/approve", approvalID)
	var resp developer.ApprovalInstance
	err := c.sendRequest("POST", path, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// RejectApproval rejects a pending approval request.
// Rejected approvals will cancel the associated resource operation.
//
// Parameters:
//   - approvalID: The unique identifier of the approval instance to reject
//
// Returns:
//   - *developer.ApprovalInstance: The updated approval instance with rejected status
//   - error: nil on success, error on failure (e.g., insufficient permissions or already processed)
func (c *Client) RejectApproval(approvalID string) (*developer.ApprovalInstance, error) {
	path := fmt.Sprintf("/v1/approval/%s/reject", approvalID)
	var resp developer.ApprovalInstance
	err := c.sendRequest("POST", path, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

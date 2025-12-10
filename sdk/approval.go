package martianpay

import (
	"fmt"

	"github.com/MartianPay/martianpay-go-sample/pkg/developer"
)

// GetApprovalDetail gets approval details
func (c *Client) GetApprovalDetail(params *developer.ApprovalGetRequest) (*developer.ApprovalInstance, error) {
	var resp developer.ApprovalInstance
	err := c.sendRequestWithQuery("GET", "/v1/approval/detail", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ApproveApproval approves approval
func (c *Client) ApproveApproval(approvalID string) (*developer.ApprovalInstance, error) {
	path := fmt.Sprintf("/v1/approval/%s/approve", approvalID)
	var resp developer.ApprovalInstance
	err := c.sendRequest("POST", path, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// RejectApproval rejects approval
func (c *Client) RejectApproval(approvalID string) (*developer.ApprovalInstance, error) {
	path := fmt.Sprintf("/v1/approval/%s/reject", approvalID)
	var resp developer.ApprovalInstance
	err := c.sendRequest("POST", path, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

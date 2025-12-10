package martianpay

import (
	"fmt"

	"github.com/MartianPay/martianpay-go-sample/pkg/developer"
)

// PreviewPayout previews a payout before creation
func (c *Client) PreviewPayout(req *developer.PayoutPreviewRequest) (*developer.PayoutPreviewResp, error) {
	var response developer.PayoutPreviewResp
	err := c.sendRequest("POST", "/v1/payouts/preview", req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// CreatePayout creates a new payout
func (c *Client) CreatePayout(req *developer.PayoutCreateRequest) (*developer.PayoutCreateResp, error) {
	var response developer.PayoutCreateResp
	err := c.sendRequest("POST", "/v1/payouts", req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// GetPayout retrieves payout details
func (c *Client) GetPayout(payoutID string) (*developer.PayoutGetResp, error) {
	var response developer.PayoutGetResp
	err := c.sendRequest("GET", fmt.Sprintf("/v1/payouts/%s", payoutID), nil, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// ListPayouts lists all payouts with pagination
func (c *Client) ListPayouts(req *developer.PayoutListRequest) (*developer.PayoutListResp, error) {
	var response developer.PayoutListResp
	err := c.sendRequestWithQuery("GET", "/v1/payouts", req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// CancelPayout cancels a pending payout
func (c *Client) CancelPayout(payoutID string) (*developer.Payout, error) {
	var response developer.Payout
	err := c.sendRequest("POST", fmt.Sprintf("/v1/payouts/%s/cancel", payoutID), nil, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// GetApprovalInstance retrieves approval instance details
func (c *Client) GetApprovalInstance(resourceID string) (*developer.ApprovalInstance, error) {
	params := map[string]string{"resource_id": resourceID}
	var response developer.ApprovalInstance
	err := c.sendRequestWithQuery("GET", "/v1/approval/detail", params, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// ApprovePayout approves a payout
func (c *Client) ApprovePayout(approvalID string, comment string) error {
	requestBody := map[string]string{"comment": comment}
	return c.sendRequest("POST", fmt.Sprintf("/v1/approval/%s/approve", approvalID), requestBody, nil)
}

// RejectPayout rejects a payout
func (c *Client) RejectPayout(approvalID string, reason string) error {
	requestBody := map[string]string{"comment": reason}
	return c.sendRequest("POST", fmt.Sprintf("/v1/approval/%s/reject", approvalID), requestBody, nil)
}

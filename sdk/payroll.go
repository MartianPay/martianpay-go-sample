package martianpay

import (
	"fmt"

	"github.com/MartianPay/martianpay-go-sample/pkg/developer"
)

// CreateDirectPayroll creates a direct payroll
func (c *Client) CreateDirectPayroll(req *developer.PayrollDirectCreateRequest) (*developer.PayrollDirectCreateResponse, error) {
	var response developer.PayrollDirectCreateResponse
	err := c.sendRequest("POST", "/v1/payrolls/direct", req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// ConfirmPayroll confirms a payroll
func (c *Client) ConfirmPayroll(payrollID string, req *developer.PayrollConfirmRequest) (*developer.PayrollConfirmResponse, error) {
	var response developer.PayrollConfirmResponse
	err := c.sendRequest("POST", fmt.Sprintf("/v1/payrolls/%s/confirm", payrollID), req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// GetPayroll retrieves payroll details
func (c *Client) GetPayroll(payrollID string) (*developer.PayrollGetResponse, error) {
	var response developer.PayrollGetResponse
	err := c.sendRequest("GET", fmt.Sprintf("/v1/payrolls/%s", payrollID), nil, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// ListPayrolls lists all payrolls with pagination
func (c *Client) ListPayrolls(req *developer.PayrollListRequest) (*developer.PayrollListResponse, error) {
	var response developer.PayrollListResponse
	err := c.sendRequestWithQuery("GET", "/v1/payrolls", req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// ListPayrollItems lists payroll items with pagination
func (c *Client) ListPayrollItems(req *developer.PayrollItemsListRequest) (*developer.PayrollItemsListResponse, error) {
	var response developer.PayrollItemsListResponse
	err := c.sendRequestWithQuery("GET", "/v1/payrolls/items/list", req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

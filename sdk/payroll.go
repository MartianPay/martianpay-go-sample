// Package martianpay provides SDK methods for managing payroll operations.
// Payroll enables merchants to process bulk payments to employees or contractors.
package martianpay

import (
	"fmt"

	"github.com/MartianPay/martianpay-go-sample/pkg/developer"
)

// CreateDirectPayroll creates a new direct payroll batch.
// Direct payroll allows merchants to pay multiple recipients in a single batch transaction.
//
// Parameters:
//   - req: Request containing payroll details, recipient list, and payment amounts
//
// Returns:
//   - *developer.PayrollDirectCreateResponse: The created payroll batch with ID and status
//   - error: nil on success, error on failure (e.g., insufficient balance or invalid recipients)
func (c *Client) CreateDirectPayroll(req *developer.PayrollDirectCreateRequest) (*developer.PayrollDirectCreateResponse, error) {
	var response developer.PayrollDirectCreateResponse
	err := c.sendRequest("POST", "/v1/payrolls/direct", req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// ConfirmPayroll confirms and executes a pending payroll batch.
// Once confirmed, the payroll will be processed and payments will be sent to recipients.
//
// Parameters:
//   - payrollID: The unique identifier of the payroll batch to confirm
//   - req: Confirmation request containing any final adjustments or approvals
//
// Returns:
//   - *developer.PayrollConfirmResponse: The confirmed payroll with updated status
//   - error: nil on success, error on failure (e.g., insufficient balance or validation errors)
func (c *Client) ConfirmPayroll(payrollID string, req *developer.PayrollConfirmRequest) (*developer.PayrollConfirmResponse, error) {
	var response developer.PayrollConfirmResponse
	err := c.sendRequest("POST", fmt.Sprintf("/v1/payrolls/%s/confirm", payrollID), req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// GetPayroll retrieves detailed information about a specific payroll batch.
// Includes batch status, total amount, recipient count, and payment progress.
//
// Parameters:
//   - payrollID: The unique identifier of the payroll batch
//
// Returns:
//   - *developer.PayrollGetResponse: Complete payroll batch details
//   - error: nil on success, error on failure (e.g., payroll not found)
func (c *Client) GetPayroll(payrollID string) (*developer.PayrollGetResponse, error) {
	var response developer.PayrollGetResponse
	err := c.sendRequest("GET", fmt.Sprintf("/v1/payrolls/%s", payrollID), nil, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// ListPayrolls retrieves a paginated list of payroll batches.
// Can be filtered by status, date range, and other criteria.
//
// Parameters:
//   - req: Query parameters including pagination and filters (status, date range, etc.)
//
// Returns:
//   - *developer.PayrollListResponse: List of payroll batches with pagination metadata
//   - error: nil on success, error on failure
func (c *Client) ListPayrolls(req *developer.PayrollListRequest) (*developer.PayrollListResponse, error) {
	var response developer.PayrollListResponse
	err := c.sendRequestWithQuery("GET", "/v1/payrolls", req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// ListPayrollItems retrieves individual payment items within payroll batches.
// Shows detailed status and tracking for each recipient payment.
//
// Parameters:
//   - req: Query parameters including pagination and filters (payroll ID, recipient, status, etc.)
//
// Returns:
//   - *developer.PayrollItemsListResponse: List of individual payroll payment items with pagination metadata
//   - error: nil on success, error on failure
func (c *Client) ListPayrollItems(req *developer.PayrollItemsListRequest) (*developer.PayrollItemsListResponse, error) {
	var response developer.PayrollItemsListResponse
	err := c.sendRequestWithQuery("GET", "/v1/payrolls/items/list", req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

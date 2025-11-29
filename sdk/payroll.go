package martianpay

import (
	"fmt"

	"github.com/MartianPay/martianpay-go-sample/pkg/developer"
)

// PayrollDirectCreateRequest represents a request to create a payroll directly without CSV upload
type PayrollDirectCreateRequest struct {
	// ExternalID is an optional external identifier for idempotency and tracking
	ExternalID string `json:"external_id,omitempty"`
	// Items contains the list of payroll recipients and their payment details
	Items []PayrollDirectItem `json:"items" binding:"required,min=1"`
	// AutoApprove determines whether to automatically approve and execute the payroll without manual review
	AutoApprove bool `json:"auto_approve,omitempty"`
}

// PayrollDirectItem represents a single payroll recipient in direct payroll creation
type PayrollDirectItem struct {
	// ExternalID is an optional external identifier for this payroll item
	ExternalID string `json:"external_id,omitempty"`
	// Name is the recipient's name (optional, will use address suffix if empty)
	Name string `json:"name,omitempty"`
	// Email is the recipient's email address for notifications
	Email string `json:"email,omitempty"`
	// Phone is the recipient's phone number for notifications
	Phone string `json:"phone,omitempty"`
	// Coin is the cryptocurrency symbol (e.g., "USDC", "USDT", "ETH")
	Coin string `json:"coin" binding:"required"`
	// Network is the blockchain network (e.g., "ETH", "TRON", "SOL")
	Network string `json:"network" binding:"required"`
	// Address is the recipient's wallet address for payment
	Address string `json:"address" binding:"required"`
	// Amount is the payment amount in the specified currency
	Amount string `json:"amount" binding:"required"`
	// PaymentMethod is the payment method ("normal" or "binance", defaults to "normal")
	PaymentMethod string `json:"payment_method,omitempty"`
}

// PayrollDirectCreateResponse represents the response after creating a direct payroll
type PayrollDirectCreateResponse struct {
	// Payroll contains the created payroll header information including status and totals
	Payroll *developer.Payroll `json:"payroll"`
	// Items contains the individual payroll items/employees with their payment details
	Items []*developer.PayrollItems `json:"items"`
	// SwapItems contains currency swap information if automatic conversion is enabled
	SwapItems []*developer.PayrollSwapItems `json:"swap_items"`
	// BinanceFromItems contains Binance exchange information if Binance swap is used
	BinanceFromItems []*developer.BinanceFromItems `json:"binance_from_items"`
}

// CreateDirectPayroll creates a payroll with optional auto-approval in a single API call
func (c *Client) CreateDirectPayroll(req PayrollDirectCreateRequest) (*PayrollDirectCreateResponse, error) {
	var response PayrollDirectCreateResponse
	err := c.sendRequest("POST", "/v1/payrolls/direct", req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// PayrollGetReq represents a request to get a specific payroll
type PayrollGetReq struct {
	ID string
}

// PayrollGetResponse represents the response when retrieving a specific payroll
type PayrollGetResponse struct {
	Payroll   *developer.Payroll            `json:"payroll"`
	Items     []*developer.PayrollItems     `json:"items"`
	SwapItems []*developer.PayrollSwapItems `json:"swap_items"`
}

// GetPayroll retrieves a specific payroll by ID
func (c *Client) GetPayroll(req PayrollGetReq) (*PayrollGetResponse, error) {
	var response PayrollGetResponse
	err := c.sendRequest("GET", fmt.Sprintf("/v1/payrolls/%s", req.ID), nil, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// PayrollListReq represents a request to list payrolls
type PayrollListReq struct {
	// Pagination
	Page     int32 `json:"page" binding:"min=0"`               // Page number, starting from 0
	PageSize int32 `json:"page_size" binding:"required,min=1"` // Items per page

	// Filters
	StartDate  *string `json:"start_date,omitempty"`  // Filter by start date (YYYY-MM-DD)
	EndDate    *string `json:"end_date,omitempty"`    // Filter by end date (YYYY-MM-DD)
	ExternalID *string `json:"external_id,omitempty"` // Optional filter by external ID
	PayrollID  *string `json:"payroll_id,omitempty"`  // Optional filter by payroll ID
	Status     *string `json:"status,omitempty"`      // Optional filter by status
}

// PayrollListResponse represents the response when listing payrolls
type PayrollListResponse struct {
	Payrolls []*developer.Payroll `json:"payrolls"`
	Total    int64                `json:"total"`
}

// ListPayrolls retrieves a list of payrolls based on the provided parameters
func (c *Client) ListPayrolls(req PayrollListReq) (*PayrollListResponse, error) {
	var response PayrollListResponse
	err := c.sendRequestWithQuery("GET", "/v1/payrolls", req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// PayrollItemsListReq represents a request to list payroll items
type PayrollItemsListReq struct {
	// Pagination
	Page     int32 `json:"page" binding:"min=0"`               // Page number, starting from 0
	PageSize int32 `json:"page_size" binding:"required,min=1"` // Items per page

	// Filters
	StartDate      *string `json:"start_date,omitempty"`
	EndDate        *string `json:"end_date,omitempty"`
	EmployeeName   *string `json:"employee_name,omitempty"`
	ExternalID     *string `json:"external_id,omitempty"`      // Optional filter by payroll external ID
	PayrollID      *string `json:"payroll_id,omitempty"`       // Optional filter by payroll ID
	ItemExternalID *string `json:"item_external_id,omitempty"` // Optional filter by payroll item external ID
}

// PayrollItemsListResponse represents the response when listing payroll items
type PayrollItemsListResponse struct {
	PayrollItems []*developer.PayrollItems `json:"payroll_items"`
	Total        int64                     `json:"total"`
}

// ListPayrollItems retrieves a list of payroll items based on the provided parameters
func (c *Client) ListPayrollItems(req PayrollItemsListReq) (*PayrollItemsListResponse, error) {
	var response PayrollItemsListResponse
	err := c.sendRequestWithQuery("GET", "/v1/payrolls/items/list", req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

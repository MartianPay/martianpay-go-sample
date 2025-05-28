package martianpay

import (
	"fmt"

	"github.com/MartianPay/martianpay-go-sample/pkg/developer"
)

type CustomerCreateRequest struct {
	developer.CustomerParams
}

type CustomerCreateResp struct {
	developer.Customer
}

// CreateCustomer creates a new customer
func (c *Client) CreateCustomer(req CustomerCreateRequest) (*CustomerCreateResp, error) {
	var response CustomerCreateResp
	err := c.sendRequest("POST", "/v1/customers", req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

type CustomerUpdateRequest struct {
	ID string
	developer.CustomerParams
}

type CustomerUpdateResp struct {
	developer.Customer
}

// UpdateCustomer updates an existing customer
func (c *Client) UpdateCustomer(req CustomerUpdateRequest) (*CustomerUpdateResp, error) {
	var response CustomerUpdateResp
	err := c.sendRequest("POST", fmt.Sprintf("/v1/customers/%s", req.ID), req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

type CustomerGetRequest struct {
	ID string
}

type CustomerGetResp struct {
	developer.Customer
}

type CustomerGetHandler struct {
}

// GetCustomer retrieves a specific customer by ID
func (c *Client) GetCustomer(req CustomerGetRequest) (*CustomerGetResp, error) {
	var response CustomerGetResp
	err := c.sendRequest("GET", fmt.Sprintf("/v1/customers/%s", req.ID), nil, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

type CustomerListRequest struct {
	// Pagination
	Page     int32   `json:"page" binding:"min=0"`               // Page number, starting from 0
	PageSize int32   `json:"page_size" binding:"required,min=1"` // Items per page
	Email    *string `json:"email,omitempty"`                    // Filter by email
}

type CustomerListResp struct {
	Customers []developer.Customer `json:"customers"` // List of customers
	Total     int32                `json:"total"`     // Total number of records matching the filters
	Page      int32                `json:"page"`      // Current page number
	PageSize  int32                `json:"page_size"` // Items per page
}

// ListCustomers retrieves a list of customers based on the provided parameters
func (c *Client) ListCustomers(req CustomerListRequest) (*CustomerListResp, error) {
	var response CustomerListResp
	err := c.sendRequest("POST", "/v1/customers/list", req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

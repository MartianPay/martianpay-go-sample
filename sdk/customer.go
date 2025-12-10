package martianpay

import (
	"fmt"

	"github.com/MartianPay/martianpay-go-sample/pkg/developer"
)

// CreateCustomer creates a new customer
func (c *Client) CreateCustomer(req *developer.CustomerCreateRequest) (*developer.Customer, error) {
	var response developer.Customer
	err := c.sendRequest("POST", "/v1/customers", req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// UpdateCustomer updates an existing customer
func (c *Client) UpdateCustomer(customerID string, req *developer.CustomerUpdateRequest) (*developer.Customer, error) {
	var response developer.Customer
	err := c.sendRequest("POST", fmt.Sprintf("/v1/customers/%s", customerID), req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// GetCustomer retrieves a specific customer by ID
func (c *Client) GetCustomer(customerID string) (*developer.Customer, error) {
	var response developer.Customer
	err := c.sendRequest("GET", fmt.Sprintf("/v1/customers/%s", customerID), nil, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// ListCustomers retrieves a list of customers based on the provided parameters
func (c *Client) ListCustomers(req *developer.CustomerListRequest) (*developer.CustomerListResponse, error) {
	var response developer.CustomerListResponse
	err := c.sendRequestWithQuery("GET", "/v1/customers", req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// DeleteCustomer deletes a customer by ID
func (c *Client) DeleteCustomer(customerID string) error {
	err := c.sendRequest("DELETE", fmt.Sprintf("/v1/customers/%s", customerID), nil, nil)
	if err != nil {
		return err
	}
	return nil
}

// ListCustomerPaymentMethods retrieves a list of saved payment methods for a customer
func (c *Client) ListCustomerPaymentMethods(customerID string) (*developer.PaymentMethodListResponse, error) {
	req := &developer.CustomerPaymentMethodListRequest{
		CustomerID: customerID,
	}
	var response developer.PaymentMethodListResponse
	err := c.sendRequestWithQuery("GET", "/v1/customers/payment_methods", req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

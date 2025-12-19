// Package martianpay provides SDK methods for managing customers.
// Customers represent individuals or businesses that make purchases from merchants.
package martianpay

import (
	"fmt"

	"github.com/MartianPay/martianpay-go-sample/pkg/developer"
)

// CreateCustomer creates a new customer record.
// Customers can be associated with payment methods, subscriptions, and orders.
//
// Parameters:
//   - req: Request containing customer details (name, email, phone, shipping address, metadata)
//
// Returns:
//   - *developer.Customer: The created customer with assigned ID
//   - error: nil on success, error on failure
func (c *Client) CreateCustomer(req *developer.CustomerCreateRequest) (*developer.Customer, error) {
	var response developer.Customer
	err := c.sendRequest("POST", "/v1/customers", req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// UpdateCustomer updates an existing customer's information.
// Can modify customer name, email, phone, address, and metadata.
//
// Parameters:
//   - customerID: The unique identifier of the customer to update
//   - req: Updated customer fields (name, email, phone, address, metadata)
//
// Returns:
//   - *developer.Customer: The updated customer details
//   - error: nil on success, error on failure
func (c *Client) UpdateCustomer(customerID string, req *developer.CustomerUpdateRequest) (*developer.Customer, error) {
	var response developer.Customer
	err := c.sendRequest("POST", fmt.Sprintf("/v1/customers/%s", customerID), req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// GetCustomer retrieves detailed information about a specific customer.
// Includes customer contact details, payment methods, and metadata.
//
// Parameters:
//   - customerID: The unique identifier of the customer
//
// Returns:
//   - *developer.Customer: Complete customer details
//   - error: nil on success, error on failure (e.g., customer not found)
func (c *Client) GetCustomer(customerID string) (*developer.Customer, error) {
	var response developer.Customer
	err := c.sendRequest("GET", fmt.Sprintf("/v1/customers/%s", customerID), nil, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// ListCustomers retrieves a paginated list of customers.
// Can be filtered by email, creation date, and other criteria.
//
// Parameters:
//   - req: Query parameters including pagination and filters (email, creation date, etc.)
//
// Returns:
//   - *developer.CustomerListResponse: List of customers with pagination metadata
//   - error: nil on success, error on failure
func (c *Client) ListCustomers(req *developer.CustomerListRequest) (*developer.CustomerListResponse, error) {
	var response developer.CustomerListResponse
	err := c.sendRequestWithQuery("GET", "/v1/customers", req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// DeleteCustomer permanently deletes a customer record.
// Deletion may be restricted if the customer has active subscriptions or pending payments.
//
// Parameters:
//   - customerID: The unique identifier of the customer to delete
//
// Returns:
//   - error: nil on success, error on failure (e.g., customer has active subscriptions)
func (c *Client) DeleteCustomer(customerID string) error {
	err := c.sendRequest("DELETE", fmt.Sprintf("/v1/customers/%s", customerID), nil, nil)
	if err != nil {
		return err
	}
	return nil
}

// GenerateEphemeralToken creates a short-lived token for customer authentication in checkout flows.
// Ephemeral tokens allow social media integrations and third-party systems to authenticate customers
// without exposing long-lived credentials.
//
// Parameters:
//   - req: Request containing identity provider info (idp_key, idp_subject), provider, return URL, etc.
//
// Returns:
//   - *developer.EphemeralTokenResponse: The ephemeral token with expiration time
//   - error: nil on success, error on failure
func (c *Client) GenerateEphemeralToken(req *developer.EphemeralTokenRequest) (*developer.EphemeralTokenResponse, error) {
	var response developer.EphemeralTokenResponse
	err := c.sendRequest("POST", "/v1/customers/ephemeral_token", req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// ListCustomerPaymentMethods retrieves all saved payment methods for a specific customer.
// Includes credit cards, bank accounts, and other payment instruments on file.
//
// Parameters:
//   - customerID: The unique identifier of the customer
//
// Returns:
//   - *developer.PaymentMethodListResponse: List of payment methods associated with the customer
//   - error: nil on success, error on failure
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

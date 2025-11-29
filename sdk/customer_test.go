package martianpay

import (
	"testing"

	"github.com/MartianPay/martianpay-go-sample/pkg/developer"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestCreateAndUpdateCustomer(t *testing.T) {
	// Initialize the client with your API key
	client := NewClient(apiKey)

	// Create a customer
	email := "newcustomer@example.com"
	name := "John Doe"
	description := "Test customer"
	createReq := CustomerCreateRequest{
		CustomerParams: developer.CustomerParams{
			Email:       &email,
			Name:        &name,
			Description: &description,
		},
	}

	createResp, err := client.CreateCustomer(createReq)
	assert.NoError(t, err)
	assert.NotNil(t, createResp)
	assert.Equal(t, createReq.Email, createResp.Email)
	assert.Equal(t, createReq.Name, createResp.Name)

	// Update the customer
	newName := "Jane Doe"
	newDescription := "Updated test customer"
	updateReq := CustomerUpdateRequest{
		ID: createResp.ID,
		CustomerParams: developer.CustomerParams{
			Name:        &newName,
			Description: &newDescription,
		},
	}

	updateResp, err := client.UpdateCustomer(updateReq)
	assert.NoError(t, err)
	assert.NotNil(t, updateResp)
	assert.Equal(t, updateReq.ID, updateResp.ID)
	assert.Equal(t, updateReq.Name, updateResp.Name)
	assert.Equal(t, updateReq.Description, updateResp.Description)

	logrus.WithFields(logrus.Fields{
		"id":    updateResp.ID,
		"email": updateResp.Email,
		"name":  updateResp.Name,
	}).Info("Successfully created and updated customer")
}

func TestGetCustomer(t *testing.T) {
	// Initialize the client with your API key
	client := NewClient(apiKey)

	// Create request
	req := CustomerGetRequest{
		ID: "cus_7Xpxk2n22WAEBDaBVlG5dnRf",
	}

	// Call API
	response, err := client.GetCustomer(req)

	// Verify response
	if assert.NoError(t, err) {
		assert.NotNil(t, response)
		assert.Equal(t, req.ID, response.ID)
	}
}

func TestListCustomers(t *testing.T) {
	// Initialize the client with your API key
	client := NewClient(apiKey)

	email := "user@example.com"

	// Create request
	req := CustomerListRequest{
		Page:     0,
		PageSize: 10,
		Email:    &email, // Optional
	}

	// Call API
	response, err := client.ListCustomers(req)

	// Verify response
	if assert.NoError(t, err) {
		assert.NotNil(t, response)
		assert.Equal(t, req.Page, response.Page)
		assert.Equal(t, req.PageSize, response.PageSize)

		// Log response for debugging
		logrus.WithFields(logrus.Fields{
			"total_records": response.Total,
			"page":          response.Page,
			"page_size":     response.PageSize,
		}).Info("Successfully retrieved customers")

		// Log individual customers
		for _, customer := range response.Customers {
			logrus.WithFields(logrus.Fields{
				"id":            customer.ID,
				"name":          customer.Name,
				"email":         customer.Email,
				"total_expense": customer.TotalExpense,
			}).Info("Customer details")
		}
	}
}

func TestDeleteCustomer(t *testing.T) {
	// Initialize the client with your API key
	client := NewClient(apiKey)

	// First, create a customer to delete
	email := "delete_test@example.com"
	name := "Delete Test Customer"
	createReq := CustomerCreateRequest{
		CustomerParams: developer.CustomerParams{
			Email: &email,
			Name:  &name,
		},
	}

	createResp, err := client.CreateCustomer(createReq)
	assert.NoError(t, err)
	assert.NotNil(t, createResp)

	// Now delete the customer
	deleteReq := CustomerDeleteRequest{
		ID: createResp.ID,
	}

	deleteResp, err := client.DeleteCustomer(deleteReq)
	assert.NoError(t, err)
	assert.NotNil(t, deleteResp)
	assert.True(t, deleteResp.Deleted)
	assert.Equal(t, createResp.ID, deleteResp.ID)

	logrus.WithFields(logrus.Fields{
		"id":      deleteResp.ID,
		"deleted": deleteResp.Deleted,
	}).Info("Successfully deleted customer")
}

func TestListCustomerPaymentMethods(t *testing.T) {
	// Initialize the client with your API key
	client := NewClient(apiKey)

	// Use an existing customer with saved cards
	customerID := "cus_7Xpxk2n22WAEBDaBVlG5dnRf" // Replace with actual customer ID that has saved cards

	// List customer's saved payment methods
	req := CustomerPaymentMethodListRequest{
		CustomerID: customerID,
	}

	resp, err := client.ListCustomerPaymentMethods(req)
	if err != nil {
		logrus.WithError(err).Warn("Failed to list payment methods")
		t.Skip("Skipping test - failed to list payment methods")
		return
	}

	assert.NotNil(t, resp)
	logrus.WithFields(logrus.Fields{
		"customer_id":     customerID,
		"payment_methods": len(resp.PaymentMethods),
	}).Info("Successfully listed customer payment methods")

	// Log each payment method
	for _, pm := range resp.PaymentMethods {
		logrus.WithFields(logrus.Fields{
			"payment_method_id": pm.ID,
			"provider":          pm.Provider,
			"type":              pm.Type,
			"brand":             pm.Brand,
			"last4":             pm.Last4,
			"exp_month":         pm.ExpMonth,
			"exp_year":          pm.ExpYear,
			"funding":           pm.Funding,
			"country":           pm.Country,
		}).Info("Payment method details")
	}
}

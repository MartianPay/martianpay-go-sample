package martianpay

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

const apiKey = "your_api_key_here"

func TestGetPaymentIntent(t *testing.T) {
	// Initialize the client with your API key
	client := NewClient(apiKey)

	// Create request
	req := PaymentIntentGetReq{
		ID: "pi_nKSxQrU2Pjh9KGzyIRJcWpGJ",
	}

	// Call API
	response, err := client.GetPaymentIntent(req)

	// Verify response
	if assert.NoError(t, err) {
		assert.NotNil(t, response)
		assert.Equal(t, req.ID, response.ID)
		logrus.WithFields(logrus.Fields{
			"id":     response.ID,
			"amount": response.Amount,
		}).Info("Successfully retrieved payment intent")
	}
}

func TestListPaymentIntents(t *testing.T) {
	// Initialize the client with your API key
	client := NewClient(apiKey)

	email := "user@example.com"

	// Create request
	req := PaymentIntentListReq{
		Page:          0,
		PageSize:      10,
		CustomerEmail: &email, // Optional
	}

	// Call API
	response, err := client.ListPaymentIntents(req)

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
		}).Info("Successfully retrieved payment intents")

		// Log individual payment intents
		for _, intent := range response.PaymentIntents {
			logrus.WithFields(logrus.Fields{
				"id":     intent.ID,
				"amount": intent.Amount,
			}).Info("Payment intent details")
		}
	}
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

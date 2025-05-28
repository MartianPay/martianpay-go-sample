package martianpay

import (
	"testing"

	"github.com/MartianPay/martianpay-go-sample/pkg/developer"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

const apiKey = "your_api_key_here"

func TestCreateAndUpdatePaymentIntent(t *testing.T) {
	// Initialize the client with your API key
	client := NewClient(apiKey)

	// Create a payment intent
	description := "Test payment intent"
	createReq := PaymentIntentCreateRequest{
		PaymentIntentParams: developer.PaymentIntentParams{
			Amount:          "10.00",
			Currency:        "USD",
			MerchantOrderId: "order_123",
			Description:     &description,
		},
	}

	createResp, err := client.CreatePaymentIntent(createReq)
	assert.NoError(t, err)
	assert.NotNil(t, createResp)

	// Update the payment intent
	paymentMethodType := developer.PaymentMethodTypeCrypto
	assetId := "USDC-Ethereum-TEST"
	updateReq := PaymentIntentUpdateRequest{
		ID:                createResp.ID,
		PaymentMethodType: &paymentMethodType,
		PaymentMethodData: &developer.PaymentMethodOptions{
			Crypto: &developer.Crypto{
				AssetId: &assetId,
			},
		},
	}

	updateResp, err := client.UpdatePaymentIntent(updateReq)
	assert.NoError(t, err)
	assert.NotNil(t, updateResp)
	logrus.WithFields(logrus.Fields{
		"id":              updateResp.ID,
		"payment_address": *updateResp.Charges[0].PaymentMethodOptions.Crypto.DepositAddress,
	}).Info("Updated payment intent with payment address")
}

func TestGetPaymentIntent(t *testing.T) {
	// Initialize the client with your API key
	client := NewClient(apiKey)

	// Create request
	req := PaymentIntentGetReq{
		ID: "pi_bLwFdNH862WcOwryK0ThCpHL",
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

func TestCreateRefund(t *testing.T) {
	// Initialize the client with your API key
	client := NewClient(apiKey)

	// Create request
	reason := string(developer.PaymentIntentCancellationReasonRequestedByCustomer)
	paymentIntentId := "pi_m3cL1pixO1D3c8A5iC18iGZo"
	req := RefundCreateRequest{
		RefundParams: developer.RefundParams{
			PaymentIntent: &paymentIntentId,
			Amount:        "1",
			Reason:        &reason,
		},
	}

	// Call API
	response, err := client.CreateRefund(req)

	// Verify response
	if assert.NoError(t, err) {
		assert.NotNil(t, response)
		assert.NotEmpty(t, response.Refunds)
		logrus.WithFields(logrus.Fields{
			"refund_id":      response.Refunds[0].ID,
			"amount":         response.Refunds[0].Amount,
			"payment_intent": response.Refunds[0].PaymentIntent,
		}).Info("Successfully created refund")
	}
}

func TestGetRefund(t *testing.T) {
	// Initialize the client with your API key
	client := NewClient(apiKey)

	// Create request
	req := RefundGetRequest{
		ID: "rf_ZTqX7HuMtGxaJ1KnizURLU9m", // Replace with a valid refund ID
	}

	// Call API
	response, err := client.GetRefund(req)

	// Verify response
	if assert.NoError(t, err) {
		assert.NotNil(t, response)
		assert.Equal(t, req.ID, response.ID)
		logrus.WithFields(logrus.Fields{
			"id":     response.ID,
			"amount": response.Amount,
			"status": response.Status,
		}).Info("Successfully retrieved refund")
	}
}

func TestListRefunds(t *testing.T) {
	// Initialize the client with your API key
	client := NewClient(apiKey)

	// Create request
	req := RefundListRequest{
		Page:     0,
		PageSize: 10,
	}

	// Call API
	response, err := client.ListRefunds(req)

	// Verify response
	if assert.NoError(t, err) {
		assert.NotNil(t, response)
		assert.Equal(t, req.Page, response.Page)
		assert.Equal(t, req.PageSize, response.PageSize)

		logrus.WithFields(logrus.Fields{
			"total_records": response.Total,
			"page":          response.Page,
			"page_size":     response.PageSize,
		}).Info("Successfully retrieved refunds")

		for _, refund := range response.Refunds {
			logrus.WithFields(logrus.Fields{
				"id":             refund.ID,
				"amount":         refund.Amount,
				"status":         refund.Status,
				"payment_intent": refund.PaymentIntent,
			}).Info("Refund details")
		}
	}
}

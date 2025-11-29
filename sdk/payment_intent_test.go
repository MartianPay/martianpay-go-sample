package martianpay

import (
	"testing"

	"github.com/MartianPay/martianpay-go-sample/pkg/developer"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestCreateAndUpdatePaymentIntent(t *testing.T) {
	// This test demonstrates the API-only integration approach for crypto payments.
	//
	// INTEGRATION OPTIONS:
	// --------------------
	// Option 1 (Recommended): MartianPay.js Widget
	// - Use MartianPay.js widget on your frontend (https://docs.martianpay.com/v1/docs/martianpay-js-usage)
	// - After creating payment intent, display the widget with payment_intent.client_secret
	// - The widget handles payment method selection and automatically calls UpdatePaymentIntent API
	// - NO NEED to call UpdatePaymentIntent from your backend
	//
	// Option 2 (API-only): Direct Backend Integration (this test)
	// - Create payment intent via API
	// - Call UpdatePaymentIntent with crypto asset_id
	// - API returns deposit address
	// - Display deposit address to user for payment
	// - Listen to webhooks for payment confirmation

	// Initialize the client with your API key
	client := NewClient(apiKey)

	// Step 1: Create a payment intent
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

	// Step 2: Update payment intent to specify crypto payment method
	// NOTE: Skip this step if using MartianPay.js widget (widget handles it automatically)
	paymentMethodType := developer.PaymentMethodTypeCrypto
	assetId := "USDC-Ethereum-TEST"
	updateReq := PaymentIntentUpdateRequest{
		ID:                createResp.ID,
		PaymentMethodType: &paymentMethodType,
		PaymentMethodData: &developer.PaymentMethodConfirmOptions{
			Crypto: &developer.CryptoOption{
				AssetId: &assetId,
			},
		},
	}

	updateResp, err := client.UpdatePaymentIntent(updateReq)
	assert.NoError(t, err)
	assert.NotNil(t, updateResp)

	// Step 3: Extract deposit address from response
	// The API returns a unique deposit address for this payment
	// User sends crypto to this address to complete the payment
	logrus.WithFields(logrus.Fields{
		"payment_intent_id": updateResp.ID,
		"deposit_address":   *updateResp.Charges[0].PaymentMethodOptions.Crypto.DepositAddress,
		"asset_id":          *updateResp.Charges[0].PaymentMethodOptions.Crypto.AssetId,
		"network":           *updateResp.Charges[0].PaymentMethodOptions.Crypto.Network,
		"amount":            *updateResp.Charges[0].PaymentMethodOptions.Crypto.Amount,
	}).Info("Crypto payment details - display deposit address to user")
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

func TestCancelPaymentIntent(t *testing.T) {
	// Initialize the client with your API key
	client := NewClient(apiKey)

	// First, create a payment intent to cancel
	description := "Test payment intent for cancellation"
	createReq := PaymentIntentCreateRequest{
		PaymentIntentParams: developer.PaymentIntentParams{
			Amount:          "10.00",
			Currency:        "USD",
			MerchantOrderId: "order_cancel_test",
			Description:     &description,
		},
	}

	createResp, err := client.CreatePaymentIntent(createReq)
	assert.NoError(t, err)
	assert.NotNil(t, createResp)

	// Now cancel the payment intent
	reason := "Customer requested cancellation"
	cancelReq := PaymentIntentCancelReq{
		ID:     createResp.ID,
		Reason: &reason,
	}

	cancelResp, err := client.CancelPaymentIntent(cancelReq)
	assert.NoError(t, err)
	assert.NotNil(t, cancelResp)
	assert.Equal(t, "Cancelled", string(cancelResp.PaymentIntentStatus))

	logrus.WithFields(logrus.Fields{
		"id":                    cancelResp.ID,
		"status":                cancelResp.Status,
		"payment_intent_status": cancelResp.PaymentIntentStatus,
		"canceled_at":           cancelResp.CanceledAt,
	}).Info("Successfully canceled payment intent")
}

func TestListPaymentIntentsWithPermanentDeposit(t *testing.T) {
	// Initialize the client with your API key
	client := NewClient(apiKey)

	permanentDeposit := true
	assetId := "USDC-Ethereum-TEST"

	// Create request with permanent deposit filters
	req := PaymentIntentListReq{
		Page:                    0,
		PageSize:                10,
		PermanentDeposit:        &permanentDeposit,
		PermanentDepositAssetId: &assetId,
	}

	// Call API
	response, err := client.ListPaymentIntents(req)

	// Verify response
	if assert.NoError(t, err) {
		assert.NotNil(t, response)
		assert.Equal(t, req.Page, response.Page)
		assert.Equal(t, req.PageSize, response.PageSize)

		logrus.WithFields(logrus.Fields{
			"total_records": response.Total,
			"page":          response.Page,
			"page_size":     response.PageSize,
		}).Info("Successfully retrieved payment intents with permanent deposit filter")

		// Log individual payment intents
		for _, intent := range response.PaymentIntents {
			logrus.WithFields(logrus.Fields{
				"id":                         intent.ID,
				"amount":                     intent.Amount,
				"permanent_deposit":          intent.PermanentDeposit,
				"permanent_deposit_asset_id": intent.PermanentDepositAssetId,
			}).Info("Payment intent details")
		}
	}
}

func TestFiatPaymentWithNewCard(t *testing.T) {
	// This test demonstrates two integration approaches for card payments:
	//
	// APPROACH 1 (Recommended): MartianPay.js Integration
	// -------------------------------------------------------
	// After creating a payment intent, use MartianPay.js on the frontend:
	// - Documentation: https://docs.martianpay.com/v1/docs/martianpay-js-usage
	// - The frontend widget handles payment method selection (cards or crypto)
	// - The widget automatically calls the update payment intent API
	// - No need to manually call UpdatePaymentIntent from backend
	//
	// APPROACH 2 (API-only): Direct API Integration (demonstrated in this test)
	// -------------------------------------------------------------------------
	// If you don't use MartianPay.js, you need to handle everything via API:
	//
	// For CRYPTO payments:
	// - Call UpdatePaymentIntent with crypto asset_id
	// - API returns deposit address for the selected crypto
	// - Display the address to user for payment
	//
	// For CARD payments (this test):
	// - Call UpdatePaymentIntent with cards payment method
	// - API returns stripe_payload (client_secret, public_key)
	// - Use Stripe.js on frontend to complete payment with the client_secret
	// - Integration guide: https://stripe.com/docs/payments/accept-a-payment
	//
	// Stripe Integration Flow:
	// 1. Backend: Create payment intent via MartianPay API
	// 2. Backend: Call UpdatePaymentIntent with cards payment method
	// 3. Backend: Get stripe_payload from response (client_secret, public_key)
	// 4. Frontend: Initialize Stripe.js with public_key
	// 5. Frontend: Use Stripe Elements to collect card details
	// 6. Frontend: Call stripe.confirmCardPayment(client_secret) with card details
	// 7. Stripe: Processes the payment and returns result
	// 8. MartianPay: Receives webhook from Stripe and updates payment intent status

	// Initialize the client with your API key
	client := NewClient(apiKey)

	// Step 1: Query or create a customer
	email := "fiat_test@example.com"
	name := "Fiat Test Customer"

	// Try to find existing customer by email
	listCustomerReq := CustomerListRequest{
		Page:     0,
		PageSize: 10,
		Email:    &email,
	}

	listResp, err := client.ListCustomers(listCustomerReq)
	assert.NoError(t, err)

	var customerID string
	if listResp != nil && len(listResp.Customers) > 0 {
		// Customer exists, use the existing one
		customerID = listResp.Customers[0].ID
		logrus.WithFields(logrus.Fields{
			"customer_id": customerID,
			"email":       email,
		}).Info("Using existing customer for fiat payment test")
	} else {
		// Customer doesn't exist, create a new one
		createCustomerReq := CustomerCreateRequest{
			CustomerParams: developer.CustomerParams{
				Email: &email,
				Name:  &name,
			},
		}

		customerResp, err := client.CreateCustomer(createCustomerReq)
		assert.NoError(t, err)
		assert.NotNil(t, customerResp)
		customerID = customerResp.ID
		logrus.WithFields(logrus.Fields{
			"customer_id": customerID,
			"email":       *customerResp.Email,
		}).Info("Created new customer for fiat payment test")
	}

	// Step 2: Create a payment intent
	description := "Test fiat payment with new card"
	createPIReq := PaymentIntentCreateRequest{
		PaymentIntentParams: developer.PaymentIntentParams{
			Amount:          "25.00",
			Currency:        "USD",
			MerchantOrderId: "order_fiat_new_card",
			Description:     &description,
			Customer:        &customerID,
		},
	}

	piResp, err := client.CreatePaymentIntent(createPIReq)
	assert.NoError(t, err)
	assert.NotNil(t, piResp)
	logrus.WithFields(logrus.Fields{
		"payment_intent_id": piResp.ID,
		"amount":            piResp.Amount,
		"currency":          piResp.Currency,
	}).Info("Created payment intent for fiat payment")

	// Step 3: Update payment intent with cards payment method
	// This triggers MartianPay to prepare Stripe payment and return stripe_payload
	paymentMethodType := developer.PaymentMethodTypeCards
	currency := "USD"
	savePaymentMethod := true // Save the card for future use (creates reusable payment method)

	updatePIReq := PaymentIntentUpdateRequest{
		ID:                piResp.ID,
		PaymentMethodType: &paymentMethodType,
		PaymentMethodData: &developer.PaymentMethodConfirmOptions{
			Fiat: &developer.FiatOption{
				Currency:          &currency,
				SavePaymentMethod: &savePaymentMethod, // Save card for future payments
				// Note: PaymentMethodID is NOT provided for new card payments
				// It's only used when charging a saved card (see TestFiatPaymentWithSavedCard)
			},
		},
	}

	updateResp, err := client.UpdatePaymentIntent(updatePIReq)
	// The API returns stripe_payload containing:
	// - client_secret: Used by Stripe.js to confirm payment on frontend
	// - public_key: Stripe publishable key for initializing Stripe.js
	// - customer_id: Stripe customer ID (if SavePaymentMethod is true)
	if err != nil {
		logrus.WithError(err).Warn("Payment intent update failed (expected in test without Stripe setup)")
	} else {
		assert.NotNil(t, updateResp)
		logrus.WithFields(logrus.Fields{
			"payment_intent_id": updateResp.ID,
			"status":            updateResp.Status,
		}).Info("Updated payment intent with fiat payment method")

		// Step 4: Extract and display stripe_payload
		// These values should be sent to frontend for Stripe.js integration
		if len(updateResp.Charges) > 0 {
			for i, charge := range updateResp.Charges {
				if charge.StripePayload != nil {
					logrus.WithFields(logrus.Fields{
						"charge_index":     i,
						"charge_id":        charge.ID,
						"client_secret":    charge.StripePayload.ClientSecret, // Pass to frontend Stripe.js
						"public_key":       charge.StripePayload.PublicKey,    // Pass to frontend to initialize Stripe
						"status":           charge.StripePayload.Status,       // Current Stripe payment status
						"customer_id":      charge.StripePayload.CustomerID,   // Stripe customer ID (for saved cards)
						"payment_provider": charge.PaymentProvider,            // Should be "stripe"
					}).Info("Stripe payload details - use these values for frontend integration")

					// Frontend integration example (not executable in test):
					// const stripe = Stripe(public_key);
					// const {error} = await stripe.confirmCardPayment(client_secret, {
					//   payment_method: {
					//     card: cardElement,
					//     billing_details: { name: 'Customer Name' }
					//   }
					// });
				}
			}
		} else {
			logrus.Info("No charges found in response - stripe_payload will be in charges array")
		}
	}
}

func TestFiatPaymentWithSavedCard(t *testing.T) {
	// This test demonstrates charging a saved card (previously saved payment method).
	//
	// INTEGRATION OPTIONS:
	// --------------------
	// Option 1 (Recommended): MartianPay.js Widget
	// - Use MartianPay.js widget on your frontend (https://docs.martianpay.com/v1/docs/martianpay-js-usage)
	// - After creating payment intent, display the widget with payment_intent.client_secret
	// - The widget shows saved payment methods and handles payment automatically
	// - NO NEED to call UpdatePaymentIntent from your backend
	//
	// Option 2 (API-only): Direct Backend Integration (this test)
	// - List customer's saved payment methods
	// - Create payment intent
	// - Call UpdatePaymentIntent with saved payment_method_id
	// - Payment is processed immediately with the saved card

	// Initialize the client with your API key
	client := NewClient(apiKey)

	// Step 1: Use an existing customer with saved cards
	// In a real scenario, this customer would already have payment methods saved
	customerID := "cus_7Xpxk2n22WAEBDaBVlG5dnRf" // Replace with actual customer ID that has saved cards

	// Step 2: List customer's saved payment methods
	listPMReq := CustomerPaymentMethodListRequest{
		CustomerID: customerID,
	}

	pmListResp, err := client.ListCustomerPaymentMethods(listPMReq)
	if err != nil {
		logrus.WithError(err).Warn("Failed to list payment methods (customer may not have saved cards)")
		t.Skip("Skipping test - customer has no saved payment methods")
		return
	}

	assert.NotNil(t, pmListResp)
	if len(pmListResp.PaymentMethods) == 0 {
		logrus.Warn("No saved payment methods found for customer")
		t.Skip("Skipping test - no saved payment methods")
		return
	}

	// Log saved cards
	for _, pm := range pmListResp.PaymentMethods {
		logrus.WithFields(logrus.Fields{
			"payment_method_id": pm.ID,
			"brand":             pm.Brand,
			"last4":             pm.Last4,
			"exp_month":         pm.ExpMonth,
			"exp_year":          pm.ExpYear,
		}).Info("Saved payment method")
	}

	// Step 3: Create a payment intent
	description := "Test fiat payment with saved card"
	createPIReq := PaymentIntentCreateRequest{
		PaymentIntentParams: developer.PaymentIntentParams{
			Amount:          "15.00",
			Currency:        "USD",
			MerchantOrderId: "order_fiat_saved_card",
			Description:     &description,
			Customer:        &customerID,
		},
	}

	piResp, err := client.CreatePaymentIntent(createPIReq)
	assert.NoError(t, err)
	assert.NotNil(t, piResp)
	logrus.WithFields(logrus.Fields{
		"payment_intent_id": piResp.ID,
		"amount":            piResp.Amount,
	}).Info("Created payment intent for saved card payment")

	// Step 4: Update payment intent with saved card
	// NOTE: Skip this step if using MartianPay.js widget (widget handles it automatically)
	paymentMethodType := developer.PaymentMethodTypeCards
	currency := "USD"
	savedPaymentMethodID := pmListResp.PaymentMethods[0].ID // Use first saved card

	updatePIReq := PaymentIntentUpdateRequest{
		ID:                piResp.ID,
		PaymentMethodType: &paymentMethodType,
		PaymentMethodData: &developer.PaymentMethodConfirmOptions{
			Fiat: &developer.FiatOption{
				Currency:        &currency,
				PaymentMethodID: &savedPaymentMethodID, // Provide saved payment_method_id
				// SavePaymentMethod is not needed here since card is already saved
			},
		},
	}

	updateResp, err := client.UpdatePaymentIntent(updatePIReq)
	if err != nil {
		logrus.WithError(err).Warn("Payment intent update failed")
	} else {
		assert.NotNil(t, updateResp)
		logrus.WithFields(logrus.Fields{
			"payment_intent_id": updateResp.ID,
			"status":            updateResp.Status,
			"payment_method_id": savedPaymentMethodID,
		}).Info("Successfully charged saved card - payment processed immediately")
	}
}

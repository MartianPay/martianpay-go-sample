package main

import (
	"fmt"

	"github.com/MartianPay/martianpay-go-sample/pkg/developer"
	martianpay "github.com/MartianPay/martianpay-go-sample/sdk"
)

// Payment Intent Examples
func createAndUpdatePaymentIntent(client *martianpay.Client) {
	// This example demonstrates the API-only integration approach for crypto payments.
	//
	// INTEGRATION OPTIONS:
	// --------------------
	// Option 1 (Recommended): MartianPay.js Widget
	// - Use MartianPay.js widget on your frontend (https://docs.martianpay.com/v1/docs/martianpay-js-usage)
	// - After creating payment intent, display the widget with payment_intent.client_secret
	// - The widget handles payment method selection and automatically calls UpdatePaymentIntent API
	// - NO NEED to call UpdatePaymentIntent from your backend
	//
	// Option 2 (API-only): Direct Backend Integration (this example)
	// - Create payment intent via API
	// - Call UpdatePaymentIntent with crypto asset_id
	// - API returns deposit address
	// - Display deposit address to user for payment
	// - Listen to webhooks for payment confirmation

	fmt.Println("Creating and Updating Payment Intent (Crypto)...")

	description := "Test payment intent"
	createReq := martianpay.PaymentIntentCreateRequest{
		PaymentIntentParams: developer.PaymentIntentParams{
			Amount:          "10.00",
			Currency:        "USD",
			MerchantOrderId: "order_123",
			Description:     &description,
		},
	}

	createResp, err := client.CreatePaymentIntent(createReq)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("✓ Step 1 - Payment Intent Created:\n")
	fmt.Printf("  ID: %s\n", createResp.ID)
	fmt.Printf("  Amount: %s %s\n", createResp.Amount.Amount, createResp.Amount.AssetId)
	fmt.Printf("  Currency: %s\n", createResp.Currency)
	fmt.Printf("  Status: %s\n\n", createResp.Status)

	// Step 2: Update payment intent to specify crypto payment method
	// NOTE: Skip this step if using MartianPay.js widget (widget handles it automatically)
	paymentMethodType := developer.PaymentMethodTypeCrypto
	assetId := "USDC-Ethereum-TEST"
	updateReq := martianpay.PaymentIntentUpdateRequest{
		ID:                createResp.ID,
		PaymentMethodType: &paymentMethodType,
		PaymentMethodData: &developer.PaymentMethodConfirmOptions{
			Crypto: &developer.CryptoOption{
				AssetId: &assetId,
			},
		},
	}

	updateResp, err := client.UpdatePaymentIntent(updateReq)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	// Step 3: Extract deposit address from response
	// The API returns a unique deposit address for this payment
	// User sends crypto to this address to complete the payment
	fmt.Printf("✓ Step 2 - Payment Intent Updated (Crypto Payment Details):\n")
	fmt.Printf("  Deposit Address: %s\n", *updateResp.Charges[0].PaymentMethodOptions.Crypto.DepositAddress)
	fmt.Printf("  Asset ID: %s\n", *updateResp.Charges[0].PaymentMethodOptions.Crypto.AssetId)
	fmt.Printf("  Network: %s\n", *updateResp.Charges[0].PaymentMethodOptions.Crypto.Network)
	fmt.Printf("  Amount: %s\n", *updateResp.Charges[0].PaymentMethodOptions.Crypto.Amount)

	fmt.Printf("\n  💡 Integration Tips:\n")
	fmt.Printf("  • API-only approach: Display this deposit address to user\n")
	fmt.Printf("  • MartianPay.js Widget (Recommended): Use payment_intent.client_secret\n")
	fmt.Printf("    - Widget handles UpdatePaymentIntent automatically\n")
	fmt.Printf("    - No need to call UpdatePaymentIntent from backend\n")
	fmt.Printf("    - Docs: https://docs.martianpay.com/v1/docs/martianpay-js-usage\n")
}

func getPaymentIntent(client *martianpay.Client) {
	fmt.Println("Getting Payment Intent...")
	fmt.Print("Enter Payment Intent ID: ")

	var id string
	fmt.Scanln(&id)
	if id == "" {
		id = "pi_example_id" // Default for demo
	}

	req := martianpay.PaymentIntentGetReq{ID: id}
	response, err := client.GetPaymentIntent(req)

	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("✓ Payment Intent Retrieved:\n")
	fmt.Printf("  ID: %s\n", response.ID)
	fmt.Printf("  Amount: %s %s\n", response.Amount.Amount, response.Amount.AssetId)
	fmt.Printf("  Currency: %s\n", response.Currency)
	fmt.Printf("  Status: %s\n", response.Status)
}

func listPaymentIntents(client *martianpay.Client) {
	fmt.Println("Listing Payment Intents...")

	email := "user@example.com"
	req := martianpay.PaymentIntentListReq{
		Page:          0,
		PageSize:      10,
		CustomerEmail: &email,
	}

	response, err := client.ListPaymentIntents(req)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("✓ Total Records: %d\n", response.Total)
	fmt.Printf("  Page: %d, Page Size: %d\n\n", response.Page, response.PageSize)

	for i, intent := range response.PaymentIntents {
		fmt.Printf("  [%d] ID: %s, Amount: %s %s\n",
			i+1, intent.ID, intent.Amount.Amount, intent.Amount.AssetId)
	}
}

func cancelPaymentIntent(client *martianpay.Client) {
	fmt.Println("Canceling Payment Intent...")

	// First create a payment intent to cancel
	description := "Test payment intent for cancellation"
	createReq := martianpay.PaymentIntentCreateRequest{
		PaymentIntentParams: developer.PaymentIntentParams{
			Amount:          "10.00",
			Currency:        "USD",
			MerchantOrderId: "order_cancel_test",
			Description:     &description,
		},
	}

	createResp, err := client.CreatePaymentIntent(createReq)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("  Created Payment Intent: %s\n", createResp.ID)

	reason := "Customer requested cancellation"
	cancelReq := martianpay.PaymentIntentCancelReq{
		ID:     createResp.ID,
		Reason: &reason,
	}

	cancelResp, err := client.CancelPaymentIntent(cancelReq)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("✓ Payment Intent Canceled:\n")
	fmt.Printf("  ID: %s\n", cancelResp.ID)
	fmt.Printf("  Status: %s\n", cancelResp.PaymentIntentStatus)
	fmt.Printf("  Canceled At: %d\n", cancelResp.CanceledAt)
}

func listPaymentIntentsWithPermanentDeposit(client *martianpay.Client) {
	fmt.Println("Listing Payment Intents with Permanent Deposit...")

	permanentDeposit := true
	assetId := "USDC-Ethereum-TEST"

	req := martianpay.PaymentIntentListReq{
		Page:                    0,
		PageSize:                10,
		PermanentDeposit:        &permanentDeposit,
		PermanentDepositAssetId: &assetId,
	}

	response, err := client.ListPaymentIntents(req)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("✓ Total Records: %d\n\n", response.Total)

	for i, intent := range response.PaymentIntents {
		fmt.Printf("  [%d] ID: %s, Permanent Deposit: %t, Asset: %s\n",
			i+1, intent.ID, intent.PermanentDeposit, intent.PermanentDepositAssetId)
	}
}

func fiatPaymentWithNewCard(client *martianpay.Client) {
	// This example demonstrates two integration approaches for card payments:
	//
	// APPROACH 1 (Recommended): MartianPay.js Integration
	// -------------------------------------------------------
	// After creating a payment intent, use MartianPay.js on the frontend:
	// - Documentation: https://docs.martianpay.com/v1/docs/martianpay-js-usage
	// - The frontend widget handles payment method selection (cards or crypto)
	// - The widget automatically calls the update payment intent API
	// - No need to manually call UpdatePaymentIntent from backend
	//
	// APPROACH 2 (API-only): Direct API Integration (demonstrated in this example)
	// -------------------------------------------------------------------------
	// If you don't use MartianPay.js, you need to handle everything via API:
	//
	// For CRYPTO payments:
	// - Call UpdatePaymentIntent with crypto asset_id
	// - API returns deposit address for the selected crypto
	// - Display the address to user for payment
	//
	// For CARD payments (this example):
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

	fmt.Println("Fiat Payment with New Card...")
	fmt.Println("Note: This example shows the API-only approach with Stripe integration")

	email := "fiat_test@example.com"
	listCustomerReq := martianpay.CustomerListRequest{
		Page:     0,
		PageSize: 10,
		Email:    &email,
	}

	listResp, err := client.ListCustomers(listCustomerReq)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	// Step 1: Query or create a customer
	var customerID string
	if listResp != nil && len(listResp.Customers) > 0 {
		customerID = listResp.Customers[0].ID
		fmt.Printf("✓ Step 1 - Using existing customer: %s\n", customerID)
	} else {
		name := "Fiat Test Customer"
		createCustomerReq := martianpay.CustomerCreateRequest{
			CustomerParams: developer.CustomerParams{
				Email: &email,
				Name:  &name,
			},
		}

		customerResp, err := client.CreateCustomer(createCustomerReq)
		if err != nil {
			fmt.Printf("✗ API Error: %v\n", err)
			return
		}
		customerID = customerResp.ID
		fmt.Printf("✓ Step 1 - Created new customer: %s\n", customerID)
	}

	// Step 2: Create a payment intent
	description := "Test fiat payment with new card"
	createPIReq := martianpay.PaymentIntentCreateRequest{
		PaymentIntentParams: developer.PaymentIntentParams{
			Amount:          "25.00",
			Currency:        "USD",
			MerchantOrderId: "order_fiat_new_card",
			Description:     &description,
			Customer:        &customerID,
		},
	}

	piResp, err := client.CreatePaymentIntent(createPIReq)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("\n✓ Step 2 - Payment Intent Created:\n")
	fmt.Printf("  ID: %s\n", piResp.ID)
	fmt.Printf("  Amount: %s %s\n", piResp.Amount.Amount, piResp.Amount.AssetId)
	fmt.Printf("  Currency: %s\n", piResp.Currency)

	// Step 3: Update payment intent with cards payment method
	// This triggers MartianPay to prepare Stripe payment and return stripe_payload
	paymentMethodType := developer.PaymentMethodTypeCards
	currency := "USD"
	savePaymentMethod := true // Save the card for future use (creates reusable payment method)

	updatePIReq := martianpay.PaymentIntentUpdateRequest{
		ID:                piResp.ID,
		PaymentMethodType: &paymentMethodType,
		PaymentMethodData: &developer.PaymentMethodConfirmOptions{
			Fiat: &developer.FiatOption{
				Currency:          &currency,
				SavePaymentMethod: &savePaymentMethod, // Save card for future payments
				// Note: PaymentMethodID is NOT provided for new card payments
				// It's only used when charging a saved card (see fiatPaymentWithSavedCard)
			},
		},
	}

	updateResp, err := client.UpdatePaymentIntent(updatePIReq)
	if err != nil {
		fmt.Printf("✗ API Error updating payment intent: %v\n", err)
		return
	}

	// Step 4: Extract and display stripe_payload
	// The API returns stripe_payload containing:
	// - client_secret: Used by Stripe.js to confirm payment on frontend
	// - public_key: Stripe publishable key for initializing Stripe.js
	// - customer_id: Stripe customer ID (if SavePaymentMethod is true)
	fmt.Printf("\n✓ Step 3 - Payment Intent Updated with Card Payment Method\n")
	fmt.Printf("  Status: %s\n", updateResp.Status)

	if len(updateResp.Charges) > 0 {
		for i, charge := range updateResp.Charges {
			if charge.StripePayload != nil {
				fmt.Printf("\n  Stripe Payload Details (Charge %d):\n", i)
				fmt.Printf("  ────────────────────────────────────────────\n")
				fmt.Printf("  Client Secret:    %s\n", charge.StripePayload.ClientSecret)
				fmt.Printf("  Public Key:       %s\n", charge.StripePayload.PublicKey)
				fmt.Printf("  Stripe Status:    %s\n", charge.StripePayload.Status)
				fmt.Printf("  Customer ID:      %s\n", charge.StripePayload.CustomerID)
				fmt.Printf("  Payment Provider: %s\n", charge.PaymentProvider)
				fmt.Printf("  ────────────────────────────────────────────\n")

				fmt.Printf("\n  💡 Frontend Integration Example:\n")
				fmt.Printf("  ────────────────────────────────────────────\n")
				fmt.Printf("  // Initialize Stripe with public_key\n")
				fmt.Printf("  const stripe = Stripe('%s');\n", charge.StripePayload.PublicKey)
				fmt.Printf("\n")
				fmt.Printf("  // Confirm card payment with client_secret\n")
				fmt.Printf("  const {error} = await stripe.confirmCardPayment(\n")
				fmt.Printf("    '%s',\n", charge.StripePayload.ClientSecret)
				fmt.Printf("    {\n")
				fmt.Printf("      payment_method: {\n")
				fmt.Printf("        card: cardElement,\n")
				fmt.Printf("        billing_details: { name: 'Customer Name' }\n")
				fmt.Printf("      }\n")
				fmt.Printf("    }\n")
				fmt.Printf("  );\n")
				fmt.Printf("  ────────────────────────────────────────────\n")

				fmt.Printf("\n  💡 Integration Note:\n")
				fmt.Printf("  This example shows API-only approach with manual UpdatePaymentIntent call.\n")
				fmt.Printf("  If using MartianPay.js Widget (Recommended):\n")
				fmt.Printf("  • Simply pass payment_intent.client_secret to the widget\n")
				fmt.Printf("  • Widget handles UpdatePaymentIntent automatically\n")
				fmt.Printf("  • No backend UpdatePaymentIntent call needed\n")
				fmt.Printf("  • Docs: https://docs.martianpay.com/v1/docs/martianpay-js-usage\n")
			}
		}
	} else {
		fmt.Printf("  No charges found - stripe_payload will be in charges array\n")
	}
}

func fiatPaymentWithSavedCard(client *martianpay.Client) {
	// This example demonstrates charging a saved card (previously saved payment method).
	//
	// INTEGRATION OPTIONS:
	// --------------------
	// Option 1 (Recommended): MartianPay.js Widget
	// - Use MartianPay.js widget on your frontend (https://docs.martianpay.com/v1/docs/martianpay-js-usage)
	// - After creating payment intent, display the widget with payment_intent.client_secret
	// - The widget shows saved payment methods and handles payment automatically
	// - NO NEED to call UpdatePaymentIntent from your backend
	//
	// Option 2 (API-only): Direct Backend Integration (this example)
	// - List customer's saved payment methods
	// - Create payment intent
	// - Call UpdatePaymentIntent with saved payment_method_id
	// - Payment is processed immediately with the saved card

	fmt.Println("Fiat Payment with Saved Card...")
	fmt.Println("Note: This example shows charging a saved payment method via API")

	fmt.Print("Enter Customer ID (or press Enter for demo): ")
	var customerID string
	fmt.Scanln(&customerID)
	if customerID == "" {
		customerID = "cus_7Xpxk2n22WAEBDaBVlG5dnRf" // Default customer with saved cards
	}

	// Step 1: List customer's saved payment methods
	pmListReq := martianpay.CustomerPaymentMethodListRequest{
		CustomerID: customerID,
	}

	pmListResp, err := client.ListCustomerPaymentMethods(pmListReq)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	if len(pmListResp.PaymentMethods) == 0 {
		fmt.Println("  No saved payment methods found")
		return
	}

	fmt.Printf("✓ Found %d saved payment method(s)\n", len(pmListResp.PaymentMethods))
	for i, pm := range pmListResp.PaymentMethods {
		fmt.Printf("  [%d] %s ending in %s (expires %d/%d)\n",
			i+1, pm.Brand, pm.Last4, pm.ExpMonth, pm.ExpYear)
	}

	// Step 2: Create a payment intent
	description := "Test fiat payment with saved card"
	createPIReq := martianpay.PaymentIntentCreateRequest{
		PaymentIntentParams: developer.PaymentIntentParams{
			Amount:          "15.00",
			Currency:        "USD",
			MerchantOrderId: "order_fiat_saved_card",
			Description:     &description,
			Customer:        &customerID,
		},
	}

	piResp, err := client.CreatePaymentIntent(createPIReq)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("\n✓ Payment Intent Created: %s\n", piResp.ID)

	// Step 3: Update payment intent with saved card
	// NOTE: Skip this step if using MartianPay.js widget (widget handles it automatically)
	paymentMethodType := developer.PaymentMethodTypeCards
	currency := "USD"
	savedPaymentMethodID := pmListResp.PaymentMethods[0].ID // Use first saved card

	updatePIReq := martianpay.PaymentIntentUpdateRequest{
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
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("\n✓ Step 3 - Payment Processed with Saved Card:\n")
	fmt.Printf("  Payment Intent ID: %s\n", updateResp.ID)
	fmt.Printf("  Status: %s\n", updateResp.Status)
	fmt.Printf("  Payment Method ID: %s\n", savedPaymentMethodID)

	fmt.Printf("\n  💡 Integration Note:\n")
	fmt.Printf("  • API-only approach: Payment processed immediately with saved card\n")
	fmt.Printf("  • MartianPay.js Widget (Recommended): Use payment_intent.client_secret\n")
	fmt.Printf("    - Widget shows saved payment methods automatically\n")
	fmt.Printf("    - Widget handles UpdatePaymentIntent for you\n")
	fmt.Printf("    - No backend UpdatePaymentIntent call needed\n")
	fmt.Printf("    - Docs: https://docs.martianpay.com/v1/docs/martianpay-js-usage\n")
}

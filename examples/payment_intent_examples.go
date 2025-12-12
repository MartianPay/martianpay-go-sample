package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/MartianPay/martianpay-go-sample/pkg/developer"
	martianpay "github.com/MartianPay/martianpay-go-sample/sdk"
)

// generateOrderID generates a random merchant order ID
func generateOrderID(prefix string) string {
	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Intn(1000000)
	timestamp := time.Now().Unix()
	return fmt.Sprintf("%s_%d_%d", prefix, timestamp, randomNum)
}

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
	createReq := &developer.PaymentIntentCreateRequest{
		PaymentIntentParams: developer.PaymentIntentParams{
			Amount:          "0.10",
			Currency:        "USD",
			MerchantOrderId: generateOrderID("order"),
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
	updateReq := &developer.PaymentIntentUpdateRequest{
		PaymentMethodType: &paymentMethodType,
		PaymentMethodData: &developer.PaymentMethodConfirmOptions{
			Crypto: &developer.CryptoOption{
				AssetId: &assetId,
			},
		},
	}

	updateResp, err := client.UpdatePaymentIntent(createResp.ID, updateReq)
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
	fmt.Println("  Fetching payment intents...")

	// List payment intents first
	listReq := &developer.PaymentIntentListRequest{
		Pagination: developer.Pagination{
			Page:     0,
			PageSize: 10,
		},
	}
	listResp, err := client.ListPaymentIntents(listReq)
	if err == nil && len(listResp.PaymentIntents) > 0 {
		fmt.Printf("\n  Recent Payment Intents:\n")
		for i, pi := range listResp.PaymentIntents {
			fmt.Printf("  [%d] ID: %s - %s", i+1, pi.ID, pi.Status)
			if pi.Amount != nil {
				fmt.Printf(" - %s %s", pi.Amount.Amount, pi.Amount.AssetId)
			}
			fmt.Println()
		}
		fmt.Print("\nEnter payment intent number or ID: ")
	} else {
		fmt.Print("\nEnter Payment Intent ID: ")
	}

	var choice string
	fmt.Scanln(&choice)

	var id string
	if choice != "" && listResp != nil && len(listResp.PaymentIntents) > 0 {
		// Try to find by ID first
		foundByID := false
		for _, pi := range listResp.PaymentIntents {
			if pi.ID == choice {
				id = choice
				foundByID = true
				break
			}
		}
		// If not found by ID, try as number
		if !foundByID {
			var idx int
			fmt.Sscanf(choice, "%d", &idx)
			if idx > 0 && idx <= len(listResp.PaymentIntents) {
				id = listResp.PaymentIntents[idx-1].ID
			}
		}
		if id == "" {
			id = choice
		}
	} else if choice != "" {
		id = choice
	} else {
		id = "pi_example_id"
	}

	response, err := client.GetPaymentIntent(id)

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

	req := &developer.PaymentIntentListRequest{
		Pagination: developer.Pagination{
			Page:     0,
			PageSize: 10,
		},
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
	createReq := &developer.PaymentIntentCreateRequest{
		PaymentIntentParams: developer.PaymentIntentParams{
			Amount:          "0.10",
			Currency:        "USD",
			MerchantOrderId: generateOrderID("order_cancel"),
			Description:     &description,
		},
	}

	createResp, err := client.CreatePaymentIntent(createReq)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("  Created Payment Intent: %s\n", createResp.ID)

	cancelReq := &developer.PaymentIntentCancelRequest{
		Reason: developer.PaymentIntentCancellationReasonRequestedByCustomer,
	}

	cancelResp, err := client.CancelPaymentIntent(createResp.ID, cancelReq)
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

	req := &developer.PaymentIntentListRequest{
		Pagination: developer.Pagination{
			Page:     0,
			PageSize: 10,
		},
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

	email := generateRandomEmail("fiat_test")
	listCustomerReq := &developer.CustomerListRequest{
		Pagination: developer.Pagination{
			Page:     0,
			PageSize: 10,
		},
		Email: &email,
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
		createCustomerReq := &developer.CustomerCreateRequest{
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
	createPIReq := &developer.PaymentIntentCreateRequest{
		PaymentIntentParams: developer.PaymentIntentParams{
			Amount:          "1.00",
			Currency:        "USD",
			MerchantOrderId: generateOrderID("order_fiat_new_card"),
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

	updatePIReq := &developer.PaymentIntentUpdateRequest{
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

	updateResp, err := client.UpdatePaymentIntent(piResp.ID, updatePIReq)
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
	pmListResp, err := client.ListCustomerPaymentMethods(customerID)
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
	createPIReq := &developer.PaymentIntentCreateRequest{
		PaymentIntentParams: developer.PaymentIntentParams{
			Amount:          "1.00",
			Currency:        "USD",
			MerchantOrderId: generateOrderID("order_fiat_saved_card"),
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

	fmt.Printf("\n✓ Step 2 - Preparing to charge saved card:\n")
	fmt.Printf("  Currency: %s\n", currency)
	fmt.Printf("  Payment Method ID: %s\n", savedPaymentMethodID)

	updatePIReq := &developer.PaymentIntentUpdateRequest{
		PaymentMethodType: &paymentMethodType,
		PaymentMethodData: &developer.PaymentMethodConfirmOptions{
			Fiat: &developer.FiatOption{
				Currency:        &currency,
				PaymentMethodID: &savedPaymentMethodID, // Provide saved payment_method_id
				// SavePaymentMethod is not needed here since card is already saved
			},
		},
	}

	updateResp, err := client.UpdatePaymentIntent(piResp.ID, updatePIReq)
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

// createPaymentIntentWithPaymentLink creates a payment intent using a payment link
func createPaymentIntentWithPaymentLink(client *martianpay.Client) {
	fmt.Println("Creating Payment Intent with Payment Link...")
	fmt.Println("This example creates a payment intent from a payment link with variant and subscription")

	// Step 1: List payment links
	fmt.Println("\n  Step 1 - Fetching payment links...")
	plReq := &developer.PaymentLinkListRequest{
		Page:     0,
		PageSize: 10,
	}
	plResp, err := client.ListPaymentLinks(plReq)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	if len(plResp.PaymentLinks) == 0 {
		fmt.Println("✗ No payment links found. Please create one first.")
		return
	}

	// Categorize payment links: with selling plan vs without
	var withSellingPlan []*developer.PaymentLink
	var withoutSellingPlan []*developer.PaymentLink

	for _, link := range plResp.PaymentLinks {
		hasSellingPlan := false
		if link.Product != nil && len(link.Product.SellingPlanGroups) > 0 {
			hasSellingPlan = true
		}

		if hasSellingPlan {
			withSellingPlan = append(withSellingPlan, link)
		} else {
			withoutSellingPlan = append(withoutSellingPlan, link)
		}
	}

	fmt.Printf("\n  Available Payment Links:\n")

	// Display payment links with selling plans first
	if len(withSellingPlan) > 0 {
		fmt.Printf("\n  📅 With Subscription Plans:\n")
		for i, link := range withSellingPlan {
			fmt.Printf("  [%d] ", i+1)
			if link.Product != nil {
				fmt.Printf("%s", link.Product.Name)
			}
			fmt.Printf(" (ID: %s)\n", link.ID)
			if link.PriceRange != nil && link.PriceRange.Min != nil {
				fmt.Printf("      Price: %s %s", link.PriceRange.Min.Amount, link.PriceRange.Min.AssetId)
				if link.PriceRange.Max != nil && link.PriceRange.Max.Amount != link.PriceRange.Min.Amount {
					fmt.Printf(" - %s %s", link.PriceRange.Max.Amount, link.PriceRange.Max.AssetId)
				}
				fmt.Println()
			}
			if link.Product != nil && len(link.Product.SellingPlanGroups) > 0 {
				fmt.Printf("      Plans: %d selling plan group(s)\n", len(link.Product.SellingPlanGroups))
			}
		}
	}

	// Display payment links without selling plans
	if len(withoutSellingPlan) > 0 {
		fmt.Printf("\n  💳 One-Time Payment:\n")
		startIdx := len(withSellingPlan)
		for i, link := range withoutSellingPlan {
			fmt.Printf("  [%d] ", startIdx+i+1)
			if link.Product != nil {
				fmt.Printf("%s", link.Product.Name)
			}
			fmt.Printf(" (ID: %s)\n", link.ID)
			if link.PriceRange != nil && link.PriceRange.Min != nil {
				fmt.Printf("      Price: %s %s", link.PriceRange.Min.Amount, link.PriceRange.Min.AssetId)
				if link.PriceRange.Max != nil && link.PriceRange.Max.Amount != link.PriceRange.Min.Amount {
					fmt.Printf(" - %s %s", link.PriceRange.Max.Amount, link.PriceRange.Max.AssetId)
				}
				fmt.Println()
			}
		}
	}

	// Rebuild full list in display order for selection
	allLinks := append(withSellingPlan, withoutSellingPlan...)

	fmt.Print("\nEnter payment link number (or press Enter for first): ")
	var linkChoice string
	fmt.Scanln(&linkChoice)

	selectedIdx := 0
	if linkChoice != "" {
		var idx int
		fmt.Sscanf(linkChoice, "%d", &idx)
		if idx > 0 && idx <= len(allLinks) {
			selectedIdx = idx - 1
		}
	}

	selectedLink := allLinks[selectedIdx]
	fmt.Printf("  Selected: %s\n", selectedLink.ID)

	// Step 2: Get payment link details to show variants
	fmt.Println("\n  Step 2 - Getting payment link details...")
	linkDetails, err := client.GetPaymentLink(selectedLink.ID)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	if len(linkDetails.PrimaryVariants) == 0 {
		fmt.Println("✗ No variants found in this payment link")
		return
	}

	// Step 3: Select variant
	fmt.Printf("\n  Available Variants:\n")
	for i, pv := range linkDetails.PrimaryVariants {
		fmt.Printf("  [%d] Variant ID: %s\n", i+1, pv.VariantID)
		if pv.Variant != nil {
			fmt.Printf("      Options: ")
			for optName, optValue := range pv.Variant.OptionValues {
				fmt.Printf("%s=%s ", optName, optValue)
			}
			fmt.Println()
			if pv.Variant.Price != nil {
				fmt.Printf("      Price: %s %s\n", pv.Variant.Price.Amount, pv.Variant.Price.AssetId)
			}
		}
	}

	fmt.Print("\nEnter variant number (or press Enter for first): ")
	var variantChoice string
	fmt.Scanln(&variantChoice)

	selectedVariantIdx := 0
	if variantChoice != "" {
		var idx int
		fmt.Sscanf(variantChoice, "%d", &idx)
		if idx > 0 && idx <= len(linkDetails.PrimaryVariants) {
			selectedVariantIdx = idx - 1
		}
	}

	selectedVariant := linkDetails.PrimaryVariants[selectedVariantIdx]
	fmt.Printf("  Selected Variant: %s\n", selectedVariant.VariantID)

	// Step 4: Check for selling plans
	var sellingPlanID *string
	if linkDetails.Product != nil && len(linkDetails.Product.SellingPlanGroups) > 0 {
		fmt.Printf("\n  This product has selling plans (subscriptions)\n")

		// Build a flat list of all selling plans for easy selection
		type PlanChoice struct {
			ID          string
			Name        string
			GroupName   string
			Interval    string
			IntervalCnt string
			Policies    []developer.PricingPolicyItem
		}
		var allPlans []PlanChoice

		for _, spg := range linkDetails.Product.SellingPlanGroups {
			fmt.Printf("\n  Selling Plan Group: %s\n", spg.Name)
			if len(spg.SellingPlans) > 0 {
				for _, sp := range spg.SellingPlans {
					allPlans = append(allPlans, PlanChoice{
						ID:          sp.ID,
						Name:        sp.Name,
						GroupName:   spg.Name,
						Interval:    sp.BillingPolicy.Interval,
						IntervalCnt: sp.BillingPolicy.IntervalCount,
						Policies:    sp.PricingPolicy,
					})

					fmt.Printf("    [%d] ID: %s\n", len(allPlans), sp.ID)
					fmt.Printf("        Name: %s\n", sp.Name)
					fmt.Printf("        Billing: Every %s %s\n",
						sp.BillingPolicy.IntervalCount,
						sp.BillingPolicy.Interval)
					if len(sp.PricingPolicy) > 0 {
						fmt.Printf("        Discounts:\n")
						for _, policy := range sp.PricingPolicy {
							fmt.Printf("          - %s%% off (%s)\n",
								policy.AdjustmentValue,
								policy.PolicyType)
						}
					}
				}
			}
		}

		fmt.Print("\nEnter selling plan number or ID (or press Enter to skip): ")
		var spChoice string
		fmt.Scanln(&spChoice)

		if spChoice != "" {
			// Try to find by ID first
			foundByID := false
			for _, plan := range allPlans {
				if plan.ID == spChoice {
					sellingPlanID = &spChoice
					foundByID = true
					fmt.Printf("  Selected: %s (%s)\n", plan.Name, plan.ID)
					break
				}
			}

			// If not found by ID, try as number
			if !foundByID {
				var idx int
				fmt.Sscanf(spChoice, "%d", &idx)
				if idx > 0 && idx <= len(allPlans) {
					sellingPlanID = &allPlans[idx-1].ID
					fmt.Printf("  Selected: %s (%s)\n", allPlans[idx-1].Name, allPlans[idx-1].ID)
				} else {
					// Still treat as ID even if not found in list
					sellingPlanID = &spChoice
					fmt.Printf("  Using ID: %s\n", spChoice)
				}
			}
		}
	}

	// Step 5: Get quantity
	fmt.Print("\nEnter quantity [1]: ")
	var qtyStr string
	fmt.Scanln(&qtyStr)
	quantity := 1
	if qtyStr != "" {
		fmt.Sscanf(qtyStr, "%d", &quantity)
	}

	// Step 6: Get shipping address
	fmt.Println("\n  Step 3 - Shipping Address:")
	fmt.Print("Country [US]: ")
	var country string
	fmt.Scanln(&country)
	if country == "" {
		country = "US"
	}

	fmt.Print("Line1 [123 Main St]: ")
	var line1 string
	fmt.Scanln(&line1)
	if line1 == "" {
		line1 = "123 Main St"
	}

	fmt.Print("City [New York]: ")
	var city string
	fmt.Scanln(&city)
	if city == "" {
		city = "New York"
	}

	fmt.Print("State [NY]: ")
	var state string
	fmt.Scanln(&state)
	if state == "" {
		state = "NY"
	}

	fmt.Print("Postal Code [10001]: ")
	var postalCode string
	fmt.Scanln(&postalCode)
	if postalCode == "" {
		postalCode = "10001"
	}

	// Step 7: Get customer ID
	fmt.Println("\n  Step 7 - Select Customer (optional):")
	fmt.Println("  Fetching customers...")

	custReq := &developer.CustomerListRequest{
		Pagination: developer.Pagination{
			Page:     0,
			PageSize: 10,
		},
	}
	custResp, err := client.ListCustomers(custReq)
	if err == nil && len(custResp.Customers) > 0 {
		fmt.Printf("\n  Available Customers:\n")
		for i, cust := range custResp.Customers {
			fmt.Printf("  [%d] ID: %s", i+1, cust.ID)
			if cust.Email != nil && *cust.Email != "" {
				fmt.Printf(" - %s", *cust.Email)
			}
			if cust.Name != nil && *cust.Name != "" {
				fmt.Printf(" (%s)", *cust.Name)
			}
			fmt.Println()
		}
		fmt.Print("\nEnter customer number or ID (or press Enter to skip): ")
	} else {
		fmt.Print("\nEnter Customer ID (or press Enter to skip): ")
	}

	var customerChoice string
	fmt.Scanln(&customerChoice)

	var customerID string
	if customerChoice != "" && custResp != nil && len(custResp.Customers) > 0 {
		// Try to find by ID first
		foundByID := false
		for _, cust := range custResp.Customers {
			if cust.ID == customerChoice {
				customerID = cust.ID
				foundByID = true
				break
			}
		}
		// If not found by ID, try as number
		if !foundByID {
			var idx int
			fmt.Sscanf(customerChoice, "%d", &idx)
			if idx > 0 && idx <= len(custResp.Customers) {
				customerID = custResp.Customers[idx-1].ID
			}
		}
		if customerID == "" {
			customerID = customerChoice
		}
	} else if customerChoice != "" {
		customerID = customerChoice
	}

	if customerID != "" {
		fmt.Printf("  Selected Customer: %s\n", customerID)
	}

	// Step 8: Get receipt email
	fmt.Print("Enter Receipt Email: ")
	var receiptEmail string
	fmt.Scanln(&receiptEmail)

	// Step 9: Get merchant order ID
	merchantOrderID := generateOrderID("order")
	fmt.Printf("Generated Merchant Order ID: %s\n", merchantOrderID)
	fmt.Print("Press Enter to use this, or enter custom order ID: ")
	var customOrderID string
	fmt.Scanln(&customOrderID)
	if customOrderID != "" {
		merchantOrderID = customOrderID
	}

	// Step 10: Get metadata
	fmt.Println("\n  Metadata (key-value pairs):")
	fmt.Print("Enter number of metadata entries [0]: ")
	var metadataCountStr string
	fmt.Scanln(&metadataCountStr)
	metadataCount := 0
	if metadataCountStr != "" {
		fmt.Sscanf(metadataCountStr, "%d", &metadataCount)
	}

	metadata := make(map[string]string)
	for i := 0; i < metadataCount; i++ {
		fmt.Printf("  Entry %d:\n", i+1)
		fmt.Print("    Key: ")
		var key string
		fmt.Scanln(&key)
		fmt.Print("    Value: ")
		var value string
		fmt.Scanln(&value)
		if key != "" {
			metadata[key] = value
		}
	}

	// Step 11: Create payment intent
	fmt.Println("\n  Step 11 - Creating Payment Intent...")

	req := &developer.PaymentIntentCreateRequest{
		PaymentIntentParams: developer.PaymentIntentParams{
			MerchantOrderId: merchantOrderID,
			ReceiptEmail:    receiptEmail,
			Metadata:        metadata,
		},
		PaymentLinkID: &selectedLink.ID,
		PrimaryVariant: &developer.VariantSelectionRequest{
			VariantID:     selectedVariant.VariantID,
			Quantity:      quantity,
			SellingPlanID: sellingPlanID,
		},
		Addons: []developer.VariantSelectionRequest{},
		ShippingAddress: &developer.PaymentIntentShippingAddress{
			Country:    country,
			State:      &state,
			City:       city,
			Line1:      line1,
			PostalCode: postalCode,
		},
	}

	if customerID != "" {
		req.Customer = &customerID
	}

	if linkDetails.Product != nil {
		req.Description = &linkDetails.Product.Name
		if linkDetails.Product.Version > 0 {
			req.ProductVersion = &linkDetails.Product.Version
		}
	}

	response, err := client.CreatePaymentIntent(req)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("\n✓ Payment Intent Created:\n")
	fmt.Printf("  ID: %s\n", response.ID)
	fmt.Printf("  Status: %s\n", response.Status)
	fmt.Printf("  Merchant Order ID: %s\n", response.MerchantOrderId)
	if response.Amount != nil {
		fmt.Printf("  Amount: %s %s\n", response.Amount.Amount, response.Amount.AssetId)
	}
	if response.ClientSecret != "" {
		fmt.Printf("  Client Secret: %s\n", response.ClientSecret)
	}
	if sellingPlanID != nil {
		fmt.Printf("  Subscription: Yes (Selling Plan ID: %s)\n", *sellingPlanID)
	}
	if response.Subscription != nil {
		fmt.Printf("  Subscription ID: %s\n", *response.Subscription)
	}
	if len(metadata) > 0 {
		fmt.Printf("  Metadata: %d entries\n", len(metadata))
		for k, v := range metadata {
			fmt.Printf("    %s: %s\n", k, v)
		}
	}

	fmt.Printf("\n  💡 Next Steps:\n")
	fmt.Printf("  • Use MartianPay.js Widget with client_secret: %s\n", response.ClientSecret)
	fmt.Printf("  • Widget URL: https://docs.martianpay.com/v1/docs/martianpay-js-usage\n")
	fmt.Printf("  • Customer will select payment method in the widget\n")
}

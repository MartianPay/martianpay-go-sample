// Package main provides examples for the MartianPay Refund API.
// Refunds allow merchants to return funds to customers for completed payments.
// Both full and partial refunds are supported for cryptocurrency and fiat payments.
package main

import (
	"fmt"

	"github.com/MartianPay/martianpay-go-sample/pkg/developer"
	martianpay "github.com/MartianPay/martianpay-go-sample/sdk"
)

// Refund Examples

// createRefund demonstrates creating a refund for a completed payment intent.
// Refunds can be full or partial depending on the amount specified.
//
// Steps:
// 1. Prompt for payment intent ID (the original payment to refund)
// 2. Prompt for refund amount
// 3. Create refund with reason
// 4. Display refund details including ID and status
//
// Note: Refunds may take time to process depending on the payment method.
// For crypto refunds, funds are sent back to the original payment address.
// For fiat refunds, funds are returned via Stripe.
//
// API Endpoints Used:
//   - POST /v1/refunds
func createRefund(client *martianpay.Client) {
	fmt.Println("Creating Refund...")
	fmt.Print("Enter Payment Intent ID (or press Enter for demo): ")

	var paymentIntentID string
	fmt.Scanln(&paymentIntentID)
	if paymentIntentID == "" {
		paymentIntentID = "pi_example_id"
	}

	fmt.Print("Enter refund amount (or press Enter for default 10.00): ")
	var amountInput string
	fmt.Scanln(&amountInput)
	if amountInput == "" {
		amountInput = "10.00"
	}

	reason := "Customer requested refund"
	req := &developer.RefundCreateRequest{
		RefundParams: developer.RefundParams{
			PaymentIntent: &paymentIntentID,
			Amount:        amountInput,
			Reason:        &reason,
		},
	}

	response, err := client.CreateRefund(req)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("✓ Refund Created:\n")
	if len(response.Refunds) > 0 {
		fmt.Printf("  ID: %s\n", response.Refunds[0].ID)
		fmt.Printf("  Amount: %s %s\n", response.Refunds[0].Amount.Amount, response.Refunds[0].Amount.AssetId)
		fmt.Printf("  Status: %s\n", response.Refunds[0].Status)
	}
}

// getRefund retrieves and displays details of a specific refund by ID.
//
// Displayed Information:
//   - Refund ID
//   - Refund amount and currency
//   - Refund status (pending, succeeded, failed, canceled)
//   - Associated payment intent ID
//
// API Endpoints Used:
//   - GET /v1/refunds/:id
func getRefund(client *martianpay.Client) {
	fmt.Println("Getting Refund...")
	fmt.Print("Enter Refund ID: ")

	var id string
	fmt.Scanln(&id)
	if id == "" {
		id = "rf_example_id"
	}

	response, err := client.GetRefund(id)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("✓ Refund Retrieved:\n")
	fmt.Printf("  ID: %s\n", response.ID)
	fmt.Printf("  Amount: %s %s\n", response.Amount.Amount, response.Amount.AssetId)
	fmt.Printf("  Status: %s\n", response.Status)
	fmt.Printf("  Payment Intent: %s\n", response.PaymentIntent)
}

// listRefunds retrieves and displays a paginated list of all refunds.
//
// Features:
//   - Pagination support (page and page size)
//   - Displays refund ID, amount, currency, and status
//   - Shows total count of refunds
//
// Use Cases:
//   - View refund history
//   - Reconcile refund transactions
//   - Monitor refund status
//
// API Endpoints Used:
//   - GET /v1/refunds
func listRefunds(client *martianpay.Client) {
	fmt.Println("Listing Refunds...")

	req := &developer.RefundListRequest{
		Pagination: developer.Pagination{
			Page:     0,
			PageSize: 10,
		},
	}

	response, err := client.ListRefunds(req)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("✓ Total Refunds: %d\n", response.Total)
	fmt.Printf("  Page: %d, Page Size: %d\n\n", response.Page, response.PageSize)

	for i, refund := range response.Refunds {
		fmt.Printf("  [%d] ID: %s, Amount: %s %s, Status: %s\n",
			i+1, refund.ID, refund.Amount.Amount, refund.Amount.AssetId, refund.Status)
	}
}

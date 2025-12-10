package main

import (
	"fmt"

	"github.com/MartianPay/martianpay-go-sample/pkg/developer"
	martianpay "github.com/MartianPay/martianpay-go-sample/sdk"
)

// Refund Examples
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

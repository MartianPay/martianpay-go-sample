// Package main provides examples for the MartianPay Payroll API.
// Payroll allows merchants to send batch cryptocurrency payments to multiple recipients,
// perfect for paying employees, contractors, or distributing funds.
//
// Two payment methods are supported:
// 1. Normal - Direct blockchain transfer (may incur network fees)
// 2. Binance - Transfer via Binance Pay (lower fees for Binance users)
package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/MartianPay/martianpay-go-sample/pkg/developer"
	martianpay "github.com/MartianPay/martianpay-go-sample/sdk"
)

// generatePayrollEmail generates a unique random email for payroll item recipients.
//
// Parameters:
//   - prefix: A prefix string to identify the recipient type
//
// Returns:
//   - A unique email address in the format: prefix_timestamp_randomNumber@example.com
func generatePayrollEmail(prefix string) string {
	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Intn(1000000)
	timestamp := time.Now().Unix()
	return fmt.Sprintf("%s_%d_%d@example.com", prefix, timestamp, randomNum)
}

// Payroll Examples

// createDirectPayroll demonstrates creating a direct payroll with normal payment method.
// This function shows the complete workflow for batch cryptocurrency payments.
//
// Steps:
// 1. Query available crypto assets from merchant balance
// 2. Filter assets with sufficient balance (>= 0.1)
// 3. Let user select the cryptocurrency and network
// 4. Prompt for recipient address and amount
// 5. Create payroll with auto-approval enabled
// 6. Display payroll details and status
//
// Payment Method: Normal (direct blockchain transfer)
//
// Use Cases:
//   - Employee salary payments in crypto
//   - Contractor payments
//   - Bulk fund distribution
//
// API Endpoints Used:
//   - GET /v1/assets (query available assets)
//   - GET /v1/balance (check merchant balance)
//   - POST /v1/payrolls/direct (create payroll)
func createDirectPayroll(client *martianpay.Client) {
	fmt.Println("Creating Direct Payroll...")

	// Query assets
	assetsResp, err := client.ListAssets()
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	var cryptoAssets []*developer.Asset
	for _, asset := range assetsResp.Assets {
		if !asset.IsFiat && asset.Payable && asset.CryptoAssetParams != nil {
			cryptoAssets = append(cryptoAssets, asset)
		}
	}

	fmt.Printf("  Found %d crypto assets\n", len(cryptoAssets))

	// Query balance
	balance, err := client.GetBalance()
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	// Match assets with balance and let user choose
	type AssetOption struct {
		Coin             string
		Network          string
		AssetId          string
		AvailableBalance string
	}

	var availableAssets []AssetOption
	for _, asset := range cryptoAssets {
		for _, detail := range balance.BalanceDetails {
			if detail.Currency == asset.Id && detail.AvailableBalance != "0" && detail.AvailableBalance != "" {
				availableFloat, err := strconv.ParseFloat(detail.AvailableBalance, 64)
				if err != nil {
					continue
				}
				if availableFloat >= 0.1 {
					availableAssets = append(availableAssets, AssetOption{
						Coin:             asset.Coin,
						Network:          asset.CryptoAssetParams.Network,
						AssetId:          asset.Id,
						AvailableBalance: detail.AvailableBalance,
					})
				}
			}
		}
	}

	if len(availableAssets) == 0 {
		fmt.Println("  No sufficient balance found (>= 0.1)")
		return
	}

	// Display available assets for user to choose
	fmt.Println("\n  Available assets with sufficient balance:")
	for i, option := range availableAssets {
		fmt.Printf("  [%d] %s on %s (Balance: %s)\n", i+1, option.Coin, option.Network, option.AvailableBalance)
	}

	fmt.Print("\nSelect asset number: ")
	var assetChoice int
	fmt.Scanln(&assetChoice)
	if assetChoice < 1 || assetChoice > len(availableAssets) {
		fmt.Println("Invalid choice")
		return
	}

	selectedAsset := availableAssets[assetChoice-1]
	fmt.Printf("\n✓ Selected: %s on %s\n", selectedAsset.Coin, selectedAsset.Network)

	// Ask user to input recipient address
	fmt.Printf("\nEnter recipient address for %s on %s: ", selectedAsset.Coin, selectedAsset.Network)
	var recipientAddress string
	fmt.Scanln(&recipientAddress)
	if recipientAddress == "" {
		fmt.Println("Error: Address cannot be empty")
		return
	}

	fmt.Printf("✓ Recipient address: %s\n", recipientAddress)

	// Ask user to input amount
	fmt.Printf("\nEnter amount to send (or press Enter for default 0.1): ")
	var amountInput string
	fmt.Scanln(&amountInput)
	if amountInput == "" {
		amountInput = "0.1"
	}

	// Validate amount
	amountFloat, err := strconv.ParseFloat(amountInput, 64)
	if err != nil || amountFloat <= 0 {
		fmt.Println("Error: Invalid amount")
		return
	}

	fmt.Printf("✓ Amount: %s\n", amountInput)

	timestamp := time.Now().UnixNano()
	externalID := fmt.Sprintf("ORDER-%d", timestamp)

	req := &developer.PayrollDirectCreateRequest{
		ExternalID:  externalID,
		AutoApprove: true,
		Items: []developer.PayrollDirectItem{
			{
				ExternalID:    fmt.Sprintf("ITEM-%d-001", timestamp),
				Name:          "John Doe",
				Email:         generatePayrollEmail("john"),
				Phone:         "+1234567890",
				Coin:          selectedAsset.Coin,
				Network:       selectedAsset.Network,
				Address:       recipientAddress,
				Amount:        amountInput,
				PaymentMethod: "normal",
			},
		},
	}

	response, err := client.CreateDirectPayroll(req)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("\n✓ Payroll Created (Normal Payment Method):\n")
	fmt.Printf("  ID: %s\n", response.Payroll.ID)
	fmt.Printf("  External ID: %s\n", response.Payroll.ExternalID)
	fmt.Printf("  Status: %s\n", response.Payroll.Status)
	fmt.Printf("  Total Amount: %s\n", response.Payroll.TotalAmount)
}

// createDirectPayrollBinance demonstrates creating a direct payroll with Binance payment method.
// Binance payment method offers lower fees and faster transfers for recipients with Binance accounts.
//
// Payment Method: Binance (Binance Pay transfer)
//
// Benefits of Binance Payment:
//   - Lower transaction fees compared to blockchain transfers
//   - Faster transfer speed
//   - Requires recipient to have a Binance account
//
// Steps are similar to createDirectPayroll but uses "binance" as payment method.
//
// API Endpoints Used:
//   - GET /v1/assets
//   - GET /v1/balance
//   - POST /v1/payrolls/direct (with payment_method: "binance")
func createDirectPayrollBinance(client *martianpay.Client) {
	fmt.Println("Creating Direct Payroll (Binance Payment Method)...")

	// Query assets
	assetsResp, err := client.ListAssets()
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	var cryptoAssets []*developer.Asset
	for _, asset := range assetsResp.Assets {
		if !asset.IsFiat && asset.Payable && asset.CryptoAssetParams != nil {
			cryptoAssets = append(cryptoAssets, asset)
		}
	}

	fmt.Printf("  Found %d crypto assets\n", len(cryptoAssets))

	// Query balance
	balance, err := client.GetBalance()
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	// Match assets with balance and let user choose
	type AssetOption struct {
		Coin             string
		Network          string
		AssetId          string
		AvailableBalance string
	}

	var availableAssets []AssetOption
	for _, asset := range cryptoAssets {
		for _, detail := range balance.BalanceDetails {
			if detail.Currency == asset.Id && detail.AvailableBalance != "0" && detail.AvailableBalance != "" {
				availableFloat, err := strconv.ParseFloat(detail.AvailableBalance, 64)
				if err != nil {
					continue
				}
				if availableFloat >= 0.1 {
					availableAssets = append(availableAssets, AssetOption{
						Coin:             asset.Coin,
						Network:          asset.CryptoAssetParams.Network,
						AssetId:          asset.Id,
						AvailableBalance: detail.AvailableBalance,
					})
				}
			}
		}
	}

	if len(availableAssets) == 0 {
		fmt.Println("  No sufficient balance found (>= 0.1)")
		return
	}

	// Display available assets for user to choose
	fmt.Println("\n  Available assets with sufficient balance:")
	for i, option := range availableAssets {
		fmt.Printf("  [%d] %s on %s (Balance: %s)\n", i+1, option.Coin, option.Network, option.AvailableBalance)
	}

	fmt.Print("\nSelect asset number: ")
	var assetChoice int
	fmt.Scanln(&assetChoice)
	if assetChoice < 1 || assetChoice > len(availableAssets) {
		fmt.Println("Invalid choice")
		return
	}

	selectedAsset := availableAssets[assetChoice-1]
	fmt.Printf("\n✓ Selected: %s on %s\n", selectedAsset.Coin, selectedAsset.Network)

	// Ask user to input recipient address
	fmt.Printf("\nEnter recipient address for %s on %s (Binance): ", selectedAsset.Coin, selectedAsset.Network)
	var recipientAddress string
	fmt.Scanln(&recipientAddress)
	if recipientAddress == "" {
		fmt.Println("Error: Address cannot be empty")
		return
	}

	fmt.Printf("✓ Recipient address: %s\n", recipientAddress)

	// Ask user to input amount
	fmt.Printf("\nEnter amount to send (or press Enter for default 0.1): ")
	var amountInput string
	fmt.Scanln(&amountInput)
	if amountInput == "" {
		amountInput = "0.1"
	}

	// Validate amount
	amountFloat, err := strconv.ParseFloat(amountInput, 64)
	if err != nil || amountFloat <= 0 {
		fmt.Println("Error: Invalid amount")
		return
	}

	fmt.Printf("✓ Amount: %s\n", amountInput)

	timestamp := time.Now().UnixNano()
	externalID := fmt.Sprintf("ORDER-%d", timestamp)

	req := &developer.PayrollDirectCreateRequest{
		ExternalID:  externalID,
		AutoApprove: true,
		Items: []developer.PayrollDirectItem{
			{
				ExternalID:    fmt.Sprintf("ITEM-%d-001", timestamp),
				Name:          "Jane Smith",
				Email:         generatePayrollEmail("jane"),
				Phone:         "+1234567891",
				Coin:          selectedAsset.Coin,
				Network:       selectedAsset.Network,
				Address:       recipientAddress,
				Amount:        amountInput,
				PaymentMethod: "binance",
			},
		},
	}

	response, err := client.CreateDirectPayroll(req)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("\n✓ Payroll Created (Binance Payment Method):\n")
	fmt.Printf("  ID: %s\n", response.Payroll.ID)
	fmt.Printf("  External ID: %s\n", response.Payroll.ExternalID)
	fmt.Printf("  Status: %s\n", response.Payroll.Status)
	fmt.Printf("  Total Amount: %s\n", response.Payroll.TotalAmount)
}

// getPayroll retrieves and displays details of a specific payroll batch.
//
// Displayed Information:
//   - Payroll ID and External ID
//   - Status (pending, processing, completed, failed)
//   - Total amount to be distributed
//   - Number of payment items in the batch
//
// API Endpoints Used:
//   - GET /v1/payrolls/:id
func getPayroll(client *martianpay.Client) {
	fmt.Println("Getting Payroll...")
	fmt.Print("Enter Payroll ID: ")

	var id string
	fmt.Scanln(&id)
	if id == "" {
		id = "payroll_example_id"
	}

	response, err := client.GetPayroll(id)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("✓ Payroll Retrieved:\n")
	fmt.Printf("  ID: %s\n", response.Payroll.ID)
	fmt.Printf("  External ID: %s\n", response.Payroll.ExternalID)
	fmt.Printf("  Status: %s\n", response.Payroll.Status)
	fmt.Printf("  Total Amount: %s\n", response.Payroll.TotalAmount)
	fmt.Printf("  Total Items: %d\n", response.Payroll.TotalItemNum)
}

// listPayrolls retrieves and displays a paginated list of all payroll batches.
//
// Features:
//   - Pagination support
//   - Displays payroll ID, status, and total amount
//   - Shows total count of payroll batches
//
// API Endpoints Used:
//   - GET /v1/payrolls
func listPayrolls(client *martianpay.Client) {
	fmt.Println("Listing Payrolls...")

	req := &developer.PayrollListRequest{
		Page:     0,
		PageSize: 10,
	}

	response, err := client.ListPayrolls(req)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("✓ Total Payrolls: %d\n\n", response.Total)

	for i, payroll := range response.Payrolls {
		fmt.Printf("  [%d] ID: %s, Status: %s, Amount: %s\n",
			i+1, payroll.ID, payroll.Status, payroll.TotalAmount)
	}
}

// listPayrollItems retrieves and displays individual payment items within a payroll batch.
// Each item represents a single payment to one recipient.
//
// Displayed Information for Each Item:
//   - Item ID
//   - Payment amount
//   - Payment status (pending, processing, sent, failed)
//   - Network used for transfer
//
// Use Cases:
//   - Track individual payment status within a batch
//   - Identify failed payments for retry
//   - Reconcile individual transfers
//
// API Endpoints Used:
//   - GET /v1/payroll_items
func listPayrollItems(client *martianpay.Client) {
	fmt.Println("Listing Payroll Items...")
	fmt.Print("Enter Payroll ID: ")

	var payrollID string
	fmt.Scanln(&payrollID)
	if payrollID == "" {
		payrollID = "payroll_example_id"
	}

	req := &developer.PayrollItemsListRequest{
		PayrollID: &payrollID,
		Page:      0,
		PageSize:  10,
	}

	response, err := client.ListPayrollItems(req)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("✓ Total Items: %d\n\n", response.Total)

	for i, item := range response.PayrollItems {
		fmt.Printf("  [%d] ID: %s\n", i+1, item.ID)
		fmt.Printf("      Amount: %s, Status: %s\n", item.Amount, item.Status)
		fmt.Printf("      Network: %s\n", item.Network)
	}
}

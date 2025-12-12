package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/MartianPay/martianpay-go-sample/pkg/developer"
	martianpay "github.com/MartianPay/martianpay-go-sample/sdk"
)

// generatePayrollEmail generates a random email for payroll items
func generatePayrollEmail(prefix string) string {
	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Intn(1000000)
	timestamp := time.Now().Unix()
	return fmt.Sprintf("%s_%d_%d@example.com", prefix, timestamp, randomNum)
}

// Payroll Examples
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

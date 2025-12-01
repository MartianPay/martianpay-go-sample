package main

import (
	"fmt"
	"strconv"
	"time"

	martianpay "github.com/MartianPay/martianpay-go-sample/sdk"
	"github.com/MartianPay/martianpay-go-sample/pkg/developer"
)

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

	// Match assets with balance
	var selectedCoin, selectedNetwork string
	for _, asset := range cryptoAssets {
		for _, detail := range balance.BalanceDetails {
			if detail.Currency == asset.Id && detail.AvailableBalance != "0" && detail.AvailableBalance != "" {
				availableFloat, err := strconv.ParseFloat(detail.AvailableBalance, 64)
				if err != nil {
					continue
				}
				if availableFloat >= 0.1 {
					selectedCoin = asset.Coin
					selectedNetwork = asset.CryptoAssetParams.Network
					fmt.Printf("  Selected: %s on %s (Balance: %s)\n", selectedCoin, selectedNetwork, detail.AvailableBalance)
					break
				}
			}
		}
		if selectedCoin != "" {
			break
		}
	}

	if selectedCoin == "" {
		fmt.Println("  No sufficient balance found (>= 0.1)")
		return
	}

	timestamp := time.Now().UnixNano()
	externalID := fmt.Sprintf("ORDER-%d", timestamp)

	req := martianpay.PayrollDirectCreateRequest{
		ExternalID:  externalID,
		AutoApprove: true,
		Items: []martianpay.PayrollDirectItem{
			{
				ExternalID:    fmt.Sprintf("ITEM-%d-001", timestamp),
				Name:          "John Doe",
				Email:         "john@example.com",
				Phone:         "+1234567890",
				Coin:          selectedCoin,
				Network:       selectedNetwork,
				Address:       "TN9RRaXkCFtTXRso2GdTZxSxxwufzxLQPP",
				Amount:        "0.1",
				PaymentMethod: "normal",
			},
		},
	}

	response, err := client.CreateDirectPayroll(req)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("✓ Payroll Created:\n")
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

	response, err := client.GetPayroll(martianpay.PayrollGetReq{ID: id})
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

	req := martianpay.PayrollListReq{
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

	req := martianpay.PayrollItemsListReq{
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

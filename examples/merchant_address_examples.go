package main

import (
	"bufio"
	"fmt"
	"os"

	martianpay "github.com/MartianPay/martianpay-go-sample/sdk"
)

// Merchant Address Examples
func createMerchantAddress(client *martianpay.Client) {
	fmt.Println("Creating Merchant Address...")

	req := martianpay.MerchantAddressCreateRequest{
		Network: "Ethereum Sepolia",
		Address: "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb",
	}

	resp, err := client.CreateMerchantAddress(req)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("✓ Merchant Address Created:\n")
	fmt.Printf("  ID: %s\n", resp.ID)
	fmt.Printf("  Network: %s\n", resp.Network)
	fmt.Printf("  Address: %s\n", resp.Address)
	fmt.Printf("  Status: %s\n", resp.Status)
}

func getMerchantAddress(client *martianpay.Client) {
	fmt.Println("Getting Merchant Address...")
	fmt.Print("Enter Address ID: ")

	var id string
	fmt.Scanln(&id)
	if id == "" {
		id = "ma_example_id"
	}

	resp, err := client.GetMerchantAddress(id)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("✓ Merchant Address Retrieved:\n")
	fmt.Printf("  ID: %s\n", resp.ID)
	fmt.Printf("  Network: %s\n", resp.Network)
	fmt.Printf("  Address: %s\n", resp.Address)
	fmt.Printf("  Status: %s\n", resp.Status)
	fmt.Printf("  Alias: %s\n", resp.Alias)
}

func updateMerchantAddress(client *martianpay.Client) {
	fmt.Println("Updating Merchant Address...")
	fmt.Print("Enter Address ID: ")

	var id string
	fmt.Scanln(&id)
	if id == "" {
		id = "ma_example_id"
	}

	alias := "My Main Wallet"
	req := martianpay.MerchantAddressUpdateRequest{
		ID:    id,
		Alias: &alias,
	}

	resp, err := client.UpdateMerchantAddress(req)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("✓ Merchant Address Updated:\n")
	fmt.Printf("  ID: %s\n", resp.ID)
	fmt.Printf("  Alias: %s\n", resp.Alias)
}

func verifyMerchantAddress(client *martianpay.Client) {
	fmt.Println("Verifying Merchant Address...")
	fmt.Println("Note: You need to send a test transaction first")
	fmt.Print("Enter Address ID: ")

	var id string
	fmt.Scanln(&id)
	if id == "" {
		id = "ma_example_id"
	}

	req := martianpay.MerchantAddressVerifyRequest{
		ID:     id,
		Amount: "0.01",
	}

	resp, err := client.VerifyMerchantAddress(req)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("✓ Verification Initiated:\n")
	fmt.Printf("  ID: %s\n", resp.ID)
	fmt.Printf("  Status: %s\n", resp.Status)
}

func listMerchantAddresses(client *martianpay.Client) {
	fmt.Println("Listing Merchant Addresses...")

	req := martianpay.MerchantAddressListRequest{
		Page:     0,
		PageSize: 10,
	}

	resp, err := client.ListMerchantAddresses(req)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("✓ Total Addresses: %d\n\n", resp.Total)

	for i, addr := range resp.MerchantAddresses {
		fmt.Printf("  [%d] ID: %s\n", i+1, addr.ID)
		fmt.Printf("      Network: %s, Status: %s\n", addr.Network, addr.Status)
		fmt.Printf("      Address: %s\n", addr.Address)
	}
}

func listMerchantAddressesByNetwork(client *martianpay.Client) {
	fmt.Println("Listing Merchant Addresses by Network...")
	fmt.Print("Enter Network (e.g., Ethereum Sepolia): ")

	scanner := bufio.NewScanner(os.Stdin)
	var network string
	if scanner.Scan() {
		network = scanner.Text()
	}
	if network == "" {
		network = "Ethereum Sepolia"
	}

	req := martianpay.MerchantAddressListRequest{
		Page:     0,
		PageSize: 10,
		Network:  &network,
	}

	resp, err := client.ListMerchantAddresses(req)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("✓ Total Addresses for %s: %d\n\n", network, resp.Total)

	for i, addr := range resp.MerchantAddresses {
		fmt.Printf("  [%d] ID: %s, Address: %s\n", i+1, addr.ID, addr.Address)
	}
}

func deleteMerchantAddress(client *martianpay.Client) {
	fmt.Println("Deleting Merchant Address...")
	fmt.Print("Enter Address ID: ")

	var id string
	fmt.Scanln(&id)
	if id == "" {
		fmt.Println("  Address ID required")
		return
	}

	err := client.DeleteMerchantAddress(id)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("✓ Merchant Address Deleted: %s\n", id)
}

func createAndVerifyMerchantAddress(client *martianpay.Client) {
	fmt.Println("Create and Verify Merchant Address (Full Workflow)...")

	req := martianpay.MerchantAddressCreateRequest{
		Network: "Ethereum Sepolia",
		Address: "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb",
	}

	resp, err := client.CreateMerchantAddress(req)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("✓ Step 1: Address Created\n")
	fmt.Printf("  ID: %s\n", resp.ID)
	fmt.Printf("  Address: %s\n", resp.Address)
	fmt.Printf("\nStep 2: Send a test transaction to this address\n")
	fmt.Printf("Step 3: Call verify API with the transaction amount\n")
	fmt.Println("\n(Skipping verification - requires manual transaction)")
}

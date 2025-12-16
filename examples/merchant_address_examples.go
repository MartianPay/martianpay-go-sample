// Package main provides examples for the MartianPay Merchant Address API.
// Merchant addresses are cryptocurrency wallet addresses used to receive payouts.
// Each address must be verified to ensure the merchant owns it before it can be used.
//
// Verification Process:
// 1. Create a merchant address
// 2. Send a test transaction to the address
// 3. Call the verify API with the transaction amount
// 4. Address is verified and ready for payouts
package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/MartianPay/martianpay-go-sample/pkg/developer"
	martianpay "github.com/MartianPay/martianpay-go-sample/sdk"
)

// Merchant Address Examples

// createMerchantAddress demonstrates creating a new merchant address for receiving payouts.
// The address must be verified before it can be used for withdrawals.
//
// Steps:
// 1. Prompt for network (e.g., "Ethereum Sepolia", "Solana Mainnet")
// 2. Prompt for wallet address
// 3. Create the merchant address
// 4. Display address details and status
//
// Note: Newly created addresses have "unverified" status and must be verified
// before they can receive payouts.
//
// API Endpoints Used:
//   - POST /v1/merchant_addresses
func createMerchantAddress(client *martianpay.Client) {
	fmt.Println("Creating Merchant Address...")

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter Network (e.g., Ethereum Sepolia): ")
	var network string
	if scanner.Scan() {
		network = scanner.Text()
	}
	if network == "" {
		network = "Ethereum Sepolia"
		fmt.Printf("Using default network: %s\n", network)
	}

	fmt.Print("Enter Address: ")
	var address string
	if scanner.Scan() {
		address = scanner.Text()
	}
	if address == "" {
		fmt.Println("Error: Address cannot be empty")
		return
	}

	req := &developer.MerchantAddressCreateRequest{
		Network: network,
		Address: address,
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

// getMerchantAddress retrieves and displays details of a specific merchant address.
//
// Displayed Information:
//   - Address ID
//   - Blockchain network
//   - Wallet address
//   - Verification status
//   - Alias (optional friendly name)
//
// API Endpoints Used:
//   - GET /v1/merchant_addresses/:id
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

// updateMerchantAddress demonstrates updating merchant address metadata.
// You can update the alias (friendly name) for easier identification.
//
// Updateable Fields:
//   - Alias: A friendly name for the address (e.g., "Main Wallet", "Cold Storage")
//
// Note: The actual wallet address and network cannot be changed after creation.
//
// API Endpoints Used:
//   - PUT /v1/merchant_addresses/:id
func updateMerchantAddress(client *martianpay.Client) {
	fmt.Println("Updating Merchant Address...")
	fmt.Print("Enter Address ID: ")

	var id string
	fmt.Scanln(&id)
	if id == "" {
		id = "ma_example_id"
	}

	alias := "My Main Wallet"
	req := &developer.MerchantAddressUpdateRequest{
		Alias: &alias,
	}

	resp, err := client.UpdateMerchantAddress(id, req)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("✓ Merchant Address Updated:\n")
	fmt.Printf("  ID: %s\n", resp.ID)
	fmt.Printf("  Alias: %s\n", resp.Alias)
}

// verifyMerchantAddress demonstrates the address verification process.
// Verification proves that the merchant owns the wallet address.
//
// Verification Steps:
// 1. Create a merchant address (if not already created)
// 2. Send a test transaction to the address from your wallet
// 3. Call this verification API with the exact amount sent
// 4. MartianPay checks the blockchain for the transaction
// 5. If found, the address is marked as "verified"
//
// Important:
//   - The verification amount must match the transaction exactly
//   - The transaction must be confirmed on the blockchain
//   - Use a small amount (e.g., 0.01) for testing
//
// API Endpoints Used:
//   - POST /v1/merchant_addresses/:id/verify
func verifyMerchantAddress(client *martianpay.Client) {
	fmt.Println("Verifying Merchant Address...")
	fmt.Println("Note: You need to send a test transaction first")
	fmt.Print("Enter Address ID: ")

	var id string
	fmt.Scanln(&id)
	if id == "" {
		id = "ma_example_id"
	}

	req := &developer.MerchantAddressVerifyRequest{
		Amount: "0.01",
	}

	resp, err := client.VerifyMerchantAddress(id, req)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("✓ Verification Initiated:\n")
	fmt.Printf("  ID: %s\n", resp.ID)
	fmt.Printf("  Status: %s\n", resp.Status)
}

// listMerchantAddresses retrieves and displays all merchant addresses.
//
// Features:
//   - Pagination support
//   - Displays address ID, network, address, and status
//   - Shows total count
//
// Use Cases:
//   - View all payout addresses
//   - Check verification status
//   - Select an address for payout
//
// API Endpoints Used:
//   - GET /v1/merchant_addresses
func listMerchantAddresses(client *martianpay.Client) {
	fmt.Println("Listing Merchant Addresses...")

	req := &developer.MerchantAddressListRequest{
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

// listMerchantAddressesByNetwork retrieves addresses filtered by blockchain network.
// This is useful when you want to find addresses for a specific network.
//
// Example Networks:
//   - "Ethereum Sepolia"
//   - "Ethereum Mainnet"
//   - "Solana Mainnet"
//   - "Polygon Mainnet"
//
// API Endpoints Used:
//   - GET /v1/merchant_addresses?network=<network_name>
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

	req := &developer.MerchantAddressListRequest{
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

// deleteMerchantAddress demonstrates deleting a merchant address.
//
// Important Restrictions:
//   - Cannot delete an address that has been used in payouts
//   - Only delete addresses that are no longer needed
//   - Deletion is permanent and cannot be undone
//
// API Endpoints Used:
//   - DELETE /v1/merchant_addresses/:id
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

// createAndVerifyMerchantAddress demonstrates the complete workflow for setting up
// a verified merchant address. This is an educational example showing all steps.
//
// Complete Workflow:
// 1. Create merchant address
// 2. Send test transaction (manual step - done outside this function)
// 3. Verify address with transaction amount (would be done separately)
//
// Note: This example only completes step 1 (creation) since steps 2 and 3
// require manual intervention (sending crypto from your wallet).
//
// API Endpoints Used:
//   - POST /v1/merchant_addresses
func createAndVerifyMerchantAddress(client *martianpay.Client) {
	fmt.Println("Create and Verify Merchant Address (Full Workflow)...")

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter Network (e.g., Ethereum Sepolia): ")
	var network string
	if scanner.Scan() {
		network = scanner.Text()
	}
	if network == "" {
		network = "Ethereum Sepolia"
		fmt.Printf("Using default network: %s\n", network)
	}

	fmt.Print("Enter Address: ")
	var address string
	if scanner.Scan() {
		address = scanner.Text()
	}
	if address == "" {
		fmt.Println("Error: Address cannot be empty")
		return
	}

	req := &developer.MerchantAddressCreateRequest{
		Network: network,
		Address: address,
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

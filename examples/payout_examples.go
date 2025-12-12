package main

import (
	"fmt"

	"github.com/MartianPay/martianpay-go-sample/pkg/developer"
	martianpay "github.com/MartianPay/martianpay-go-sample/sdk"
)

// Payout Examples

// previewPayout previews a payout to estimate fees and amounts
func previewPayout(client *martianpay.Client) {
	fmt.Println("Previewing Payout...")

	// First, get merchant addresses to select one
	fmt.Println("  Fetching merchant addresses...")
	addressReq := &developer.MerchantAddressListRequest{
		Page:     0,
		PageSize: 10,
	}
	addressResp, err := client.ListMerchantAddresses(addressReq)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	if len(addressResp.MerchantAddresses) == 0 {
		fmt.Println("✗ No merchant addresses found. Please create one first.")
		return
	}

	// Display available addresses
	fmt.Printf("\n  Available Merchant Addresses:\n")
	for i, addr := range addressResp.MerchantAddresses {
		fmt.Printf("  [%d] ID: %s, Network: %s, Address: %s\n", i+1, addr.ID, addr.Network, addr.Address)
		if addr.Verification != nil {
			fmt.Printf("      Status: %s, AML Status: %s\n", addr.Verification.Status, addr.Verification.AmlStatus)
		}
	}

	fmt.Print("\nEnter address number (or press Enter for first): ")
	var addrChoice string
	fmt.Scanln(&addrChoice)

	selectedIdx := 0
	if addrChoice != "" {
		var idx int
		fmt.Sscanf(addrChoice, "%d", &idx)
		if idx > 0 && idx <= len(addressResp.MerchantAddresses) {
			selectedIdx = idx - 1
		}
	}

	selectedAddr := addressResp.MerchantAddresses[selectedIdx]
	fmt.Printf("  Selected: %s (%s)\n", selectedAddr.ID, selectedAddr.Network)

	// Get asset ID from the selected address verification
	var assetID string
	if selectedAddr.Verification != nil && selectedAddr.Verification.AssetID != "" {
		assetID = selectedAddr.Verification.AssetID
	} else {
		fmt.Print("\nEnter Asset ID (e.g., USDC-Solana-TEST): ")
		fmt.Scanln(&assetID)
	}

	// Get amount
	fmt.Print("Enter source amount: ")
	var amount string
	fmt.Scanln(&amount)
	if amount == "" {
		amount = "15"
	}

	req := &developer.PayoutPreviewRequest{
		PayoutParams: developer.PayoutParams{
			SourceCoin:             "USDC",
			SourceAmount:           amount,
			DestinationAssetId:     assetID,
			DestinationAccountType: "wallet",
			DestinationAccountId:   selectedAddr.ID,
		},
	}

	response, err := client.PreviewPayout(req)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("\n✓ Payout Preview:\n")
	fmt.Printf("  ID: %s\n", response.ID)
	fmt.Printf("  Source Amount: %s %s\n", response.SourceAmount, response.SourceCoin)
	fmt.Printf("  Receive Amount: %s %s\n", response.ReceiveAmount, response.ReceiveCoin)
	fmt.Printf("  Exchange Rate: %s\n", response.ExchangeRate)
	fmt.Printf("  Network Fee: %s\n", response.PaymentNetworkFee)
	fmt.Printf("  Service Fee: %s\n", response.PaymentServiceFee)
	fmt.Printf("  Total Fee: %s\n", response.PaymentTotalFee)
	fmt.Printf("  Net Amount: %s\n", response.PaymentNetAmount)
	fmt.Printf("  Status: %s\n", response.Status)
	fmt.Printf("  Approval Status: %s\n", response.ApprovalStatus)

	if response.ReceiveWalletAddress != nil {
		fmt.Printf("\n  Destination:\n")
		fmt.Printf("    Network: %s\n", response.ReceiveWalletAddress.Network)
		fmt.Printf("    Address: %s\n", response.ReceiveWalletAddress.Address)
	}

	if len(response.SwapItems) > 0 {
		fmt.Printf("\n  Swap Items:\n")
		for i, swap := range response.SwapItems {
			fmt.Printf("  [%d] From: %s %s -> To: %s %s\n",
				i+1,
				swap.EstimatedFromAmount,
				swap.FromAssetId,
				swap.EstimatedToAmount,
				swap.ToAssetId)
			fmt.Printf("      Quote ID: %s\n", swap.QuoteId)
		}
	}
}

// createPayout creates a new payout request
func createPayout(client *martianpay.Client) {
	fmt.Println("Creating Payout...")

	// First, preview to get quote IDs
	fmt.Println("  Step 1: Preview payout to get quote IDs...")

	// Get merchant addresses
	addressReq := &developer.MerchantAddressListRequest{
		Page:     0,
		PageSize: 10,
	}
	addressResp, err := client.ListMerchantAddresses(addressReq)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	if len(addressResp.MerchantAddresses) == 0 {
		fmt.Println("✗ No merchant addresses found. Please create one first.")
		return
	}

	// Display available addresses
	fmt.Printf("\n  Available Merchant Addresses:\n")
	for i, addr := range addressResp.MerchantAddresses {
		fmt.Printf("  [%d] ID: %s, Network: %s, Address: %s\n", i+1, addr.ID, addr.Network, addr.Address)
	}

	fmt.Print("\nEnter address number (or press Enter for first): ")
	var addrChoice string
	fmt.Scanln(&addrChoice)

	selectedIdx := 0
	if addrChoice != "" {
		var idx int
		fmt.Sscanf(addrChoice, "%d", &idx)
		if idx > 0 && idx <= len(addressResp.MerchantAddresses) {
			selectedIdx = idx - 1
		}
	}

	selectedAddr := addressResp.MerchantAddresses[selectedIdx]

	// Get asset ID
	var assetID string
	if selectedAddr.Verification != nil && selectedAddr.Verification.AssetID != "" {
		assetID = selectedAddr.Verification.AssetID
	} else {
		fmt.Print("\nEnter Asset ID (e.g., USDC-Solana-TEST): ")
		fmt.Scanln(&assetID)
	}

	// Get amount
	fmt.Print("Enter source amount: ")
	var amount string
	fmt.Scanln(&amount)
	if amount == "" {
		amount = "15"
	}

	// Preview first
	previewReq := &developer.PayoutPreviewRequest{
		PayoutParams: developer.PayoutParams{
			SourceCoin:             "USDC",
			SourceAmount:           amount,
			DestinationAssetId:     assetID,
			DestinationAccountType: "wallet",
			DestinationAccountId:   selectedAddr.ID,
		},
	}

	previewResp, err := client.PreviewPayout(previewReq)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	// Extract quote IDs from swap items
	var quoteIDs []string
	if len(previewResp.SwapItems) > 0 {
		fmt.Printf("  Found %d swap items\n", len(previewResp.SwapItems))
		for _, swap := range previewResp.SwapItems {
			quoteIDs = append(quoteIDs, swap.QuoteId)
		}
	}

	// Create payout
	fmt.Println("\n  Step 2: Creating payout...")
	createReq := &developer.PayoutCreateRequest{
		PayoutParams: developer.PayoutParams{
			SourceCoin:             "USDC",
			SourceAmount:           amount,
			DestinationAssetId:     assetID,
			DestinationAccountType: "wallet",
			DestinationAccountId:   selectedAddr.ID,
			QuoteIds:               quoteIDs,
		},
	}

	response, err := client.CreatePayout(createReq)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("\n✓ Payout Created:\n")
	fmt.Printf("  ID: %s\n", response.ID)
	fmt.Printf("  Source Amount: %s %s\n", response.SourceAmount, response.SourceCoin)
	fmt.Printf("  Receive Amount: %s %s\n", response.ReceiveAmount, response.ReceiveCoin)
	fmt.Printf("  Network Fee: %s\n", response.PaymentNetworkFee)
	fmt.Printf("  Total Fee: %s\n", response.PaymentTotalFee)
	fmt.Printf("  Status: %s\n", response.Status)
	fmt.Printf("  Approval Status: %s\n", response.ApprovalStatus)

	fmt.Printf("\n  Note: If amount > $200 USD, approval is required.\n")
	fmt.Printf("        Use the approval functions to approve/reject this payout.\n")
}

// getPayout retrieves details of a specific payout
func getPayout(client *martianpay.Client) {
	fmt.Println("Getting Payout...")
	fmt.Print("Enter Payout ID: ")

	var id string
	fmt.Scanln(&id)
	if id == "" {
		id = "payout_example_id"
	}

	response, err := client.GetPayout(id)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("\n✓ Payout Retrieved:\n")
	fmt.Printf("  ID: %s\n", response.ID)
	fmt.Printf("  Source Amount: %s %s\n", response.SourceAmount, response.SourceCoin)
	fmt.Printf("  Receive Amount: %s %s\n", response.ReceiveAmount, response.ReceiveCoin)
	fmt.Printf("  Exchange Rate: %s\n", response.ExchangeRate)
	fmt.Printf("  Network Fee: %s\n", response.PaymentNetworkFee)
	fmt.Printf("  Total Fee: %s\n", response.PaymentTotalFee)
	fmt.Printf("  Status: %s\n", response.Status)
	fmt.Printf("  Approval Status: %s\n", response.ApprovalStatus)
	fmt.Printf("  Created: %d\n", response.Created)
	fmt.Printf("  Updated: %d\n", response.Updated)

	if response.ReceiveWalletAddress != nil {
		fmt.Printf("\n  Destination:\n")
		fmt.Printf("    Network: %s\n", response.ReceiveWalletAddress.Network)
		fmt.Printf("    Address: %s\n", response.ReceiveWalletAddress.Address)
	}
}

// listPayouts retrieves a list of payouts
func listPayouts(client *martianpay.Client) {
	fmt.Println("Listing Payouts...")

	req := &developer.PayoutListRequest{
		Page:     0,
		PageSize: 10,
	}

	response, err := client.ListPayouts(req)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("✓ Total Payouts: %d\n", response.Total)
	fmt.Printf("  Page: %d, Page Size: %d\n\n", response.Page, response.PageSize)

	for i, payout := range response.Payouts {
		fmt.Printf("  [%d] ID: %s\n", i+1, payout.ID)
		fmt.Printf("      Amount: %s %s -> %s %s\n",
			payout.SourceAmount, payout.SourceCoin,
			payout.ReceiveAmount, payout.ReceiveCoin)
		fmt.Printf("      Status: %s, Approval: %s\n", payout.Status, payout.ApprovalStatus)
		if payout.ReceiveWalletAddress != nil {
			fmt.Printf("      To: %s (%s)\n", payout.ReceiveWalletAddress.Address, payout.ReceiveWalletAddress.Network)
		}
		fmt.Println()
	}
}

// getPayoutApprovalInstance retrieves the approval instance for a payout
func getPayoutApprovalInstance(client *martianpay.Client) {
	fmt.Println("Getting Payout Approval Instance...")
	fmt.Print("Enter Payout ID: ")

	var payoutID string
	fmt.Scanln(&payoutID)
	if payoutID == "" {
		payoutID = "payout_example_id"
	}

	response, err := client.GetApprovalInstance(payoutID)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("\n✓ Approval Instance Retrieved:\n")
	fmt.Printf("  ID: %s\n", response.ID)
	fmt.Printf("  Resource ID: %s\n", response.ResourceID)
	fmt.Printf("  Resource Type: %s\n", response.ResourceType)
	fmt.Printf("  Status: %s\n", response.Status)
	fmt.Printf("  Created: %d\n", response.CreatedAt)
	fmt.Printf("  Updated: %d\n\n", response.UpdatedAt)

	if len(response.Records) > 0 {
		fmt.Printf("  Approval Records:\n")
		for i, record := range response.Records {
			fmt.Printf("  [%d] Action: %s\n", i+1, record.Action)
			fmt.Printf("      Approver: %s (%s)\n", record.ApproverName, record.ApproverRole)
			fmt.Printf("      Comment: %s\n", record.Comment)
			fmt.Printf("      Namespace: %s\n", record.Namespace)
		}
	}

	if response.CurrentStep != nil {
		fmt.Printf("\n  Current Step:\n")
		fmt.Printf("    Step Order: %d\n", response.CurrentStep.StepOrder)
		fmt.Printf("    Namespace: %s\n", response.CurrentStep.Namespace)
		fmt.Printf("    Roles: %v\n", response.CurrentStep.Roles)
	}
}

// approvePayoutRequest approves a payout
func approvePayoutRequest(client *martianpay.Client) {
	fmt.Println("Approving Payout...")
	fmt.Print("Enter Approval Instance ID: ")

	var approvalID string
	fmt.Scanln(&approvalID)
	if approvalID == "" {
		approvalID = "approval_instance_example_id"
	}

	fmt.Print("Enter comment (optional): ")
	var comment string
	fmt.Scanln(&comment)
	if comment == "" {
		comment = "Approved via SDK example"
	}

	err := client.ApprovePayout(approvalID, comment)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("✓ Payout Approved Successfully\n")
	fmt.Printf("  Approval Instance ID: %s\n", approvalID)
	fmt.Printf("  Comment: %s\n", comment)
}

// rejectPayoutRequest rejects a payout
func rejectPayoutRequest(client *martianpay.Client) {
	fmt.Println("Rejecting Payout...")
	fmt.Print("Enter Approval Instance ID: ")

	var approvalID string
	fmt.Scanln(&approvalID)
	if approvalID == "" {
		approvalID = "approval_instance_example_id"
	}

	fmt.Print("Enter rejection reason: ")
	var reason string
	fmt.Scanln(&reason)
	if reason == "" {
		reason = "Rejected via SDK example"
	}

	err := client.RejectPayout(approvalID, reason)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("✓ Payout Rejected Successfully\n")
	fmt.Printf("  Approval Instance ID: %s\n", approvalID)
	fmt.Printf("  Reason: %s\n", reason)
}

// cancelPayout cancels a pending payout
func cancelPayout(client *martianpay.Client) {
	fmt.Println("Canceling Payout...")
	fmt.Print("Enter Payout ID: ")

	var payoutID string
	fmt.Scanln(&payoutID)
	if payoutID == "" {
		payoutID = "payout_example_id"
	}

	response, err := client.CancelPayout(payoutID)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("\n✓ Payout Canceled:\n")
	fmt.Printf("  ID: %s\n", response.ID)
	fmt.Printf("  Status: %s\n", response.Status)
	fmt.Printf("  Source Amount: %s %s\n", response.SourceAmount, response.SourceCoin)

	fmt.Printf("\n  Note: Cancellation windows:\n")
	fmt.Printf("    ≤ $200: Can cancel within 10 minutes after creation\n")
	fmt.Printf("    $200-$2,000: Can cancel before merchant approval\n")
	fmt.Printf("    > $2,000: Can cancel before all approvals complete\n")
}


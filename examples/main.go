package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	martianpay "github.com/MartianPay/martianpay-go-sample/sdk"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Get API Key from user
	fmt.Println("===========================================")
	fmt.Println("  MartianPay Go SDK Examples")
	fmt.Println("===========================================")
	fmt.Print("\nEnter your API key (or press Enter to use default): ")

	var apiKeyInput string
	if scanner.Scan() {
		apiKeyInput = strings.TrimSpace(scanner.Text())
	}

	// Use input API key or fall back to default
	if apiKeyInput == "" {
		apiKeyInput = apiKey
		fmt.Println("Using default API key from common.go")
	} else {
		fmt.Println("Using custom API key")
	}

	// Store API key globally for all examples
	currentAPIKey = apiKeyInput
	fmt.Println()

	for {
		showMainMenu()
		fmt.Print("\nEnter your choice (0 to exit): ")

		if !scanner.Scan() {
			break
		}

		choice := strings.TrimSpace(scanner.Text())
		if choice == "0" {
			fmt.Println("Goodbye!")
			break
		}

		num, err := strconv.Atoi(choice)
		if err != nil || num < 0 || num > 12 {
			fmt.Println("Invalid choice. Please try again.")
			continue
		}

		// Show submenu based on category
		if num > 0 {
			handleCategory(num, scanner)
		}
	}
}

func showMainMenu() {
	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Println("MartianPay SDK Examples - Main Menu")
	fmt.Println(strings.Repeat("=", 80))

	fmt.Println("\n1. Payment Intent Examples")
	fmt.Println("2. Customer Examples")
	fmt.Println("3. Refund Examples")
	fmt.Println("4. Payroll Examples")
	fmt.Println("5. Merchant Address Examples")
	fmt.Println("6. Payout Examples")
	fmt.Println("7. Assets Examples")
	fmt.Println("8. Balance Examples")
	fmt.Println("9. Product Examples")
	fmt.Println("10. Payment Link Examples")
	fmt.Println("11. Subscription Examples")
	fmt.Println("12. Webhook Examples")

	fmt.Println("\n0. Exit")
}

func handleCategory(category int, scanner *bufio.Scanner) {
	for {
		fmt.Println("\n" + strings.Repeat("=", 80))

		var menuItems []string
		var categoryName string

		switch category {
		case 1:
			categoryName = "Payment Intent Examples"
			menuItems = []string{
				"Create and Update Payment Intent (Crypto)",
				"Get Payment Intent",
				"List Payment Intents",
				"Cancel Payment Intent",
				"List Payment Intents with Permanent Deposit",
				"Fiat Payment with New Card",
				"Fiat Payment with Saved Card",
			"Create Payment Intent with Payment Link",
			}
		case 2:
			categoryName = "Customer Examples"
			menuItems = []string{
				"Create and Update Customer",
				"Get Customer",
				"List Customers",
				"Delete Customer",
				"List Customer Payment Methods",
			}
		case 3:
			categoryName = "Refund Examples"
			menuItems = []string{
				"Create Refund",
				"Get Refund",
				"List Refunds",
			}
		case 4:
			categoryName = "Payroll Examples"
			menuItems = []string{
				"Create Direct Payroll (Normal Payment)",
				"Create Direct Payroll (Binance Payment)",
				"Get Payroll",
				"List Payrolls",
				"List Payroll Items",
			}
		case 5:
			categoryName = "Merchant Address Examples"
			menuItems = []string{
				"Create Merchant Address",
				"Get Merchant Address",
				"Update Merchant Address",
				"Verify Merchant Address",
				"List Merchant Addresses",
				"List Merchant Addresses by Network",
				"Delete Merchant Address",
				"Create and Verify Merchant Address (Full Workflow)",
			}
		case 6:
			categoryName = "Payout Examples"
			menuItems = []string{
				"Preview Payout",
				"Create Payout",
				"Get Payout",
				"List Payouts",
				"Get Payout Approval Instance",
				"Approve Payout",
				"Reject Payout",
				"Cancel Payout",
			}
		case 7:
			categoryName = "Assets Examples"
			menuItems = []string{
				"List Enabled Assets",
				"Get All Available Assets",
				"List Asset Network Fees",
				"Show Crypto Assets by Network",
				"Show Payable Assets Only",
				"Compare Mainnet vs Testnet Assets",
			}
		case 8:
			categoryName = "Balance Examples"
			menuItems = []string{
				"Show Balance Summary",
				"Show Balance by Currency",
				"Show Available Balances Only",
				"Show Locked and Pending Balances",
				"Compare Balance Types",
			}
		case 9:
			categoryName = "Product Examples"
			menuItems = []string{
				"List Products",
				"Create Product with Variants",
			"Create Product with Selling Plan",
				"Get Product",
				"Update Product",
				"Delete Product",
				"List Active Products",
				"List Selling Plan Groups",
				"Create Selling Plan Group",
				"Get Selling Plan Group",
				"Update Selling Plan Group",
				"Delete Selling Plan Group",
				"List Selling Plans",
				"Create Selling Plan",
				"Get Selling Plan",
				"Update Selling Plan",
				"Delete Selling Plan",
				"Calculate Selling Plan Price",
			}
		case 10:
			categoryName = "Payment Link Examples"
			menuItems = []string{
				"List Payment Links",
				"Create Payment Link",
				"Get Payment Link",
				"Update Payment Link",
				"Delete Payment Link",
				"List Active Payment Links",
				"List Payment Links by Product",
			}
		case 11:
			categoryName = "Subscription Examples"
			menuItems = []string{
				"List Subscriptions",
				"List Subscriptions by Customer",
				"List Subscriptions by Status",
				"Get Subscription",
				"Cancel Subscription at Period End",
				"Cancel Subscription Immediately",
				"Pause Subscription",
				"Pause Subscription with Auto-Resume",
				"Resume Subscription",
			}
		case 12:
			categoryName = "Webhook Examples"
			menuItems = []string{
				"Start Webhook Event Receiver Server",
			}
		}

		fmt.Printf("%s\n", categoryName)
		fmt.Println(strings.Repeat("=", 80))

		for i, item := range menuItems {
			fmt.Printf("%2d. %s\n", i+1, item)
		}
		fmt.Println("\n 0. Back to Main Menu")

		fmt.Print("\nEnter your choice: ")
		if !scanner.Scan() {
			return
		}

		choice := strings.TrimSpace(scanner.Text())
		if choice == "0" {
			return
		}

		num, err := strconv.Atoi(choice)
		if err != nil || num < 0 || num > len(menuItems) {
			fmt.Println("Invalid choice. Please try again.")
			continue
		}

		// Map to original example number
		exampleNum := getExampleNumber(category, num)

		fmt.Println("\n" + strings.Repeat("=", 80))
		runExample(exampleNum)
		fmt.Println(strings.Repeat("=", 80))
		fmt.Print("\nPress Enter to continue...")
		scanner.Scan()
	}
}

func getExampleNumber(category, choice int) int {
	categoryOffsets := map[int]int{
		1:  0,  // Payment Intent: 1-8
		2:  8,  // Customer: 9-13
		3:  13, // Refund: 14-16
		4:  16, // Payroll: 17-21
		5:  21, // Merchant Address: 22-29
		6:  29, // Payout: 30-37
		7:  37, // Assets: 38-43
		8:  43, // Balance: 44-48
		9:  48, // Product: 49-66 (7 product + 11 selling plan)
		10: 66, // Payment Link: 67-73
		11: 73, // Subscription: 74-82 (9 examples)
		12: 82, // Webhook: 83
	}

	return categoryOffsets[category] + choice
}

func runExample(choice int) {
	fmt.Printf("Using API Key: %s\n", currentAPIKey)
	client := martianpay.NewClient(currentAPIKey)

	switch choice {
	case 1:
		createAndUpdatePaymentIntent(client)
	case 2:
		getPaymentIntent(client)
	case 3:
		listPaymentIntents(client)
	case 4:
		cancelPaymentIntent(client)
	case 5:
		listPaymentIntentsWithPermanentDeposit(client)
	case 6:
		fiatPaymentWithNewCard(client)
	case 7:
		fiatPaymentWithSavedCard(client)
	case 8:
		createPaymentIntentWithPaymentLink(client)
	case 9:
		createAndUpdateCustomer(client)
	case 10:
		getCustomer(client)
	case 11:
		listCustomers(client)
	case 12:
		deleteCustomer(client)
	case 13:
		listCustomerPaymentMethods(client)
	case 14:
		createRefund(client)
	case 15:
		getRefund(client)
	case 16:
		listRefunds(client)
	case 17:
		createDirectPayroll(client)
	case 18:
		createDirectPayrollBinance(client)
	case 19:
		getPayroll(client)
	case 20:
		listPayrolls(client)
	case 21:
		listPayrollItems(client)
	case 22:
		createMerchantAddress(client)
	case 23:
		getMerchantAddress(client)
	case 24:
		updateMerchantAddress(client)
	case 25:
		verifyMerchantAddress(client)
	case 26:
		listMerchantAddresses(client)
	case 27:
		listMerchantAddressesByNetwork(client)
	case 28:
		deleteMerchantAddress(client)
	case 29:
		createAndVerifyMerchantAddress(client)
	case 30:
		previewPayout(client)
	case 31:
		createPayout(client)
	case 32:
		getPayout(client)
	case 33:
		listPayouts(client)
	case 34:
		getPayoutApprovalInstance(client)
	case 35:
		approvePayoutRequest(client)
	case 36:
		rejectPayoutRequest(client)
	case 37:
		cancelPayout(client)
	case 38:
		listAssets(client)
	case 39:
		getAllAssets(client)
	case 40:
		listAssetFees(client)
	case 41:
		showCryptoAssetsByNetwork(client)
	case 42:
		showPayableAssets(client)
	case 43:
		compareMainnetVsTestnet(client)
	case 44:
		showBalance(client)
	case 45:
		showBalanceByCurrency(client)
	case 46:
		showAvailableBalancesOnly(client)
	case 47:
		showLockedAndPendingBalances(client)
	case 48:
		compareBalanceTypes(client)
	case 49:
		listProducts(client)
	case 50:
		createProductWithVariants(client)
	case 51:
		createProductWithSellingPlan(client)
	case 52:
		getProduct(client)
	case 53:
		updateProduct(client)
	case 54:
		deleteProduct(client)
	case 55:
		listActiveProducts(client)
	case 56:
		listSellingPlanGroups(client)
	case 57:
		createSellingPlanGroup(client)
	case 58:
		getSellingPlanGroup(client)
	case 59:
		updateSellingPlanGroup(client)
	case 60:
		deleteSellingPlanGroup(client)
	case 61:
		listSellingPlans(client)
	case 62:
		createSellingPlan(client)
	case 63:
		getSellingPlan(client)
	case 64:
		updateSellingPlan(client)
	case 65:
		deleteSellingPlan(client)
	case 66:
		calculateSellingPlanPrice(client)
	case 67:
		listPaymentLinks(client)
	case 68:
		createPaymentLink(client)
	case 69:
		getPaymentLink(client)
	case 70:
		updatePaymentLink(client)
	case 71:
		deletePaymentLink(client)
	case 72:
		listActivePaymentLinks(client)
	case 73:
		listPaymentLinksByProduct(client)
	case 74:
		listSubscriptions(client)
	case 75:
		listSubscriptionsByCustomer(client)
	case 76:
		listSubscriptionsByStatus(client)
	case 77:
		getSubscription(client)
	case 78:
		cancelSubscriptionAtPeriodEnd(client)
	case 79:
		cancelSubscriptionImmediately(client)
	case 80:
		pauseSubscription(client)
	case 81:
		pauseSubscriptionWithAutoResume(client)
	case 82:
		resumeSubscription(client)
	case 83:
		startWebhookServer()
	default:
		fmt.Println("Invalid choice")
	}
}

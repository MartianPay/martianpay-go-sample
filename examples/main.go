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
		if err != nil || num < 0 || num > 6 {
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
	fmt.Println("6. Webhook Examples")

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
				"Create Direct Payroll",
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
		1: 0,  // Payment Intent: 1-7
		2: 7,  // Customer: 8-12
		3: 12, // Refund: 13-15
		4: 15, // Payroll: 16-19
		5: 19, // Merchant Address: 20-27
		6: 27, // Webhook: 28
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
		createAndUpdateCustomer(client)
	case 9:
		getCustomer(client)
	case 10:
		listCustomers(client)
	case 11:
		deleteCustomer(client)
	case 12:
		listCustomerPaymentMethods(client)
	case 13:
		createRefund(client)
	case 14:
		getRefund(client)
	case 15:
		listRefunds(client)
	case 16:
		createDirectPayroll(client)
	case 17:
		getPayroll(client)
	case 18:
		listPayrolls(client)
	case 19:
		listPayrollItems(client)
	case 20:
		createMerchantAddress(client)
	case 21:
		getMerchantAddress(client)
	case 22:
		updateMerchantAddress(client)
	case 23:
		verifyMerchantAddress(client)
	case 24:
		listMerchantAddresses(client)
	case 25:
		listMerchantAddressesByNetwork(client)
	case 26:
		deleteMerchantAddress(client)
	case 27:
		createAndVerifyMerchantAddress(client)
	case 28:
		startWebhookServer()
	default:
		fmt.Println("Invalid choice")
	}
}

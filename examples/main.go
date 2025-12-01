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

	for {
		showMenu()
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
		if err != nil || num < 0 || num > 28 {
			fmt.Println("Invalid choice. Please try again.")
			continue
		}

		fmt.Println("\n" + strings.Repeat("=", 80))
		runExample(num)
		fmt.Println(strings.Repeat("=", 80))
		fmt.Print("\nPress Enter to continue...")
		scanner.Scan()
	}
}

func showMenu() {
	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Println("MartianPay SDK Examples")
	fmt.Println(strings.Repeat("=", 80))

	fmt.Println("\n=== Payment Intent Examples ===")
	fmt.Println(" 1. Create and Update Payment Intent (Crypto)")
	fmt.Println(" 2. Get Payment Intent")
	fmt.Println(" 3. List Payment Intents")
	fmt.Println(" 4. Cancel Payment Intent")
	fmt.Println(" 5. List Payment Intents with Permanent Deposit")
	fmt.Println(" 6. Fiat Payment with New Card")
	fmt.Println(" 7. Fiat Payment with Saved Card")

	fmt.Println("\n=== Customer Examples ===")
	fmt.Println(" 8. Create and Update Customer")
	fmt.Println(" 9. Get Customer")
	fmt.Println("10. List Customers")
	fmt.Println("11. Delete Customer")
	fmt.Println("12. List Customer Payment Methods")

	fmt.Println("\n=== Refund Examples ===")
	fmt.Println("13. Create Refund")
	fmt.Println("14. Get Refund")
	fmt.Println("15. List Refunds")

	fmt.Println("\n=== Payroll Examples ===")
	fmt.Println("16. Create Direct Payroll")
	fmt.Println("17. Get Payroll")
	fmt.Println("18. List Payrolls")
	fmt.Println("19. List Payroll Items")

	fmt.Println("\n=== Merchant Address Examples ===")
	fmt.Println("20. Create Merchant Address")
	fmt.Println("21. Get Merchant Address")
	fmt.Println("22. Update Merchant Address")
	fmt.Println("23. Verify Merchant Address")
	fmt.Println("24. List Merchant Addresses")
	fmt.Println("25. List Merchant Addresses by Network")
	fmt.Println("26. Delete Merchant Address")
	fmt.Println("27. Create and Verify Merchant Address (Full Workflow)")

	fmt.Println("\n=== Webhook Examples ===")
	fmt.Println("28. Start Webhook Event Receiver Server")

	fmt.Println("\n 0. Exit")
}

func runExample(choice int) {
	client := martianpay.NewClient(apiKey)

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

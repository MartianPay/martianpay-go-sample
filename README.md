# MartianPay Go SDK

Official Go SDK for the MartianPay payment platform.

> **🚀 Quick Start**: Check out the [Interactive Examples](#interactive-examples) to get started quickly!

## Features

- **Payment Intents**: Create, update, retrieve, list, and cancel payment intents
- **Customers**: Create, update, retrieve, list, and delete customers
- **Payment Methods**: List customer's saved payment methods (cards)
- **Refunds**: Create, retrieve, and list refunds
- **Payroll**: Create direct payroll, retrieve, and list payrolls and payroll items
- **Merchant Addresses (Wallets)**: Add, verify, update, list, and delete blockchain addresses for withdrawals
- **Assets**: List all available crypto and fiat assets with network details
- **Balance**: Query merchant balance across different currencies and assets
- **Crypto Payments**: Support for crypto payment methods (USDT, USDC, ETH, etc.)
- **Fiat/Card Payments**: Support for card payments via Stripe (new card and saved card)
- **Webhook Events**: Receive and verify webhook events

## Installation

```sh
go get github.com/MartianPay/martianpay-go-sample
```

## Integration Approaches

MartianPay offers flexible integration options to suit different use cases. For detailed integration guides including frontend code examples, see the [examples directory](examples/README.md).

## Interactive Examples

The fastest way to learn the SDK! We provide an interactive menu-driven interface with examples covering all features:

```bash
cd examples
make update   # Update to latest SDK version
make run      # Run interactive examples menu
```

**Features:**
- 📋 Two-level menu system organized by feature category
- 🎲 Automatic randomization of emails and order IDs to avoid duplicates
- 🎯 User-friendly prompts for interactive input
- ✅ Comprehensive coverage of all SDK methods

The examples demonstrate the **API-only integration approach** to show all SDK methods. For production, we recommend using the **MartianPay.js Widget** for simpler integration.

📁 **See [examples/README.md](examples/README.md) for full details**

## Quick Start

Here's a simple example of using the SDK to list customers:

```go
package main

import (
	"fmt"
	"log"

	martianpay "github.com/MartianPay/martianpay-go-sample/sdk"
)

func main() {
	// Initialize the MartianPay client with your API key
	apiKey := "your_api_key_here" // Replace with your actual API key
	client := martianpay.NewClient(apiKey)

	// Create request to list customers
	req := martianpay.CustomerListRequest{
		Page:     0,  // Start from page 0
		PageSize: 10, // Get 10 customers per page
	}

	// Call the ListCustomers API
	resp, err := client.ListCustomers(req)
	if err != nil {
		log.Fatalf("Failed to list customers: %v", err)
	}

	// Display results
	fmt.Printf("Total customers: %d\n", resp.Total)
	fmt.Printf("Customers on this page: %d\n\n", len(resp.Customers))

	// Print each customer
	for i, customer := range resp.Customers {
		fmt.Printf("Customer %d:\n", i+1)
		fmt.Printf("  ID: %s\n", customer.ID)
		if customer.Email != nil {
			fmt.Printf("  Email: %s\n", *customer.Email)
		}
		if customer.Name != nil {
			fmt.Printf("  Name: %s\n", *customer.Name)
		}
		fmt.Printf("  Total Payment: %.2f %s\n", customer.TotalPayment, customer.Currency)
		fmt.Println()
	}
}
```

## Testing the SDK

All SDK functionality can be tested through the interactive examples:

```bash
cd examples
make run
```

Select from organized categories:
1. **Payment Intent Examples** - Create, update, list, cancel payment intents
2. **Customer Examples** - Manage customers and payment methods
3. **Refund Examples** - Process and manage refunds
4. **Payroll Examples** - Create and manage crypto payrolls
5. **Merchant Address Examples** - Add and verify withdrawal addresses
6. **Webhook Examples** - Test webhook event handling

## Keeping SDK Up to Date

To ensure you're using the latest features and bug fixes:

```bash
# Update to the latest SDK version
go get -u github.com/MartianPay/martianpay-go-sample

# Update all dependencies
go mod tidy
```

> **💡 Tip**: Run `go get -u github.com/MartianPay/martianpay-go-sample` periodically to get the latest features, improvements, and bug fixes.

### Recent Improvements

Latest updates include:
- ✅ Two-level interactive menu for better navigation
- ✅ Automatic randomization of emails and order IDs to prevent duplicates
- ✅ User input prompts for flexible testing (addresses, amounts, networks)
- ✅ Enhanced error handling with `error_code` field support
- ✅ Support for both `form` and `json` tags in query parameters
- ✅ Fixed pointer type handling in reflection-based parameter parsing
- ✅ Comprehensive integration documentation and examples

## Documentation & Resources

- 📖 [Interactive Examples](examples/README.md) - Menu-driven examples for all features
- 📖 [MartianPay.js Widget Guide](https://docs.martianpay.com/v1/docs/martianpay-js-usage) - Recommended integration method
- 📖 [API Reference](https://docs.martianpay.com) - Full API documentation
- 🏠 [MartianPay Dashboard](https://dashboard.martianpay.com) - Get your API key

## Support

Need help? Here are your options:

- 📚 Check the [examples directory](examples/) for code samples
- 📖 Read the [API documentation](https://docs.martianpay.com)
- 💬 Contact support through the [MartianPay Dashboard](https://dashboard.martianpay.com)
// Package main provides examples for the MartianPay Balance API.
// The Balance API allows merchants to view their account balances across
// multiple currencies and cryptocurrencies.
//
// Balance Types:
//   - Available: Funds ready for immediate use (payouts, refunds)
//   - Pending: Unreconciled deposits awaiting confirmation
//   - Locked: Funds reserved for pending payouts or held transactions
//   - Frozen: Funds restricted due to compliance or security holds
//   - Total: Sum of all balance types
package main

import (
	"fmt"
	"strings"

	martianpay "github.com/MartianPay/martianpay-go-sample/sdk"
)

// Balance Examples

// showBalance displays a comprehensive summary of the merchant's balance.
// This includes total balances and detailed breakdowns by currency/asset.
//
// Displayed Information:
//   - Primary currency totals (Available, Pending, Locked, Frozen, Total)
//   - Detailed balance for each currency/asset
//   - Explanation of balance types
//
// Use Cases:
//   - Check available funds before creating payouts
//   - Monitor pending deposits
//   - Review locked funds from pending transactions
//   - Financial reporting and reconciliation
//
// API Endpoints Used:
//   - GET /v1/balance
func showBalance(client *martianpay.Client) {
	fmt.Println("Getting Merchant Balance...")

	balance, err := client.GetBalance()
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Println("Merchant Balance Summary")
	fmt.Println(strings.Repeat("=", 80))

	fmt.Printf("\nPrimary Currency: %s\n\n", balance.Currency)
	fmt.Printf("  Available Balance:  %s\n", balance.AvailableBalance)
	fmt.Printf("  Pending Balance:    %s\n", balance.PendingBalance)
	fmt.Printf("  Locked Balance:     %s\n", balance.LockedBalance)
	fmt.Printf("  Frozen Balance:     %s\n", balance.FrozenBalance)
	fmt.Printf("  ───────────────────────────\n")
	fmt.Printf("  Total Balance:      %s\n", balance.TotalBalance)

	if len(balance.BalanceDetails) > 0 {
		fmt.Printf("\n%s\n", strings.Repeat("=", 80))
		fmt.Printf("Balance Details by Currency/Asset (%d)\n", len(balance.BalanceDetails))
		fmt.Printf("%s\n", strings.Repeat("=", 80))

		for i, detail := range balance.BalanceDetails {
			fmt.Printf("\n[%d] %s\n", i+1, detail.Currency)
			fmt.Printf("    Available:      %s\n", detail.AvailableBalance)
			fmt.Printf("    Pending:        %s\n", detail.PendingBalance)
			fmt.Printf("    Locked:         %s\n", detail.LockedBalance)
			fmt.Printf("    Frozen:         %s\n", detail.FrozenBalance)
			fmt.Printf("    Total:          %s\n", detail.TotalBalance)
		}
	}

	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Println("\nBalance Types Explained:")
	fmt.Println("  • Available: Funds ready for immediate use in transactions")
	fmt.Println("  • Pending:   Unreconciled deposits awaiting confirmation")
	fmt.Println("  • Locked:    Funds reserved for pending payouts")
	fmt.Println("  • Frozen:    Funds restricted due to compliance/security")
	fmt.Println(strings.Repeat("=", 80))
}

// showBalanceByCurrency displays balances organized by currency type (crypto vs fiat).
// This helps merchants understand their balance composition across different asset types.
//
// Features:
//   - Separates cryptocurrency and fiat balances
//   - Shows detailed breakdown for each asset
//   - Displays all balance types (Available, Pending, Locked, Frozen)
//
// API Endpoints Used:
//   - GET /v1/balance
func showBalanceByCurrency(client *martianpay.Client) {
	fmt.Println("Showing Balance by Currency...")

	balance, err := client.GetBalance()
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	if len(balance.BalanceDetails) == 0 {
		fmt.Println("✓ No balance details found")
		return
	}

	// Group by currency type (detect crypto vs fiat by currency name pattern)
	cryptoBalances := make(map[string]*BalanceInfo)
	fiatBalances := make(map[string]*BalanceInfo)

	for _, detail := range balance.BalanceDetails {
		info := &BalanceInfo{
			Currency:  detail.Currency,
			Available: detail.AvailableBalance,
			Pending:   detail.PendingBalance,
			Locked:    detail.LockedBalance,
			Frozen:    detail.FrozenBalance,
			Total:     detail.TotalBalance,
		}

		// Detect if it's crypto (contains "-" separator like "USDC-Ethereum-TEST") or fiat
		if strings.Contains(detail.Currency, "-") {
			cryptoBalances[detail.Currency] = info
		} else {
			fiatBalances[detail.Currency] = info
		}
	}

	// Display crypto balances
	if len(cryptoBalances) > 0 {
		fmt.Printf("\n%s\n", strings.Repeat("=", 80))
		fmt.Printf("Cryptocurrency Balances (%d)\n", len(cryptoBalances))
		fmt.Printf("%s\n", strings.Repeat("=", 80))

		for currency, info := range cryptoBalances {
			fmt.Printf("\n%s:\n", currency)
			fmt.Printf("  Available: %-20s  Locked: %-20s\n", info.Available, info.Locked)
			fmt.Printf("  Pending:   %-20s  Frozen: %-20s\n", info.Pending, info.Frozen)
			fmt.Printf("  Total:     %s\n", info.Total)
		}
	}

	// Display fiat balances
	if len(fiatBalances) > 0 {
		fmt.Printf("\n%s\n", strings.Repeat("=", 80))
		fmt.Printf("Fiat Currency Balances (%d)\n", len(fiatBalances))
		fmt.Printf("%s\n", strings.Repeat("=", 80))

		for currency, info := range fiatBalances {
			fmt.Printf("\n%s:\n", currency)
			fmt.Printf("  Available: %-20s  Locked: %-20s\n", info.Available, info.Locked)
			fmt.Printf("  Pending:   %-20s  Frozen: %-20s\n", info.Pending, info.Frozen)
			fmt.Printf("  Total:     %s\n", info.Total)
		}
	}

	fmt.Println("\n" + strings.Repeat("=", 80))
}

// showAvailableBalancesOnly displays only assets with non-zero available balances.
// This is useful for quickly identifying which currencies you can use for payouts.
//
// Features:
//   - Filters out currencies with zero available balance
//   - Shows only immediately usable funds
//   - Helps in payout planning
//
// Use Cases:
//   - Check which currencies can be used for payouts
//   - Quick view of liquid assets
//   - Identify currencies needing deposits
//
// API Endpoints Used:
//   - GET /v1/balance
func showAvailableBalancesOnly(client *martianpay.Client) {
	fmt.Println("Showing Available Balances Only...")

	balance, err := client.GetBalance()
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Println("Available Balances (Ready for Use)")
	fmt.Println(strings.Repeat("=", 80))

	hasAvailableBalance := false

	for i, detail := range balance.BalanceDetails {
		// Skip if available balance is 0 or empty
		if detail.AvailableBalance == "0" || detail.AvailableBalance == "" {
			continue
		}

		hasAvailableBalance = true
		fmt.Printf("\n[%d] %s\n", i+1, detail.Currency)
		fmt.Printf("    Available: %s\n", detail.AvailableBalance)
	}

	if !hasAvailableBalance {
		fmt.Println("\n  No available balance found in any currency.")
	}

	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Printf("Total Available (Primary Currency): %s %s\n", balance.AvailableBalance, balance.Currency)
	fmt.Println(strings.Repeat("=", 80))
}

// showLockedAndPendingBalances displays balances that are not immediately available.
// This helps merchants track funds in transit or held for various reasons.
//
// Locked Balances:
//   - Funds reserved for pending payouts
//   - Held during transaction processing
//   - Will become available when transactions complete
//
// Pending Balances:
//   - Deposits awaiting blockchain confirmation
//   - Payments being reconciled
//   - Will become available after confirmation
//
// API Endpoints Used:
//   - GET /v1/balance
func showLockedAndPendingBalances(client *martianpay.Client) {
	fmt.Println("Showing Locked and Pending Balances...")

	balance, err := client.GetBalance()
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Println("Locked and Pending Balances")
	fmt.Println(strings.Repeat("=", 80))

	hasLockedOrPending := false

	for i, detail := range balance.BalanceDetails {
		hasLocked := detail.LockedBalance != "0" && detail.LockedBalance != ""
		hasPending := detail.PendingBalance != "0" && detail.PendingBalance != ""

		if !hasLocked && !hasPending {
			continue
		}

		hasLockedOrPending = true
		fmt.Printf("\n[%d] %s\n", i+1, detail.Currency)

		if hasLocked {
			fmt.Printf("    🔒 Locked:  %s (Reserved for pending payouts)\n", detail.LockedBalance)
		}

		if hasPending {
			fmt.Printf("    ⏳ Pending: %s (Awaiting reconciliation)\n", detail.PendingBalance)
		}
	}

	if !hasLockedOrPending {
		fmt.Println("\n  No locked or pending balances found.")
	}

	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Printf("Summary:\n")
	fmt.Printf("  Total Locked:  %s %s\n", balance.LockedBalance, balance.Currency)
	fmt.Printf("  Total Pending: %s %s\n", balance.PendingBalance, balance.Currency)
	fmt.Println(strings.Repeat("=", 80))
}

// compareBalanceTypes provides a comprehensive side-by-side comparison of all balance types.
// This tabular view makes it easy to see the distribution of funds across balance types.
//
// Features:
//   - Summary totals for all balance types
//   - Detailed breakdown by asset
//   - Tabular format for easy comparison
//
// Use Cases:
//   - Financial reporting
//   - Balance reconciliation
//   - Understanding fund allocation
//
// API Endpoints Used:
//   - GET /v1/balance
func compareBalanceTypes(client *martianpay.Client) {
	fmt.Println("Comparing Balance Types...")

	balance, err := client.GetBalance()
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Println("Balance Types Comparison")
	fmt.Println(strings.Repeat("=", 80))

	// Calculate totals
	fmt.Printf("\n%-20s: %s %s\n", "Available", balance.AvailableBalance, balance.Currency)
	fmt.Printf("%-20s: %s %s\n", "Pending", balance.PendingBalance, balance.Currency)
	fmt.Printf("%-20s: %s %s\n", "Locked", balance.LockedBalance, balance.Currency)
	fmt.Printf("%-20s: %s %s\n", "Frozen", balance.FrozenBalance, balance.Currency)
	fmt.Printf("%s\n", strings.Repeat("-", 80))
	fmt.Printf("%-20s: %s %s\n", "Total", balance.TotalBalance, balance.Currency)

	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Println("Details by Asset:")
	fmt.Println(strings.Repeat("=", 80))

	// Print header
	fmt.Printf("\n%-30s | %-12s | %-12s | %-12s | %-12s\n",
		"Currency/Asset", "Available", "Pending", "Locked", "Frozen")
	fmt.Println(strings.Repeat("-", 90))

	for _, detail := range balance.BalanceDetails {
		fmt.Printf("%-30s | %-12s | %-12s | %-12s | %-12s\n",
			detail.Currency,
			detail.AvailableBalance,
			detail.PendingBalance,
			detail.LockedBalance,
			detail.FrozenBalance)
	}

	fmt.Println(strings.Repeat("=", 90))
}

// BalanceInfo is a helper struct for organizing balance information by currency.
// It stores all balance types for a single currency/asset.
type BalanceInfo struct {
	Currency  string // Currency or asset ID
	Available string // Available balance (can be used immediately)
	Pending   string // Pending balance (awaiting confirmation)
	Locked    string // Locked balance (reserved for transactions)
	Frozen    string // Frozen balance (compliance/security hold)
	Total     string // Total balance (sum of all types)
}

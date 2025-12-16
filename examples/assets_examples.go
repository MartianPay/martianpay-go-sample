// Package main provides examples for the MartianPay Assets API.
// Assets represent the cryptocurrencies and fiat currencies that merchants can accept
// for payments or use for payouts. Each asset has specific properties like network,
// decimals, and availability for payments.
//
// Asset Types:
//   - Cryptocurrency: BTC, ETH, USDC, USDT, SOL, etc. (on various networks)
//   - Fiat Currency: USD (currently only supports USD via payment processors like Stripe)
//
// Key Properties:
//   - Payable: Whether the asset can be used for accepting payments
//   - Network: Blockchain network (Ethereum, Solana, Polygon, etc.)
//   - Mainnet/Testnet: Production or testing environment
//   - Contract Address: Smart contract address for tokens
package main

import (
	"fmt"
	"strings"

	martianpay "github.com/MartianPay/martianpay-go-sample/sdk"
)

// Assets Examples

// listAssets retrieves and displays all assets enabled for the merchant account.
// This shows assets that are currently configured and available for use.
//
// Displayed Information:
//   - Asset ID and display name
//   - Coin symbol (BTC, ETH, USDC, etc.)
//   - Decimals (precision for amounts)
//   - Payable status (can accept payments)
//   - Network and contract details (for crypto)
//   - Provider information (for fiat)
//
// Use Cases:
//   - Check which assets are enabled for your account
//   - Identify payable assets for payment intents
//   - View asset configuration details
//
// API Endpoints Used:
//   - GET /v1/assets (enabled assets only)
func listAssets(client *martianpay.Client) {
	fmt.Println("Listing Enabled Assets...")

	response, err := client.ListAssets()
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	if len(response.Assets) == 0 {
		fmt.Println("✓ No assets found")
		return
	}

	// Separate crypto and fiat assets
	var cryptoAssets, fiatAssets []string
	cryptoCount := 0
	fiatCount := 0

	for _, asset := range response.Assets {
		if asset.IsFiat {
			fiatCount++
			fiatAssets = append(fiatAssets, asset.Id)
		} else {
			cryptoCount++
			cryptoAssets = append(cryptoAssets, asset.Id)
		}
	}

	fmt.Printf("✓ Total Assets: %d (Crypto: %d, Fiat: %d)\n\n", len(response.Assets), cryptoCount, fiatCount)

	// Display crypto assets
	if cryptoCount > 0 {
		fmt.Printf("Cryptocurrency Assets (%d):\n", cryptoCount)
		fmt.Println(strings.Repeat("-", 80))

		for i, asset := range response.Assets {
			if asset.IsFiat {
				continue
			}

			fmt.Printf("\n[%d] %s\n", i+1, asset.DisplayName)
			fmt.Printf("    ID: %s\n", asset.Id)
			fmt.Printf("    Coin: %s\n", asset.Coin)
			fmt.Printf("    Decimals: %d\n", asset.Decimals)
			fmt.Printf("    Payable: %v\n", asset.Payable)

			if asset.CryptoAssetParams != nil {
				fmt.Printf("    Network: %s\n", asset.CryptoAssetParams.Network)
				fmt.Printf("    Mainnet: %v\n", asset.CryptoAssetParams.IsMainnet)
				if asset.CryptoAssetParams.ContractAddress != "" {
					fmt.Printf("    Contract: %s\n", asset.CryptoAssetParams.ContractAddress)
				}
				if asset.CryptoAssetParams.Token != "" {
					fmt.Printf("    Token: %s\n", asset.CryptoAssetParams.Token)
				}
				if asset.CryptoAssetParams.ChainId != 0 {
					fmt.Printf("    Chain ID: %d\n", asset.CryptoAssetParams.ChainId)
				}
			}
		}
		fmt.Println()
	}

	// Display fiat assets
	if fiatCount > 0 {
		fmt.Printf("\nFiat Currency Assets (%d):\n", fiatCount)
		fmt.Println(strings.Repeat("-", 80))

		for i, asset := range response.Assets {
			if !asset.IsFiat {
				continue
			}

			fmt.Printf("\n[%d] %s\n", i+1, asset.DisplayName)
			fmt.Printf("    ID: %s\n", asset.Id)
			fmt.Printf("    Coin: %s\n", asset.Coin)
			fmt.Printf("    Decimals: %d\n", asset.Decimals)
			fmt.Printf("    Payable: %v\n", asset.Payable)

			if asset.FiatAssetParams != nil {
				fmt.Printf("    Symbol: %s\n", asset.FiatAssetParams.Symbol)
				if asset.FiatAssetParams.Provider != "" {
					fmt.Printf("    Provider: %s\n", asset.FiatAssetParams.Provider)
				}
			}
		}
		fmt.Println()
	}
}

// getAllAssets retrieves all available assets including those not yet enabled.
// This is useful for discovering new assets that can be added to your account.
//
// Features:
//   - Shows both enabled and disabled assets
//   - Groups assets by network
//   - Allows filtering by crypto/fiat/payable
//   - Interactive display options
//
// Use Cases:
//   - Discover available cryptocurrencies
//   - Plan asset configuration
//   - Compare mainnet vs testnet options
//
// API Endpoints Used:
//   - GET /v1/assets/all (all available assets)
func getAllAssets(client *martianpay.Client) {
	fmt.Println("Getting All Available Assets...")

	assets, err := client.GetAllAssets()
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	if len(assets) == 0 {
		fmt.Println("✓ No assets found")
		return
	}

	// Count by type and network
	cryptoByNetwork := make(map[string]int)
	fiatCount := 0
	payableCount := 0

	for _, asset := range assets {
		if asset.IsFiat {
			fiatCount++
		} else if asset.CryptoAssetParams != nil {
			cryptoByNetwork[asset.CryptoAssetParams.Network]++
		}
		if asset.Payable {
			payableCount++
		}
	}

	fmt.Printf("✓ Total Assets: %d (Payable: %d)\n", len(assets), payableCount)
	fmt.Printf("  Fiat Assets: %d\n", fiatCount)
	fmt.Printf("  Crypto Assets: %d\n\n", len(assets)-fiatCount)

	if len(cryptoByNetwork) > 0 {
		fmt.Println("Assets by Network:")
		for network, count := range cryptoByNetwork {
			fmt.Printf("  - %s: %d\n", network, count)
		}
		fmt.Println()
	}

	// Show detailed list with option to filter
	fmt.Print("Show detailed list? (y/N): ")
	var showDetails string
	fmt.Scanln(&showDetails)

	if strings.ToLower(showDetails) != "y" {
		return
	}

	fmt.Print("\nFilter by: (1) All, (2) Crypto only, (3) Fiat only, (4) Payable only: ")
	var filterChoice string
	fmt.Scanln(&filterChoice)
	if filterChoice == "" {
		filterChoice = "1"
	}

	fmt.Println("\n" + strings.Repeat("=", 80))

	for i, asset := range assets {
		// Apply filter
		shouldShow := false
		switch filterChoice {
		case "1":
			shouldShow = true
		case "2":
			shouldShow = !asset.IsFiat
		case "3":
			shouldShow = asset.IsFiat
		case "4":
			shouldShow = asset.Payable
		}

		if !shouldShow {
			continue
		}

		assetType := "Crypto"
		if asset.IsFiat {
			assetType = "Fiat"
		}

		fmt.Printf("\n[%d] %s (%s)\n", i+1, asset.DisplayName, assetType)
		fmt.Printf("    ID: %s, Coin: %s\n", asset.Id, asset.Coin)
		fmt.Printf("    Decimals: %d, Payable: %v\n", asset.Decimals, asset.Payable)

		if asset.CryptoAssetParams != nil {
			fmt.Printf("    Network: %s (Mainnet: %v)\n",
				asset.CryptoAssetParams.Network,
				asset.CryptoAssetParams.IsMainnet)
		}

		if asset.FiatAssetParams != nil && asset.FiatAssetParams.Symbol != "" {
			fmt.Printf("    Symbol: %s\n", asset.FiatAssetParams.Symbol)
		}
	}

	fmt.Println("\n" + strings.Repeat("=", 80))
}

// listAssetFees retrieves network fee information for cryptocurrency withdrawals.
// Network fees are charged by blockchains for processing transactions.
//
// Displayed Information:
//   - Minimum payout amount (smallest withdrawal allowed)
//   - Network fee amount (cost to send transactions)
//   - Fees by blockchain network
//
// Use Cases:
//   - Calculate payout costs
//   - Determine minimum withdrawal amounts
//   - Compare fees across networks
//
// Note: Network fees vary by blockchain and can change based on network congestion.
//
// API Endpoints Used:
//   - GET /v1/assets/fees
func listAssetFees(client *martianpay.Client) {
	fmt.Println("Listing Asset Network Fees...")

	response, err := client.ListAssetFees()
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	if len(response.NetworkFees) == 0 {
		fmt.Println("✓ No network fees found")
		return
	}

	fmt.Printf("✓ Network Fees for %d Networks:\n\n", len(response.NetworkFees))
	fmt.Println(strings.Repeat("=", 80))

	// Sort networks for consistent display
	var networks []string
	for network := range response.NetworkFees {
		networks = append(networks, network)
	}

	// Display fees by network
	for i, network := range networks {
		fee := response.NetworkFees[network]

		fmt.Printf("\n[%d] %s\n", i+1, network)
		fmt.Printf("    Minimum Payout Amount: %s\n", fee.MinPayoutAmount.String())
		fmt.Printf("    Network Fee Amount: %s\n", fee.FeeAmount.String())
	}

	fmt.Println("\n" + strings.Repeat("=", 80))
}

// showCryptoAssetsByNetwork organizes and displays cryptocurrencies grouped by network.
// This helps visualize which tokens are available on each blockchain.
//
// Features:
//   - Groups assets by blockchain (Ethereum, Solana, Polygon, etc.)
//   - Shows payable status for each asset
//   - Counts assets per network
//
// Use Cases:
//   - Select optimal network for payments
//   - View multi-chain token availability
//   - Compare network options
//
// API Endpoints Used:
//   - GET /v1/assets/all
func showCryptoAssetsByNetwork(client *martianpay.Client) {
	fmt.Println("Showing Crypto Assets Grouped by Network...")

	assets, err := client.GetAllAssets()
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	// Group assets by network
	assetsByNetwork := make(map[string][]string)
	for _, asset := range assets {
		if !asset.IsFiat && asset.CryptoAssetParams != nil {
			network := asset.CryptoAssetParams.Network
			displayInfo := fmt.Sprintf("%s (%s)", asset.DisplayName, asset.Id)
			if asset.Payable {
				displayInfo += " [Payable]"
			}
			assetsByNetwork[network] = append(assetsByNetwork[network], displayInfo)
		}
	}

	if len(assetsByNetwork) == 0 {
		fmt.Println("✓ No crypto assets found")
		return
	}

	fmt.Printf("✓ Found assets on %d blockchain networks:\n\n", len(assetsByNetwork))

	for network, assets := range assetsByNetwork {
		fmt.Printf("━━━ %s (%d assets) ━━━\n", network, len(assets))
		for i, asset := range assets {
			fmt.Printf("  %d. %s\n", i+1, asset)
		}
		fmt.Println()
	}
}

// showPayableAssets displays only assets that can be used for accepting payments.
// These are the currencies/tokens that customers can use to pay you.
//
// Features:
//   - Filters to show only payable assets
//   - Separates crypto and fiat assets
//   - Shows network and provider information
//
// Use Cases:
//   - List payment options for customers
//   - Configure payment intent currencies
//   - Plan payment method offerings
//
// API Endpoints Used:
//   - GET /v1/assets (with payable filter)
func showPayableAssets(client *martianpay.Client) {
	fmt.Println("Showing Payable Assets...")

	response, err := client.ListAssets()
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	var payableCrypto, payableFiat int
	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Println("Payable Assets")
	fmt.Println(strings.Repeat("=", 80))

	for i, asset := range response.Assets {
		if !asset.Payable {
			continue
		}

		if asset.IsFiat {
			payableFiat++
		} else {
			payableCrypto++
		}

		assetType := "Crypto"
		networkInfo := ""
		if asset.IsFiat {
			assetType = "Fiat"
			if asset.FiatAssetParams != nil && asset.FiatAssetParams.Provider != "" {
				networkInfo = fmt.Sprintf(" [%s]", asset.FiatAssetParams.Provider)
			}
		} else if asset.CryptoAssetParams != nil {
			networkInfo = fmt.Sprintf(" [%s]", asset.CryptoAssetParams.Network)
		}

		fmt.Printf("\n[%d] %s (%s)%s\n", i+1, asset.DisplayName, assetType, networkInfo)
		fmt.Printf("    ID: %s\n", asset.Id)
		fmt.Printf("    Coin: %s, Decimals: %d\n", asset.Coin, asset.Decimals)
	}

	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Printf("Total Payable: %d (Crypto: %d, Fiat: %d)\n", payableCrypto+payableFiat, payableCrypto, payableFiat)
	fmt.Println(strings.Repeat("=", 80))
}

// compareMainnetVsTestnet compares production and testing cryptocurrency assets.
// Mainnet assets use real blockchain networks, while testnet assets use test networks.
//
// Mainnet Assets:
//   - Production blockchain networks
//   - Real cryptocurrency with value
//   - Used for actual transactions
//
// Testnet Assets:
//   - Test blockchain networks (e.g., Sepolia, Goerli)
//   - Test cryptocurrency (no real value)
//   - Used for development and testing
//
// Use Cases:
//   - Choose between production and testing environments
//   - Development and testing workflows
//   - Network availability comparison
//
// API Endpoints Used:
//   - GET /v1/assets/all
func compareMainnetVsTestnet(client *martianpay.Client) {
	fmt.Println("Comparing Mainnet vs Testnet Assets...")

	assets, err := client.GetAllAssets()
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	mainnetCount := 0
	testnetCount := 0
	mainnetByNetwork := make(map[string]int)
	testnetByNetwork := make(map[string]int)

	for _, asset := range assets {
		if asset.IsFiat || asset.CryptoAssetParams == nil {
			continue
		}

		if asset.CryptoAssetParams.IsMainnet {
			mainnetCount++
			mainnetByNetwork[asset.CryptoAssetParams.Network]++
		} else {
			testnetCount++
			testnetByNetwork[asset.CryptoAssetParams.Network]++
		}
	}

	fmt.Printf("\n✓ Mainnet Assets: %d\n", mainnetCount)
	fmt.Printf("✓ Testnet Assets: %d\n\n", testnetCount)

	if len(mainnetByNetwork) > 0 {
		fmt.Println("Mainnet Assets by Network:")
		for network, count := range mainnetByNetwork {
			fmt.Printf("  - %s: %d\n", network, count)
		}
		fmt.Println()
	}

	if len(testnetByNetwork) > 0 {
		fmt.Println("Testnet Assets by Network:")
		for network, count := range testnetByNetwork {
			fmt.Printf("  - %s: %d\n", network, count)
		}
		fmt.Println()
	}
}

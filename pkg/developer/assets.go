// assets.go contains types for managing supported cryptocurrencies and fiat currencies.
// It provides structures for asset definitions, exchange rates, and asset-related operations.
package developer

import "github.com/shopspring/decimal"

// ================================
// Core Types
// ================================

// Asset represents a cryptocurrency or fiat asset supported by the platform
type Asset struct {
	// Id is the unique identifier of the asset (e.g., "USDC_ETH", "USD")
	Id string `json:"id"`
	// DisplayName is the human-readable name of the asset (e.g., "USD Coin on Ethereum")
	DisplayName string `json:"display_name"`
	// Coin is the base currency code (e.g., "USDC", "BTC", "USD")
	Coin string `json:"coin"`
	// IsFiat indicates whether this is a fiat currency (true) or cryptocurrency (false)
	IsFiat bool `json:"is_fiat"`
	// Decimals is the number of decimal places for this asset
	Decimals int `json:"decimals"`
	// Payable indicates whether this asset can be used for payments
	Payable bool `json:"payable"`
	// CryptoAssetParams contains cryptocurrency-specific parameters (nil for fiat assets)
	*CryptoAssetParams
	// FiatAssetParams contains fiat currency-specific parameters (nil for crypto assets)
	*FiatAssetParams
}

// CryptoAssetParams contains blockchain-specific parameters for cryptocurrency assets
type CryptoAssetParams struct {
	// Network is the blockchain network name (e.g., "Ethereum", "Polygon", "Solana")
	Network string `json:"network"`
	// IsMainnet indicates if this is a mainnet (true) or testnet (false) asset
	IsMainnet bool `json:"is_mainnet"`
	// ContractAddress is the smart contract address for token assets (empty for native coins)
	ContractAddress string `json:"contract_address"`
	// AddressUrlTemplate is the URL template for viewing an address on a block explorer
	AddressUrlTemplate string `json:"address_url_template"`
	// TxUrlTemplate is the URL template for viewing a transaction on a block explorer
	TxUrlTemplate string `json:"tx_url_template"`
	// Token is the token symbol (e.g., "USDC", "USDT")
	Token string `json:"token"`
	// ChainId is the blockchain chain ID for EVM-compatible networks
	ChainId int64 `json:"chain_id"`
}

// FiatAssetParams contains parameters for fiat currency assets
type FiatAssetParams struct {
	// Symbol is the currency symbol (e.g., "$", "€", "£")
	Symbol string `json:"symbol"`
	// Provider is the payment provider for this fiat currency (e.g., "Stripe", "PayPal")
	Provider string `json:"provider"`
}

// NetworkFee represents fee information for a specific network
type NetworkFee struct {
	// MinPayoutAmount is the minimum payout amount allowed for this network
	MinPayoutAmount decimal.Decimal `json:"min_payout_amount" example:"10.00"`
	// FeeAmount is the fixed fee amount for transactions on this network
	FeeAmount decimal.Decimal `json:"fee_amount" example:"0.50"`
}

// ================================
// Request Types
// ================================

// AssetListRequest represents a request to list assets enabled for a specific merchant
type AssetListRequest struct {
	// MerchantId is the unique identifier of the merchant to filter assets
	MerchantId string `json:"merchant_id" form:"merchant_id" example:"merchant_123456"`
}

// NetworkFeeByAssetRequest represents a request to get network fees for a specific asset
type NetworkFeeByAssetRequest struct {
	// AssetID is the unique identifier of the asset to get fee information for
	AssetID string `form:"asset_id" binding:"required" example:"USDC"`
}

// ================================
// Response Types
// ================================

// AssetListResponse contains the list of assets enabled for the merchant
type AssetListResponse struct {
	// Assets is the list of cryptocurrency and fiat assets available to the merchant
	Assets []*Asset `json:"assets"`
}

// NetworkFeesResponse contains network fee information for all supported networks
type NetworkFeesResponse struct {
	// NetworkFees is a map of network names to their fee information
	NetworkFees map[string]*NetworkFee `json:"network_fees"`
}

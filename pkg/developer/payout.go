// payout.go contains types for managing cryptocurrency payouts and currency swaps.
// It provides structures for creating payouts, tracking payout status, and managing
// currency conversions through the MartianPay payout system.
package developer

import (
	"github.com/dchest/uniuri"
	"github.com/shopspring/decimal"
)

const (
	// PayoutIdPrefix is the prefix for payout IDs
	PayoutIdPrefix = "payout_"
	// PayoutIdLength is the length of the payout ID suffix
	PayoutIdLength = 24

	// PayoutObject is the object type identifier for payouts
	PayoutObject = "payout"
)

// Current status of the payout: `paid`, `pending`, `in_transit`, `canceled` or `failed`. A payout is `pending` until it's submitted to the bank, when it becomes `in_transit`. The status changes to `paid` if the transaction succeeds, or to `failed` or `canceled` (within 5 business days). Some payouts that fail might initially show as `paid`, then change to `failed`.
type PayoutStatus string

const (
	// PayoutStatusCanceled indicates the payout was canceled
	PayoutStatusCanceled PayoutStatus = "canceled"
	// PayoutStatusFailed indicates the payout failed
	PayoutStatusFailed PayoutStatus = "failed"
	// PayoutStatusInSwap indicates the payout is currently swapping currencies
	PayoutStatusInSwap PayoutStatus = "in_swap"
	// PayoutStatusInTransit indicates the payout is in transit to the bank
	PayoutStatusInTransit PayoutStatus = "in_transit"
	// PayoutStatusPaid indicates the payout was successfully paid
	PayoutStatusPaid PayoutStatus = "paid"
	// PayoutStatusPending indicates the payout is pending submission
	PayoutStatusPending PayoutStatus = "pending"
	// PayoutStatusApproved indicates the payout has been approved
	PayoutStatusApproved PayoutStatus = "approved"
	// PayoutStatusRejected indicates the payout was rejected
	PayoutStatusRejected PayoutStatus = "rejected"
)

// PayoutParams contains parameters for creating a payout
type PayoutParams struct {
	// SourceCoin is the source coin for the swap
	SourceCoin string `json:"source_coin"`
	// SourceAmount is the source amount for the swap
	SourceAmount string `json:"source_amount"`

	// QuoteIds is the list of quote IDs for the swap
	QuoteIds []string `json:"quote_ids"`

	// DestinationAssetId is the asset ID of the destination currency
	DestinationAssetId string `json:"receive_asset_id"`
	// DestinationAmount is the amount to be received
	DestinationAmount string `json:"receive_amount"`
	// DestinationAccountType is the type of the destination account (bank or wallet)
	DestinationAccountType string `json:"receive_account_type"`
	// DestinationAccountId is the ID of the destination account
	DestinationAccountId string `json:"receive_account_id"`
	// DestinationAddress is the optional address for receiving funds
	DestinationAddress *string `json:"receive_address"`
	// ToMerchantId is the destination merchant ID for internal transfers
	ToMerchantId string `json:"to_merchant_id"`

	// InternalNote contains internal notes for record keeping
	InternalNote string `json:"internal_note"`
	// StatementDescriptor is the description that appears on the statement
	StatementDescriptor string `json:"statement_descriptor"`
	// ExternalId is the external reference ID
	ExternalId string `json:"external_id"`
	// Metadata contains additional metadata as key-value pairs
	Metadata map[string]string `json:"metadata"`
}

// Payout represents a payout transaction
type Payout struct {
	// ID is the unique identifier for the payout
	ID string `json:"id"`
	// Object is the type of object, always "payout"
	Object string `json:"object"`

	// Livemode indicates whether this payout was created in live mode or test mode
	Livemode bool `json:"livemode"`

	// ArrivalDate is the date the payout is expected to arrive in the bank, factoring in delays for weekends or bank holidays
	ArrivalDate int64 `json:"arrival_date"`
	// Automatic indicates whether the payout is created by an automated payout schedule (true) or requested manually (false)
	Automatic bool `json:"automatic"`

	// Transactions is the list of transactions associated with this payout
	Transactions []*TransactionDetails `json:"transactions"`
	// Created is the Unix timestamp when the payout was created
	Created int64 `json:"created"`
	// Updated is the Unix timestamp when the payout was last updated
	Updated int64 `json:"updated"`
	// MerchantId is the ID of the merchant who owns this payout
	MerchantId string `json:"merchant_id"`

	// SourceAmount is the original amount to be paid out
	SourceAmount decimal.Decimal `json:"source_amount"`
	// SourceCoin is the currency code of the source amount
	SourceCoin string `json:"source_coin"`

	// ExchangeRate is the exchange rate used for currency conversion
	ExchangeRate decimal.Decimal `json:"exchange_rate"`

	// ReceiveCoin is the currency code of the receive amount
	ReceiveCoin string `json:"receive_coin"`
	// ReceiveAssetId is the asset ID of the destination currency
	ReceiveAssetId string `json:"receive_asset_id"`
	// ReceiveAccountType is the type of receiving account (bank or wallet)
	ReceiveAccountType string `json:"receive_account_type"`
	// ReceiveBankAccount contains the bank account details if receiving to bank
	ReceiveBankAccount *MerchantAccount `json:"receive_bank_account"`
	// ReceiveWalletAddress contains the wallet address if receiving to crypto wallet
	ReceiveWalletAddress *MerchantAddress `json:"receive_wallet_address"`
	// ReceiveAmount is the amount to be received after conversion
	ReceiveAmount decimal.Decimal `json:"receive_amount"`
	// ReceiveAmountMin is the minimum amount guaranteed to be received
	ReceiveAmountMin decimal.Decimal `json:"receive_amount_min"`

	// PaymentMaxAmount is the maximum amount that can be paid
	PaymentMaxAmount decimal.Decimal `json:"payment_max_amount"`
	// PaymentNetworkFee is the network fee for the transaction
	PaymentNetworkFee decimal.Decimal `json:"payment_network_fee"`
	// PaymentServiceFee is the service fee charged for the payout
	PaymentServiceFee decimal.Decimal `json:"payment_service_fee"`
	// PaymentTotalFee is the total of all fees
	PaymentTotalFee decimal.Decimal `json:"payment_total_fee"`
	// PaymentNetAmount is the net amount after deducting fees
	PaymentNetAmount decimal.Decimal `json:"payment_net_amount"`

	// Status is the current status (paid, pending, in_transit, canceled, failed, in_swap, approved, or rejected)
	Status PayoutStatus `json:"status"`

	// ApprovalStatus is the approval status (in_progress, approved, or rejected)
	ApprovalStatus *string `json:"approval_status"`

	// InternalNote contains internal notes for record keeping
	InternalNote string `json:"internal_note"`
	// StatementDescriptor is the description that appears on the statement
	StatementDescriptor string `json:"statement_descriptor"`

	// ExternalId is the external reference ID provided by the client
	ExternalId string `json:"external_id"`

	// FailureMessage is the detailed message explaining reason for failure if applicable
	FailureMessage string `json:"failure_message"`

	// AmlStatus is the AML check status
	AmlStatus string `json:"aml_status"`
	// AmlInfo contains AML information (semicolon-separated rule names)
	AmlInfo *string `json:"aml_info"`

	// Metadata contains additional metadata as key-value pairs
	Metadata map[string]string `json:"metadata"`
}

// GeneratePayoutId generates a new unique payout identifier with the 'payout_' prefix
func GeneratePayoutId() string {
	uniqueString := uniuri.NewLen(PayoutIdLength)
	return PayoutIdPrefix + uniqueString
}

// Withdraw represents a withdrawal transaction
type Withdraw struct {
	// ID is the unique identifier for the withdrawal
	ID string `json:"id"`
	// MerchantID is the ID of the merchant initiating the withdrawal
	MerchantID string `json:"merchant_id"`
	// PayoutID is the ID of the associated payout
	PayoutID string `json:"payout_id"`
	// AssetID is the ID of the asset being withdrawn
	AssetID string `json:"asset_id"`
	// Decimals is the decimal precision of the asset
	Decimals int32 `json:"decimals"`
	// Amount is the total amount to be withdrawn
	Amount decimal.Decimal `json:"amount"`
	// TxFee is the transaction fee for the withdrawal
	TxFee decimal.Decimal `json:"tx_fee"`
	// NetworkFee is the network fee for processing the withdrawal
	NetworkFee decimal.Decimal `json:"network_fee"`
	// Status is the current status of the withdrawal
	Status string `json:"status"`
	// SubStatus is the sub status (e.g. "pending_approval")
	SubStatus string `json:"sub_status"`
	// Address is the destination address for the withdrawal
	Address string `json:"address"`
	// Type is the withdrawal type ("standard" or "unfreeze")
	Type string `json:"type,omitempty"`
	// OriginalFrozenTxID is the original AML-rejected transaction ID (for unfreeze withdrawals)
	OriginalFrozenTxID string `json:"original_frozen_tx_id,omitempty"`
	// PaymentIntentID is the payment intent ID (for unfreeze withdrawals)
	PaymentIntentID string `json:"payment_intent_id,omitempty"`
	// CreatedAt is the Unix timestamp when the withdrawal was created
	CreatedAt int64 `json:"created_at"`
	// UpdatedAt is the Unix timestamp when the withdrawal was last updated
	UpdatedAt int64 `json:"updated_at"`
	// Version is the version number for optimistic locking
	Version int64 `json:"version"`
}

// PayoutSwapItem represents a swap item for a payout
type PayoutSwapItem struct {
	// ID is the unique identifier for the swap item
	ID string `json:"id"`
	// QuoteId is the ID of the associated quote
	QuoteId string `json:"quote_id"`
	// CreatedAt is the Unix timestamp when the swap item was created
	CreatedAt int64 `json:"created_at"`
	// UpdatedAt is the Unix timestamp when the swap item was last updated
	UpdatedAt int64 `json:"updated_at"`
	// MerchantID is the ID of the merchant
	MerchantID string `json:"merchant_id"`
	// PayoutID is the ID of the associated payout
	PayoutID string `json:"payout_id"`
	// Payout contains the payout details (if requested)
	Payout *Payout `json:"payout,omitempty"`
	// FromAssetId is the source asset ID for the swap
	FromAssetId string `json:"from_asset_id"`
	// ExpectedFromAmount is the expected amount to swap from
	ExpectedFromAmount string `json:"expected_from_amount"`
	// EstimatedFromAmount is the estimated amount to be swapped from
	EstimatedFromAmount string `json:"estimated_from_amount"`
	// EstimatedFromAmountUsd is the USD value of the estimated from amount
	EstimatedFromAmountUsd string `json:"estimated_from_amount_usd"`
	// ActualFromAmount is the actual amount swapped from
	ActualFromAmount string `json:"actual_from_amount"`
	// NetworkFee is the network fee for the swap transaction
	NetworkFee string `json:"network_fee"`
	// ToAssetId is the destination asset ID for the swap
	ToAssetId string `json:"to_asset_id"`
	// ToAddress is the destination address for Binance swap (funds sent directly here)
	ToAddress string `json:"to_address,omitempty"`
	// TxFee is the service fee for the swap transaction
	TxFee string `json:"tx_fee,omitempty"`
	// ExpectedToAmount is the expected amount to receive
	ExpectedToAmount string `json:"expected_to_amount"`
	// EstimatedToAmount is the estimated amount to be received
	EstimatedToAmount string `json:"estimated_to_amount"`
	// EstimatedToAmountMin is the minimum guaranteed receive amount
	EstimatedToAmountMin string `json:"estimated_to_amount_min"`
	// EstimatedToAmountUsd is the USD value of the estimated to amount
	EstimatedToAmountUsd string `json:"estimated_to_amount_usd"`
	// ActualToAmount is the actual amount received
	ActualToAmount string `json:"actual_to_amount"`
	// Status is the current status of the swap
	Status string `json:"status"`
}

// ================================
// Request Types
// ================================

// PayoutPreviewRequest previews payout before creation
type PayoutPreviewRequest struct {
	PayoutParams
}

// PayoutCreateRequest creates a new payout
type PayoutCreateRequest struct {
	PayoutParams
}

// PayoutGetRequest gets a specific payout
type PayoutGetRequest struct {
	// ID is the payout ID from URL path
	ID string `json:"-"`
}

// PayoutListRequest lists payouts with filters
type PayoutListRequest struct {
	// Page is the page number, starting from 0
	Page int32 `json:"page" binding:"min=0" form:"page"`
	// PageSize is the number of items per page (max 50)
	PageSize int32 `json:"page_size" binding:"required,min=1,max=50" form:"page_size"`
	// Status filters payouts by status
	Status *string `json:"status" form:"status"`
	// MerchantID filters payouts by merchant ID
	MerchantID *string `json:"merchant_id" form:"merchant_id"`
	// StartTime filters payouts created after this Unix timestamp
	StartTime *int64 `json:"start_time" form:"start_time"`
	// EndTime filters payouts created before this Unix timestamp
	EndTime *int64 `json:"end_time" form:"end_time"`
	// ExternalID filters payouts by external ID
	ExternalID *string `json:"external_id" form:"external_id"`
}

// PayoutSwapPreviewRequest previews swap for payout
type PayoutSwapPreviewRequest struct {
	// FromAssetId is the source asset ID for the swap
	FromAssetId string `json:"from_asset_id"`
	// FromAmount is the amount to swap from
	FromAmount string `json:"from_amount"`
	// QuoteId is the ID of the quote to use
	QuoteId string `json:"quote_id"`
	// ToAssetId is the destination asset ID for the swap
	ToAssetId string `json:"to_asset_id"`
	// ToAmount is the amount to receive
	ToAmount string `json:"to_amount"`
}

// ================================
// Response Types
// ================================

// PayoutPreviewResp represents payout preview response
type PayoutPreviewResp struct {
	Payout
	// Withdraw contains the withdrawal details
	Withdraw *Withdraw `json:"withdraw,omitempty"`
	// SwapItems contains the swap items for the payout
	SwapItems []*PayoutSwapItem `json:"swap_items,omitempty"`
}

// PayoutCreateResp represents payout creation response
type PayoutCreateResp struct {
	Payout
	// Withdraw contains the withdrawal details
	Withdraw *Withdraw `json:"withdraw,omitempty"`
	// SwapItems contains the swap items for the payout
	SwapItems []*PayoutSwapItem `json:"swap_items,omitempty"`
}

// PayoutGetResp represents payout get response
type PayoutGetResp struct {
	Payout
	// Withdraw contains the withdrawal details
	Withdraw *Withdraw `json:"withdraw,omitempty"`
	// SwapItems contains the swap items for the payout
	SwapItems []*PayoutSwapItem `json:"swap_items,omitempty"`
}

// PayoutListResp represents paginated payout list
type PayoutListResp struct {
	// Payouts is the list of payouts
	Payouts []Payout `json:"payouts"`
	// Total is the total number of payouts matching the filters
	Total int32 `json:"total"`
	// Page is the current page number
	Page int32 `json:"page"`
	// PageSize is the number of items per page
	PageSize int32 `json:"page_size"`
}

// AdminPayoutListResp represents admin payout list response
type AdminPayoutListResp struct {
	// Payouts is the list of payouts
	Payouts []Payout `json:"payouts"`
	// Total is the total number of payouts matching the filters
	Total int32 `json:"total"`
	// Page is the current page number
	Page int32 `json:"page"`
	// PageSize is the number of items per page
	PageSize int32 `json:"page_size"`
}

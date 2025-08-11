package developer

import (
	"github.com/dchest/uniuri"
	"github.com/shopspring/decimal"
)

const (
	PayoutIdPrefix = "payout_"
	PayoutIdLength = 24

	PayoutObject = "payout"
)

// Current status of the payout: `paid`, `pending`, `in_transit`, `canceled` or `failed`. A payout is `pending` until it's submitted to the bank, when it becomes `in_transit`. The status changes to `paid` if the transaction succeeds, or to `failed` or `canceled` (within 5 business days). Some payouts that fail might initially show as `paid`, then change to `failed`.
type PayoutStatus string

// List of values that PayoutStatus can take
const (
	PayoutStatusCanceled  PayoutStatus = "canceled"
	PayoutStatusFailed    PayoutStatus = "failed"
	PayoutStatusInSwap    PayoutStatus = "in_swap"
	PayoutStatusInTransit PayoutStatus = "in_transit"
	PayoutStatusPaid      PayoutStatus = "paid"
	PayoutStatusPending   PayoutStatus = "pending"
	PayoutStatusApproved  PayoutStatus = "approved"
	PayoutStatusRejected  PayoutStatus = "rejected"
)

type PayoutParams struct {
	SourceCoin string `json:"source_coin"` // Source coin for the swap
	SourceAmount string `json:"source_amount"` // Source amount for the swap

	// Quote IDs for the swap
	QuoteIds []string `json:"quote_ids"` // List of quote IDs for the swap

	// Destination/Receive asset information
	DestinationAssetId     string  `json:"receive_asset_id"`     // Asset ID of the destination currency
	DestinationAmount      string  `json:"receive_amount"`       // Amount to be received
	DestinationAccountType string  `json:"receive_account_type"` // Type of the destination account (bank or wallet)
	DestinationAccountId   string  `json:"receive_account_id"`   // ID of the destination account
	DestinationAddress     *string `json:"receive_address"`      // Optional address for receiving funds
	ToMerchantId           string  `json:"to_merchant_id"`       // For internal transfers - destination merchant ID

	// Additional information
	InternalNote        string            `json:"internal_note"`        // Internal notes for record keeping
	StatementDescriptor string            `json:"statement_descriptor"` // Description that appears on the statement
	ExternalId          string            `json:"external_id"`          // External reference ID
	Metadata            map[string]string `json:"metadata"`             // Additional metadata as key-value pairs
}

type Payout struct {
	ID     string `json:"id"`     // Unique identifier for the payout
	Object string `json:"object"` // Type of object, always "payout"

	Livemode bool `json:"livemode"` // Whether this payout was created in live mode or test mode

	// Date that you can expect the payout to arrive in the bank. This factors in delays to account for weekends or bank holidays.
	ArrivalDate int64 `json:"arrival_date"`
	// Returns `true` if the payout is created by an automated payout schedule and `false` if it's requested manually.
	Automatic bool `json:"automatic"`

	Transactions []*TransactionDetails `json:"transactions"` // List of transactions associated with this payout
	Created      int64                 `json:"created"`      // Timestamp when the payout was created
	Updated      int64                 `json:"updated"`      // Timestamp when the payout was last updated
	MerchantId   string                `json:"merchant_id"`  // ID of the merchant who owns this payout

	SourceAmount decimal.Decimal `json:"source_amount"` // Original amount to be paid out
	SourceCoin   string          `json:"source_coin"`   // Currency code of the source amount

	ExchangeRate decimal.Decimal `json:"exchange_rate"` // Exchange rate used for currency conversion

	ReceiveCoin          string           `json:"receive_coin"`           // Currency code of the receive amount
	ReceiveAssetId       string           `json:"receive_asset_id"`       // Asset ID of the destination currency
	ReceiveAccountType   string           `json:"receive_account_type"`   // Type of receiving account (bank or wallet)
	ReceiveBankAccount   *MerchantAccount `json:"receive_bank_account"`   // Bank account details if receiving to bank
	ReceiveWalletAddress *MerchantAddress `json:"receive_wallet_address"` // Wallet address if receiving to crypto wallet
	ReceiveAmount        decimal.Decimal  `json:"receive_amount"`         // Amount to be received after conversion
	ReceiveAmountMin     decimal.Decimal  `json:"receive_amount_min"`     // Minimum amount guaranteed to be received
	PaymentMaxAmount  decimal.Decimal `json:"payment_max_amount"`  // Maximum amount that can be paid

	PaymentNetworkFee decimal.Decimal `json:"payment_network_fee"` // Network fee for the transaction
	PaymentServiceFee decimal.Decimal `json:"payment_service_fee"` // Service fee charged for the payout
	PaymentTotalFee   decimal.Decimal `json:"payment_total_fee"`   // Total of all fees
	PaymentNetAmount  decimal.Decimal `json:"payment_net_amount"`  // Net amount after deducting fees

	// Current status of the payout: `paid`, `pending`, `in_transit`, `canceled` or `failed`. A payout is `pending` until it's submitted to the bank, when it becomes `in_transit`. The status changes to `paid` if the transaction succeeds, or to `failed` or `canceled` (within 5 business days). Some payouts that fail might initially show as `paid`, then change to `failed`.
	Status PayoutStatus `json:"status"`

	// Approval status of the payout: `in_progress`, `approved` or `rejected`.
	ApprovalStatus *string `json:"approval_status"`

	InternalNote        string `json:"internal_note"`        // Internal notes for record keeping
	StatementDescriptor string `json:"statement_descriptor"` // Description that appears on the statement

	ExternalId string `json:"external_id"` // External reference ID provided by the client

	FailureMessage string `json:"failure_message"` // Detailed message explaining reason for failure if applicable

	Metadata map[string]string `json:"metadata"` // Additional metadata as key-value pairs
}

func GeneratePayoutId() string {
	uniqueString := uniuri.NewLen(PayoutIdLength)
	return PayoutIdPrefix + uniqueString
}

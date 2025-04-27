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
	PayoutStatusInTransit PayoutStatus = "in_transit"
	PayoutStatusPaid      PayoutStatus = "paid"
	PayoutStatusPending   PayoutStatus = "pending"
	PayoutStatusApproved  PayoutStatus = "approved"
	PayoutStatusRejected  PayoutStatus = "rejected"
)

type PayoutParams struct {
	SourceCoin             string            `json:"source_coin"`
	DestinationAmount      string            `json:"receive_amount"`
	DestinationCoin        string            `json:"receive_coin"`
	DestinationAccountType string            `json:"receive_account_type"`
	DestinationAccountId   string            `json:"receive_account_id"`
	DestinationAssetId     string            `json:"receive_asset_id"`
	DestinationAddress     *string           `json:"receive_address"`
	InternalNote           string            `json:"internal_note"`
	StatementDescriptor    string            `json:"statement_descriptor"`
	ExternalId             string            `json:"external_id"`
	Metadata               map[string]string `json:"metadata"`
}

type Payout struct {
	ID     string `json:"id"`
	Object string `json:"object"`

	Livemode bool `json:"livemode"`

	// Date that you can expect the payout to arrive in the bank. This factors in delays to account for weekends or bank holidays.
	ArrivalDate int64 `json:"arrival_date"`
	// Returns `true` if the payout is created by an automated payout schedule and `false` if it's requested manually.
	Automatic bool `json:"automatic"`

	Transactions []*TransactionDetails `json:"transactions"` // 区块链交易列表
	Created      int64                 `json:"created"`
	Updated      int64                 `json:"updated"`
	MerchantId   string                `json:"merchant_id"`

	SourceAmount decimal.Decimal `json:"source_amount"`
	SourceCoin   string          `json:"source_coin"`

	ExchangeRate decimal.Decimal `json:"exchange_rate"`

	PaymentCoin          string           `json:"receive_coin"`
	PaymentAssetId       string           `json:"receive_asset_id"`
	PaymentAccountType   string           `json:"receive_account_type"`
	PaymentBankAccount   *MerchantAccount `json:"receive_bank_account"`
	PaymentWalletAddress *MerchantAddress `json:"receive_wallet_address"`
	PaymentAmount        decimal.Decimal  `json:"receive_amount"`
	ReceiveMaxAmount     decimal.Decimal  `json:"receive_max_amount"`

	PaymentNetworkFee decimal.Decimal `json:"payment_network_fee"`
	PaymentServiceFee decimal.Decimal `json:"payment_service_fee"`
	PaymentTotalFee   decimal.Decimal `json:"payment_total_fee"`
	PaymentNetAmount  decimal.Decimal `json:"payment_net_amount"`

	// Current status of the payout: `paid`, `pending`, `in_transit`, `canceled` or `failed`. A payout is `pending` until it's submitted to the bank, when it becomes `in_transit`. The status changes to `paid` if the transaction succeeds, or to `failed` or `canceled` (within 5 business days). Some payouts that fail might initially show as `paid`, then change to `failed`.
	Status PayoutStatus `json:"status"`

	InternalNote        string `json:"internal_note"`
	StatementDescriptor string `json:"statement_descriptor"`

	ExternalId string `json:"external_id"`

	FailureMessage string `json:"failure_message"`

	Metadata map[string]string `json:"metadata"`
}

func GeneratePayoutId() string {
	uniqueString := uniuri.NewLen(PayoutIdLength)
	return PayoutIdPrefix + uniqueString
}

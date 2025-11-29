package developer

import (
	"github.com/dchest/uniuri"
)

const (
	PayrollIdPrefix = "payroll_"
	PayrollIdLength = 24

	PayrollItemIdPrefix = "payroll_item_"
	PayrollItemIdLength = 24

	PayrollSwapItemIdPrefix = "payroll_swap_item_"
	PayrollSwapItemIdLength = 24

	PayrollBinanceFromItemPrefix = "payroll_bf_"
	PayrollBinanceFromItemLength = 24

	PayrollAddressVerifyIDPrefix = "payroll_av_"
	PayrollAddressVerifyIDLength = 24
)

type Payroll struct {
	ID              string `json:"id"`                // Unique identifier for the payroll
	CreatedAt       int64  `json:"created_at"`        // Timestamp when the payroll was created
	UpdatedAt       int64  `json:"updated_at"`        // Timestamp when the payroll was last updated
	CanceledAt      int64  `json:"canceled_at"`       // Timestamp when the payroll was canceled, if applicable
	MerchantID      string `json:"merchant_id"`       // ID of the merchant who created the payroll
	ExternalID      string `json:"external_id"`       // External ID for idempotency
	ApprovalStatus  string `json:"approval_status"`   // Current approval status of the payroll
	Status          string `json:"status"`            // Current processing status of the payroll
	TotalItemNum    int64  `json:"total_item_num"`    // Total number of items in the payroll
	TotalAmount     string `json:"total_amount"`      // Total amount to be paid in the payroll
	TotalServiceFee string `json:"total_service_fee"` // Total service fee for the payroll
	Currency        string `json:"currency"`          // Currency used for the payroll
}

type PayrollValidation struct {
	NameValid     bool `json:"name_valid"`     // Whether the recipient name is valid
	EmailValid    bool `json:"email_valid"`    // Whether the recipient email is valid
	PhoneValid    bool `json:"phone_valid"`    // Whether the recipient phone number is valid
	AmountValid   bool `json:"amount_valid"`   // Whether the payment amount is valid
	CoinValid     bool `json:"coin_valid"`     // Whether the cryptocurrency is valid
	NetworkValid  bool `json:"network_valid"`  // Whether the blockchain network is valid
	AddressValid  bool `json:"address_valid"`  // Whether the wallet address is valid
	BalanceEnough bool `json:"balance_enough"` // Whether there's enough balance to complete the payment
}

func (p *PayrollValidation) IsWalletInfoValid() bool {
	return p.AmountValid && p.CoinValid && p.NetworkValid && p.AddressValid
}

type PayrollDifferenceField struct {
	Previous string `json:"previous"` // Previous value before changes
	Current  string `json:"current"`  // Current value after changes
	Status   string `json:"status"`   // "added", "modified", or "unchanged"
}

const (
	PayrollDifferenceStatusAdded     = "added"
	PayrollDifferenceStatusModified  = "modified"
	PayrollDifferenceStatusUnchanged = "unchanged"
)

type PayrollDifference struct {
	Name    PayrollDifferenceField `json:"name"`
	Email   PayrollDifferenceField `json:"email"`
	Phone   PayrollDifferenceField `json:"phone"`
	Amount  PayrollDifferenceField `json:"amount"`
	Coin    PayrollDifferenceField `json:"coin"`
	Network PayrollDifferenceField `json:"network"`
	Address PayrollDifferenceField `json:"address"`
}

type PayrollItemParams struct {
	ExternalID    string `json:"external_id"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	Phone         string `json:"phone"`
	Amount        string `json:"amount"`
	Coin          string `json:"coin"`
	Network       string `json:"network"`
	Address       string `json:"address"`
	PaymentMethod string `json:"payment_method"` // "normal" or "binance"
}
type PayrollItems struct {
	ID                string                `json:"id"`                    // Unique identifier for the payroll item
	CreatedAt         int64                 `json:"created_at"`            // Timestamp when the item was created
	UpdatedAt         int64                 `json:"updated_at"`            // Timestamp when the item was last updated
	PayrollID         string                `json:"payroll_id"`            // ID of the parent payroll
	Payroll           *Payroll              `json:"payroll"`               // Reference to the parent payroll
	ExternalID        string                `json:"external_id"`           // External ID for idempotency
	Name              string                `json:"name"`                  // Recipient name
	Email             string                `json:"email"`                 // Recipient email address
	Phone             string                `json:"phone"`                 // Recipient phone number
	Amount            string                `json:"amount"`                // Payment amount
	ServiceFee        string                `json:"service_fee"`           // Service fee charged for the payment
	ExchangeRateToUSD string                `json:"exchange_rate_to_usd"`  // Exchange rate to USD at time of payment
	Coin              string                `json:"coin"`                  // Cryptocurrency used for payment
	Network           string                `json:"network"`               // Blockchain network used for the transaction
	AssetID           string                `json:"asset_id"`              // Asset identifier in the blockchain
	Address           string                `json:"address"`               // Recipient wallet address
	AddressVerified   bool                  `json:"address_verified"`      // Whether the address has been verified
	Status            string                `json:"status"`                // Current status of the payroll item
	Transactions      []*TransactionDetails `json:"transactions"`          // Details of related transactions
	Validation        *PayrollValidation    `json:"validation,omitempty"`  // Validation status for the item fields
	Differences       *PayrollDifference    `json:"differences,omitempty"` // Changes made to item since creation
	HasMonthlyPayment bool                  `json:"has_monthly_payment"`   // Whether this is a recurring monthly payment
	PaymentMethod     string                `json:"payment_method"`        // Payment method (normal or binance)
	IsBinance         bool                  `json:"is_binance"`            // Whether this is a binance payroll item
	BinanceTag        string                `json:"binance_tag"`           // Binance tag for the item
	BinanceTaskID     string                `json:"binance_task_id"`       // Binance task ID for the item
	AmlInfo           string                `json:"aml_info"`              // AML info for the item
	TxId              string                `json:"tx_id"`                 // Transaction ID (Binance order ID or blockchain tx hash)
}

type PayrollSwapItems struct {
	ID                     string   `json:"id"`                        // Unique identifier for the swap item
	LifiQuoteId            string   `json:"lifi_quote_id"`             // Quote ID from the LiFi service
	CreatedAt              int64    `json:"created_at"`                // Timestamp when the swap was created
	UpdatedAt              int64    `json:"updated_at"`                // Timestamp when the swap was last updated
	SentAt                 *int64   `json:"sent_at"`                   // Timestamp when the swap was sent to execution
	MerchantID             string   `json:"merchant_id"`               // ID of the merchant who initiated the swap
	PayrollID              string   `json:"payroll_id"`                // ID of the related payroll
	Payroll                *Payroll `json:"payroll,omitempty"`         // Reference to the related payroll
	FromAssetId            string   `json:"from_asset_id"`             // Source asset identifier
	ExpectedFromAmount     string   `json:"expected_from_amount"`      // Expected amount to be swapped from source
	EstimatedFromAmount    string   `json:"estimated_from_amount"`     // Estimated amount to be swapped from source
	EstimatedFromAmountUsd string   `json:"estimated_from_amount_usd"` // USD value of estimated source amount
	ActualFromAmount       string   `json:"actual_from_amount"`        // Actual amount swapped from source
	NetworkFee             string   `json:"network_fee"`               // Network fee paid for the swap
	ToAssetId              string   `json:"to_asset_id"`               // Destination asset identifier
	ExpectedToAmount       string   `json:"expected_to_amount"`        // Expected amount to be received
	EstimatedToAmount      string   `json:"estimated_to_amount"`       // Estimated amount to be received
	EstimatedToAmountMin   string   `json:"estimated_to_amount_min"`   // Minimum estimated amount to be received
	EstimatedToAmountUsd   string   `json:"estimated_to_amount_usd"`   // USD value of estimated destination amount
	ActualToAmount         string   `json:"actual_to_amount"`          // Actual amount received after swap
	Status                 string   `json:"status"`                    // Current status of the swap
}

type BinanceFromItems struct {
	ID            string `json:"id"`              // Unique identifier for the Binance from item
	PayrollID     string `json:"payroll_id"`      // ID of the payroll this item belongs to
	MerchantID    string `json:"merchant_id"`     // ID of the merchant
	AssetID       string `json:"asset_id"`        // Asset ID to transfer from
	Coin          string `json:"coin"`            // Coin type
	Network       string `json:"network"`         // Network type
	Amount        string `json:"amount"`          // Amount to transfer to Binance
	BinanceTaskID string `json:"binance_task_id"` // Associated Binance task ID
	Status        string `json:"status"`          // Current status of the transfer
	CreatedAt     int64  `json:"created_at"`      // Timestamp when the item was created
	UpdatedAt     int64  `json:"updated_at"`      // Timestamp when the item was last updated
}

func GeneratePayrollId() string {
	uniqueString := uniuri.NewLen(PayrollIdLength)
	return PayrollIdPrefix + uniqueString
}

func GeneratePayrollItemId() string {
	uniqueString := uniuri.NewLen(PayrollItemIdLength)
	return PayrollItemIdPrefix + uniqueString
}

func GeneratePayrollSwapItemId() string {
	uniqueString := uniuri.NewLen(PayrollSwapItemIdLength)
	return PayrollSwapItemIdPrefix + uniqueString
}

func GeneratePayrollBinanceFromItemId() string {
	uniqueString := uniuri.NewLen(PayrollBinanceFromItemLength)
	return PayrollBinanceFromItemPrefix + uniqueString
}

func GeneratePayrollAddressVerifyID() string {
	uniqueString := uniuri.NewLen(PayrollAddressVerifyIDLength)
	return PayrollAddressVerifyIDPrefix + uniqueString
}

func GeneratePayrollAddressVerificationCode() string {
	uniqueString := uniuri.NewLen(PayrollAddressVerifyIDLength)
	return uniqueString
}

const (
	PayrollPaymentMethodTypeBinance = "binance"
	PayrollPaymentMethodTypeNormal  = "normal"

	BinanceTaskIdPrefix = "binance_task_"
	BinanceTaskIdLength = 24

	BinanceDepositIdPrefix = "binance_deposit_"
	BinanceDepositIdLength = 24

	BinanceWithdrawIdPrefix = "binance_withdraw_"
	BinanceWithdrawIdLength = 24

	BinanceTradeIdPrefix = "binance_trade_"
	BinanceTradeIdLength = 24
)

// GenerateBinanceTaskId generates a unique Binance task ID
func GenerateBinanceTaskId() string {
	uniqueString := uniuri.NewLen(BinanceTaskIdLength)
	return BinanceTaskIdPrefix + uniqueString
}

// GenerateBinanceDepositId generates a unique Binance deposit ID
func GenerateBinanceDepositId() string {
	uniqueString := uniuri.NewLen(BinanceDepositIdLength)
	return BinanceDepositIdPrefix + uniqueString
}

// GenerateBinanceWithdrawId generates a unique Binance withdraw ID
func GenerateBinanceWithdrawId() string {
	uniqueString := uniuri.NewLen(BinanceWithdrawIdLength)
	return BinanceWithdrawIdPrefix + uniqueString
}

// GenerateBinanceTradeId generates a unique Binance trade ID
func GenerateBinanceTradeId() string {
	uniqueString := uniuri.NewLen(BinanceTradeIdLength)
	return BinanceTradeIdPrefix + uniqueString
}

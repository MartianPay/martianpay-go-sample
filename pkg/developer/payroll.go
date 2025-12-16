// payroll.go contains types for managing batch payroll payments.
// It provides structures for creating payroll batches, managing payroll items,
// and tracking address verification for payroll recipients.
package developer

import (
	"github.com/dchest/uniuri"
)

// ================================
// Request Types
// ================================

// PayrollConfirmRequest confirms a draft payroll
type PayrollConfirmRequest struct {
	// Id is the payroll ID from URL path
	Id string `json:"-"`
}

// PayrollGetRequest gets a specific payroll
type PayrollGetRequest struct {
	// Id is the payroll ID from URL path
	Id string `json:"-"`
}

// PayrollListRequest lists payrolls with filters
type PayrollListRequest struct {
	// Page is the page number, starting from 0
	Page int32 `json:"page" binding:"min=0" form:"page"`
	// PageSize is the number of items per page (max 50)
	PageSize int32 `json:"page_size" binding:"required,min=1,max=50" form:"page_size"`
	// StartDate filters payrolls created after this date
	StartDate *string `json:"start_date" form:"start_date"`
	// EndDate filters payrolls created before this date
	EndDate *string `json:"end_date" form:"end_date"`
	// ExternalID filters payrolls by external ID
	ExternalID *string `json:"external_id" form:"external_id"`
	// PayrollID filters by specific payroll ID
	PayrollID *string `json:"payroll_id" form:"payroll_id"`
	// Status filters payrolls by status
	Status *string `json:"status" form:"status"`
}

// PayrollAddressVerificationRequest verifies an address with token and amount
type PayrollAddressVerificationRequest struct {
	// VerificationCode is the verification token
	VerificationCode string `json:"token" binding:"required"`
	// Amount is the verification amount
	Amount string `json:"amount" binding:"required"`
}

// PayrollItemsListRequest lists payroll items with filters
type PayrollItemsListRequest struct {
	// Page is the page number, starting from 0
	Page int32 `json:"page" binding:"min=0" form:"page"`
	// PageSize is the number of items per page (max 50)
	PageSize int32 `json:"page_size" binding:"required,min=1,max=50" form:"page_size"`
	// StartDate filters items created after this date
	StartDate *string `json:"start_date" form:"start_date"`
	// EndDate filters items created before this date
	EndDate *string `json:"end_date" form:"end_date"`
	// EmployeeName filters items by employee name
	EmployeeName *string `json:"employee_name" form:"employee_name"`
	// ExternalID filters items by external ID
	ExternalID *string `json:"external_id" form:"external_id"`
	// PayrollID filters items by payroll ID
	PayrollID *string `json:"payroll_id" form:"payroll_id"`
	// ItemExternalID filters items by item external ID
	ItemExternalID *string `json:"item_external_id" form:"item_external_id"`
}

// PayrollGetAddressVerificationRequest gets address verification details
type PayrollGetAddressVerificationRequest struct {
	// VerificationCode is the verification code from URL path
	VerificationCode string `json:"-"`
}

// PayrollDirectItem represents a single payroll recipient in direct payroll creation
type PayrollDirectItem struct {
	// ExternalID is the external ID for the item
	ExternalID string `json:"external_id"`
	// Name is the recipient name
	Name string `json:"name"`
	// Email is the recipient email
	Email string `json:"email"`
	// Phone is the recipient phone number
	Phone string `json:"phone"`
	// Amount is the payment amount
	Amount string `json:"amount" binding:"required"`
	// Coin is the cryptocurrency to use for payment
	Coin string `json:"coin" binding:"required"`
	// Network is the blockchain network to use
	Network string `json:"network" binding:"required"`
	// Address is the recipient wallet address
	Address string `json:"address" binding:"required"`
	// PaymentMethod is the payment method ("normal" or "binance")
	PaymentMethod string `json:"payment_method"`
}

// PayrollDirectCreateRequest creates a payroll directly without CSV
type PayrollDirectCreateRequest struct {
	// ExternalID is the external ID for idempotency
	ExternalID string `json:"external_id"`
	// Items is the list of payroll recipients
	Items []PayrollDirectItem `json:"items" binding:"required,min=1"`
	// AutoApprove indicates whether to automatically approve the payroll
	AutoApprove bool `json:"auto_approve"`
}

// ================================
// Response Types
// ================================

// PayrollCreateResponse represents the response after creating a new payroll
type PayrollCreateResponse struct {
	// Payroll contains the created payroll details
	Payroll *Payroll `json:"payroll"`
	// Items contains the list of payroll items
	Items []*PayrollItems `json:"items"`
	// SwapItems contains the list of swap items
	SwapItems []*PayrollSwapItems `json:"swap_items"`
	// BinanceFromItems contains the list of Binance from items
	BinanceFromItems []*BinanceFromItems `json:"binance_from_items"`
}

// PayrollConfirmResponse represents the response after confirming a payroll
type PayrollConfirmResponse struct {
	// Payroll contains the confirmed payroll details
	Payroll *Payroll `json:"payroll"`
	// Items contains the list of payroll items
	Items []*PayrollItems `json:"items"`
	// SwapItems contains the list of swap items
	SwapItems []*PayrollSwapItems `json:"swap_items"`
	// BinanceFromItems contains the list of Binance from items
	BinanceFromItems []*BinanceFromItems `json:"binance_from_items"`
}

// PayrollGetResponse represents the response when getting a specific payroll
type PayrollGetResponse struct {
	// Payroll contains the payroll details
	Payroll *Payroll `json:"payroll"`
	// Items contains the list of payroll items
	Items []*PayrollItems `json:"items"`
	// SwapItems contains the list of swap items
	SwapItems []*PayrollSwapItems `json:"swap_items"`
}

// PayrollListResponse represents a paginated list of payrolls
type PayrollListResponse struct {
	// Payrolls is the list of payrolls
	Payrolls []*Payroll `json:"payrolls"`
	// Total is the total number of payrolls matching the filters
	Total int64 `json:"total"`
}

// PayrollAddressVerificationResponse represents address verification response
type PayrollAddressVerificationResponse struct {
	// Empty response on success
}

// PayrollItemsListResponse represents a paginated list of payroll items
type PayrollItemsListResponse struct {
	// PayrollItems is the list of payroll items
	PayrollItems []*PayrollItems `json:"payroll_items"`
	// Total is the total number of payroll items matching the filters
	Total int64 `json:"total"`
}

// PayrollGetAddressVerificationResponse represents address verification details
type PayrollGetAddressVerificationResponse struct {
	// Coin is the cryptocurrency type
	Coin string `json:"coin"`
	// Network is the blockchain network
	Network string `json:"network"`
	// Address is the wallet address
	Address string `json:"address"`
	// AmountStartDecimal is the start of the amount range for verification
	AmountStartDecimal string `json:"amount_start_decimal"`
	// AmountEndDecimal is the end of the amount range for verification
	AmountEndDecimal string `json:"amount_end_decimal"`
}

// PayrollDirectCreateResponse represents the response after direct payroll creation
type PayrollDirectCreateResponse struct {
	// Payroll contains the created payroll details
	Payroll *Payroll `json:"payroll"`
	// Items contains the list of payroll items
	Items []*PayrollItems `json:"items"`
	// SwapItems contains the list of swap items
	SwapItems []*PayrollSwapItems `json:"swap_items,omitempty"`
	// BinanceFromItems contains the list of Binance from items
	BinanceFromItems []*BinanceFromItems `json:"binance_from_items,omitempty"`
}

// ================================
// Core Types
// ================================

const (
	// PayrollIdPrefix is the prefix for payroll IDs
	PayrollIdPrefix = "payroll_"
	// PayrollIdLength is the length of the payroll ID suffix
	PayrollIdLength = 24

	// PayrollItemIdPrefix is the prefix for payroll item IDs
	PayrollItemIdPrefix = "payroll_item_"
	// PayrollItemIdLength is the length of the payroll item ID suffix
	PayrollItemIdLength = 24

	// PayrollSwapItemIdPrefix is the prefix for payroll swap item IDs
	PayrollSwapItemIdPrefix = "payroll_swap_item_"
	// PayrollSwapItemIdLength is the length of the payroll swap item ID suffix
	PayrollSwapItemIdLength = 24

	// PayrollBinanceFromItemPrefix is the prefix for Binance from item IDs
	PayrollBinanceFromItemPrefix = "payroll_bf_"
	// PayrollBinanceFromItemLength is the length of the Binance from item ID suffix
	PayrollBinanceFromItemLength = 24

	// PayrollAddressVerifyIDPrefix is the prefix for address verification IDs
	PayrollAddressVerifyIDPrefix = "payroll_av_"
	// PayrollAddressVerifyIDLength is the length of the address verification ID suffix
	PayrollAddressVerifyIDLength = 24
)

// Payroll represents a payroll batch
type Payroll struct {
	// ID is the unique identifier for the payroll
	ID string `json:"id"`
	// CreatedAt is the Unix timestamp when the payroll was created
	CreatedAt int64 `json:"created_at"`
	// UpdatedAt is the Unix timestamp when the payroll was last updated
	UpdatedAt int64 `json:"updated_at"`
	// CanceledAt is the Unix timestamp when the payroll was canceled, if applicable
	CanceledAt int64 `json:"canceled_at"`
	// MerchantID is the ID of the merchant who created the payroll
	MerchantID string `json:"merchant_id"`
	// ExternalID is the external ID for idempotency
	ExternalID string `json:"external_id"`
	// ApprovalStatus is the current approval status of the payroll
	ApprovalStatus string `json:"approval_status"`
	// Status is the current processing status of the payroll
	Status string `json:"status"`
	// TotalItemNum is the total number of items in the payroll
	TotalItemNum int64 `json:"total_item_num"`
	// TotalAmount is the total amount to be paid in the payroll
	TotalAmount string `json:"total_amount"`
	// TotalServiceFee is the total service fee for the payroll
	TotalServiceFee string `json:"total_service_fee"`
	// Currency is the currency used for the payroll
	Currency string `json:"currency"`
}

// PayrollValidation represents validation status for payroll item fields
type PayrollValidation struct {
	// NameValid indicates whether the recipient name is valid
	NameValid bool `json:"name_valid"`
	// EmailValid indicates whether the recipient email is valid
	EmailValid bool `json:"email_valid"`
	// PhoneValid indicates whether the recipient phone number is valid
	PhoneValid bool `json:"phone_valid"`
	// AmountValid indicates whether the payment amount is valid
	AmountValid bool `json:"amount_valid"`
	// CoinValid indicates whether the cryptocurrency is valid
	CoinValid bool `json:"coin_valid"`
	// NetworkValid indicates whether the blockchain network is valid
	NetworkValid bool `json:"network_valid"`
	// AddressValid indicates whether the wallet address is valid
	AddressValid bool `json:"address_valid"`
	// PaymentMethodValid indicates whether the payment method is valid (must be "normal" or "binance")
	PaymentMethodValid bool `json:"payment_method_valid"`
	// BalanceEnough indicates whether there's enough balance to complete the payment
	BalanceEnough bool `json:"balance_enough"`
}

// IsWalletInfoValid checks if all wallet-related fields are valid
func (p *PayrollValidation) IsWalletInfoValid() bool {
	return p.AmountValid && p.CoinValid && p.NetworkValid && p.AddressValid && p.PaymentMethodValid
}

// PayrollDifferenceField represents a single field's change in a payroll item
type PayrollDifferenceField struct {
	// Previous is the value before changes
	Previous string `json:"previous"`
	// Current is the value after changes
	Current string `json:"current"`
	// Status is the change status ("added", "modified", or "unchanged")
	Status string `json:"status"`
}

const (
	// PayrollDifferenceStatusAdded indicates a field was added
	PayrollDifferenceStatusAdded = "added"
	// PayrollDifferenceStatusModified indicates a field was modified
	PayrollDifferenceStatusModified = "modified"
	// PayrollDifferenceStatusUnchanged indicates a field was unchanged
	PayrollDifferenceStatusUnchanged = "unchanged"
)

// PayrollDifference represents changes made to a payroll item
type PayrollDifference struct {
	// Name contains the name field changes
	Name PayrollDifferenceField `json:"name"`
	// Email contains the email field changes
	Email PayrollDifferenceField `json:"email"`
	// Phone contains the phone field changes
	Phone PayrollDifferenceField `json:"phone"`
	// Amount contains the amount field changes
	Amount PayrollDifferenceField `json:"amount"`
	// Coin contains the coin field changes
	Coin PayrollDifferenceField `json:"coin"`
	// Network contains the network field changes
	Network PayrollDifferenceField `json:"network"`
	// Address contains the address field changes
	Address PayrollDifferenceField `json:"address"`
}

// PayrollItemParams contains parameters for a payroll item
type PayrollItemParams struct {
	// ExternalID is the external ID for the item
	ExternalID string `json:"external_id"`
	// Name is the recipient name
	Name string `json:"name"`
	// Email is the recipient email
	Email string `json:"email"`
	// Phone is the recipient phone number
	Phone string `json:"phone"`
	// Amount is the payment amount
	Amount string `json:"amount"`
	// Coin is the cryptocurrency to use
	Coin string `json:"coin"`
	// Network is the blockchain network to use
	Network string `json:"network"`
	// Address is the recipient wallet address
	Address string `json:"address"`
	// PaymentMethod is the payment method ("normal" or "binance")
	PaymentMethod string `json:"payment_method"`
}

// PayrollItems represents a single item in a payroll batch
type PayrollItems struct {
	// ID is the unique identifier for the payroll item
	ID string `json:"id"`
	// CreatedAt is the Unix timestamp when the item was created
	CreatedAt int64 `json:"created_at"`
	// UpdatedAt is the Unix timestamp when the item was last updated
	UpdatedAt int64 `json:"updated_at"`
	// PayrollID is the ID of the parent payroll
	PayrollID string `json:"payroll_id"`
	// Payroll is the reference to the parent payroll
	Payroll *Payroll `json:"payroll"`
	// ExternalID is the external ID for idempotency
	ExternalID string `json:"external_id"`
	// Name is the recipient name
	Name string `json:"name"`
	// Email is the recipient email address
	Email string `json:"email"`
	// Phone is the recipient phone number
	Phone string `json:"phone"`
	// Amount is the payment amount
	Amount string `json:"amount"`
	// ServiceFee is the service fee charged for the payment
	ServiceFee string `json:"service_fee"`
	// ExchangeRateToUSD is the exchange rate to USD at time of payment
	ExchangeRateToUSD string `json:"exchange_rate_to_usd"`
	// Coin is the cryptocurrency used for payment
	Coin string `json:"coin"`
	// Network is the blockchain network used for the transaction
	Network string `json:"network"`
	// AssetID is the asset identifier in the blockchain
	AssetID string `json:"asset_id"`
	// Address is the recipient wallet address
	Address string `json:"address"`
	// AddressVerified indicates whether the address has been verified
	AddressVerified bool `json:"address_verified"`
	// Status is the current status of the payroll item
	Status string `json:"status"`
	// Transactions contains the details of related transactions
	Transactions []*TransactionDetails `json:"transactions"`
	// Validation contains the validation status for the item fields
	Validation *PayrollValidation `json:"validation,omitempty"`
	// Differences contains the changes made to item since creation
	Differences *PayrollDifference `json:"differences,omitempty"`
	// HasMonthlyPayment indicates whether this is a recurring monthly payment
	HasMonthlyPayment bool `json:"has_monthly_payment"`
	// PaymentMethod is the payment method (normal or binance)
	PaymentMethod string `json:"payment_method"`
	// IsBinance indicates whether this is a binance payroll item
	IsBinance bool `json:"is_binance"`
	// BinanceTag is the Binance tag for the item
	BinanceTag string `json:"binance_tag"`
	// BinanceTaskID is the Binance task ID for the item
	BinanceTaskID string `json:"binance_task_id"`
	// AmlInfo is the AML information for the item
	AmlInfo string `json:"aml_info"`
	// TxId is the transaction ID (Binance order ID or blockchain tx hash)
	TxId string `json:"tx_id"`
}

// PayrollSwapItems represents a swap item for a payroll
type PayrollSwapItems struct {
	// ID is the unique identifier for the swap item
	ID string `json:"id"`
	// LifiQuoteId is the quote ID from the LiFi service
	LifiQuoteId string `json:"lifi_quote_id"`
	// CreatedAt is the Unix timestamp when the swap was created
	CreatedAt int64 `json:"created_at"`
	// UpdatedAt is the Unix timestamp when the swap was last updated
	UpdatedAt int64 `json:"updated_at"`
	// SentAt is the Unix timestamp when the swap was sent to execution
	SentAt *int64 `json:"sent_at"`
	// MerchantID is the ID of the merchant who initiated the swap
	MerchantID string `json:"merchant_id"`
	// PayrollID is the ID of the related payroll
	PayrollID string `json:"payroll_id"`
	// Payroll is the reference to the related payroll
	Payroll *Payroll `json:"payroll,omitempty"`
	// FromAssetId is the source asset identifier
	FromAssetId string `json:"from_asset_id"`
	// ExpectedFromAmount is the expected amount to be swapped from source
	ExpectedFromAmount string `json:"expected_from_amount"`
	// EstimatedFromAmount is the estimated amount to be swapped from source
	EstimatedFromAmount string `json:"estimated_from_amount"`
	// EstimatedFromAmountUsd is the USD value of estimated source amount
	EstimatedFromAmountUsd string `json:"estimated_from_amount_usd"`
	// ActualFromAmount is the actual amount swapped from source
	ActualFromAmount string `json:"actual_from_amount"`
	// NetworkFee is the network fee paid for the swap
	NetworkFee string `json:"network_fee"`
	// ToAssetId is the destination asset identifier
	ToAssetId string `json:"to_asset_id"`
	// ExpectedToAmount is the expected amount to be received
	ExpectedToAmount string `json:"expected_to_amount"`
	// EstimatedToAmount is the estimated amount to be received
	EstimatedToAmount string `json:"estimated_to_amount"`
	// EstimatedToAmountMin is the minimum estimated amount to be received
	EstimatedToAmountMin string `json:"estimated_to_amount_min"`
	// EstimatedToAmountUsd is the USD value of estimated destination amount
	EstimatedToAmountUsd string `json:"estimated_to_amount_usd"`
	// ActualToAmount is the actual amount received after swap
	ActualToAmount string `json:"actual_to_amount"`
	// Status is the current status of the swap
	Status string `json:"status"`
}

// BinanceFromItems represents a Binance transfer item for a payroll
type BinanceFromItems struct {
	// ID is the unique identifier for the Binance from item
	ID string `json:"id"`
	// PayrollID is the ID of the payroll this item belongs to
	PayrollID string `json:"payroll_id"`
	// MerchantID is the ID of the merchant
	MerchantID string `json:"merchant_id"`
	// AssetID is the asset ID to transfer from
	AssetID string `json:"asset_id"`
	// Coin is the coin type
	Coin string `json:"coin"`
	// Network is the network type
	Network string `json:"network"`
	// Amount is the amount to transfer to Binance
	Amount string `json:"amount"`
	// BinanceTaskID is the associated Binance task ID
	BinanceTaskID string `json:"binance_task_id"`
	// Status is the current status of the transfer
	Status string `json:"status"`
	// CreatedAt is the Unix timestamp when the item was created
	CreatedAt int64 `json:"created_at"`
	// UpdatedAt is the Unix timestamp when the item was last updated
	UpdatedAt int64 `json:"updated_at"`
}

// GeneratePayrollId generates a new unique payroll identifier
func GeneratePayrollId() string {
	uniqueString := uniuri.NewLen(PayrollIdLength)
	return PayrollIdPrefix + uniqueString
}

// GeneratePayrollItemId generates a new unique payroll item identifier
func GeneratePayrollItemId() string {
	uniqueString := uniuri.NewLen(PayrollItemIdLength)
	return PayrollItemIdPrefix + uniqueString
}

// GeneratePayrollSwapItemId generates a new unique payroll swap item identifier
func GeneratePayrollSwapItemId() string {
	uniqueString := uniuri.NewLen(PayrollSwapItemIdLength)
	return PayrollSwapItemIdPrefix + uniqueString
}

// GeneratePayrollBinanceFromItemId generates a new unique Binance from item identifier
func GeneratePayrollBinanceFromItemId() string {
	uniqueString := uniuri.NewLen(PayrollBinanceFromItemLength)
	return PayrollBinanceFromItemPrefix + uniqueString
}

// GeneratePayrollAddressVerifyID generates a new unique address verification identifier
func GeneratePayrollAddressVerifyID() string {
	uniqueString := uniuri.NewLen(PayrollAddressVerifyIDLength)
	return PayrollAddressVerifyIDPrefix + uniqueString
}

// GeneratePayrollAddressVerificationCode generates a new verification code for address verification
func GeneratePayrollAddressVerificationCode() string {
	uniqueString := uniuri.NewLen(PayrollAddressVerifyIDLength)
	return uniqueString
}

const (
	// PayrollPaymentMethodTypeBinance indicates a Binance payment method
	PayrollPaymentMethodTypeBinance = "binance"
	// PayrollPaymentMethodTypeNormal indicates a normal payment method
	PayrollPaymentMethodTypeNormal = "normal"

	// BinanceTaskIdPrefix is the prefix for Binance task IDs
	BinanceTaskIdPrefix = "binance_task_"
	// BinanceTaskIdLength is the length of the Binance task ID suffix
	BinanceTaskIdLength = 24

	// BinanceDepositIdPrefix is the prefix for Binance deposit IDs
	BinanceDepositIdPrefix = "binance_deposit_"
	// BinanceDepositIdLength is the length of the Binance deposit ID suffix
	BinanceDepositIdLength = 24

	// BinanceWithdrawIdPrefix is the prefix for Binance withdraw IDs
	BinanceWithdrawIdPrefix = "binance_withdraw_"
	// BinanceWithdrawIdLength is the length of the Binance withdraw ID suffix
	BinanceWithdrawIdLength = 24

	// BinanceTradeIdPrefix is the prefix for Binance trade IDs
	BinanceTradeIdPrefix = "binance_trade_"
	// BinanceTradeIdLength is the length of the Binance trade ID suffix
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

// settlement.go contains types for managing payment settlements and balance flows.
// It provides structures for tracking transaction settlements, fees, exchange rates,
// and merchant balance changes.
package developer

import (
	"github.com/shopspring/decimal"
)

const (
	// SettlementObject is the object type identifier for settlement records
	SettlementObject = "statement"
)

// ================================
// Request Types
// ================================

// SettlementListRequest represents a request to list settlement records with filtering options
type SettlementListRequest struct {
	// StartDate is the start date for filtering settlements (format: YYYY-MM-DD)
	StartDate string `json:"start_date" example:"2024-01-01"`
	// EndDate is the end date for filtering settlements (format: YYYY-MM-DD)
	EndDate string `json:"end_date" example:"2024-12-31"`
	// Page is the page number for pagination (zero-based)
	Page int32 `json:"page" binding:"min=0" example:"0"`
	// PageSize is the number of items per page
	PageSize int32 `json:"page_size" binding:"required,min=1" example:"20"`
	// Status is the settlement status to filter by (e.g., "pending", "completed", "failed")
	Status string `json:"status" example:"completed"`
}

// SettlementExportRequest represents a request to export settlement data
type SettlementExportRequest struct {
	// StartDate is the start date for the export period (format: YYYY-MM-DD)
	StartDate string `json:"start_date" example:"2024-01-01"`
	// EndDate is the end date for the export period (format: YYYY-MM-DD)
	EndDate string `json:"end_date" example:"2024-12-31"`
}

// AdminSettlementResultsListRequest represents an admin request to list settlement results
type AdminSettlementResultsListRequest struct {
	// StartDate is the start date for filtering settlements (format: YYYY-MM-DD)
	StartDate string `json:"start_date" example:"2024-01-01"`
	// EndDate is the end date for filtering settlements (format: YYYY-MM-DD)
	EndDate string `json:"end_date" example:"2024-12-31"`
	// Page is the page number for pagination (zero-based)
	Page int32 `json:"page" binding:"min=0" example:"0"`
	// PageSize is the number of items per page
	PageSize int32 `json:"page_size" binding:"required,min=1" example:"20"`
	// MerchantId is the merchant ID to filter by (optional, admin can filter by any merchant)
	MerchantId string `json:"merchant_id" example:"merchant_123456"`
	// Status is the settlement status to filter by (e.g., "pending", "completed", "failed")
	Status string `json:"status" example:"completed"`
}

// AdminSettlementResultsDetailsRequest represents a request to get detailed settlement information
type AdminSettlementResultsDetailsRequest struct {
	// SettlementId is the unique identifier of the settlement to retrieve details for
	SettlementId string `json:"statement_id" example:"settlement_123456"`
}

// AdminSettlementResultsReconcileRequest represents a request to reconcile settlement results
type AdminSettlementResultsReconcileRequest struct {
	// StartDate is the start date for the reconciliation period (format: YYYY-MM-DD)
	StartDate string `json:"start_date" example:"2024-01-01"`
	// EndDate is the end date for the reconciliation period (format: YYYY-MM-DD)
	EndDate string `json:"end_date" example:"2024-12-31"`
}

// ================================
// Response Types
// ================================

// SettlementListResponse contains paginated settlement records
type SettlementListResponse struct {
	// Settlements is the list of settlement records matching the filter criteria
	Settlements []*Settlement `json:"settlements"`
	// Total is the total number of records matching the filters
	Total int64 `json:"total" example:"150"`
	// Page is the current page number (zero-based)
	Page int32 `json:"page" example:"0"`
	// PageSize is the number of items per page
	PageSize int32 `json:"page_size" example:"20"`
}

// AdminSettlementResultsListResponse contains paginated settlement results for admin viewing
type AdminSettlementResultsListResponse struct {
	// Statements is the list of settlement result records
	Statements []*SettlementResult `json:"statements"`
	// Total is the total number of records matching the filters
	Total int64 `json:"total" example:"150"`
	// Page is the current page number (zero-based)
	Page int32 `json:"page" example:"0"`
	// PageSize is the number of items per page
	PageSize int32 `json:"page_size" example:"20"`
}

// SettlementAdjustOperationType defines types of settlement adjustments
type SettlementAdjustOperationType string

const (
	// SettlementAdjustOperationTypeManualAmount represents a manual adjustment to the settlement amount
	SettlementAdjustOperationTypeManualAmount SettlementAdjustOperationType = "MANUAL_AMOUNT"
	// SettlementAdjustOperationTypeManualFrozenAmount represents a manual adjustment to the frozen amount
	SettlementAdjustOperationTypeManualFrozenAmount SettlementAdjustOperationType = "MANUAL_FROZEN_AMOUNT"
	// SettlementAdjustOperationTypeManualTxFee represents a manual adjustment to the transaction fee
	SettlementAdjustOperationTypeManualTxFee SettlementAdjustOperationType = "MANUAL_TX_FEE"
	// SettlementAdjustOperationTypeManualTaxFee represents a manual adjustment to the tax fee
	SettlementAdjustOperationTypeManualTaxFee SettlementAdjustOperationType = "MANUAL_TAX_FEE"
	// SettlementAdjustOperationTypeManualGasFee represents a manual adjustment to the gas fee
	SettlementAdjustOperationTypeManualGasFee SettlementAdjustOperationType = "MANUAL_GAS_FEE"
)

// SettlementAdjustStatus defines status of settlement adjustments
type SettlementAdjustStatus string

const (
	// SettlementAdjustStatusPending indicates the adjustment is pending approval
	SettlementAdjustStatusPending SettlementAdjustStatus = "PENDING"
	// SettlementAdjustStatusApproved indicates the adjustment has been approved
	SettlementAdjustStatusApproved SettlementAdjustStatus = "APPROVED"
	// SettlementAdjustStatusRejected indicates the adjustment has been rejected
	SettlementAdjustStatusRejected SettlementAdjustStatus = "REJECTED"
)

// SettlementAdjust represents settlement adjustment history
type SettlementAdjust struct {
	// ID is the unique identifier for the adjustment
	ID string `json:"id"`
	// SettlementID is the ID of the settlement being adjusted
	SettlementID string `json:"statement_id"`
	// OperationType is the type of adjustment operation (e.g., MANUAL_AMOUNT, MANUAL_TX_FEE)
	OperationType SettlementAdjustOperationType `json:"operation_type"`
	// Amount is the adjusted amount as a decimal string
	Amount string `json:"amount"`
	// OriginalAmount is the original amount before adjustment as a decimal string
	OriginalAmount string `json:"original_amount"`
	// Status is the current status of the adjustment (PENDING, APPROVED, REJECTED)
	Status SettlementAdjustStatus `json:"status"`
	// Operator is the user who initiated the adjustment
	Operator string `json:"operator"`
	// Remark contains the reason or notes for the adjustment
	Remark string `json:"remark"`
	// Approver is the user who approved or rejected the adjustment
	Approver string `json:"approver"`
	// ApproveRemark contains the approver's comments
	ApproveRemark string `json:"approve_remark"`
	// CreatedAt is the Unix timestamp when the adjustment was created
	CreatedAt int64 `json:"created_at"`
	// UpdatedAt is the Unix timestamp when the adjustment was last updated
	UpdatedAt int64 `json:"updated_at"`
}

// AdminSettlementResultsDetailsResponse contains detailed settlement information for admin viewing
type AdminSettlementResultsDetailsResponse struct {
	// Settlement contains the main settlement header information
	Settlement *Settlement `json:"statement"`
	// SettlementResults contains the detailed settlement results and calculations
	SettlementResults *SettlementResult `json:"statement_results"`
	// AdjustHistory contains the history of any adjustments made to this settlement
	AdjustHistory []*SettlementAdjust `json:"adjust_history"`
}

// ================================
// Core Types
// ================================

// Settlement represents a settlement record for a payment transaction
type Settlement struct {
	// ID is the unique identifier for the settlement
	ID string `json:"id"`
	// Object is the type identifier, always "statement"
	Object string `json:"object"`
	// Type is the settlement type (e.g., "charge", "refund")
	Type string `json:"type"`
	// PaymentIntentID is the ID of the associated payment intent
	PaymentIntentID string `json:"payment_intent_id"`
	// OriginalAmount is the original transaction amount before conversion
	OriginalAmount string `json:"original_amount"`
	// OriginalCurrency is the currency of the original transaction
	OriginalCurrency string `json:"original_currency"`
	// SettlementAmount is the amount after currency conversion
	SettlementAmount string `json:"settlement_amount"`
	// SettlementCurrency is the currency used for settlement
	SettlementCurrency string `json:"settlement_currency"`
	// ExchangeRate is the exchange rate applied for currency conversion
	ExchangeRate string `json:"exchange_rate"`
	// ChargeID is the ID of the associated charge (if applicable)
	ChargeID string `json:"charge_id"`
	// RefundID is the ID of the associated refund (if applicable)
	RefundID string `json:"refund_id"`
	// FrozenAmount is the amount frozen due to AML or risk checks
	FrozenAmount string `json:"frozen_amount"`
	// GasFee is the blockchain gas fee paid for the transaction
	GasFee string `json:"gas_fee"`
	// GasFeeCurrency is the currency used for gas fee payment
	GasFeeCurrency string `json:"gas_fee_currency"`
	// GasFeeExchangeRate is the exchange rate for gas fee conversion
	GasFeeExchangeRate string `json:"gas_fee_exchange_rate"`
	// TxFee is the transaction processing fee
	TxFee string `json:"tx_fee"`
	// TaxFee is the tax amount charged on the transaction
	TaxFee string `json:"tax_fee"`
	// NetAmount is the net settlement amount after all fees and deductions
	NetAmount string `json:"net_amount"`
	// Status is the current status of the settlement
	Status string `json:"status"`
	// TxHashes is the list of blockchain transaction hashes
	TxHashes []string `json:"tx_hashes"`
	// Customer is the customer identifier for this settlement
	Customer string `json:"customer"`
	// TxTime is the Unix timestamp of the transaction
	TxTime int64 `json:"tx_time"`
}

// SettlementAbstractInfo represents summary information for a settlement transaction
type SettlementAbstractInfo struct {
	// CreatedAt is the Unix timestamp when the settlement was created
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Currency is the currency code for the settlement
	Currency *string `json:"currency,omitempty"`
	// Amount is the settlement amount
	Amount *decimal.Decimal `json:"amount,omitempty"`
	// FrozenAmount is the amount frozen due to AML or risk checks
	FrozenAmount *decimal.Decimal `json:"frozen_amount,omitempty"`
	// TxFee is the transaction processing fee
	TxFee *decimal.Decimal `json:"tx_fee,omitempty"`
	// TaxFee is the tax amount charged
	TaxFee *decimal.Decimal `json:"tax_fee,omitempty"`
	// NetAmount is the net amount after all fees and deductions
	NetAmount *decimal.Decimal `json:"net_amount,omitempty"`
	// GasCurrency is the currency used for gas fee payment
	GasCurrency *string `json:"gas_currency,omitempty"`
	// GasAmount is the blockchain gas fee amount
	GasAmount *decimal.Decimal `json:"gas_amount,omitempty"`
	// GasExchangeToUSD is the exchange rate from gas currency to USD
	GasExchangeToUSD *decimal.Decimal `json:"gas_exchange_to_usd,omitempty"`
	// ExchangeRateToPi is the exchange rate to the platform internal currency
	ExchangeRateToPi *decimal.Decimal `json:"exchange_rate_to_pi,omitempty"`
	// ExchangeRateToUSD is the exchange rate to USD
	ExchangeRateToUSD *decimal.Decimal `json:"exchange_rate_to_usd,omitempty"`
	// ExpectedAmount is the expected settlement amount
	ExpectedAmount *decimal.Decimal `json:"expected_amount,omitempty"`
	// SettleWarnings contains any warnings or issues related to the settlement
	SettleWarnings *string `json:"settle_warnings,omitempty"`
}

// SettlementResult represents detailed settlement result information with charge and refund details
type SettlementResult struct {
	// CreatedAt is the Unix timestamp when the settlement result was created
	CreatedAt int64 `json:"created_at"`
	// SettlementID is the unique identifier for the settlement
	SettlementID string `json:"statement_id"`
	// SettlementAbstractInfo contains summary information for the settlement
	SettlementAbstractInfo SettlementAbstractInfo `json:"statement_abstract_info"`
	// Merchant contains the merchant information for this settlement
	Merchant Merchant `json:"merchant"`
	// Type is the settlement type (e.g., "charge", "refund")
	Type string `json:"type"`
	// ChargeID is the ID of the associated charge (if applicable)
	ChargeID string `json:"charge_id"`
	// ChargeAbstractInfo contains summary information for the charge
	ChargeAbstractInfo SettlementAbstractInfo `json:"charge_abstract_info"`
	// RefundID is the ID of the associated refund (if applicable)
	RefundID string `json:"refund_id"`
	// RefundAbstractInfo contains summary information for the refund
	RefundAbstractInfo SettlementAbstractInfo `json:"refund_abstract_info"`
	// Status is the current status of the settlement
	Status string `json:"status"`
	// ReviewComment contains admin review comments for the settlement
	ReviewComment string `json:"review_comment"`
	// SettlementDate is the settlement date in string format (YYYY-MM-DD)
	SettlementDate string `json:"statement_date"`
}

// MerchantBalanceFlow represents a record of balance changes for a merchant account
type MerchantBalanceFlow struct {
	// CreatedAt is the Unix timestamp when the balance flow was created
	CreatedAt int64 `json:"created_at"`
	// MerchantID is the ID of the merchant whose balance changed
	MerchantID string `json:"merchant_id"`
	// Currency is the currency code for this balance flow
	Currency string `json:"currency"`
	// BusinessType is the type of business operation (e.g., "charge", "refund", "payout")
	BusinessType string `json:"business_type"`
	// FlowType is the direction of the balance flow (e.g., "in", "out")
	FlowType string `json:"flow_type"`
	// RelatedID is the ID of the related transaction or operation
	RelatedID string `json:"related_id"`
	// BeforeAvailableBalance is the available balance before this flow
	BeforeAvailableBalance decimal.Decimal `json:"before_available_balance"`
	// BeforePendingBalance is the pending balance before this flow
	BeforePendingBalance decimal.Decimal `json:"before_pending_balance"`
	// BeforeLockedBalance is the locked balance before this flow
	BeforeLockedBalance decimal.Decimal `json:"before_locked_balance"`
	// BeforeFrozenBalance is the frozen balance before this flow
	BeforeFrozenBalance decimal.Decimal `json:"before_frozen_balance"`
	// AfterAvailableBalance is the available balance after this flow
	AfterAvailableBalance decimal.Decimal `json:"after_available_balance"`
	// AfterPendingBalance is the pending balance after this flow
	AfterPendingBalance decimal.Decimal `json:"after_pending_balance"`
	// AfterLockedBalance is the locked balance after this flow
	AfterLockedBalance decimal.Decimal `json:"after_locked_balance"`
	// AfterFrozenBalance is the frozen balance after this flow
	AfterFrozenBalance decimal.Decimal `json:"after_frozen_balance"`
	// Remark contains additional notes or comments about this balance flow
	Remark string `json:"remark"`
}

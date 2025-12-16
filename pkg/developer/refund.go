// refund.go contains types for managing payment refunds.
// It provides structures for creating refunds, tracking refund status, and managing
// refund transactions for both cryptocurrency and fiat payments.
package developer

const (
	// RefundObject is the type identifier for refund objects
	RefundObject = "refund"
)

// RefundParams contains parameters for creating a refund
type RefundParams struct {
	// Amount is the amount to refund in the smallest currency unit
	Amount string `json:"amount" binding:"required"`
	// Metadata is additional metadata to attach to the refund
	Metadata map[string]string `json:"metadata"`
	// PaymentIntent is the ID of payment intent to refund
	PaymentIntent *string `json:"payment_intent_id" binding:"required"`
	// Reason is the reason for the refund
	Reason *string `json:"reason"`
	// Description is the description of the refund
	Description *string `json:"description"`
}

// Refund represents a refund transaction
type Refund struct {
	// ID is the unique identifier for the refund
	ID string `json:"id"`
	// Object is the type identifier, always "refund"
	Object string `json:"object"`
	// Amount is the amount refunded in the smallest currency unit
	Amount *AssetAmount `json:"amount"`
	// NetworkFee is the network fee charged for processing the refund
	NetworkFee *AssetAmount `json:"network_fee"`
	// NetAmount is the net amount after deducting fees
	NetAmount *AssetAmount `json:"net_amount"`
	// Created is the Unix timestamp when the refund was created
	Created int64 `json:"created"`
	// Description is the description of the refund
	Description string `json:"description"`
	// Transactions is the list of blockchain transactions associated with the refund
	Transactions []*TransactionDetails `json:"transactions"`
	// FailureReason is the reason for refund failure if applicable
	FailureReason string `json:"failure_reason"`
	// Metadata is additional metadata attached to the refund
	Metadata map[string]string `json:"metadata"`
	// Charge is the ID of charge that's refunded
	Charge *string `json:"charge"`
	// PaymentIntent is the ID of the payment intent
	PaymentIntent *string `json:"payment_intent"`
	// RefundAddress is the address where funds are refunded to
	RefundAddress *string `json:"refund_address"`
	// Reason is the reason for the refund
	Reason string `json:"reason"`
	// Status is the status of the refund (Pending, Success, Failed, Canceled)
	Status string `json:"status"`
}

// ================================
// Request Types
// ================================

// RefundCreateRequest represents a request to create a new refund
type RefundCreateRequest struct {
	RefundParams
}

// RefundGetRequest represents a request to get a refund by ID
type RefundGetRequest struct {
	// Unique identifier of the refund to retrieve
	ID string
}

// RefundListRequest represents a request to list refunds with pagination and filters
type RefundListRequest struct {
	Pagination

	// Filter refunds by payment intent ID
	PaymentIntent *string `json:"payment_intent,omitempty" form:"payment_intent"`
}

// ================================
// Response Types
// ================================

// RefundCreateResp represents the response containing newly created refunds
type RefundCreateResp struct {
	// List of created refunds
	Refunds []Refund `json:"refunds"`
}

// RefundGetResp represents the response containing a single refund's details
type RefundGetResp struct {
	Refund
}

// RefundListResp represents the response containing a paginated list of refunds
type RefundListResp struct {
	// List of refunds
	Refunds []Refund `json:"refunds"`
	// Total number of records matching the filters
	Total int64 `json:"total"`
	// Current page number
	Page int32 `json:"page"`
	// Items per page
	PageSize int32 `json:"page_size"`
}

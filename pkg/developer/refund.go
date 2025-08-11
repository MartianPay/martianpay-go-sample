package developer

const (
	RefundObject = "refund"
)

type RefundParams struct {
	Amount string `json:"amount" binding:"required"` // Amount to refund in the smallest currency unit
	// Currency string `json:"currency"`

	Metadata map[string]string `json:"metadata"` // Additional metadata to attach to the refund

	PaymentIntent *string `json:"payment_intent_id" binding:"required"` // ID of payment intent to refund

	Reason      *string `json:"reason"`      // Reason for the refund
	Description *string `json:"description"` // Description of the refund
}

type Refund struct {
	ID          string       `json:"id"`          // Unique identifier for the refund
	Object      string       `json:"object"`      // Object type, always "refund"
	Amount      *AssetAmount `json:"amount"`      // Amount refunded in the smallest currency unit
	NetworkFee  *AssetAmount `json:"network_fee"` // Network fee charged for processing the refund
	NetAmount   *AssetAmount `json:"net_amount"`  // Net amount after deducting fees
	Created     int64        `json:"created"`     // Unix timestamp when the refund was created
	Description string       `json:"description"` // Description of the refund

	Transactions  []*TransactionDetails `json:"transactions"`   // List of blockchain transactions associated with the refund
	FailureReason string                `json:"failure_reason"` // Reason for refund failure if applicable
	Metadata      map[string]string     `json:"metadata"`       // Additional metadata attached to the refund

	Charge        *string `json:"charge"`         // ID of charge that's refunded
	PaymentIntent *string `json:"payment_intent"` // ID of the payment intent

	RefundAddress *string `json:"refund_address"` // Address where funds are refunded to

	Reason string `json:"reason"` // Reason for the refund
	Status string `json:"status"` // Status of the refund (Pending, Success, Failed, Canceled)
}

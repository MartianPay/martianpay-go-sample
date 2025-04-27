package developer

const (
	RefundObject = "refund"
)

type RefundParams struct {
	Amount string `json:"amount"`
	// Currency string `json:"currency"`

	Metadata map[string]string `json:"metadata"`

	Charge        *string `json:"charge"`            // id of charge that's refunded.
	PaymentIntent *string `json:"payment_intent_id"` // paymentintent id

	Reason      *string `json:"reason"`
	Description *string `json:"description"`
}

type Refund struct {
	ID          string       `json:"id"`
	Object      string       `json:"object"` // refund
	Amount      *AssetAmount `json:"amount"`
	NetworkFee  *AssetAmount `json:"network_fee"`
	NetAmount   *AssetAmount `json:"net_amount"`
	Created     int64        `json:"created"`
	Description string       `json:"description"`

	Transactions  []*TransactionDetails `json:"transactions"` // 区块链交易列表
	FailureReason string                `json:"failure_reason"`
	Metadata      map[string]string     `json:"metadata"`

	Charge        *string `json:"charge"`         // id of charge that's refunded.
	PaymentIntent *string `json:"payment_intent"` // paymentintent id

	RefundAddress *string `json:"refund_address"` // refund address

	Reason string `json:"reason"`
	Status string `json:"status"`
}

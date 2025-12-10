package developer

// Unfreeze types
const (
	// UnfreezeTypeReverse sends unfrozen funds back to sender address (default)
	UnfreezeTypeReverse = "unfreeze_reverse"
	// UnfreezeTypeRelease adds unfrozen funds to merchant's available balance
	UnfreezeTypeRelease = "unfreeze_release"
)

// ================================
// Request Types
// ================================

// UnfreezeCreateRequest creates an unfreeze request
type UnfreezeCreateRequest struct {
	// PaymentIntentID is the ID of the payment intent with frozen funds
	PaymentIntentID string `json:"payment_intent_id" binding:"required"`
	// Type is the unfreeze type (unfreeze_reverse or unfreeze_release)
	Type string `json:"type" binding:"omitempty,oneof=unfreeze_reverse unfreeze_release" default:"unfreeze_reverse"`
	// Address is the destination address for reversed funds (required when Type is unfreeze_reverse)
	Address string `json:"address"`
	// ExternalID is an optional external identifier for idempotency
	ExternalID string `json:"external_id"`
	// Description is an optional description of the unfreeze reason
	Description string `json:"description"`
}

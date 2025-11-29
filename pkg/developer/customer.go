package developer

import "github.com/dchest/uniuri"

const (
	CustomerObject   = "customer"
	CustomerIDLength = 24
	CustomerIDPrefix = "cus_"
)

type CustomerParams struct {
	// Name is the customer's full name
	Name *string `json:"name" example:"John Doe"`

	// Email is the customer's email address (must be unique per merchant)
	Email *string `json:"email" example:"john@example.com"`

	// Description is an optional text description for internal reference
	Description *string `json:"description" example:"VIP customer"`

	// Metadata is a set of key-value pairs for storing additional information
	Metadata map[string]string `json:"metadata"`

	// Phone is the customer's phone number in E.164 format (e.g., +1234567890)
	Phone *string `json:"phone" example:"+1234567890"`
}

type Customer struct {
	// ID is the unique customer identifier (e.g., "cus_abc123")
	ID string `json:"id" example:"cus_abc123"`

	// Object is always "customer" for this resource type
	Object string `json:"object" example:"customer"`

	// TotalExpense is the total amount this customer has spent (in smallest currency unit, e.g., cents)
	TotalExpense int64 `json:"total_expense" example:"10000"`

	// TotalPayment is the total amount of successful payments (in smallest currency unit)
	TotalPayment int64 `json:"total_payment" example:"10000"`

	// TotalRefund is the total amount refunded to this customer (in smallest currency unit)
	TotalRefund int64 `json:"total_refund" example:"0"`

	// Currency is the three-letter ISO currency code
	Currency string `json:"currency,omitempty" example:"USD"`

	// Created is the Unix timestamp when this customer was created
	Created int64 `json:"created,omitempty" example:"1640000000"`

	// Name is the customer's full name
	Name *string `json:"name,omitempty" example:"John Doe"`

	// Email is the customer's email address
	Email *string `json:"email,omitempty" example:"john@example.com"`

	// Description is an optional text description for internal reference
	Description *string `json:"description,omitempty" example:"VIP customer"`

	// Metadata is a set of key-value pairs for storing additional information
	Metadata map[string]string `json:"metadata,omitempty"`

	// Phone is the customer's phone number
	Phone *string `json:"phone,omitempty" example:"+1234567890"`
}

func (c *Customer) GenerateCustomerID() string {
	return CustomerIDPrefix + uniuri.NewLen(CustomerIDLength)
}

// PaymentMethodCard represents a saved payment method card
type PaymentMethodCard struct {
	ID          string `json:"id"`          // Payment method ID (e.g., pm_1234567890)
	Provider    string `json:"provider"`    // Payment provider (stripe, paypal, etc.)
	Type        string `json:"type"`        // Payment method type (card, wallet, bank_account, etc.)
	Last4       string `json:"last4"`       // Last 4 digits of the card number
	Brand       string `json:"brand"`       // Card brand (visa, mastercard, amex, etc.)
	ExpMonth    int64  `json:"exp_month"`   // Expiration month
	ExpYear     int64  `json:"exp_year"`    // Expiration year
	Funding     string `json:"funding"`     // Card funding type (credit, debit, etc.)
	Country     string `json:"country"`     // Card issuing country
	Fingerprint string `json:"fingerprint"` // Card fingerprint for deduplication
}

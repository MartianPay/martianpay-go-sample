// charge.go contains types for managing payment charges and fraud detection.
// It provides structures for charge records and fraud assessment results.
package developer

// ChargeFraudMartianReport represents the fraud assessment result from MartianPay's system
type ChargeFraudMartianReport string

const (
	// ChargeFraudMartianReportFraudulent indicates the charge was flagged as fraudulent by MartianPay
	ChargeFraudMartianReportFraudulent ChargeFraudMartianReport = "fraudulent"
)

// ChargeFraudUserReport represents the fraud report status from the user/merchant
type ChargeFraudUserReport string

const (
	// ChargeFraudUserReportFraudulent indicates the user reported the charge as fraudulent
	ChargeFraudUserReportFraudulent ChargeFraudUserReport = "fraudulent"
	// ChargeFraudUserReportSafe indicates the user verified the charge as legitimate
	ChargeFraudUserReportSafe ChargeFraudUserReport = "safe"
)

// ChargeFraudDetails contains fraud-related information for a charge
type ChargeFraudDetails struct {
	// MartianReport is the fraud assessment from MartianPay's system
	MartianReport ChargeFraudMartianReport `json:"martian_report"`
	// UserReport is the fraud status reported by the user/merchant
	UserReport ChargeFraudUserReport `json:"user_report"`
}

// ReviewClosedReason represents the reason why a review was closed
type ReviewClosedReason string

const (
	// ReviewClosedReasonApproved indicates the review was closed because it was approved
	ReviewClosedReasonApproved ReviewClosedReason = "approved"
	// ReviewClosedReasonDisputed indicates the review was closed due to a dispute
	ReviewClosedReasonDisputed ReviewClosedReason = "disputed"
	// ReviewClosedReasonRedacted indicates the review was closed and redacted
	ReviewClosedReasonRedacted ReviewClosedReason = "redacted"
	// ReviewClosedReasonRefunded indicates the review was closed due to a refund
	ReviewClosedReasonRefunded ReviewClosedReason = "refunded"
	// ReviewClosedReasonRefundedAsFraud indicates the review was closed with a fraud refund
	ReviewClosedReasonRefundedAsFraud ReviewClosedReason = "refunded_as_fraud"
)

type ReviewIPAddressLocation struct {
	// The city where the payment originated.
	City string `json:"city"`
	// Two-letter ISO code representing the country where the payment originated.
	Country string `json:"country"`
	// The geographic latitude where the payment originated.
	Latitude float64 `json:"latitude"`
	// The geographic longitude where the payment originated.
	Longitude float64 `json:"longitude"`
	// The state/county/province/region where the payment originated.
	Region string `json:"region"`
}

// ReviewOpenedReason represents the reason why a review was opened
type ReviewOpenedReason string

// List of values that ReviewOpenedReason can take
const (
	// ReviewOpenedReasonManual indicates the review was manually opened
	ReviewOpenedReasonManual ReviewOpenedReason = "manual"
	// ReviewOpenedReasonRule indicates the review was opened by an automated rule
	ReviewOpenedReasonRule ReviewOpenedReason = "rule"
)

// ReviewReason represents the current reason for the review's status
type ReviewReason string

// List of values that ReviewReason can take
const (
	// ReviewReasonApproved indicates the review was approved
	ReviewReasonApproved ReviewReason = "approved"
	// ReviewReasonDisputed indicates the review is disputed
	ReviewReasonDisputed ReviewReason = "disputed"
	// ReviewReasonManual indicates the review is under manual review
	ReviewReasonManual ReviewReason = "manual"
	// ReviewReasonRefunded indicates the review resulted in a refund
	ReviewReasonRefunded ReviewReason = "refunded"
	// ReviewReasonRefundedAsFraud indicates the review resulted in a fraud refund
	ReviewReasonRefundedAsFraud ReviewReason = "refunded_as_fraud"
	// ReviewReasonRedacted indicates the review was redacted
	ReviewReasonRedacted ReviewReason = "redacted"
	// ReviewReasonRule indicates the review is being processed by rules
	ReviewReasonRule ReviewReason = "rule"
)

type ReviewSession struct {
	// The browser used in this browser session (e.g., `Chrome`).
	Browser string `json:"browser"`
	// Information about the device used for the browser session (e.g., `Samsung SM-G930T`).
	Device string `json:"device"`
	// The platform for the browser session (e.g., `Macintosh`).
	Platform string `json:"platform"`
	// The version for the browser session (e.g., `61.0.3163.100`).
	Version string `json:"version"`
}

// Review represents supplementary information for disputed transactions
type Review struct {
	// ID is the unique identifier for the review
	ID string `json:"id"`
	// Object is the type identifier, always "review"
	Object string `json:"object"`
	// Charge is the ID of the associated charge
	Charge *string `json:"charge"`
	// ClosedReason indicates why the review was closed
	ClosedReason ReviewClosedReason `json:"closed_reason"`
	// Created is the Unix timestamp when the review was created
	Created int64 `json:"created"`
	// IPAddress is the IP address from which the payment was made
	IPAddress string `json:"ip_address"`
	// IPAddressLocation contains the geographic location of the IP address
	IPAddressLocation *ReviewIPAddressLocation `json:"ip_address_location"`
	// Livemode indicates if this is a live transaction (true) or test (false)
	Livemode bool `json:"livemode"`
	// Open indicates whether the review is currently open
	Open bool `json:"open"`
	// OpenedReason is the reason the review was opened (rule or manual)
	OpenedReason ReviewOpenedReason `json:"opened_reason"`
	// PaymentIntent is the ID of the associated payment intent, if one exists
	PaymentIntent *string `json:"payment_intent"`
	// Reason is the current status reason (rule, manual, approved, refunded, refunded_as_fraud, disputed, or redacted)
	Reason ReviewReason `json:"reason"`
	// Session contains information about the browsing session that initiated the payment
	Session *ReviewSession `json:"session"`
}

const (
	// ChargeObject is the object type identifier for charges
	ChargeObject = "charge"
)

// Charge represents a payment charge in the system
type Charge struct {
	// ID is the unique identifier for the charge
	ID string `json:"id"`
	// Object is the type identifier, always "charge"
	Object string `json:"object"`
	// Amount contains the charged amount with asset information
	Amount *AssetAmount `json:"amount"`

	// PaymentDetails contains detailed payment information
	PaymentDetails *PaymentDetails `json:"payment_details"`
	// ExchangeRate is the exchange rate applied for currency conversion
	ExchangeRate string `json:"exchange_rate"`

	// CalculatedStatementDescriptor is the description shown on the customer's statement
	CalculatedStatementDescriptor string `json:"calculated_statement_descriptor"`
	// Captured indicates whether the charge has been captured
	Captured bool `json:"captured"`
	// Created is the Unix timestamp when the charge was created
	Created int64 `json:"created"`
	// Customer is the ID of the customer associated with this charge
	Customer string `json:"customer"`
	// Description is a description of the charge
	Description string `json:"description"`
	// Disputed indicates whether the charge has been disputed
	Disputed bool `json:"disputed"`
	// FailureCode is the error code if the charge failed
	FailureCode string `json:"failure_code"`
	// FailureMessage is the error message if the charge failed
	FailureMessage string `json:"failure_message"`
	// FraudDetails contains fraud assessment information
	FraudDetails *ChargeFraudDetails `json:"fraud_details"`
	// Livemode indicates if this is a live transaction (true) or test (false)
	Livemode bool `json:"livemode"`
	// Metadata is a set of key-value pairs for storing additional information
	Metadata map[string]string `json:"metadata"`

	// Paid indicates whether the charge was successfully paid
	Paid bool `json:"paid"`
	// PaymentIntent is the ID of the associated payment intent
	PaymentIntent *string `json:"payment_intent"`

	// PaymentMethodType indicates the type of payment method used
	PaymentMethodType PaymentMethodType `json:"payment_method_type"`
	// PaymentMethodOptions contains payment method-specific options
	PaymentMethodOptions *PaymentMethodOptions `json:"payment_method_options"`
	// Transactions contains the list of blockchain transactions for this charge
	Transactions []*TransactionDetails `json:"transactions"`

	// ReceiptEmail is the email address to send the receipt to
	ReceiptEmail string `json:"receipt_email"`
	// ReceiptURL is the URL to view the receipt for this charge
	ReceiptURL string `json:"receipt_url"`

	// Refunded indicates whether the charge has been refunded
	Refunded bool `json:"refunded"`

	// Refunds contains the list of refunds associated with this charge
	Refunds []*Refund `json:"refunds"`
	// Review contains the review information if one exists for this charge
	Review *Review `json:"review"`

	// PayerMaxPayload contains PayerMax-specific payment data
	PayerMaxPayload *PayerMaxPayload `json:"payer_max_payload,omitempty"`
	// StripePayload contains Stripe-specific payment data
	StripePayload *StripePayload `json:"stripe_payload,omitempty"`
	// PaymentProvider is the name of the payment provider (e.g., "Crypto", "Stripe", "PayPal")
	PaymentProvider string `json:"payment_provider"`
}

// PayerMaxPayload contains PayerMax-specific payment information
type PayerMaxPayload struct {
	// CashierURL is the URL to the PayerMax cashier page
	CashierURL string `json:"cashier_url"`
	// Status is the current payment status
	Status string `json:"status"`
	// SessionKey is the session key for the payment
	SessionKey string `json:"session_key"`
	// ClientKey is the client key for the payment
	ClientKey string `json:"client_key"`
}

// StripePayload contains Stripe-specific payment information
type StripePayload struct {
	// ClientSecret is the client secret for confirming the payment
	ClientSecret string `json:"client_secret"`
	// PublicKey is the Stripe publishable key
	PublicKey string `json:"public_key"`
	// Status is the current payment status
	Status string `json:"status"`
	// CustomerID is the Stripe customer ID if one exists
	CustomerID *string `json:"customer_id,omitempty"`
}

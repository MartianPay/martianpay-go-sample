// customer.go contains types for managing customers and customer authentication.
// It provides structures for customer profiles, authentication tokens, payment methods,
// and customer resolution across different authentication channels.
package developer

import "github.com/dchest/uniuri"

const (
	CustomerObject   = "customer"
	CustomerIDLength = 24
	CustomerIDPrefix = "cus_"
)

// CustomerParams contains parameters for creating or updating a customer
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

// Customer represents a customer in the MartianPay system
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

// GenerateCustomerID generates a new unique customer identifier with the 'cus_' prefix
func (c *Customer) GenerateCustomerID() string {
	return CustomerIDPrefix + uniuri.NewLen(CustomerIDLength)
}

// CustomerAddress represents a shipping or billing address from payment history
type CustomerAddress struct {
	// Country is the two-letter ISO country code
	Country string `json:"country" example:"US"`

	// State is the state/province code (optional for countries without states)
	State *string `json:"state,omitempty" example:"CA"`

	// City is the city name
	City string `json:"city" example:"San Francisco"`

	// PostalCode is the ZIP/Postal code
	PostalCode string `json:"postal_code" example:"94102"`

	// Line1 is the primary street address
	Line1 string `json:"line1" example:"123 Main St"`

	// Line2 is the secondary address information (apartment, suite, etc.)
	Line2 *string `json:"line2,omitempty" example:"Apt 4B"`
}

// CustomerTaxRegion represents a tax jurisdiction from payment history
type CustomerTaxRegion struct {
	// Country is the two-letter ISO country code for the tax jurisdiction
	Country string `json:"country" example:"US"`

	// State is the state/province code for the tax jurisdiction (optional for national tax only)
	State *string `json:"state,omitempty" example:"CA"`
}

// CustomerAuthToken represents an issued authentication token
type CustomerAuthToken struct {
	// Token is the JWT authentication token string
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`

	// TokenID is the unique identifier for this token
	TokenID string `json:"token_id" example:"token_xyz789"`

	// ExpiresAt is the Unix timestamp when the token expires
	ExpiresAt int64 `json:"expires_at" example:"1735689600"`
}

// CustomerResolveResponse contains the resolved customer identity and optional auth token
type CustomerResolveResponse struct {
	// CustomerID is the unique MartianPay customer identifier
	CustomerID string `json:"customer_id" example:"cust_abc123"`

	// Email is the customer's email address
	Email *string `json:"email,omitempty" example:"customer@example.com"`

	// Phone is the customer's phone number
	Phone *string `json:"phone,omitempty" example:"+1234567890"`

	// UUID is the merchant-assigned unique identifier for the customer
	UUID *string `json:"uuid,omitempty" example:"2b3deaec-bfe4-4519-9fcf-6248215f158e"`

	// StripeCustomerID is the associated Stripe customer ID if customer has been synced with Stripe
	StripeCustomerID *string `json:"stripe_customer_id,omitempty" example:"cus_stripe123"`

	// PayermaxCustomerID is the associated PayerMax customer ID if applicable
	PayermaxCustomerID *string `json:"payermax_customer_id,omitempty" example:"pm_cust_123"`

	// Provider identifies the channel or platform
	Provider *string `json:"provider,omitempty" example:"instagram"`

	// ChannelMetadata stores arbitrary key-value data specific to the authentication channel
	ChannelMetadata map[string]any `json:"channel_metadata,omitempty"`

	// IssuedBy tracks the issuer of the authentication request
	IssuedBy *string `json:"issued_by,omitempty" example:"instagram_bot"`

	// OrderID tracks the merchant order ID associated with this authentication
	OrderID *string `json:"order_id,omitempty" example:"ord_abc123"`

	// ReturnURL is the URL to redirect the customer (from ephemeral token)
	ReturnURL *string `json:"return_url,omitempty" example:"https://example.com/payment/success"`

	// AuthToken contains the authentication token for session management (only issued when refresh_token=true)
	AuthToken *CustomerAuthToken `json:"auth_token,omitempty"`

	// Addresses contains the customer's historical addresses from payment link orders, sorted by time (newest first), deduplicated
	Addresses []CustomerAddress `json:"addresses"`

	// TaxRegions contains the customer's historical tax regions from payment link orders, sorted by time (newest first), deduplicated
	TaxRegions []CustomerTaxRegion `json:"tax_regions"`
}

// CustomerListResponse contains paginated customer list
type CustomerListResponse struct {
	// Customers is the list of customers
	Customers []Customer `json:"customers"`
	// Total is the total number of records matching the filters
	Total int32 `json:"total"`
	// Page is the current page number
	Page int32 `json:"page"`
	// PageSize is the number of items per page
	PageSize int32 `json:"page_size"`
}

// EphemeralTokenResponse contains the ephemeral token details
type EphemeralTokenResponse struct {
	// Token is the ephemeral token string (short-lived, typically 5-15 minutes)
	Token string `json:"token" example:"eph_abc123xyz789"`
	// ExpiresAt is the Unix timestamp when the ephemeral token expires
	ExpiresAt int64 `json:"expires_at" example:"1735689600"`
}

// CustomerResolveRequest contains parameters for resolving customer identity
type CustomerResolveRequest struct {
	// PaymentLinkID is required for public (unauthenticated) requests to identify the merchant context
	PaymentLinkID *string `json:"payment_link_id,omitempty" example:"pl_abc123"`
	// MerchantID is used for subscription email scenarios to identify the merchant context
	MerchantID *string `json:"merchant_id,omitempty" example:"merch_abc123"`
	// Email is the customer's email address for email-based identity resolution
	Email *string `json:"email,omitempty" example:"customer@example.com"`
	// Phone is the customer's phone number for phone-based identity resolution (e.g., WhatsApp, Instagram DM flows)
	Phone *string `json:"phone,omitempty" example:"+1234567890"`
	// AuthToken is a previously issued authentication token for session resumption
	AuthToken *string `json:"auth_token,omitempty" example:"cust_token_abc123"`
	// EphemeralToken is a short-lived token (5-15 min) issued by merchant backend for social media integrations
	EphemeralToken *string `json:"ephemeral_token,omitempty" example:"eph_token_xyz789"`
	// Provider identifies the channel or platform (e.g., "instagram", "whatsapp", "wechat") for tracking purposes
	Provider *string `json:"provider,omitempty" example:"instagram"`
	// VerificationCode is the OTP code sent via email/SMS for saved customers requiring re-authentication
	VerificationCode *string `json:"verification_code,omitempty" example:"123456"`
	// ChannelMetadata stores arbitrary key-value data specific to the authentication channel (e.g., Instagram user ID)
	ChannelMetadata map[string]any `json:"channel_metadata,omitempty"`
	// RefreshToken indicates whether to issue a new long-lived auth token (default: false)
	RefreshToken bool `json:"refresh_token,omitempty" example:"true"`
	// AllowCreate determines whether to create a new customer if identity not found (default: true)
	AllowCreate *bool `json:"allow_create,omitempty" example:"true"`
	// IssuedBy tracks the issuer of the authentication request for audit purposes
	IssuedBy *string `json:"issued_by,omitempty" example:"instagram_bot"`
}

// CustomerCreateRequest contains parameters for creating a new customer
type CustomerCreateRequest struct {
	CustomerParams
}

// CustomerUpdateRequest contains parameters for updating an existing customer
type CustomerUpdateRequest struct {
	// ID is the customer ID
	ID string `json:"id"`
	CustomerParams
}

// CustomerGetRequest contains parameters for retrieving a customer
type CustomerGetRequest struct {
	// ID is the customer ID
	ID string `json:"id"`
}

// CustomerListRequest contains parameters for listing customers
type CustomerListRequest struct {
	Pagination
	// Email is the filter by email
	Email *string `json:"email" form:"email"`
}

// EphemeralTokenRequest contains parameters for issuing ephemeral tokens
type EphemeralTokenRequest struct {
	// IDPKey is the identity provider type - determines how to identify the customer
	// Valid values: "email", "phone", "uuid"
	// Example: "email" means the customer is identified by their email address
	// Optional: If not provided, creates an anonymous ephemeral token (only order_id based)
	IDPKey string `json:"idp_key,omitempty" example:"email"`

	// IDPSubject is the unique identifier under the specified IDP
	// When idp_key="email", this should be the email address (e.g., "user@example.com")
	// When idp_key="phone", this should be the phone number (e.g., "+1234567890")
	// When idp_key="uuid", this should be a UUID
	// Optional: If not provided along with idp_key, creates an anonymous ephemeral token
	IDPSubject string `json:"idp_subject,omitempty" example:"user@example.com"`

	// Provider identifies the channel or platform initiating the request
	// Examples: "instagram", "whatsapp", "wechat", "telegram"
	// Used for tracking and channel-specific business logic
	// Optional: If not provided, creates an anonymous ephemeral token
	Provider string `json:"provider,omitempty" example:"instagram"`

	// AllowCreate determines whether to create a new customer if not found (default: true)
	// Set to false to return an error if customer doesn't exist
	AllowCreate *bool `json:"allow_create,omitempty" example:"true"`

	// ChannelMetadata stores arbitrary channel-specific data
	// Examples: {"instagram_user_id": "12345", "dm_thread_id": "67890"}
	ChannelMetadata map[string]any `json:"channel_metadata,omitempty"`

	// IssuedBy identifies the system or service issuing this token request
	// Used for audit logging and tracking token origin
	IssuedBy *string `json:"issued_by,omitempty" example:"instagram_commerce_bot"`

	// ReturnURL is the URL to redirect the customer after authentication/payment
	// Used in social media integrations where the flow needs to return to a specific page
	ReturnURL *string `json:"return_url,omitempty" example:"https://example.com/payment/success"`

	// OrderID is an optional order identifier to associate with this ephemeral token
	// When provided, this order_id will be returned when the token is resolved
	// Useful for tracking which order the customer authentication is for
	OrderID *string `json:"order_id,omitempty" example:"ord_abc123"`
}

// ================================
// Payment Method Types
// ================================

// PaymentMethodCard represents a saved payment method card
type PaymentMethodCard struct {
	// ID is the payment method ID (e.g., pm_1234567890)
	ID string `json:"id"`
	// Provider is the payment provider (stripe, paypal, etc.)
	Provider string `json:"provider"`
	// Type is the payment method type (card, wallet, bank_account, etc.)
	Type string `json:"type"`
	// Last4 is the last 4 digits of the card number
	Last4 string `json:"last4"`
	// Brand is the card brand (visa, mastercard, amex, etc.)
	Brand string `json:"brand"`
	// ExpMonth is the expiration month
	ExpMonth int64 `json:"exp_month"`
	// ExpYear is the expiration year
	ExpYear int64 `json:"exp_year"`
	// Funding is the card funding type (credit, debit, etc.)
	Funding string `json:"funding"`
	// Country is the card issuing country
	Country string `json:"country"`
	// Fingerprint is the card fingerprint for deduplication
	Fingerprint string `json:"fingerprint"`
}

// PaymentMethodListResponse represents the response for listing payment methods
type PaymentMethodListResponse struct {
	// PaymentMethods is the list of saved payment methods
	PaymentMethods []*PaymentMethodCard `json:"payment_methods"`
}

// CustomerPaymentMethodListRequest lists payment methods for a specific customer (merchant auth)
type CustomerPaymentMethodListRequest struct {
	// CustomerID is the customer ID to list payment methods for
	CustomerID string `form:"customer_id" binding:"required"`
}

// CustomerPaymentMethodPublicListRequest lists payment methods for a customer (public/no auth)
type CustomerPaymentMethodPublicListRequest struct {
	// CustomerID is the customer ID
	CustomerID string `form:"customer_id" binding:"required"`
	// PaymentLinkID is the optional payment link ID (either this or public_key is required)
	PaymentLinkID string `form:"payment_link_id"`
	// PublicKey is the optional merchant public key (either this or payment_link_id is required)
	PublicKey string `form:"public_key"`
	// ClientSecret is the payment intent client secret (required for authentication)
	ClientSecret string `form:"client_secret" binding:"required"`
}

// ListCustomerPaymentMethodsRequest lists payment methods for authenticated customer
type ListCustomerPaymentMethodsRequest struct {
	// No fields - customer is identified via Authorization Bearer token in middleware
}

// GetCustomerPaymentMethodRequest gets a specific payment method by ID
type GetCustomerPaymentMethodRequest struct {
	// PaymentMethodID is the payment method ID from URL path
	PaymentMethodID string `json:"-" form:"-"`
}

// DeleteCustomerPaymentMethodRequest deletes a saved payment method
type DeleteCustomerPaymentMethodRequest struct {
	// PaymentMethodID is the payment method ID from URL path
	PaymentMethodID string `json:"-"`
}

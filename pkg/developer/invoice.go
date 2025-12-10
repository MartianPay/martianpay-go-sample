package developer

// ================================
// Request Types
// ================================

// ListCustomerInvoicesRequest represents a request to list a customer's invoices
type ListCustomerInvoicesRequest struct {
	// SubscriptionID filters invoices by subscription ID
	SubscriptionID *string `json:"subscription_id" form:"subscription_id"`
	// Status filters invoices by status
	Status *string `json:"status" form:"status"`
	// Offset is the number of records to skip for pagination
	Offset int `json:"offset" form:"offset"`
	// Limit is the maximum number of records to return
	Limit int `json:"limit" form:"limit"`
}

// GetCustomerInvoiceRequest gets specific invoice for authenticated customer
type GetCustomerInvoiceRequest struct {
	// InvoiceID is the invoice ID from URL path
	InvoiceID string `json:"-"`
}

// GetCustomerInvoicePaymentIntentRequest gets payment intent for customer invoice
type GetCustomerInvoicePaymentIntentRequest struct {
	// InvoiceID is the invoice ID from URL path
	InvoiceID string `json:"-"`
}

// ListMerchantInvoicesRequest represents a request to list merchant invoices with filters
type ListMerchantInvoicesRequest struct {
	// CustomerID filters invoices by customer ID
	CustomerID *string `json:"customer_id" form:"customer_id"`
	// SubscriptionID filters invoices by subscription ID
	SubscriptionID *string `json:"subscription_id" form:"subscription_id"`
	// Status filters invoices by status
	Status *string `json:"status" form:"status"`
	// ExternalID filters invoices by external ID
	ExternalID *string `json:"external_id" form:"external_id"`
	// Offset is the number of records to skip for pagination
	Offset int `json:"offset" form:"offset"`
	// Limit is the maximum number of records to return
	Limit int `json:"limit" form:"limit"`
}

// GetMerchantInvoiceRequest gets specific invoice for merchant management
type GetMerchantInvoiceRequest struct {
	// InvoiceID is the invoice ID from URL path
	InvoiceID string `json:"-"`
}

// GetMerchantInvoicePaymentIntentRequest gets payment intent for merchant invoice
type GetMerchantInvoicePaymentIntentRequest struct {
	// InvoiceID is the invoice ID from URL path
	InvoiceID string `json:"-"`
}

// SendMerchantInvoiceRequest sends invoice to customer
type SendMerchantInvoiceRequest struct {
	// InvoiceID is the invoice ID from URL path
	InvoiceID string `json:"-"`
}

// VoidMerchantInvoiceRequest voids an invoice
type VoidMerchantInvoiceRequest struct {
	// InvoiceID is the invoice ID from URL path
	InvoiceID string `json:"-"`
}

// DownloadInvoicePDFPublicRequest downloads invoice PDF (public access)
type DownloadInvoicePDFPublicRequest struct {
	// InvoiceID is the invoice ID from URL path
	InvoiceID string `json:"-"`
}

// DownloadCustomerInvoicePDFRequest downloads invoice PDF for customer
type DownloadCustomerInvoicePDFRequest struct {
	// InvoiceID is the invoice ID from URL path
	InvoiceID string `json:"-"`
}

// DownloadMerchantInvoicePDFRequest downloads invoice PDF for merchant
type DownloadMerchantInvoicePDFRequest struct {
	// InvoiceID is the invoice ID from URL path
	InvoiceID string `json:"-"`
}

// ================================
// Response Types
// ================================

// ListInvoicesResponse represents a paginated list of invoices
type ListInvoicesResponse struct {
	Data   []*InvoiceDetails `json:"data"`
	Total  int64             `json:"total"`
	Offset int               `json:"offset"`
	Limit  int               `json:"limit"`
}

// InvoiceLineItemDetails represents a line item in API responses with converted amounts
type InvoiceLineItemDetails struct {
	// Description is the line item description
	Description string `json:"description"`
	// ProductID is the ID of the associated product
	ProductID *string `json:"product_id,omitempty"`
	// VariantID is the ID of the associated product variant
	VariantID *string `json:"variant_id,omitempty"`
	// Quantity is the number of items
	Quantity int `json:"quantity"`
	// UnitAmount is the unit amount converted from cents to actual value
	UnitAmount string `json:"unit_amount"`
	// Amount is the total amount converted from cents to actual value
	Amount string `json:"amount"`
	// Currency is the currency code
	Currency string `json:"currency"`
	// PeriodStart is the Unix timestamp when the billing period starts
	PeriodStart *int64 `json:"period_start,omitempty"`
	// PeriodEnd is the Unix timestamp when the billing period ends
	PeriodEnd *int64 `json:"period_end,omitempty"`
	// Metadata is additional line item metadata
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// ProductName is the product name for display (fetched via JOIN/lookup, not stored)
	ProductName *string `json:"product_name,omitempty"`
	// VariantTitle is the variant title for display (fetched via JOIN/lookup, not stored)
	VariantTitle *string `json:"variant_title,omitempty"`
}

// InvoiceDetails represents an invoice in API responses
type InvoiceDetails struct {
	// ID is the unique identifier for the invoice
	ID string `json:"id"`
	// MerchantID is the ID of the merchant who owns this invoice
	MerchantID string `json:"merchant_id"`
	// SubscriptionID is the ID of the associated subscription
	SubscriptionID *string `json:"subscription_id,omitempty"`
	// PaymentIntentID is the ID of the associated payment intent
	PaymentIntentID *string `json:"payment_intent_id,omitempty"`
	// CustomerID is the ID of the customer this invoice belongs to
	CustomerID string `json:"customer_id"`
	// Amount is the invoice amount converted from cents to actual value
	Amount string `json:"amount"`
	// Currency is the currency code
	Currency string `json:"currency"`
	// Status is the current status of the invoice
	Status string `json:"status"`
	// BillingReason is the reason this invoice was created
	BillingReason *string `json:"billing_reason,omitempty"`
	// AttemptCount is the number of payment attempts
	AttemptCount int `json:"attempt_count"`
	// DueDate is the Unix timestamp when payment is due
	DueDate *int64 `json:"due_date,omitempty"`
	// PaidAt is the Unix timestamp when the invoice was paid
	PaidAt *int64 `json:"paid_at,omitempty"`
	// VoidedAt is the Unix timestamp when the invoice was voided
	VoidedAt *int64 `json:"voided_at,omitempty"`
	// Lines is the list of line items in the invoice
	Lines []InvoiceLineItemDetails `json:"lines,omitempty"`
	// Metadata is additional metadata for the invoice
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// ExternalID is the merchant's external identifier for this invoice
	ExternalID *string `json:"external_id,omitempty"`
	// CreatedAt is the Unix timestamp when the invoice was created
	CreatedAt int64 `json:"created_at"`
	// UpdatedAt is the Unix timestamp when the invoice was last updated
	UpdatedAt int64 `json:"updated_at"`
	// Version is the version number for optimistic locking
	Version int64 `json:"version"`
	// MerchantName is the merchant name for display (fetched via JOIN/lookup, not stored)
	MerchantName *string `json:"merchant_name,omitempty"`
	// PublicKey is the merchant public key for client-side integration (fetched via JOIN/lookup, not stored)
	PublicKey *string `json:"public_key,omitempty"`
	// CustomerEmail is the customer email for display (fetched via JOIN/lookup, not stored)
	CustomerEmail *string `json:"customer_email,omitempty"`
	// CustomerName is the customer name for display (fetched via JOIN/lookup, not stored)
	CustomerName *string `json:"customer_name,omitempty"`
	// SubscriptionProductName is the product name from subscription (fetched via JOIN/lookup, not stored)
	SubscriptionProductName *string `json:"subscription_product_name,omitempty"`
	// SubscriptionProductVariant is the variant title from subscription (fetched via JOIN/lookup, not stored)
	SubscriptionProductVariant *string `json:"subscription_product_variant,omitempty"`
}

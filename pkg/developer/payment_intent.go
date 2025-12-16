// payment_intent.go contains types for managing payment intents and payment processing.
// It provides structures for payment intents, charges, transactions, and payment method options
// for both cryptocurrency and fiat payment processing.
package developer

import (
	"errors"
	"fmt"
	"strings"
)

const (
	PaymentIntentObject = "payment_intent"
)

// PaymentIntentStatus represents the current state of a payment intent
type PaymentIntentStatus string

const (
	// PaymentIntentStatusCreated indicates the payment intent has just been created and is waiting for user to select payment method
	PaymentIntentStatusCreated PaymentIntentStatus = "Created"
	// PaymentIntentStatusWaiting indicates the payment method has been selected and is waiting for payment
	PaymentIntentStatusWaiting PaymentIntentStatus = "Waiting"
	// PaymentIntentStatusPartiallyPaid indicates at least one blockchain tx with status Submitted/Completed/Confirmed exists, but total amount is insufficient
	PaymentIntentStatusPartiallyPaid PaymentIntentStatus = "Partially Paid"
	// PaymentIntentStatusPaid indicates at least one blockchain tx with status Submitted/Completed/Confirmed exists, with sufficient total amount
	PaymentIntentStatusPaid PaymentIntentStatus = "Paid"
	// PaymentIntentStatusCompleted indicates at least one blockchain tx with status Completed/Confirmed exists, with sufficient total amount
	PaymentIntentStatusCompleted PaymentIntentStatus = "Completed"
	// PaymentIntentStatusConfirmed indicates at least one blockchain tx with status Confirmed exists, with sufficient total amount
	PaymentIntentStatusConfirmed PaymentIntentStatus = "Confirmed"
	// PaymentIntentStatusFrozen indicates payment has been received but funds are frozen due to AML rejection
	PaymentIntentStatusFrozen PaymentIntentStatus = "Frozen"
	// PaymentIntentStatusUnfrozen indicates previously frozen funds have been successfully unfrozen
	PaymentIntentStatusUnfrozen PaymentIntentStatus = "Unfrozen"
	// PaymentIntentStatusCancelled indicates the payment intent was cancelled due to timeout (20min for Created, 60min for Partially Paid/Paid, 120min for Completed)
	PaymentIntentStatusCancelled PaymentIntentStatus = "Cancelled"
)

// TransactionDetails represents detailed information about a blockchain transaction
type TransactionDetails struct {
	// TxId is the unique transaction identifier
	TxId string `json:"tx_id"`
	// SourceAddress is the sender's blockchain address
	SourceAddress string `json:"source_address"`
	// DestinationAddress is the recipient's blockchain address
	DestinationAddress string `json:"destination_address"`
	// TxHash is the blockchain transaction hash
	TxHash string `json:"tx_hash"`
	// Amount is the transaction amount
	Amount string `json:"amount"`
	// Decimals is the number of decimal places for the asset
	Decimals int `json:"decimals"`
	// AssetId is the asset identifier for the token
	AssetId string `json:"asset_id"`
	// Token is the token name (e.g., USDT, BTC, ETH)
	Token string `json:"token"`
	// Network is the blockchain network type (e.g., TRC20, ERC20)
	Network string `json:"network"`
	// Type is the transaction type (0: deposit, 1: refund)
	Type string `json:"type"`
	// CreatedAt is the Unix timestamp when the transaction was created
	CreatedAt int64 `json:"created_at"`
	// Status is the current transaction status
	Status string `json:"status"`
	// AmlStatus is the AML check status ('', 'approved', 'rejected')
	AmlStatus string `json:"aml_status"`
	// AmlInfo contains AML check details
	AmlInfo *string `json:"aml_info"`
	// ChargeId is the ID of the associated charge
	ChargeId string `json:"charge_id"`
	// RefundId is the ID of the associated refund (only valid for type==1)
	RefundId string `json:"refund_id"`
	// FeeInfo contains details of the transaction's fee
	FeeInfo *string `json:"fee_info"`
	// FeeCurrency is the asset type used to pay the fee (ETH for ERC-20 tokens, BTC for Omni, etc.)
	FeeCurrency string `json:"fee_currency"`
}

// PaymentMethodType represents the type of payment method used for a payment intent
type PaymentMethodType string

const (
	PaymentMethodTypeCrypto        PaymentMethodType = "crypto"
	PaymentMethodTypeTron          PaymentMethodType = "tron"
	PaymentMethodTypeWalletConnect PaymentMethodType = "wallet_connect"
	PaymentMethodTypeCards         PaymentMethodType = "cards"
	PaymentMethodTypeOthers        PaymentMethodType = "others"
)

// PaymentMethodOptions contains available payment method options for a payment intent
type PaymentMethodOptions struct {
	// Crypto contains cryptocurrency payment options
	Crypto *Crypto `json:"crypto"`
	// Fiat contains fiat currency payment options
	Fiat *Fiat `json:"fiat"`
}

// PaymentMethodConfirmOptions contains simplified options for confirming a payment method
type PaymentMethodConfirmOptions struct {
	// Crypto contains cryptocurrency confirmation options
	Crypto *CryptoOption `json:"crypto"`
	// Fiat contains fiat currency confirmation options
	Fiat *FiatOption `json:"fiat"`
}

// FiatOption contains fiat payment confirmation parameters
type FiatOption struct {
	// Currency is the currency code for the fiat payment
	Currency *string `json:"currency" binding:"required"`
	// PaymentMethodID is the saved payment method ID for Stripe (used with cards payment type)
	PaymentMethodID *string `json:"payment_method_id,omitempty"`
	// SavePaymentMethod indicates whether to save the payment method for future use (Stripe only)
	SavePaymentMethod *bool `json:"save_payment_method,omitempty"`
}

// CryptoOption contains cryptocurrency payment confirmation parameters
type CryptoOption struct {
	// AssetId is the asset identifier for the token
	AssetId *string `json:"asset_id" binding:"required"`
}

// Fiat contains fiat currency payment method details
type Fiat struct {
	// Amount is the amount of fiat to be paid
	Amount *string `json:"amount"`
	// AssetId is the asset identifier for the fiat currency
	AssetId *string `json:"asset_id" binding:"required"`
	// Currency is the currency code for the fiat
	Currency *string `json:"currency" binding:"required"`
	// Decimals is the number of decimal places for the fiat currency
	Decimals *int `json:"decimals"`
	// ExchangeRate is the exchange rate for the fiat currency
	ExchangeRate *string `json:"exchange_rate"`
}

// Crypto contains cryptocurrency payment method details
type Crypto struct {
	// Amount is the amount of crypto to be paid
	Amount *string `json:"amount"`
	// Token is the token name (e.g., USDT, BTC, ETH)
	Token *string `json:"token"`
	// AssetId is the asset identifier for the token
	AssetId *string `json:"asset_id" binding:"required"`
	// Network is the blockchain network type (e.g., TRC20, ERC20)
	Network *string `json:"network"`
	// Decimals is the number of decimal places for the token
	Decimals *int `json:"decimals"`
	// ExchangeRate is the exchange rate for the token
	ExchangeRate *string `json:"exchange_rate"`
	// DepositAddress is the deposit address for receiving the crypto payment
	DepositAddress *string `json:"deposit_address"`
}

// PaymentIntentParams contains parameters for creating a payment intent
type PaymentIntentParams struct {
	// Amount is the amount to be charged in the specified currency (required for traditional mode, optional for payment link mode)
	Amount string `json:"amount" example:"1.0"`
	// Currency is the ISO 4217 currency code (required for traditional mode, optional for payment link mode)
	Currency string `json:"currency" example:"USD"`
	// Customer is the identifier of the customer this payment intent belongs to
	Customer *string `json:"customer"`
	// Description is an arbitrary string attached to the object, often useful for displaying to users
	Description *string `json:"description"`
	// Metadata is a set of key-value pairs that you can attach to the object
	Metadata map[string]string `json:"metadata"`
	// MerchantOrderId is an optional unique identifier for the order on the merchant's side (must be unique per merchant if provided)
	MerchantOrderId string `json:"merchant_order_id" example:"sn-1234567890"`
	// ReceiptEmail is the email address to send the receipt to
	ReceiptEmail string `json:"receipt_email" binding:"omitempty,email"`
	// ReturnURL is the URL to redirect the customer to after payment
	ReturnURL *string `json:"return_url"`
	// PaymentMethodID is the ID of a saved payment method to use for this payment (Stripe only)
	PaymentMethodID *string `json:"payment_method_id"`
}

// PaymentDetails contains detailed payment information including amounts and fees
type PaymentDetails struct {
	// Settled indicates whether the payment has been settled
	Settled bool `json:"-"`
	// AmountCaptured is the amount that has been captured from the customer
	AmountCaptured *AssetAmount `json:"amount_captured"`
	// AmountRefunded is the amount that has been refunded to the customer
	AmountRefunded *AssetAmount `json:"amount_refunded"`
	// TxFee is the transaction fee charged by the payment processor
	TxFee *AssetAmount `json:"tx_fee"`
	// TaxFee is the tax amount collected for the transaction
	TaxFee *AssetAmount `json:"tax_fee"`
	// FrozenAmount is the amount temporarily held/frozen for pending operations
	FrozenAmount *AssetAmount `json:"frozen_amount"`
	// NetAmount is the net amount after deducting all fees
	NetAmount *AssetAmount `json:"net_amount"`
	// GasFee contains gas fees for blockchain transactions by asset type
	GasFee map[string]*AssetAmount `json:"gas_fee"`
	// NetworkFee is the network fee for processing the transaction
	NetworkFee *AssetAmount `json:"network_fee"`
}

// UnfreezeWithdraw represents an unfreeze withdrawal of previously frozen funds
type UnfreezeWithdraw struct {
	// ID is the unique identifier for the withdraw
	ID string `json:"id"`
	// PayoutID is the associated payout ID
	PayoutID string `json:"payout_id"`
	// AssetID is the asset ID of the withdrawn funds
	AssetID string `json:"asset_id"`
	// Amount is the amount being withdrawn
	Amount *AssetAmount `json:"amount"`
	// NetworkFee is the network fee for the withdrawal
	NetworkFee *AssetAmount `json:"network_fee"`
	// Status is the current status of the withdrawal
	Status string `json:"status"`
	// Address is the destination address for withdrawal
	Address string `json:"address"`
	// Type is the type of withdraw (unfreeze_reverse or unfreeze_release)
	Type string `json:"type"`
	// Description is the description from the associated payout
	Description string `json:"description"`
	// OriginalFrozenTxID is the original AML-rejected transaction ID
	OriginalFrozenTxID string `json:"original_frozen_tx_id"`
	// CreatedAt is the Unix timestamp when the withdraw was created
	CreatedAt int64 `json:"created_at"`
	// UpdatedAt is the Unix timestamp when the withdraw was last updated
	UpdatedAt int64 `json:"updated_at"`
}

/*
PaymentIntentStatus represents the lifecycle status of a payment intent:

1. **Created**
  - Payment intent has just been created and is waiting for the user to select a payment method.

2. **Waiting**
  - User has selected a payment method and is waiting to make the payment.

3. **Partially Paid**
  - At least one associated Charge contains a blockchain transaction with status Submitted, Completed, or Confirmed,
    but the total amount (sum of all transactions if multiple) is insufficient.

4. **Paid**
  - At least one associated Charge contains a blockchain transaction with status Submitted, Completed, or Confirmed,
    and the total amount (sum of all transactions if multiple) is sufficient.

5. **Completed**
  - At least one associated Charge contains a blockchain transaction with status Completed or Confirmed,
    and the total amount (sum of all transactions if multiple) is sufficient.

6. **Confirmed**
  - At least one associated Charge contains a blockchain transaction with status Confirmed,
    and the total amount (sum of all transactions if multiple) is sufficient.

7. **Cancelled**
  - After 20 minutes (or configured timeout), payment intent status is still Created with no payment transaction generated.
  - After 60 minutes (or configured timeout), payment intent status is still Partially Paid or Paid with no valid transaction or insufficient gas to be included on-chain.
  - After 120 minutes (or configured timeout), payment intent status is still Completed but the on-chain transaction disappeared due to blockchain reorganization.
*/
type PaymentIntent struct {
	// ID is the unique identifier for the payment intent
	ID string `json:"id"`
	// Object is the object type, always "payment_intent"
	Object string `json:"object"`
	// Amount is the amount to be collected by this payment intent
	Amount *AssetAmount `json:"amount"`
	// PaymentDetails contains details about the payment including settled amounts and fees
	PaymentDetails *PaymentDetails `json:"payment_details"`
	// CanceledAt is the Unix timestamp when the payment intent was canceled
	CanceledAt int64 `json:"canceled_at"`
	// CancellationReason is the reason why the payment intent was canceled
	CancellationReason string `json:"cancellation_reason"`
	// ClientSecret is the client secret used for client-side confirmation
	ClientSecret string `json:"client_secret"`
	// Created is the Unix timestamp when the payment intent was created
	Created int64 `json:"created"`
	// Updated is the Unix timestamp when the payment intent was last updated
	Updated int64 `json:"updated"`
	// Currency is the currency of the payment intent
	Currency string `json:"currency"`
	// Customer is the customer associated with this payment intent
	Customer *Customer `json:"customer"`
	// Description is the description of the payment intent
	Description string `json:"description"`
	// Livemode indicates whether the object exists in live mode (true) or test mode (false)
	Livemode bool `json:"livemode"`
	// Metadata contains key-value pairs attached to the payment intent
	Metadata map[string]interface{} `json:"metadata"`
	// PaymentLinkDetails contains details about the payment link if this payment intent was created through a payment link
	PaymentLinkDetails *PaymentLinkDetails `json:"payment_link_details"`
	// MerchantOrderId is the merchant-specified order ID associated with this payment intent
	MerchantOrderId string `json:"merchant_order_id"`
	// Charges is the list of charges associated with this payment intent
	Charges []*Charge `json:"charges"`
	// ReceiptEmail is the email address to send the receipt to
	ReceiptEmail string `json:"receipt_email"`
	// ReturnURL is the URL to redirect the customer to after payment
	ReturnURL string `json:"return_url"`
	// Status is the current status of the payment intent
	Status string `json:"status"`
	// ProcessingStatus is the internal processing status of the payment intent
	ProcessingStatus string `json:"-"`
	// PaymentIntentStatus is the standardized status of the payment intent
	PaymentIntentStatus PaymentIntentStatus `json:"payment_intent_status"`
	// CompleteOnFirstPayment indicates whether to complete payment immediately when first transaction arrives, regardless of amount
	CompleteOnFirstPayment bool `json:"complete_on_first_payment"`
	// PermanentDeposit indicates if this payment is a permanent deposit
	PermanentDeposit bool `json:"permanent_deposit"`
	// PermanentDepositAssetId is the asset ID for the permanent deposit
	PermanentDepositAssetId string `json:"permanent_deposit_asset_id"`
	// ExpiredAt is the Unix timestamp when the payment intent expires
	ExpiredAt uint64 `json:"expired_at"`
	// UnfreezeWithdraws is the list of unfreeze withdrawals for previously frozen funds
	UnfreezeWithdraws []*UnfreezeWithdraw `json:"unfreeze_withdraws"`
	// Invoice is the invoice ID if this payment is for an invoice (subscriptions)
	Invoice *string `json:"invoice,omitempty"`
	// InvoiceDetails contains expanded invoice details
	InvoiceDetails *InvoiceDetails `json:"invoice_details,omitempty"`
	// Subscription is the subscription ID (resolved through invoice)
	Subscription *string `json:"subscription,omitempty"`
	// SubscriptionDetails contains expanded subscription details
	SubscriptionDetails *SubscriptionDetails `json:"subscription_details,omitempty"`
	// MerchantID is the merchant ID who owns this payment intent
	MerchantID *string `json:"merchant_id,omitempty"`
	// MerchantName is the merchant name for display
	MerchantName *string `json:"merchant_name,omitempty"`
}

// PaymentIntentCancellationReason represents the reason why a payment intent was cancelled
type PaymentIntentCancellationReason string

const (
	// PaymentIntentCancellationReasonAbandoned indicates the payment intent was abandoned by the customer
	PaymentIntentCancellationReasonAbandoned PaymentIntentCancellationReason = "abandoned"
	// PaymentIntentCancellationReasonAutomatic indicates the payment intent was automatically cancelled by the system
	PaymentIntentCancellationReasonAutomatic PaymentIntentCancellationReason = "automatic"
	// PaymentIntentCancellationReasonDuplicate indicates the payment intent was cancelled due to duplication
	PaymentIntentCancellationReasonDuplicate PaymentIntentCancellationReason = "duplicate"
	// PaymentIntentCancellationReasonFailedInvoice indicates the payment intent was cancelled due to a failed invoice
	PaymentIntentCancellationReasonFailedInvoice PaymentIntentCancellationReason = "failed_invoice"
	// PaymentIntentCancellationReasonFraudulent indicates the payment intent was cancelled due to suspected fraud
	PaymentIntentCancellationReasonFraudulent PaymentIntentCancellationReason = "fraudulent"
	// PaymentIntentCancellationReasonRequestedByCustomer indicates the payment intent was cancelled at the customer's request
	PaymentIntentCancellationReasonRequestedByCustomer PaymentIntentCancellationReason = "requested_by_customer"
	// PaymentIntentCancellationReasonVoidInvoice indicates the payment intent was cancelled due to a voided invoice
	PaymentIntentCancellationReasonVoidInvoice PaymentIntentCancellationReason = "void_invoice"
)

// ================================
// Request Types
// ================================

// VariantSelectionRequest represents a product variant selection
type VariantSelectionRequest struct {
	// VariantID is the ID of the product variant being selected
	VariantID string `json:"variant_id" binding:"required"`
	// Quantity is the quantity of the variant being purchased
	Quantity int `json:"quantity" binding:"required"`
	// SellingPlanID is the optional ID of the selling plan (subscription) for this variant
	SellingPlanID *string `json:"selling_plan_id,omitempty"`
}

// PaymentIntentShippingAddress represents shipping destination
type PaymentIntentShippingAddress struct {
	// Country is the country code for the shipping address
	Country string `json:"country" binding:"required"`
	// State is the optional state or province for the shipping address
	State *string `json:"state,omitempty"`
	// City is the city for the shipping address
	City string `json:"city" binding:"required"`
	// PostalCode is the postal or ZIP code for the shipping address
	PostalCode string `json:"postal_code" binding:"required"`
	// Line1 is the first line of the street address
	Line1 string `json:"line1" binding:"required"`
	// Line2 is the optional second line of the street address
	Line2 *string `json:"line2,omitempty"`
}

// PaymentIntentTaxRegion represents tax jurisdiction
type PaymentIntentTaxRegion struct {
	// Country is the country code for the tax jurisdiction
	Country string `json:"country" binding:"required"`
	// State is the optional state or province for the tax jurisdiction
	State *string `json:"state,omitempty"`
}

// PaymentIntentCreateRequest creates a new payment intent
type PaymentIntentCreateRequest struct {
	PaymentIntentParams
	// PaymentLinkID is the optional ID of the payment link used to create this payment intent
	PaymentLinkID *string `json:"payment_link_id,omitempty"`
	// ProductVersion is the optional version of the product catalog to use
	ProductVersion *int64 `json:"product_version,omitempty"`
	// PrimaryVariant is the optional primary variant selection
	PrimaryVariant *VariantSelectionRequest `json:"primary_variant,omitempty"`
	// Addons is the optional list of addon variant selections
	Addons []VariantSelectionRequest `json:"addons,omitempty"`
	// ShippingAddress is the optional shipping address for the payment intent
	ShippingAddress *PaymentIntentShippingAddress `json:"shipping_address,omitempty"`
	// TaxRegion is the optional tax jurisdiction for the payment intent
	TaxRegion *PaymentIntentTaxRegion `json:"tax_region,omitempty"`
}

// PaymentIntentUpdateRequest confirms payment intent with payment method
type PaymentIntentUpdateRequest struct {
	// ID is the payment intent ID from URL path
	ID string `json:"-"`
	// PaymentMethodType is the type of payment method being used
	PaymentMethodType *PaymentMethodType `json:"payment_method_type" binding:"required"`
	// PaymentMethodData contains the payment method confirmation options
	PaymentMethodData *PaymentMethodConfirmOptions `json:"payment_method_options" binding:"required"`
	// SaveCustomer indicates whether to save the customer information for future use
	SaveCustomer *bool `json:"save_customer,omitempty"`
}

// PaymentIntentLinkUpdateRequest updates payment link payment intent
type PaymentIntentLinkUpdateRequest struct {
	// ID is the payment intent ID from URL path
	ID string `json:"-"`
	// Key is the merchant public API key
	Key string `json:"key"`
	// PaymentLinkId is the ID of the payment link
	PaymentLinkId string `json:"payment_link_id"`
	// ClientSecret is the client secret for the payment intent
	ClientSecret string `json:"client_secret"`
	// PaymentMethodType is the type of payment method being used
	PaymentMethodType *PaymentMethodType `json:"payment_method_type" binding:"required"`
	// PaymentMethodData contains the payment method confirmation options
	PaymentMethodData *PaymentMethodConfirmOptions `json:"payment_method_options" binding:"required"`
	// SaveCustomer indicates whether to save the customer information for future use
	SaveCustomer *bool `json:"save_customer,omitempty"`
}

// PaymentIntentLinkCreateRequest creates payment intent from payment link
type PaymentIntentLinkCreateRequest struct {
	// PaymentLinkID is the ID of the payment link to create the payment intent from
	PaymentLinkID string `json:"payment_link_id" binding:"required"`
	// ProductVersion is the optional version of the product catalog to use
	ProductVersion *int64 `json:"product_version,omitempty"`
	// PrimaryVariant is the primary variant selection
	PrimaryVariant VariantSelectionRequest `json:"primary_variant" binding:"required"`
	// Addons is the optional list of addon variant selections
	Addons []VariantSelectionRequest `json:"addons,omitempty"`
	// Customer is the optional customer ID
	Customer *string `json:"customer,omitempty"`
	// CustomerAuthToken is the optional customer authentication token
	CustomerAuthToken *string `json:"customer_auth_token,omitempty"`
	// MerchantOrderId is the optional merchant order ID for reference
	MerchantOrderId *string `json:"merchant_order_id,omitempty"`
	// Description is the optional description of the payment intent
	Description *string `json:"description,omitempty"`
	// ReceiptEmail is the optional email address to send the receipt to
	ReceiptEmail *string `json:"receipt_email,omitempty"`
	// ReturnUrl is the optional URL to redirect to after payment completion
	ReturnUrl *string `json:"return_url,omitempty"`
	// ShippingAddress is the optional shipping address for the payment intent
	ShippingAddress *PaymentIntentShippingAddress `json:"shipping_address,omitempty"`
	// TaxRegion is the optional tax jurisdiction for the payment intent
	TaxRegion *PaymentIntentTaxRegion `json:"tax_region,omitempty"`
}

// PaymentIntentCancelRequest cancels a payment intent
type PaymentIntentCancelRequest struct {
	// Reason is the reason for cancelling the payment intent
	Reason PaymentIntentCancellationReason `json:"reason"`
}

// PayerMaxDropinOrderRequest creates PayerMax drop-in order
type PayerMaxDropinOrderRequest struct {
	// PaymentLinkId is the ID of the payment link
	PaymentLinkId string `json:"payment_link_id"`
	// Key is the merchant public API key
	Key string `json:"key"`
	// PaymentIntentID is the ID of the payment intent
	PaymentIntentID string `json:"payment_intent_id"`
	// ChargeID is the ID of the charge to process
	ChargeID string `json:"charge_id" binding:"required"`
	// ClientSecret is the client secret for the payment intent
	ClientSecret string `json:"client_secret" binding:"required"`
	// PaymentToken is the PayerMax payment token
	PaymentToken string `json:"payment_token" binding:"required"`
}

// PaymentIntentInvoiceCreateRequest creates payment intent from invoice
type PaymentIntentInvoiceCreateRequest struct {
	// InvoiceID is the ID of the invoice to create the payment intent from
	InvoiceID string `json:"invoice_id" binding:"required"`
}

// PaymentIntentGetRequest gets a payment intent by ID
type PaymentIntentGetRequest struct {
	// ID is the payment intent ID from URL path
	ID string `json:"-"`
	// Key is the optional publishable key for public access
	Key *string `json:"key"`
	// PaymentLinkId is required if accessing via payment link
	PaymentLinkId *string `json:"payment_link_id"`
	// ClientSecret is required if using a publishable key
	ClientSecret string `json:"client_secret"`
}

// PaymentIntentListRequest lists payment intents with filters
type PaymentIntentListRequest struct {
	Pagination
	// Customer filters payment intents by customer ID
	Customer *string `json:"customer,omitempty" form:"customer"`
	// CustomerEmail filters payment intents by customer email
	CustomerEmail *string `json:"customer_email,omitempty" form:"customer_email"`
	// MerchantOrderId filters payment intents by merchant order ID
	MerchantOrderId *string `json:"merchant_order_id,omitempty" form:"merchant_order_id"`
	// PermanentDeposit filters payment intents by permanent deposit flag
	PermanentDeposit *bool `json:"permanent_deposit,omitempty" form:"permanent_deposit"`
	// PermanentDepositAssetId filters payment intents by permanent deposit asset ID
	PermanentDepositAssetId *string `json:"permanent_deposit_asset_id,omitempty" form:"permanent_deposit_asset_id"`
}

// ================================
// Response Types
// ================================

// PaymentIntentCreateResp represents payment intent creation response
type PaymentIntentCreateResp struct {
	PaymentIntent
}

// PaymentIntentUpdateResp represents payment intent update response
type PaymentIntentUpdateResp struct {
	PaymentIntent
}

// PaymentIntentGetResp represents payment intent get response
type PaymentIntentGetResp struct {
	PaymentIntent
}

// PaymentIntentListResp represents paginated payment intent list
type PaymentIntentListResp struct {
	// PaymentIntents is the list of payment intents
	PaymentIntents []*PaymentIntent `json:"payment_intents" form:"list"`
	// Total is the total number of payment intents matching the filters
	Total int64 `json:"total" form:"total"`
	// Page is the current page number
	Page int32 `json:"page" form:"page"`
	// PageSize is the number of items per page
	PageSize int32 `json:"page_size" form:"page_size"`
}

// PaymentIntentLinkCreateResp represents payment link creation response
type PaymentIntentLinkCreateResp struct {
	// PaymentIntent is the created payment intent
	PaymentIntent *PaymentIntent `json:"payment_intent"`
}

// PayerMaxDropinOrderResponse represents PayerMax drop-in order response
type PayerMaxDropinOrderResponse struct {
	// Success indicates whether the order creation was successful
	Success bool `json:"success"`
	// ChargeID is the ID of the charge
	ChargeID string `json:"charge_id"`
	// PaymentIntent is the payment intent associated with the order
	PaymentIntent *PaymentIntent `json:"payment_intent"`
}

// PaymentIntentInvoiceCreateResponse represents invoice payment intent creation response
type PaymentIntentInvoiceCreateResponse struct {
	// PaymentIntent is the created payment intent for the invoice
	PaymentIntent *PaymentIntent `json:"payment_intent"`
}

// ================================
// Validation Methods
// ================================

// Validate validates variant selection request
func (s *VariantSelectionRequest) Validate(kind string) error {
	// Note: String trimming would mutate the original request
	// Validation should be done without mutation in SDK types
	if len(strings.TrimSpace(s.VariantID)) == 0 {
		return fmt.Errorf("%s variant_id is required", kind)
	}
	if s.Quantity <= 0 {
		return fmt.Errorf("%s quantity must be greater than 0", kind)
	}
	return nil
}

// Validate validates shipping address
func (addr *PaymentIntentShippingAddress) Validate() error {
	if addr == nil {
		return errors.New("shipping address is required")
	}
	if len(strings.TrimSpace(addr.Country)) == 0 {
		return errors.New("shipping address country is required")
	}
	if len(strings.TrimSpace(addr.City)) == 0 {
		return errors.New("shipping address city is required")
	}
	if len(strings.TrimSpace(addr.PostalCode)) == 0 {
		return errors.New("shipping address postal_code is required")
	}
	if len(strings.TrimSpace(addr.Line1)) == 0 {
		return errors.New("shipping address line1 is required")
	}
	return nil
}

// Validate validates tax region
func (region *PaymentIntentTaxRegion) Validate() error {
	if region == nil {
		return errors.New("tax region is required")
	}
	if len(strings.TrimSpace(region.Country)) == 0 {
		return errors.New("tax region country is required")
	}
	return nil
}

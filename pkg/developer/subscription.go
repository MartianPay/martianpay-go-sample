package developer

// ================================
// Constants and Types
// ================================

// SubscriptionStatus represents the status of a subscription
type SubscriptionStatus string

const (
	// SubscriptionStatusIncomplete indicates the subscription is in initial state, awaiting first payment
	SubscriptionStatusIncomplete SubscriptionStatus = "incomplete"
	// SubscriptionStatusActive indicates the subscription is active and billing normally
	SubscriptionStatusActive SubscriptionStatus = "active"
	// SubscriptionStatusPaused indicates the subscription is temporarily paused
	SubscriptionStatusPaused SubscriptionStatus = "paused"
	// SubscriptionStatusPastDue indicates the latest invoice payment has failed
	SubscriptionStatusPastDue SubscriptionStatus = "past_due"
	// SubscriptionStatusCanceled indicates the subscription has been canceled
	SubscriptionStatusCanceled SubscriptionStatus = "canceled"
)

// CollectionMethod represents how payment is collected
type CollectionMethod string

const (
	// CollectionMethodChargeAutomatically indicates payment intents are automatically created
	CollectionMethodChargeAutomatically CollectionMethod = "charge_automatically"
	// CollectionMethodSendInvoice indicates invoices are sent to the customer
	CollectionMethodSendInvoice CollectionMethod = "send_invoice"
)

// PauseCollectionBehavior defines how to handle pending invoices when pausing a subscription
type PauseCollectionBehavior string

const (
	// PauseCollectionBehaviorVoid indicates pending invoices should be canceled or voided
	PauseCollectionBehaviorVoid PauseCollectionBehavior = "void"
	// PauseCollectionBehaviorKeepAsDraft indicates invoices should be kept as draft
	PauseCollectionBehaviorKeepAsDraft PauseCollectionBehavior = "keep_as_draft"
)

// ================================
// Request Types
// ================================

// ListCustomerSubscriptionsRequest lists subscriptions for authenticated customer
type ListCustomerSubscriptionsRequest struct {
	// Status filters subscriptions by status
	Status *string `json:"status" form:"status"`
	// Offset is the number of records to skip for pagination
	Offset int `json:"offset" form:"offset"`
	// Limit is the maximum number of records to return
	Limit int `json:"limit" form:"limit"`
}

// GetCustomerSubscriptionRequest gets specific subscription for authenticated customer
type GetCustomerSubscriptionRequest struct {
	// SubscriptionID is the subscription ID from URL path
	SubscriptionID string `json:"-" form:"-"`
}

// CancelCustomerSubscriptionRequest cancels subscription for authenticated customer
type CancelCustomerSubscriptionRequest struct {
	// SubscriptionID is the subscription ID from URL path
	SubscriptionID string `json:"-"`
	// CancelAtPeriodEnd indicates whether to cancel at period end (true) or immediately (false), default true
	CancelAtPeriodEnd *bool `json:"cancel_at_period_end"`
	// CancelReason is the reason for cancellation
	CancelReason *string `json:"cancel_reason,omitempty"`
}

// PauseCustomerSubscriptionRequest pauses subscription for authenticated customer
type PauseCustomerSubscriptionRequest struct {
	// SubscriptionID is the subscription ID from URL path
	SubscriptionID string `json:"-"`
	// Behavior defines how to handle pending invoices ("void" | "keep_as_draft"), defaults to "void"
	Behavior *PauseCollectionBehavior `json:"behavior"`
	// ResumesAt is the Unix timestamp for automatic resumption
	ResumesAt *int64 `json:"resumes_at,omitempty"`
}

// ResumeCustomerSubscriptionRequest resumes subscription for authenticated customer
type ResumeCustomerSubscriptionRequest struct {
	// SubscriptionID is the subscription ID from URL path
	SubscriptionID string `json:"-"`
}

// UpdateCustomerSubscriptionPaymentMethodRequest updates payment method for customer subscription
type UpdateCustomerSubscriptionPaymentMethodRequest struct {
	// SubscriptionID is the ID of the subscription (from URL path)
	SubscriptionID string `json:"-"`
	// DefaultPaymentMethodID is the ID of the payment method to set as default
	DefaultPaymentMethodID string `json:"default_payment_method_id" binding:"required"`
}

// GetMerchantSubscriptionRequest gets specific subscription for merchant management
type GetMerchantSubscriptionRequest struct {
	// SubscriptionID is the subscription ID from URL path
	SubscriptionID string `json:"-"`
}

// ListMerchantSubscriptionsRequest lists subscriptions for merchant with filters
type ListMerchantSubscriptionsRequest struct {
	// CustomerID filters subscriptions by customer ID
	CustomerID *string `json:"customer_id" form:"customer_id"`
	// Status filters subscriptions by status
	Status *string `json:"status" form:"status"`
	// ExternalID filters subscriptions by external ID
	ExternalID *string `json:"external_id" form:"external_id"`
	// Offset is the number of records to skip for pagination
	Offset int `json:"offset" form:"offset"`
	// Limit is the maximum number of records to return
	Limit int `json:"limit" form:"limit"`
}

// CancelMerchantSubscriptionRequest cancels subscription for merchant management
type CancelMerchantSubscriptionRequest struct {
	// SubscriptionID is the subscription ID from URL path
	SubscriptionID string `json:"-"`
	// CancelAtPeriodEnd indicates whether to cancel at period end (true) or immediately (false), default true
	CancelAtPeriodEnd *bool `json:"cancel_at_period_end"`
	// CancelReason is the reason for cancellation
	CancelReason *string `json:"cancel_reason,omitempty"`
}

// PauseMerchantSubscriptionRequest pauses subscription for merchant management
type PauseMerchantSubscriptionRequest struct {
	// SubscriptionID is the subscription ID from URL path
	SubscriptionID string `json:"-"`
	// Behavior defines how to handle pending invoices ("void" | "keep_as_draft"), defaults to "void"
	Behavior *PauseCollectionBehavior `json:"behavior"`
	// ResumesAt is the Unix timestamp for automatic resumption
	ResumesAt *int64 `json:"resumes_at,omitempty"`
}

// ResumeMerchantSubscriptionRequest resumes subscription for merchant management
type ResumeMerchantSubscriptionRequest struct {
	// SubscriptionID is the subscription ID from URL path
	SubscriptionID string `json:"-"`
}

// ================================
// Response Types
// ================================

// SubscriptionCurrentPricingTier represents the pricing information for a specific billing cycle
type SubscriptionCurrentPricingTier struct {
	// CycleNumber is which cycle this pricing applies to
	CycleNumber int `json:"cycle_number"`
	// BasePrice is the variant base price
	BasePrice string `json:"base_price"`
	// SellingPlanDiscount is the discount amount from selling plan
	SellingPlanDiscount string `json:"selling_plan_discount"`
	// FinalPrice is the final price after discount
	FinalPrice string `json:"final_price"`
	// Currency is the currency code
	Currency string `json:"currency"`
	// PolicyType is the discount policy type ("FIXED" | "RECURRING")
	PolicyType *string `json:"policy_type,omitempty"`
	// CycleDescription is the cycle description (e.g., "Cycle 1" or "Cycle 2 onwards")
	CycleDescription *string `json:"cycle_description,omitempty"`
	// DiscountPercentage is the discount percentage if discount is percentage-based
	DiscountPercentage *string `json:"discount_percentage,omitempty"`
	// BillingCycle is the billing cycle period ("month" | "year" | "3 months")
	BillingCycle string `json:"billing_cycle"`
	// BillingCycleInterval is the interval value (e.g., 1 for monthly, 3 for quarterly)
	BillingCycleInterval int `json:"billing_cycle_interval,omitempty"`
}

// ListSubscriptionsResponse represents a paginated list of subscriptions
type ListSubscriptionsResponse struct {
	// Data is the list of subscription details
	Data []*SubscriptionDetails `json:"data"`
	// Total is the total number of subscriptions matching the filter
	Total int64 `json:"total"`
	// Offset is the number of records skipped
	Offset int `json:"offset"`
	// Limit is the maximum number of records returned
	Limit int `json:"limit"`
}

// SubscriptionDetails represents a subscription in API responses
type SubscriptionDetails struct {
	// ID is the unique identifier for the subscription
	ID string `json:"id"`
	// MerchantID is the ID of the merchant who owns this subscription
	MerchantID string `json:"merchant_id"`
	// CustomerID is the ID of the customer this subscription belongs to
	CustomerID string `json:"customer_id"`
	// SellingPlanID is the ID of the selling plan (pricing schedule)
	SellingPlanID string `json:"selling_plan_id"`
	// ProductID is the ID of the subscribed product
	ProductID *string `json:"product_id,omitempty"`
	// VariantID is the ID of the subscribed product variant
	VariantID *string `json:"variant_id,omitempty"`
	// Status is the current subscription status (incomplete, active, paused, past_due, canceled)
	Status string `json:"status"`
	// CollectionMethod defines how payment is collected (charge_automatically or send_invoice)
	CollectionMethod string `json:"collection_method"`
	// BillingCycleAnchor is the Unix timestamp that defines the billing cycle start date
	BillingCycleAnchor int64 `json:"billing_cycle_anchor"`
	// CurrentPeriodStart is the Unix timestamp when the current billing period started
	CurrentPeriodStart int64 `json:"current_period_start"`
	// CurrentPeriodEnd is the Unix timestamp when the current billing period ends
	CurrentPeriodEnd int64 `json:"current_period_end"`
	// TrialStart is the Unix timestamp when the trial period started
	TrialStart *int64 `json:"trial_start,omitempty"`
	// TrialEnd is the Unix timestamp when the trial period ends
	TrialEnd *int64 `json:"trial_end,omitempty"`
	// CanceledAt is the Unix timestamp when the subscription was canceled
	CanceledAt *int64 `json:"canceled_at,omitempty"`
	// CancelAtPeriodEnd indicates whether the subscription will be canceled at the end of the current period
	CancelAtPeriodEnd bool `json:"cancel_at_period_end"`
	// CancelReason is the customer-provided reason for cancellation
	CancelReason *string `json:"cancel_reason,omitempty"`
	// PauseCollectionBehavior defines how to handle pending invoices when paused (void or keep_as_draft)
	PauseCollectionBehavior *string `json:"pause_collection_behavior,omitempty"`
	// PausedAt is the Unix timestamp when the subscription was paused
	PausedAt *int64 `json:"paused_at,omitempty"`
	// ResumesAt is the Unix timestamp when the subscription will automatically resume
	ResumesAt *int64 `json:"resumes_at,omitempty"`
	// LatestInvoiceID is the ID of the most recent invoice for this subscription
	LatestInvoiceID *string `json:"latest_invoice_id,omitempty"`
	// DefaultPaymentMethodID is the ID of the default payment method
	DefaultPaymentMethodID *string `json:"default_payment_method_id,omitempty"`
	// DefaultProviderType is the payment provider type (e.g., "stripe", "crypto")
	DefaultProviderType *string `json:"default_provider_type,omitempty"`
	// DefaultPaymentMethodType is the payment method type (e.g., "card", "crypto")
	DefaultPaymentMethodType *string `json:"default_payment_method_type,omitempty"`
	// Metadata is a set of key-value pairs for storing additional information
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// ExternalID is the merchant's external identifier for this subscription
	ExternalID *string `json:"external_id,omitempty"`
	// CreatedAt is the Unix timestamp when the subscription was created
	CreatedAt int64 `json:"created_at"`
	// UpdatedAt is the Unix timestamp when the subscription was last updated
	UpdatedAt int64 `json:"updated_at"`
	// PaymentRequired indicates whether payment is required for an incomplete subscription
	PaymentRequired *bool `json:"payment_required,omitempty"`
	// PaymentURL is the URL where the customer can complete the initial payment for an incomplete subscription
	PaymentURL *string `json:"payment_url,omitempty"`
	// PaymentExpiresAt is the Unix timestamp when the payment URL expires for an incomplete subscription
	PaymentExpiresAt *int64 `json:"payment_expires_at,omitempty"`
	// HoursSinceCreation is the number of hours since the subscription was created (for incomplete subscriptions)
	HoursSinceCreation *float64 `json:"hours_since_creation,omitempty"`
	// CurrentCycleNumber is which billing cycle this subscription is currently in
	CurrentCycleNumber *int `json:"current_cycle_number,omitempty"`
	// CurrentPricingTier is the applicable pricing tier for the current billing cycle
	CurrentPricingTier *SubscriptionCurrentPricingTier `json:"current_pricing_tier,omitempty"`
	// UpcomingPricingTier is the pricing tier for the next billing cycle (if different from current)
	UpcomingPricingTier *SubscriptionCurrentPricingTier `json:"upcoming_pricing_tier,omitempty"`
	// NextChargeAmount is the amount that will be charged in the next billing cycle
	NextChargeAmount *string `json:"next_charge_amount,omitempty"`
	// NextChargeAmountDisplay is the human-readable format of the next charge amount
	NextChargeAmountDisplay *string `json:"next_charge_amount_display,omitempty"`
	// SellingPlanPricing contains the complete pricing tiers for all cycles (same format as product API)
	SellingPlanPricing *SellingPlanPricing `json:"selling_plan_pricing,omitempty"`
	// MerchantName is the merchant's display name (fetched via JOIN/lookup, not stored)
	MerchantName *string `json:"merchant_name,omitempty"`
	// CustomerEmail is the customer's email address for display (fetched via JOIN/lookup, not stored)
	CustomerEmail *string `json:"customer_email,omitempty"`
	// CustomerName is the customer's display name (fetched via JOIN/lookup, not stored)
	CustomerName *string `json:"customer_name,omitempty"`
	// ProductName is the product's display name (fetched via JOIN/lookup, not stored)
	ProductName *string `json:"product_name,omitempty"`
	// ProductDescription is the product's description for display (fetched via JOIN/lookup, not stored)
	ProductDescription *string `json:"product_description,omitempty"`
	// ProductImageURL is the first product image URL for display (fetched via JOIN/lookup, not stored)
	ProductImageURL *string `json:"product_image_url,omitempty"`
	// VariantTitle is the variant's title for display (e.g., "Large / Blue") (fetched via JOIN/lookup, not stored)
	VariantTitle *string `json:"variant_title,omitempty"`
	// VariantOptionValues is the variant's option values map (e.g., {"size": "Large"}) (fetched via JOIN/lookup, not stored)
	VariantOptionValues map[string]string `json:"variant_option_values,omitempty"`
	// VariantPrice is the variant's original price for comparison (fetched via JOIN/lookup, not stored)
	VariantPrice *string `json:"variant_price,omitempty"`
	// SellingPlanName is the selling plan's display name (fetched via JOIN/lookup, not stored)
	SellingPlanName *string `json:"selling_plan_name,omitempty"`
	// SellingPlanDescription is the selling plan's description for display (fetched via JOIN/lookup, not stored)
	SellingPlanDescription *string `json:"selling_plan_description,omitempty"`
	// PaymentMethodBrand is the payment method's brand (e.g., "visa") (fetched via JOIN/lookup, not stored)
	PaymentMethodBrand *string `json:"payment_method_brand,omitempty"`
	// PaymentMethodLast4 is the payment method's last 4 digits for display (fetched via JOIN/lookup, not stored)
	PaymentMethodLast4 *string `json:"payment_method_last4,omitempty"`
	// PendingUpdate contains scheduled plan changes (e.g., downgrade at period end)
	PendingUpdate *SubscriptionPendingUpdate `json:"pending_update,omitempty"`
	// ProrationBehavior indicates how proration was handled for the last plan change
	ProrationBehavior *string `json:"proration_behavior,omitempty"`
	// ProrationDate is the timestamp used for proration calculation
	ProrationDate *int64 `json:"proration_date,omitempty"`

	// === Proration fields (populated for update/preview operations) ===
	// Applied indicates if the change was actually applied (true=update, false=preview, nil=GET)
	Applied *bool `json:"applied,omitempty"`
	// IsUpgrade indicates if this is an upgrade (true) or downgrade (false)
	IsUpgrade *bool `json:"is_upgrade,omitempty"`
	// EffectiveDate is the Unix timestamp when the change takes/would take effect
	EffectiveDate *int64 `json:"effective_date,omitempty"`
	// ProrationCredit is the credit amount for unused time on old plan
	ProrationCredit *string `json:"proration_credit,omitempty"`
	// ChargeToday is the net amount charged/to be charged today
	ChargeToday *string `json:"charge_today,omitempty"`
	// NextChargeDate is the Unix timestamp of the next billing date
	NextChargeDate *int64 `json:"next_charge_date,omitempty"`
	// ProrationDetails contains detailed proration breakdown (for update/preview)
	ProrationDetails *ProrationDetails `json:"proration_details,omitempty"`
}

// SubscriptionPendingUpdate represents a scheduled subscription change
type SubscriptionPendingUpdate struct {
	// TargetSellingPlanID is the selling plan to change to
	TargetSellingPlanID string `json:"target_selling_plan_id"`
	// TargetSellingPlanName is the name of the target selling plan
	TargetSellingPlanName *string `json:"target_selling_plan_name,omitempty"`
	// TargetVariantID is the variant to change to (optional)
	TargetVariantID *string `json:"target_variant_id,omitempty"`
	// TargetVariantTitle is the display title of the target variant
	TargetVariantTitle *string `json:"target_variant_title,omitempty"`
	// TargetVariantPrice is the base price of the target variant (formatted)
	TargetVariantPrice *string `json:"target_variant_price,omitempty"`
	// TargetVariantOptionValues contains the option values of the target variant
	TargetVariantOptionValues map[string]string `json:"target_variant_option_values,omitempty"`
	// NextChargeAmount is the calculated amount that will be charged after the change takes effect
	NextChargeAmount *string `json:"next_charge_amount,omitempty"`
	// ChangeType indicates whether this is an "upgrade" or "downgrade"
	ChangeType *string `json:"change_type,omitempty"`
	// EffectiveDate is the Unix timestamp when the change will take effect
	EffectiveDate int64 `json:"effective_date"`
	// ProrationBehavior indicates how proration will be handled
	ProrationBehavior string `json:"proration_behavior"`
	// BillingCycleAnchor controls billing cycle timing ("now", "unchanged") (optional)
	BillingCycleAnchor *string `json:"billing_cycle_anchor,omitempty"`
	// ProrationDate is the Unix timestamp used for proration calculation (optional)
	ProrationDate *int64 `json:"proration_date,omitempty"`
	// Metadata contains additional information for this pending update
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// ScheduledAt is the Unix timestamp when this update was scheduled
	ScheduledAt int64 `json:"scheduled_at"`
}

// ================================
// Subscription Price Change APIs
// ================================

// ProrationBehavior constants
const (
	// ProrationBehaviorAlwaysInvoice creates invoice immediately with proration
	ProrationBehaviorAlwaysInvoice = "always_invoice"
	// ProrationBehaviorCreateProrations creates proration items added to next invoice
	ProrationBehaviorCreateProrations = "create_prorations"
	// ProrationBehaviorNone does not create any proration
	ProrationBehaviorNone = "none"
)

// BillingCycleAnchor constants
const (
	// BillingCycleAnchorNow resets billing cycle to current time
	BillingCycleAnchorNow = "now"
	// BillingCycleAnchorUnchanged keeps existing billing cycle date
	BillingCycleAnchorUnchanged = "unchanged"
)

// UpdateSubscriptionPlanRequest updates or previews a subscription plan change (upgrade/downgrade)
// This request type is used by both update and preview endpoints
type UpdateSubscriptionPlanRequest struct {
	// SubscriptionID is the subscription ID from URL path
	SubscriptionID string `json:"-"`
	// PrimaryVariant is the primary variant selection to update
	PrimaryVariant *SubscriptionItemUpdate `json:"primary_variant,omitempty"`
	// Addons is the optional list of addon variant selections (reserved for future use)
	Addons []SubscriptionItemUpdate `json:"addons,omitempty"`
	// ProrationBehavior controls how proration is handled ("always_invoice", "create_prorations", "none")
	ProrationBehavior *string `json:"proration_behavior,omitempty"`
	// BillingCycleAnchor controls billing cycle timing ("now", "unchanged")
	BillingCycleAnchor *string `json:"billing_cycle_anchor,omitempty"`
	// ProrationDate is the Unix timestamp for custom proration calculation (backdating)
	ProrationDate *int64 `json:"proration_date,omitempty"`
	// Metadata is a set of key-value pairs for storing additional information (ignored in preview)
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// SubscriptionItemUpdate represents a subscription item update
type SubscriptionItemUpdate struct {
	// ID is the existing subscription item ID (optional for updates)
	ID *string `json:"id,omitempty"`
	// SellingPlanID is the new selling plan ID to switch to
	SellingPlanID string `json:"selling_plan_id" binding:"required"`
	// VariantID is the new variant ID (optional)
	VariantID *string `json:"variant_id,omitempty"`
	// Quantity is the quantity (default: 1)
	Quantity *int `json:"quantity,omitempty"`
}

// PreviewSubscriptionUpdateRequest is an alias for UpdateSubscriptionPlanRequest
// Deprecated: Use UpdateSubscriptionPlanRequest instead
type PreviewSubscriptionUpdateRequest = UpdateSubscriptionPlanRequest

// SubscriptionUpdateResponse is deprecated - use SubscriptionDetails directly
// The proration fields are now included in SubscriptionDetails
// Deprecated: Use SubscriptionDetails instead
type SubscriptionUpdateResponse = SubscriptionDetails

// SubscriptionPreviewResponse is deprecated - use SubscriptionDetails directly
// The proration fields are now included in SubscriptionDetails
// Deprecated: Use SubscriptionDetails instead
type SubscriptionPreviewResponse = SubscriptionDetails

// ProrationDetails contains detailed proration calculation breakdown
type ProrationDetails struct {
	// CreditedAmount is the credit for unused time on old plan (in cents)
	CreditedAmount string `json:"credited_amount"`
	// ChargedAmount is the charge for new plan (in cents)
	ChargedAmount string `json:"charged_amount"`
	// NetAmount is the final net amount (charged - credited) (in cents)
	NetAmount string `json:"net_amount"`
	// DaysRemaining is the number of days remaining in current period
	DaysRemaining int `json:"days_remaining"`
	// TotalDays is the total days in the current period
	TotalDays int `json:"total_days"`
	// CurrentPrice is the current plan price (in cents)
	CurrentPrice string `json:"current_price"`
	// TargetPrice is the target plan price (in cents)
	TargetPrice string `json:"target_price"`
}

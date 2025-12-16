// selling_plan.go contains types for managing subscription selling plans.
// It provides structures for defining recurring billing schedules, pricing tiers,
// and subscription plan configurations.
package developer

// ================================
// Constants and Enums
// ================================

// SellingPlanStatus represents subscription plan status
type SellingPlanStatus string

const (
	SellingPlanStatusActive   SellingPlanStatus = "active"
	SellingPlanStatusDisabled SellingPlanStatus = "disabled"
)

// BillingPolicyType represents billing policy type
type BillingPolicyType string

const (
	// BillingPolicyTypeRecurring indicates charge per billing cycle
	BillingPolicyTypeRecurring BillingPolicyType = "RECURRING"
	// BillingPolicyTypePrepaid indicates prepaid billing
	BillingPolicyTypePrepaid BillingPolicyType = "PREPAID"
)

// PolicyType represents pricing policy type
type PolicyType string

const (
	// PolicyTypeFixed indicates fixed policy that applies to all cycles unless overridden by RECURRING
	PolicyTypeFixed PolicyType = "FIXED"
	// PolicyTypeRecurring indicates subsequent cycle override policy
	PolicyTypeRecurring PolicyType = "RECURRING"
)

// AdjustmentType represents discount type
type AdjustmentType string

const (
	// AdjustmentTypePercentage indicates percentage-based discount
	AdjustmentTypePercentage AdjustmentType = "PERCENTAGE"
	// AdjustmentTypeFixedAmount indicates fixed amount discount
	AdjustmentTypeFixedAmount AdjustmentType = "FIXED_AMOUNT"
	// AdjustmentTypePrice indicates direct pricing
	AdjustmentTypePrice AdjustmentType = "PRICE"
)

// BillingInterval represents billing cycle unit
type BillingInterval string

const (
	// BillingIntervalDay indicates daily billing cycle
	BillingIntervalDay BillingInterval = "day"
	// BillingIntervalWeek indicates weekly billing cycle
	BillingIntervalWeek BillingInterval = "week"
	// BillingIntervalMonth indicates monthly billing cycle
	BillingIntervalMonth BillingInterval = "month"
	// BillingIntervalYear indicates yearly billing cycle
	BillingIntervalYear BillingInterval = "year"
)

// ================================
// Struct Types
// ================================

// BillingPolicyConfig represents billing cycle configuration
type BillingPolicyConfig struct {
	// Interval is the billing interval ("day" | "week" | "month" | "year")
	Interval string `json:"interval"`
	// IntervalCount is the cycle multiplier as string (e.g., "1"=monthly, "3"=quarterly)
	IntervalCount string `json:"interval_count"`
	// MinCycles is the minimum cycles for PREPAID (as string)
	MinCycles *string `json:"min_cycles,omitempty"`
}

// PricingPolicyItem represents a single pricing policy item
type PricingPolicyItem struct {
	// PolicyType is the policy type ("FIXED" | "RECURRING", FIXED=first cycle, RECURRING=subsequent cycles)
	PolicyType string `json:"policy_type"`
	// AdjustmentType is the adjustment type ("PERCENTAGE" | "FIXED_AMOUNT" | "PRICE")
	AdjustmentType string `json:"adjustment_type"`
	// AdjustmentValue is the discount value as string (e.g., "10.5")
	AdjustmentValue string `json:"adjustment_value"`
	// AfterCycle is only for RECURRING type (as string)
	AfterCycle *string `json:"after_cycle,omitempty"`
}

// PricingPolicyConfig is an array of pricing policy items
type PricingPolicyConfig []PricingPolicyItem

// SellingPlanGroupWithPlans represents a Selling Plan Group with associated Plans
type SellingPlanGroupWithPlans struct {
	// ID is the unique identifier for the selling plan group
	ID string `json:"id"`
	// Name is the name of the selling plan group
	Name string `json:"name"`
	// Description is the description of the selling plan group
	Description string `json:"description,omitempty"`
	// Options is the list of option names
	Options []string `json:"options,omitempty"`
	// Status is the current status of the selling plan group
	Status string `json:"status"`
	// SellingPlans is the list of associated selling plans
	SellingPlans []*SellingPlan `json:"selling_plans"`
}

// SellingPlan represents a subscription plan response structure
type SellingPlan struct {
	// ID is the unique identifier for the selling plan
	ID string `json:"id"`
	// Name is the name of the selling plan
	Name string `json:"name"`
	// Description is the description of the selling plan
	Description string `json:"description,omitempty"`
	// BillingPolicyType is the billing policy type ("RECURRING" | "PREPAID")
	BillingPolicyType string `json:"billing_policy_type"`
	// BillingPolicy contains the billing cycle configuration
	BillingPolicy BillingPolicyConfig `json:"billing_policy"`
	// PricingPolicy is the list of pricing policy items
	PricingPolicy []PricingPolicyItem `json:"pricing_policy"`
	// TrialPeriodDays is the trial period in days as string
	TrialPeriodDays string `json:"trial_period_days"`
	// ValidFrom is the Unix timestamp when the plan becomes valid
	ValidFrom *int64 `json:"valid_from,omitempty"`
	// ValidUntil is the Unix timestamp when the plan expires
	ValidUntil *int64 `json:"valid_until,omitempty"`
	// Priority is the priority of the selling plan
	Priority string `json:"priority"`
	// Status is the current status of the selling plan
	Status string `json:"status"`
	// CreatedAt is the Unix timestamp when the plan was created
	CreatedAt int64 `json:"created_at"`
	// UpdatedAt is the Unix timestamp when the plan was last updated
	UpdatedAt int64 `json:"updated_at"`
	// Version is the version number for optimistic locking
	Version int64 `json:"version"`
}

// SellingPlanPricing represents pre-calculated subscription pricing (for Variant)
type SellingPlanPricing struct {
	// SellingPlanID is the ID of the selling plan
	SellingPlanID string `json:"selling_plan_id"`
	// SellingPlanName is the name of the selling plan
	SellingPlanName string `json:"selling_plan_name"`
	// BillingCycle is the billing cycle ("month" | "year" | "3 months")
	BillingCycle string `json:"billing_cycle"`
	// Currency is the currency code
	Currency string `json:"currency"`
	// TrialPeriodDays is the trial period in days
	TrialPeriodDays int `json:"trial_period_days"`
	// PricingTiers contains pricing tiers (includes FIXED and RECURRING policies)
	PricingTiers []PricingTier `json:"pricing_tiers"`
}

// PricingTier represents a single pricing tier (corresponding to a pricing policy)
type PricingTier struct {
	// BasePrice is the variant base price
	BasePrice string `json:"base_price"`
	// TotalCycles is the total number of cycles
	TotalCycles int `json:"total_cycles"`
	// SubtotalBeforePolicy is the subtotal before discount (base_price × cycles)
	SubtotalBeforePolicy string `json:"subtotal_before_policy"`
	// SellingPlanDiscount is the selling plan discount amount
	SellingPlanDiscount string `json:"selling_plan_discount"`
	// SubtotalAfterPolicy is the subtotal after discount (final subscription price)
	SubtotalAfterPolicy string `json:"subtotal_after_policy"`
	// PolicyType indicates which pricing policy applies ("FIXED" | "RECURRING")
	PolicyType *string `json:"policy_type,omitempty"`
	// AfterCycle is the cycle number this pricing starts from (for RECURRING policies)
	AfterCycle *int `json:"after_cycle,omitempty"`
	// CycleDescription is the cycle description (e.g., "Cycle 1" or "Cycle 2 onwards")
	CycleDescription *string `json:"cycle_description,omitempty"`
}

// ================================
// Request Types
// ================================

// CreateSellingPlanGroupRequest represents a request to create a new selling plan group
type CreateSellingPlanGroupRequest struct {
	// Name of the selling plan group
	Name string `json:"name" binding:"required"`
	// Description of the selling plan group
	Description string `json:"description"`
	// Options for the selling plan group (e.g., ["Size", "Color"])
	Options []string `json:"options"`
}

// UpdateSellingPlanGroupRequest represents a request to update an existing selling plan group
type UpdateSellingPlanGroupRequest struct {
	// Name of the selling plan group
	Name string `json:"name"`
	// Description of the selling plan group
	Description string `json:"description"`
	// Options for the selling plan group
	Options []string `json:"options"`
	// Status of the selling plan group: "active" | "disabled"
	Status string `json:"status"`
}

// BillingPolicyRequest represents billing policy configuration in a request
type BillingPolicyRequest struct {
	// Interval for billing: "day" | "week" | "month" | "year"
	Interval string `json:"interval" binding:"required"`
	// Interval count (string format, e.g., "1", "3", "12")
	IntervalCount string `json:"interval_count" binding:"required"`
	// Suggested minimum cycles (optional, string format)
	MinCycles *string `json:"min_cycles,omitempty"`
}

// PricingPolicyItemRequest represents a single pricing policy item in a request
type PricingPolicyItemRequest struct {
	// Policy type: "FIXED" (first cycle) | "RECURRING" (subsequent cycles)
	PolicyType string `json:"policy_type" binding:"required"`
	// Adjustment type: "PERCENTAGE" | "FIXED_AMOUNT" | "PRICE"
	AdjustmentType string `json:"adjustment_type" binding:"required"`
	// Adjustment value (string format, e.g., "10.5" means 10.5% or $10.5)
	AdjustmentValue string `json:"adjustment_value" binding:"required"`
	// Required for RECURRING only: take effect after which cycle (string format, "1" means starting from 2nd cycle)
	AfterCycle *string `json:"after_cycle,omitempty"`
}

// PricingPolicyRequest is an array of pricing policy items supporting multi-tier pricing
type PricingPolicyRequest []PricingPolicyItemRequest

// CreateSellingPlanRequest represents a request to create a new selling plan
type CreateSellingPlanRequest struct {
	// ID of the selling plan group this plan belongs to
	SellingPlanGroupID string `json:"selling_plan_group_id" binding:"required"`
	// Name of the selling plan
	Name string `json:"name" binding:"required"`
	// Description of the selling plan
	Description string `json:"description"`
	// Billing policy type: "RECURRING" | "PREPAID"
	BillingPolicyType string `json:"billing_policy_type" binding:"required"`
	// Billing policy configuration
	BillingPolicy BillingPolicyRequest `json:"billing_policy" binding:"required"`
	// Pricing policy (optional, following Shopify practice)
	PricingPolicy *PricingPolicyRequest `json:"pricing_policy,omitempty"`
	// Trial period in days (string format)
	TrialPeriodDays string `json:"trial_period_days"`
	// Valid from timestamp (Unix timestamp in seconds)
	ValidFrom *int64 `json:"valid_from,omitempty"`
	// Valid until timestamp (Unix timestamp in seconds)
	ValidUntil *int64 `json:"valid_until,omitempty"`
	// Priority of this plan (string format)
	Priority string `json:"priority"`
}

// UpdateSellingPlanRequest represents a request to update an existing selling plan
type UpdateSellingPlanRequest struct {
	// Name of the selling plan
	Name string `json:"name"`
	// Description of the selling plan
	Description string `json:"description"`
	// Billing policy type: "RECURRING" | "PREPAID"
	BillingPolicyType string `json:"billing_policy_type"`
	// Billing policy configuration
	BillingPolicy *BillingPolicyRequest `json:"billing_policy,omitempty"`
	// Pricing policy configuration
	PricingPolicy *PricingPolicyRequest `json:"pricing_policy,omitempty"`
	// Trial period in days (string format)
	TrialPeriodDays string `json:"trial_period_days"`
	// Valid from timestamp (Unix timestamp in seconds)
	ValidFrom *int64 `json:"valid_from,omitempty"`
	// Valid until timestamp (Unix timestamp in seconds)
	ValidUntil *int64 `json:"valid_until,omitempty"`
	// Priority of this plan (string format)
	Priority string `json:"priority"`
	// Status: "active" | "disabled"
	Status string `json:"status"`
}

// CalculatePriceRequest represents a request to calculate subscription price for a variant
type CalculatePriceRequest struct {
	// ID of the product variant
	VariantID string `json:"variant_id"`
	// Base price of the variant
	BasePrice string `json:"base_price"`
	// Currency code (currently only supports "USD")
	Currency string `json:"currency"`
	// ID of the selling plan to calculate price for
	PlanID string `json:"plan_id" binding:"required"`
}

// ================================
// Response Types
// ================================

// BillingPolicyResponse represents billing policy configuration in a response
type BillingPolicyResponse struct {
	// Billing interval: "day" | "week" | "month" | "year"
	Interval string `json:"interval"`
	// Interval count (string format)
	IntervalCount string `json:"interval_count"`
	// Suggested minimum cycles (string format)
	MinCycles *string `json:"min_cycles,omitempty"`
}

// PricingPolicyItemResponse represents a single pricing policy item in a response
type PricingPolicyItemResponse struct {
	// Policy type: "FIXED" | "RECURRING"
	PolicyType string `json:"policy_type"`
	// Adjustment type: "PERCENTAGE" | "FIXED_AMOUNT" | "PRICE"
	AdjustmentType string `json:"adjustment_type"`
	// Adjustment value (string format)
	AdjustmentValue string `json:"adjustment_value"`
	// Present only for RECURRING policies (string format)
	AfterCycle *string `json:"after_cycle,omitempty"`
}

// PricingPolicyResponse is an array of pricing policy items in a response
type PricingPolicyResponse []PricingPolicyItemResponse

// SellingPlanResponse represents a selling plan in a response
type SellingPlanResponse struct {
	// Unique identifier for the selling plan
	ID string `json:"id"`
	// ID of the selling plan group this plan belongs to
	SellingPlanGroupID string `json:"selling_plan_group_id"`
	// Merchant ID who owns this plan
	MerchantID string `json:"merchant_id"`
	// Name of the selling plan
	Name string `json:"name"`
	// Description of the selling plan
	Description string `json:"description,omitempty"`
	// Billing policy type: "RECURRING" | "PREPAID"
	BillingPolicyType string `json:"billing_policy_type"`
	// Billing policy configuration
	BillingPolicy BillingPolicyResponse `json:"billing_policy"`
	// Pricing policy configuration
	PricingPolicy PricingPolicyResponse `json:"pricing_policy"`
	// Trial period in days (string format)
	TrialPeriodDays string `json:"trial_period_days"`
	// Valid from timestamp (Unix timestamp in seconds)
	ValidFrom *int64 `json:"valid_from,omitempty"`
	// Valid until timestamp (Unix timestamp in seconds)
	ValidUntil *int64 `json:"valid_until,omitempty"`
	// Priority (string format)
	Priority string `json:"priority"`
	// Status: "active" | "disabled"
	Status string `json:"status"`
	// Unix timestamp when created (seconds)
	CreatedAt int64 `json:"created_at"`
	// Unix timestamp when last updated (seconds)
	UpdatedAt int64 `json:"updated_at"`
}

// SellingPlanGroupResponse represents a selling plan group in a response
type SellingPlanGroupResponse struct {
	// Unique identifier for the selling plan group
	ID string `json:"id"`
	// Merchant ID who owns this group
	MerchantID string `json:"merchant_id"`
	// Name of the selling plan group
	Name string `json:"name"`
	// Description of the selling plan group
	Description string `json:"description,omitempty"`
	// Options for the selling plan group
	Options []string `json:"options,omitempty"`
	// Status: "active" | "disabled"
	Status string `json:"status"`
	// Unix timestamp when created (seconds)
	CreatedAt int64 `json:"created_at"`
	// Unix timestamp when last updated (seconds)
	UpdatedAt int64 `json:"updated_at"`
	// Include plans if requested
	SellingPlans []*SellingPlanResponse `json:"selling_plans,omitempty"`
}

// ListSellingPlanGroupsResponse represents a paginated list of selling plan groups
type ListSellingPlanGroupsResponse struct {
	// List of selling plan groups
	Data []*SellingPlanGroupResponse `json:"data"`
	// Total number of records
	Total int64 `json:"total"`
	// Offset for pagination
	Offset int `json:"offset"`
	// Limit for pagination
	Limit int `json:"limit"`
}

// ListSellingPlansResponse represents a paginated list of selling plans
type ListSellingPlansResponse struct {
	// List of selling plans
	Data []*SellingPlanResponse `json:"data"`
	// Total number of records
	Total int64 `json:"total"`
	// Offset for pagination
	Offset int `json:"offset"`
	// Limit for pagination
	Limit int `json:"limit"`
}

// CalculatePriceResponse represents the calculated subscription price for a variant
type CalculatePriceResponse struct {
	// Base price of the variant
	BasePrice string `json:"base_price"`
	// Total number of billing cycles
	TotalCycles int `json:"total_cycles"`
	// Subtotal before applying pricing policy
	SubtotalBeforePolicy string `json:"subtotal_before_policy"`
	// Discount amount from selling plan
	SellingPlanDiscount string `json:"selling_plan_discount"`
	// Subtotal after applying pricing policy
	SubtotalAfterPolicy string `json:"subtotal_after_policy"`
	// Currency code
	Currency string `json:"currency"`
	// Billing cycle description
	BillingCycle string `json:"billing_cycle"`
	// Trial period in days
	TrialPeriodDays int `json:"trial_period_days"`
}

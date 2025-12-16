// product.go contains types for managing products and product variants.
// It provides structures for product catalogs, variants, options, pricing,
// and inventory management for e-commerce operations.
package developer

// ================================
// Core Product Types
// ================================

// ProductOptionSwatch represents visual swatch metadata for product options
type ProductOptionSwatch struct {
	// Type is the swatch type ("color" or "image")
	Type string `json:"type"`
	// Value is the hex color or display value
	Value string `json:"value,omitempty"`
	// MediaID is the reference to media asset when type=image
	MediaID string `json:"media_id,omitempty"`
}

// ProductOptionValue represents a single value for a product option
type ProductOptionValue struct {
	// Value is the option value label
	Value string `json:"value"`
	// SortOrder is the optional order for display
	SortOrder int `json:"sort_order,omitempty"`
	// Swatch is the optional swatch metadata
	Swatch *ProductOptionSwatch `json:"swatch,omitempty"`
	// Metadata contains additional metadata
	Metadata map[string]string `json:"metadata,omitempty"`
}

// ProductOption represents a product option (e.g., Color, Size)
type ProductOption struct {
	// Name is the option name (e.g., "Color")
	Name string `json:"name"`
	// SortOrder is the optional order for display
	SortOrder int `json:"sort_order,omitempty"`
	// Values is the list of allowed values for this option
	Values []*ProductOptionValue `json:"values"`
	// Metadata contains optional metadata
	Metadata map[string]string `json:"metadata,omitempty"`
}

// ProductVariant represents a specific variant of a product
type ProductVariant struct {
	// ID is the unique identifier for the variant
	ID string `json:"id"`
	// OptionValues contains the selected option value per option name
	OptionValues map[string]string `json:"option_values"`
	// Price is the variant price
	Price *AssetAmount `json:"price,omitempty"`
	// MediaOrder contains ordered media asset IDs for the variant
	MediaOrder []string `json:"media_order,omitempty"`
	// InventoryQuantity is the optional inventory quantity
	InventoryQuantity *int `json:"inventory_quantity,omitempty"`
	// Active indicates if the variant is active/available
	Active bool `json:"active"`
	// Version is the catalog version that produced this variant
	Version int64 `json:"version"`
	// SellingPlanGroupIDs contains variant-level selling plan group associations
	SellingPlanGroupIDs []string `json:"selling_plan_group_ids,omitempty"`
	// SellingPlanGroups contains expanded selling plan groups with plans (only populated when variant has override)
	SellingPlanGroups []*SellingPlanGroupWithPlans `json:"selling_plan_groups,omitempty"`
	// SellingPlanPricing contains pre-calculated subscription pricing per selling plan
	SellingPlanPricing []*SellingPlanPricing `json:"selling_plan_pricing,omitempty"`
}

// Product represents a product in the catalog
type Product struct {
	// ID is the unique identifier for the product
	ID string `json:"id"`
	// Name is the name of the product
	Name string `json:"name"`
	// Price is the legacy price field for simple products
	Price *AssetAmount `json:"price,omitempty"`
	// FixedPrice is the explicit price for simple products
	FixedPrice *AssetAmount `json:"fixed_price,omitempty"`
	// Description is the detailed description of the product
	Description string `json:"description"`
	// TaxCode is the tax classification code for the product
	TaxCode string `json:"tax_code"`
	// Metadata contains additional custom data for the product
	Metadata map[string]string `json:"metadata"`
	// DefaultCurrency is the base currency for the product/variants
	DefaultCurrency string `json:"default_currency"`
	// MediaOrder contains ordered media asset IDs at product level
	MediaOrder []string `json:"media_order,omitempty"`
	// CollectShippingAddress indicates whether to collect shipping address during checkout
	CollectShippingAddress bool `json:"collect_shipping_address"`
	// CollectTaxAddress indicates whether to collect billing/tax address during checkout
	CollectTaxAddress bool `json:"collect_tax_address"`
	// RequiresSellingPlan indicates if subscription is required (true=subscription only, false=one-time or subscription)
	RequiresSellingPlan bool `json:"requires_selling_plan"`
	// Options contains variant option definitions
	Options []*ProductOption `json:"options,omitempty"`
	// Variants contains variant combinations
	Variants []*ProductVariant `json:"variants,omitempty"`
	// Active indicates if the product is currently available
	Active bool `json:"active"`
	// UpdatedAt is the Unix timestamp when the product was last updated
	UpdatedAt int64 `json:"updated_at"`
	// CreatedAt is the Unix timestamp when the product was created
	CreatedAt int64 `json:"created_at"`
	// Version is the catalog version associated with this product
	Version int64 `json:"version"`
	// Includes contains included media assets with signed URLs
	Includes *PaymentLinkIncludes `json:"includes,omitempty"`
	// SellingPlanGroupIDs contains product-level selling plan group associations
	SellingPlanGroupIDs []string `json:"selling_plan_group_ids,omitempty"`
	// SellingPlanGroups contains expanded selling plan groups with plans (use expand=selling_plans query param)
	SellingPlanGroups []*SellingPlanGroupWithPlans `json:"selling_plan_groups,omitempty"`
}

// ProductItem represents a product with quantity
type ProductItem struct {
	// Product contains the referenced product information
	Product *Product `json:"product"`
	// Quantity is the number of product items
	Quantity int `json:"quantity"`
}

// ================================
// Request Types
// ================================

// ProductCreateRequest creates a new product
type ProductCreateRequest struct {
	Product `json:",inline"`
}

// ProductUpdateRequest updates an existing product
type ProductUpdateRequest struct {
	ProductCreateRequest `json:",inline"`
}

// ProductPublishRequest publishes a product version
type ProductPublishRequest struct {
	ExpectedVersion int64 `json:"expected_version"`
}

// ProductDeleteRequest deletes a product (no fields needed)
type ProductDeleteRequest struct{}

// ProductListRequest lists products with filters
type ProductListRequest struct {
	Page     int32 `json:"page" binding:"min=0" form:"page"`
	PageSize int32 `json:"page_size" binding:"required,min=1,max=50" form:"page_size"`
	Active   *bool `json:"active" form:"active"`
}

// ================================
// Response Types
// ================================

// ProductListResp represents paginated product list
type ProductListResp struct {
	Products []*Product `json:"products"`
	Total    int64      `json:"total"`
	Page     int32      `json:"page"`
	PageSize int32      `json:"page_size"`
}

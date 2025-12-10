package developer

// ================================
// Core PaymentLink Types
// ================================

// PaymentLinkVariant represents a variant available in a payment link
type PaymentLinkVariant struct {
	// VariantID is the identifier of the product variant
	VariantID string `json:"variant_id"`
	// Variant contains the variant details returned for convenience
	Variant *ProductVariant `json:"variant,omitempty"`
	// Quantity is the default quantity for this variant
	Quantity int `json:"quantity"`
	// MinQuantity is the optional minimum quantity for add-ons
	MinQuantity *int `json:"min_quantity,omitempty"`
	// MaxQuantity is the optional maximum quantity for add-ons
	MaxQuantity *int `json:"max_quantity,omitempty"`
	// IsPrimary indicates if this variant is part of the primary selection
	IsPrimary bool `json:"is_primary"`
}

// PaymentLinkPriceRange represents the price range for a payment link
type PaymentLinkPriceRange struct {
	// Min is the minimum price across primary variants
	Min *AssetAmount `json:"min"`
	// Max is the maximum price across primary variants
	Max *AssetAmount `json:"max"`
}

// PaymentLinkMedia represents media assets for a payment link
type PaymentLinkMedia struct {
	// ID is the media identifier
	ID string `json:"id"`
	// URL is the CDN URL for the asset
	URL string `json:"url"`
	// Width is the width in pixels
	Width int `json:"width,omitempty"`
	// Height is the height in pixels
	Height int `json:"height,omitempty"`
	// AltText is the alternative text for accessibility
	AltText string `json:"alt_text,omitempty"`
	// ContentType is the MIME type of the asset
	ContentType string `json:"content_type,omitempty"`
}

// PaymentLinkIncludes contains included resources for a payment link
type PaymentLinkIncludes struct {
	// Media contains media assets referenced by the product/variants
	Media []*PaymentLinkMedia `json:"media,omitempty"`
}

// PaymentLink represents a payment link for purchasing products
type PaymentLink struct {
	// ID is the unique identifier for the payment link
	ID string `json:"id"`
	// ProductItems is the legacy flattened product list
	ProductItems []*ProductItem `json:"product_items,omitempty"`
	// Product is the product associated with the payment link
	Product *Product `json:"product,omitempty"`
	// PrimaryVariants are primary variants selectable by the buyer
	PrimaryVariants []*PaymentLinkVariant `json:"primary_variants,omitempty"`
	// AddonVariants are optional add-on variants
	AddonVariants []*PaymentLinkVariant `json:"addon_variants,omitempty"`
	// VariantConfig contains UI and default selection metadata
	VariantConfig map[string]any `json:"variant_config,omitempty"`
	// PriceRange is the price range across primary variants
	PriceRange *PaymentLinkPriceRange `json:"price_range,omitempty"`
	// TotalPrice is the computed total price (legacy behaviour)
	TotalPrice *AssetAmount `json:"total_price,omitempty"`
	// Active indicates if the payment link is currently active
	Active bool `json:"active"`
	// UpdatedAt is the Unix timestamp when the link was last updated
	UpdatedAt int64 `json:"updated_at"`
	// CreatedAt is the Unix timestamp when the link was created
	CreatedAt int64 `json:"created_at"`
	// URL is the URL to access the payment link
	URL *string `json:"url"`
	// Includes contains included shared resources (e.g., media)
	Includes *PaymentLinkIncludes `json:"includes,omitempty"`
}

// PaymentLinkDetails contains detailed payment link information including merchant info
type PaymentLinkDetails struct {
	// MerchantID is the ID of the merchant who created the payment link
	MerchantID string `json:"merchant_id"`
	// MerchantName is the name of the merchant
	MerchantName string `json:"merchant_name"`
	// PublicKey is the public API key for this merchant
	PublicKey string `json:"public_key"`
	// PaymentLink contains the details of the payment link
	PaymentLink *PaymentLink `json:"payment_link"`
}

// ================================
// Request Types
// ================================

// PaymentLinkAddonVariantRequest represents addon variant configuration
type PaymentLinkAddonVariantRequest struct {
	// VariantID is the identifier of the variant
	VariantID string `json:"variant_id"`
	// MinQuantity is the optional minimum quantity for this addon
	MinQuantity *int `json:"min_quantity,omitempty"`
	// MaxQuantity is the optional maximum quantity for this addon
	MaxQuantity *int `json:"max_quantity,omitempty"`
}

// PaymentLinkProductItemRequest represents a product item in payment link
type PaymentLinkProductItemRequest struct {
	// ProductID is the ID of the product
	ProductID string `json:"product_id" binding:"required"`
	// Quantity is the quantity for this product item
	Quantity int `json:"quantity" binding:"required,min=1"`
}

// PaymentLinkCreateRequest creates a new payment link
type PaymentLinkCreateRequest struct {
	// ProductID is the ID of the product for this payment link
	ProductID string `json:"product_id"`
	// DefaultVariantID is the default variant to select
	DefaultVariantID string `json:"default_variant_id"`
	// PrimaryVariantIDs are the primary variants available for selection
	PrimaryVariantIDs []string `json:"primary_variant_ids"`
	// AddonVariants are the optional addon variants
	AddonVariants []PaymentLinkAddonVariantRequest `json:"addon_variants"`
	// VariantConfig contains UI and default selection metadata
	VariantConfig map[string]any `json:"variant_config"`
}

// PaymentLinkUpdateRequest updates a payment link
type PaymentLinkUpdateRequest struct {
	// Active indicates whether the payment link is active
	Active *bool `json:"active"`
}

// PaymentLinkListRequest lists payment links with filters
type PaymentLinkListRequest struct {
	// Page is the page number for pagination
	Page int32 `json:"page" binding:"min=0" form:"page"`
	// PageSize is the number of items per page
	PageSize int32 `json:"page_size" binding:"required,min=1,max=50" form:"page_size"`
	// Active filters by active status
	Active *bool `json:"active" form:"active"`
	// Product filters by product ID
	Product string `json:"product" form:"product"`
}

// PaymentLinkPublicGetRequest gets public payment link details
type PaymentLinkPublicGetRequest struct {
	// ID is the payment link ID from URL path
	ID string `json:"-"`
}

// ================================
// Response Types
// ================================

// PaymentLinkListResponse represents paginated payment link list
type PaymentLinkListResponse struct {
	// PaymentLinks is the list of payment links
	PaymentLinks []*PaymentLink `json:"payment_links"`
	// Total is the total number of payment links matching the filter
	Total int64 `json:"total"`
	// Page is the current page number
	Page int32 `json:"page"`
	// PageSize is the number of items per page
	PageSize int32 `json:"page_size"`
}

// PaymentLinkPublicGetResponse is an alias to PaymentLinkDetails
type PaymentLinkPublicGetResponse = PaymentLinkDetails

package developer

// MerchantAddress represents a merchant's blockchain address
type MerchantAddress struct {
	// ID is the unique identifier for the merchant address
	ID string `json:"id"`
	// MerchantID is the ID of the merchant who owns this address
	MerchantID string `json:"merchant_id"`
	// Network is the blockchain network (e.g., ETH, BTC)
	Network string `json:"network"`
	// Address is the actual blockchain address
	Address string `json:"address"`
	// VerificationID is the reference to the verification record ID
	VerificationID *string `json:"verification_id"`
	// Status is the current status of the address (created, verifying, success, failed)
	Status string `json:"status"`
	// Alias is the user-defined name for this address
	Alias string `json:"alias"`
	// CreatedAt is the Unix timestamp when the address was created
	CreatedAt int64 `json:"created_at"`
	// UpdatedAt is the Unix timestamp when the address was last updated
	UpdatedAt int64 `json:"updated_at"`
	// Verification contains nested verification details
	Verification *AddressVerification `json:"verification"`
}

// AddressVerification represents a verification record for a merchant's address
type AddressVerification struct {
	// ID is the unique identifier for the verification record
	ID string `json:"id"`
	// Status is the current verification status
	Status string `json:"status"`
	// MerchantAddressID is the ID of the merchant address being verified
	MerchantAddressID string `json:"merchant_address_id"`
	// AssetID is the asset used for verification
	AssetID string `json:"asset_id"`
	// VerifiedAt is the Unix timestamp when verification was completed
	VerifiedAt *int64 `json:"verified_at"`
	// TriedTimes is the number of verification attempts
	TriedTimes int `json:"tried_times"`
	// AmlStatus is the Anti-Money Laundering status
	AmlStatus string `json:"aml_status"`
}

// ================================
// Request Types
// ================================

// MerchantAddressCreateRequest creates a new merchant address
type MerchantAddressCreateRequest struct {
	// Network is the blockchain network (e.g., "ETH", "BTC", "TRX")
	Network string `json:"network"`
	// Address is the blockchain address to register
	Address string `json:"address"`
}

// MerchantAddressVerifyRequest verifies a merchant address
type MerchantAddressVerifyRequest struct {
	// Amount is the verification amount sent to the address
	Amount string `json:"amount" binding:"required"`
}

// MerchantAddressUpdateRequest updates a merchant address
type MerchantAddressUpdateRequest struct {
	// Alias is the user-defined name for the address
	Alias *string `json:"alias"`
}

// MerchantAddressListRequest lists merchant addresses with filters
type MerchantAddressListRequest struct {
	// Network filters addresses by blockchain network
	Network *string `json:"network,omitempty" form:"network"`
	// Page is the page number (zero-based)
	Page int32 `json:"page" binding:"min=0" form:"page"`
	// PageSize is the number of items per page (max 50)
	PageSize int32 `json:"page_size" binding:"required,min=1,max=50" form:"page_size"`
}

// ================================
// Response Types
// ================================

// MerchantAddressListResp represents paginated merchant address list
type MerchantAddressListResp struct {
	// MerchantAddresses is the list of merchant addresses
	MerchantAddresses []*MerchantAddress `json:"merchant_addresses"`
	// Total is the total number of addresses matching the filter
	Total int64 `json:"total"`
	// Page is the current page number
	Page int32 `json:"page"`
	// PageSize is the number of items per page
	PageSize int32 `json:"page_size"`
}

package developer

// MerchantAddress represents a merchant's blockchain address
type MerchantAddress struct {
	ID             string  `json:"id"`              // Unique identifier for the merchant address
	MerchantID     string  `json:"merchant_id"`     // ID of the merchant who owns this address
	Network        string  `json:"network"`         // Blockchain network (e.g., ETH, BTC)
	Address        string  `json:"address"`         // The actual blockchain address
	VerificationID *string `json:"verification_id"` // Reference to the verification record ID
	Status         string  `json:"status"`          // Current status of the address (created, verifying, success, failed)
	Alias          string  `json:"alias"`           // User-defined name for this address
	// Active         bool                 `json:"active"`
	CreatedAt    int64                `json:"created_at"`   // Timestamp when the address was created
	UpdatedAt    int64                `json:"updated_at"`   // Timestamp when the address was last updated
	Verification *AddressVerification `json:"verification"` // Nested verification details
}

// AddressVerification represents a verification record for a merchant's address
type AddressVerification struct {
	ID                string `json:"id"`                  // Unique identifier for the verification record
	Status            string `json:"status"`              // Current verification status
	MerchantAddressID string `json:"merchant_address_id"` // ID of the merchant address being verified
	AssetID           string `json:"asset_id"`            // Asset used for verification
	VerifiedAt        *int64 `json:"verified_at"`         // Timestamp when verification was completed
	TriedTimes        int    `json:"tried_times"`         // Number of verification attempts
	AmlStatus         string `json:"aml_status"`          // Anti-Money Laundering status
}

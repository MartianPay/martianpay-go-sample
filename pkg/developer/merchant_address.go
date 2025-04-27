package developer

type MerchantAddress struct {
	ID             string  `json:"id"`
	MerchantID     string  `json:"merchant_id"`
	Network        string  `json:"network"`
	Address        string  `json:"address"`
	VerificationID *string `json:"verification_id"`
	Status         string  `json:"status"`
	Alias          string  `json:"alias"`
	// Active         bool                 `json:"active"`
	CreatedAt    int64                `json:"created_at"`
	UpdatedAt    int64                `json:"updated_at"`
	Verification *AddressVerification `json:"verification"`
}

type AddressVerification struct {
	ID                string `json:"id"`
	Status            string `json:"status"`
	MerchantAddressID string `json:"merchant_address_id"`
	AssetID           string `json:"asset_id"`
	VerifiedAt        *int64 `json:"verified_at"`
	TriedTimes        int    `json:"tried_times"`
	AmlStatus         string `json:"aml_status"`
}

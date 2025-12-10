package developer

import "github.com/dchest/uniuri"

const (
	// MerchantIDLength is the length of the merchant ID suffix
	MerchantIDLength = 24
	// MerchantIDPrefix is the prefix for merchant IDs
	MerchantIDPrefix = "accu_"
	// MerchantObject is the object type identifier for merchants
	MerchantObject = "merchant"
)

// MerchantParams contains parameters for creating or updating a merchant
type MerchantParams struct {
	// Name is the merchant name
	Name *string `json:"name" binding:"required"`
	// Country is the country code
	Country *string `json:"country" binding:"required"`
	// Timezone is the timezone for the merchant
	Timezone *string `json:"timezone" binding:"required"`
}

// BalanceDetail represents the balance breakdown for a specific currency
type BalanceDetail struct {
	// Currency is the currency code
	Currency string `json:"currency"`
	// AvailableBalance is the available balance amount
	AvailableBalance string `json:"available_balance"`
	// PendingBalance is the pending balance amount
	PendingBalance string `json:"pending_balance"`
	// LockedBalance is the locked balance amount
	LockedBalance string `json:"locked_balance"`
	// FrozenBalance is the frozen balance amount
	FrozenBalance string `json:"frozen_balance"`
	// TotalBalance is the total balance amount
	TotalBalance string `json:"total_balance"`
}

// MerchantBalance represents the merchant's balance information
type MerchantBalance struct {
	// Currency is the currency code
	Currency string `json:"currency"`
	// AvailableBalance is the available balance amount
	AvailableBalance string `json:"available_balance"`
	// PendingBalance is the pending balance amount
	PendingBalance string `json:"pending_balance"`
	// LockedBalance is the locked balance amount
	LockedBalance string `json:"locked_balance"`
	// FrozenBalance is the frozen balance amount
	FrozenBalance string `json:"frozen_balance"`
	// TotalBalance is the total balance amount
	TotalBalance string `json:"total_balance"`
	// BalanceDetails is the list of balance details by currency
	BalanceDetails []*BalanceDetail `json:"balance_details"`
}

// Merchant represents a merchant account
type Merchant struct {
	// ID is the unique identifier for the merchant
	ID string `json:"id"`
	// Object is the object type, always "merchant"
	Object string `json:"object"`
	// CreatedAt is the Unix timestamp when the merchant was created
	CreatedAt int64 `json:"created_at"`
	// Name is the merchant name
	Name string `json:"name"`
	// Uid is the merchant unique identifier
	Uid string `json:"uid"`
	// Email is the merchant email address
	Email string `json:"email"`
	// Country is the country or region
	Country string `json:"country"`
	// Timezone is the timezone
	Timezone string `json:"timezone"`
	// Owner indicates whether the merchant is the owner
	Owner bool `json:"owner"`
	// Roles is the list of roles assigned to the merchant
	Roles []Role `json:"roles"`
}

func GenerateMerchantID() string {
	uniqueString := uniuri.NewLen(MerchantIDLength)
	return MerchantIDPrefix + uniqueString
}

// MerchantContractParams contains parameters for creating or updating a merchant contract
type MerchantContractParams struct {
	// Name is the merchant name
	Name *string `json:"name" binding:"required"`
	// Description is the merchant description
	Description *string `json:"description"`
	// Email is the merchant email
	Email *string `json:"email"`
	// Phone is the merchant phone
	Phone *string `json:"phone"`
	// CompanyNumber is the company registration number
	CompanyNumber *string `json:"company_number"`
	// LegalPerson is the legal representative name
	LegalPerson *string `json:"legal_person"`
	// BankAccount is the bank account number
	BankAccount *string `json:"bank_account"`
	// BankName is the bank name
	BankName *string `json:"bank_name"`
	// BankSwiftCode is the SWIFT code
	BankSwiftCode *string `json:"bank_swift_code"`
	// Country is the country or region
	Country *string `json:"country"`
	// Address is the company address
	Address *string `json:"address"`
	// TaxNumber is the tax number
	TaxNumber *string `json:"tax_number"`
	// Metadata contains additional metadata
	Metadata map[string]string `json:"metadata"`
	// ExpiresAt is the contract expiration time as Unix timestamp
	ExpiresAt *int64 `json:"expires_at"`
}

const (
	// MerchantContractStatusPending indicates the contract is pending approval
	MerchantContractStatusPending = "pending"
	// MerchantContractStatusActive indicates the contract is active
	MerchantContractStatusActive = "active"
	// MerchantContractStatusRejected indicates the contract was rejected
	MerchantContractStatusRejected = "rejected"
	// MerchantContractStatusTerminated indicates the contract was terminated
	MerchantContractStatusTerminated = "terminated"

	// MerchantContractIDLength is the length of the merchant contract ID suffix
	MerchantContractIDLength = 24
	// MerchantContractIDPrefix is the prefix for merchant contract IDs
	MerchantContractIDPrefix = "mc_"
	// MerchantContractObject is the object type identifier for merchant contracts
	MerchantContractObject = "merchant_contract"
)

// MerchantContract represents a merchant contract
type MerchantContract struct {
	// ID is the unique identifier for the merchant contract
	ID string `json:"id"`
	// Object is the object type, always "merchant_contract"
	Object string `json:"object"`
	// MerchantID is the ID of the merchant this contract belongs to
	MerchantID string `json:"merchant_id"`
	// Name is the name of the merchant
	Name string `json:"name"`
	// Description is the description of the merchant business
	Description string `json:"description"`
	// Email is the contact email address
	Email string `json:"email"`
	// Phone is the contact phone number
	Phone string `json:"phone"`
	// CompanyNumber is the company registration number
	CompanyNumber string `json:"company_number"`
	// LegalPerson is the legal representative name
	LegalPerson string `json:"legal_person"`
	// BankAccount is the bank account number
	BankAccount string `json:"bank_account"`
	// BankName is the bank name
	BankName string `json:"bank_name"`
	// BankSwiftCode is the bank SWIFT/BIC code for international transfers
	BankSwiftCode string `json:"bank_swift_code"`
	// Country is the country of registration
	Country string `json:"country"`
	// Address is the company address
	Address string `json:"address"`
	// TaxNumber is the tax identification number
	TaxNumber string `json:"tax_number"`
	// Metadata contains additional metadata for the contract
	Metadata map[string]string `json:"metadata"`
	// Status is the contract status (pending, active, rejected, terminated)
	Status string `json:"status"`
	// Created is the Unix timestamp when the contract was created
	Created int64 `json:"created"`
	// Deleted indicates whether the contract has been deleted
	Deleted bool `json:"deleted"`
	// ExpiresAt is the Unix timestamp when the contract expires
	ExpiresAt int64 `json:"expires_at"`
	// FeeRate is the transaction fee rate in ten thousandths (e.g. 200 = 2%)
	FeeRate int64 `json:"fee_rate"`
	// FixedFee is the fixed amount fee per transaction in cents
	FixedFee int64 `json:"fixed_fee"`
	// Reason is the reason for rejection if status is "rejected"
	Reason string `json:"reason"`
	// Livemode indicates whether this contract is in live mode or test mode
	Livemode bool `json:"livemode"`
}

func GenerateMerchantContractID() string {
	uniqueString := uniuri.NewLen(MerchantContractIDLength)
	return MerchantContractIDPrefix + uniqueString
}

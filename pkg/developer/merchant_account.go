package developer

import (
	"github.com/dchest/uniuri"
)

const (
	MerchantAccountTypeWallet = "wallet"
	MerchantAccountTypeBank   = "bank"

	MerchantAccountObject   = "merchant_account"
	MerchantAccountPrefix   = "ma_"
	MerchantAccountIDLength = 24
)

func GenerateMerchantAccountID() string {
	return MerchantAccountPrefix + uniuri.NewLen(MerchantAccountIDLength)
}

type MerchantAccount struct {
	ID            string       `json:"id"`
	Object        string       `json:"object"`
	CreatedAt     int64        `json:"created_at"`
	AccountType   string       `json:"account_type"`
	Alias         string       `json:"alias"`
	IsDefault     bool         `json:"is_default"`
	AccountStatus string       `json:"account_status"`
	BankAccount   *BankAccount `json:"bank_account,omitempty"`
}

type BankAccount struct {
	AccountHolderName string `json:"account_holder_name"`
	BankName          string `json:"bank_name"`
	BranchName        string `json:"branch_name"`
	SwiftCode         string `json:"swift_code"`
	AccountNumber     string `json:"account_number"`
}

type BalanceDetail struct {
	Currency         string `json:"currency"`          // Currency code (Asset ID, e.g., "USDT-Tron", "USDC-Ethereum")
	AvailableBalance string `json:"available_balance"` // Available balance
	PendingBalance   string `json:"pending_balance"`   // Pending balance
	LockedBalance    string `json:"locked_balance"`    // Locked balance
	FrozenBalance    string `json:"frozen_balance"`    // Frozen balance
	TotalBalance     string `json:"total_balance"`     // Total balance
}

type MerchantBalance struct {
	Currency         string           `json:"currency"`          // Currency code
	AvailableBalance string           `json:"available_balance"` // Available balance
	PendingBalance   string           `json:"pending_balance"`   // Pending balance
	LockedBalance    string           `json:"locked_balance"`    // Locked balance
	FrozenBalance    string           `json:"frozen_balance"`    // Frozen balance
	TotalBalance     string           `json:"total_balance"`     // Total balance
	BalanceDetails   []*BalanceDetail `json:"balance_details"`   // List of balance details
}

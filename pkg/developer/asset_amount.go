// asset_amount.go contains types for representing asset amounts with precision.
// It provides structures for handling cryptocurrency and fiat amounts with proper
// decimal precision and formatting.
package developer

import "github.com/shopspring/decimal"

// AssetAmount represents a cryptocurrency asset with its amount and decimal precision
type AssetAmount struct {
	// AssetId is the unique identifier for the asset/currency
	AssetId string `json:"asset_id"`
	// Amount is the decimal amount of the asset
	Amount decimal.Decimal `json:"amount"`
	// DecimalDigits is the number of decimal places for this asset (not serialized to JSON)
	DecimalDigits int `json:"-"`
}

// NewAssetAmount creates a new AssetAmount with the specified decimal amount
func NewAssetAmount(amount decimal.Decimal, assetId string, decimals int) *AssetAmount {
	return &AssetAmount{
		AssetId:       assetId,
		Amount:        amount,
		DecimalDigits: decimals,
	}
}

// NewAssetAmountFromBigInt creates a new AssetAmount from an integer amount by applying decimal shift
func NewAssetAmountFromBigInt(intAmount decimal.Decimal, assetId string, decimals int) *AssetAmount {
	amount := intAmount.Shift(int32(-decimals))
	return &AssetAmount{
		AssetId:       assetId,
		Amount:        amount,
		DecimalDigits: int(decimals),
	}
}

// BigInt returns the amount as an integer by shifting by decimal digits
func (a *AssetAmount) BigInt() decimal.Decimal {
	return a.Amount.Shift(int32(a.DecimalDigits))
}

// Value returns the decimal amount value
func (a *AssetAmount) Value() decimal.Decimal {
	return a.Amount
}

// DecimalOverflow checks if the amount has decimal overflow
func (a *AssetAmount) DecimalOverflow() bool {
	inter := a.BigInt().Truncate(0)
	return a.Amount.Cmp(inter) == 0
}

// IsValidPrice checks if the asset amount represents a valid price (positive and no decimal overflow)
func (a *AssetAmount) IsValidPrice() bool {
	if a == nil {
		return false
	}
	return a.Value().Sign() > 0 && !a.DecimalOverflow()
}

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

func NewAssetAmount(amount decimal.Decimal, assetId string, decimals int) *AssetAmount {
	return &AssetAmount{
		AssetId:       assetId,
		Amount:        amount,
		DecimalDigits: decimals,
	}
}

func NewAssetAmountFromBigInt(intAmount decimal.Decimal, assetId string, decimals int) *AssetAmount {
	amount := intAmount.Shift(int32(-decimals))
	return &AssetAmount{
		AssetId:       assetId,
		Amount:        amount,
		DecimalDigits: int(decimals),
	}
}
func (a *AssetAmount) BigInt() decimal.Decimal {
	return a.Amount.Shift(int32(a.DecimalDigits))
}

func (a *AssetAmount) Value() decimal.Decimal {
	return a.Amount
}

func (a *AssetAmount) DecimalOverflow() bool {
	inter := a.BigInt().Truncate(0)
	return a.Amount.Cmp(inter) == 0
}

func (a *AssetAmount) IsValidPrice() bool {
	if a == nil {
		return false
	}
	return a.Value().Sign() > 0 && !a.DecimalOverflow()
}

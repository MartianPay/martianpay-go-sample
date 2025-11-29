package developer

type Asset struct {
	Id          string `json:"id"`
	DisplayName string `json:"display_name"`
	Coin        string `json:"coin"`
	IsFiat      bool   `json:"is_fiat"`
	Decimals    int    `json:"decimals"`
	Payable     bool   `json:"payable"`
	*CryptoAssetParams
	*FiatAssetParams
}

type CryptoAssetParams struct {
	Network            string `json:"network"`
	IsMainnet          bool   `json:"is_mainnet"`
	ContractAddress    string `json:"contract_address"`
	AddressUrlTemplate string `json:"address_url_template"`
	TxUrlTemplate      string `json:"tx_url_template"`
	Token              string `json:"token"`
	ChainId            int64  `json:"chain_id"`
}

type FiatAssetParams struct {
	Symbol   string `json:"symbol"`
	Provider string `json:"provider"`
}

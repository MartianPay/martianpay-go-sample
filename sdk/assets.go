package martianpay

import "github.com/MartianPay/martianpay-go-sample/pkg/developer"

type AssetListResponse struct {
	Assets []*developer.Asset `json:"assets"`
}

// ListAssets retrieves all available assets
func (c *Client) ListAssets() (*AssetListResponse, error) {
	var response AssetListResponse
	err := c.sendRequest("GET", "/v1/assets", nil, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

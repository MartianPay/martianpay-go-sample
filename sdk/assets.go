package martianpay

import "github.com/MartianPay/martianpay-go-sample/pkg/developer"

// ListAssets retrieves all available assets
func (c *Client) ListAssets() (*developer.AssetListResponse, error) {
	var response developer.AssetListResponse
	err := c.sendRequest("GET", "/v1/assets", nil, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// GetAllAssets gets all available assets
func (c *Client) GetAllAssets() ([]*developer.Asset, error) {
	var response []*developer.Asset
	err := c.sendRequest("GET", "/v1/assets/all", nil, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// ListAssetFees lists asset network fees
func (c *Client) ListAssetFees() (*developer.NetworkFeesResponse, error) {
	var response developer.NetworkFeesResponse
	err := c.sendRequest("GET", "/v1/assets/fees", nil, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

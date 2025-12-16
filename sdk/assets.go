// Package martianpay provides SDK methods for managing cryptocurrency assets.
// Assets represent the supported cryptocurrencies and tokens available on the MartianPay platform.
package martianpay

import "github.com/MartianPay/martianpay-go-sample/pkg/developer"

// ListAssets retrieves a list of all available cryptocurrency assets.
// Includes asset details such as symbol, name, supported networks, and trading status.
//
// Returns:
//   - *developer.AssetListResponse: List of available assets with details
//   - error: nil on success, error on failure
func (c *Client) ListAssets() (*developer.AssetListResponse, error) {
	var response developer.AssetListResponse
	err := c.sendRequest("GET", "/v1/assets", nil, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// GetAllAssets retrieves complete details for all available cryptocurrency assets.
// Returns raw asset data without pagination or additional metadata.
//
// Returns:
//   - []*developer.Asset: Array of all available assets
//   - error: nil on success, error on failure
func (c *Client) GetAllAssets() ([]*developer.Asset, error) {
	var response []*developer.Asset
	err := c.sendRequest("GET", "/v1/assets/all", nil, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// ListAssetFees retrieves current network fees for all supported cryptocurrency assets.
// Network fees are required for blockchain transactions and vary by asset and network congestion.
//
// Returns:
//   - *developer.NetworkFeesResponse: Current network fees for each asset and network
//   - error: nil on success, error on failure
func (c *Client) ListAssetFees() (*developer.NetworkFeesResponse, error) {
	var response developer.NetworkFeesResponse
	err := c.sendRequest("GET", "/v1/assets/fees", nil, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

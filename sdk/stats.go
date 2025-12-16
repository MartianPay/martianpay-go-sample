// Package martianpay provides SDK methods for accessing merchant statistics and balance information.
// Stats endpoints provide real-time data about merchant account balances and financial metrics.
package martianpay

import "github.com/MartianPay/martianpay-go-sample/pkg/developer"

// GetBalance retrieves the current balance for the merchant account.
// Shows available balance, pending balance, and reserved funds for each cryptocurrency asset.
//
// Returns:
//   - *developer.MerchantBalance: Current balance details across all assets
//   - error: nil on success, error on failure
func (c *Client) GetBalance() (*developer.MerchantBalance, error) {
	var response developer.MerchantBalance
	err := c.sendRequest("GET", "/v1/stats/balance", nil, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

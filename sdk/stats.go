package martianpay

import "github.com/MartianPay/martianpay-go-sample/pkg/developer"

// GetBalance retrieves the merchant's balance
func (c *Client) GetBalance() (*developer.MerchantBalance, error) {
	var response developer.MerchantBalance
	err := c.sendRequest("GET", "/v1/stats/balance", nil, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

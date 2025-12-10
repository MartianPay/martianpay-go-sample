package martianpay

import (
	"fmt"

	"github.com/MartianPay/martianpay-go-sample/pkg/developer"
)

// ListOrders lists orders
func (c *Client) ListOrders(params *developer.OrderListRequest) (*developer.OrderListResponse, error) {
	var resp developer.OrderListResponse
	err := c.sendRequestWithQuery("GET", "/v1/orders", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetOrder gets order details
func (c *Client) GetOrder(orderNumber string) (*developer.OrderDetail, error) {
	path := fmt.Sprintf("/v1/orders/%s", orderNumber)
	var resp developer.OrderDetail
	err := c.sendRequest("GET", path, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

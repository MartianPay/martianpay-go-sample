// Package martianpay provides SDK methods for managing orders.
// Orders represent completed purchases containing line items, customer information, and payment details.
package martianpay

import (
	"fmt"

	"github.com/MartianPay/martianpay-go-sample/pkg/developer"
)

// ListOrders retrieves a paginated list of orders.
// Orders represent completed purchases and contain line items, customer info, and payment details.
//
// Parameters:
//   - params: Query parameters including pagination, filters by status, customer, date range, etc.
//
// Returns:
//   - *developer.OrderListResponse: List of orders with pagination metadata
//   - error: nil on success, error on failure
func (c *Client) ListOrders(params *developer.OrderListRequest) (*developer.OrderListResponse, error) {
	var resp developer.OrderListResponse
	err := c.sendRequestWithQuery("GET", "/v1/orders", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetOrder retrieves detailed information about a specific order.
// Includes line items, customer details, payment information, and fulfillment status.
//
// Parameters:
//   - orderNumber: The unique order number/ID
//
// Returns:
//   - *developer.OrderDetail: Complete order details
//   - error: nil on success, error on failure (e.g., order not found)
func (c *Client) GetOrder(orderNumber string) (*developer.OrderDetail, error) {
	path := fmt.Sprintf("/v1/orders/%s", orderNumber)
	var resp developer.OrderDetail
	err := c.sendRequest("GET", path, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

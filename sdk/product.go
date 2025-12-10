package martianpay

import (
	"fmt"

	"github.com/MartianPay/martianpay-go-sample/pkg/developer"
)

// ListProducts lists products with pagination
func (c *Client) ListProducts(params *developer.ProductListRequest) (*developer.ProductListResp, error) {
	var resp developer.ProductListResp
	err := c.sendRequestWithQuery("GET", "/v1/products", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateProduct creates a product with variants
func (c *Client) CreateProduct(params *developer.ProductCreateRequest) (*developer.Product, error) {
	var resp developer.Product
	err := c.sendRequest("POST", "/v1/products", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetProduct retrieves product details
func (c *Client) GetProduct(productID string) (*developer.Product, error) {
	path := fmt.Sprintf("/v1/products/%s", productID)
	var resp developer.Product
	err := c.sendRequest("GET", path, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateProduct updates product configuration
func (c *Client) UpdateProduct(productID string, params *developer.ProductUpdateRequest) (*developer.Product, error) {
	path := fmt.Sprintf("/v1/products/%s", productID)
	var resp developer.Product
	err := c.sendRequest("POST", path, params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteProduct deletes inactive product
func (c *Client) DeleteProduct(productID string) error {
	path := fmt.Sprintf("/v1/products/%s", productID)
	return c.sendRequest("DELETE", path, nil, nil)
}

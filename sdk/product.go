// Package martianpay provides SDK methods for managing products.
// Products represent items or services that merchants sell, including variants, options, and pricing.
package martianpay

import (
	"fmt"

	"github.com/MartianPay/martianpay-go-sample/pkg/developer"
)

// ListProducts retrieves a paginated list of products.
// Products can be filtered by active status and other parameters.
//
// Parameters:
//   - params: Query parameters including pagination (page, page_size) and filters
//
// Returns:
//   - *developer.ProductListResp: List of products with pagination metadata
//   - error: nil on success, error on failure
func (c *Client) ListProducts(params *developer.ProductListRequest) (*developer.ProductListResp, error) {
	var resp developer.ProductListResp
	err := c.sendRequestWithQuery("GET", "/v1/products", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateProduct creates a new product with options and variants.
// Products can have multiple options (e.g., size, color) and variants (combinations of options).
//
// Parameters:
//   - params: Product creation request including name, description, options, variants, etc.
//
// Returns:
//   - *developer.Product: The created product with assigned ID
//   - error: nil on success, error on failure
func (c *Client) CreateProduct(params *developer.ProductCreateRequest) (*developer.Product, error) {
	var resp developer.Product
	err := c.sendRequest("POST", "/v1/products", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetProduct retrieves detailed information about a specific product.
// Includes product options, variants, selling plan groups, and metadata.
//
// Parameters:
//   - productID: The unique identifier of the product
//
// Returns:
//   - *developer.Product: Complete product details
//   - error: nil on success, error on failure (e.g., product not found)
func (c *Client) GetProduct(productID string) (*developer.Product, error) {
	path := fmt.Sprintf("/v1/products/%s", productID)
	var resp developer.Product
	err := c.sendRequest("GET", path, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateProduct updates an existing product's configuration.
// Requires the product's version field for optimistic locking to prevent concurrent modifications.
//
// Parameters:
//   - productID: The unique identifier of the product to update
//   - params: Updated product fields (must include version for optimistic locking)
//
// Returns:
//   - *developer.Product: The updated product with new version number
//   - error: nil on success, error on failure (e.g., version conflict)
func (c *Client) UpdateProduct(productID string, params *developer.ProductUpdateRequest) (*developer.Product, error) {
	path := fmt.Sprintf("/v1/products/%s", productID)
	var resp developer.Product
	err := c.sendRequest("POST", path, params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteProduct permanently deletes an inactive product.
// Only products with active=false can be deleted. Active products must be deactivated first.
//
// Parameters:
//   - productID: The unique identifier of the product to delete
//
// Returns:
//   - error: nil on success, error on failure (e.g., product is active or has dependencies)
func (c *Client) DeleteProduct(productID string) error {
	path := fmt.Sprintf("/v1/products/%s", productID)
	return c.sendRequest("DELETE", path, nil, nil)
}

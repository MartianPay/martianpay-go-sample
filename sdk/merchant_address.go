// Package martianpay provides SDK methods for managing merchant addresses.
// Merchant addresses are used for shipping, billing, and business location information.
package martianpay

import (
	"fmt"

	"github.com/MartianPay/martianpay-go-sample/pkg/developer"
)

// CreateMerchantAddress creates a new merchant address record.
// Addresses can be used for shipping origins, return addresses, or business locations.
//
// Parameters:
//   - req: Request containing address details (street, city, state, postal code, country, etc.)
//
// Returns:
//   - *developer.MerchantAddress: The created address with assigned ID
//   - error: nil on success, error on failure
func (c *Client) CreateMerchantAddress(req *developer.MerchantAddressCreateRequest) (*developer.MerchantAddress, error) {
	var response developer.MerchantAddress
	err := c.sendRequest("POST", "/v1/addresses", req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// GetMerchantAddress retrieves detailed information about a specific merchant address.
// Includes full address details and verification status.
//
// Parameters:
//   - id: The unique identifier of the merchant address
//
// Returns:
//   - *developer.MerchantAddress: Complete address details
//   - error: nil on success, error on failure (e.g., address not found)
func (c *Client) GetMerchantAddress(id string) (*developer.MerchantAddress, error) {
	var response developer.MerchantAddress
	err := c.sendRequest("GET", "/v1/addresses/"+id, nil, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// UpdateMerchantAddress updates an existing merchant address.
// Can modify street address, city, state, postal code, and other address fields.
//
// Parameters:
//   - id: The unique identifier of the merchant address to update
//   - req: Updated address fields
//
// Returns:
//   - *developer.MerchantAddress: The updated address details
//   - error: nil on success, error on failure
func (c *Client) UpdateMerchantAddress(id string, req *developer.MerchantAddressUpdateRequest) (*developer.MerchantAddress, error) {
	var response developer.MerchantAddress
	err := c.sendRequest("POST", fmt.Sprintf("/v1/addresses/%s", id), req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// VerifyMerchantAddress verifies a merchant address using address validation services.
// Verification checks if the address is valid and deliverable.
//
// Parameters:
//   - id: The unique identifier of the merchant address to verify
//   - req: Verification request parameters
//
// Returns:
//   - *developer.MerchantAddress: The address with updated verification status
//   - error: nil on success, error on failure (e.g., address cannot be verified)
func (c *Client) VerifyMerchantAddress(id string, req *developer.MerchantAddressVerifyRequest) (*developer.MerchantAddress, error) {
	var response developer.MerchantAddress
	err := c.sendRequest("POST", fmt.Sprintf("/v1/addresses/%s/verify", id), req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// DeleteMerchantAddress permanently deletes a merchant address.
// Addresses in use by orders or shipping configurations cannot be deleted.
//
// Parameters:
//   - id: The unique identifier of the merchant address to delete
//
// Returns:
//   - error: nil on success, error on failure (e.g., address is in use)
func (c *Client) DeleteMerchantAddress(id string) error {
	return c.sendRequest("DELETE", "/v1/addresses/"+id, nil, nil)
}

// ListMerchantAddresses retrieves a paginated list of merchant addresses.
// Can be filtered by country, state, or verification status.
//
// Parameters:
//   - req: Query parameters including pagination and filters
//
// Returns:
//   - *developer.MerchantAddressListResp: List of merchant addresses with pagination metadata
//   - error: nil on success, error on failure
func (c *Client) ListMerchantAddresses(req *developer.MerchantAddressListRequest) (*developer.MerchantAddressListResp, error) {
	var response developer.MerchantAddressListResp
	err := c.sendRequestWithQuery("GET", "/v1/addresses", req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

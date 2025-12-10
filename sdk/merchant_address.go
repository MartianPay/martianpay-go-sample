package martianpay

import (
	"fmt"

	"github.com/MartianPay/martianpay-go-sample/pkg/developer"
)

// CreateMerchantAddress creates a new merchant address
func (c *Client) CreateMerchantAddress(req *developer.MerchantAddressCreateRequest) (*developer.MerchantAddress, error) {
	var response developer.MerchantAddress
	err := c.sendRequest("POST", "/v1/addresses", req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// GetMerchantAddress retrieves a merchant address by ID
func (c *Client) GetMerchantAddress(id string) (*developer.MerchantAddress, error) {
	var response developer.MerchantAddress
	err := c.sendRequest("GET", "/v1/addresses/"+id, nil, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// UpdateMerchantAddress updates a merchant address
func (c *Client) UpdateMerchantAddress(id string, req *developer.MerchantAddressUpdateRequest) (*developer.MerchantAddress, error) {
	var response developer.MerchantAddress
	err := c.sendRequest("POST", fmt.Sprintf("/v1/addresses/%s", id), req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// VerifyMerchantAddress verifies a merchant address
func (c *Client) VerifyMerchantAddress(id string, req *developer.MerchantAddressVerifyRequest) (*developer.MerchantAddress, error) {
	var response developer.MerchantAddress
	err := c.sendRequest("POST", fmt.Sprintf("/v1/addresses/%s/verify", id), req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// DeleteMerchantAddress deletes a merchant address by ID
func (c *Client) DeleteMerchantAddress(id string) error {
	return c.sendRequest("DELETE", "/v1/addresses/"+id, nil, nil)
}

// ListMerchantAddresses retrieves a paginated list of merchant addresses
func (c *Client) ListMerchantAddresses(req *developer.MerchantAddressListRequest) (*developer.MerchantAddressListResp, error) {
	var response developer.MerchantAddressListResp
	err := c.sendRequestWithQuery("GET", "/v1/addresses", req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

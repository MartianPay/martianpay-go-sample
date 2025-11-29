package martianpay

import "github.com/MartianPay/martianpay-go-sample/pkg/developer"

// MerchantAddressCreateRequest represents the request to create a merchant address
type MerchantAddressCreateRequest struct {
	Network string `json:"network"` // Blockchain network (e.g., "Ethereum", "Tron", "Bitcoin")
	Address string `json:"address"` // The blockchain address to add
}

// MerchantAddressUpdateRequest represents the request to update a merchant address
type MerchantAddressUpdateRequest struct {
	ID    string  `json:"-"`     // Address ID (passed in URL path)
	Alias *string `json:"alias"` // Optional alias/label for the address
}

// MerchantAddressVerifyRequest represents the request to verify a merchant address
type MerchantAddressVerifyRequest struct {
	ID     string `json:"-"`      // Address ID (passed in URL path)
	Amount string `json:"amount"` // Amount to verify (small test transaction)
}

// MerchantAddressListRequest represents the request to list merchant addresses
type MerchantAddressListRequest struct {
	Network  *string `json:"network,omitempty" form:"network"`                           // Optional network filter
	Page     int32   `json:"page" form:"page"`                                           // Page number, starting from 0
	PageSize int32   `json:"page_size" binding:"required,min=1,max=50" form:"page_size"` // Number of items per page
}

// MerchantAddressListResponse represents the response when listing merchant addresses
type MerchantAddressListResponse struct {
	MerchantAddresses []*developer.MerchantAddress `json:"merchant_addresses"`
	Total             int64                        `json:"total"`
	Page              int32                        `json:"page"`
	PageSize          int32                        `json:"page_size"`
}

// CreateMerchantAddress creates a new merchant address
func (c *Client) CreateMerchantAddress(req MerchantAddressCreateRequest) (*developer.MerchantAddress, error) {
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
func (c *Client) UpdateMerchantAddress(req MerchantAddressUpdateRequest) (*developer.MerchantAddress, error) {
	var response developer.MerchantAddress
	err := c.sendRequest("POST", "/v1/addresses/"+req.ID, req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// VerifyMerchantAddress verifies a merchant address
func (c *Client) VerifyMerchantAddress(req MerchantAddressVerifyRequest) (*developer.MerchantAddress, error) {
	var response developer.MerchantAddress
	err := c.sendRequest("POST", "/v1/addresses/"+req.ID+"/verify", req, &response)
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
func (c *Client) ListMerchantAddresses(req MerchantAddressListRequest) (*MerchantAddressListResponse, error) {
	var response MerchantAddressListResponse
	err := c.sendRequestWithQuery("GET", "/v1/addresses", req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

package martianpay

import (
	"fmt"

	"github.com/MartianPay/martianpay-go-sample/pkg/developer"
)

// ListSellingPlanGroups lists selling plan groups
func (c *Client) ListSellingPlanGroups(params *developer.Pagination) (*developer.ListSellingPlanGroupsResponse, error) {
	var resp developer.ListSellingPlanGroupsResponse
	err := c.sendRequestWithQuery("GET", "/v1/selling_plan_groups", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateSellingPlanGroup creates a selling plan group
func (c *Client) CreateSellingPlanGroup(params *developer.CreateSellingPlanGroupRequest) (*developer.SellingPlanGroupResponse, error) {
	var resp developer.SellingPlanGroupResponse
	err := c.sendRequest("POST", "/v1/selling_plan_groups", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSellingPlanGroup gets a selling plan group
func (c *Client) GetSellingPlanGroup(groupID string) (*developer.SellingPlanGroupResponse, error) {
	path := fmt.Sprintf("/v1/selling_plan_groups/%s", groupID)
	var resp developer.SellingPlanGroupResponse
	err := c.sendRequest("GET", path, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateSellingPlanGroup updates a selling plan group
func (c *Client) UpdateSellingPlanGroup(groupID string, params *developer.UpdateSellingPlanGroupRequest) (*developer.SellingPlanGroupResponse, error) {
	path := fmt.Sprintf("/v1/selling_plan_groups/%s", groupID)
	var resp developer.SellingPlanGroupResponse
	err := c.sendRequest("POST", path, params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteSellingPlanGroup deletes a selling plan group
func (c *Client) DeleteSellingPlanGroup(groupID string) error {
	path := fmt.Sprintf("/v1/selling_plan_groups/%s", groupID)
	return c.sendRequest("DELETE", path, nil, nil)
}

// ListSellingPlans lists selling plans
func (c *Client) ListSellingPlans(params *developer.Pagination) (*developer.ListSellingPlansResponse, error) {
	var resp developer.ListSellingPlansResponse
	err := c.sendRequestWithQuery("GET", "/v1/selling_plans", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateSellingPlan creates a selling plan
func (c *Client) CreateSellingPlan(params *developer.CreateSellingPlanRequest) (*developer.SellingPlanResponse, error) {
	var resp developer.SellingPlanResponse
	err := c.sendRequest("POST", "/v1/selling_plans", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CalculateSellingPlanPrice calculates selling plan price
func (c *Client) CalculateSellingPlanPrice(params map[string]interface{}) (*developer.CalculatePriceResponse, error) {
	var resp developer.CalculatePriceResponse
	err := c.sendRequest("POST", "/v1/selling_plans/calculate_price", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSellingPlan gets a selling plan
func (c *Client) GetSellingPlan(planID string) (*developer.SellingPlanResponse, error) {
	path := fmt.Sprintf("/v1/selling_plans/%s", planID)
	var resp developer.SellingPlanResponse
	err := c.sendRequest("GET", path, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateSellingPlan updates a selling plan
func (c *Client) UpdateSellingPlan(planID string, params *developer.UpdateSellingPlanRequest) (*developer.SellingPlanResponse, error) {
	path := fmt.Sprintf("/v1/selling_plans/%s", planID)
	var resp developer.SellingPlanResponse
	err := c.sendRequest("POST", path, params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteSellingPlan deletes a selling plan
func (c *Client) DeleteSellingPlan(planID string) error {
	path := fmt.Sprintf("/v1/selling_plans/%s", planID)
	return c.sendRequest("DELETE", path, nil, nil)
}

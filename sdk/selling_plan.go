// Package martianpay provides SDK methods for managing selling plans and selling plan groups.
// Selling plans enable subscription and recurring payment functionality for products.
package martianpay

import (
	"fmt"

	"github.com/MartianPay/martianpay-go-sample/pkg/developer"
)

// ListSellingPlanGroups retrieves a paginated list of selling plan groups.
// Selling plan groups organize multiple selling plans (subscription options) together.
//
// Parameters:
//   - params: Pagination parameters including page number and page size
//
// Returns:
//   - *developer.ListSellingPlanGroupsResponse: List of selling plan groups with pagination metadata
//   - error: nil on success, error on failure
func (c *Client) ListSellingPlanGroups(params *developer.Pagination) (*developer.ListSellingPlanGroupsResponse, error) {
	var resp developer.ListSellingPlanGroupsResponse
	err := c.sendRequestWithQuery("GET", "/v1/selling_plan_groups", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateSellingPlanGroup creates a new selling plan group.
// A selling plan group contains multiple selling plans and can be associated with products.
//
// Parameters:
//   - params: Request containing group name, description, merchant code, and associated selling plans
//
// Returns:
//   - *developer.SellingPlanGroupResponse: The created selling plan group with assigned ID
//   - error: nil on success, error on failure
func (c *Client) CreateSellingPlanGroup(params *developer.CreateSellingPlanGroupRequest) (*developer.SellingPlanGroupResponse, error) {
	var resp developer.SellingPlanGroupResponse
	err := c.sendRequest("POST", "/v1/selling_plan_groups", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSellingPlanGroup retrieves detailed information about a specific selling plan group.
// Includes all selling plans within the group and product associations.
//
// Parameters:
//   - groupID: The unique identifier of the selling plan group
//
// Returns:
//   - *developer.SellingPlanGroupResponse: Complete selling plan group details
//   - error: nil on success, error on failure (e.g., group not found)
func (c *Client) GetSellingPlanGroup(groupID string) (*developer.SellingPlanGroupResponse, error) {
	path := fmt.Sprintf("/v1/selling_plan_groups/%s", groupID)
	var resp developer.SellingPlanGroupResponse
	err := c.sendRequest("GET", path, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateSellingPlanGroup updates an existing selling plan group's configuration.
// Can modify group name, description, and associated selling plans.
//
// Parameters:
//   - groupID: The unique identifier of the selling plan group to update
//   - params: Updated group fields including name, description, and selling plan associations
//
// Returns:
//   - *developer.SellingPlanGroupResponse: The updated selling plan group details
//   - error: nil on success, error on failure
func (c *Client) UpdateSellingPlanGroup(groupID string, params *developer.UpdateSellingPlanGroupRequest) (*developer.SellingPlanGroupResponse, error) {
	path := fmt.Sprintf("/v1/selling_plan_groups/%s", groupID)
	var resp developer.SellingPlanGroupResponse
	err := c.sendRequest("POST", path, params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteSellingPlanGroup permanently deletes a selling plan group.
// Can only delete groups that are not associated with any active products.
//
// Parameters:
//   - groupID: The unique identifier of the selling plan group to delete
//
// Returns:
//   - error: nil on success, error on failure (e.g., group has product associations)
func (c *Client) DeleteSellingPlanGroup(groupID string) error {
	path := fmt.Sprintf("/v1/selling_plan_groups/%s", groupID)
	return c.sendRequest("DELETE", path, nil, nil)
}

// ListSellingPlans retrieves a paginated list of selling plans.
// Selling plans define subscription billing intervals, pricing, and terms.
//
// Parameters:
//   - params: Pagination parameters including page number and page size
//
// Returns:
//   - *developer.ListSellingPlansResponse: List of selling plans with pagination metadata
//   - error: nil on success, error on failure
func (c *Client) ListSellingPlans(params *developer.Pagination) (*developer.ListSellingPlansResponse, error) {
	var resp developer.ListSellingPlansResponse
	err := c.sendRequestWithQuery("GET", "/v1/selling_plans", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateSellingPlan creates a new selling plan.
// Defines subscription billing terms including interval, pricing adjustments, and delivery frequency.
//
// Parameters:
//   - params: Request containing plan name, billing interval, pricing policy, and delivery settings
//
// Returns:
//   - *developer.SellingPlanResponse: The created selling plan with assigned ID
//   - error: nil on success, error on failure
func (c *Client) CreateSellingPlan(params *developer.CreateSellingPlanRequest) (*developer.SellingPlanResponse, error) {
	var resp developer.SellingPlanResponse
	err := c.sendRequest("POST", "/v1/selling_plans", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CalculateSellingPlanPrice calculates the price for a selling plan configuration.
// Useful for previewing subscription pricing before creating a subscription.
//
// Parameters:
//   - params: Map containing selling plan ID, quantity, and any applicable discounts
//
// Returns:
//   - *developer.CalculatePriceResponse: Calculated pricing details including total, subtotal, and discounts
//   - error: nil on success, error on failure
func (c *Client) CalculateSellingPlanPrice(params map[string]interface{}) (*developer.CalculatePriceResponse, error) {
	var resp developer.CalculatePriceResponse
	err := c.sendRequest("POST", "/v1/selling_plans/calculate_price", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSellingPlan retrieves detailed information about a specific selling plan.
// Includes billing interval, pricing adjustments, delivery frequency, and other terms.
//
// Parameters:
//   - planID: The unique identifier of the selling plan
//
// Returns:
//   - *developer.SellingPlanResponse: Complete selling plan details
//   - error: nil on success, error on failure (e.g., plan not found)
func (c *Client) GetSellingPlan(planID string) (*developer.SellingPlanResponse, error) {
	path := fmt.Sprintf("/v1/selling_plans/%s", planID)
	var resp developer.SellingPlanResponse
	err := c.sendRequest("GET", path, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateSellingPlan updates an existing selling plan's configuration.
// Can modify billing intervals, pricing policies, and delivery settings.
//
// Parameters:
//   - planID: The unique identifier of the selling plan to update
//   - params: Updated plan fields including name, billing interval, and pricing adjustments
//
// Returns:
//   - *developer.SellingPlanResponse: The updated selling plan details
//   - error: nil on success, error on failure
func (c *Client) UpdateSellingPlan(planID string, params *developer.UpdateSellingPlanRequest) (*developer.SellingPlanResponse, error) {
	path := fmt.Sprintf("/v1/selling_plans/%s", planID)
	var resp developer.SellingPlanResponse
	err := c.sendRequest("POST", path, params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteSellingPlan permanently deletes a selling plan.
// Can only delete plans that are not part of any active selling plan groups or subscriptions.
//
// Parameters:
//   - planID: The unique identifier of the selling plan to delete
//
// Returns:
//   - error: nil on success, error on failure (e.g., plan is in use)
func (c *Client) DeleteSellingPlan(planID string) error {
	path := fmt.Sprintf("/v1/selling_plans/%s", planID)
	return c.sendRequest("DELETE", path, nil, nil)
}

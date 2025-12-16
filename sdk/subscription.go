// Package martianpay provides SDK methods for managing subscriptions.
// Subscriptions enable recurring billing and automatic payment collection for customers.
package martianpay

import (
	"fmt"

	"github.com/MartianPay/martianpay-go-sample/pkg/developer"
)

// ListSubscriptions retrieves a paginated list of subscriptions.
// Subscriptions represent recurring billing arrangements for customers.
//
// Parameters:
//   - params: Query parameters including pagination, filters by customer, status, etc.
//
// Returns:
//   - *developer.ListSubscriptionsResponse: List of subscriptions with pagination metadata
//   - error: nil on success, error on failure
func (c *Client) ListSubscriptions(params *developer.ListMerchantSubscriptionsRequest) (*developer.ListSubscriptionsResponse, error) {
	var resp developer.ListSubscriptionsResponse
	err := c.sendRequestWithQuery("GET", "/v1/subscriptions", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSubscription retrieves detailed information about a specific subscription.
// Includes billing cycle, payment method, customer details, and subscription status.
//
// Parameters:
//   - subscriptionID: The unique identifier of the subscription
//
// Returns:
//   - *developer.SubscriptionDetails: Complete subscription details
//   - error: nil on success, error on failure (e.g., subscription not found)
func (c *Client) GetSubscription(subscriptionID string) (*developer.SubscriptionDetails, error) {
	path := fmt.Sprintf("/v1/subscriptions/%s", subscriptionID)
	var resp developer.SubscriptionDetails
	err := c.sendRequest("GET", path, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CancelSubscription cancels a subscription either immediately or at the end of the billing period.
// Cancelled subscriptions cannot be resumed.
//
// Parameters:
//   - subscriptionID: The unique identifier of the subscription to cancel
//   - params: Cancellation parameters (e.g., cancel_at_period_end flag)
//
// Returns:
//   - *developer.SubscriptionDetails: Updated subscription details with cancelled status
//   - error: nil on success, error on failure
func (c *Client) CancelSubscription(subscriptionID string, params *developer.CancelMerchantSubscriptionRequest) (*developer.SubscriptionDetails, error) {
	path := fmt.Sprintf("/v1/subscriptions/%s/cancel", subscriptionID)
	var resp developer.SubscriptionDetails
	err := c.sendRequest("POST", path, params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// PauseSubscription temporarily pauses a subscription.
// Paused subscriptions can be resumed later without creating a new subscription.
//
// Parameters:
//   - subscriptionID: The unique identifier of the subscription to pause
//   - params: Pause parameters (e.g., auto_resume_at timestamp)
//
// Returns:
//   - *developer.SubscriptionDetails: Updated subscription details with paused status
//   - error: nil on success, error on failure
func (c *Client) PauseSubscription(subscriptionID string, params *developer.PauseMerchantSubscriptionRequest) (*developer.SubscriptionDetails, error) {
	path := fmt.Sprintf("/v1/subscriptions/%s/pause", subscriptionID)
	var resp developer.SubscriptionDetails
	err := c.sendRequest("POST", path, params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ResumeSubscription resumes a previously paused subscription.
// Billing will continue from the next scheduled billing date.
//
// Parameters:
//   - subscriptionID: The unique identifier of the subscription to resume
//
// Returns:
//   - *developer.SubscriptionDetails: Updated subscription details with active status
//   - error: nil on success, error on failure (e.g., subscription not paused)
func (c *Client) ResumeSubscription(subscriptionID string) (*developer.SubscriptionDetails, error) {
	path := fmt.Sprintf("/v1/subscriptions/%s/resume", subscriptionID)
	var resp developer.SubscriptionDetails
	err := c.sendRequest("POST", path, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

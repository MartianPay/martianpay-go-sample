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

// UpdateSubscription updates a subscription's plan (upgrade or downgrade).
// Supports automatic proration calculation for mid-cycle plan changes.
//
// For upgrades (new price > current price):
//   - Applied immediately
//   - Customer charged prorated difference today
//   - Billing cycle can be reset with billing_cycle_anchor: "now"
//
// For downgrades (new price < current price):
//   - Scheduled as pending update
//   - Takes effect at end of current billing period
//   - Customer continues on current plan until effective_date
//
// Parameters:
//   - subscriptionID: The unique identifier of the subscription to update
//   - params: Update parameters including new selling plan, proration behavior, etc.
//
// Returns:
//   - *developer.SubscriptionDetails: Updated subscription details with proration info
//   - error: nil on success, error on failure
func (c *Client) UpdateSubscription(subscriptionID string, params *developer.UpdateSubscriptionPlanRequest) (*developer.SubscriptionDetails, error) {
	path := fmt.Sprintf("/v1/subscriptions/%s", subscriptionID)
	var resp developer.SubscriptionDetails
	err := c.sendRequest("POST", path, params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// PreviewSubscriptionUpdate previews the proration calculation for a plan change.
// Use this to show customers what they'll be charged before confirming the change.
//
// The preview does NOT:
//   - Modify the subscription
//   - Create invoices
//   - Charge the customer
//   - Send any webhooks
//
// Parameters:
//   - subscriptionID: The unique identifier of the subscription
//   - params: Preview parameters (same as UpdateSubscription)
//
// Returns:
//   - *developer.SubscriptionDetails: Preview with proration calculation (applied=false)
//   - error: nil on success, error on failure
func (c *Client) PreviewSubscriptionUpdate(subscriptionID string, params *developer.UpdateSubscriptionPlanRequest) (*developer.SubscriptionDetails, error) {
	path := fmt.Sprintf("/v1/subscriptions/%s/preview", subscriptionID)
	var resp developer.SubscriptionDetails
	err := c.sendRequest("POST", path, params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// RevokeCancelSubscription revokes a pending cancellation for a subscription.
// This is used when a subscription was previously scheduled to cancel (cancel_at_period_end=true)
// but the customer or merchant wants to continue the subscription.
//
// Parameters:
//   - subscriptionID: The unique identifier of the subscription to revoke cancellation
//
// Returns:
//   - *developer.SubscriptionDetails: Updated subscription details with cancellation revoked
//   - error: nil on success, error on failure (e.g., subscription not pending cancellation)
func (c *Client) RevokeCancelSubscription(subscriptionID string) (*developer.SubscriptionDetails, error) {
	path := fmt.Sprintf("/v1/subscriptions/%s/revoke-cancel", subscriptionID)
	var resp developer.SubscriptionDetails
	err := c.sendRequest("POST", path, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

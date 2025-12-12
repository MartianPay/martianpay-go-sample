package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/MartianPay/martianpay-go-sample/pkg/developer"
	martianpay "github.com/MartianPay/martianpay-go-sample/sdk"
)

// Subscription Examples

// listSubscriptions lists all subscriptions
func listSubscriptions(client *martianpay.Client) {
	fmt.Println("Listing Subscriptions...")

	req := &developer.ListMerchantSubscriptionsRequest{
		Offset: 0,
		Limit:  10,
	}

	response, err := client.ListSubscriptions(req)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("✓ Total Subscriptions: %d\n", response.Total)
	fmt.Printf("  Offset: %d, Limit: %d\n\n", response.Offset, response.Limit)

	if len(response.Data) == 0 {
		fmt.Println("  No subscriptions found")
		return
	}

	for i, sub := range response.Data {
		fmt.Printf("[%d] ID: %s\n", i+1, sub.ID)
		fmt.Printf("    Customer ID: %s\n", sub.CustomerID)
		if sub.CustomerName != nil {
			fmt.Printf("    Customer: %s", *sub.CustomerName)
			if sub.CustomerEmail != nil {
				fmt.Printf(" (%s)", *sub.CustomerEmail)
			}
			fmt.Println()
		}
		if sub.ProductName != nil {
			fmt.Printf("    Product: %s\n", *sub.ProductName)
		}
		fmt.Printf("    Status: %s\n", sub.Status)
		if sub.CurrentPricingTier != nil {
			fmt.Printf("    Current Price: %s %s\n",
				sub.CurrentPricingTier.FinalPrice,
				sub.CurrentPricingTier.Currency)
		}
		if sub.NextChargeAmount != nil && *sub.NextChargeAmount != "" {
			fmt.Printf("    Next Charge: %s\n", *sub.NextChargeAmountDisplay)
		}
		fmt.Printf("    Current Period: %s - %s\n",
			formatTimestamp(sub.CurrentPeriodStart),
			formatTimestamp(sub.CurrentPeriodEnd))
		if sub.CancelAtPeriodEnd {
			fmt.Printf("    ⚠️  Will cancel at period end\n")
		}
		if sub.PausedAt != nil {
			fmt.Printf("    ⏸️  Paused since %s\n", formatTimestamp(*sub.PausedAt))
		}
		fmt.Println()
	}
}

// listSubscriptionsByCustomer lists subscriptions for a specific customer
func listSubscriptionsByCustomer(client *martianpay.Client) {
	fmt.Println("Listing Subscriptions by Customer...")
	fmt.Print("Enter Customer ID: ")

	var customerID string
	fmt.Scanln(&customerID)
	if customerID == "" {
		fmt.Println("✗ Customer ID is required")
		return
	}

	req := &developer.ListMerchantSubscriptionsRequest{
		CustomerID: &customerID,
		Offset:     0,
		Limit:      10,
	}

	response, err := client.ListSubscriptions(req)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("\n✓ Subscriptions for Customer %s: %d\n\n", customerID, response.Total)

	if len(response.Data) == 0 {
		fmt.Println("  No subscriptions found for this customer")
		return
	}

	for i, sub := range response.Data {
		fmt.Printf("[%d] %s\n", i+1, sub.ID)
		if sub.ProductName != nil {
			fmt.Printf("    Product: %s\n", *sub.ProductName)
		}
		if sub.VariantTitle != nil {
			fmt.Printf("    Variant: %s\n", *sub.VariantTitle)
		}
		fmt.Printf("    Status: %s\n", sub.Status)
		if sub.CurrentPricingTier != nil {
			fmt.Printf("    Price: %s %s / %s\n",
				sub.CurrentPricingTier.FinalPrice,
				sub.CurrentPricingTier.Currency,
				sub.CurrentPricingTier.BillingCycle)
		}
		fmt.Println()
	}
}

// listSubscriptionsByStatus lists subscriptions filtered by status
func listSubscriptionsByStatus(client *martianpay.Client) {
	fmt.Println("Listing Subscriptions by Status...")
	fmt.Println("\nAvailable statuses:")
	fmt.Println("  1. incomplete - Initial state, awaiting first payment")
	fmt.Println("  2. active - Active and billing normally")
	fmt.Println("  3. paused - Temporarily paused")
	fmt.Println("  4. past_due - Payment failed")
	fmt.Println("  5. canceled - Canceled")

	fmt.Print("\nSelect status (1-5): ")
	var choice string
	fmt.Scanln(&choice)

	var status string
	switch choice {
	case "1":
		status = "incomplete"
	case "2":
		status = "active"
	case "3":
		status = "paused"
	case "4":
		status = "past_due"
	case "5":
		status = "canceled"
	default:
		status = "active"
	}

	req := &developer.ListMerchantSubscriptionsRequest{
		Status: &status,
		Offset: 0,
		Limit:  10,
	}

	response, err := client.ListSubscriptions(req)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("\n✓ %s Subscriptions: %d\n\n", strings.ToUpper(status), response.Total)

	if len(response.Data) == 0 {
		fmt.Printf("  No %s subscriptions found\n", status)
		return
	}

	for i, sub := range response.Data {
		fmt.Printf("[%d] ID: %s\n", i+1, sub.ID)
		fmt.Printf("    Customer ID: %s\n", sub.CustomerID)
		if sub.CustomerName != nil {
			fmt.Printf("    Customer: %s\n", *sub.CustomerName)
		}
		if sub.ProductName != nil {
			fmt.Printf("    Product: %s\n", *sub.ProductName)
		}

		// Show status-specific information
		switch status {
		case "incomplete":
			if sub.PaymentRequired != nil && *sub.PaymentRequired {
				fmt.Printf("    ⚠️  Payment Required\n")
				if sub.PaymentURL != nil {
					fmt.Printf("    Payment URL: %s\n", *sub.PaymentURL)
				}
				if sub.HoursSinceCreation != nil {
					fmt.Printf("    Hours Since Creation: %.1f\n", *sub.HoursSinceCreation)
				}
			}
		case "paused":
			if sub.PausedAt != nil {
				fmt.Printf("    Paused: %s\n", formatTimestamp(*sub.PausedAt))
			}
			if sub.ResumesAt != nil {
				fmt.Printf("    Auto-Resume: %s\n", formatTimestamp(*sub.ResumesAt))
			}
		case "canceled":
			if sub.CanceledAt != nil {
				fmt.Printf("    Canceled: %s\n", formatTimestamp(*sub.CanceledAt))
			}
			if sub.CancelReason != nil {
				fmt.Printf("    Reason: %s\n", *sub.CancelReason)
			}
		}

		fmt.Println()
	}
}

// getSubscription retrieves details of a specific subscription
func getSubscription(client *martianpay.Client) {
	fmt.Println("Getting Subscription...")
	fmt.Print("Enter Subscription ID: ")

	var id string
	fmt.Scanln(&id)
	if id == "" {
		id = "sub_example_id"
	}

	response, err := client.GetSubscription(id)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("\n✓ Subscription Details:\n")
	fmt.Printf("  ID: %s\n", response.ID)
	fmt.Printf("  Status: %s\n", response.Status)

	// Customer Information
	fmt.Printf("\n  Customer:\n")
	if response.CustomerName != nil {
		fmt.Printf("    Name: %s\n", *response.CustomerName)
	}
	if response.CustomerEmail != nil {
		fmt.Printf("    Email: %s\n", *response.CustomerEmail)
	}
	fmt.Printf("    ID: %s\n", response.CustomerID)

	// Product Information
	if response.ProductName != nil {
		fmt.Printf("\n  Product:\n")
		fmt.Printf("    Name: %s\n", *response.ProductName)
		if response.ProductDescription != nil {
			fmt.Printf("    Description: %s\n", *response.ProductDescription)
		}
		if response.VariantTitle != nil {
			fmt.Printf("    Variant: %s\n", *response.VariantTitle)
		}
		if response.VariantOptionValues != nil && len(response.VariantOptionValues) > 0 {
			fmt.Printf("    Options: ")
			first := true
			for k, v := range response.VariantOptionValues {
				if !first {
					fmt.Printf(", ")
				}
				fmt.Printf("%s: %s", k, v)
				first = false
			}
			fmt.Println()
		}
	}

	// Selling Plan Information
	if response.SellingPlanName != nil {
		fmt.Printf("\n  Selling Plan:\n")
		fmt.Printf("    Name: %s\n", *response.SellingPlanName)
		if response.SellingPlanDescription != nil {
			fmt.Printf("    Description: %s\n", *response.SellingPlanDescription)
		}
	}

	// Pricing Information
	if response.CurrentPricingTier != nil {
		fmt.Printf("\n  Current Pricing:\n")
		fmt.Printf("    Cycle: %d\n", response.CurrentPricingTier.CycleNumber)
		fmt.Printf("    Base Price: %s %s\n",
			response.CurrentPricingTier.BasePrice,
			response.CurrentPricingTier.Currency)
		if response.CurrentPricingTier.DiscountPercentage != nil {
			fmt.Printf("    Discount: %s%%\n", *response.CurrentPricingTier.DiscountPercentage)
		}
		fmt.Printf("    Discount Amount: %s\n", response.CurrentPricingTier.SellingPlanDiscount)
		fmt.Printf("    Final Price: %s %s\n",
			response.CurrentPricingTier.FinalPrice,
			response.CurrentPricingTier.Currency)
		fmt.Printf("    Billing Frequency: %s\n", response.CurrentPricingTier.BillingCycle)
	}

	if response.UpcomingPricingTier != nil {
		fmt.Printf("\n  Upcoming Pricing (Next Cycle):\n")
		fmt.Printf("    Cycle: %d\n", response.UpcomingPricingTier.CycleNumber)
		fmt.Printf("    Final Price: %s %s\n",
			response.UpcomingPricingTier.FinalPrice,
			response.UpcomingPricingTier.Currency)
	}

	// Billing Information
	fmt.Printf("\n  Billing:\n")
	fmt.Printf("    Current Period: %s - %s\n",
		formatTimestamp(response.CurrentPeriodStart),
		formatTimestamp(response.CurrentPeriodEnd))
	if response.NextChargeAmount != nil {
		fmt.Printf("    Next Charge: %s\n", *response.NextChargeAmountDisplay)
	}
	fmt.Printf("    Billing Cycle Anchor: %s\n", formatTimestamp(response.BillingCycleAnchor))

	// Trial Information
	if response.TrialStart != nil && response.TrialEnd != nil {
		fmt.Printf("\n  Trial Period:\n")
		fmt.Printf("    Start: %s\n", formatTimestamp(*response.TrialStart))
		fmt.Printf("    End: %s\n", formatTimestamp(*response.TrialEnd))
	}

	// Payment Method
	if response.DefaultPaymentMethodID != nil {
		fmt.Printf("\n  Payment Method:\n")
		if response.PaymentMethodBrand != nil && response.PaymentMethodLast4 != nil {
			fmt.Printf("    Card: %s ending in %s\n",
				strings.ToUpper(*response.PaymentMethodBrand),
				*response.PaymentMethodLast4)
		}
		if response.DefaultPaymentMethodType != nil {
			fmt.Printf("    Type: %s\n", *response.DefaultPaymentMethodType)
		}
	}

	// Cancellation Information
	if response.CancelAtPeriodEnd {
		fmt.Printf("\n  ⚠️  Cancellation:\n")
		fmt.Printf("    Will cancel at: %s\n", formatTimestamp(response.CurrentPeriodEnd))
		if response.CancelReason != nil {
			fmt.Printf("    Reason: %s\n", *response.CancelReason)
		}
	} else if response.CanceledAt != nil {
		fmt.Printf("\n  Cancellation:\n")
		fmt.Printf("    Canceled at: %s\n", formatTimestamp(*response.CanceledAt))
		if response.CancelReason != nil {
			fmt.Printf("    Reason: %s\n", *response.CancelReason)
		}
	}

	// Pause Information
	if response.PausedAt != nil {
		fmt.Printf("\n  Pause:\n")
		fmt.Printf("    Paused at: %s\n", formatTimestamp(*response.PausedAt))
		if response.PauseCollectionBehavior != nil {
			fmt.Printf("    Behavior: %s\n", *response.PauseCollectionBehavior)
		}
		if response.ResumesAt != nil {
			fmt.Printf("    Auto-Resume: %s\n", formatTimestamp(*response.ResumesAt))
		}
	}

	// Timestamps
	fmt.Printf("\n  Timestamps:\n")
	fmt.Printf("    Created: %s\n", formatTimestamp(response.CreatedAt))
	fmt.Printf("    Updated: %s\n", formatTimestamp(response.UpdatedAt))
}

// cancelSubscriptionAtPeriodEnd cancels a subscription at the end of the current period
func cancelSubscriptionAtPeriodEnd(client *martianpay.Client) {
	fmt.Println("Canceling Subscription at Period End...")
	fmt.Print("Enter Subscription ID: ")

	var id string
	fmt.Scanln(&id)
	if id == "" {
		fmt.Println("✗ Subscription ID is required")
		return
	}

	fmt.Print("Enter cancellation reason (optional): ")
	var reason string
	fmt.Scanln(&reason)
	if reason == "" {
		reason = "Customer requested cancellation"
	}

	cancelAtPeriodEnd := true
	req := &developer.CancelMerchantSubscriptionRequest{
		CancelAtPeriodEnd: &cancelAtPeriodEnd,
		CancelReason:      &reason,
	}

	response, err := client.CancelSubscription(id, req)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("\n✓ Subscription Scheduled for Cancellation:\n")
	fmt.Printf("  ID: %s\n", response.ID)
	fmt.Printf("  Status: %s (remains active until period end)\n", response.Status)
	fmt.Printf("  Will cancel at: %s\n", formatTimestamp(response.CurrentPeriodEnd))
	fmt.Printf("  Reason: %s\n", reason)
	fmt.Printf("\n  Note: Customer will retain access until %s\n",
		formatTimestamp(response.CurrentPeriodEnd))
}

// cancelSubscriptionImmediately cancels a subscription immediately
func cancelSubscriptionImmediately(client *martianpay.Client) {
	fmt.Println("Canceling Subscription Immediately...")
	fmt.Print("Enter Subscription ID: ")

	var id string
	fmt.Scanln(&id)
	if id == "" {
		fmt.Println("✗ Subscription ID is required")
		return
	}

	fmt.Print("Enter cancellation reason (optional): ")
	var reason string
	fmt.Scanln(&reason)
	if reason == "" {
		reason = "Immediate cancellation requested"
	}

	fmt.Print("\n⚠️  This will cancel the subscription immediately. Continue? (yes/no): ")
	var confirm string
	fmt.Scanln(&confirm)
	if strings.ToLower(confirm) != "yes" {
		fmt.Println("  Cancelled")
		return
	}

	cancelAtPeriodEnd := false
	req := &developer.CancelMerchantSubscriptionRequest{
		CancelAtPeriodEnd: &cancelAtPeriodEnd,
		CancelReason:      &reason,
	}

	response, err := client.CancelSubscription(id, req)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("\n✓ Subscription Canceled Immediately:\n")
	fmt.Printf("  ID: %s\n", response.ID)
	fmt.Printf("  Status: %s\n", response.Status)
	if response.CanceledAt != nil {
		fmt.Printf("  Canceled at: %s\n", formatTimestamp(*response.CanceledAt))
	}
	fmt.Printf("  Reason: %s\n", reason)
	fmt.Printf("\n  Note: Customer access has been revoked\n")
}

// pauseSubscription pauses a subscription indefinitely
func pauseSubscription(client *martianpay.Client) {
	fmt.Println("Pausing Subscription...")
	fmt.Print("Enter Subscription ID: ")

	var id string
	fmt.Scanln(&id)
	if id == "" {
		fmt.Println("✗ Subscription ID is required")
		return
	}

	fmt.Println("\nPause behaviors:")
	fmt.Println("  1. void - Cancel pending invoices (recommended)")
	fmt.Println("  2. keep_as_draft - Keep invoices as draft")
	fmt.Print("\nSelect behavior (1-2, default: 1): ")

	var behaviorChoice string
	fmt.Scanln(&behaviorChoice)

	behavior := developer.PauseCollectionBehaviorVoid
	if behaviorChoice == "2" {
		behavior = developer.PauseCollectionBehaviorKeepAsDraft
	}

	req := &developer.PauseMerchantSubscriptionRequest{
		Behavior: &behavior,
	}

	response, err := client.PauseSubscription(id, req)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("\n✓ Subscription Paused:\n")
	fmt.Printf("  ID: %s\n", response.ID)
	fmt.Printf("  Status: %s\n", response.Status)
	if response.PausedAt != nil {
		fmt.Printf("  Paused at: %s\n", formatTimestamp(*response.PausedAt))
	}
	if response.PauseCollectionBehavior != nil {
		fmt.Printf("  Behavior: %s\n", *response.PauseCollectionBehavior)
	}
	fmt.Printf("\n  Note: Subscription will remain paused until manually resumed\n")
}

// pauseSubscriptionWithAutoResume pauses a subscription with automatic resumption
func pauseSubscriptionWithAutoResume(client *martianpay.Client) {
	fmt.Println("Pausing Subscription with Auto-Resume...")
	fmt.Print("Enter Subscription ID: ")

	var id string
	fmt.Scanln(&id)
	if id == "" {
		fmt.Println("✗ Subscription ID is required")
		return
	}

	fmt.Print("Enter number of days to pause (e.g., 30): ")
	var days string
	fmt.Scanln(&days)
	if days == "" {
		days = "30"
	}

	var daysInt int
	fmt.Sscanf(days, "%d", &daysInt)
	if daysInt <= 0 {
		daysInt = 30
	}

	resumesAt := time.Now().AddDate(0, 0, daysInt).Unix()

	behavior := developer.PauseCollectionBehaviorVoid
	req := &developer.PauseMerchantSubscriptionRequest{
		Behavior:  &behavior,
		ResumesAt: &resumesAt,
	}

	response, err := client.PauseSubscription(id, req)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("\n✓ Subscription Paused with Auto-Resume:\n")
	fmt.Printf("  ID: %s\n", response.ID)
	fmt.Printf("  Status: %s\n", response.Status)
	if response.PausedAt != nil {
		fmt.Printf("  Paused at: %s\n", formatTimestamp(*response.PausedAt))
	}
	if response.ResumesAt != nil {
		fmt.Printf("  Will resume at: %s\n", formatTimestamp(*response.ResumesAt))
		fmt.Printf("  Days paused: %d\n", daysInt)
	}
}

// resumeSubscription resumes a paused subscription
func resumeSubscription(client *martianpay.Client) {
	fmt.Println("Resuming Subscription...")
	fmt.Print("Enter Subscription ID: ")

	var id string
	fmt.Scanln(&id)
	if id == "" {
		fmt.Println("✗ Subscription ID is required")
		return
	}

	response, err := client.ResumeSubscription(id)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("\n✓ Subscription Resumed:\n")
	fmt.Printf("  ID: %s\n", response.ID)
	fmt.Printf("  Status: %s\n", response.Status)
	fmt.Printf("  New Period: %s - %s\n",
		formatTimestamp(response.CurrentPeriodStart),
		formatTimestamp(response.CurrentPeriodEnd))
	if response.NextChargeAmount != nil {
		fmt.Printf("  Next Charge: %s\n", *response.NextChargeAmountDisplay)
	}
	fmt.Printf("\n  Note: Billing has resumed according to the original schedule\n")
}

// formatTimestamp converts Unix timestamp to human-readable format
func formatTimestamp(timestamp int64) string {
	t := time.Unix(timestamp, 0)
	return t.Format("2006-01-02 15:04:05")
}

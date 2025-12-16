// Package main provides examples for the MartianPay Customer API.
// Customers represent buyers in your system and can have saved payment methods,
// making it easy to process repeat purchases.
package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/MartianPay/martianpay-go-sample/pkg/developer"
	martianpay "github.com/MartianPay/martianpay-go-sample/sdk"
)

// generateRandomEmail generates a unique random email address for testing.
// This is useful for creating test customers without email collisions.
//
// Parameters:
//   - prefix: A prefix string to identify the customer type (e.g., "customer", "test_user")
//
// Returns:
//   - A unique email address in the format: prefix_timestamp_randomNumber@example.com
//
// Example: "customer_1234567890_456789@example.com"
func generateRandomEmail(prefix string) string {
	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Intn(1000000)
	timestamp := time.Now().Unix()
	return fmt.Sprintf("%s_%d_%d@example.com", prefix, timestamp, randomNum)
}

// Customer Examples

// createAndUpdateCustomer demonstrates creating a new customer and then updating their information.
// This is a common workflow when managing customer data.
//
// Steps:
// 1. Create a new customer with email, name, and description
// 2. Display the created customer details
// 3. Update the customer's name
// 4. Display the updated customer details
//
// API Endpoints Used:
//   - POST /v1/customers (create)
//   - PUT /v1/customers/:id (update)
func createAndUpdateCustomer(client *martianpay.Client) {
	fmt.Println("Creating and Updating Customer...")

	email := generateRandomEmail("customer")
	name := "John Doe"
	description := "Test customer"

	createReq := &developer.CustomerCreateRequest{
		CustomerParams: developer.CustomerParams{
			Email:       &email,
			Name:        &name,
			Description: &description,
		},
	}

	createResp, err := client.CreateCustomer(createReq)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("✓ Customer Created:\n")
	fmt.Printf("  ID: %s\n", createResp.ID)
	fmt.Printf("  Email: %s\n", *createResp.Email)
	fmt.Printf("  Name: %s\n\n", *createResp.Name)

	newName := "John Updated"
	updateReq := &developer.CustomerUpdateRequest{
		CustomerParams: developer.CustomerParams{
			Name: &newName,
		},
	}

	updateResp, err := client.UpdateCustomer(createResp.ID, updateReq)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("✓ Customer Updated:\n")
	fmt.Printf("  ID: %s\n", updateResp.ID)
	fmt.Printf("  Name: %s\n", *updateResp.Name)
}

// getCustomer retrieves and displays details of a specific customer by ID.
// The function first lists available customers to help with ID selection.
//
// Steps:
// 1. List recent customers for reference
// 2. Prompt user to select a customer (by number or ID)
// 3. Retrieve and display customer details including balance and statistics
//
// API Endpoints Used:
//   - GET /v1/customers (list)
//   - GET /v1/customers/:id (get)
func getCustomer(client *martianpay.Client) {
	fmt.Println("Getting Customer...")
	fmt.Println("  Fetching customers...")

	// List customers first
	listReq := &developer.CustomerListRequest{
		Pagination: developer.Pagination{
			Page:     0,
			PageSize: 10,
		},
	}
	listResp, err := client.ListCustomers(listReq)
	if err == nil && len(listResp.Customers) > 0 {
		fmt.Printf("\n  Available Customers:\n")
		for i, cust := range listResp.Customers {
			fmt.Printf("  [%d] ID: %s", i+1, cust.ID)
			if cust.Email != nil {
				fmt.Printf(" - %s", *cust.Email)
			}
			if cust.Name != nil {
				fmt.Printf(" (%s)", *cust.Name)
			}
			fmt.Println()
		}
		fmt.Print("\nEnter customer number or ID: ")
	} else {
		fmt.Print("\nEnter Customer ID: ")
	}

	var choice string
	fmt.Scanln(&choice)

	var id string
	if choice != "" && listResp != nil && len(listResp.Customers) > 0 {
		// Try to find by ID first
		foundByID := false
		for _, cust := range listResp.Customers {
			if cust.ID == choice {
				id = choice
				foundByID = true
				break
			}
		}
		// If not found by ID, try as number
		if !foundByID {
			var idx int
			fmt.Sscanf(choice, "%d", &idx)
			if idx > 0 && idx <= len(listResp.Customers) {
				id = listResp.Customers[idx-1].ID
			}
		}
		if id == "" {
			id = choice
		}
	} else if choice != "" {
		id = choice
	} else {
		id = "cus_example_id"
	}

	response, err := client.GetCustomer(id)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("✓ Customer Retrieved:\n")
	fmt.Printf("  ID: %s\n", response.ID)
	if response.Email != nil {
		fmt.Printf("  Email: %s\n", *response.Email)
	}
	if response.Name != nil {
		fmt.Printf("  Name: %s\n", *response.Name)
	}
	fmt.Printf("  Currency: %s\n", response.Currency)
	fmt.Printf("  Total Payment: %.2f\n", response.TotalPayment)
}

// listCustomers retrieves and displays a paginated list of customers.
// Useful for viewing all customers in your merchant account.
//
// Features:
//   - Pagination support (page and page size)
//   - Displays customer ID, email, and name
//   - Shows total count of customers
//
// API Endpoints Used:
//   - GET /v1/customers (list with pagination)
func listCustomers(client *martianpay.Client) {
	fmt.Println("Listing Customers...")

	req := &developer.CustomerListRequest{
		Pagination: developer.Pagination{
			Page:     0,
			PageSize: 10,
		},
	}

	response, err := client.ListCustomers(req)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("✓ Total Customers: %d\n", response.Total)
	fmt.Printf("  Page: %d, Page Size: %d\n\n", response.Page, response.PageSize)

	for i, customer := range response.Customers {
		fmt.Printf("  [%d] ID: %s", i+1, customer.ID)
		if customer.Email != nil {
			fmt.Printf(", Email: %s", *customer.Email)
		}
		if customer.Name != nil {
			fmt.Printf(", Name: %s", *customer.Name)
		}
		fmt.Println()
	}
}

// deleteCustomer demonstrates deleting a customer from the system.
// For demonstration purposes, this function creates a test customer first,
// then immediately deletes it.
//
// Note: In production, only delete customers when necessary (e.g., GDPR compliance).
// Deleted customers cannot be recovered.
//
// Steps:
// 1. Create a test customer
// 2. Delete the newly created customer
// 3. Confirm successful deletion
//
// API Endpoints Used:
//   - POST /v1/customers (create test customer)
//   - DELETE /v1/customers/:id (delete)
func deleteCustomer(client *martianpay.Client) {
	fmt.Println("Deleting Customer...")

	email := generateRandomEmail("delete_test")
	name := "Delete Test Customer"
	createReq := &developer.CustomerCreateRequest{
		CustomerParams: developer.CustomerParams{
			Email: &email,
			Name:  &name,
		},
	}

	createResp, err := client.CreateCustomer(createReq)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("  Created customer: %s\n", createResp.ID)

	err = client.DeleteCustomer(createResp.ID)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("✓ Customer Deleted: %s\n", createResp.ID)
}

// listCustomerPaymentMethods retrieves all saved payment methods for a customer.
// Saved payment methods allow for faster checkout on repeat purchases.
//
// Payment methods include:
//   - Credit/debit cards (Stripe)
//   - Bank accounts
//   - Other payment methods supported by your configuration
//
// Displayed Information:
//   - Payment method ID
//   - Card brand (Visa, Mastercard, etc.)
//   - Last 4 digits of card
//   - Expiration date
//
// API Endpoints Used:
//   - GET /v1/customers/:id/payment_methods
func listCustomerPaymentMethods(client *martianpay.Client) {
	fmt.Println("Listing Customer Payment Methods...")
	fmt.Print("Enter Customer ID: ")

	var customerID string
	fmt.Scanln(&customerID)
	if customerID == "" {
		customerID = "cus_example_id"
	}

	resp, err := client.ListCustomerPaymentMethods(customerID)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("✓ Found %d payment method(s)\n\n", len(resp.PaymentMethods))

	for i, pm := range resp.PaymentMethods {
		fmt.Printf("  [%d] ID: %s\n", i+1, pm.ID)
		fmt.Printf("      Brand: %s, Last4: %s\n", pm.Brand, pm.Last4)
		fmt.Printf("      Expires: %d/%d\n", pm.ExpMonth, pm.ExpYear)
	}
}

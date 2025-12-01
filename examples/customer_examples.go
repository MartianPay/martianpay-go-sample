package main

import (
	"fmt"

	"github.com/MartianPay/martianpay-go-sample/pkg/developer"
	martianpay "github.com/MartianPay/martianpay-go-sample/sdk"
)

// Customer Examples
func createAndUpdateCustomer(client *martianpay.Client) {
	fmt.Println("Creating and Updating Customer...")

	email := "newcustomer@example.com"
	name := "John Doe"
	description := "Test customer"

	createReq := martianpay.CustomerCreateRequest{
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
	updateReq := martianpay.CustomerUpdateRequest{
		ID: createResp.ID,
		CustomerParams: developer.CustomerParams{
			Name: &newName,
		},
	}

	updateResp, err := client.UpdateCustomer(updateReq)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("✓ Customer Updated:\n")
	fmt.Printf("  ID: %s\n", updateResp.ID)
	fmt.Printf("  Name: %s\n", *updateResp.Name)
}

func getCustomer(client *martianpay.Client) {
	fmt.Println("Getting Customer...")
	fmt.Print("Enter Customer ID: ")

	var id string
	fmt.Scanln(&id)
	if id == "" {
		id = "cus_example_id"
	}

	response, err := client.GetCustomer(martianpay.CustomerGetRequest{ID: id})
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

func listCustomers(client *martianpay.Client) {
	fmt.Println("Listing Customers...")

	req := martianpay.CustomerListRequest{
		Page:     0,
		PageSize: 10,
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

func deleteCustomer(client *martianpay.Client) {
	fmt.Println("Deleting Customer...")

	email := "delete_test@example.com"
	name := "Delete Test Customer"
	createReq := martianpay.CustomerCreateRequest{
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

	_, err = client.DeleteCustomer(martianpay.CustomerDeleteRequest{ID: createResp.ID})
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("✓ Customer Deleted: %s\n", createResp.ID)
}

func listCustomerPaymentMethods(client *martianpay.Client) {
	fmt.Println("Listing Customer Payment Methods...")
	fmt.Print("Enter Customer ID: ")

	var customerID string
	fmt.Scanln(&customerID)
	if customerID == "" {
		customerID = "cus_example_id"
	}

	req := martianpay.CustomerPaymentMethodListRequest{
		CustomerID: customerID,
	}

	resp, err := client.ListCustomerPaymentMethods(req)
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

// Package main provides examples for the MartianPay Product API.
// Products represent items or services merchants sell, with support for variants,
// options (like size/color), pricing, and inventory management.
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/MartianPay/martianpay-go-sample/pkg/developer"
	martianpay "github.com/MartianPay/martianpay-go-sample/sdk"
	"github.com/shopspring/decimal"
)

// Product Examples

// listProducts lists all products with pagination
func listProducts(client *martianpay.Client) {
	fmt.Println("Listing Products...")

	req := &developer.ProductListRequest{
		Page:     0,
		PageSize: 10,
	}

	response, err := client.ListProducts(req)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("✓ Total Products: %d\n", response.Total)
	fmt.Printf("  Page: %d, Page Size: %d\n\n", response.Page, response.PageSize)

	if len(response.Products) == 0 {
		fmt.Println("  No products found")
		return
	}

	for i, product := range response.Products {
		fmt.Printf("[%d] %s\n", i+1, product.Name)
		fmt.Printf("    ID: %s\n", product.ID)
		fmt.Printf("    Active: %v\n", product.Active)
		if product.Description != "" {
			desc := product.Description
			if len(desc) > 60 {
				desc = desc[:60] + "..."
			}
			fmt.Printf("    Description: %s\n", desc)
		}
		if product.Price != nil {
			fmt.Printf("    Price: %s %s\n", product.Price.Amount, product.Price.AssetId)
		}
		if len(product.Variants) > 0 {
			fmt.Printf("    Variants: %d\n", len(product.Variants))
		}
		if product.RequiresSellingPlan {
			fmt.Printf("    Type: Subscription Only\n")
		}
		fmt.Printf("    Created: %d\n", product.CreatedAt)
		fmt.Println()
	}
}

// createProductWithVariants creates a product with multiple variants
func createProductWithVariants(client *martianpay.Client) {
	fmt.Println("Creating Product with Variants...")
	fmt.Println("This example creates a T-shirt with Size and Color variants")

	name := "Premium T-Shirt"
	description := "High-quality cotton t-shirt available in multiple sizes and colors"

	// Define options (sort_order starts from 0)
	options := []*developer.ProductOption{
		{
			Name: "Color",
			// SortOrder: 0 (default, first option)
			Values: []*developer.ProductOptionValue{
				{
					Value:     "Black",
					SortOrder: 0,
					Swatch:    &developer.ProductOptionSwatch{Type: "color", Value: "#000000"},
				},
				{
					Value:     "White",
					SortOrder: 1,
					Swatch:    &developer.ProductOptionSwatch{Type: "color", Value: "#FFFFFF"},
				},
				{
					Value:     "Blue",
					SortOrder: 2,
					Swatch:    &developer.ProductOptionSwatch{Type: "color", Value: "#0000FF"},
				},
			},
		},
		{
			Name:      "Size",
			SortOrder: 1, // Second option
			Values: []*developer.ProductOptionValue{
				{
					Value:     "Small",
					SortOrder: 0,
					Swatch:    &developer.ProductOptionSwatch{Type: "text"},
				},
				{
					Value:     "Medium",
					SortOrder: 1,
					Swatch:    &developer.ProductOptionSwatch{Type: "text"},
				},
				{
					Value:     "Large",
					SortOrder: 2,
					Swatch:    &developer.ProductOptionSwatch{Type: "text"},
				},
			},
		},
	}

	// Define variants - Must cover all option combinations (3 sizes × 3 colors = 9 variants)
	inventoryQty := 100
	variants := []*developer.ProductVariant{
		// Black variants
		{
			OptionValues:      map[string]string{"Size": "Small", "Color": "Black"},
			Price:             &developer.AssetAmount{Amount: decimal.NewFromFloat(25.00), AssetId: "USD"},
			InventoryQuantity: &inventoryQty,
			Active:            true,
		},
		{
			OptionValues:      map[string]string{"Size": "Medium", "Color": "Black"},
			Price:             &developer.AssetAmount{Amount: decimal.NewFromFloat(27.00), AssetId: "USD"},
			InventoryQuantity: &inventoryQty,
			Active:            true,
		},
		{
			OptionValues:      map[string]string{"Size": "Large", "Color": "Black"},
			Price:             &developer.AssetAmount{Amount: decimal.NewFromFloat(29.00), AssetId: "USD"},
			InventoryQuantity: &inventoryQty,
			Active:            true,
		},
		// White variants
		{
			OptionValues:      map[string]string{"Size": "Small", "Color": "White"},
			Price:             &developer.AssetAmount{Amount: decimal.NewFromFloat(25.00), AssetId: "USD"},
			InventoryQuantity: &inventoryQty,
			Active:            true,
		},
		{
			OptionValues:      map[string]string{"Size": "Medium", "Color": "White"},
			Price:             &developer.AssetAmount{Amount: decimal.NewFromFloat(27.00), AssetId: "USD"},
			InventoryQuantity: &inventoryQty,
			Active:            true,
		},
		{
			OptionValues:      map[string]string{"Size": "Large", "Color": "White"},
			Price:             &developer.AssetAmount{Amount: decimal.NewFromFloat(29.00), AssetId: "USD"},
			InventoryQuantity: &inventoryQty,
			Active:            true,
		},
		// Blue variants
		{
			OptionValues:      map[string]string{"Size": "Small", "Color": "Blue"},
			Price:             &developer.AssetAmount{Amount: decimal.NewFromFloat(25.00), AssetId: "USD"},
			InventoryQuantity: &inventoryQty,
			Active:            true,
		},
		{
			OptionValues:      map[string]string{"Size": "Medium", "Color": "Blue"},
			Price:             &developer.AssetAmount{Amount: decimal.NewFromFloat(27.00), AssetId: "USD"},
			InventoryQuantity: &inventoryQty,
			Active:            true,
		},
		{
			OptionValues:      map[string]string{"Size": "Large", "Color": "Blue"},
			Price:             &developer.AssetAmount{Amount: decimal.NewFromFloat(29.00), AssetId: "USD"},
			InventoryQuantity: &inventoryQty,
			Active:            true,
		},
	}

	req := &developer.ProductCreateRequest{
		Product: developer.Product{
			Name:                   name,
			Description:            description,
			Active:                 true,
			DefaultCurrency:        "USD",
			TaxCode:                "txcd_99999999", // General tangible goods tax code
			CollectShippingAddress: true,
			CollectTaxAddress:      true,
			Options:                options,
			Variants:               variants,
			Metadata: map[string]string{
				"category": "apparel",
				"sku":      "TSHIRT-001",
			},
		},
	}

	response, err := client.CreateProduct(req)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("\n✓ Product with Variants Created:\n")
	fmt.Printf("  ID: %s\n", response.ID)
	fmt.Printf("  Name: %s\n", response.Name)
	fmt.Printf("  Options: %d\n", len(response.Options))
	for _, opt := range response.Options {
		fmt.Printf("    - %s: %d values\n", opt.Name, len(opt.Values))
	}
	fmt.Printf("  Variants: %d\n", len(response.Variants))
	for i, variant := range response.Variants {
		fmt.Printf("    [%d] ", i+1)
		for optName, optValue := range variant.OptionValues {
			fmt.Printf("%s: %s, ", optName, optValue)
		}
		if variant.Price != nil {
			fmt.Printf("Price: %s %s", variant.Price.Amount, variant.Price.AssetId)
		}
		fmt.Println()
	}
}

// createProductWithSellingPlan creates a product with selling plan group attached
func createProductWithSellingPlan(client *martianpay.Client) {
	fmt.Println("Creating Product with Selling Plan...")
	fmt.Println("This example creates a subscription product with selling plan group")

	// First, list selling plan groups to select from
	fmt.Println("\n  Fetching selling plan groups...")
	groupParams := &developer.Pagination{
		Page:     0,
		PageSize: 10,
	}
	groupsResp, err := client.ListSellingPlanGroups(groupParams)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	if len(groupsResp.Data) == 0 {
		fmt.Println("✗ No selling plan groups found. Please create one first.")
		return
	}

	fmt.Printf("\n  Available Selling Plan Groups:\n")
	for i, group := range groupsResp.Data {
		fmt.Printf("  [%d] %s (ID: %s)\n", i+1, group.Name, group.ID)
	}

	fmt.Print("\nEnter group number or ID (or press Enter for first): ")
	var groupChoice string
	fmt.Scanln(&groupChoice)

	selectedIdx := 0
	if groupChoice != "" {
		// Try to find by ID first
		foundByID := false
		for i, group := range groupsResp.Data {
			if group.ID == groupChoice {
				selectedIdx = i
				foundByID = true
				break
			}
		}
		// If not found by ID, try as number
		if !foundByID {
			var idx int
			fmt.Sscanf(groupChoice, "%d", &idx)
			if idx > 0 && idx <= len(groupsResp.Data) {
				selectedIdx = idx - 1
			}
		}
	}

	selectedGroup := groupsResp.Data[selectedIdx]
	fmt.Printf("  Selected: %s (ID: %s)\n", selectedGroup.Name, selectedGroup.ID)

	// Ask if subscription is required
	fmt.Print("\nRequire subscription? (y=subscription only, n=one-time or subscription) [n]: ")
	var requiresSubStr string
	fmt.Scanln(&requiresSubStr)
	requiresSellingPlan := strings.ToLower(strings.TrimSpace(requiresSubStr)) == "y"

	name := "Monthly Newsletter Subscription"
	description := "Premium newsletter subscription with exclusive content"

	// Define simple variant structure for subscription product
	options := []*developer.ProductOption{
		{
			Name: "Tier",
			Values: []*developer.ProductOptionValue{
				{
					Value:     "Basic",
					SortOrder: 0,
					Swatch:    &developer.ProductOptionSwatch{Type: "text"},
				},
				{
					Value:     "Premium",
					SortOrder: 1,
					Swatch:    &developer.ProductOptionSwatch{Type: "text"},
				},
			},
		},
	}

	inventoryQty := 999
	variants := []*developer.ProductVariant{
		{
			OptionValues:      map[string]string{"Tier": "Basic"},
			Price:             &developer.AssetAmount{Amount: decimal.NewFromFloat(9.99), AssetId: "USD"},
			InventoryQuantity: &inventoryQty,
			Active:            true,
		},
		{
			OptionValues:      map[string]string{"Tier": "Premium"},
			Price:             &developer.AssetAmount{Amount: decimal.NewFromFloat(19.99), AssetId: "USD"},
			InventoryQuantity: &inventoryQty,
			Active:            true,
		},
	}

	req := &developer.ProductCreateRequest{
		Product: developer.Product{
			Name:                   name,
			Description:            description,
			Active:                 true,
			DefaultCurrency:        "USD",
			TaxCode:                "txcd_10000000", // Digital goods tax code
			CollectShippingAddress: false,           // No shipping for digital products
			CollectTaxAddress:      true,
			RequiresSellingPlan:    requiresSellingPlan,
			SellingPlanGroupIDs:    []string{selectedGroup.ID},
			Options:                options,
			Variants:               variants,
			Metadata: map[string]string{
				"category": "subscription",
				"type":     "newsletter",
			},
		},
	}

	response, err := client.CreateProduct(req)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("\n✓ Product with Selling Plan Created:\n")
	fmt.Printf("  ID: %s\n", response.ID)
	fmt.Printf("  Name: %s\n", response.Name)
	fmt.Printf("  Description: %s\n", response.Description)
	fmt.Printf("  Requires Selling Plan: %v\n", response.RequiresSellingPlan)
	if len(response.SellingPlanGroupIDs) > 0 {
		fmt.Printf("  Selling Plan Groups: %d\n", len(response.SellingPlanGroupIDs))
		for _, spgID := range response.SellingPlanGroupIDs {
			fmt.Printf("    - %s\n", spgID)
		}
	}
	fmt.Printf("  Options: %d\n", len(response.Options))
	for _, opt := range response.Options {
		fmt.Printf("    - %s: %d values\n", opt.Name, len(opt.Values))
	}
	fmt.Printf("  Variants: %d\n", len(response.Variants))
	for i, variant := range response.Variants {
		fmt.Printf("    [%d] ", i+1)
		for optName, optValue := range variant.OptionValues {
			fmt.Printf("%s: %s, ", optName, optValue)
		}
		if variant.Price != nil {
			fmt.Printf("Price: %s %s", variant.Price.Amount, variant.Price.AssetId)
		}
		fmt.Println()
	}
}

// getProduct retrieves details of a specific product
func getProduct(client *martianpay.Client) {
	fmt.Println("Getting Product...")
	fmt.Println("  Fetching products...")

	// List products first
	listReq := &developer.ProductListRequest{
		Page:     0,
		PageSize: 10,
	}
	listResp, err := client.ListProducts(listReq)
	if err == nil && len(listResp.Products) > 0 {
		fmt.Printf("\n  Available Products:\n")
		for i, prod := range listResp.Products {
			fmt.Printf("  [%d] %s (ID: %s)\n", i+1, prod.Name, prod.ID)
		}
		fmt.Print("\nEnter product number or ID: ")
	} else {
		fmt.Print("\nEnter Product ID: ")
	}

	var choice string
	fmt.Scanln(&choice)

	var id string
	if choice != "" && listResp != nil && len(listResp.Products) > 0 {
		// Try to find by ID first
		foundByID := false
		for _, prod := range listResp.Products {
			if prod.ID == choice {
				id = choice
				foundByID = true
				break
			}
		}
		// If not found by ID, try as number
		if !foundByID {
			var idx int
			fmt.Sscanf(choice, "%d", &idx)
			if idx > 0 && idx <= len(listResp.Products) {
				id = listResp.Products[idx-1].ID
			}
		}
		if id == "" {
			id = choice
		}
	} else if choice != "" {
		id = choice
	} else {
		id = "prod_example_id"
	}

	response, err := client.GetProduct(id, nil)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("\n✓ Product Details:\n")
	fmt.Printf("  ID: %s\n", response.ID)
	fmt.Printf("  Name: %s\n", response.Name)
	fmt.Printf("  Description: %s\n", response.Description)
	fmt.Printf("  Active: %v\n", response.Active)
	fmt.Printf("  Default Currency: %s\n", response.DefaultCurrency)
	if response.Price != nil {
		fmt.Printf("  Price: %s %s\n", response.Price.Amount, response.Price.AssetId)
	}
	fmt.Printf("  Collect Shipping: %v\n", response.CollectShippingAddress)
	fmt.Printf("  Collect Tax: %v\n", response.CollectTaxAddress)
	fmt.Printf("  Requires Subscription: %v\n", response.RequiresSellingPlan)
	fmt.Printf("  Version: %d\n", response.Version)
	fmt.Printf("  Created: %d\n", response.CreatedAt)
	fmt.Printf("  Updated: %d\n\n", response.UpdatedAt)

	if len(response.Options) > 0 {
		fmt.Printf("  Options:\n")
		for _, opt := range response.Options {
			fmt.Printf("    - %s:\n", opt.Name)
			for _, val := range opt.Values {
				fmt.Printf("      • %s\n", val.Value)
			}
		}
		fmt.Println()
	}

	if len(response.Variants) > 0 {
		fmt.Printf("  Variants (%d):\n", len(response.Variants))
		for i, variant := range response.Variants {
			fmt.Printf("    [%d] ID: %s\n", i+1, variant.ID)
			fmt.Printf("        Options: ")
			for optName, optValue := range variant.OptionValues {
				fmt.Printf("%s=%s ", optName, optValue)
			}
			fmt.Println()
			if variant.Price != nil {
				fmt.Printf("        Price: %s %s\n", variant.Price.Amount, variant.Price.AssetId)
			}
			fmt.Printf("        Active: %v\n", variant.Active)
		}
	}

	if len(response.Metadata) > 0 {
		fmt.Printf("\n  Metadata:\n")
		for key, value := range response.Metadata {
			fmt.Printf("    %s: %s\n", key, value)
		}
	}
}

// updateProduct updates an existing product
func updateProduct(client *martianpay.Client) {
	fmt.Println("Updating Product...")
	fmt.Print("Enter Product ID: ")

	var id string
	fmt.Scanln(&id)
	if id == "" {
		fmt.Println("✗ Product ID is required")
		return
	}

	// First get the current product
	current, err := client.GetProduct(id, nil)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("\nCurrent Product: %s\n", current.Name)
	fmt.Print("Enter new name (press Enter to keep current): ")
	var newName string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		newName = scanner.Text()
	}
	if newName == "" {
		newName = current.Name
	}

	fmt.Printf("Active status (current: %v). Change? (y/n): ", current.Active)
	var changeActive string
	fmt.Scanln(&changeActive)

	active := current.Active
	if strings.ToLower(changeActive) == "y" {
		active = !active
	}

	req := &developer.ProductUpdateRequest{
		ProductCreateRequest: developer.ProductCreateRequest{
			Product: developer.Product{
				Name:                   newName,
				Description:            current.Description,
				Active:                 active,
				DefaultCurrency:        current.DefaultCurrency,
				TaxCode:                current.TaxCode,
				CollectShippingAddress: current.CollectShippingAddress,
				CollectTaxAddress:      current.CollectTaxAddress,
				RequiresSellingPlan:    current.RequiresSellingPlan,
				SellingPlanGroupIDs:    current.SellingPlanGroupIDs,
				Options:                current.Options,
				Variants:               current.Variants,
				Metadata:               current.Metadata,
				Version:                current.Version, // Required for optimistic locking
			},
		},
	}

	response, err := client.UpdateProduct(id, req)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("\n✓ Product Updated:\n")
	fmt.Printf("  ID: %s\n", response.ID)
	fmt.Printf("  Name: %s\n", response.Name)
	fmt.Printf("  Active: %v\n", response.Active)
	fmt.Printf("  Version: %d\n", response.Version)
}

// deleteProduct deletes an inactive product
func deleteProduct(client *martianpay.Client) {
	fmt.Println("Deleting Product...")
	fmt.Print("Enter Product ID: ")

	var id string
	fmt.Scanln(&id)
	if id == "" {
		fmt.Println("✗ Product ID is required")
		return
	}

	fmt.Print("Are you sure? (yes/no): ")
	var confirm string
	fmt.Scanln(&confirm)
	if strings.ToLower(confirm) != "yes" {
		fmt.Println("  Cancelled")
		return
	}

	err := client.DeleteProduct(id)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		fmt.Println("  Note: Only inactive products can be deleted")
		return
	}

	fmt.Printf("✓ Product Deleted: %s\n", id)
}

// listActiveProducts lists only active products
func listActiveProducts(client *martianpay.Client) {
	fmt.Println("Listing Active Products...")

	active := true
	req := &developer.ProductListRequest{
		Page:     0,
		PageSize: 10,
		Active:   &active,
	}

	response, err := client.ListProducts(req)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("✓ Active Products: %d\n\n", response.Total)

	if len(response.Products) == 0 {
		fmt.Println("  No active products found")
		return
	}

	for i, product := range response.Products {
		fmt.Printf("[%d] %s (ID: %s)\n", i+1, product.Name, product.ID)
		if product.Price != nil {
			fmt.Printf("    Price: %s %s\n", product.Price.Amount, product.Price.AssetId)
		}
		if len(product.Variants) > 0 {
			fmt.Printf("    Variants: %d\n", len(product.Variants))
		}
	}
}

// Selling Plan Group Examples

// listSellingPlanGroups lists all selling plan groups
func listSellingPlanGroups(client *martianpay.Client) {
	fmt.Println("Listing Selling Plan Groups...")

	params := &developer.Pagination{
		Page:     0,
		PageSize: 10,
	}

	response, err := client.ListSellingPlanGroups(params)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("✓ Total Selling Plan Groups: %d\n", response.Total)
	fmt.Printf("  Offset: %d, Limit: %d\n\n", response.Offset, response.Limit)

	if len(response.Data) == 0 {
		fmt.Println("  No selling plan groups found")
		return
	}

	for i, group := range response.Data {
		fmt.Printf("[%d] %s\n", i+1, group.Name)
		fmt.Printf("    ID: %s\n", group.ID)
		fmt.Printf("    Status: %s\n", group.Status)
		if group.Description != "" {
			fmt.Printf("    Description: %s\n", group.Description)
		}
		if len(group.Options) > 0 {
			fmt.Printf("    Options: %v\n", group.Options)
		}
		if len(group.SellingPlans) > 0 {
			fmt.Printf("    Selling Plans: %d\n", len(group.SellingPlans))
		}
		fmt.Printf("    Created: %d\n", group.CreatedAt)
		fmt.Println()
	}
}

// createSellingPlanGroup creates a new selling plan group
func createSellingPlanGroup(client *martianpay.Client) {
	fmt.Println("Creating Selling Plan Group...")

	fmt.Print("Enter group name: ")
	var name string
	fmt.Scanln(&name)
	if name == "" {
		name = "Subscription Plans"
	}

	fmt.Print("Enter description: ")
	var description string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		description = scanner.Text()
	}
	if description == "" {
		description = "Flexible subscription options for recurring purchases"
	}

	req := &developer.CreateSellingPlanGroupRequest{
		Name:        name,
		Description: description,
		Options:     []string{"Billing Frequency"},
	}

	response, err := client.CreateSellingPlanGroup(req)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("\n✓ Selling Plan Group Created:\n")
	fmt.Printf("  ID: %s\n", response.ID)
	fmt.Printf("  Name: %s\n", response.Name)
	fmt.Printf("  Description: %s\n", response.Description)
	fmt.Printf("  Status: %s\n", response.Status)
	if len(response.Options) > 0 {
		fmt.Printf("  Options: %v\n", response.Options)
	}
}

// getSellingPlanGroup retrieves details of a specific selling plan group
func getSellingPlanGroup(client *martianpay.Client) {
	fmt.Println("Getting Selling Plan Group...")
	fmt.Print("Enter Selling Plan Group ID: ")

	var id string
	fmt.Scanln(&id)
	if id == "" {
		id = "spg_example_id"
	}

	response, err := client.GetSellingPlanGroup(id)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("\n✓ Selling Plan Group Details:\n")
	fmt.Printf("  ID: %s\n", response.ID)
	fmt.Printf("  Name: %s\n", response.Name)
	fmt.Printf("  Description: %s\n", response.Description)
	fmt.Printf("  Status: %s\n", response.Status)
	if len(response.Options) > 0 {
		fmt.Printf("  Options: %v\n", response.Options)
	}
	fmt.Printf("  Created: %d\n", response.CreatedAt)
	fmt.Printf("  Updated: %d\n\n", response.UpdatedAt)

	if len(response.SellingPlans) > 0 {
		fmt.Printf("  Selling Plans (%d):\n", len(response.SellingPlans))
		for i, plan := range response.SellingPlans {
			fmt.Printf("    [%d] %s (ID: %s)\n", i+1, plan.Name, plan.ID)
			fmt.Printf("        Type: %s\n", plan.BillingPolicyType)
			fmt.Printf("        Interval: Every %s %s\n",
				plan.BillingPolicy.IntervalCount,
				plan.BillingPolicy.Interval)
			if plan.TrialPeriodDays != "" && plan.TrialPeriodDays != "0" {
				fmt.Printf("        Trial: %s days\n", plan.TrialPeriodDays)
			}
			fmt.Printf("        Status: %s\n", plan.Status)
		}
	}
}

// updateSellingPlanGroup updates an existing selling plan group
func updateSellingPlanGroup(client *martianpay.Client) {
	fmt.Println("Updating Selling Plan Group...")
	fmt.Print("Enter Selling Plan Group ID: ")

	var id string
	fmt.Scanln(&id)
	if id == "" {
		fmt.Println("✗ Selling Plan Group ID is required")
		return
	}

	// First get the current group
	current, err := client.GetSellingPlanGroup(id)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("\nCurrent Group: %s\n", current.Name)
	fmt.Print("Enter new name (press Enter to keep current): ")
	var newName string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		newName = scanner.Text()
	}
	if newName == "" {
		newName = current.Name
	}

	fmt.Printf("Status (current: %s). Change? (active/disabled/keep): ", current.Status)
	var statusStr string
	fmt.Scanln(&statusStr)

	status := current.Status
	if statusStr == "active" || statusStr == "disabled" {
		status = statusStr
	}

	req := &developer.UpdateSellingPlanGroupRequest{
		Name:        newName,
		Description: current.Description,
		Options:     current.Options,
		Status:      status,
	}

	response, err := client.UpdateSellingPlanGroup(id, req)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("\n✓ Selling Plan Group Updated:\n")
	fmt.Printf("  ID: %s\n", response.ID)
	fmt.Printf("  Name: %s\n", response.Name)
	fmt.Printf("  Status: %s\n", response.Status)
}

// deleteSellingPlanGroup deletes a selling plan group
func deleteSellingPlanGroup(client *martianpay.Client) {
	fmt.Println("Deleting Selling Plan Group...")
	fmt.Print("Enter Selling Plan Group ID: ")

	var id string
	fmt.Scanln(&id)
	if id == "" {
		fmt.Println("✗ Selling Plan Group ID is required")
		return
	}

	fmt.Print("Are you sure? (yes/no): ")
	var confirm string
	fmt.Scanln(&confirm)
	if strings.ToLower(confirm) != "yes" {
		fmt.Println("  Cancelled")
		return
	}

	err := client.DeleteSellingPlanGroup(id)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		fmt.Println("  Note: Cannot delete a group that has active selling plans")
		return
	}

	fmt.Printf("✓ Selling Plan Group Deleted: %s\n", id)
}

// Selling Plan Examples

// listSellingPlans lists all selling plans
func listSellingPlans(client *martianpay.Client) {
	fmt.Println("Listing Selling Plans...")

	params := &developer.Pagination{
		Page:     0,
		PageSize: 20,
	}

	response, err := client.ListSellingPlans(params)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("\n✓ Total Selling Plans: %d\n", response.Total)
	fmt.Printf("  Page: %d/%d (Page Size: %d)\n\n",
		response.Offset/response.Limit+1,
		(response.Total+int64(response.Limit)-1)/int64(response.Limit),
		response.Limit)

	if len(response.Data) == 0 {
		fmt.Println("  No selling plans found")
		return
	}

	for i, plan := range response.Data {
		fmt.Printf("[%d] %s\n", i+1, plan.Name)
		fmt.Printf("    ID: %s\n", plan.ID)
		fmt.Printf("    Group ID: %s\n", plan.SellingPlanGroupID)
		fmt.Printf("    Type: %s\n", plan.BillingPolicyType)
		fmt.Printf("    Interval: Every %s %s\n",
			plan.BillingPolicy.IntervalCount,
			plan.BillingPolicy.Interval)
		if plan.TrialPeriodDays != "" && plan.TrialPeriodDays != "0" {
			fmt.Printf("    Trial: %s days\n", plan.TrialPeriodDays)
		}
		if len(plan.PricingPolicy) > 0 {
			fmt.Printf("    Pricing Policies:\n")
			for _, policy := range plan.PricingPolicy {
				fmt.Printf("      - %s: %s%% %s",
					policy.PolicyType,
					policy.AdjustmentValue,
					policy.AdjustmentType)
				if policy.AfterCycle != nil {
					fmt.Printf(" (after cycle %s)", *policy.AfterCycle)
				}
				fmt.Println()
			}
		}
		fmt.Printf("    Priority: %s\n", plan.Priority)
		fmt.Printf("    Status: %s\n", plan.Status)
		fmt.Printf("    Created: %d\n", plan.CreatedAt)
		fmt.Println()
	}
}

// createSellingPlan creates a new selling plan
func createSellingPlan(client *martianpay.Client) {
	fmt.Println("Creating Selling Plan...")

	// First, list groups to select from
	fmt.Println("  Fetching selling plan groups...")
	groupParams := &developer.Pagination{
		Page:     0,
		PageSize: 10,
	}
	groupsResp, err := client.ListSellingPlanGroups(groupParams)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	if len(groupsResp.Data) == 0 {
		fmt.Println("✗ No selling plan groups found. Please create one first.")
		return
	}

	fmt.Printf("\n  Available Selling Plan Groups:\n")
	for i, group := range groupsResp.Data {
		fmt.Printf("  [%d] %s (ID: %s)\n", i+1, group.Name, group.ID)
	}

	fmt.Print("\nEnter group number or ID (or press Enter for first): ")
	var groupChoice string
	fmt.Scanln(&groupChoice)

	selectedIdx := 0
	if groupChoice != "" {
		// Try to find by ID first
		foundByID := false
		for i, group := range groupsResp.Data {
			if group.ID == groupChoice {
				selectedIdx = i
				foundByID = true
				break
			}
		}

		// If not found by ID, try as number
		if !foundByID {
			var idx int
			fmt.Sscanf(groupChoice, "%d", &idx)
			if idx > 0 && idx <= len(groupsResp.Data) {
				selectedIdx = idx - 1
			}
		}
	}

	selectedGroup := groupsResp.Data[selectedIdx]
	fmt.Printf("  Selected: %s (ID: %s)\n\n", selectedGroup.Name, selectedGroup.ID)

	// Provide example options
	fmt.Println("  Example Selling Plans:")
	fmt.Println("  [1] Monthly Subscription (10% off)")
	fmt.Println("  [2] Quarterly Subscription (15% off)")
	fmt.Println("  [3] Annual Subscription (20% off)")
	fmt.Print("\nSelect example (or press Enter for #1): ")

	var exampleChoice string
	fmt.Scanln(&exampleChoice)

	var name, description, interval, intervalCount, trialDays string
	var pricingPolicy developer.PricingPolicyRequest

	switch exampleChoice {
	case "2":
		name = "Quarterly Subscription"
		description = "Save 15% with quarterly billing"
		interval = "month"
		intervalCount = "3"
		trialDays = "7"
		// FIXED policy should not have after_cycle
		pricingPolicy = developer.PricingPolicyRequest{
			{
				PolicyType:      "FIXED",
				AdjustmentType:  "PERCENTAGE",
				AdjustmentValue: "15",
			},
		}
	case "3":
		name = "Annual Subscription"
		description = "Save 20% with annual billing"
		interval = "year"
		intervalCount = "1"
		trialDays = "14"
		// FIXED policy should not have after_cycle
		pricingPolicy = developer.PricingPolicyRequest{
			{
				PolicyType:      "FIXED",
				AdjustmentType:  "PERCENTAGE",
				AdjustmentValue: "20",
			},
		}
	default:
		name = "Monthly Subscription"
		description = "Save 10% with monthly billing"
		interval = "month"
		intervalCount = "1"
		trialDays = "7"
		// FIXED policy should not have after_cycle
		pricingPolicy = developer.PricingPolicyRequest{
			{
				PolicyType:      "FIXED",
				AdjustmentType:  "PERCENTAGE",
				AdjustmentValue: "10",
			},
		}
	}

	req := &developer.CreateSellingPlanRequest{
		SellingPlanGroupID: selectedGroup.ID,
		Name:               name,
		Description:        description,
		BillingPolicyType:  "RECURRING",
		BillingPolicy: developer.BillingPolicyRequest{
			Interval:      interval,
			IntervalCount: intervalCount,
		},
		PricingPolicy:   &pricingPolicy,
		TrialPeriodDays: trialDays,
		Priority:        "1",
	}

	fmt.Printf("\nSending request:\n")
	fmt.Printf("  Group ID: %s\n", req.SellingPlanGroupID)
	fmt.Printf("  Name: %s\n", req.Name)
	fmt.Printf("  Billing: Every %s %s\n", req.BillingPolicy.IntervalCount, req.BillingPolicy.Interval)
	fmt.Printf("  Trial Days: %s\n", req.TrialPeriodDays)
	fmt.Printf("  Priority: %s\n", req.Priority)
	if req.PricingPolicy != nil && len(*req.PricingPolicy) > 0 {
		fmt.Printf("  Pricing Policy: %d policies\n", len(*req.PricingPolicy))
	}

	// Debug: Print the actual JSON being sent
	jsonBytes, _ := json.Marshal(req)
	fmt.Printf("\nJSON Payload:\n%s\n\n", string(jsonBytes))

	response, err := client.CreateSellingPlan(req)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("\n✓ Selling Plan Created:\n")
	fmt.Printf("  ID: %s\n", response.ID)
	fmt.Printf("  Name: %s\n", response.Name)
	fmt.Printf("  Description: %s\n", response.Description)
	fmt.Printf("  Type: %s\n", response.BillingPolicyType)
	fmt.Printf("  Interval: Every %s %s\n",
		response.BillingPolicy.IntervalCount,
		response.BillingPolicy.Interval)
	fmt.Printf("  Trial Period: %s days\n", response.TrialPeriodDays)
	if len(response.PricingPolicy) > 0 {
		fmt.Printf("  Pricing Policy:\n")
		for _, policy := range response.PricingPolicy {
			fmt.Printf("    - Type: %s, Adjustment: %s%% %s\n",
				policy.PolicyType,
				policy.AdjustmentValue,
				policy.AdjustmentType)
		}
	}
	fmt.Printf("  Status: %s\n", response.Status)
}

// getSellingPlan retrieves details of a specific selling plan
func getSellingPlan(client *martianpay.Client) {
	fmt.Println("Getting Selling Plan...")
	fmt.Print("Enter Selling Plan ID: ")

	var id string
	fmt.Scanln(&id)
	if id == "" {
		id = "sp_example_id"
	}

	response, err := client.GetSellingPlan(id)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("\n✓ Selling Plan Details:\n")
	fmt.Printf("  ID: %s\n", response.ID)
	fmt.Printf("  Group ID: %s\n", response.SellingPlanGroupID)
	fmt.Printf("  Name: %s\n", response.Name)
	fmt.Printf("  Description: %s\n", response.Description)
	fmt.Printf("  Type: %s\n", response.BillingPolicyType)
	fmt.Printf("  Status: %s\n", response.Status)
	fmt.Printf("  Priority: %s\n", response.Priority)

	fmt.Printf("\n  Billing Policy:\n")
	fmt.Printf("    Interval: Every %s %s\n",
		response.BillingPolicy.IntervalCount,
		response.BillingPolicy.Interval)
	if response.BillingPolicy.MinCycles != nil {
		fmt.Printf("    Min Cycles: %s\n", *response.BillingPolicy.MinCycles)
	}

	if response.TrialPeriodDays != "" && response.TrialPeriodDays != "0" {
		fmt.Printf("\n  Trial Period: %s days\n", response.TrialPeriodDays)
	}

	if len(response.PricingPolicy) > 0 {
		fmt.Printf("\n  Pricing Policy:\n")
		for i, policy := range response.PricingPolicy {
			fmt.Printf("    [%d] Policy Type: %s\n", i+1, policy.PolicyType)
			fmt.Printf("        Adjustment: %s %s\n",
				policy.AdjustmentValue,
				policy.AdjustmentType)
			if policy.AfterCycle != nil {
				fmt.Printf("        After Cycle: %s\n", *policy.AfterCycle)
			}
		}
	}

	if response.ValidFrom != nil {
		fmt.Printf("\n  Valid From: %d\n", *response.ValidFrom)
	}
	if response.ValidUntil != nil {
		fmt.Printf("  Valid Until: %d\n", *response.ValidUntil)
	}

	fmt.Printf("\n  Created: %d\n", response.CreatedAt)
	fmt.Printf("  Updated: %d\n", response.UpdatedAt)
}

// updateSellingPlan updates an existing selling plan
func updateSellingPlan(client *martianpay.Client) {
	fmt.Println("Updating Selling Plan...")
	fmt.Print("Enter Selling Plan ID: ")

	var id string
	fmt.Scanln(&id)
	if id == "" {
		fmt.Println("✗ Selling Plan ID is required")
		return
	}

	// First get the current plan
	current, err := client.GetSellingPlan(id)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("\nCurrent Plan: %s\n", current.Name)
	fmt.Printf("Status (current: %s). Change? (active/disabled/keep): ", current.Status)
	var statusStr string
	fmt.Scanln(&statusStr)

	status := current.Status
	if statusStr == "active" || statusStr == "disabled" {
		status = statusStr
	}

	req := &developer.UpdateSellingPlanRequest{
		Name:              current.Name,
		Description:       current.Description,
		BillingPolicyType: current.BillingPolicyType,
		TrialPeriodDays:   current.TrialPeriodDays,
		Priority:          current.Priority,
		Status:            status,
	}

	response, err := client.UpdateSellingPlan(id, req)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("\n✓ Selling Plan Updated:\n")
	fmt.Printf("  ID: %s\n", response.ID)
	fmt.Printf("  Name: %s\n", response.Name)
	fmt.Printf("  Status: %s\n", response.Status)
}

// deleteSellingPlan deletes a selling plan
func deleteSellingPlan(client *martianpay.Client) {
	fmt.Println("Deleting Selling Plan...")
	fmt.Print("Enter Selling Plan ID: ")

	var id string
	fmt.Scanln(&id)
	if id == "" {
		fmt.Println("✗ Selling Plan ID is required")
		return
	}

	fmt.Print("Are you sure? (yes/no): ")
	var confirm string
	fmt.Scanln(&confirm)
	if strings.ToLower(confirm) != "yes" {
		fmt.Println("  Cancelled")
		return
	}

	err := client.DeleteSellingPlan(id)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("✓ Selling Plan Deleted: %s\n", id)
}

// calculateSellingPlanPrice calculates subscription price for a variant
func calculateSellingPlanPrice(client *martianpay.Client) {
	fmt.Println("Calculating Selling Plan Price...")

	fmt.Print("Enter Variant ID: ")
	var variantID string
	fmt.Scanln(&variantID)
	if variantID == "" {
		fmt.Println("✗ Variant ID is required")
		return
	}

	fmt.Print("Enter Base Price: ")
	var basePrice string
	fmt.Scanln(&basePrice)
	if basePrice == "" {
		basePrice = "29.99"
	}

	fmt.Print("Enter Currency (currently only supports USD): ")
	var currency string
	fmt.Scanln(&currency)
	if currency == "" {
		currency = "USD"
	}

	fmt.Print("Enter Selling Plan ID: ")
	var planID string
	fmt.Scanln(&planID)
	if planID == "" {
		fmt.Println("✗ Selling Plan ID is required")
		return
	}

	params := map[string]interface{}{
		"variant_id": variantID,
		"base_price": basePrice,
		"currency":   currency,
		"plan_id":    planID,
	}

	response, err := client.CalculateSellingPlanPrice(params)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("\n✓ Subscription Price Calculation:\n")
	fmt.Printf("  Base Price: %s %s\n", response.BasePrice, response.Currency)
	fmt.Printf("  Billing Cycle: %s\n", response.BillingCycle)
	fmt.Printf("  Total Cycles: %d\n", response.TotalCycles)
	if response.TrialPeriodDays > 0 {
		fmt.Printf("  Trial Period: %d days\n", response.TrialPeriodDays)
	}
	fmt.Printf("\n  Subtotal Before Discount: %s %s\n",
		response.SubtotalBeforePolicy, response.Currency)
	fmt.Printf("  Selling Plan Discount: %s %s\n",
		response.SellingPlanDiscount, response.Currency)
	fmt.Printf("  Final Subscription Price: %s %s\n",
		response.SubtotalAfterPolicy, response.Currency)
}

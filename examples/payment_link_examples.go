package main

import (
	"fmt"
	"strings"

	"github.com/MartianPay/martianpay-go-sample/pkg/developer"
	martianpay "github.com/MartianPay/martianpay-go-sample/sdk"
)

// Payment Link Examples

// listPaymentLinks lists all payment links with pagination
func listPaymentLinks(client *martianpay.Client) {
	fmt.Println("Listing Payment Links...")

	req := &developer.PaymentLinkListRequest{
		Page:     0,
		PageSize: 10,
	}

	response, err := client.ListPaymentLinks(req)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("✓ Total Payment Links: %d\n", response.Total)
	fmt.Printf("  Page: %d, Page Size: %d\n\n", response.Page, response.PageSize)

	if len(response.PaymentLinks) == 0 {
		fmt.Println("  No payment links found")
		return
	}

	for i, link := range response.PaymentLinks {
		fmt.Printf("[%d] ID: %s\n", i+1, link.ID)
		if link.Product != nil {
			fmt.Printf("    Product: %s\n", link.Product.Name)
		}
		fmt.Printf("    Active: %v\n", link.Active)
		if link.URL != nil {
			fmt.Printf("    URL: %s\n", *link.URL)
		}
		if link.PriceRange != nil {
			if link.PriceRange.Min != nil {
				fmt.Printf("    Price Range: %s", link.PriceRange.Min.Amount)
				if link.PriceRange.Max != nil && link.PriceRange.Max.Amount != link.PriceRange.Min.Amount {
					fmt.Printf(" - %s", link.PriceRange.Max.Amount)
				}
				fmt.Printf(" %s\n", link.PriceRange.Min.AssetId)
			}
		}
		if len(link.PrimaryVariants) > 0 {
			fmt.Printf("    Primary Variants: %d\n", len(link.PrimaryVariants))
		}
		if len(link.AddonVariants) > 0 {
			fmt.Printf("    Addon Variants: %d\n", len(link.AddonVariants))
		}
		fmt.Printf("    Created: %d\n", link.CreatedAt)
		fmt.Println()
	}
}

// createPaymentLink creates a new payment link for a product
func createPaymentLink(client *martianpay.Client) {
	fmt.Println("Creating Payment Link...")

	// First, list products to select from
	fmt.Println("  Fetching products...")
	productsReq := &developer.ProductListRequest{
		Page:     0,
		PageSize: 10,
	}

	productsResp, err := client.ListProducts(productsReq)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	if len(productsResp.Products) == 0 {
		fmt.Println("✗ No products found. Please create a product first.")
		return
	}

	fmt.Printf("\n  Available Products:\n")
	for i, product := range productsResp.Products {
		fmt.Printf("  [%d] %s (ID: %s)\n", i+1, product.Name, product.ID)
		if len(product.Variants) > 0 {
			fmt.Printf("      Variants: %d\n", len(product.Variants))
		}
	}

	fmt.Print("\nEnter product number (or press Enter for first): ")
	var productChoice string
	fmt.Scanln(&productChoice)

	selectedIdx := 0
	if productChoice != "" {
		var idx int
		fmt.Sscanf(productChoice, "%d", &idx)
		if idx > 0 && idx <= len(productsResp.Products) {
			selectedIdx = idx - 1
		}
	}

	selectedProduct := productsResp.Products[selectedIdx]
	fmt.Printf("  Selected: %s\n", selectedProduct.Name)

	req := &developer.PaymentLinkCreateRequest{
		ProductID: selectedProduct.ID,
	}

	// If product has variants, select which ones to include
	if len(selectedProduct.Variants) > 0 {
		fmt.Printf("\n  Product has %d variants. Including all as primary variants.\n", len(selectedProduct.Variants))

		var primaryVariantIDs []string
		for _, variant := range selectedProduct.Variants {
			if variant.Active {
				primaryVariantIDs = append(primaryVariantIDs, variant.ID)
			}
		}
		req.PrimaryVariantIDs = primaryVariantIDs

		// Set first variant as default
		if len(primaryVariantIDs) > 0 {
			req.DefaultVariantID = primaryVariantIDs[0]
		}
	}

	paymentLink, err := client.CreatePaymentLink(req)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("\n✓ Payment Link Created:\n")
	fmt.Printf("  ID: %s\n", paymentLink.ID)
	if paymentLink.URL != nil {
		fmt.Printf("  URL: %s\n", *paymentLink.URL)
		fmt.Printf("\n  🔗 Share this URL to accept payments!\n")
	}
	if paymentLink.Product != nil {
		fmt.Printf("  Product: %s\n", paymentLink.Product.Name)
	}
	fmt.Printf("  Active: %v\n", paymentLink.Active)
	if len(paymentLink.PrimaryVariants) > 0 {
		fmt.Printf("  Primary Variants: %d\n", len(paymentLink.PrimaryVariants))
	}
}

// getPaymentLink retrieves details of a specific payment link
func getPaymentLink(client *martianpay.Client) {
	fmt.Println("Getting Payment Link...")
	fmt.Println("  Fetching payment links...")

	// List payment links first
	listReq := &developer.PaymentLinkListRequest{
		Page:     0,
		PageSize: 10,
	}
	listResp, err := client.ListPaymentLinks(listReq)
	if err == nil && len(listResp.PaymentLinks) > 0 {
		fmt.Printf("\n  Available Payment Links:\n")
		for i, link := range listResp.PaymentLinks {
			fmt.Printf("  [%d] ", i+1)
			if link.Product != nil {
				fmt.Printf("%s ", link.Product.Name)
			}
			fmt.Printf("(ID: %s)\n", link.ID)
		}
		fmt.Print("\nEnter payment link number or ID: ")
	} else {
		fmt.Print("\nEnter Payment Link ID: ")
	}

	var choice string
	fmt.Scanln(&choice)

	var id string
	if choice != "" && listResp != nil && len(listResp.PaymentLinks) > 0 {
		// Try to find by ID first
		foundByID := false
		for _, link := range listResp.PaymentLinks {
			if link.ID == choice {
				id = choice
				foundByID = true
				break
			}
		}
		// If not found by ID, try as number
		if !foundByID {
			var idx int
			fmt.Sscanf(choice, "%d", &idx)
			if idx > 0 && idx <= len(listResp.PaymentLinks) {
				id = listResp.PaymentLinks[idx-1].ID
			}
		}
		if id == "" {
			id = choice
		}
	} else if choice != "" {
		id = choice
	} else {
		id = "plink_example_id"
	}

	link, err := client.GetPaymentLink(id)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("\n✓ Payment Link Details:\n")
	fmt.Printf("  ID: %s\n", link.ID)
	if link.URL != nil {
		fmt.Printf("  URL: %s\n", *link.URL)
	}
	fmt.Printf("  Active: %v\n", link.Active)
	fmt.Printf("  Created: %d\n", link.CreatedAt)
	fmt.Printf("  Updated: %d\n\n", link.UpdatedAt)

	if link.Product != nil {
		fmt.Printf("Product:\n")
		fmt.Printf("  Name: %s\n", link.Product.Name)
		fmt.Printf("  ID: %s\n", link.Product.ID)
		if link.Product.Description != "" {
			fmt.Printf("  Description: %s\n", link.Product.Description)
		}
	}

	if link.PriceRange != nil {
		fmt.Printf("\nPrice Range:\n")
		if link.PriceRange.Min != nil {
			fmt.Printf("  Min: %s %s\n", link.PriceRange.Min.Amount, link.PriceRange.Min.AssetId)
		}
		if link.PriceRange.Max != nil {
			fmt.Printf("  Max: %s %s\n", link.PriceRange.Max.Amount, link.PriceRange.Max.AssetId)
		}
	}

	if len(link.PrimaryVariants) > 0 {
		fmt.Printf("\nPrimary Variants (%d):\n", len(link.PrimaryVariants))
		for i, pv := range link.PrimaryVariants {
			fmt.Printf("  [%d] Variant ID: %s\n", i+1, pv.VariantID)
			if pv.Variant != nil {
				fmt.Printf("      Options: ")
				for optName, optValue := range pv.Variant.OptionValues {
					fmt.Printf("%s=%s ", optName, optValue)
				}
				fmt.Println()
				if pv.Variant.Price != nil {
					fmt.Printf("      Price: %s %s\n", pv.Variant.Price.Amount, pv.Variant.Price.AssetId)
				}
			}
			fmt.Printf("      Quantity: %d\n", pv.Quantity)
		}
	}

	if len(link.AddonVariants) > 0 {
		fmt.Printf("\nAddon Variants (%d):\n", len(link.AddonVariants))
		for i, av := range link.AddonVariants {
			fmt.Printf("  [%d] Variant ID: %s\n", i+1, av.VariantID)
			if av.Variant != nil && av.Variant.Price != nil {
				fmt.Printf("      Price: %s %s\n", av.Variant.Price.Amount, av.Variant.Price.AssetId)
			}
			if av.MinQuantity != nil {
				fmt.Printf("      Min Quantity: %d\n", *av.MinQuantity)
			}
			if av.MaxQuantity != nil {
				fmt.Printf("      Max Quantity: %d\n", *av.MaxQuantity)
			}
		}
	}

	if link.Includes != nil && len(link.Includes.Media) > 0 {
		fmt.Printf("\nMedia Assets (%d):\n", len(link.Includes.Media))
		for i, media := range link.Includes.Media {
			fmt.Printf("  [%d] ID: %s\n", i+1, media.ID)
			if media.URL != "" {
				fmt.Printf("      URL: %s\n", media.URL)
			}
			if media.ContentType != "" {
				fmt.Printf("      Type: %s\n", media.ContentType)
			}
		}
	}
}

// updatePaymentLink updates a payment link's active status
func updatePaymentLink(client *martianpay.Client) {
	fmt.Println("Updating Payment Link...")
	fmt.Print("Enter Payment Link ID: ")

	var id string
	fmt.Scanln(&id)
	if id == "" {
		fmt.Println("✗ Payment Link ID is required")
		return
	}

	// First get the current link
	current, err := client.GetPaymentLink(id)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("\nCurrent Status: Active = %v\n", current.Active)
	fmt.Print("Set active status (true/false/toggle): ")

	var activeStr string
	fmt.Scanln(&activeStr)

	var active bool
	switch strings.ToLower(strings.TrimSpace(activeStr)) {
	case "true", "t", "1", "yes", "y":
		active = true
	case "false", "f", "0", "no", "n":
		active = false
	case "toggle", "":
		active = !current.Active
	default:
		fmt.Printf("Invalid input '%s', toggling current status\n", activeStr)
		active = !current.Active
	}

	req := &developer.PaymentLinkUpdateRequest{
		Active: &active,
	}

	fmt.Printf("Sending update request with Active = %v\n", *req.Active)

	link, err := client.UpdatePaymentLink(id, req)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("\n✓ Payment Link Updated:\n")
	fmt.Printf("  ID: %s\n", link.ID)
	fmt.Printf("  Active: %v\n", link.Active)
	if link.URL != nil {
		fmt.Printf("  URL: %s\n", *link.URL)
	}
}

// deletePaymentLink deletes an inactive payment link
func deletePaymentLink(client *martianpay.Client) {
	fmt.Println("Deleting Payment Link...")
	fmt.Print("Enter Payment Link ID: ")

	var id string
	fmt.Scanln(&id)
	if id == "" {
		fmt.Println("✗ Payment Link ID is required")
		return
	}

	fmt.Print("Are you sure? (yes/no): ")
	var confirm string
	fmt.Scanln(&confirm)
	if strings.ToLower(confirm) != "yes" {
		fmt.Println("  Cancelled")
		return
	}

	err := client.DeletePaymentLink(id)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		fmt.Println("  Note: Only inactive payment links can be deleted")
		return
	}

	fmt.Printf("✓ Payment Link Deleted: %s\n", id)
}

// listActivePaymentLinks lists only active payment links
func listActivePaymentLinks(client *martianpay.Client) {
	fmt.Println("Listing Active Payment Links...")

	active := true
	req := &developer.PaymentLinkListRequest{
		Page:     0,
		PageSize: 10,
		Active:   &active,
	}

	response, err := client.ListPaymentLinks(req)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("✓ Active Payment Links: %d\n\n", response.Total)

	if len(response.PaymentLinks) == 0 {
		fmt.Println("  No active payment links found")
		return
	}

	for i, link := range response.PaymentLinks {
		fmt.Printf("[%d] ID: %s\n", i+1, link.ID)
		if link.Product != nil {
			fmt.Printf("    Product: %s\n", link.Product.Name)
		}
		if link.URL != nil {
			fmt.Printf("    URL: %s\n", *link.URL)
		}
		if link.PriceRange != nil && link.PriceRange.Min != nil {
			fmt.Printf("    Price: %s %s", link.PriceRange.Min.Amount, link.PriceRange.Min.AssetId)
			if link.PriceRange.Max != nil && link.PriceRange.Max.Amount != link.PriceRange.Min.Amount {
				fmt.Printf(" - %s %s", link.PriceRange.Max.Amount, link.PriceRange.Max.AssetId)
			}
			fmt.Println()
		}
		fmt.Println()
	}
}

// listPaymentLinksByProduct lists payment links for a specific product
func listPaymentLinksByProduct(client *martianpay.Client) {
	fmt.Println("Listing Payment Links by Product...")
	fmt.Print("Enter Product ID: ")

	var productID string
	fmt.Scanln(&productID)
	if productID == "" {
		fmt.Println("✗ Product ID is required")
		return
	}

	req := &developer.PaymentLinkListRequest{
		Page:     0,
		PageSize: 10,
		Product:  productID,
	}

	response, err := client.ListPaymentLinks(req)
	if err != nil {
		fmt.Printf("✗ API Error: %v\n", err)
		return
	}

	fmt.Printf("✓ Payment Links for Product %s: %d\n\n", productID, response.Total)

	if len(response.PaymentLinks) == 0 {
		fmt.Println("  No payment links found for this product")
		return
	}

	for i, link := range response.PaymentLinks {
		fmt.Printf("[%d] ID: %s\n", i+1, link.ID)
		fmt.Printf("    Active: %v\n", link.Active)
		if link.URL != nil {
			fmt.Printf("    URL: %s\n", *link.URL)
		}
		if len(link.PrimaryVariants) > 0 {
			fmt.Printf("    Primary Variants: %d\n", len(link.PrimaryVariants))
		}
		fmt.Println()
	}
}

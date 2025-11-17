// Author: William Provost
// Version: 1.0.0
// Date: 2025-11-15
// Fileoverview: This program calculates the total cost of a car based on selected features.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Base price
	basePrice := 25000.00
	totalCost := basePrice

	// Feature choices
	var addFloorMats string
	var addNavigation string
	var addHeatedSeats string
	var addWarranty string

	reader := bufio.NewReader(os.Stdin)

	// Input feature choices
	fmt.Print("Would you like floor mats for $500? (yes/no) ")
	addFloorMats, _ = reader.ReadString('\n')
	addFloorMats = strings.TrimSpace(addFloorMats)

	fmt.Print("Would you like a navigation system for $1000? (yes/no) ")
	addNavigation, _ = reader.ReadString('\n')
	addNavigation = strings.TrimSpace(addNavigation)

	fmt.Print("Would you like heated leather seats for $500? (yes/no) ")
	addHeatedSeats, _ = reader.ReadString('\n')
	addHeatedSeats = strings.TrimSpace(addHeatedSeats)

	fmt.Print("Would you like a 5-year extended warranty for $2500? (yes/no) ")
	addWarranty, _ = reader.ReadString('\n')
	addWarranty = strings.TrimSpace(addWarranty)

	// Output base
	fmt.Printf("Base Price                              %.2f\n", basePrice)

	// Add features
	if strings.ToLower(addFloorMats) == "yes" {
		fmt.Println("Floor mats                              500.00")
		totalCost += 500
	}

	if strings.ToLower(addNavigation) == "yes" {
		fmt.Println("Navigation system                      1000.00")
		totalCost += 1000
	}

	if strings.ToLower(addHeatedSeats) == "yes" {
		fmt.Println("Heated leather seats                    500.00")
		totalCost += 500
	}

	if strings.ToLower(addWarranty) == "yes" {
		fmt.Println("5-Year extended warranty               2500.00")
		totalCost += 2500
	}

	// Tax
	taxAmount := totalCost * 0.13
	fmt.Printf("13%% Taxes                              %.2f\n", taxAmount)

	// Final total
	finalCost := totalCost + taxAmount
	fmt.Printf("Final cost of car                      %.2f\n", finalCost)

	fmt.Println("\nDone.")
}

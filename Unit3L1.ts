/**
 * @author William Provost
 * @version 1.0.0
 * @date 2025-11-15
 * @fileoverview This program calculates the total cost of a car based on selected features.
 */

// variables
const basePrice: number = 25000.0;
let totalCost1: number = basePrice;

// ask for features
const addFloorMats: string = prompt("Would you like floor mats for $500? (yes/no)") || "no";
const addNavigation: string = prompt("Would you like a navigation system for $1000? (yes/no)") || "no";
const addHeatedSeats: string = prompt("Would you like heated leather seats for $500? (yes/no)") || "no";
const addWarranty: string = prompt("Would you like a 5-year extended warranty for $2500? (yes/no)") || "no";

// output base price
console.log("Base Price                              " + basePrice.toFixed(2));

// check each feature
if (addFloorMats.toLowerCase() === "yes") {
  console.log("Floor mats                              500.00");
  totalCost1 += 500;
}

if (addNavigation.toLowerCase() === "yes") {
  console.log("Navigation system                      1000.00");
  totalCost1 += 1000;
}

if (addHeatedSeats.toLowerCase() === "yes") {
  console.log("Heated leather seats                    500.00");
  totalCost1 += 500;
}

if (addWarranty.toLowerCase() === "yes") {
  console.log("5-Year extended warranty               2500.00");
  totalCost1 += 2500;
}

// tax calculation
const taxAmount: number = totalCost * 0.13;
console.log("13% Taxes                              " + taxAmount.toFixed(2));

// final cost
const finalCost: number = totalCost + taxAmount;
console.log("Final cost of car                      " + finalCost.toFixed(2));

console.log("\nDone.");

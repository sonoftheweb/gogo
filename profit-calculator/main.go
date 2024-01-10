package main

import "fmt"

func main() {
	revenue := getUserInput("Enter revenue: ")
	totalExpenses := getUserInput("Enter total expenses: ")
	taxRate := getUserInput("Enter tax rate: ")

	earningBeforeTax, earningAfterTax, ratioEBTtoEAT := computeValues(revenue, totalExpenses, taxRate)

	fmt.Printf("Earnings Before Tax: %.2f\n", earningBeforeTax)
	fmt.Printf("Earnings After Tax: %.2f\n", earningAfterTax)
	fmt.Printf("Ratio of Earnings Before Tax to Earnings After Tax: %.2f\n", ratioEBTtoEAT)
}

func computeValues(revenue, totalExpenses, taxRate float64) (earningBeforeTax, earningAfterTax, ratioEBTtoEAT float64) {
	earningBeforeTax = revenue - totalExpenses
	earningAfterTax = earningBeforeTax - (1 - taxRate/100)
	ratioEBTtoEAT = earningBeforeTax / earningAfterTax

	return
}

func getUserInput(text string) float64 {
	var userInput float64

	fmt.Print(text)
	_, err := fmt.Scan(&userInput)
	if err != nil {
		fmt.Println("Invalid input. When prompted, please enter valid values.")
		return 0
	}

	return userInput
}

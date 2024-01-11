package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	revenue, revenueError := getUserInput("Enter revenue: ")
	if revenueError != nil {
		fmt.Println(revenueError)
		return
	}

	totalExpenses, totalExpensesError := getUserInput("Enter total expenses: ")
	if totalExpensesError != nil {
		fmt.Println(totalExpensesError)
		return
	}

	taxRate, taxRateError := getUserInput("Enter tax rate: ")
	if taxRateError != nil {
		fmt.Println(taxRateError)
		return
	}

	earningBeforeTax, earningAfterTax, ratioEBTtoEAT := computeValues(revenue, totalExpenses, taxRate)

	// store to file
	storeToFile(revenue, totalExpenses, taxRate, earningBeforeTax, earningAfterTax, ratioEBTtoEAT)

	fmt.Printf("Earnings Before Tax: %.2f\n", earningBeforeTax)
	fmt.Printf("Earnings After Tax: %.2f\n", earningAfterTax)
	fmt.Printf("Ratio of Earnings Before Tax to Earnings After Tax: %.2f\n", ratioEBTtoEAT)
}

func storeToFile(revenue, expenses, rate, tax, tax2, eat float64) {
	err := os.WriteFile("profit.txt", []byte(fmt.Sprintf("Revenue: %.2f\nExpenses: %.2f\nTax Rate: %.2f\nEarnings Before Tax: %.2f\nEarnings After Tax: %.2f\nRatio of Earnings Before Tax to Earnings After Tax: %.2f\n", revenue, expenses, rate, tax, tax2, eat)), 0644)
	if err != nil {
		panic(err)
	}
}

func computeValues(revenue, totalExpenses, taxRate float64) (earningBeforeTax, earningAfterTax, ratioEBTtoEAT float64) {
	earningBeforeTax = revenue - totalExpenses
	earningAfterTax = earningBeforeTax - (1 - taxRate/100)
	ratioEBTtoEAT = earningBeforeTax / earningAfterTax

	return
}

func getUserInput(text string) (float64, error) {
	var userInput float64

	fmt.Print(text)
	_, err := fmt.Scan(&userInput)
	if err != nil {
		// Clear the buffer
		var discard string
		_, err := fmt.Scanln(&discard)
		if err != nil {
			return 0, err
		}

		return 0, errors.New("invalid input. when prompted, please enter valid values")
	} else if userInput <= 0 {
		return 0, errors.New("invalid input. value must be greater than 0")
	} else {
		return userInput, nil
	}
}

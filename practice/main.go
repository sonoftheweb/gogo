package main

import (
	"fmt"
	"math"
)

const inflationRate float64 = 3.5

func main() {
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	outputText("Enter investment amount: ")
	_, errInvestmentAmount := fmt.Scan(&investmentAmount)

	outputText("Enter expected return rate: ")
	_, errExpectedReturnRate := fmt.Scan(&expectedReturnRate)

	outputText("Enter number of years: ")
	_, errYears := fmt.Scan(&years)

	if errInvestmentAmount != nil || errExpectedReturnRate != nil || errYears != nil {
		fmt.Println("Invalid input. When prompted, please enter valid values.")
		return
	}

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	// fmt.Println("Future value:", futureValue)
	//fmt.Printf("Future value: %v\n", futureValue)
	//fmt.Println("Future value (adjusted for inflation):", futureRealValue)
	fmt.Printf("Future value: %.2f\nFuture value (adjusted for inflation): %.2f\n", futureValue, futureRealValue)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)

	return fv, rfv
}

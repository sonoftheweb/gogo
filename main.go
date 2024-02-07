package main

import (
	"fmt"

	//"example.com/practice-calculator/cmdmanager"
	"example.com/practice-calculator/filemanager"
	"example.com/practice-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	var e error = nil
	for _, taxRate := range taxRates {
		if e != nil {
			fmt.Println("could not process job")
			fmt.Println(e)
			break
		}

		//io := cmdmanager.New()
		io := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		priceJob := prices.NewTaxIncludedPriceJob(taxRate, io)
		err := priceJob.Process()

		if err != nil {
			e = err
		}
	}
}

package main

import (
	"fmt"

	//"example.com/practice-calculator/cmdmanager"
	"example.com/practice-calculator/filemanager"
	"example.com/practice-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	doneChannels := make([]chan bool, len(taxRates))
	errorChannels := make([]chan error, len(taxRates))

	var e error = nil
	for idx, taxRate := range taxRates {
		doneChannels[idx] = make(chan bool)
		errorChannels[idx] = make(chan error)
		if e != nil {
			fmt.Println("could not process job")
			fmt.Println(e)
			break
		}

		//io := cmdmanager.New()
		io := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		priceJob := prices.NewTaxIncludedPriceJob(taxRate, io)
		go priceJob.Process(doneChannels[idx], errorChannels[idx])
	}

	for idx := range taxRates {
		select {
		case err := <-errorChannels[idx]:
			if err != nil {
				e = err
			}
		case <-doneChannels[idx]:
			fmt.Println("Done!")
		}
	}
}

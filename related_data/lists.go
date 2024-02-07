package main

import "fmt"

func main() {
	//var productNames = [4]string{"A book"}
	//prices := [4]float64{10.99, 9.99, 45.99, 20.0}
	//fmt.Println(prices)
	//fmt.Println(productNames)
	//
	//slice := prices[1:]
	//slice[0] = 199.99
	//slice2 := slice[:1]
	//fmt.Println(slice)
	//fmt.Println(slice2)
	//fmt.Println(prices)

	prices := []float64{10.99, 8.99}
	fmt.Println(prices[0:1])
	prices[1] = 9.99
	prices = append(prices, 5.99)
	prices = prices[1:]
	fmt.Println(prices)
}

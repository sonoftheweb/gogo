package main

import "fmt"

func main() {
	age := 32

	fmt.Println("Age:", age)
	getAdultYears(&age)
	fmt.Println("Age:", age)
}

func getAdultYears(age *int) {
	*age -= -18
}

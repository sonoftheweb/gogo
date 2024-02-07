package main

import (
	"fmt"
	"github.com/Pallinder/go-randomdata"
	"practice2/structs"
)

func main() {
	hobbies := [3]string{"Gaming", "Reading", "Programming"}
	newList := [2]string{hobbies[1], hobbies[2]}
	sliceFromHobbies := hobbies[0:2]
	//sliceFromHobbies := hobbies[:2]

	fmt.Println(hobbies)
	fmt.Println("First element standalone:", hobbies[0])
	fmt.Println("Second and Third element in a new list:", newList)
	fmt.Println("Slice from hobbies with first and second elements:", sliceFromHobbies)
	fmt.Println("Re-slice from sliceFromHobbies getting the second and last element of original array:", sliceFromHobbies[1:3])

	dynamicArray := []string{"Structs", "Maps"}
	dynamicArray[1] = "Pointers"
	dynamicArray = append(dynamicArray, "Interfaces")

	// creates a new products slice with 2 products
	products := structs.NewProducts()
	for i := 0; i < 2; i++ {
		product := structs.New(1, fmt.Sprintf("Product-%v", i), float64(randomdata.Number(10, 30)))
		products.Add(product)
	}

	// add a third product to the list
	product := structs.New(1, "New-Product-3", float64(randomdata.Number(10, 30)))
	products.Add(product)

	fmt.Println("Products:", *products)
}

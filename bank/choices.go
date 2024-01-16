package main

import "fmt"

func choicesMessage() {
	fmt.Println("What do you want to do today?")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit")
	fmt.Println("3. Withdraw")
	fmt.Println("4. Exit")

	var choice int
	choice = getChoice()
	performChoiceAction(choice)
}

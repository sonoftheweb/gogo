package main

import (
	"bank/fileOps"
	"fmt"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile string = "balance.txt"

var accountBalance = fileOps.GetFloatFromFile(accountBalanceFile)

func main() {
	welcomeMessage()
	choicesMessage()
}

func welcomeMessage() {
	fmt.Println("Hi", randomdata.FullName(randomdata.RandomGender))
	fmt.Println("Welcome to Go Bank!")
	fmt.Println("You can reach us 24/7 at", randomdata.PhoneNumber())
}

func getChoice() int {
	var choice int
	fmt.Print("Enter your choice: ")
	_, err := fmt.Scanln(&choice)
	if err != nil {
		fmt.Println("Invalid choice selected")
		choicesMessage()
	}
	return choice
}

func performChoiceAction(choice int) {
	switch choice {
	case 1:
		checkBalance()
	case 2:
		deposit()
	case 3:
		withdraw()
	case 4:
		exit()
	}
}

func exit() {
	fmt.Println("Bye!")
}

func withdraw() {
	var withdrawAmount float64
	fmt.Print("Your Withdraw: ")
	_, err := fmt.Scanln(&withdrawAmount)
	if err != nil {
		println("Invalid amount")
		withdraw()
	}
	withdrawInCents := dollarToCents(withdrawAmount)
	if withdrawInCents > accountBalance {
		fmt.Println("Insufficient funds")
		withdraw()
	}
	accountBalance -= withdrawInCents
	fileOps.WriteFloatToFile(accountBalanceFile, accountBalance)
	checkBalance()
}

func deposit() {
	var depositAmount float64
	fmt.Print("Your Deposit: ")
	_, err := fmt.Scanln(&depositAmount)
	if err != nil {
		println("Invalid amount")
		deposit()
	}
	depositInCents := dollarToCents(depositAmount)
	accountBalance += depositInCents
	fileOps.WriteFloatToFile(accountBalanceFile, accountBalance)
	checkBalance()
}

func checkBalance() {
	fmt.Printf("Your balance is %.2f dollars\n", centsToDollar(accountBalance))
	choicesMessage()
}

func dollarToCents(dollar float64) int64 {
	return int64(dollar * 100)
}

func centsToDollar(cents int64) float64 {
	return float64(cents) / 100
}

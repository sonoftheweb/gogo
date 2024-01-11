package main

import (
	"fmt"
	"os"
	"strconv"
)

// value in cents
var accountBalance int64 = getBalanceFromFile()

const accountBalanceFile string = "balance.txt"

func main() {
	welcomeMessage()
	choicesMessage()
}

func welcomeMessage() {
	fmt.Println("Welcome to Go Bank!")
}

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
	writeBalanceToFile(accountBalance)
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
	writeBalanceToFile(accountBalance)
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

func writeBalanceToFile(balance int64) {
	balanceText := fmt.Sprint(balance)
	err := os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
	if err != nil {
		return
	}
}

func getBalanceFromFile() int64 {
	data, err := os.ReadFile(accountBalanceFile)
	if err != nil {
		return 0
	}
	balanceText := string(data)
	balance, convError := strconv.ParseInt(balanceText, 64, 10)
	if convError != nil {
		return 0
	}

	return balance
}

package main

import (
	"fmt"
	"structs/structs"
)

var user *structs.User
var admin structs.Admin

func main() {
	email := getUserData("Please enter your email: ")
	password := getUserData("Please enter your password: ")

	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var err error = nil
	admin, err = structs.NewAdmin(&email, &password)
	if err != nil {
		fmt.Println(err)
	}

	user, err = structs.NewUser(&firstName, &lastName, &birthdate)
	if err != nil {
		fmt.Println(err)
	}

	admin.User = *user

	if user != nil {
		user.OutputUserData()
		user.SetUserAge("30")
		user.OutputUserData()
	} else {
		return
	}
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	_, err := fmt.Scanln(&value)
	if err != nil {
		return ""
	}
	return value
}

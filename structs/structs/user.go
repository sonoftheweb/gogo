package structs

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	age       string
	createdAt time.Time
}

func NewUser(firstName, lastName, birthdate *string) (*User, error) {
	if *firstName == "" || *lastName == "" || *birthdate == "" {
		return nil, errors.New("please enter all required fields")
	}

	return &User{
		firstName: *firstName,
		lastName:  *lastName,
		birthdate: *birthdate,
		createdAt: time.Now(),
	}, nil
}

func (user *User) OutputUserData() {
	// fmt.Println((*user).firstName, (*user).lastName, (*user).birthdate, (*user).createdAt) // This is the same as the line below
	fmt.Println(user.firstName, user.lastName, user.birthdate, user.age, user.createdAt)
	fmt.Println(*user)
}

func (user *User) SetUserAge(age string) {
	user.age = age
}

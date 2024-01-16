package structs

import (
	"errors"
)

type Admin struct {
	email    string
	password string
	User     User
}

func NewAdmin(email, password *string) (Admin, error) {
	if *email == "" || *password == "" {
		return Admin{}, errors.New("please enter all required fields")
	}

	return Admin{
		email:    *email,
		password: *password,
		User:     User{},
	}, nil
}

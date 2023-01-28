package ch

import "net/mail"

type User struct {
	FirstName string
	LastName  string
	Email     mail.Address
	Password  string
}

func NewUser(firstName string, lastName string, emailStr string, password string) (*User, error) {
	email, err := mail.ParseAddress(emailStr)
	if err != nil {
		return nil, err
	}

	return &User{
		FirstName: firstName, LastName: lastName,
		Email: *email, Password: password,
	}, nil
}

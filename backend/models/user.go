package models

import (
	uuid "github.com/satori/go.uuid"
	"net/mail"
)

type User struct {
	ID uuid.UUID `gorm:"type:uuid;primary_key"`

	Username  string `gorm:"index:username_idx,unique"`
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Email     string `gorm:"index:username_idx,unique"`
	Password  string `gorm:"not null"`
}

func NewUser(username string, firstName string, lastName string, emailStr string, password string) (*User, error) {
	email, err := mail.ParseAddress(emailStr)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:       uuid.NewV4(),
		Username: username, FirstName: firstName, LastName: lastName,
		Email: email.Address, Password: password,
	}, nil
}

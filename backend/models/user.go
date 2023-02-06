package models

import (
	"CardHero/db"
	uuid "github.com/satori/go.uuid"
	"net/mail"
)

type User struct {
	ID uuid.UUID `gorm:"type:uuid;primary_key"`

	Username  string `gorm:"index:username_idx,unique"`
	FirstName string
	LastName  string
	Email     string `gorm:"index:username_idx,unique"`
	Password  string
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

func FetchUser(username string) (*User, error) {
	conn := db.GetConn()

	var user User
	err := conn.Find(&user, "username = ?", username).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

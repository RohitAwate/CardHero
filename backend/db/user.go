package db

import (
	"CardHero/models"
	"fmt"
	uuid "github.com/satori/go.uuid"
)

func GetUserByUsername(username string) (*models.User, error) {
	conn := getConn()

	var user models.User
	err := conn.Find(&user, "username = ?", username).Error
	if err != nil {
		return nil, err
	}

	if user.ID == uuid.Nil {
		return nil, fmt.Errorf("user not found with username: %s", username)
	}

	return &user, nil
}

func GetUserByEmail(email string) (*models.User, error) {
	conn := getConn()

	var user models.User
	err := conn.Find(&user, "email = ?", email).Error
	if err != nil {
		return nil, err
	}

	if user.ID == uuid.Nil {
		return nil, fmt.Errorf("user not found with email: %s", email)
	}

	return &user, nil
}

func SaveUser(user models.User) {
	conn := getConn()
	conn.Create(&user)
}

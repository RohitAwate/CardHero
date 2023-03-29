package db

import (
	"CardHero/models"
	"fmt"
	uuid "github.com/satori/go.uuid"
)

func GetUser(username string) (*models.User, error) {
	conn := getConn()

	var user models.User
	err := conn.Find(&user, "username = ?", username).Error
	if err != nil {
		return nil, err
	}

	if user.ID == uuid.Nil {
		return nil, fmt.Errorf("user not found with username: " + username)
	}

	return &user, nil
}

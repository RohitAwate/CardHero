package db

import "CardHero/models"

func GetUser(username string) (*models.User, error) {
	conn := getConn()

	var user models.User
	err := conn.Find(&user, "username = ?", username).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

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

const (
	getUserByLoginCredentialsQuery = `
		SELECT * FROM users
		WHERE username = ?
		AND password = crypt(?, password);
	`
)

func GetUserByLoginCredentials(username, password string) (*models.User, error) {
	conn := getConn()

	var user models.User
	err := conn.Raw(getUserByLoginCredentialsQuery, username, password).Scan(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

const (
	savePasswordQuery = `
		UPDATE users
		SET password = crypt(?, gen_salt('bf'))
		WHERE username = ?;
	`
)

func SaveUser(user models.User) error {
	// Remove the password from the user since we will be hashing + salting it
	password := user.Password
	user.Password = ""

	// Create user without password
	conn := getConn()
	err := conn.Create(&user).Error
	if err != nil {
		return err
	}

	// Hash and save the password
	err = conn.Exec(savePasswordQuery, password, user.Username).Error
	if err != nil {
		return err
	}

	return nil
}

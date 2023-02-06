package db

import "CardHero/models"

func GetCardsBy(user models.User) ([]models.Card, error) {
	conn := getConn()

	var cards []models.Card
	err := conn.Find(&cards, "owner_id = ?", user.ID).Error
	if err != nil {
		return nil, err
	}

	return cards, nil
}

func SaveCard(card models.Card) {
	conn := getConn()
	conn.Create(&card)
}

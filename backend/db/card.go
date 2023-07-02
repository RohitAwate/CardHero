package db

import (
	"CardHero/models"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm/clause"
)

func GetCardsBy(user models.User) ([]models.Card, error) {
	conn := getConn()

	var cards []models.Card
	err := conn.Preload(clause.Associations).Order("timestamp").Find(&cards, "owner_id = ?", user.ID).Error
	if err != nil {
		return nil, err
	}

	return cards, nil
}

func GetCardByID(user models.User, cardID uuid.UUID) (*models.Card, error) {
	conn := getConn()

	var card models.Card
	err := conn.Preload(clause.Associations).Order("timestamp").Find(&card, "owner_id = ? and id = ?", user.ID, cardID).Error
	if err != nil {
		return nil, err
	}

	return &card, nil
}

func SaveCard(card models.Card) {
	conn := getConn()
	conn.Create(&card)
}

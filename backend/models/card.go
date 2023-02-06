package models

import (
	"CardHero/db"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Card struct {
	ID        uuid.UUID `json:"id,omitempty" gorm:"type:uuid;primary_key"`
	Timestamp time.Time `json:"timestamp"`
	OwnerID   uuid.UUID `json:"-"`
	Owner     User      `json:"-"`
	Contents  string    `json:"contents,omitempty"`
}

func NewCard(owner User, contents string, timestamp time.Time) Card {
	return Card{
		ID:        uuid.NewV4(),
		Timestamp: timestamp,
		Owner:     owner,
		Contents:  contents,
	}
}

func GetCardsBy(user User) ([]Card, error) {
	conn := db.GetConn()

	var cards []Card
	err := conn.Find(&cards, "owner_id = ?", user.ID).Error
	if err != nil {
		return nil, err
	}

	return cards, nil
}

func SaveCard(card Card) {
	conn := db.GetConn()
	conn.Create(&card)
}

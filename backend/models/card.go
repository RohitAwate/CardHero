package models

import (
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

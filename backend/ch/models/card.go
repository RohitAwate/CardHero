package models

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type Card struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Timestamp time.Time `json:"timestamp"`
	Owner     User      `json:"-"`
	Contents  string    `json:"contents,omitempty"`
}

func NewCard(owner User, contents string) Card {
	return Card{
		ID:        uuid.NewV4(),
		Timestamp: time.Now(),
		Owner:     owner,
		Contents:  contents,
	}
}

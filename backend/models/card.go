package models

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type Card struct {
	ID        uuid.UUID `json:"id,omitempty" gorm:"type:uuid;primary_key"`
	Timestamp time.Time `json:"timestamp" gorm:"not null"`
	OwnerID   uuid.UUID `json:"-" gorm:"not null"`
	Owner     User      `json:"-"`
	Contents  string    `json:"contents,omitempty"`
	FolderID  uuid.UUID `json:"-" gorm:"not null"`
	Folder    Folder    `json:"-"`
}

func NewCard(owner User, contents string, timestamp time.Time) Card {
	return Card{
		ID:        uuid.NewV4(),
		Timestamp: timestamp,
		OwnerID:   owner.ID,
		Owner:     owner,
		Contents:  contents,
	}
}

func (c *Card) AssignFolder(folder Folder) {
	c.FolderID = folder.ID
	c.Folder = Folder{}
}

package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"time"
)

type Card struct {
	ID        uuid.UUID `json:"id,omitempty" gorm:"type:uuid;primary_key"`
	Timestamp time.Time `json:"timestamp" gorm:"not null"`

	OwnerID  uuid.UUID `json:"-" gorm:"not null"`
	Owner    User      `json:"-"`
	FolderID uuid.UUID `json:"-" gorm:"not null"`
	Folder   Folder    `json:"-"`

	Contents      string `json:"contents,omitempty"`
	ContentsIndex string `gorm:"type:tsvector"`
}

func SetupSearchIndexTrigger(conn *gorm.DB) {
	// Setup full-text search trigger
	query := `
		CREATE OR REPLACE FUNCTION card_tsvector_trigger() RETURNS trigger AS $$
		begin
		  -- Replace all non-word characters with a space
		  new.contents := regexp_replace(new.contents, '\W+', ' ', 'g');
		  new.contents_index := to_tsvector(new.contents);
		  return new;
		end
		$$ LANGUAGE plpgsql;

		CREATE TRIGGER cards_search_index_update BEFORE INSERT OR UPDATE
		ON cards FOR EACH ROW EXECUTE PROCEDURE card_tsvector_trigger();
	`

	if err := conn.Exec(query).Error; err != nil {
		panic(err)
	}
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

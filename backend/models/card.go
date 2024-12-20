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
	Owner    User      `json:"-" gorm:"constraint:OnDelete:CASCADE"`
	FolderID uuid.UUID `json:"-" gorm:"not null"`
	Folder   Folder    `json:"-" gorm:"constraint:OnDelete:CASCADE"`

	Contents       string `json:"contents,omitempty"`
	SearchContents string `json:"-" gorm:"type:tsvector;index"`
}

func SetupSearchIndexTrigger(conn *gorm.DB) {
	// Setup full-text search trigger
	query := `
		CREATE INDEX IF NOT EXISTS search_contents_index
			ON cards
			USING GIN (search_contents);

		CREATE OR REPLACE FUNCTION card_tsvector_trigger() RETURNS trigger AS $$
		DECLARE
			processed_contents TEXT;
		begin
			-- Replace all non-word characters with a space
			processed_contents := regexp_replace(new.contents, '\W+', ' ', 'g');
			new.search_contents := to_tsvector(processed_contents);
			return new;
		end
		$$ LANGUAGE plpgsql;

		CREATE OR REPLACE TRIGGER cards_search_index_update
			BEFORE INSERT OR UPDATE
			ON cards FOR EACH ROW
			EXECUTE PROCEDURE card_tsvector_trigger();
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

package db

import "CardHero/models"

const (
	searchCardsQuery = `
		SELECT id, timestamp, contents
		FROM cards
		WHERE owner_id = ? AND
		    search_contents @@ plainto_tsquery(?);
	`
)

func Search(query string, user models.User) ([]models.Card, error) {
	conn := getConn()

	cards := make([]models.Card, 0)
	err := conn.Raw(searchCardsQuery, user.ID, query).Scan(&cards).Error
	if err != nil {
		return cards, err
	}

	return cards, nil
}

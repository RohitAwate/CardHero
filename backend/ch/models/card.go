package models

type Card struct {
	Owner    User   `json:"-"`
	Contents string `json:"contents,omitempty"`
}

func NewCard(owner User, contents string) Card {
	return Card{Owner: owner, Contents: contents}
}

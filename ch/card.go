package ch

type Card struct {
	Owner    User
	Contents string
}

func NewCard(owner User, contents string) Card {
	return Card{Owner: owner, Contents: contents}
}

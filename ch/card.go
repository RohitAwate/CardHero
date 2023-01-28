package ch

type Card struct {
	Owner    User
	Contents string
}

func newCard(owner User, contents string) Card {
	return Card{Owner: owner, Contents: contents}
}

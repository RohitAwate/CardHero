package ch

import "fmt"

type Log struct {
	Owner User
	Cards []*Card
}

func NewLog(user User) Log {
	return Log{Owner: user}
}

func (log *Log) Append(card *Card) {
	log.Cards = append(log.Cards, card)
	fmt.Println(log.Cards)
}

func (log *Log) Size() int {
	return len(log.Cards)
}

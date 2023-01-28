package models

import (
	"encoding/json"
)

type CardLog struct {
	Owner User
	cards []*Card
}

func NewCardLog(user User) CardLog {
	return CardLog{Owner: user}
}

func (clog *CardLog) Append(cards ...*Card) {
	clog.cards = append(clog.cards, cards...)
}

func (clog *CardLog) Size() int {
	return len(clog.cards)
}

func (clog *CardLog) JSON() ([]byte, error) {
	bytes, err := json.Marshal(clog.cards)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

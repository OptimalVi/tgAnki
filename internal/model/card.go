package model

import "time"

type Card struct {
	ID        int       `json:"id"`
	deckID    int       `json:"deckId`
	frontText string    `json:"frontText"`
	backText  string    `json:"backText"`
	createdAt time.Time `json:"createdAt"`
	checkAt   time.Time `json:"checkAt"`
	repeatAt  time.Time `json:"repeatAt"`
}

package model

import "time"

type Card struct {
	ID        int       `json:"id"`
	DeckID    int       `json:"deckId`
	FrontText string    `json:"frontText"`
	BackText  string    `json:"backText"`
	createdAt time.Time `json:"createdAt"`
	checkAt   time.Time `json:"checkAt"`
	repeatAt  time.Time `json:"repeatAt"`
}

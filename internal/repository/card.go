package repository

import "optimal_vi/tg_anki/internal/model"

type ICardRepository interface {
	InsertCard(deckId int, front string, back string) error
	GetCardsByDeckId(deckId int) ([]*model.Card, error)
}
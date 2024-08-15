package repository

import (
	"database/sql"
	"optimal_vi/tg_anki/internal/model"
)

type IDeckRepository interface {
	GetDeck(id int) (*model.Deck, error)
	GetDecksByChatID(chatId int64) ([]model.Deck, error)
	InsertDeck(chatId int64, name string) (sql.Result, error)
	DeleteDeck(chatId int64, id int) error
}

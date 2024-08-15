package sqlite

import (
	"database/sql"
	"optimal_vi/tg_anki/internal/model"
)

type CardRepository struct {
	DB *sql.DB
}

func (r *CardRepository) InsertCard(deckId int, front string, back string) error {
	return nil
}
func (r *CardRepository) GetCardsByDeckId(deckId int) ([]*model.Card, error) {
	return nil, nil
}

func (c *CardRepository) GetCards(deckId int64) ([]model.Card, error) {
	rows, err := c.DB.Query("SELECT id, question, answer FROM cards WHERE deckId =?", deckId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cards := make([]model.Card, 0)
	for rows.Next() {
		card := model.Card{}
		err := rows.Scan(&card)
		if err != nil {
			return nil, err
		}
		cards = append(cards, card)
	}

	return cards, nil
}

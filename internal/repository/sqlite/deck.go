package sqlite

import (
	"database/sql"
	"log"
	"optimal_vi/tg_anki/internal/model"
)

type DeckRepository struct {
	DB *sql.DB
}

func (r *DeckRepository) GetDeck(id int) (*model.Deck, error) {
	row := r.DB.QueryRow("SELECT id, name, chatId, createdAt FROM decks WHERE id =?", id)
	deck := model.Deck{}
	err := row.Scan(&deck.ID, &deck.Name, &deck.ChatID, &deck.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &deck, err
}

func (r *DeckRepository) GetDecksByChatID(chatId int64) ([]model.Deck, error) {
	rows, err := r.DB.Query("SELECT id, name, chatId, createdAt FROM decks WHERE chatId =?", chatId)
	if err != nil {

		return nil, err
	}
	defer rows.Close()

	decks := make([]model.Deck, 0)
	for rows.Next() {
		deck := model.Deck{}
		err := rows.Scan(&deck.ID, &deck.Name, &deck.ChatID, &deck.CreatedAt)
		if err != nil {
			log.Println("Error r-s-d-26", err)
			return nil, err
		}
		decks = append(decks, deck)
	}
	return decks, nil
}

func (r *DeckRepository) InsertDeck(chatId int64, name string) (sql.Result, error) {
	stmp, err := r.DB.Prepare("INSERT INTO decks (name, chatId) VALUES (?,?)")
	if err != nil {
		return nil, err
	}
	if res, err := stmp.Exec(name, chatId); err == nil {
		return res, nil
	}
	return nil, err
}

func (r *DeckRepository) DeleteDeck(chatId int64, deckId int) error {
	_, err := r.DB.Exec("DELETE FROM decks WHERE chatId =? AND id =?", chatId, deckId)
	return err
}

func (r *DeckRepository) UpdateName(deckId int, name string) error {
	smtp, err := r.DB.Prepare("UPDATE decks SET name =? WHERE id =?")
	if err != nil {
		return err
	}
	_, err = smtp.Exec(name, deckId)
	return err
}

package controller

import (
	"log"
	"optimal_vi/tg_anki/internal/repository"
	"optimal_vi/tg_anki/internal/repository/sqlite"
	"optimal_vi/tg_anki/internal/view"
	"optimal_vi/tg_anki/pkg/db"
	"optimal_vi/tg_anki/pkg/tg"
)

var (
	deckRepository repository.IDeckRepository = &sqlite.DeckRepository{DB: db.GetDB()}
	cardRepository repository.ICardRepository = &sqlite.CardRepository{DB: db.GetDB()}
)

func HomeController(chat *tg.ChatContext) error {
	var err error

	decks, err := deckRepository.GetDecksByChatID(chat.ID)
	if err != nil {
		log.Print(err)
		return err
	}

	if len(decks) > 0 {
		_, err = tg.Send(view.TgViewDecksListButton(chat.ID, decks))
		if err != nil {
			log.Print(err)
			return err
		}
	}

	chat.SetAction(tg.ChatContextHome)
	return nil
}

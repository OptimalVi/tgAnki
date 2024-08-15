package controller

import (
	"log"
	"optimal_vi/tg_anki/internal/bot/view"
	"optimal_vi/tg_anki/internal/repository"
	"optimal_vi/tg_anki/internal/repository/sqlite"
	"optimal_vi/tg_anki/pkg/db"
	"optimal_vi/tg_anki/pkg/tg"
)

var (
	deckRepository repository.IDeckRepository = &sqlite.DeckRepository{DB: db.GetDB()}
	cardRepository repository.ICardRepository = &sqlite.CardRepository{DB: db.GetDB()}
)

func HomeController(chat *tg.ChatContext, udp *tg.BotUpdate) error {
	var err error

	decks, err := deckRepository.GetDecksByChatID(chat.ID)
	if err != nil {
		log.Print(err)
		return err
	}

	if len(decks) > 0 {
		_, err = tg.Send(view.DecksListView(chat.ID, decks))
		if err != nil {
			log.Print(err)
			return err
		}
	}

	chat.SetAction(tg.ChatContextHome)
	return nil
}

// func OpenDeckController(chat *tg.ChatContext, dockId int) error {
// 	fmt.Printf("\nOpenDeckController triggered for chat ID %d, deck ID %d\n", chat.ID, dockId)
// 	deck, err := deckRepository.GetDeck(dockId)
// 	if err != nil {
// 		log.Printf("Error b-c-44", err)
// 		return err
// 	}
// 	if deck == nil {
// 		fmt.Printf("Deck not found for chat ID %d, deck ID %d\n", chat.ID, dockId)
// 		return fmt.Errorf("Deck not found")
// 	}

// 	view := view.OpenDeckView(chat.ID, deck)
// 	if _, err := tg.Send(view); err != nil {
// 		fmt.Print("Error sending message b-c-48: ", err)
// 		return err
// 	}

// 	chat.SetData("deck", deck)
// 	chat.SetAction(tg.ChatContextDeck)
// 	return nil
// }

// func CreateDeckController(chat *tg.ChatContext, udp *tg.BotUpdate) error {
// 	fmt.Printf("\nCreateDeckController triggered for chat ID %d\n", chat.ID)

// 	view := view.CreateDeckWriteNameView(chat.ID)
// 	if _, err := tg.Send(view); err != nil {
// 		fmt.Print("Error sending message b-c-53: ", err)
// 	}

// 	chat.SetAction(tg.ChatContextDeckCreate)
// 	return nil
// }

// func SaveDeckController(chat *tg.ChatContext, deckName string) error {
// 	fmt.Printf("\nSaveDeckController triggered for chat ID %d, deck name '%s'\n", chat.ID, deckName)

// 	deckRepository.InsertDeck(chat.ID, deckName)

// 	HomeController(chat, udp)
// 	return nil
// }

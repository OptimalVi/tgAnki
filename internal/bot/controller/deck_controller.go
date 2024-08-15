package controller

import (
	"fmt"
	"optimal_vi/tg_anki/internal/bot/view"
	"optimal_vi/tg_anki/internal/model"
	"optimal_vi/tg_anki/pkg/tg"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func DeckCreateController(chat *tg.ChatContext, udp *tg.BotUpdate) error {
	name := udp.Message.Text

	_, err := deckRepository.InsertDeck(chat.ID, name)
	if err != nil {
		return err
	}

	tg.Send(tgbotapi.NewMessage(chat.ID, "Deck created successfully: "+name))

	return HomeController(chat, udp)
}

func DeckController(chat *tg.ChatContext, udp *tg.BotUpdate) error {
	data := udp.CallbackQuery.Data
	deck, ok := chat.GetData("deck").(model.Deck)
	if ok {
		fmt.Printf("DeckController: %s\n", data)
		return tg.InvalidMessageData{}
	}

	// if deck == nil {
	// 	HomeController(chat, udp)
	// 	return fmt.Errorf("Deck not found in chat context")
	// }

	switch data {
	case "drop-deck":
		view := view.DeckDropView(chat.ID, deck.Name)
		_, err := tg.Send(view)
		if err != nil {
			return err
		}
		chat.SetAction(tg.ChatContextDeckDrop)
	default:
		return fmt.Errorf("Invalid DeckController data: %s", data)
	}

	return nil
}

func DeckDropController(chat *tg.ChatContext, udp *tg.BotUpdate) error {
	data := udp.CallbackQuery.Data
	fmt.Printf("INFO> DeckDropController: %s\n", data)
	deck, ok := chat.GetData("deck").(*model.Deck)
	if ok {
		fmt.Printf("DeckDropController: %s\n", data)
		return tg.InvalidMessageData{}
	}
	fmt.Println("INFO> DeckDropController:", deck)

	switch data {
	case "confirm":
		err := deckRepository.DeleteDeck(chat.ID, deck.ID)
		if err != nil {
			return err
		}
		_, err = tg.Send(tgbotapi.NewMessage(chat.ID, "Deck deleted successfully: "+deck.Name))
		if err != nil {
			return err
		}

		chat.SetData("deck", nil)
		return HomeController(chat, udp)
	case "back":
		view := view.OpenDeckView(chat.ID, deck)
		_, err := tg.Send(view)
		if err != nil {
			return err
		}
		chat.SetAction(tg.ChatContextDeck)
	default:
		return fmt.Errorf("Invalid DeckDropController data: %v", data)
	}

	return nil
}

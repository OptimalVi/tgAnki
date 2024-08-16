package controller

import (
	"fmt"
	"optimal_vi/tg_anki/internal/bot/state"
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

	switch data {
	case "add-card":
		_, err := tg.Send(
			view.AddCardView(chat.ID, deck.Name, state.CardCreateFront),
		)
		if err != nil {
			return err
		}
		chat.SetAction(tg.ChatContextCardCreate)
		chat.SetData("card-create", state.CardCreateFront)
		return nil
	case "review-cards":
	case "edit-deck":
		view := view.DeckEditView(chat.ID, deck.Name)
		_, err := tg.Send(view)
		if err != nil {
			return err
		}
		chat.SetAction(tg.ChatContextDeckEdit)

	case "drop-deck":
		view := view.DeckDropView(chat.ID, deck.Name)
		_, err := tg.Send(view)
		if err != nil {
			return err
		}
		chat.SetAction(tg.ChatContextDeckDrop)

	case "home":
		chat.SetData("deck", nil)
		return HomeController(chat, udp)

	default:
		return fmt.Errorf("Invalid DeckController data: %s", data)
	}

	return nil
}

func DeckDropController(chat *tg.ChatContext, udp *tg.BotUpdate) error {
	data := udp.CallbackQuery.Data
	deck, ok := chat.GetData("deck").(*model.Deck)
	if !ok {
		return fmt.Errorf("Deck not found in chat context")
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

func DeckEditController(chat *tg.ChatContext, udp *tg.BotUpdate) error {
	tp, err := udp.GetType()
	if err != nil {
		return fmt.Errorf("ERROR> GetType: %v", err)
	}
	deck, ok := chat.GetData("deck").(*model.Deck)
	if !ok {
		return fmt.Errorf("Deck not found in chat context")
	}

	switch tp {
	case tg.BotUpdateText:
		newName := udp.Message.Text
		err := deckRepository.UpdateName(deck.ID, newName)
		if err != nil {
			return err
		}
		_, err = tg.Send(tgbotapi.NewMessage(chat.ID, "Deck updated successfully: "+newName))
		if err != nil {
			return err
		}

		deck, err = deckRepository.GetDeck(deck.ID)
		if err != nil {
			return err
		}

		chat.SetData("deck", deck)
		chat.SetAction(tg.ChatContextDeck)
		view := view.OpenDeckView(chat.ID, deck)
		_, err = tg.Send(view)
		return err

	case tg.BotUpdateCallbackQuery:
		view := view.OpenDeckView(chat.ID, deck)
		_, err := tg.Send(view)
		chat.SetAction(tg.ChatContextDeck)
		return err
	}

	return fmt.Errorf("ERROR> Invalid DeckEditController type: %v", tp)
}

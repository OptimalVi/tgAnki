package controller

import (
	"fmt"
	"optimal_vi/tg_anki/internal/bot/view"
	"optimal_vi/tg_anki/pkg/tg"
	"strconv"
)

func StartController(chat *tg.ChatContext, udp *tg.BotUpdate) error {
	fmt.Printf("\nStart controller triggered for chat ID %d\n", chat.ID)

	data := udp.CallbackQuery.Data

	if data == "create-deck" {
		view := view.CreateDeckWriteNameView(chat.ID)
		if _, err := tg.Send(view); err != nil {
			return fmt.Errorf("Error sending message StartController create-deck: %w", err)
		}

		chat.SetAction(tg.ChatContextDeckCreate)

	} else {
		deckId, err := strconv.Atoi(data)
		if err != nil {
			return fmt.Errorf("Invalid StartController data: %w", err)
		}

		deck, err := deckRepository.GetDeck(deckId)
		if err != nil {
			return fmt.Errorf("Error getting deck by id %d: %w", deckId, err)
		}

		view := view.OpenDeckView(deck.ChatID, deck)
		if _, err := tg.Send(view); err != nil {
			return err
		}

		chat.SetData("deck", deck)
		chat.SetAction(tg.ChatContextDeck)
	}

	return nil
}

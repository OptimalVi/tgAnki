package handler

import (
	"fmt"
	"optimal_vi/tg_anki/pkg/tg"
)

func DeckHandler(chat *tg.ChatContext, udp tg.BotUpdate) error {
	fmt.Printf("\nDeckHandler triggered for chat ID %d\n", chat.ID)
	return nil
}

// func DeckCreateHandler(chat *tg.ChatContext, udp tg.BotUpdate) error {
// 	fmt.Printf("\nDeckCreateHandler triggered for chat ID %d\n", chat.ID)

// 	if udp.Message.Text == "" {
// 		return tg.InvalidMessage{}
// 	}

// 	deckName := udp.Message.Text

// 	// controller.SaveDeckController(chat, deckName)

// 	return nil
// }

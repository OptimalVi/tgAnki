package view

import (
	"fmt"
	"optimal_vi/tg_anki/internal/bot/state"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func AddCardView(chatId int64, deckName string, st int) *tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(chatId, "")
	switch st {
	case state.CardCreateFront:
		msg.Text = fmt.Sprintf("Creating new card for deck: %s\nWrite front view:", deckName)
	case state.CardCreateBack:
		msg.Text = "Write back view:"
	}

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData(
			"Cancel", "cancel",
		)),
	)

	return &msg
}

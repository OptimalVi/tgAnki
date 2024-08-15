package view

import (
	"optimal_vi/tg_anki/internal/model"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func TgViewCardList(msg tgbotapi.MessageConfig, cards []model.Card) tgbotapi.MessageConfig {
	msg.Text = "Cards:"
	cardsButtons := make([][]tgbotapi.InlineKeyboardButton, 0, len(cards))

	// for _, card := range cards {
		// row := tgbotapi.NewInlineKeyboardRow(
		// 	tgbotapi.NewInlineKeyboardButtonData(
		// 		card.FrontText,
		// 		strconv.Itoa(card.ID),
		// 	),
		// )
		// cardsButtons = append(cardsButtons, row)
	// }

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(cardsButtons...)
	return msg
}

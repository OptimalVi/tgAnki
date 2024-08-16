package view

import (
	"fmt"
	"optimal_vi/tg_anki/internal/model"
	"optimal_vi/tg_anki/pkg/tg"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func TgViewAddDeckButton(chat *tg.ChatContext) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(chat.ID, "Decks:")
	msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Add new deck"),
		),
	)
	return msg
}

func CreateDeckWriteNameView(chatId int64) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(chatId, "Enter the name of the new deck:")
}

func TgViewCreateDeckSuccess(chatId int64, deckName string) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(
		chatId,
		fmt.Sprintf("Deck '%s' created successfully!", deckName),
	)
}

func DecksListView(chatId int64, decks []model.Deck) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(chatId, "Decks:")
	decksButtons := make([][]tgbotapi.InlineKeyboardButton, 0, len(decks))

	for _, deck := range decks {
		row := tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				deck.Name,
				strconv.Itoa(deck.ID),
			),
		)
		decksButtons = append(decksButtons, row)
	}

	decksButtons = append(decksButtons, tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(
			"🖋️ Create new deck",
			"create-deck",
		),
	))

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(decksButtons...)

	return msg
}

func OpenDeckView(chatId int64, deck *model.Deck) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(chatId, fmt.Sprintf("Deck: %s", deck.Name))

	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData(
			"👟 Treining", "training",
		)),
		tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData(
			"📑 Card list", "show-cards",
		)),
		tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData(
			"🖋️ Add card", "add-card",
		)),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"🗑️ Drop deck", "drop-deck",
			),
			tgbotapi.NewInlineKeyboardButtonData(
				"✍️ Edit deck", "edit-deck",
			),
		),
		tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData(
			"🏚️ Return to home", "home",
		)),
	)

	msg.ReplyMarkup = keyboard
	return msg
}

func DeckDropView(chatId int64, deckName string) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(chatId, fmt.Sprintf("Are you sure you want delete deck '%s'?", deckName))
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Yes, delete", "confirm"),
			tgbotapi.NewInlineKeyboardButtonData("Cancel", "back"),
		),
	)

	return msg
}

func DeckEditView(chatId int64, deckName string) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(chatId, fmt.Sprintf("Enter the new name for deck '%s':", deckName))

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Cancel", "back"),
		))

	return msg
}

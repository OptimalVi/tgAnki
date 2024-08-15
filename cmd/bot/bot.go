package main

import (
	"log"
	"optimal_vi/tg_anki/internal/bot/handler"
	"optimal_vi/tg_anki/pkg/conf"
	"optimal_vi/tg_anki/pkg/db"
	"optimal_vi/tg_anki/pkg/tg"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	chatContexts = make(map[int64]*tg.ChatContext)
)

func main() {
	db.Migrate()
	bot := tg.GetBot()

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		var err error

		fromChat := update.FromChat()
		if fromChat == nil {
			log.Printf("ERROR> Not defined chat for update: %v\n", update)
			continue
		}

		chat, ok := chatContexts[fromChat.ID]
		if !ok {
			chat = tg.NewChatContext(fromChat.ID)
			chatContexts[fromChat.ID] = chat
		}
		log.Printf("> Action: %d\n", chat.GetAction())

		udp := &tg.BotUpdate{Update: update}

		err = handler.Handler(chat, udp)

		if err != nil {
			log.Printf("ERROR> Handler returned error: %v\n", err)
			continue
		}
	}
}

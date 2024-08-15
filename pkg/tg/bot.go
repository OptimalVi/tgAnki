package tg

import (
	"optimal_vi/tg_anki/pkg/conf"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	bot  *tgbotapi.BotAPI
	once sync.Once
)

func GetBot() *tgbotapi.BotAPI {
	once.Do(func() {
		bot, _ = tgbotapi.NewBotAPI(conf.GetConfig().BotToken)
		bot.Debug = true
	})
	return bot
}

func Send(msg tgbotapi.Chattable) (tgbotapi.Message, error) {
	return bot.Send(msg)
}

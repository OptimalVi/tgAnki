package tg

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	BotUpdateUndefined = iota
	BotUpdateText
	BotUpdateCommand
	BotUpdateCallbackQuery
	BotUpdateAny
)

type BotUpdate struct {
	tgbotapi.Update
	msgType int
}

func (b *BotUpdate) GetType() (int, error) {
	if b.msgType > BotUpdateUndefined {
		return b.msgType, nil
	}

	switch {
	case b.Message != nil && b.Message.IsCommand():
		b.msgType = BotUpdateCommand
	case b.CallbackQuery != nil:
		b.msgType = BotUpdateCallbackQuery
	case b.Message != nil && b.Message.Text != "":
		b.msgType = BotUpdateText
	}

	if b.msgType > 0 {
		return b.msgType, nil
	}
	return -1, UnknownMessageType{}
}

func (b *BotUpdate) TypeIs(BotUpdateType int) (bool, error) {
	tp, err := b.GetType()
	if err != nil {
		return false, err
	}

	return (tp == BotUpdateType), nil
}

package tg

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

const (
	ChatContextNew = iota
	ChatContextHome
	ChatContextDeckCreate
	ChatContextDeck
	ChatContextDeckEdit
	ChatContextDeckDrop
	ChatContextCardCreate
	ChatContextCardEdit
	ChatContextTraining
)

type ChatContext struct {
	ID     int64
	Udp    *tgbotapi.Update
	action int
	data   map[string]interface{}
}

func NewChatContext(chatId int64) *ChatContext {
	return &ChatContext{
		ID:     chatId,
		action: ChatContextNew,
	}
}

func (c *ChatContext) SetAction(action int) {
	c.action = action
}

func (c *ChatContext) GetAction() int {
	return c.action
}

func (c *ChatContext) ActionIs(action int) bool {
	return c.action == action
}

func (c *ChatContext) SetData(key string, value interface{}) {
	if c.data == nil {
		c.data = make(map[string]interface{})
	}
	c.data[key] = value
}

func (c *ChatContext) GetData(key string) interface{} {
	return c.data[key]
}

func (c *ChatContext) ResetData() {
	c.data = make(map[string]interface{})
}

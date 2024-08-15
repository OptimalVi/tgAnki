package tg

import (
	"log"
	"reflect"
	"runtime"
)

type BotUpdateHandleData struct {
	Action     int
	UpdateType int
	Controller ControllerFunc
}

type ControllerFunc func(*ChatContext, *BotUpdate) error

func (c BotUpdateHandleData) CallControllerFunc(chat *ChatContext, udp *BotUpdate) error {
	log.Printf(
		"\nINFO> Matched handler Action %d, UpdateType %d, Controller %s\n",
		c.Action,
		c.UpdateType,
		runtime.FuncForPC(reflect.ValueOf(c.Controller).Pointer()).Name(),
	)

	return c.Controller(chat, udp)
}

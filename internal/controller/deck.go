package controller

import (
	"fmt"
	"log"
	"optimal_vi/tg_anki/internal/view"
	"optimal_vi/tg_anki/pkg/tg"
	"strconv"
)

var (
	err error
)

func CreateDeckWriteName(ch *tg.ChatContext) error {

	msg := view.TgViewCreateDeckWriteName(ch.ID)

	if _, err := tg.Send(msg); err != nil {
		log.Print("Error c-d-15", err)
		return err
	}

	ch.SetAction(tg.ChatContextDeckCreate)

	return nil
}
func CreateDeck(ch *tg.ChatContext, text string) error {
	deckRepository.InsertDeck(ch.ID, text)
	_, err := tg.Send(view.TgViewCreateDeckSuccess(ch.ID, text))
	if err != nil {
		log.Print("Error c-d-27", err)
	}

	HomeController(ch)
	return nil
}

func OpenDeck(ch *tg.ChatContext, data string) error {
	var deckId int

	if deckId, err = strconv.Atoi(data); err != nil {
		log.Print("Error c-d-37", err)
		return err
	}

	fmt.Println("Opening deck with ID:", deckId)
	return nil
}

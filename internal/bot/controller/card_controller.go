package controller

import (
	"fmt"
	"optimal_vi/tg_anki/internal/bot/state"
	"optimal_vi/tg_anki/internal/bot/view"
	"optimal_vi/tg_anki/internal/model"
	"optimal_vi/tg_anki/pkg/tg"
)

func CardCreateController(chat *tg.ChatContext, udp *tg.BotUpdate) error {
	deck := chat.GetData("deck").(*model.Deck)
	tp, err := udp.GetType()
	if err != nil {
		return err
	}
	switch tp {
	case tg.BotUpdateCallbackQuery:
		if udp.CallbackQuery.Data == "cancel" {
			_, err := tg.Send(view.OpenDeckView(chat.ID, deck))
			if err != nil {
				return err
			}

			chat.DeleteData("card-create")
			chat.SetAction(tg.ChatContextDeck)
			return nil
		}
		return fmt.Errorf("Invalid CardCreateController callback query data: %s", udp.CallbackQuery)
	case tg.BotUpdateText:

		switch chat.GetData("card-create").(int) {
		case state.CardCreateFront:
			card := &model.Card{
				DeckID:    deck.ID,
				FrontText: udp.Message.Text,
			}

			chat.SetData("card", card)
			_, err := tg.Send(view.AddCardView(chat.ID, deck.Name, state.CardCreateBack))
			if err != nil {
				return err
			}

		case state.CardCreateBack:
			card := chat.GetData("card").(*model.Card)
			card.BackText = udp.Message.Text

			err := cardRepository.InsertCard(card)
			if err != nil {
				return err
			}

			_, err = tg.Send(view.OpenDeckView(chat.ID, deck))
			if err != nil {
				return err
			}

			chat.DeleteData("card-create")
			chat.SetAction(tg.ChatContextDeck)
			return nil
		}
	}

	return nil
}

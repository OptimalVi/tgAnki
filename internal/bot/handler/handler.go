package handler

import (
	"fmt"
	"optimal_vi/tg_anki/internal/bot/controller"
	"optimal_vi/tg_anki/pkg/tg"
)

var handlerData []*tg.BotUpdateHandleData

func Handler(chat *tg.ChatContext, udp *tg.BotUpdate) error {
	Action := chat.GetAction()
	UdpType, err := udp.GetType()
	if err != nil {
		return err
	}

	AddHandlerData(&tg.BotUpdateHandleData{
		Action:     tg.ChatContextNew,
		UpdateType: tg.BotUpdateAny,
		Controller: controller.HomeController,
	})
	// Home display
	AddHandlerData(&tg.BotUpdateHandleData{
		Action:     tg.ChatContextHome,
		UpdateType: tg.BotUpdateCallbackQuery,
		Controller: controller.StartController,
	})
	AddHandlerData(&tg.BotUpdateHandleData{
		Action:     tg.ChatContextDeckCreate,
		UpdateType: tg.BotUpdateText,
		Controller: controller.DeckCreateController,
	})

	// Deck display
	AddHandlerData(&tg.BotUpdateHandleData{
		Action:     tg.ChatContextDeck,
		UpdateType: tg.BotUpdateCallbackQuery,
		Controller: controller.DeckController,
	})
	AddHandlerData(&tg.BotUpdateHandleData{
		Action:     tg.ChatContextDeckDrop,
		UpdateType: tg.BotUpdateCallbackQuery,
		Controller: controller.DeckDropController,
	})
	AddHandlerData(&tg.BotUpdateHandleData{
		Action:     tg.ChatContextDeckEdit,
		UpdateType: tg.BotUpdateText,
		Controller: controller.DeckEditController,
	})
	AddHandlerData(&tg.BotUpdateHandleData{
		Action:     tg.ChatContextDeckEdit,
		UpdateType: tg.BotUpdateCallbackQuery,
		Controller: controller.DeckEditController,
	})

	AddHandlerData(&tg.BotUpdateHandleData{
		Action:     tg.ChatContextCardCreate,
		UpdateType: tg.BotUpdateText,
		Controller: controller.CardCreateController,
	})
	AddHandlerData(&tg.BotUpdateHandleData{
		Action:     tg.ChatContextCardCreate,
		UpdateType: tg.BotUpdateCallbackQuery,
		Controller: controller.CardCreateController,
	})

	matched, err := DefineController(Action, UdpType)
	if err != nil {
		controller.HomeController(chat, udp)
		return err
	}

	return matched.CallControllerFunc(chat, udp)
}

func AddHandlerData(data *tg.BotUpdateHandleData) {
	handlerData = append(handlerData, data)
}

func DefineController(chatAction int, updateType int) (*tg.BotUpdateHandleData, error) {
	for _, data := range handlerData {
		if data.Action == chatAction {
			switch {
			case data.UpdateType == tg.BotUpdateAny:
				return data, nil
			case data.UpdateType == updateType:
				return data, nil
			}
		}
	}

	return nil, fmt.Errorf("No matching handler Action %d, UpdateType %d", chatAction, updateType)
}

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

// func StartHandler(chat *tg.ChatContext, udp tg.BotUpdate) error {
// 	fmt.Printf("\nStart handler triggered for chat ID %d\n", chat.ID)

// 	tp, _ := udp.GetType()
// 	fmt.Printf("Type: %d\n", tp)
// 	isCallback, err := udp.TypeIs(tg.BotUpdateCallbackQuery)
// 	if err != nil {
// 		fmt.Print("\nError b-h-t-17:", err)
// 	}
// 	if !isCallback {
// 		fmt.Print("Error b-h-t-23: Invalid update type")
// 		return tg.InvalidMessageType{}
// 	}

// 	if udp.CallbackQuery.Data == "create-deck" {
// 		controller.CreateDeckController(chat)
// 	} else {
// 		deckId, err := strconv.Atoi(udp.CallbackQuery.Data)
// 		if err != nil {
// 			fmt.Print("\nError b-h-t-29:", err)
// 			return err
// 		}
// 		controller.OpenDeckController(chat, deckId)
// 	}

// 	return nil
// }

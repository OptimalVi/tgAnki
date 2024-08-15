package tg

type UnknownMessageType struct{}

func (u UnknownMessageType) Error() string { return "Unknown message type" }

type InvalidMessageType struct{}

func (i InvalidMessageType) Error() string { return "Invalid message type" }

type InvalidMessage struct{}

func (i InvalidMessage) Error() string { return "Invalid message" }

type InvalidMessageData struct{}

func (i InvalidMessageData) Error() string { return "Invalid message data" }

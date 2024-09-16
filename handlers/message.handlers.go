package handlers

import (
	"nugu.dev/rd-vigor/repositories"
	"nugu.dev/rd-vigor/services"
)

type MessageService interface {
	CreateMessage(sender_id string, content string, chatroom_id string) *services.ServiceLayerErr

	GetLastChatroomMessage(chatroom_id string) (repositories.Message, *services.ServiceLayerErr)
	GetChatroomMessages(chatroom_id string) ([]repositories.Message, *services.ServiceLayerErr)
}

type MessageHandler struct {
	Service MessageService
}

func NewMessageHandler(ms MessageService) *MessageHandler {
	return &MessageHandler{
		Service: ms,
	}
}

package services

import (
	"net/http"

	"nugu.dev/rd-vigor/repositories"
)

type MessageRepository interface {
	CreateMessage(sender_id string, content string, chatroom_id string) *repositories.RepositoryLayerErr

	GetLastChatroomMessage(chatroom_id string) (repositories.Message, *repositories.RepositoryLayerErr)
	GetChatroomMessages(chatroom_id string) ([]repositories.Message, *repositories.RepositoryLayerErr)
}

type MessageService struct {
	Repository MessageRepository
}

func NewMessageService(mr MessageRepository) *MessageService {
	return &MessageService{
		Repository: mr,
	}
}

func (ms *MessageService) GetLastChatroomMessage(chatroom_id string) (repositories.Message, *ServiceLayerErr) {

	var msg repositories.Message
	msg, err := ms.Repository.GetLastChatroomMessage(chatroom_id)

	if err != nil {
		return msg, &ServiceLayerErr{err.Error, "Query Err", http.StatusInternalServerError}
	}

	return msg, nil
}

func (ms *MessageService) GetChatroomMessages(chatroom_id string) ([]repositories.Message, *ServiceLayerErr) {

	messages, err := ms.Repository.GetChatroomMessages(chatroom_id)

	if err != nil {
		return messages, &ServiceLayerErr{err.Error, "Query Err", http.StatusInternalServerError}
	}

	return messages, nil
}

func (ms *MessageService) CreateMessage(sender_id string, content string, chatroom_id string) *ServiceLayerErr {

	err := ms.Repository.CreateMessage(sender_id, content, chatroom_id)

	if err != nil {
		return &ServiceLayerErr{err.Error, "Query Err", http.StatusInternalServerError}
	}

	return nil
}

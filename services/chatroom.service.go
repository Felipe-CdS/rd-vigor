package services

import (
	"nugu.dev/rd-vigor/repositories"
)

type ChatroomRepository interface {
	CreateChatroom(sender_id string, recipient_id string) (string, *repositories.RepositoryLayerErr)

	GetChatroomsByUserID(id string) ([]repositories.Chatroom, *repositories.RepositoryLayerErr)
	GetChatroomRecipient(user1_id string, chatroom_id string) (string, *repositories.RepositoryLayerErr)
}

type ChatroomLastData struct {
	Chatroom    *repositories.Chatroom
	LastMessage *repositories.Message
	LastSender  *repositories.User
	Recipient   *repositories.User
}

type ChatroomService struct {
	Repository ChatroomRepository
}

func NewChatroomService(cr ChatroomRepository) *ChatroomService {
	return &ChatroomService{
		Repository: cr,
	}
}

func (cs *ChatroomService) CreateChatroom(sender_id string, recipient_id string) (string, *ServiceLayerErr) {

	chatroom_id, err := cs.Repository.CreateChatroom(sender_id, recipient_id)

	if err != nil {
		return "", &ServiceLayerErr{Error: err.Error, Message: "Service Layer Error"}
	}
	return chatroom_id, nil
}

func (cs *ChatroomService) GetChatroomsByUserID(id string) ([]repositories.Chatroom, *ServiceLayerErr) {

	chatrooms, err := cs.Repository.GetChatroomsByUserID(id)

	if err != nil {
		return chatrooms, &ServiceLayerErr{Error: err.Error, Message: "Service Layer Error"}
	}
	return chatrooms, nil
}

func (cs *ChatroomService) GetChatroomRecipient(user1_id string, chatroom_id string) (string, *ServiceLayerErr) {

	recipientId, err := cs.Repository.GetChatroomRecipient(user1_id, chatroom_id)

	if err != nil {
		return recipientId, &ServiceLayerErr{Error: err.Error, Message: "Service Layer Error"}
	}
	return recipientId, nil
}

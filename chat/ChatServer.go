package chat

import (
	"github.com/labstack/echo/v4"
	"nugu.dev/rd-vigor/repositories"
)

type WsServer struct {
	chatrooms map[*Hub]bool
	mr        *repositories.MessageRepository
}

func NewWsSever(mr *repositories.MessageRepository) *WsServer {
	return &WsServer{
		mr:        mr,
		chatrooms: make(map[*Hub]bool),
	}
}

func (w *WsServer) NewHub(c echo.Context, chatroomId string) *Hub {

	for k := range w.chatrooms {
		if k.Id == chatroomId {
			return k
		}
	}

	messages, err := w.mr.GetChatroomMessages("2")

	if err != nil {

	}

	hub := NewHub(chatroomId, messages)

	w.chatrooms[hub] = true

	go hub.Run(c)

	return hub
}

package handlers

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"nugu.dev/rd-vigor/repositories"
	"nugu.dev/rd-vigor/services"
	"nugu.dev/rd-vigor/views/inbox_views"
)

type ChatroomService interface {
	CreateChatroom(sender_id string, recipient_id string) (string, *services.ServiceLayerErr)

	GetChatroomsByUserID(id string) ([]repositories.Chatroom, *services.ServiceLayerErr)
	GetChatroomRecipient(user1_id string, chatroom_id string) (string, *services.ServiceLayerErr)
}

func NewChatroomHandler(cs ChatroomService, us UserService, ms MessageService) *ChatroomHandler {
	return &ChatroomHandler{
		ChatroomServices: cs,
		UserServices:     us,
		MessageServices:  ms,
	}
}

type ChatroomHandler struct {
	ChatroomServices ChatroomService
	UserServices     UserService
	MessageServices  MessageService
}

func (ch *ChatroomHandler) GetInboxBase(c echo.Context) error {

	loggedUser := c.Get("user").(repositories.User)
	chatroomId := c.QueryParam("chatroom")

	return ch.View(c,
		inbox_views.Base(
			"Mensagens",
			loggedUser,
			chatroomId,
		))
}

func (ch *ChatroomHandler) NewChatroomModal(c echo.Context) error {
	return ch.View(c, inbox_views.NewChatModal())
}

func (ch *ChatroomHandler) SelectRecipient(c echo.Context) error {

	if c.QueryParam("username") == "" {
		return ch.View(c, inbox_views.RecipientToBeSelectedDiv())
	}

	usr, queryErr := ch.UserServices.GetUserByUsername(c.QueryParam("username"))

	if queryErr != nil {
	}

	return ch.View(c, inbox_views.RecipientSelectedDiv(usr))
}

func (ch *ChatroomHandler) GetUserChatroomsList(c echo.Context) error {

	data := []services.ChatroomLastData{}

	loggedUser := c.Get("user").(repositories.User)

	chatrooms, queryErr := ch.ChatroomServices.GetChatroomsByUserID(loggedUser.ID)

	if queryErr != nil {
	}

	for _, c := range chatrooms {
		m, queryErr := ch.MessageServices.GetLastChatroomMessage(c.ChatroomId)

		if queryErr != nil {
		}

		u, queryErr := ch.UserServices.GetUserByID(m.UserId)

		if queryErr != nil {
		}

		rId, queryErr := ch.ChatroomServices.GetChatroomRecipient(loggedUser.ID, c.ChatroomId)

		if queryErr != nil {
		}

		r, queryErr := ch.UserServices.GetUserByID(rId)

		if queryErr != nil {
		}

		d := &services.ChatroomLastData{
			Chatroom:    &c,
			LastSender:  &u,
			LastMessage: &m,
			Recipient:   &r,
		}

		data = append(data, *d)
	}

	return ch.View(c, inbox_views.RecipientsList(data))
}

func (ch *ChatroomHandler) CreateChatroom(c echo.Context) error {

	loggedUser := c.Get("user").(repositories.User)
	content := c.FormValue("content")

	recipient, queryErr := ch.UserServices.GetUserByUsername(c.FormValue("recipient"))

	if queryErr != nil {
	}

	chatroomId, err := ch.ChatroomServices.CreateChatroom(loggedUser.ID, recipient.ID)

	if err != nil {
		fmt.Println(err)
	}

	if err = ch.MessageServices.CreateMessage(loggedUser.ID, content, chatroomId); err != nil {

	}

	return c.Redirect(http.StatusSeeOther, fmt.Sprintf("/inbox?chatroom=%s", chatroomId))
}

func (ch *ChatroomHandler) GetChatroomDetails(c echo.Context) error {

	loggedUser := c.Get("user").(repositories.User)
	chatroomId := c.Param("chatroom_id")

	rId, queryErr := ch.ChatroomServices.GetChatroomRecipient(loggedUser.ID, chatroomId)

	if queryErr != nil {
	}

	r, queryErr := ch.UserServices.GetUserByID(rId)

	if queryErr != nil {
	}

	return ch.View(c, inbox_views.RecipientDetails(r))
}

func (ch *ChatroomHandler) GetChat(c echo.Context) error {

	loggedUser := c.Get("user").(repositories.User)
	chatroomId := c.Param("chatroom_id")

	rId, queryErr := ch.ChatroomServices.GetChatroomRecipient(loggedUser.ID, chatroomId)

	if queryErr != nil {
	}

	r, queryErr := ch.UserServices.GetUserByID(rId)

	if queryErr != nil {
	}

	m, queryErr := ch.MessageServices.GetChatroomMessages(chatroomId)

	if queryErr != nil {
		fmt.Println(queryErr)
	}

	return ch.View(c, inbox_views.ChatDiv(r, m, chatroomId))
}

func (ch *ChatroomHandler) View(c echo.Context, cmp templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

	return cmp.Render(c.Request().Context(), c.Response().Writer)
}

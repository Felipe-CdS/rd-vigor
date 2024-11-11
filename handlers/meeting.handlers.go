package handlers

import (
	"net/http"
	"time"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"nugu.dev/rd-vigor/services"
	"nugu.dev/rd-vigor/views/inbox_views"
)

type MeetingService interface {
	CreateMeeting(user1ID string, user2ID string, timestamp time.Time) *services.ServiceLayerErr
}

type MeetingHandler struct {
	Service     MeetingService
	UserService UserService
}

func NewMeetingHandler(ms MeetingService, us UserService) *MeetingHandler {
	return &MeetingHandler{
		Service:     ms,
		UserService: us,
	}
}

func (mh *MeetingHandler) GetMeetingCreationModal(c echo.Context) error {

	recipient, err := mh.UserService.GetUser(c.QueryParam("recipient"))

	if err != nil {
		c.Response().WriteHeader(http.StatusBadRequest)
		return nil
	}

	return mh.View(c, inbox_views.MeetingRecipientForm(recipient))
}

func (mh *MeetingHandler) View(c echo.Context, cmp templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

	return cmp.Render(c.Request().Context(), c.Response().Writer)
}

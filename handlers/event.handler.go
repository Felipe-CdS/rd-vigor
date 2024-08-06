package handlers

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"nugu.dev/rd-vigor/repositories"
	"nugu.dev/rd-vigor/views/events_views"
)

type EventService interface {
}

func NewEventHandler(es EventService) *EventHandler {
	return &EventHandler{
		EventServices: es,
	}
}

type EventHandler struct {
	EventServices EventService
}

func (eh *EventHandler) GetEventSearchPage(c echo.Context) error {
	usr := repositories.User{}

	return eh.View(c, events_views.EventSearch(usr))
}

func (eh *EventHandler) View(c echo.Context, cmp templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

	return cmp.Render(c.Request().Context(), c.Response().Writer)
}

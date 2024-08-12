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

	loggedUser := c.Get("user").(repositories.User)
	return eh.View(c, events_views.EventSearch(loggedUser))
}

func (eh *EventHandler) GetEventDetails(c echo.Context) error {

	loggedUser := c.Get("user").(repositories.User)
	return eh.View(c, events_views.EventDetails(loggedUser))
}

func (eh *EventHandler) View(c echo.Context, cmp templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

	return cmp.Render(c.Request().Context(), c.Response().Writer)
}

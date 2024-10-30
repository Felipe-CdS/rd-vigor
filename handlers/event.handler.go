package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"nugu.dev/rd-vigor/repositories"
	"nugu.dev/rd-vigor/services"
	"nugu.dev/rd-vigor/views/courses_views"
	"nugu.dev/rd-vigor/views/events_views"
)

type EventService interface {
	CreateEvent(e repositories.Event) *services.ServiceLayerErr

	GetEvents(past bool) ([]repositories.Event, *services.ServiceLayerErr)
	GetEventByID(id string) (repositories.Event, *services.ServiceLayerErr)
	GetNextEvent() (repositories.Event, *services.ServiceLayerErr)
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
	past := c.QueryParam("p") == "true"

	events, err := eh.EventServices.GetEvents(past)

	if err != nil {
		fmt.Println(err)
		return c.Redirect(http.StatusMovedPermanently, "/signin")
	}

	return eh.View(c, events_views.EventSearch(loggedUser, events, past))
}

func (eh *EventHandler) CreateNewEvent(c echo.Context) error {

	loggedUser := c.Get("user").(repositories.User)

	if loggedUser.Role != "admin" {
		return c.Redirect(http.StatusMovedPermanently, "/signin")
	}

	if c.FormValue("title") == "" {
		c.Response().WriteHeader(http.StatusBadRequest)
		return eh.View(c, events_views.ErrorAlert("Nome Inválido."))
	}

	if c.FormValue("description") == "" {
		c.Response().WriteHeader(http.StatusBadRequest)
		return eh.View(c, events_views.ErrorAlert("Descrição inválida."))
	}

	if c.FormValue("date") == "" {
		c.Response().WriteHeader(http.StatusBadRequest)
		return eh.View(c, events_views.ErrorAlert("Data inválida."))
	}

	i, err := strconv.ParseInt(c.FormValue("date"), 10, 64)

	if err != nil {
		c.Response().WriteHeader(http.StatusBadRequest)
		return eh.View(c, events_views.ErrorAlert("Data inválida."))
	}

	e := repositories.Event{
		Title:       c.FormValue("title"),
		Description: c.FormValue("description"),
		Date:        time.Unix(i, 0),
	}

	if err := eh.EventServices.CreateEvent(e); err != nil {
		if err.Code == http.StatusInternalServerError {
			c.Response().WriteHeader(http.StatusInternalServerError)
			return eh.View(c, events_views.ErrorAlert("Um erro inesperado ocorreu no servidor. Por favor, tente novamente mais tarde."))
		}

		c.Response().WriteHeader(err.Code)
		return eh.View(c, events_views.ErrorAlert(err.Message))
	}
	return c.Redirect(http.StatusSeeOther, "/events")
}

func (eh *EventHandler) GetEventDetails(c echo.Context) error {

	loggedUser := c.Get("user").(repositories.User)

	e, queryErr := eh.EventServices.GetEventByID(c.Param("id"))

	if queryErr != nil {
		return c.Redirect(http.StatusMovedPermanently, "/events")
	}

	return eh.View(c, events_views.EventDetails(loggedUser, e))
}

func (eh *EventHandler) GetCoursesSearchPage(c echo.Context) error {
	loggedUser := c.Get("user").(repositories.User)
	return eh.View(c, courses_views.CoursesSearch(loggedUser))
}

func (eh *EventHandler) GetEventDashboard(c echo.Context) error {
	loggedUser := c.Get("user").(repositories.User)
	return eh.View(c, events_views.EventsDashboard("dashboard", loggedUser))
}

func (eh *EventHandler) View(c echo.Context, cmp templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

	return cmp.Render(c.Request().Context(), c.Response().Writer)
}

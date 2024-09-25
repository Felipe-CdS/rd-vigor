package handlers

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"nugu.dev/rd-vigor/repositories"
	"nugu.dev/rd-vigor/services"
	"nugu.dev/rd-vigor/views/settings_views"
)

type PortifolioService interface {
	CreatePortifolio(u repositories.User, t string, d string) *services.ServiceLayerErr
	EditPortifolio(u repositories.User, portifolioId string, t string, d string) *services.ServiceLayerErr
	DeletePortifolio(u repositories.User, portifolioId string) *services.ServiceLayerErr
	GetUserPortifolios(u repositories.User) ([]repositories.Portifolio, *services.ServiceLayerErr)
}

type PortifolioHandler struct {
	Service PortifolioService
}

func NewPortifolioHandler(ps PortifolioService) *PortifolioHandler {
	return &PortifolioHandler{
		Service: ps,
	}
}

func (ph *PortifolioHandler) CreatePortifolio(c echo.Context) error {

	loggedUser := c.Get("user").(repositories.User)
	t := c.FormValue("title")
	d := c.FormValue("description")

	if err := ph.Service.CreatePortifolio(loggedUser, t, d); err != nil {
		c.Response().WriteHeader(err.Code)
		return nil
	}

	return ph.View(c, settings_views.PortifolioSection())
}

func (ph *PortifolioHandler) EditPortifolio(c echo.Context) error {
	loggedUser := c.Get("user").(repositories.User)
	id := c.FormValue("id")
	t := c.FormValue("title")
	d := c.FormValue("description")

	if err := ph.Service.EditPortifolio(loggedUser, id, t, d); err != nil {
		c.Response().WriteHeader(err.Code)
		return nil
	}

	return ph.View(c, settings_views.PortifolioSection())
}

func (ph *PortifolioHandler) DeletePortifolio(c echo.Context) error {
	loggedUser := c.Get("user").(repositories.User)
	portifolioId := c.QueryParam("id")

	if err := ph.Service.DeletePortifolio(loggedUser, portifolioId); err != nil {
		c.Response().WriteHeader(err.Code)
		return nil
	}

	return ph.View(c, settings_views.PortifolioSection())
}

func (ph *PortifolioHandler) GetUserPortifolios(c echo.Context) error {

	loggedUser := c.Get("user").(repositories.User)

	p, err := ph.Service.GetUserPortifolios(loggedUser)

	if err != nil {
		c.Response().WriteHeader(err.Code)
		return nil
	}

	return ph.View(c, settings_views.PortifolioList(p))
}

func (ph *PortifolioHandler) View(c echo.Context, cmp templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

	return cmp.Render(c.Request().Context(), c.Response().Writer)
}

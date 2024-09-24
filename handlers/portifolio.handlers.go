package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"nugu.dev/rd-vigor/repositories"
	"nugu.dev/rd-vigor/services"
)

type PortifolioService interface {
	CreatePortifolio(u repositories.User, t string, d string) *services.ServiceLayerErr
	GetUserPortifolios(u repositories.User) *services.ServiceLayerErr
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
		// return th.View(c, tags_views.ErrorAlert("Nome Inválido."))
		return nil
	}

	return c.Redirect(http.StatusSeeOther, "/admin/dashboard/tags")
}

func (ph *PortifolioHandler) EditPortifolio(c echo.Context) error {
	return nil
}

func (ph *PortifolioHandler) GetUserPortifolios(c echo.Context) error {

	loggedUser := c.Get("user").(repositories.User)

	if err := ph.Service.GetUserPortifolios(loggedUser); err != nil {
		c.Response().WriteHeader(err.Code)
		// return th.View(c, tags_views.ErrorAlert("Nome Inválido."))
		return nil
	}

	return c.Redirect(http.StatusSeeOther, "/admin/dashboard/tags")
}

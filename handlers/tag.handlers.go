package handlers

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"nugu.dev/rd-vigor/repositories"
	"nugu.dev/rd-vigor/services"
	"nugu.dev/rd-vigor/views/tags_views"
)

type TagService interface {
	CreateTag(n string) *services.ServiceLayerErr
	GetAllTags() ([]repositories.Tag, *services.ServiceLayerErr)
}

type TagHandler struct {
	Service TagService
}

func NewTagHandler(ts TagService) *TagHandler {
	return &TagHandler{
		Service: ts,
	}
}

func (th *TagHandler) GetTagDashboard(c echo.Context) error {

	loggedUser := c.Get("user").(repositories.User)

	if loggedUser.Role != "admin" {
		return c.Redirect(http.StatusMovedPermanently, "/signin")
	}

	tags, err := th.Service.GetAllTags()

	if err != nil {

	}

	c.Response().Header().Set("HX-Retarget", "body")
	c.Response().Header().Set("HX-Push-Url", "/admin/dashboard/users")
	return th.View(c, tags_views.TagsDashboard("Dashboard", loggedUser, tags))
}

func (th *TagHandler) CreateNewTag(c echo.Context) error {

	loggedUser := c.Get("user").(repositories.User)

	if loggedUser.Role != "admin" {
		return c.Redirect(http.StatusMovedPermanently, "/signin")
	}

	if c.FormValue("tag-name") == "" {
		c.Response().WriteHeader(http.StatusBadRequest)
		return th.View(c, tags_views.ErrorAlert("Nome Inválido."))
	}

	if err := th.Service.CreateTag(c.FormValue("tag-name")); err != nil {
		if err.Code == http.StatusInternalServerError {
			c.Response().WriteHeader(http.StatusInternalServerError)
			return th.View(c, tags_views.ErrorAlert("Um erro inesperado ocorreu no servidor. Por favor, tente novamente mais tarde."))
		}

		c.Response().WriteHeader(err.Code)
		return th.View(c, tags_views.ErrorAlert(err.Message))
	}
	return c.Redirect(http.StatusSeeOther, "/admin/dashboard/tags")
}

func (th *TagHandler) View(c echo.Context, cmp templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

	return cmp.Render(c.Request().Context(), c.Response().Writer)
}

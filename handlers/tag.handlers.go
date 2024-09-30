package handlers

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"nugu.dev/rd-vigor/repositories"
	"nugu.dev/rd-vigor/services"
	admin_views "nugu.dev/rd-vigor/views/admin_views/dashboard"
	"nugu.dev/rd-vigor/views/layout"
	"nugu.dev/rd-vigor/views/settings_views"
	"nugu.dev/rd-vigor/views/tags_views"
)

type TagService interface {
	CreateTag(n string) *services.ServiceLayerErr
	SearchTagByName(n string) ([]repositories.Tag, *services.ServiceLayerErr)
	GetAllTags() ([]repositories.Tag, *services.ServiceLayerErr)
	GetUserTags(u repositories.User) ([]repositories.Tag, *services.ServiceLayerErr)
}

type TagHandler struct {
	Service     TagService
	UserService UserService
}

func NewTagHandler(ts TagService, us UserService) *TagHandler {
	return &TagHandler{
		Service:     ts,
		UserService: us,
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

func (th *TagHandler) SearchTagByName(c echo.Context) error {

	loggedUser := c.Get("user").(repositories.User)

	// if triggerInput == settings-tag-search:
	// request from settings
	// if triggerInput == search:
	// request from admin dashboard

	triggerInput := c.Request().Header.Get("HX-Trigger-Name")
	tagName := c.FormValue(triggerInput)

	list, err := th.Service.SearchTagByName(tagName)

	if err != nil {
		if err.Code == http.StatusInternalServerError {
			c.Response().WriteHeader(http.StatusInternalServerError)
			return th.View(c, tags_views.ErrorAlert("Um erro inesperado ocorreu no servidor. Por favor, tente novamente mais tarde."))
		}

		c.Response().WriteHeader(err.Code)
		return th.View(c, tags_views.ErrorAlert(err.Message))

	}

	if triggerInput == "settings-tag-search" {
		return th.View(c, settings_views.UserTagsList(list, loggedUser))
	}

	user, err := th.UserService.GetUserByUsername(c.FormValue("user"))

	if err != nil {
		if err.Code == http.StatusInternalServerError {
			c.Response().WriteHeader(http.StatusInternalServerError)
			return th.View(c, tags_views.ErrorAlert("Um erro inesperado ocorreu no servidor. Por favor, tente novamente mais tarde."))
		}

		c.Response().WriteHeader(err.Code)
		return th.View(c, tags_views.ErrorAlert(err.Message))
	}

	return th.View(c, admin_views.TagsListResponse(list, user))
}

func (th *TagHandler) SearchTagNavbar(c echo.Context) error {

	searchInput := c.FormValue("search")
	loggedUser := c.Get("user").(repositories.User)

	if loggedUser.Role != "admin" {
		return c.Redirect(http.StatusMovedPermanently, "/signin")
	}

	list, err := th.Service.SearchTagByName(searchInput)

	if err != nil {
		if err.Code == http.StatusInternalServerError {
			c.Response().WriteHeader(http.StatusInternalServerError)
			return th.View(c, tags_views.ErrorAlert("Um erro inesperado ocorreu no servidor. Por favor, tente novamente mais tarde."))
		}

		c.Response().WriteHeader(err.Code)
		return th.View(c, tags_views.ErrorAlert(err.Message))
	}

	return th.View(c, layout.SearchModalResponsePartial(list, searchInput))
}

func (th *TagHandler) GetUserTags(c echo.Context) error {

	loggedUser := c.Get("user").(repositories.User)

	list, err := th.Service.GetUserTags(loggedUser)

	if err != nil {
		c.Response().WriteHeader(err.Code)
		return nil
	}

	return th.View(c, settings_views.UserTagsList(list, loggedUser))
}

func (th *TagHandler) View(c echo.Context, cmp templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

	return cmp.Render(c.Request().Context(), c.Response().Writer)
}

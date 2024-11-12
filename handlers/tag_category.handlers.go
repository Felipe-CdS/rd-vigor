package handlers

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"nugu.dev/rd-vigor/repositories"
	"nugu.dev/rd-vigor/services"
	"nugu.dev/rd-vigor/views/settings_views"
	"nugu.dev/rd-vigor/views/tags_views"
)

type TagCategoryService interface {
	CreateTagCategory(n string) *services.ServiceLayerErr

	GetAllTagCategories() ([]repositories.TagCategory, *services.ServiceLayerErr)
	GetAllTagsByCategory(categoryId string) ([]repositories.Tag, *services.ServiceLayerErr)
}

type TagCategoryHandler struct {
	Service TagCategoryService
}

func NewTagCategoryHandler(ts TagCategoryService) *TagCategoryHandler {
	return &TagCategoryHandler{
		Service: ts,
	}
}

func (th *TagCategoryHandler) GetTagCategoryDashboard(c echo.Context) error {

	loggedUser := c.Get("user").(repositories.User)

	if loggedUser.Role != "admin" {
		return c.Redirect(http.StatusMovedPermanently, "/signin")
	}

	tc, err := th.Service.GetAllTagCategories()

	if err != nil {

	}

	c.Response().Header().Set("HX-Retarget", "body")
	c.Response().Header().Set("HX-Push-Url", "/admin/dashboard/users")
	return th.View(c, tags_views.TagCategoriesDashboard("Dashboard", loggedUser, tc))
}

func (th *TagCategoryHandler) CreateNewTagCategory(c echo.Context) error {

	loggedUser := c.Get("user").(repositories.User)

	if loggedUser.Role != "admin" {
		return c.Redirect(http.StatusMovedPermanently, "/signin")
	}

	if c.FormValue("category-name") == "" {
		c.Response().WriteHeader(http.StatusBadRequest)
		return th.View(c, tags_views.ErrorAlert("Nome Inválido."))
	}

	if err := th.Service.CreateTagCategory(c.FormValue("category-name")); err != nil {
		if err.Code == http.StatusInternalServerError {
			c.Response().WriteHeader(http.StatusInternalServerError)
			return th.View(c, tags_views.ErrorAlert("Um erro inesperado ocorreu no servidor. Por favor, tente novamente mais tarde."))
		}

		c.Response().WriteHeader(err.Code)
		return th.View(c, tags_views.ErrorAlert(err.Message))
	}
	return c.Redirect(http.StatusSeeOther, "/admin/dashboard/tag-categories")
}

func (th *TagCategoryHandler) GetAllTagsByCategory(c echo.Context) error {

	loggedUser := c.Get("user").(repositories.User)

	if c.FormValue("category-name") == "" {
		c.Response().WriteHeader(http.StatusBadRequest)
		return th.View(c, tags_views.ErrorAlert("Nome Inválido."))
	}

	tags, err := th.Service.GetAllTagsByCategory(c.FormValue("category-name"))

	if err != nil {
		if err.Code == http.StatusInternalServerError {
			c.Response().WriteHeader(http.StatusInternalServerError)
			return th.View(c, tags_views.ErrorAlert("Um erro inesperado ocorreu no servidor. Por favor, tente novamente mais tarde."))
		}

		c.Response().WriteHeader(err.Code)
		return th.View(c, tags_views.ErrorAlert(err.Message))
	}
	return th.View(c, settings_views.TagsByCategoryList(loggedUser, tags))
}

func (th *TagCategoryHandler) View(c echo.Context, cmp templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

	return cmp.Render(c.Request().Context(), c.Response().Writer)
}

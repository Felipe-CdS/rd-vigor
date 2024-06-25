package handlers

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"nugu.dev/rd-vigor/services"
	"nugu.dev/rd-vigor/views/login"
)

type UserService interface {
	CreateUser(u services.User) error
}

func NewUserHandler(us UserService) *UserHandler {
	return &UserHandler{
		UserServices: us,
	}
}

type UserHandler struct {
	UserServices UserService
}

func (uh *UserHandler) CreateNewUser(c echo.Context) error {

	if c.Request().Method == "POST" {
		// if c.FormValue("password") != c.FormValue("password") {
		// 	return echo.NewHTTPError(http.StatusNotFound)
		// }

		user := services.User{
			FirstName:      c.FormValue("first_name"),
			LastName:       c.FormValue("last_name"),
			Email:          c.FormValue("email"),
			OccupationArea: c.FormValue("occupation_area"),
			Password:       c.FormValue("password"),
			CreatedAt:      1532009163,
		}

		fmt.Printf("%+v\n", user)
		err := uh.UserServices.CreateUser(user)

		if err != nil {
			return err
		}
		return c.Redirect(http.StatusSeeOther, "/signup-done")
	}

	return uh.View(c, login.LoginForm())
}

func (uh *UserHandler) View(c echo.Context, cmp templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

	return cmp.Render(c.Request().Context(), c.Response().Writer)
}

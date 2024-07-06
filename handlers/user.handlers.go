package handlers

import (
	"net/http"
	"net/mail"
	"time"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"nugu.dev/rd-vigor/repositories"
	"nugu.dev/rd-vigor/services"
	"nugu.dev/rd-vigor/views/auth_views"
)

type UserService interface {
	CreateUser(u repositories.User) *services.ServiceLayerErr
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

		if c.FormValue("first_name") == "" {
			c.Response().WriteHeader(http.StatusBadRequest)
			return uh.View(c, auth_views.SignupFormErrorAlert("Por favor, insira um primeiro nome válido."))
		}

		if c.FormValue("last_name") == "" {
			c.Response().WriteHeader(http.StatusBadRequest)
			return uh.View(c, auth_views.SignupFormErrorAlert("Por favor, insira um sobrenome válido."))
		}

		if c.FormValue("email") == "" {
			c.Response().WriteHeader(http.StatusBadRequest)
			return uh.View(c, auth_views.SignupFormErrorAlert("Por favor, insira um email válido."))
		}

		if _, err := mail.ParseAddress(c.FormValue("email")); err != nil {
			c.Response().WriteHeader(http.StatusBadRequest)
			return uh.View(c, auth_views.SignupFormErrorAlert("E-mail inválido."))
		}

		if c.FormValue("password") == "" {
			c.Response().WriteHeader(http.StatusBadRequest)
			return uh.View(c, auth_views.SignupFormErrorAlert("Por favor, insira uma senha válida."))
		}

		if c.FormValue("password") != c.FormValue("repeat-password") {
			c.Response().WriteHeader(http.StatusBadRequest)
			return uh.View(c, auth_views.SignupFormErrorAlert("As senhas inseridas não são idênticas."))
		}

		if c.FormValue("occupation_area") == "" {
			c.Response().WriteHeader(http.StatusBadRequest)
			return uh.View(c, auth_views.SignupFormErrorAlert("Por favor, insira uma área de atuação válida."))
		}

		user := repositories.User{
			FirstName:      c.FormValue("first_name"),
			LastName:       c.FormValue("last_name"),
			Email:          c.FormValue("email"),
			OccupationArea: c.FormValue("occupation_area"),
			Password:       c.FormValue("password"),
			Telephone:      c.FormValue("telephone"),
			ReferFriend:    c.FormValue("refer_friend"),
			CreatedAt:      time.Now(),
		}

		if err := uh.UserServices.CreateUser(user); err != nil {
			if err.Code == http.StatusInternalServerError {
				c.Response().WriteHeader(http.StatusInternalServerError)
				return uh.View(c, auth_views.SignupFormErrorAlert("Um erro inesperado ocorreu no servidor. Por favor, tente novamente mais tarde."))
			}

			c.Response().WriteHeader(err.Code)
			return uh.View(c, auth_views.SignupFormErrorAlert(err.Message))
		}
		return c.Redirect(http.StatusSeeOther, "/signup-done")
	}

	return uh.View(c, auth_views.SigninForm())
}

func (uh *UserHandler) View(c echo.Context, cmp templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

	return cmp.Render(c.Request().Context(), c.Response().Writer)
}

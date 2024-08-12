package handlers

import (
	"fmt"
	"net/http"
	"net/mail"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"nugu.dev/rd-vigor/auth"
	"nugu.dev/rd-vigor/repositories"
	"nugu.dev/rd-vigor/services"
	admin_views "nugu.dev/rd-vigor/views/admin_views/dashboard"
	"nugu.dev/rd-vigor/views/auth_views"
	"nugu.dev/rd-vigor/views/user_views"
)

type UserService interface {
	CreateUser(data map[string]string) *services.ServiceLayerErr
	AuthUser(login string, password string) (repositories.User, *services.ServiceLayerErr)
	GetAllUsers() ([]repositories.User, *services.ServiceLayerErr)
	GetUserByID(id string) (repositories.User, *services.ServiceLayerErr)
	GetUserByUsername(username string) (repositories.User, *services.ServiceLayerErr)
	SetNewTagUser(username string, tag_name string) *services.ServiceLayerErr
	GetUserTags(user repositories.User) ([]repositories.Tag, *services.ServiceLayerErr)
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

	if c.Request().Method == "GET" {
		cmp := auth_views.Base("Sign up", auth_views.SignupForm())
		return cmp.Render(c.Request().Context(), c.Response().Writer)
	}

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
			return uh.View(c, auth_views.SignupFormErrorAlert("Por favor, insira um email válido."))
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

		if c.FormValue("telephone") == "" {
			c.Response().WriteHeader(http.StatusBadRequest)
			return uh.View(c, auth_views.SignupFormErrorAlert("Por favor, insira um telefone válido."))
		}

		userData := map[string]string{
			"FirstName":      c.FormValue("first_name"),
			"LastName":       c.FormValue("last_name"),
			"Email":          c.FormValue("email"),
			"OccupationArea": c.FormValue("occupation_area"),
			"Telephone":      c.FormValue("telephone"),
			"Password":       c.FormValue("password"),
			"ReferFriend":    c.FormValue("refer_friend"),
		}

		if err := uh.UserServices.CreateUser(userData); err != nil {
			if err.Code == http.StatusInternalServerError {
				c.Response().WriteHeader(http.StatusInternalServerError)
				return uh.View(c, auth_views.SignupFormErrorAlert("Um erro inesperado ocorreu no servidor. Por favor, tente novamente mais tarde."))
			}

			c.Response().WriteHeader(err.Code)
			return uh.View(c, auth_views.SignupFormErrorAlert(err.Message))
		}
		return c.Redirect(http.StatusSeeOther, "/signup-done")
	}
	return c.Redirect(http.StatusMethodNotAllowed, "/signup")
}

func (uh *UserHandler) SigninUser(c echo.Context) error {

	if c.Request().Method == "GET" {
		cookieToken, err := c.Cookie("access-token")

		if err == nil && cookieToken.Value != "" {
			claims, err := auth.DecodeToken(cookieToken.Value)

			if err != nil {
				auth.ResetAuthCookies(c)
				return c.Redirect(http.StatusMovedPermanently, "/signin")
			}

			if claims.Role == "admin" {
				return c.Redirect(http.StatusMovedPermanently, "/admin/dashboard/users")
			}

			return c.Redirect(http.StatusMovedPermanently, "/user/home")
		}

		cmp := auth_views.Base("Login or sign up", auth_views.SigninForm())
		return cmp.Render(c.Request().Context(), c.Response().Writer)
	}

	if c.Request().Method == "POST" {
		usr, err := uh.UserServices.AuthUser(c.FormValue("login"), c.FormValue("password"))

		if err != nil {
			c.Response().WriteHeader(err.Code)
			return uh.View(c, auth_views.SigninFormErrorAlert(err.Message))
		}

		if usr.Role == "admin" {
			if err := auth.GenerateTokensAndSetCookies(usr, c); err != nil {
				c.Response().WriteHeader(http.StatusInternalServerError)
				return uh.View(c, auth_views.SigninFormErrorAlert("Um erro inesperado ocorreu no servidor. Por favor, tente novamente mais tarde."))
			}
			return c.Redirect(http.StatusMovedPermanently, "/admin/dashboard/users")
		} else {
			cmp := auth_views.SigninFormDone()
			return cmp.Render(c.Request().Context(), c.Response().Writer)
		}

	}

	return c.Redirect(http.StatusMethodNotAllowed, "/signup")
}

func (uh *UserHandler) GetAdminUserList(c echo.Context) error {

	loggedUser := c.Get("user").(repositories.User)

	if loggedUser.Role != "admin" {
		return c.Redirect(http.StatusMovedPermanently, "/signin")
	}

	users, queryErr := uh.UserServices.GetAllUsers()

	if queryErr != nil {
	}

	c.Response().Header().Set("HX-Retarget", "body")
	c.Response().Header().Set("HX-Push-Url", "/admin/dashboard/users")
	return uh.View(c, admin_views.Base("Dashboard", users, loggedUser))
}

func (uh *UserHandler) GetUserDetails(c echo.Context) error {

	if c.Request().Header.Get("HX-Request") != "true" {
		c.Response().Header().Set("HX-redirect", "/admin/dashboard/users")
		return c.NoContent(http.StatusMovedPermanently)
	}

	usr, queryErr := uh.UserServices.GetUserByID(c.QueryParam("user"))

	if queryErr != nil {
		c.Response().Header().Set("HX-redirect", "/admin/dashboard/users")
		return c.NoContent(http.StatusMovedPermanently)
	}

	return uh.View(c, admin_views.UserInfoDiv(usr))
}

func (uh *UserHandler) GetUserProfile(c echo.Context) error {

	loggedUser := c.Get("user").(repositories.User)

	usr, queryErr := uh.UserServices.GetUserByUsername(c.Param("username"))

	fmt.Printf("%+v\n", usr)

	if queryErr != nil {
		c.Response().Header().Set("HX-redirect", "/admin/dashboard/users")
		return c.NoContent(http.StatusMovedPermanently)
	}

	tags, queryErr := uh.UserServices.GetUserTags(usr)

	if queryErr != nil {
		c.Response().Header().Set("HX-redirect", "/admin/dashboard/users")
		return c.NoContent(http.StatusMovedPermanently)
	}

	fmt.Printf("%+v\n", tags)

	return uh.View(c,
		user_views.UserProfile(
			fmt.Sprintf("%s %s", usr.FirstName, usr.LastName),
			loggedUser,
			usr,
			tags,
		))
}

func (uh *UserHandler) SetUserTag(c echo.Context) error {
	username := c.FormValue("user")
	tagName := c.FormValue("tag")

	uh.UserServices.SetNewTagUser(username, tagName)
	return nil
}

func (uh *UserHandler) View(c echo.Context, cmp templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

	return cmp.Render(c.Request().Context(), c.Response().Writer)
}

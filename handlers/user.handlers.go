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
	"nugu.dev/rd-vigor/views/home_views"
	"nugu.dev/rd-vigor/views/inbox_views"
	"nugu.dev/rd-vigor/views/search_views"
	"nugu.dev/rd-vigor/views/settings_views"
	"nugu.dev/rd-vigor/views/user_views"
)

type UserService interface {
	CreateUser(data map[string]string) *services.ServiceLayerErr
	UpdateUser(user repositories.User, newUserData repositories.User) *services.ServiceLayerErr
	AuthUser(login string, password string) (repositories.User, *services.ServiceLayerErr)

	GetAllUsers() ([]repositories.User, *services.ServiceLayerErr)
	GetUsersByAny(any string) ([]repositories.User, *services.ServiceLayerErr)
	GetUsersByTagID(any string) ([]repositories.User, *services.ServiceLayerErr)

	GetUserByID(id string) (repositories.User, *services.ServiceLayerErr)
	GetUserByStripeID(id string) (repositories.User, *services.ServiceLayerErr)
	GetUserByUsername(username string) (repositories.User, *services.ServiceLayerErr)

	GetUserTags(user repositories.User) ([]repositories.Tag, *services.ServiceLayerErr)
	GetUserNotTags(user repositories.User) ([]repositories.Tag, *services.ServiceLayerErr)
	SetNewTagUser(username string, tagId string) *services.ServiceLayerErr
	DeleteUserTag(user repositories.User, tagId string) *services.ServiceLayerErr

	SearchTagByNameAvaiableToUser(user repositories.User, query string) ([]repositories.Tag, *services.ServiceLayerErr)
}

func NewUserHandler(us UserService, es EventService, ps PortifolioService, ts TagService) *UserHandler {
	return &UserHandler{
		UserServices:       us,
		EventServices:      es,
		PortifolioServices: ps,
		TagServices:        ts,
	}
}

type UserHandler struct {
	UserServices       UserService
	EventServices      EventService
	PortifolioServices PortifolioService
	TagServices        TagService
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
			}

			return c.Redirect(http.StatusMovedPermanently, "/home")
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

		if err := auth.GenerateTokensAndSetCookies(usr, c); err != nil {
			c.Response().WriteHeader(http.StatusInternalServerError)
			return uh.View(c, auth_views.SigninFormErrorAlert("Um erro inesperado ocorreu no servidor. Por favor, tente novamente mais tarde."))
		}
		return c.Redirect(http.StatusMovedPermanently, "/home")
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

	if queryErr != nil {
		c.Response().Header().Set("HX-redirect", "/admin/dashboard/users")
		return c.NoContent(http.StatusMovedPermanently)
	}

	tags, queryErr := uh.UserServices.GetUserTags(usr)

	if queryErr != nil {
		c.Response().Header().Set("HX-redirect", "/admin/dashboard/users")
		return c.NoContent(http.StatusMovedPermanently)
	}

	portifolios, queryErr := uh.PortifolioServices.GetUserPortifolios(usr)

	if queryErr != nil {
		c.Response().Header().Set("HX-redirect", "/admin/dashboard/users")
		return c.NoContent(http.StatusMovedPermanently)
	}

	return uh.View(c,
		user_views.UserProfile(
			fmt.Sprintf("%s %s", usr.FirstName, usr.LastName),
			loggedUser,
			usr,
			tags,
			portifolios,
		))
}

func (uh *UserHandler) GetHome(c echo.Context) error {

	loggedUser := c.Get("user").(repositories.User)

	upcoming, err := uh.EventServices.GetNextEvent()

	if err != nil {
		upcoming = repositories.Event{}
	}

	return uh.View(c,
		home_views.Base(
			"Home",
			loggedUser,
			upcoming,
		))
}

func (uh *UserHandler) GetCalendar(c echo.Context) error {

	loggedUser := c.Get("user").(repositories.User)

	return uh.View(c,
		user_views.Calendar(
			"Agenda",
			loggedUser,
		))
}

func (uh *UserHandler) SetUserTag(c echo.Context) error {

	username := c.FormValue("user")
	tagId := c.FormValue("tag")

	triggerInput := c.Request().Header.Get("HX-Trigger-Name")

	if triggerInput == "settings-set-tag" {
		loggedUser := c.Get("user").(repositories.User)
		username = loggedUser.Username
		tagId = c.QueryParam("tag")
	}

	if err := uh.UserServices.SetNewTagUser(username, tagId); err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		return uh.View(c, auth_views.SignupFormErrorAlert("Um erro inesperado ocorreu no servidor. Por favor, tente novamente mais tarde."))
	}

	return c.Redirect(http.StatusSeeOther, "/settings/profile/tags")
}

func (uh *UserHandler) SearchTagByNameAvaiableToUser(c echo.Context) error {

	loggedUser := c.Get("user").(repositories.User)

	// if triggerInput == settings-tag-search:
	// request from settings
	// if triggerInput == search:
	// request from admin dashboard

	triggerInput := c.Request().Header.Get("HX-Trigger-Name")
	tagName := c.FormValue(triggerInput)

	list, err := uh.UserServices.SearchTagByNameAvaiableToUser(loggedUser, tagName)

	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		return nil
	}

	return uh.View(c, settings_views.SearchTagsList(tagName, list, loggedUser))
}

func (uh *UserHandler) DeleteUserTag(c echo.Context) error {

	loggedUser := c.Get("user").(repositories.User)
	tagId := c.QueryParam("tag")

	if err := uh.UserServices.DeleteUserTag(loggedUser, tagId); err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		return uh.View(c, auth_views.SignupFormErrorAlert("Um erro inesperado ocorreu no servidor. Por favor, tente novamente mais tarde."))
	}

	return c.Redirect(http.StatusSeeOther, "/settings/profile/tags")
}

func (uh *UserHandler) SearchUserByAny(c echo.Context) error {
	query := c.FormValue("query")

	if query == "" {
		return uh.View(c, inbox_views.SearchUserFormOptionsUndefined())
	}

	found, err := uh.UserServices.GetUsersByAny(query)

	if err != nil {
	}

	return uh.View(c, inbox_views.SearchUserFormOptions(found))
}

func (uh *UserHandler) SearchUsersByTag(c echo.Context) error {

	loggedUser := c.Get("user").(repositories.User)

	t := c.QueryParam("t")

	tag, err := uh.TagServices.GetTagByID(t)

	if err != nil {
	}

	foundUsers, err := uh.UserServices.GetUsersByTagID(t)

	if err != nil {
	}

	userPlusTags := []search_views.UserAndTagsStruct{}

	for _, f := range foundUsers {

		foundTags, err := uh.UserServices.GetUserTags(f)

		if err != nil {
			foundTags = []repositories.Tag{}
		}

		x := search_views.UserAndTagsStruct{
			User: f,
			Tags: foundTags,
		}

		userPlusTags = append(userPlusTags, x)
	}

	return uh.View(c, search_views.Base(loggedUser, userPlusTags, tag.Name))
}

func (uh *UserHandler) UpdateUserAccountInfo(c echo.Context) error {

	loggedUser := c.Get("user").(repositories.User)

	formData := repositories.User{
		Username:  c.FormValue("username"),
		FirstName: c.FormValue("first_name"),
		LastName:  c.FormValue("last_name"),
		Email:     c.FormValue("email"),
	}

	if formData.FirstName == "" {
		c.Response().WriteHeader(http.StatusBadRequest)
		return uh.View(c, settings_views.UpdateErrorAlert("Por favor, insira um Nome válido."))
	}

	if formData.LastName == "" {
		c.Response().WriteHeader(http.StatusBadRequest)
		return uh.View(c, settings_views.UpdateErrorAlert("Por favor, insira um Sobrenome válido."))
	}

	if formData.Email == "" {
		c.Response().WriteHeader(http.StatusBadRequest)
		return uh.View(c, settings_views.UpdateErrorAlert("Por favor, insira um email válido."))
	}

	if _, err := mail.ParseAddress(formData.Email); err != nil {
		c.Response().WriteHeader(http.StatusBadRequest)
		return uh.View(c, settings_views.UpdateErrorAlert("Por favor, insira um email válido."))
	}

	if formData.Username == "" {
		c.Response().WriteHeader(http.StatusBadRequest)
		return uh.View(c, settings_views.UpdateErrorAlert("Por favor, insira um nome de usuário válido."))
	}

	if err := uh.UserServices.UpdateUser(loggedUser, formData); err != nil {
		if err.Code == http.StatusInternalServerError {
			c.Response().WriteHeader(http.StatusInternalServerError)
			return uh.View(c, auth_views.SignupFormErrorAlert("Um erro inesperado ocorreu no servidor. Por favor, tente novamente mais tarde."))
		}

		c.Response().WriteHeader(err.Code)
		return uh.View(c, settings_views.UpdateErrorAlert(err.Message))
	}

	updatedUser, queryErr := uh.UserServices.GetUserByID(loggedUser.ID)

	if queryErr != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		return uh.View(c, auth_views.SignupFormErrorAlert("Um erro inesperado ocorreu no servidor. Por favor, tente novamente mais tarde."))
	}

	if err := auth.GenerateTokensAndSetCookies(updatedUser, c); err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		return uh.View(c, auth_views.SignupFormErrorAlert("Um erro inesperado ocorreu no servidor. Por favor, tente novamente mais tarde."))
	}
	return uh.View(c, settings_views.ContactInfoSettings(updatedUser))
}

func (uh *UserHandler) UpdateUserLocationInfo(c echo.Context) error {

	loggedUser := c.Get("user").(repositories.User)

	formData := repositories.User{
		Username:  loggedUser.Username,
		Email:     loggedUser.Email,
		Address:   c.FormValue("address"),
		Address2:  c.FormValue("address2"),
		City:      c.FormValue("city"),
		State:     c.FormValue("state"),
		Zipcode:   c.FormValue("zipcode"),
		Telephone: c.FormValue("telephone"),
	}

	if err := uh.UserServices.UpdateUser(loggedUser, formData); err != nil {
		if err.Code == http.StatusInternalServerError {
			c.Response().WriteHeader(http.StatusInternalServerError)
			return uh.View(c, auth_views.SignupFormErrorAlert("Um erro inesperado ocorreu no servidor. Por favor, tente novamente mais tarde."))
		}

		c.Response().WriteHeader(err.Code)
		return uh.View(c, settings_views.UpdateErrorAlert(err.Message))
	}

	updatedUser, queryErr := uh.UserServices.GetUserByID(loggedUser.ID)

	if queryErr != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		return uh.View(c, auth_views.SignupFormErrorAlert("Um erro inesperado ocorreu no servidor. Por favor, tente novamente mais tarde."))
	}

	return uh.View(c, settings_views.ContactInfoSettings(updatedUser))
}

func (uh *UserHandler) View(c echo.Context, cmp templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

	return cmp.Render(c.Request().Context(), c.Response().Writer)
}

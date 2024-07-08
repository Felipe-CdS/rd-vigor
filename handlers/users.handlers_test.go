package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"nugu.dev/rd-vigor/db"
	"nugu.dev/rd-vigor/repositories"
	"nugu.dev/rd-vigor/services"
)

func TestCreateNewUser(t *testing.T) {

	f := make(url.Values)
	f.Set("first_name", "Drake")
	f.Set("last_name", "Josh")
	f.Set("email", "drakejosh@gmail.com")
	f.Set("password", "123456")
	f.Set("repeat-password", "123456")
	f.Set("occupation_area", "Web Design")
	f.Set("telephone", "21 9751222231")

	t.Setenv("APP_ENV", "TESTING")

	e := echo.New()

	store := db.NewStore()
	ur := repositories.NewUserRepository(repositories.User{}, store)
	us := services.NewUserService(ur)
	uh := NewUserHandler(us)

	t.Run("Sucess", func(t *testing.T) {

		req := httptest.NewRequest(http.MethodPost, "/signup", strings.NewReader(f.Encode()))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)

		if err := uh.CreateNewUser(c); err != nil {
			t.Error(err)
		}

		if rec.Result().StatusCode != http.StatusSeeOther {
			t.Errorf("Expected 200, found %d\n", c.Response().Status)
			t.Errorf("%s\n", rec.Body.String())
		}
	})

	t.Run("Wrong First Name", func(t *testing.T) {

		f.Set("first_name", "")

		req := httptest.NewRequest(http.MethodPost, "/signup", strings.NewReader(f.Encode()))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)

		if err := uh.CreateNewUser(c); err != nil {
			t.Error(err)
		}

		if rec.Result().StatusCode != http.StatusBadRequest {
			t.Errorf("Expected 400, found %d\n", c.Response().Status)
		}
	})

	t.Run("Password Mismatch", func(t *testing.T) {

		f.Set("password", "123")
		f.Set("repeat-password", "1234")

		req := httptest.NewRequest(http.MethodPost, "/signup", strings.NewReader(f.Encode()))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)

		if err := uh.CreateNewUser(c); err != nil {
			t.Error(err)
		}

		if rec.Result().StatusCode != http.StatusBadRequest {
			t.Errorf("Expected 400, found %d\n", c.Response().Status)
		}
	})
}

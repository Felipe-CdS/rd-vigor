package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"nugu.dev/rd-vigor/auth"
	"nugu.dev/rd-vigor/views/auth_views"
)

func SetupRoutes(e *echo.Echo, uh *UserHandler) {

	e.GET("/", func(c echo.Context) error {
		return c.Redirect(http.StatusPermanentRedirect, "/signin")
	})

	e.GET("/admin", func(c echo.Context) error {
		return c.Redirect(http.StatusPermanentRedirect, "/admin/dashboard")
	})

	e.GET("/signup-done", signupFormDone)

	e.GET("/signup", uh.CreateNewUser)
	e.POST("/signup", uh.CreateNewUser)

	e.GET("/signin", uh.SigninUser)
	e.POST("/signin", uh.SigninUser)

	e.GET("/admin/dashboard", uh.GetAdminUserList)
	e.GET("/admin/dashboard/details", uh.GetUserDetails)
	e.GET("/logout", func(c echo.Context) error {
		auth.ResetAuthCookies(c)
		c.Response().Header().Set("HX-redirect", "/signin")
		return c.NoContent(http.StatusPermanentRedirect)
	})
}

func signupFormDone(c echo.Context) error {
	cmp := auth_views.SignupFormDone()
	return cmp.Render(c.Request().Context(), c.Response().Writer)
}

func signinFormDone(c echo.Context) error {
	cmp := auth_views.SigninFormDone()
	return cmp.Render(c.Request().Context(), c.Response().Writer)
}

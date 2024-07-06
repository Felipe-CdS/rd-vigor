package handlers

import (
	"github.com/labstack/echo/v4"
	"nugu.dev/rd-vigor/views/auth_views"
	"nugu.dev/rd-vigor/views/landing_page"
)

func SetupRoutes(e *echo.Echo, uh *UserHandler) {
	e.GET("/", authViewBase)

	e.GET("/signup", signupForm)
	e.GET("/signup-done", signupFormDone)

	e.POST("/signup", uh.CreateNewUser)

	e.GET("/not-ready", signinFormDone)

	// e.GET("/home", userHome)
}

func userHome(c echo.Context) error {
	cmp := landing_page.LandingIndex("Login or sign up")
	return cmp.Render(c.Request().Context(), c.Response().Writer)
}

func authViewBase(c echo.Context) error {
	cmp := auth_views.Base("Login or sign up")
	return cmp.Render(c.Request().Context(), c.Response().Writer)
}

func signupForm(c echo.Context) error {
	cmp := auth_views.SignupForm()
	return cmp.Render(c.Request().Context(), c.Response().Writer)
}

func signupFormDone(c echo.Context) error {
	cmp := auth_views.SignupFormDone()
	return cmp.Render(c.Request().Context(), c.Response().Writer)
}

func signinFormDone(c echo.Context) error {
	cmp := auth_views.SigninFormDone()
	return cmp.Render(c.Request().Context(), c.Response().Writer)
}

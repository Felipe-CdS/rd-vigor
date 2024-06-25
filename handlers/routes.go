package handlers

import (
	"github.com/labstack/echo/v4"
	"nugu.dev/rd-vigor/views/landing_page"
	"nugu.dev/rd-vigor/views/login"
)

func SetupRoutes(e *echo.Echo, uh *UserHandler) {
	e.GET("/", loginLanding)

	e.GET("/signup", signupFormGET)
	e.POST("/signup", uh.CreateNewUser)
	e.GET("/signup-done", signupFormDone)

	e.GET("/login", loginForm)
	e.GET("/not-ready", notReady)

	// e.GET("/home", userHome)
}

func userHome(c echo.Context) error {
	cmp := landing_page.LandingIndex("Login or sign up")
	return cmp.Render(c.Request().Context(), c.Response().Writer)
}

func loginLanding(c echo.Context) error {
	cmp := login.LoginIndex("Login or sign up")
	return cmp.Render(c.Request().Context(), c.Response().Writer)
}

func signupFormGET(c echo.Context) error {
	cmp := login.SignupForm()
	return cmp.Render(c.Request().Context(), c.Response().Writer)
}

func signupFormDone(c echo.Context) error {
	cmp := login.SignupFormDone()
	return cmp.Render(c.Request().Context(), c.Response().Writer)
}

func loginForm(c echo.Context) error {
	cmp := login.LoginForm()
	return cmp.Render(c.Request().Context(), c.Response().Writer)
}

func notReady(c echo.Context) error {
	cmp := login.NotReady()
	return cmp.Render(c.Request().Context(), c.Response().Writer)
}

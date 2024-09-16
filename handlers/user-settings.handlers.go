package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"nugu.dev/rd-vigor/repositories"
	"nugu.dev/rd-vigor/views/user_views"
)

func (uh *UserHandler) GetSettingsPage(c echo.Context) error {

	loggedUser := c.Get("user").(repositories.User)

	return uh.View(c,
		user_views.Settings(
			"Ajustes",
			loggedUser,
		))
}

func (uh *UserHandler) GetContactInfoSettings(c echo.Context) error {

	loggedUser := c.Get("user").(repositories.User)

	if c.Request().Header.Get("HX-Request") != "true" {
		c.Response().Header().Set("HX-redirect", "/settings")
		return c.NoContent(http.StatusMovedPermanently)
	}

	return uh.View(c, user_views.ContactInfoSettings(loggedUser))
}

func (uh *UserHandler) GetBillingSettings(c echo.Context) error {

	if c.Request().Header.Get("HX-Request") != "true" {
		c.Response().Header().Set("HX-redirect", "/settings")
		return c.NoContent(http.StatusMovedPermanently)
	}

	return uh.View(c, user_views.BillingSettings())
}

func (uh *UserHandler) GetProfileSettings(c echo.Context) error {

	if c.Request().Header.Get("HX-Request") != "true" {
		c.Response().Header().Set("HX-redirect", "/settings")
		return c.NoContent(http.StatusMovedPermanently)
	}

	return uh.View(c, user_views.ProfileSettings())
}

func (uh *UserHandler) GetSecuritySettings(c echo.Context) error {

	if c.Request().Header.Get("HX-Request") != "true" {
		c.Response().Header().Set("HX-redirect", "/settings")
		return c.NoContent(http.StatusMovedPermanently)
	}

	return uh.View(c, user_views.SecuritySettings())
}

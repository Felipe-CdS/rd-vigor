package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"nugu.dev/rd-vigor/repositories"
	"nugu.dev/rd-vigor/views/settings_views"
)

func (uh *UserHandler) GetSettingsPage(c echo.Context) error {

	loggedUser := c.Get("user").(repositories.User)

	return uh.View(c,
		settings_views.Base(
			"Ajustes",
			loggedUser,
			nil,
		))
}

func (uh *UserHandler) GetContactInfoSettings(c echo.Context) error {

	loggedUser := c.Get("user").(repositories.User)

	if c.Request().Header.Get("HX-Request") != "true" {
		c.Response().Header().Set("HX-redirect", "/settings")
		return c.NoContent(http.StatusMovedPermanently)
	}

	return uh.View(c, settings_views.ContactInfoSettings(loggedUser))
}

func (uh *UserHandler) GetBillingSettings(c echo.Context) error {

	if c.Request().Header.Get("HX-Request") != "true" {

		loggedUser := c.Get("user").(repositories.User)
		txResult := c.QueryParam("payment_intent_client_secret")

		return uh.View(c,
			settings_views.Base(
				"Ajustes",
				loggedUser,
				settings_views.BillingSettingsComplete(txResult),
			))
	}

	return uh.View(c, settings_views.BillingSettings())
}

func (uh *UserHandler) GetProfileSettings(c echo.Context) error {

	if c.Request().Header.Get("HX-Request") != "true" {
		c.Response().Header().Set("HX-redirect", "/settings")
		return c.NoContent(http.StatusMovedPermanently)
	}

	return uh.View(c, settings_views.ProfileSettings())
}

func (uh *UserHandler) GetSecuritySettings(c echo.Context) error {

	loggedUser := c.Get("user").(repositories.User)

	if c.Request().Header.Get("HX-Request") != "true" {
		c.Response().Header().Set("HX-redirect", "/settings")
		return c.NoContent(http.StatusMovedPermanently)
	}

	return uh.View(c, settings_views.SecuritySettings(loggedUser))
}

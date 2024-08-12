package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"nugu.dev/rd-vigor/auth"
	"nugu.dev/rd-vigor/views/auth_views"
)

func SetupRoutes(e *echo.Echo, uh *UserHandler, eh *EventHandler, th *TagHandler) {

	e.GET("/", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, "/signin")
	})

	e.GET("/admin/dashboard", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, "/admin/dashboard/users")
	})

	e.GET("/admin", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, "/admin/dashboard/users")
	})

	e.GET("/signup-done", signupFormDone)

	e.GET("/signup", uh.CreateNewUser)
	e.POST("/signup", uh.CreateNewUser)

	e.GET("/signin", uh.SigninUser)
	e.POST("/signin", uh.SigninUser)

	e.GET("/logout", func(c echo.Context) error {
		auth.ResetAuthCookies(c)
		c.Response().Header().Set("HX-redirect", "/signin")
		return c.NoContent(http.StatusMovedPermanently)
	})

	e.GET("/user/:username", authMiddleware(uh, uh.GetUserProfile))
	e.POST("/user/tags", authMiddleware(uh, uh.SetUserTag))
	e.GET("/admin/dashboard/users", authMiddleware(uh, uh.GetAdminUserList))
	e.GET("/admin/dashboard/users/details", authMiddleware(uh, uh.GetUserDetails))
	e.GET("/admin/dashboard/tags", authMiddleware(uh, th.GetTagDashboard))
	e.POST("/admin/dashboard/tags", authMiddleware(uh, th.CreateNewTag))
	/* EVENTS ROUTES*/

	e.GET("/events", authMiddleware(uh, eh.GetEventSearchPage))
	e.GET("/event/:event_id", authMiddleware(uh, eh.GetEventDetails))

	e.POST("/tags/search", authMiddleware(uh, th.SearchTagByName))
}

func signupFormDone(c echo.Context) error {
	cmp := auth_views.SignupFormDone()
	return cmp.Render(c.Request().Context(), c.Response().Writer)
}

func signinFormDone(c echo.Context) error {
	cmp := auth_views.SigninFormDone()
	return cmp.Render(c.Request().Context(), c.Response().Writer)
}

func authMiddleware(uh *UserHandler, next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {

		cookieToken, err := c.Cookie("access-token")

		if err != nil {
			return c.Redirect(http.StatusMovedPermanently, "/signin")
		}

		claims, err := auth.DecodeToken(cookieToken.Value)

		if err != nil {
			return c.Redirect(http.StatusMovedPermanently, "/signin")
		}

		loggedUser, queryErr := uh.UserServices.GetUserByUsername(claims.Username)

		if queryErr != nil {
			fmt.Printf("!!!!!!%+v\n", queryErr)
		}

		c.Set("user", loggedUser)
		return next(c)
	}
}

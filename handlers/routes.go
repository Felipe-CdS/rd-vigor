package handlers

import (
	"github.com/labstack/echo/v4"
    "nugu.dev/rd-vigor/views/landing_page"
)

func SetupRoutes(e *echo.Echo){
    e.GET("/", Test)
}

func Test(c echo.Context) error{
    cmp := landing_page.LandingIndex("Home")
    return cmp.Render(c.Request().Context(), c.Response().Writer)
}

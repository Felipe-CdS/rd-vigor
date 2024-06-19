package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"nugu.dev/rd-vigor/handlers"
)

func main() {
    e := echo.New()

    e.Static("/static", "assets")

    e.Use(middleware.Logger())

    handlers.SetupRoutes(e)
    e.Logger.Fatal(e.Start(":42069"))    
}

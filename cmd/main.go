package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"nugu.dev/rd-vigor/db"
	"nugu.dev/rd-vigor/handlers"
	"nugu.dev/rd-vigor/services"
)

func main() {
	e := echo.New()

	e.Static("/static", "assets")
	e.Use(middleware.Logger())

	store := db.NewStore("APP_DATA.db")

	us := services.NewUserService(services.User{}, store)
	uh := handlers.NewUserHandler(us)

	handlers.SetupRoutes(e, uh)
	e.Logger.Fatal(e.Start(":42069"))
}

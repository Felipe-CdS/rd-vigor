package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"nugu.dev/rd-vigor/db"
	"nugu.dev/rd-vigor/handlers"
	"nugu.dev/rd-vigor/services"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")

	if port == "" {
		port = ":42069"
	}

	e := echo.New()

	e.Static("/static", "assets")
	e.Use(middleware.Logger())

	store := db.NewStore()

	us := services.NewUserService(services.User{}, store)
	uh := handlers.NewUserHandler(us)

	handlers.SetupRoutes(e, uh)
	e.Logger.Fatal(e.Start(port))
}

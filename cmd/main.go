package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"nugu.dev/rd-vigor/db"
	"nugu.dev/rd-vigor/handlers"
	"nugu.dev/rd-vigor/repositories"
	"nugu.dev/rd-vigor/services"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")

	fmt.Printf("%s\n", port)

	if port == "" {
		port = "42069"
	}

	e := echo.New()

	e.Static("/static", "assets")
	e.Use(middleware.Logger())

	store := db.NewStore()

	ur := repositories.NewUserRepository(repositories.User{}, store)
	us := services.NewUserService(ur)
	uh := handlers.NewUserHandler(us)

	handlers.SetupRoutes(e, uh)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}

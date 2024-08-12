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

	if port == "" {
		port = "42069"
	}

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Static("/static", "assets")

	e.Use(middleware.Logger())

	if os.Getenv("APP_ENV") != "PROD" {
	}

	store := db.NewStore()

	ur := repositories.NewUserRepository(repositories.User{}, store)
	us := services.NewUserService(ur)
	uh := handlers.NewUserHandler(us)

	er := repositories.NewEventRepository(repositories.Event{}, store)
	es := services.NewEventService(er)
	eh := handlers.NewEventHandler(es)

	tr := repositories.NewTagRepository(repositories.Tag{}, store)
	ts := services.NewTagService(tr)
	th := handlers.NewTagHandler(ts)

	handlers.SetupRoutes(e, uh, eh, th)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}

package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"nugu.dev/rd-vigor/chat"
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

	// e.Use(middleware.Logger())

	if os.Getenv("APP_ENV") != "PROD" {
	}

	store := db.NewStore()

	tr := repositories.NewTagRepository(repositories.Tag{}, store)
	ur := repositories.NewUserRepository(repositories.User{}, store)
	er := repositories.NewEventRepository(repositories.Event{}, store)
	cr := repositories.NewChatroomRepository(repositories.Chatroom{}, store)
	mr := repositories.NewMessageRepository(repositories.Message{}, store)

	ts := services.NewTagService(tr)
	us := services.NewUserService(ur, tr)
	es := services.NewEventService(er)
	cs := services.NewChatroomService(cr)
	ms := services.NewMessageService(mr)

	th := handlers.NewTagHandler(ts, us)
	uh := handlers.NewUserHandler(us)
	eh := handlers.NewEventHandler(es)
	ch := handlers.NewChatroomHandler(cs, us, ms)

	wsServer := chat.NewWsSever(mr)

	handlers.SetupRoutes(e, wsServer, uh, eh, th, ch)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}

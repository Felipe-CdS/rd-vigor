package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/stripe/stripe-go/v79"
	"github.com/stripe/stripe-go/v79/paymentintent"
	"nugu.dev/rd-vigor/auth"
	"nugu.dev/rd-vigor/chat"
	"nugu.dev/rd-vigor/repositories"
	"nugu.dev/rd-vigor/views/auth_views"
)

func SetupRoutes(e *echo.Echo,
	wsServer *chat.WsServer,
	uh *UserHandler,
	eh *EventHandler,
	th *TagHandler,
	ch *ChatroomHandler,
) {

	e.GET("/", authMiddleware(uh, uh.GetHome))

	e.GET("/ws-chatroom/:chatroom_id", authMiddleware(
		uh,
		func(c echo.Context) error {

			loggedUser := c.Get("user").(repositories.User)

			hub := wsServer.NewHub(c, c.Param("chatroom_id"))

			chat.ServeWs(hub, c.Response().Writer, c.Request(), loggedUser.ID)
			return nil
		}),
	)

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

	e.POST("/users/search", authMiddleware(uh, uh.SearchUserByAny))
	e.POST("/tags/search", authMiddleware(uh, th.SearchTagByName))
	e.POST("/navbar/search", authMiddleware(uh, th.SearchTagNavbar))

	/* INBOX ROUTES*/
	e.GET("/calendar", authMiddleware(uh, uh.GetCalendar))

	/* INBOX ROUTES*/

	e.GET("/inbox", authMiddleware(uh, ch.GetInboxBase))
	e.POST("/chatroom/new", authMiddleware(uh, ch.CreateChatroom))
	e.GET("/chatroom/new/select-recipient", authMiddleware(uh, ch.SelectRecipient))
	e.GET("/chatroom/list", authMiddleware(uh, ch.GetUserChatroomsList))

	e.GET("/chatroom/:chatroom_id", authMiddleware(uh, ch.GetChat))
	e.GET("/chatroom/:chatroom_id/chat", authMiddleware(uh, ch.GetChat))
	e.GET("/chatroom/:chatroom_id/details", authMiddleware(uh, ch.GetChatroomDetails))

	/* SETTINGS ROUTES*/
	e.GET("/settings", authMiddleware(uh, uh.GetSettingsPage))
	e.GET("/settings/billing", authMiddleware(uh, uh.GetBillingSettings))
	e.GET("/settings/contact-info", authMiddleware(uh, uh.GetContactInfoSettings))
	e.GET("/settings/profile", authMiddleware(uh, uh.GetProfileSettings))
	e.GET("/settings/security", authMiddleware(uh, uh.GetSecuritySettings))

	e.POST("/settings/contact-info/account", authMiddleware(uh, uh.UpdateUserAccountInfo))
	e.POST("/settings/contact-info/location", authMiddleware(uh, uh.UpdateUserLocationInfo))

	/* STRIPE ROUTES*/
	e.POST("/create-payment-intent", authMiddleware(uh, CreatePaymentIntent))
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
			fmt.Printf("Middleware: %+v\n", queryErr)
		}

		c.Set("user", loggedUser)
		return next(c)
	}
}

func CreatePaymentIntent(c echo.Context) error {
	var i int64 = 10000

	loggedUser := c.Get("user").(repositories.User)
	stripe.Key = "sk_test_51PtsLPP4MxIMgAthuIYfpAKo8RHjeg5Ny4tIILV02TOzUZOE79M18WTbdQcnJavci6HmlgZLNqIABm5CD7J9vaog00sS478WGK"

	params := &stripe.PaymentIntentParams{
		Amount:   &i,
		Currency: stripe.String(string(stripe.CurrencyBRL)),
		Customer: &loggedUser.Username,
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
	}

	pi, err := paymentintent.New(params)

	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		return err
	}

	type v struct {
		ClientSecret   string `json:"clientSecret"`
		DpmCheckerLink string `json:"dpmCheckerLink"`
	}

	x := &v{
		ClientSecret:   pi.ClientSecret,
		DpmCheckerLink: fmt.Sprintf("https://dashboard.stripe.com/settings/payment_methods/review?transaction_id=%s", pi.ID),
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(x); err != nil {
		return err
	}

	return c.JSONBlob(http.StatusOK, buf.Bytes())
}

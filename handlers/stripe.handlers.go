package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/stripe/stripe-go/v79"
	"github.com/stripe/stripe-go/v79/customer"
	"github.com/stripe/stripe-go/v79/paymentintent"
	"github.com/stripe/stripe-go/v79/subscription"
	"github.com/stripe/stripe-go/v79/webhook"
	"nugu.dev/rd-vigor/repositories"
)

func CreatePaymentIntent(c echo.Context) error {
	var i int64 = 10000

	loggedUser := c.Get("user").(repositories.User)

	// Never should enter this state because a stripeID is created
	// when the user log in.
	if loggedUser.StripeID == "" {
		return &stripe.Error{}
	}

	params := &stripe.PaymentIntentParams{
		Amount:   &i,
		Currency: stripe.String(string(stripe.CurrencyBRL)),
		Customer: &loggedUser.StripeID,
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
	}

	pi, err := paymentintent.New(params)

	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		return err
	}

	x := struct {
		ClientSecret string `json:"clientSecret"`
	}{ClientSecret: pi.ClientSecret}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(x); err != nil {
		return err
	}

	return c.JSONBlob(http.StatusOK, buf.Bytes())
}

func HandleCreateSubscription(c echo.Context) error {

	loggedUser := c.Get("user").(repositories.User)

	aaa := "card"

	paymentSettings := &stripe.SubscriptionPaymentSettingsParams{
		PaymentMethodTypes:       []*string{&aaa},
		SaveDefaultPaymentMethod: stripe.String("on_subscription"),
	}

	subscriptionParams := &stripe.SubscriptionParams{
		Customer: stripe.String(loggedUser.StripeID),
		Items: []*stripe.SubscriptionItemsParams{
			{Price: stripe.String(os.Getenv("STRIPE_PRICE_ID"))},
		},
		PaymentSettings: paymentSettings,
		PaymentBehavior: stripe.String("default_incomplete"),
	}

	subscriptionParams.AddExpand("latest_invoice.payment_intent")

	s, err := subscription.New(subscriptionParams)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	x := struct {
		SubscriptionID string `json:"subscriptionId"`
		ClientSecret   string `json:"clientSecret"`
	}{
		SubscriptionID: s.ID,
		ClientSecret:   s.LatestInvoice.PaymentIntent.ClientSecret,
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(x); err != nil {
		return err
	}

	return c.JSONBlob(http.StatusOK, buf.Bytes())
}

func CreateStripeCostumer(loggedUser repositories.User) (string, error) {

	p := &stripe.CustomerParams{
		Name:     stripe.String(fmt.Sprintf("%s %s", loggedUser.FirstName, loggedUser.LastName)),
		Email:    stripe.String(loggedUser.Email),
		Metadata: map[string]string{"rdvigor_ID": loggedUser.ID},
	}

	result, err := customer.New(p)

	if err != nil {
		return "", echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return result.ID, nil
}

func HandleWebhook(c echo.Context, uh *UserHandler) error {

	if c.Request().Method != "POST" {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	b, err := io.ReadAll(c.Request().Body)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	event, err := webhook.ConstructEventWithOptions(b,
		c.Request().Header.Get("Stripe-Signature"),
		os.Getenv("STRIPE_WEBHOOK_SECRET"),
		webhook.ConstructEventOptions{IgnoreAPIVersionMismatch: true},
	)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	customer := fmt.Sprintf("%v", event.Data.Object["customer"])

	u, queryErr := uh.UserServices.GetUserByStripeID(customer)
	updatedUser := u

	if queryErr != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	if event.Type == "payment_intent.succeeded" {
		updatedUser.SubscriptionStatus = true
	}

	if event.Type == "invoice.payment_failed" {
		updatedUser.SubscriptionStatus = false
	}

	if event.Type == "customer.subscription.deleted" {
		updatedUser.SubscriptionStatus = false
	}

	if err := uh.UserServices.UpdateUser(u, updatedUser); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return nil
}

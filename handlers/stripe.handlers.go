package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/stripe/stripe-go/v79"
	"github.com/stripe/stripe-go/v79/customer"
	"github.com/stripe/stripe-go/v79/paymentintent"
	"github.com/stripe/stripe-go/v79/subscription"
	"nugu.dev/rd-vigor/repositories"
)

func CreatePaymentIntent(c echo.Context) error {
	var i int64 = 10000

	loggedUser := c.Get("user").(repositories.User)

	if loggedUser.StripeID == "" {
		params := &stripe.CustomerParams{
			Name:     stripe.String(fmt.Sprintf("%s %s", loggedUser.FirstName, loggedUser.LastName)),
			Email:    stripe.String(loggedUser.Email),
			Metadata: map[string]string{"rdvigor_ID": loggedUser.ID},
		}

		result, err := customer.New(params)

		if err != nil {
			c.Response().WriteHeader(http.StatusInternalServerError)
			return err
		}

		loggedUser.StripeID = result.ID
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

func HandleCreateSubscrition(c echo.Context) error {

	loggedUser := c.Get("user").(repositories.User)

	if loggedUser.StripeID == "" {
		p := &stripe.CustomerParams{
			Name:     stripe.String(fmt.Sprintf("%s %s", loggedUser.FirstName, loggedUser.LastName)),
			Email:    stripe.String(loggedUser.Email),
			Metadata: map[string]string{"rdvigor_ID": loggedUser.ID},
		}

		result, err := customer.New(p)

		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		loggedUser.StripeID = result.ID
	}

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

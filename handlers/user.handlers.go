package handlers

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/a-h/templ"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/labstack/echo/v4"
	"nugu.dev/rd-vigor/services"
	"nugu.dev/rd-vigor/views/login"
)

type UserService interface {
	CreateUser(u services.User) error
}

func NewUserHandler(us UserService) *UserHandler {
	return &UserHandler{
		UserServices: us,
	}
}

type UserHandler struct {
	UserServices UserService
}

func (uh *UserHandler) CreateNewUser(c echo.Context) error {

	if c.Request().Method == "POST" {
		// if c.FormValue("password") != c.FormValue("password") {
		// 	return echo.NewHTTPError(http.StatusNotFound)
		// }

		user := services.User{
			FirstName:      c.FormValue("first_name"),
			LastName:       c.FormValue("last_name"),
			Email:          c.FormValue("email"),
			OccupationArea: c.FormValue("occupation_area"),
			Password:       c.FormValue("password"),
			CreatedAt:      1532009163,
		}

		if err := uh.UserServices.CreateUser(user); err != nil {
			return err
		}

		file, err := c.FormFile("occupation_file")

		if err != nil {
			return err
		}

		src, err := file.Open()

		if err != nil {
			return err
		}

		defer src.Close()

		sess, err := session.NewSession(&aws.Config{
			Region: aws.String("us-east-2"),
		})

		if err != nil {
			return err
		}

		svc := s3.New(sess)

		var buf bytes.Buffer

		if _, err := io.Copy(&buf, src); err != nil {
			return err
		}

		_, err = svc.PutObject(&s3.PutObjectInput{
			Bucket: aws.String("rd-vigor-20050411"),
			Key:    aws.String(fmt.Sprintf("%s_%s_%s", c.FormValue("first_name"), c.FormValue("last_name"), file.Filename)),
			Body:   bytes.NewReader(buf.Bytes()),
		})

		if err != nil {
			return err
		}

		return c.Redirect(http.StatusSeeOther, "/signup-done")
	}

	return uh.View(c, login.LoginForm())
}

func (uh *UserHandler) View(c echo.Context, cmp templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

	return cmp.Render(c.Request().Context(), c.Response().Writer)
}

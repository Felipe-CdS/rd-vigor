package auth

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"

	"nugu.dev/rd-vigor/repositories"
)

type CustomClaims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

func GenerateTokensAndSetCookies(user repositories.User, c echo.Context) error {

	expirationTime := time.Now().Add(24 * time.Hour)

	claims := CustomClaims{
		Username: user.Username,
		Role:     user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(GetJWTSecret())

	if err != nil {
		return err
	}

	tokenCookie := new(http.Cookie)
	tokenCookie.Name = "access-token"
	tokenCookie.Value = tokenString
	tokenCookie.Expires = expirationTime
	tokenCookie.Path = "/"
	tokenCookie.HttpOnly = true
	tokenCookie.SameSite = http.SameSiteStrictMode

	c.SetCookie(tokenCookie)

	userCookie := new(http.Cookie)
	userCookie.Name = "user"
	userCookie.Value = user.Username
	userCookie.Expires = expirationTime
	userCookie.Path = "/"
	userCookie.HttpOnly = true
	userCookie.SameSite = http.SameSiteStrictMode

	c.SetCookie(userCookie)

	return nil
}

func ResetAuthCookies(c echo.Context) {

	tokenCookie := new(http.Cookie)
	tokenCookie.Name = "access-token"
	tokenCookie.Value = ""
	tokenCookie.Path = "/"
	tokenCookie.Expires = time.Unix(0, 0)
	tokenCookie.HttpOnly = true

	c.SetCookie(tokenCookie)

	userCookie := new(http.Cookie)
	userCookie.Name = "user"
	userCookie.Value = ""
	userCookie.Expires = time.Unix(0, 0)
	userCookie.Path = "/"
	userCookie.HttpOnly = true

	c.SetCookie(userCookie)
}

func DecodeToken(tokenString string) (*CustomClaims, error) {

	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJWTSecret(), nil
	})

	claims, ok := token.Claims.(*CustomClaims)

	if ok && token.Valid {
		return claims, nil
	}
	return nil, err
}

func JWTErrorChecker(err error, c echo.Context) error {
	return c.Redirect(http.StatusMovedPermanently, c.Echo().Reverse("userSignInForm"))
}

func GetJWTSecret() []byte {
	return []byte("secret")
}

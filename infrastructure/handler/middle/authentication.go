package middle

import (
	"errors"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"

	"github.com/alexyslozada/ecommerce/infrastructure/handler/response"
	"github.com/alexyslozada/ecommerce/model"
)

type AuthMiddleware struct {
	responser response.API
}

func New() AuthMiddleware {
	return AuthMiddleware{}
}

func (am AuthMiddleware) IsValid(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, err := getTokenFromRequest(c.Request())
		if err != nil {
			return am.responser.BindFailed(err)
		}

		isValid, claims := am.validate(token)
		if !isValid {
			err = errors.New("the token is not valid")
			return am.responser.BindFailed(err)
		}

		c.Set("userID", claims.UserID)
		c.Set("email", claims.Email)
		c.Set("isAdmin", claims.IsAdmin)

		return next(c)
	}
}

func (am AuthMiddleware) IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		isAdmin, ok := c.Get("isAdmin").(bool)
		if !isAdmin || !ok {
			err := errors.New("you are not admin")
			return am.responser.BindFailed(err)
		}

		return next(c)
	}
}

func (am AuthMiddleware) validate(token string) (bool, model.JWTCustomClaims) {
	claims, err := jwt.ParseWithClaims(token, &model.JWTCustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		log.Println(token)
		log.Println(os.Getenv("JWT_SECRET_KEY"))
		log.Println(err)
		return false, model.JWTCustomClaims{}
	}

	data, ok := claims.Claims.(*model.JWTCustomClaims)
	if !ok {
		log.Println("is not a jwtcustomclaims")
		return false, model.JWTCustomClaims{}
	}

	return true, *data
}

func getTokenFromRequest(r *http.Request) (string, error) {
	data := r.Header.Get("Authorization")
	if data == "" {
		return "", errors.New("el header de autorización está vacío")
	}

	if strings.HasPrefix(data, "Bearer") {
		return data[7:], nil
	}

	return data, nil
}

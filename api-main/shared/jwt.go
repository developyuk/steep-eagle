package shared

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

type Auth struct {
	Id   float64
	Role string
}

var CurrentAuth Auth

func GetAuth(c echo.Context) Auth {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return Auth{
		Id:   claims["id"].(float64),
		Role: claims["role"].(string),
	}
}

func GetAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {
    CurrentAuth = GetAuth(c)
		// c.Response().Header().Set(echo.HeaderServer, "Echo/3.0")
		return next(c)
	}
}

package jwt

import (
  "github.com/labstack/echo"
  "github.com/dgrijalva/jwt-go"
  "os"
)

type (
  Auth struct {
    Id   float64
    Role string
  }
)

var (
  Key         = os.Getenv("JWT_KEY")
  CurrentAuth Auth
  AuthHeader  string
)

func GetAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {

  return func(c echo.Context) error {
    AuthHeader = c.Request().Header.Get("Authorization")
    user := c.Get("user").(*jwt.Token)
    claims := user.Claims.(jwt.MapClaims)

    CurrentAuth = Auth{
      Id:   claims["id"].(float64),
      Role: claims["role"].(string),
    }
    // c.Response().Header().Set(echo.HeaderServer, "Echo/3.0")
    return next(c)
  }
}

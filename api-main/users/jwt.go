package users

import (
  // "fmt"
  // "github.com/jmoiron/sqlx"
  myShared "../shared"
  "github.com/dgrijalva/jwt-go"
  "github.com/labstack/echo"
  "net/http"
  "time"
  "log"
)

type UserLoginRequest struct {
  Email string `json:"email" xml:"email" form:"email" query:"email"`
  Pwd   string `json:"pwd" xml:"pwd" form:"pwd" query:"pwd"`
}

func Sign(c echo.Context) error {
  p := new(UserLoginRequest)
  if err := c.Bind(p); err != nil {
    return err
  }
  // email := c.FormValue("email")
  // pwd := c.FormValue("pwd")
  item, err := itemByEmail(p)
  if err != nil {
    return err
  }
  // Create token
  token := jwt.New(jwt.SigningMethodHS256)

  // Set claims
  claims := token.Claims.(jwt.MapClaims)
  claims["id"] = item.Id
  claims["name"] = item.UsersProfile[0].Name
  claims["role"] = item.Role
  claims["photo"] = item.UsersProfile[0].Photo
  claims["email"] = item.Email
  claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

  // Generate encoded token and send it as response.
  t, err2 := token.SignedString([]byte(myShared.JwtKey))
  log.Println(t)
  if err2 != nil {
    return c.JSON(http.StatusUnauthorized, err2)
  }
  return c.JSON(http.StatusOK, map[string]string{
    "token": t,
  })

  return echo.ErrUnauthorized
}
func Auth(c echo.Context) error {
  user := c.Get("user").(*jwt.Token)
  claims := user.Claims.(jwt.MapClaims)
  log.Println(claims)
  return c.JSON(http.StatusOK, map[string]interface{}{
    "id": claims["id"].(float64),
    "name":  claims["name"].(string),
    "role": claims["role"].(string),
    "photo": claims["photo"].(string),
    "email": claims["email"].(string),
  })
}

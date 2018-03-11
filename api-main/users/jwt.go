package users

import (
  // "fmt"
  // "github.com/jmoiron/sqlx"
  myShared "../shared"
  myJwt "../shared/jwt"
  "github.com/dgrijalva/jwt-go"
  "github.com/labstack/echo"
  "net/http"
  "time"
  "log"
)

type UserLoginRequest struct {
  Username string `json:"username" xml:"username" form:"username" query:"username"`
}

func Sign(c echo.Context) error {
  p := new(UserLoginRequest)
  if err := c.Bind(p); err != nil {
    return err
  }
  item, err := itemByUsername(p)

  if err != nil {
    return c.JSON(http.StatusUnauthorized, myShared.Response{Message: err.Error()})
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
  t, err2 := token.SignedString([]byte(myJwt.Key))
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
    "id":    claims["id"].(float64),
    "name":  claims["name"].(string),
    "role":  claims["role"].(string),
    "photo": claims["photo"].(string),
    "email": claims["email"].(string),
  })
}

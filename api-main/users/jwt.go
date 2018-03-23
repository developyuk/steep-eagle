package users

import (
  myShared "../shared"
  myJwt "../shared/jwt"
  //mySessionsLogins "../sessions/logins"
  "github.com/dgrijalva/jwt-go"
  "github.com/labstack/echo"
  "net/http"
  "time"
  "github.com/davecgh/go-spew/spew"
)

type (
  UserLoginRequest struct {
    Username string `json:"username" xml:"username" form:"username" query:"username"`
  }
)

func Sign(c echo.Context) error {
  p := new(UserLoginRequest)
  if err := c.Bind(p); err != nil {
    return err
  }
  //var userId uint64

  // Create token
  token := jwt.New(jwt.SigningMethodHS256)
  if item, err := itemByUsername(p); err != nil {
    return c.JSON(http.StatusUnauthorized, myShared.CreateResponse(err.Error()))
  } else {
    spew.Dump(item)

    // Set claims
    claims := token.Claims.(jwt.MapClaims)
    //userId = item.Id
    claims["id"] = item.Id
    claims["username"] = item.Username
    claims["name"] = item.Name
    claims["role"] = item.Role
    claims["photo"] = item.Photo
    claims["email"] = item.Email
    claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
  }

  // Generate encoded token and send it as response.
  if t, err2 := token.SignedString([]byte(myJwt.Key)); err2 != nil {
    return c.JSON(http.StatusUnauthorized, err2)
  } else {
    //mySessionsLogins.Create(userId)
    return c.JSON(http.StatusOK, map[string]string{
      "token": t,
    })
  }

  return echo.ErrUnauthorized
}
func Auth(c echo.Context) error {
  user := c.Get("user").(*jwt.Token)
  claims := user.Claims.(jwt.MapClaims)
  return c.JSON(http.StatusOK, map[string]interface{}{
    "id":       claims["id"].(float64),
    "username": claims["username"].(string),
    "name":     claims["name"].(string),
    "role":     claims["role"].(string),
    "photo":    claims["photo"].(string),
    "email":    claims["email"].(string),
  })
}

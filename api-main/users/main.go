package users

import (
  myShared "../shared"
  mySharedRest "../shared/rest"
  "github.com/labstack/echo"
  "net/http"
)

func List(c echo.Context) error {
  req := new(mySharedRest.Request)
  if err := c.Bind(req); err != nil {
    return c.JSON(http.StatusBadRequest, myShared.CreateResponse(err.Error()))
  }
  rest := mySharedRest.New().GetItems("/_users_profile").ParseRequest(req)
  role := getRole(c.Path())
  if role != "user" {
    rest.SetQuery("role=eq." + role)
  }

  var list []User
  if resp, err := rest.EndStruct(&list); err != nil {
    return c.JSON(resp.StatusCode, myShared.CreateResponse(err.Error()))
  }

  // log.Fatal()
  for i := range list {
    list[i].setItemLinks(getPath(role))
  }

  response := myShared.Hal{
    Links:    myShared.CreateHalLinks(c.Request().RequestURI, c.Path(), rest),
    Embedded: myShared.CreateEmbeddedItems(list),
    Count:    rest.Count,
    Total:    rest.Total,
  }
  return c.JSON(http.StatusOK, response)
}

func Item(c echo.Context) error {
  req := new(mySharedRest.Request)
  if err := c.Bind(req); err != nil {
    return c.JSON(http.StatusBadRequest, myShared.CreateResponse(err.Error()))
  }

  role := getRole(c.Path())

  var item User
  if resp, err := ItemRest(req, role, c.Param("id"), &item); err != nil {
    return c.JSON(resp.StatusCode, myShared.CreateResponse(err.Error()))
  }
  return c.JSON(http.StatusOK, item)
}

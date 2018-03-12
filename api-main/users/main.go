package users

import (
  myShared "../shared"
  mySharedRest "../shared/rest"
  "github.com/labstack/echo"
  "net/http"
  "strings"
)

func getRole(path string) string {
  role := strings.TrimRight(strings.TrimPrefix(path, "/"), "s")
  role = strings.TrimRight(strings.TrimPrefix(path, "/"), "s/:id")
  return role
}

func getPath(role string) string {
  path := PathUsers
  if role == "student" {
    path = PathStudents
  }
  if role == "tutor" {
    path = PathTutors
  }
  return path
}

func List(c echo.Context) error {
  req := new(mySharedRest.Request)
  if err := c.Bind(req); err != nil {
    return c.JSON(http.StatusBadRequest, myShared.CreateResponse(err.Error()))
  }
  rest := mySharedRest.New().GetItems("/users_full").ParseRequest(req)
  role := getRole(c.Path())
  if role != "user" {
    rest.SetQuery("role=eq." + role)
  }

  var list []myShared.User
  if resp, err := rest.End(&list); err != nil {
    return c.JSON(resp.StatusCode, myShared.CreateResponse(err.Error()))
  }

  // log.Fatal()
  for i, v := range list {
    list[i].Links = itemLinks(v, role)
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
  rest := mySharedRest.New().GetItem("/users_full").ParseRequest(req)
  role := getRole(c.Path())
  if role != "user" {
    rest.SetQuery("role=eq." + role)
  }

  var item myShared.User
  if resp, err := rest.SetQuery("id=eq." + c.Param("id")).End(&item); err != nil {
    return c.JSON(resp.StatusCode, myShared.CreateResponse(err.Error()))
  }

  item.Links = itemLinks(item, role)
  return c.JSON(http.StatusOK, item)
}

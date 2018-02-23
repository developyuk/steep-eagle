package users

import (
  myShared "../shared"
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
  path := myShared.PathUsers
  if role == "student" {
    path = myShared.PathStudents
  }
  if role == "tutor" {
    path = myShared.PathTutors
  }
  return path
}

func List(c echo.Context) error {
  params := make(map[string]string)
  role := getRole(c.Path())

  if role != "user" {
    params["role"] = "eq." + role
  }
  var list []User
  resp, err := myShared.GetItems(map[string]interface{}{
    "data":  &list,
    "path":  myShared.PathUsers,
    "query": params,
  })
  if err != nil {
    return c.JSON(resp.StatusCode, myShared.Response{
      Message: err.Error(),
    })
  }

  // log.Fatal()
  for i, v := range list {
    list[i].Links = itemLinks(v, role)
  }

  response := myShared.Hal{
    Links:    myShared.LinksSelf{Self: myShared.CreateHref(getPath(role))},
    Embedded: list,
    Count:    len(list),
    Total:    len(list),
  }
  return c.JSON(http.StatusOK, response)
}

func Item(c echo.Context) error {
  // var data Program = GetProgramTypesData(c.Param("id"))
  params := make(map[string]string)
  params["id"] = "eq." + c.Param("id")

  role := getRole(c.Path())
  if role != "user" {
    params["role"] = "eq." + role
  }

  var item User
  resp, err := myShared.GetItem(map[string]interface{}{
    "data":  &item,
    "path":  myShared.PathUsers,
    "query": params,
  })

  if err != nil {
    return c.JSON(resp.StatusCode, myShared.Response{
      Message: err.Error(),
    })
  }

  item.Links = itemLinks(item, role)
  return c.JSON(http.StatusOK, item)
}

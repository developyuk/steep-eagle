package modules

import (
  myShared "../shared"
  "fmt"
  "github.com/labstack/echo"
  "net/http"
  "gopkg.in/resty.v1"
  "encoding/json"
)

func itemLinks(v Module) myShared.LinksSelf {
  return myShared.LinksSelf{Self: myShared.Href{
    Href: fmt.Sprintf("%v/%v", myShared.PathModules, v.Id),
  }}
}

var (
  auth string
)

func List(c echo.Context) error {
  auth = c.Request().Header.Get("Authorization")

  var list []Module

  resp, _ := resty.R().
    SetHeader("Accept", "application/json").
    SetHeader("Authorization", auth).
    Get(myShared.DbApiUrl + "/modules")
  _ = json.Unmarshal(resp.Body(), &list)

  for i, v := range list {
    list[i].Links = itemLinks(v)
  }

  response := myShared.Hal{
    Links:    myShared.LinksSelf{Self: myShared.Href{Href: myShared.PathModules}},
    Embedded: list,
    Count:    len(list),
    Total:    len(list),
  }

  return c.JSON(http.StatusOK, response)
}

func Item(c echo.Context) error {
  auth = c.Request().Header.Get("Authorization")

  // var data Program = GetProgramTypesData(c.Param("id"))
  var item Module
  resp, _ := resty.R().
    SetQueryParams(map[string]string{
    "id": "eq." + c.Param("id"),
  }).
    SetHeader("Accept", "application/vnd.pgrst.object+json").
    SetHeader("Authorization", c.Request().Header.Get("Authorization")).
    Get(myShared.DbApiUrl + "/modules")
  _ = json.Unmarshal(resp.Body(), &item)

  item.Links = itemLinks(item)
  return c.JSON(http.StatusOK, item)
}

// func UpdateModule(c echo.Context) error {
// 	var data myModels.ProgramType
// 	return c.JSON(http.StatusOK, data)
// }
//
// func DeleteModule(c echo.Context) error {
// 	// id := c.Param("id")
// 	var data myModels.ProgramType
// 	return c.JSON(http.StatusOK, data)
// }

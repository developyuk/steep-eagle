package types

import (
  myShared "../../shared"
  mySharedRest "../../shared/rest"
  "github.com/labstack/echo"
  "net/http"
)

func List(c echo.Context) error {
  req := new(mySharedRest.Request)
  if err := c.Bind(req); err != nil {
    return c.JSON(http.StatusBadRequest, myShared.CreateResponse(err.Error()))
  }
  var list []ProgramType
  rest := mySharedRest.New().GetItems("/program_types").ParseRequest(req)
  if resp, err := rest.End(&list); err != nil {
    return c.JSON(resp.StatusCode, myShared.CreateResponse(err.Error()))
  }

  for i, v := range list {
    list[i].Links = itemLinks(v)
  }

  response := myShared.Hal{
    Links:    myShared.CreateHalLinks(c.Request().RequestURI, Path, rest),
    Embedded: list,
    Count:    rest.Total,
    Total:    rest.Count,
  }
  return c.JSON(http.StatusOK, response)
}

func Item(c echo.Context) error {
  req := new(mySharedRest.Request)
  if err := c.Bind(req); err != nil {
    return c.JSON(http.StatusBadRequest, myShared.CreateResponse(err.Error()))
  }
  var item ProgramType

  if resp, err := mySharedRest.New().GetItem("/program_types").ParseRequest(req).
    SetQuery("id=eq." + c.Param("id")).
    End(&item); err != nil {
    return c.JSON(resp.StatusCode, myShared.CreateResponse(err.Error()))
  }

  item.Links = itemLinks(item)
  return c.JSON(http.StatusOK, item)
}

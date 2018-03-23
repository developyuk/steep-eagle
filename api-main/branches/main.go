package branches

import (
  mySharedRest "../shared/rest"
  myShared "../shared"
  "github.com/labstack/echo"
  "net/http"
)

func List(c echo.Context) error {
  req := new(mySharedRest.Request)
  if err := c.Bind(req); err != nil {
    return c.JSON(http.StatusBadRequest, myShared.CreateResponse(err.Error()))
  }
  var list []Branch
  rest := mySharedRest.New().GetItems(Path).ParseRequest(req)
  if resp, err := rest.
  SetQuery(myShared.RequestRest{
    Select: req.Select,
  }).EndStruct(&list); err != nil {
    return c.JSON(resp.StatusCode, myShared.CreateResponse(err.Error()))
  }

  for i := range list {
    list[i].setItemLinks()
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

  var item Branch
  if resp, err := ItemRest(req, c.Param("id"), &item); err != nil {
    return c.JSON(resp.StatusCode, myShared.CreateResponse(err.Error()))
  }

  return c.JSON(http.StatusOK, item)
}

package programs

import (
  myShared "../shared"
  myRest "../shared/rest"
  "github.com/labstack/echo"
  "net/http"
)

func List(c echo.Context) error {
  var list []Program
  resp, err := myRest.GetItems(map[string]interface{}{
    "data": &list,
    "path": myShared.PathPrograms,
  })
  if err != nil {
    return c.JSON(resp.StatusCode, myShared.Response{
      Message: err.Error(),
    })
  }

  for i, v := range list {
    list[i].Links = itemLinks(v)
  }
  response := myShared.Hal{
    Links:    myShared.LinksSelf{Self: myShared.Href{Href: myShared.PathPrograms}},
    Embedded: list,
    Count:    len(list),
    Total:    uint64(len(list)),
  }

  return c.JSON(http.StatusOK, response)
}

func Item(c echo.Context) error {
  var item Program
  resp, err := myRest.GetItem(map[string]interface{}{
    "data": &item,
    "path": myShared.PathPrograms,
    "query": map[string]string{
      "id": "eq." + c.Param("id"),
    },
  })
  if err != nil {
    return c.JSON(resp.StatusCode, myShared.Response{
      Message: err.Error(),
    })
  }

  item.Links = itemLinks(item)
  return c.JSON(http.StatusOK, item)
}

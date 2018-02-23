package programs

import (
  myShared "../shared"
  "github.com/labstack/echo"
  "net/http"
)

func List(c echo.Context) error {
  var list []Program
  resp, err := myShared.GetItems(map[string]interface{}{
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
    Total:    len(list),
  }

  return c.JSON(http.StatusOK, response)
}

func Item(c echo.Context) error {
  var item Program
  resp, err := myShared.GetItem(map[string]interface{}{
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

// func UpdateProgram(c echo.Context) error {
// 	var data myModels.Program
// 	return c.JSON(http.StatusOK, data)
// }
//
// func DeleteProgram(c echo.Context) error {
// 	// id := c.Param("id")
// 	var data myModels.Program
// 	return c.JSON(http.StatusOK, data)
// }

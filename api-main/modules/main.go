package modules

import (
  myShared "../shared"
  "github.com/labstack/echo"
  "net/http"
)

func List(c echo.Context) error {
  var list []Module
  resp, err := myShared.GetItems(map[string]interface{}{
    "data": &list,
    "path": myShared.PathModules,
  })
  if err != nil {
    return c.JSON(resp.StatusCode, myShared.Response{
      Message: err.Error(),
    })
  }

  for i, v := range list {
    list[i].Links = ItemLinks(v)
  }

  response := myShared.Hal{
    Links:    myShared.LinksSelf{Self: myShared.CreateHref(myShared.PathModules)},
    Embedded: list,
    Count:    len(list),
    Total:    len(list),
  }

  return c.JSON(http.StatusOK, response)
}

func Item(c echo.Context) error {
  // var data Program = GetProgramTypesData(c.Param("id"))
  var item Module
  resp, err := myShared.GetItem(map[string]interface{}{
    "data": &item,
    "path": myShared.PathModules,
    "query": map[string]string{
      "id": "eq." + c.Param("id"),
    },
  })
  if err != nil {
    return c.JSON(resp.StatusCode, myShared.Response{
      Message: err.Error(),
    })
  }

  item.Links = ItemLinks(item)
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

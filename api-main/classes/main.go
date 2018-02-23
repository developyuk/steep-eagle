package classes

import (
  myShared "../shared"
  "github.com/labstack/echo"
  "net/http"
)

func List(c echo.Context) error {
  //params := make(map[string]string)

  var list []Class_
  resp, err := myShared.GetItems(map[string]interface{}{
    "data": &list,
    "path": "/classes_",
    "query": map[string]string{
      "order": "ts.asc",
    },
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
    Links:    myShared.LinksSelf{Self: myShared.CreateHref(myShared.PathClasses)},
    Embedded: list,
    Count:    len(list),
    Total:    len(list),
  }
  return c.JSON(http.StatusOK, response)
}

func Item(c echo.Context) error {
  var item Class_
  resp, err := myShared.GetItem(map[string]interface{}{
    "data": &item,
    "path": "/classes_",
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

// func UpdateClass(c echo.Context) error {
// 	var data myModels.ProgramType
// 	return c.JSON(http.StatusOK, data)
// }
//
// func DeleteClass(c echo.Context) error {
// 	// id := c.Param("id")
// 	var data myModels.ProgramType
// 	return c.JSON(http.StatusOK, data)
// }

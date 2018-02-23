package branches

import (
  myShared "../shared"
  "github.com/labstack/echo"
  "net/http"
)

func List(c echo.Context) error {
  var list []Branch
  resp, err := myShared.GetItems(map[string]interface{}{
    "data": &list,
    "path": myShared.PathBranches,
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
    Links:    myShared.LinksSelf{Self: myShared.CreateHref(myShared.PathBranches)},
    Embedded: list,
    Count:    len(list),
    Total:    len(list),
  }
  return c.JSON(http.StatusOK, response)
}

func Item(c echo.Context) error {
  var item Branch
  resp, err := myShared.GetItem(map[string]interface{}{
    "data": &item,
    "path": myShared.PathBranches,
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

// func UpdateBranch(c echo.Context) error {
// 	var data Branch
// 	return c.JSON(http.StatusOK, data)
// }
//
// func DeleteBranch(c echo.Context) error {
// 	// id := c.Param("id")
// 	var data Branch
// 	return c.JSON(http.StatusOK, data)
// }

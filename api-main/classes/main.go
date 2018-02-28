package classes

import (
  myShared "../shared"
  "github.com/labstack/echo"
  "net/http"
  "log"
)

func List(c echo.Context) error {
  params := make(map[string]string)

  if val := c.QueryParam("sort"); len(val) > 0 {
    params["order"] = val
  }
  log.Println(params)
  var list []myShared.Class_
  resp, err := myShared.GetItems(map[string]interface{}{
    "data":  &list,
    "path":  "/classes_ts",
    "query": params,
  })
  if err != nil {
    return c.JSON(resp.StatusCode, myShared.Response{
      Message: err.Error(),
    })
  }
  //list2 := list
  //d := 24 * time.Hour
  //list2 := list[:0]
  for i, v := range list {
    list[i].Links = myShared.ClassItemLinks(v)
    list[i].Embedded = itemEmbedded(v)

    //embedded := list[i].Embedded.(ClassEmbedded)
    //if !(embedded.LastSession != nil && embedded.LastSession.CreatedAt.Truncate(d).Equal(time.Now().Truncate(d)) ) {
    //  list2 = append(list2, list[i])
    //}
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
  var item myShared.Class_
  resp, err := myShared.GetItem(map[string]interface{}{
    "data": &item,
    "path": "/classes_ts",
    "query": map[string]string{
      "id": "eq." + c.Param("id"),
    },
  })
  if err != nil {
    return c.JSON(resp.StatusCode, myShared.Response{
      Message: err.Error(),
    })
  }

  item.Links = myShared.ClassItemLinks(item)
  item.Embedded = itemEmbedded(item)
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

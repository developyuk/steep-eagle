package classes

import (
  myShared "../shared"
  "github.com/labstack/echo"
  "net/http"
  "strconv"
  "time"
)

type (
  ClassGroupDate struct {
    myShared.Hal
    Date  string            `json:"date"`
    Day   string            `json:"day"`
    Text  string            `json:"text"`
    Items []myShared.Class_ `json:"items"`
  }
)

func list(params map[string]string) ([]myShared.Class_, error) {
  var list []myShared.Class_
  _, err := myShared.GetItems(map[string]interface{}{
    "data":  &list,
    "path":  "/classes_ts",
    "query": params,
  })
  if err != nil {
    return list, err
  }
  //list2 := list
  //d := 24 * time.Hour
  //list2 := list[:0]
  for i, v := range list {
    list[i].Links = myShared.ClassItemLinks(v)
    list[i].Embedded = myShared.ClassItemEmbedded(v)

    //embedded := list[i].Embedded.(ClassEmbedded)
    //if !(embedded.LastSession != nil && embedded.LastSession.CreatedAt.Truncate(d).Equal(time.Now().Truncate(d)) ) {
    //  list2 = append(list2, list[i])
    //}
  }
  return list, err
}
func ListGroup(c echo.Context) error {
  params := make(map[string]string)

  if val := c.QueryParam("by"); len(val) > 0 {
    params["by"] = val
  }
  list, err := list(params)
  var classGroupDate []ClassGroupDate
  var t ClassGroupDate
  for _, v := range list {
    if v.Day != t.Day {
      if t.Day != "" {
        classGroupDate = append(classGroupDate, t)
      }
      t.Date = strconv.Itoa(v.StartAtTs.Day())
      t.Day = v.Day
      t.Text = strconv.FormatFloat(time.Until(v.StartAtTs).Round(time.Hour).Hours()/24, 'f', 0, 64)
      if t.Text == "0" {
        t.Text = "Today"
      } else if t.Text == "1" {
        t.Text = "Tomorrow"
      } else {
        t.Text = "in " + t.Text + " days"
      }
      t.Items = []myShared.Class_{}
    }
    t.Items = append(t.Items, v)
  }
  if err != nil {
    return c.JSON(400, myShared.Response{
      Message: err.Error(),
    })
  }
  response := myShared.Hal{
    Links:    myShared.LinksSelf{Self: myShared.CreateHref(myShared.PathClasses)},
    Embedded: classGroupDate,
    Count:    len(classGroupDate),
    Total:    len(classGroupDate),
  }
  return c.JSON(http.StatusOK, response)
}
func List(c echo.Context) error {
  params := make(map[string]string)

  if val := c.QueryParam("sort"); len(val) > 0 {
    params["order"] = val
  }
  list, err := list(params)
  if err != nil {
    return c.JSON(400, myShared.Response{
      Message: err.Error(),
    })
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
  item.Embedded = myShared.ClassItemEmbedded(item)
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

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
  for i, v := range list {
    list[i].Links = myShared.ClassItemLinks(v)
    list[i].Embedded = myShared.ClassItemEmbedded(v)
  }
  return list, err
}
func timeToRelativeText(timeTs time.Time) string {

  text := strconv.FormatFloat(time.Until(timeTs).Round(time.Hour).Hours()/24, 'f', 0, 64)
  if text == "0" {
    text = "Today"
  } else if text == "1" {
    text = "Tomorrow"
  } else {
    text = "in " + text + " days"
  }
  return text
}
func ListGroup(c echo.Context) error {
  params := make(map[string]string)

  if val := c.QueryParam("by"); len(val) > 0 {
    params["by"] = val
  }
  list, err := list(params)
  var classGroupDates []ClassGroupDate
  var t ClassGroupDate

  for _, v := range list {
    found := false
    for i2, v2 := range classGroupDates {
      if v.Day == v2.Day {
        // Found!
        classGroupDates[i2].Items = append(classGroupDates[i2].Items,v)
        found=true
        break
      }
    }
    if !found {
      t.Date = strconv.Itoa(v.StartAtTs.Day())
      t.Day = v.Day
      t.Text = timeToRelativeText(v.StartAtTs)
      t.Items = []myShared.Class_{}
      t.Items = append(t.Items, v)
      classGroupDates = append(classGroupDates, t)
    }
  }

  if err != nil {
    return c.JSON(400, myShared.Response{
      Message: err.Error(),
    })
  }
  response := myShared.Hal{
    Links:    myShared.LinksSelf{Self: myShared.CreateHref(myShared.PathClasses)},
    Embedded: classGroupDates,
    Count:    len(classGroupDates),
    Total:    len(classGroupDates),
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

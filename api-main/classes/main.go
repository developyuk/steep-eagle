package classes

import (
  myShared "../shared"
  mySharedRest "../shared/rest"
  "github.com/labstack/echo"
  "net/http"
  "strconv"
  "time"
  //"github.com/davecgh/go-spew/spew"
)

type (
  ClassGroupDate struct {
    myShared.Hal
    Date      string    `json:"date"`
    Day       string    `json:"day"`
    StartAtTs time.Time `json:"-"`
    Text      string    `json:"text"`
    Items     []Class_  `json:"items"`
  }
)

func list(params map[string]string) (*http.Response, *mySharedRest.MyRest, []Class_, error) {
  var list []Class_
  var resp *http.Response
  path := "/_classes_ts"
  if _, ok := params["q"]; ok {
    //do something here
    path = "/_classes_ts_search"
  }
  rest := mySharedRest.New().GetItems(path)
  if resp, err := rest.SetQuery(params).EndStruct(&list); err != nil {
    //spew.Dump(params,resp,list)
    return resp, rest, list, err
  }

  for i, v := range list {
    list[i].Links = ItemLinks(&v)
  }
  return resp, rest, list, nil
}
func timeToRelativeText(timeTs time.Time) string {
  loc, _ := time.LoadLocation(myShared.TimeZone)
  timeTz := time.Now().In(loc)
  timeTsTz := timeTs.In(loc)
  text := strconv.FormatFloat((timeTsTz.Sub(timeTz).Round(time.Hour).Hours())/24, 'f', 0, 64)
  if text == "0" || text == "-0" {
    text = "today"
  } else if text == "1" {
    text = "tomorrow"
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
  if val := c.QueryParam("sort"); len(val) > 0 {
    params["order"] = val
  }
  if val := c.QueryParam("q"); len(val) > 0 {
    params["q"] = val
  }
  _, _, list, err := list(params)

  var classGroupDates []ClassGroupDate
  var t ClassGroupDate

  for _, v := range list {
    found := false
    for i2, v2 := range classGroupDates {
      if timeToRelativeText(v.StartAtTs) == timeToRelativeText(v2.StartAtTs) {
        // Found!
        classGroupDates[i2].Items = append(classGroupDates[i2].Items, v)
        found = true
        break
      }
    }
    if !found {
      t.Date = strconv.Itoa(v.StartAtTs.Day())
      t.Day = v.Day
      t.StartAtTs = v.StartAtTs
      t.Text = timeToRelativeText(v.StartAtTs)
      t.Items = []Class_{}
      t.Items = append(t.Items, v)
      classGroupDates = append(classGroupDates, t)
    }
  }

  if err != nil {
    return c.JSON(400, myShared.CreateResponse(err.Error()))
  }
  response := myShared.Hal{
    Links:    myShared.CreateHrefSelf(Path),
    Embedded: myShared.CreateEmbeddedItems(classGroupDates),
    Count:    uint64(len(classGroupDates)),
    Total:    uint64(len(classGroupDates)),
  }
  return c.JSON(http.StatusOK, response)
}
func List(c echo.Context) error {
  params := make(map[string]string)

  if val := c.QueryParam("sort"); len(val) > 0 {
    params["order"] = val
  }
  _, rest, list, err := list(params)
  if err != nil {
    return c.JSON(400, myShared.CreateResponse(err.Error()))
  }

  response := myShared.Hal{
    Links:    myShared.CreateHrefSelf(Path),
    Embedded: myShared.CreateEmbeddedItems(list),
    Count:    rest.Total,
    Total:    rest.Count,
  }
  return c.JSON(http.StatusOK, response)
}

func ItemRest(params map[string]string, id string, item *Class_) (*http.Response, error) {
  var resp *http.Response
  if resp, err := mySharedRest.New().GetItem("/_classes_ts").
    SetQuery("id=eq." + id).
    EndStruct(item); err != nil {
    return resp, err
  }

  item.Links = ItemLinks(item)
  return resp, nil
}
func Item(c echo.Context) error {
  var item Class_
  params := make(map[string]string)

  if resp, err := ItemRest(params, c.Param("id"), &item); err != nil {
    return c.JSON(resp.StatusCode, myShared.CreateResponse(err.Error()))
  }

  return c.JSON(http.StatusOK, item)
}

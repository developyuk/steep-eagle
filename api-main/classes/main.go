package classes

import (
  myShared "../shared"
  mySharedRest "../shared/rest"
  myModule "../modules"
  myBranch "../branches"
  myUser "../users"
  "github.com/labstack/echo"
  "net/http"
  "strconv"
  "time"
  //"github.com/davecgh/go-spew/spew"
  "strings"
  "fmt"
)

type (
  ClassGroupDate struct {
    myShared.Hal
    Date      string            `json:"date"`
    Day       string            `json:"day"`
    StartAtTs time.Time         `json:"-"`
    Text      string            `json:"text"`
    Items     []myShared.Class_ `json:"items"`
  }
)

func list(params map[string]string) (*http.Response, *mySharedRest.MyRest, []myShared.Class_, error) {
  var list []myShared.Class_
  var resp *http.Response
  rest := mySharedRest.New().GetItems("/classes_ts")
  embed := params["embed"]
  delete(params, "embed")

  if resp, err := rest.SetQuery(params).End(&list); err != nil {
    //spew.Dump(params,resp,list)
    return resp, rest, list, err
  }

  for i, v := range list {
    list[i].Links = ItemLinks(v)
    if len(embed) > 0 {
      classEmbedded := embedded{}
      if strings.Contains(embed, "module") {
        var module myModule.Module
        _, _ = myModule.ItemRest(&mySharedRest.Request{}, fmt.Sprint(v.ModuleId), &module)
        classEmbedded.Module = module
      }
      if strings.Contains(embed, "branch") {
        var branch myBranch.Branch
        _, _ = myBranch.ItemRest(&mySharedRest.Request{}, fmt.Sprint(v.BranchId), &branch)
        classEmbedded.Branch = branch
      }
      if strings.Contains(embed, "tutor") {
        var tutor myUser.User
        _, _ = myUser.ItemRest(&mySharedRest.Request{}, "tutor", fmt.Sprint(v.TutorId), &tutor)
        classEmbedded.Tutor = tutor
      }
      list[i].Embedded = classEmbedded
    }
  }
  return resp, rest, list, nil
}
func timeToRelativeText(timeTs time.Time) string {
  //loc,_ := time.LoadLocation("Asia/Jakarta")
  text := strconv.FormatFloat((time.Until(timeTs).Round(time.Hour).Hours()+7)/24, 'f', 0, 64)
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
  if val := c.QueryParam("embed"); len(val) > 0 {
    params["embed"] = val
  }
  if val := c.QueryParam("sort"); len(val) > 0 {
    params["order"] = val
  }
  _, _, list, err := list(params)

  var classGroupDates []ClassGroupDate
  var t ClassGroupDate

  for _, v := range list {
    found := false
    for i2, v2 := range classGroupDates {
      if v.StartAtTs.Day() == v2.StartAtTs.Day() {
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
      t.Items = []myShared.Class_{}
      t.Items = append(t.Items, v)
      classGroupDates = append(classGroupDates, t)
    }
  }

  if err != nil {
    return c.JSON(400, myShared.CreateResponse(err.Error()))
  }
  response := myShared.Hal{
    Links:    myShared.LinksSelf{Self: myShared.CreateHref(myShared.PathClasses)},
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
    Links:    myShared.LinksSelf{Self: myShared.CreateHref(myShared.PathClasses)},
    Embedded: myShared.CreateEmbeddedItems(list),
    Count:    rest.Total,
    Total:    rest.Count,
  }
  return c.JSON(http.StatusOK, response)
}

func Item(c echo.Context) error {
  var item myShared.Class_

  if resp, err := mySharedRest.New().GetItem("/classes_ts").
    SetQuery("id=eq." + c.Param("id")).
    End(&item); err != nil {
    return c.JSON(resp.StatusCode, myShared.CreateResponse(err.Error()))
  }

  item.Links = ItemLinks(item)
  //item.Embedded = ItemEmbedded(item)
  return c.JSON(http.StatusOK, item)
}

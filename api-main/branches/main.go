package branches

import (
  myShared "../shared"
  myClasses "../classes"
  "fmt"
  "github.com/labstack/echo"
  "net/http"
  "gopkg.in/resty.v1"
  "encoding/json"
  "strconv"
)

type BranchLinks struct {
  myShared.LinksSelf
  Classes []myShared.Href `json:"classes,omitempty"`
}

var (
  auth string
)

func itemClassesHref(id int) []myShared.Href {
  //classes := getClassesById(db, id)
  var classes []myClasses.Class_
  resp, _ := resty.R().
    SetQueryParams(map[string]string{
    "branch_id": "eq." + strconv.Itoa(id),
    "select":    "id",
  }).
    SetHeader("Accept", "application/json").
    SetHeader("Authorization", auth).
    Get(myShared.DbApiUrl + "/classes")
  _ = json.Unmarshal(resp.Body(), &classes)

  var data []myShared.Href
  for _, v := range classes {
    data = append(data, myShared.CreateHref(fmt.Sprintf("%v/%v", myShared.PathClasses, v.Id)))
  }

  return data
}

func itemLinks(v Branch) BranchLinks {
  return BranchLinks{
    LinksSelf: myShared.LinksSelf{Self: myShared.CreateHref(fmt.Sprintf("%v/%v", myShared.PathBranches, v.Id))},
    Classes:   itemClassesHref(v.Id),
  }
}
func List(c echo.Context) error {
  auth = c.Request().Header.Get("Authorization")

  var data []Branch
  resp, _ := resty.R().
    SetHeader("Accept", "application/json").
    SetHeader("Authorization", auth).
    Get(myShared.DbApiUrl + "/branches")
  _ = json.Unmarshal(resp.Body(), &data)

  for i, v := range data {
    data[i].Links = itemLinks(v)
  }
  response := myShared.Hal{
    Links:    myShared.LinksSelf{Self: myShared.Href{Href: myShared.PathBranches}},
    Embedded: data,
    Count:    len(data),
    Total:    len(data),
  }
  return c.JSON(http.StatusOK, response)
}

func Item(c echo.Context) error {
  auth = c.Request().Header.Get("Authorization")

  var item Branch
  resp, _ := resty.R().
    SetQueryParams(map[string]string{
    "id": "eq." + c.Param("id"),
  }).
    SetHeader("Accept", "application/vnd.pgrst.object+json").
    SetHeader("Authorization", auth).
    Get(myShared.DbApiUrl + "/branches")
  _ = json.Unmarshal(resp.Body(), &item)

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

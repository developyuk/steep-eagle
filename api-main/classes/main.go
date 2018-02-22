package classes

import (
  myShared "../shared"
  myUsers "../users"
  mySessions "../sessions"
  "fmt"
  "github.com/labstack/echo"
  "net/http"
  "gopkg.in/resty.v1"
  "encoding/json"
  "strconv"
)

type ClassLinks struct {
  myShared.LinksSelf
  Module   myShared.Href   `json:"module"`
  Branch   myShared.Href   `json:"branch"`
  Students []myShared.Href `json:"students,omitempty"`
  Sessions []myShared.Href `json:"sessions,omitempty"`
}

var (
  auth string
)

func itemSessionsHref(id int) []myShared.Href {
  var sessions []mySessions.Session

  resp, _ := resty.R().
    SetQueryParams(map[string]string{
    "class_id": "eq." + strconv.Itoa(id),
    "select":   "id",
  }).
    SetHeader("Accept", "application/json").
    SetHeader("Authorization", auth).
    Get(myShared.DbApiUrl + "/sessions")
  _ = json.Unmarshal(resp.Body(), &sessions)

  var data []myShared.Href
  for _, v := range sessions {
    data = append(data, myShared.CreateHref(fmt.Sprintf("%v/%v", myShared.PathSessions, v.Id)))
  }

  return data
}
func itemStudentsHref(id int) []myShared.Href {
  var students []myUsers.User

  resp, _ := resty.R().
    SetQueryParams(map[string]string{
    "class_id": "eq." + strconv.Itoa(id),
    "select":   "id",
  }).
    SetHeader("Accept", "application/json").
    SetHeader("Authorization", auth).
    Get(myShared.DbApiUrl + "/class_students")
  _ = json.Unmarshal(resp.Body(), &students)

  var data []myShared.Href
  for _, v := range students {
    data = append(data, myShared.CreateHref(fmt.Sprintf("%v/%v", myShared.PathStudents, v.Id.Int64)))
  }

  return data
}

func itemLinks(data Class_) ClassLinks {
  return ClassLinks{
    LinksSelf: myShared.LinksSelf{Self: myShared.CreateHref(fmt.Sprintf("%v/%v", myShared.PathClasses, data.Id))},
    Module:    myShared.CreateHref(fmt.Sprintf("%v/%v", myShared.PathModules, data.ModuleId)),
    Branch:    myShared.CreateHref(fmt.Sprintf("%v/%v", myShared.PathBranches, data.BranchId)),
    Students:  itemStudentsHref(data.Id),
    Sessions:  itemSessionsHref(data.Id),
  }
}

func List(c echo.Context) error {
  //params := make(map[string]string)

  var data []Class_
  resp, _ := resty.R().
    SetQueryParams(map[string]string{
    "order": "ts.asc",
  }).
    SetHeader("Accept", "application/json").
    SetHeader("Authorization", auth).
    Get(myShared.DbApiUrl + "/classes_")
  _ = json.Unmarshal(resp.Body(), &data)

  for i, v := range data {
    data[i].Links = itemLinks(v)
  }

  response := myShared.Hal{
    Links:    myShared.LinksSelf{Self: myShared.CreateHref(myShared.PathClasses)},
    Embedded: data,
    Count:    len(data),
    Total:    len(data),
  }
  return c.JSON(http.StatusOK, response)
}

func Item(c echo.Context) error {
  // var data Program = GetProgramTypesData(c.Param("id"))
  auth = c.Request().Header.Get("Authorization")
  var item Class_
  resp, _ := resty.R().
    SetQueryParams(map[string]string{
    "id": "eq." + c.Param("id"),
  }).
    SetHeader("Accept", "application/vnd.pgrst.object+json").
    SetHeader("Authorization", auth).
    Get(myShared.DbApiUrl + "/classes_")
  _ = json.Unmarshal(resp.Body(), &item)

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

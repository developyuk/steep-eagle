package classes

import (
  myShared "../shared"
  "fmt"
  "github.com/jmoiron/sqlx"
  "github.com/labstack/echo"
  "net/http"
)

type ClassLinks struct {
  myShared.LinksSelf
  Module   myShared.Href   `json:"module"`
  Branch   myShared.Href   `json:"branch"`
  Students []myShared.Href `json:"students,omitempty"`
  Sessions []myShared.Href `json:"sessions,omitempty"`
}

var (
  db   *sqlx.DB
  data []Class_
  item Class_
)

func itemSessionsHref(db *sqlx.DB, id int) []myShared.Href {
  sessions := getSessionsById(db, id)
  var data []myShared.Href
  for _, v := range sessions {
    data = append(data, myShared.CreateHref(fmt.Sprintf("%v/%v", myShared.PathSessions, v)))
  }

  return data
}
func itemStudentsHref(id int) []myShared.Href {
  students := getStudentsById(db, id)
  var data []myShared.Href
  for _, v := range students {
    data = append(data, myShared.CreateHref(fmt.Sprintf("%v/%v", myShared.PathStudents, v)))
  }

  return data
}

func itemLinks(data Class_) ClassLinks {
  return ClassLinks{
    LinksSelf: myShared.LinksSelf{Self: myShared.CreateHref(fmt.Sprintf("%v/%v", myShared.PathClasses, data.Id))},
    Module:    myShared.CreateHref(fmt.Sprintf("%v/%v", myShared.PathModules, data.ModuleId)),
    Branch:    myShared.CreateHref(fmt.Sprintf("%v/%v", myShared.PathBranches, data.BranchId)),
    Students:  itemStudentsHref(data.Id),
    Sessions:  itemSessionsHref(db, data.Id),
  }
}

func List(c echo.Context) error {
  params := make(map[string]interface{})

  if sort := c.QueryParam("sort"); sort != "" {
    params["sort"] = sort
  }

  if q := c.QueryParam("q"); q != "" {
    params["q"] = q
  }
  data, db = ListData(params)

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
  // User ID from path `users/:id`
  // var data Program = GetProgramTypesData(c.Param("id"))
  item, db = ItemData(c.Param("id"))
  item.Links = itemLinks(item)
  return c.JSON(http.StatusOK, item)
}

// func UpdateClass(c echo.Context) error {
// 	// User ID from path `users/:id`
// 	var data myModels.ProgramType
// 	return c.JSON(http.StatusOK, data)
// }
//
// func DeleteClass(c echo.Context) error {
// 	// User ID from path `users/:id`
// 	// id := c.Param("id")
// 	var data myModels.ProgramType
// 	return c.JSON(http.StatusOK, data)
// }

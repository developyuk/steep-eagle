package sessions

import (
  myShared "../shared"
  "github.com/labstack/echo"
  "net/http"
)

func List(c echo.Context) error {
  //params := make(map[string]interface{})

  var list []Session
  resp, err := myShared.GetItems(map[string]interface{}{
    "data": &list,
    "path": myShared.PathSessions,
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
    Links:    myShared.LinksSelf{Self: myShared.CreateHref(myShared.PathSessions)},
    Embedded: list,
    Count:    len(list),
    Total:    len(list),
  }
  return c.JSON(http.StatusOK, response)
}

func ListByTutorId(c echo.Context) error {
  tid := c.Param("id")
  var list []Session
  resp, err := myShared.GetItems(map[string]interface{}{
    "data": &list,
    "path": myShared.PathSessions,
    "query": map[string]string{
      "tutor_id": "eq." + tid,
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
    Links:    myShared.LinksSelf{Self: myShared.CreateHref(myShared.PathTutors + "/" + tid + "/" + myShared.PathSessions)},
    Embedded: list,
    Count:    len(list),
    Total:    len(list),
  }
  return c.JSON(http.StatusOK, response)
}

func ListByClassId(c echo.Context) error {
  cid := c.Param("id")

  params := make(map[string]interface{})
  params["cid"] = cid

  var list []Session
  resp, err := myShared.GetItems(map[string]interface{}{
    "data": &list,
    "path": myShared.PathSessions,
    "query": map[string]string{
      "class_id": "eq." + c.Param("id"),
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
    Links:    myShared.LinksSelf{Self: myShared.CreateHref(myShared.PathClasses + "/" + cid + "/" + myShared.PathSessions)},
    Embedded: list,
    Count:    len(list),
    Total:    len(list),
  }
  return c.JSON(http.StatusOK, response)
}

func Item(c echo.Context) error {
  var item Session
  resp, err := myShared.GetItem(map[string]interface{}{
    "data": &item,
    "path": myShared.PathSessions,
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

func ItemStudentsBySessionId(c echo.Context) error {
  id := c.Param("id")
  sid := c.Param("sid")

  var item SessionStudent
  resp, err := myShared.GetItem(map[string]interface{}{
    "data": &item,
    "path": "/sessions_students",
    "query": map[string]string{
      "session_id": "eq." + id,
      "student_id": "eq." + sid,
    },
  })
  if err != nil {
    return c.JSON(resp.StatusCode, myShared.Response{
      Message: err.Error(),
    })
  }

  item.Links = myShared.LinksSelf{Self: myShared.CreateHref(myShared.PathSessions + "/" + id + myShared.PathStudents + "/" + sid)}

  return c.JSON(http.StatusOK, item)
}

func CreateByClassId(c echo.Context) error {
  id := c.Param("id")
  response := CreateByClassIdData(id)
  return c.JSON(http.StatusOK, response)
}

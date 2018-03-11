package sessions

import (
  myShared "../shared"
  myRest "../shared/rest"
  myJwt "../shared/jwt"
  "github.com/labstack/echo"
  "net/http"
  "strconv"
  "log"
)

func List(c echo.Context) error {
  //params := make(map[string]interface{})

  var list []myShared.Session
  resp, err := myRest.GetItems(map[string]interface{}{
    "data": &list,
    "path": myShared.PathSessions,
  })
  if err != nil {
    return c.JSON(resp.StatusCode, myShared.Response{
      Message: err.Error(),
    })
  }

  for i, v := range list {
    list[i].Links = ItemLinks(v)
    list[i].Embedded = itemEmbedded(v)
  }

  response := myShared.Hal{
    Links:    myShared.LinksSelf{Self: myShared.CreateHref(myShared.PathSessions)},
    Embedded: list,
    Count:    len(list),
    Total:    uint64(len(list)),
  }
  return c.JSON(http.StatusOK, response)
}

func ListByTutorId(c echo.Context) error {
  tid := c.Param("id")
  var list []myShared.Session
  resp, err := myRest.GetItems(map[string]interface{}{
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
    list[i].Links = ItemLinks(v)
    list[i].Embedded = itemEmbedded(v)
  }

  response := myShared.Hal{
    Links:    myShared.LinksSelf{Self: myShared.CreateHref(myShared.PathTutors + "/" + tid + myShared.PathSessions)},
    Embedded: list,
    Count:    len(list),
    Total:    uint64(len(list)),
  }
  return c.JSON(http.StatusOK, response)
}

func ListByClassId(c echo.Context) error {
  cid := c.Param("id")

  params := make(map[string]interface{})
  params["cid"] = cid

  var list []myShared.Session
  resp, err := myRest.GetItems(map[string]interface{}{
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
    list[i].Links = ItemLinks(v)
    list[i].Embedded = itemEmbedded(v)
  }
  response := myShared.Hal{
    Links:    myShared.LinksSelf{Self: myShared.CreateHref(myShared.PathClasses + "/" + cid + myShared.PathSessions)},
    Embedded: list,
    Count:    len(list),
    Total:    uint64(len(list)),
  }
  return c.JSON(http.StatusOK, response)
}

func Item(c echo.Context) error {
  var item myShared.Session
  resp, err := myRest.GetItem(map[string]interface{}{
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
  item.Links = ItemLinks(item)
  item.Embedded = itemEmbedded(item)

  return c.JSON(http.StatusOK, item)
}

func ItemStudentsBySessionId(c echo.Context) error {
  id := c.Param("id")
  sid := c.Param("sid")

  var item SessionStudent
  resp, err := myRest.GetItem(map[string]interface{}{
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

func CreateByStudentId(c echo.Context) error {
  data := make(map[string]interface{});
  if err := c.Bind(&data); err != nil {
    return c.JSON(400, myShared.Response{
      Message: err.Error(),
    })
  }
  log.Println(data["status"].(bool),strconv.FormatBool(data["status"].(bool)))
  myRest.PostItem(map[string]interface{}{
    //"data": &list,
    "path": "/sessions_students",
    "query": map[string]string{
      "session_id": c.Param("id"),
      "student_id": c.Param("sid"),
      "tutor_id": strconv.FormatFloat(myJwt.CurrentAuth.Id, 'f', 0, 64),
      "feedback": data["review"].(string),
      "rating_cognition": strconv.FormatFloat(data["cognition"].(float64), 'f', 0, 64),
      "rating_creativity": strconv.FormatFloat(data["creativity"].(float64), 'f', 0, 64),
      "rating_interaction": strconv.FormatFloat(data["interaction"].(float64), 'f', 0, 64),
      "status": strconv.FormatBool(data["status"].(bool)),
    },
  })
  return c.JSON(http.StatusOK, myShared.Response{})
}
func CreateByClassId(c echo.Context) error {
  var list []myShared.Session
  resp, err := myRest.PostItem(map[string]interface{}{
    "data": &list,
    "path": myShared.PathSessions,
    "query": map[string]string{
      "class_id": c.Param("id"),
      "tutor_id": strconv.FormatFloat(myJwt.CurrentAuth.Id, 'f', 0, 64),
    },
  })
  if err != nil {
    return c.JSON(resp.StatusCode, myShared.Response{
      Message: err.Error(),
    })
  }
  item := list[0]
  item.Links = ItemLinks(item)
  item.Embedded = itemEmbedded(item)
  return c.JSON(http.StatusOK, item)
}

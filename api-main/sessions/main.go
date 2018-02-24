package sessions

import (
  myShared "../shared"
  "github.com/labstack/echo"
  "net/http"
  "strconv"
)

func List(c echo.Context) error {
  //params := make(map[string]interface{})

  var list []myShared.Session
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
    list[i].Links = ItemLinks(v)
    list[i].Embedded = itemEmbedded(v)
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
  var list []myShared.Session
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
    list[i].Links = ItemLinks(v)
    list[i].Embedded = itemEmbedded(v)
  }

  response := myShared.Hal{
    Links:    myShared.LinksSelf{Self: myShared.CreateHref(myShared.PathTutors + "/" + tid + myShared.PathSessions)},
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

  var list []myShared.Session
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
    list[i].Links = ItemLinks(v)
    list[i].Embedded = itemEmbedded(v)
  }
  response := myShared.Hal{
    Links:    myShared.LinksSelf{Self: myShared.CreateHref(myShared.PathClasses + "/" + cid + myShared.PathSessions)},
    Embedded: list,
    Count:    len(list),
    Total:    len(list),
  }
  return c.JSON(http.StatusOK, response)
}

func Item(c echo.Context) error {
  var item myShared.Session
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
  item.Links = ItemLinks(item)
  item.Embedded = itemEmbedded(item)

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
  var list []myShared.Session
  resp, err := myShared.PostItem(map[string]interface{}{
    "data": &list,
    "path": myShared.PathSessions,
    "query": map[string]string{
      "class_id": c.Param("id"),
      "tutor_id": strconv.FormatFloat(myShared.CurrentAuth.Id, 'f', 0, 64),
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

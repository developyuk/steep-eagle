package sessions

import (
  myShared "../shared"
  mySharedRest "../shared/rest"
  myUser "../users"
  myClass "../classes"
  myJwt "../shared/jwt"
  "github.com/labstack/echo"
  "net/http"
  //"log"
  "fmt"
  "time"
)

func List(c echo.Context) error {
  var list []Session
  rest := mySharedRest.New().GetItems(myShared.PathSessions)
  if resp, err := rest.End(&list); err != nil {
    return c.JSON(resp.StatusCode, myShared.CreateResponse(err.Error()))
  }

  for i, v := range list {
    list[i].Links = ItemLinks(v)
    list[i].Embedded = itemEmbedded(v)
  }

  response := myShared.Hal{
    Links:    myShared.CreateHrefSelf(myShared.PathSessions),
    Embedded: myShared.CreateEmbeddedItems(list),
    Count:    rest.Count,
    Total:    rest.Total,
  }
  return c.JSON(http.StatusOK, response)
}

func ListByTutorId(c echo.Context) error {
  tid := c.Param("id")
  var list []Session
  rest := mySharedRest.New().GetItems(myShared.PathSessions)

  if resp, err := rest.
  SetQuery("tutor_id=eq." + tid).
    End(&list); err != nil {
    return c.JSON(resp.StatusCode, myShared.CreateResponse(err.Error()))
  }

  for i, v := range list {
    list[i].Links = ItemLinks(v)
    list[i].Embedded = itemEmbedded(v)
  }

  response := myShared.Hal{
    Links:    myShared.CreateHrefSelf(myUser.PathTutors + "/" + tid + myShared.PathSessions),
    Embedded: myShared.CreateEmbeddedItems(list),
    Count:    rest.Count,
    Total:    rest.Total,
  }
  return c.JSON(http.StatusOK, response)
}

func ListByClassId(c echo.Context) error {

  params := make(map[string]interface{})
  if val := c.Param("id"); len(val) > 0 {
    params["class_id"] = "eq." + val
  }
  if val := c.QueryParam("created_at"); len(val) > 0 {
    params["created_at"] = val
  }

  var list []Session
  rest := mySharedRest.New().GetItems("/sessions__tutors")

  if resp, err := rest.
  SetQuery(params).
    End(&list); err != nil {
    return c.JSON(resp.StatusCode, myShared.CreateResponse(err.Error()))
  }

  for i, v := range list {
    list[i].Links = ItemLinks(v)
    list[i].Embedded = itemEmbedded(v)
  }
  response := myShared.Hal{
    Links:    myShared.CreateHrefSelf(myClass.Path + "/" + c.Param("id") + myShared.PathSessions),
    Embedded: myShared.CreateEmbeddedItems(list),
    Count:    rest.Count,
    Total:    rest.Total,
  }
  return c.JSON(http.StatusOK, response)
}

func Item(c echo.Context) error {
  var item Session

  if resp, err := mySharedRest.New().GetItem(myShared.PathSessions).
    SetQuery("id=eq." + c.Param("id")).
    End(&item); err != nil {
    return c.JSON(resp.StatusCode, myShared.CreateResponse(err.Error()))
  }
  item.Links = ItemLinks(item)
  item.Embedded = itemEmbedded(item)

  return c.JSON(http.StatusOK, item)
}

//func ItemStudentsBySessionId(c echo.Context) error {
//  id := c.Param("id")
//  sid := c.Param("sid")
//
//  var item SessionStudent
//  resp, err := mySharedRest.GetItem(map[string]interface{}{
//    "data": &item,
//    "path": "/sessions_students",
//    "query": map[string]string{
//      "session_id": "eq." + id,
//      "student_id": "eq." + sid,
//    },
//  })
//  if err != nil {
//    return c.JSON(resp.StatusCode, myShared.Response{
//      Message: err.Error(),
//    })
//  }
//
//  item.Links = myShared.LinksSelf{Self: myShared.CreateHref(myShared.PathSessions + "/" + id + myShared.PathStudents + "/" + sid)}
//
//  return c.JSON(http.StatusOK, item)
//}
//
//func CreateByStudentId(c echo.Context) error {
//  data := make(map[string]interface{});
//  if err := c.Bind(&data); err != nil {
//    return c.JSON(400, myShared.Response{
//      Message: err.Error(),
//    })
//  }
//  log.Println(data["status"].(bool),strconv.FormatBool(data["status"].(bool)))
//  mySharedRest.PostItem(map[string]interface{}{
//    //"data": &list,
//    "path": "/sessions_students",
//    "query": map[string]string{
//      "session_id": c.Param("id"),
//      "student_id": c.Param("sid"),
//      "tutor_id": strconv.FormatFloat(myJwt.CurrentAuth.Id, 'f', 0, 64),
//      "feedback": data["review"].(string),
//      "rating_cognition": strconv.FormatFloat(data["cognition"].(float64), 'f', 0, 64),
//      "rating_creativity": strconv.FormatFloat(data["creativity"].(float64), 'f', 0, 64),
//      "rating_interaction": strconv.FormatFloat(data["interaction"].(float64), 'f', 0, 64),
//      "status": strconv.FormatBool(data["status"].(bool)),
//    },
//  })
//  return c.JSON(http.StatusOK, myShared.Response{})
//}
func CreateByClassId(c echo.Context) error {
  //var list []Session
  var itemSessionTutor, itemSessions Session
  loc, _ := time.LoadLocation(myShared.TimeZone)
  timeTz := time.Now().In(loc)
  if _, err := mySharedRest.New().GetItem("/sessions__tutors").
    Send("class_id=" + c.Param("id")).
    Send("created_at=gte." + timeTz.Format("2016-01-01")).
    End(&itemSessions); err != nil {
    //return c.JSON(resp.StatusCode, myShared.CreateResponse(err.Error()))
    if resp, err := mySharedRest.New().PostItem(myShared.PathSessions).
      Send("class_id=" + c.Param("id")).
      End(&itemSessions); err != nil {
      return c.JSON(resp.StatusCode, myShared.CreateResponse(err.Error()))
    }
  }
  if resp, err := mySharedRest.New().PostItem("/sessions_tutors").
    Send("session_id=" + fmt.Sprint(itemSessions.Id)).
    Send("tutor_id=" + fmt.Sprint(myJwt.CurrentAuth.Id)).
    End(&itemSessionTutor); err != nil {
    return c.JSON(resp.StatusCode, myShared.CreateResponse(err.Error()))
  }
  //itemSessionTutor := itemSessions[0]
  itemSessionTutor.Links = ItemLinks(itemSessionTutor)
  itemSessionTutor.Embedded = itemEmbedded(itemSessionTutor)
  return c.JSON(http.StatusOK, itemSessionTutor)
}

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
  params := make(map[string]interface{})
  if val := c.Param("id"); len(val) > 0 {
    params["tutor_id"] = "eq." + val
  }
  var list []StudentAbsence
  rest := mySharedRest.New().GetItems("/sessions__students")

  if resp, err := rest.
  SetQuery(params).
    End(&list); err != nil {
    return c.JSON(resp.StatusCode, myShared.CreateResponse(err.Error()))
  }

  var classGroup []ClassStudentAbsence
  var t ClassStudentAbsence

  for _, v := range list {
    found := false
    for i2, v2 := range classGroup {
      if v.ClassId == v2.Class.Id {
        // Found!
        var t3 myUser.User
        myUser.ItemRest(&mySharedRest.Request{}, "student", fmt.Sprint(v.StudentId), &t3)
        classGroup[i2].Students = append(classGroup[i2].Students, t3)

        found = true
        break
      }
    }
    if !found {
      myClass.ItemRest(map[string]string{
        "embed": "tutor,module,branch",
      }, fmt.Sprint(v.ClassId), &t.Class)
      t.Id = v.Id
      t.Students = nil

      var t3 myUser.User
      myUser.ItemRest(&mySharedRest.Request{}, "student", fmt.Sprint(v.StudentId), &t3)
      t.Students = append(t.Students, t3)
      classGroup = append(classGroup, t)
    }
    //list[i].Session.Links = ItemLinks(v.Session)
    //list[i].Session.Embedded = itemEmbedded(v.Session)
  }

  response := myShared.Hal{
    Links:    myShared.CreateHrefSelf(myUser.PathTutors + "/" + c.Param("id") + myShared.PathSessions),
    Embedded: myShared.CreateEmbeddedItems(classGroup),
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
func CreateByStudentId(c echo.Context) error {
  data := make(map[string]interface{});
  if err := c.Bind(&data); err != nil {
    return c.JSON(400, myShared.CreateResponse(err.Error()))
  }
  var item SessionStudent

  if resp, err := mySharedRest.New().PostItem("/sessions_students").Send(map[string]string{
    "session_id":         c.Param("id"),
    "student_id":         c.Param("sid"),
    "tutor_id":           fmt.Sprint(myJwt.CurrentAuth.Id),
    "feedback":           data["review"].(string),
    "rating_cognition":   fmt.Sprint(data["cognition"]),
    "rating_creativity":  fmt.Sprint(data["creativity"]),
    "rating_interaction": fmt.Sprint(data["interaction"]),
    "status":             fmt.Sprint(data["status"].(bool)),
  }).End(&item); err != nil {
    return c.JSON(resp.StatusCode, myShared.CreateResponse(err.Error()))
  }
  return c.JSON(http.StatusOK, myShared.CreateResponse(""))
}

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

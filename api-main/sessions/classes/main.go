package classes

import (
  myShared "../../shared"
  mySharedRest "../../shared/rest"
  myUser "../../users"
  myClass "../../classes"
  myJwt "../../shared/jwt"
  "github.com/labstack/echo"
  "net/http"
  "fmt"
  "time"
  //"github.com/davecgh/go-spew/spew"
)

func List(c echo.Context) error {
  var list []Session
  rest := mySharedRest.New().GetItems(myShared.PathSessions)
  if resp, err := rest.EndStruct(&list); err != nil {
    return c.JSON(resp.StatusCode, myShared.CreateResponse(err.Error()))
  }

  for i, v := range list {
    list[i].Links = ItemLinks(v)
    //list[i].Embedded = itemEmbedded(v)
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
  var list []_SessionsStudentsStatusNull
  rest := mySharedRest.New().GetItems("/_sessions_students_status_null")

  if resp, err := rest.
  SetQuery(params).
    EndStruct(&list); err != nil {
    return c.JSON(resp.StatusCode, myShared.CreateResponse(err.Error()))
  }

  //var classGroup []interface{}
  classGroup := make(map[string][]_SessionsStudentsStatusNull)
  for i, v := range list {
    list[i].Session.Links = ItemStudentsStatusNullLinks(v)

    //t = append(t, list[i])
    strClassId := fmt.Sprint(v.ClassId)
    classGroup[strClassId] = append(classGroup[strClassId], list[i])
    //  found := false
    //  for i2, v2 := range classGroup {
    //    if v.ClassId == v2.Class.Id {
    //      // Found!
    //      var t3 myUser.User
    //      myUser.ItemRest(&mySharedRest.Request{}, "student", fmt.Sprint(v.StudentId), &t3)
    //      classGroup[i2].Students = append(classGroup[i2].Students, t3)
    //
    //      found = true
    //      break
    //    }
    //  }
    //  if !found {
    //    myClass.ItemRest(map[string]string{
    //      "embed": "tutor,module,branch",
    //    }, fmt.Sprint(v.ClassId), &t.Class)
    //    t.Id = v.Id
    //    t.Students = nil
    //
    //    var t3 myUser.User
    //    myUser.ItemRest(&mySharedRest.Request{}, "student", fmt.Sprint(v.StudentId), &t3)
    //    t.Students = append(t.Students, t3)
    //    classGroup = append(classGroup, t)
    //  }
  }

  var listGroups []interface{}
  for k, v := range classGroup {

    t := make(map[string]interface{})
    t["href"] = myClass.Path + "/" + k
    t2 := make(map[string]interface{})
    t2["class"] = t

    t4 := make(map[string]interface{})
    t4["items"] = v

    t3 := make(map[string]interface{})
    t3["_links"] = t2
    t3["_embedded"] = t4
    listGroups = append(listGroups, t3)
  }

  response := myShared.Hal{
    Links:    myShared.CreateHrefSelf(myUser.PathTutors + "/" + c.Param("id") + myShared.PathSessions),
    Embedded: myShared.CreateEmbeddedItems(listGroups),
    Count:    uint64(len(listGroups)),
    Total:    uint64(len(listGroups)),
  }
  return c.JSON(http.StatusOK, response)
}

func ListByClassId(c echo.Context) error {

  params := make(map[string]interface{})
  if val := c.Param("id"); len(val) > 0 {
    params["class_id"] = "eq." + val
  }
  if val := c.QueryParam("sort"); len(val) > 0 {
    params["order"] = val
  }

  if val := c.QueryParam("created_at"); len(val) > 0 {
    params["created_at"] = val
  }

  var list []_SessionTutor
  rest := mySharedRest.New().GetItems("/_sessions_tutors")

  if resp, err := rest.
  SetQuery(params).
    EndStruct(&list); err != nil {
    return c.JSON(resp.StatusCode, myShared.CreateResponse(err.Error()))
  }

  for i, v := range list {
    list[i].Links = ItemTutorLinks(v)
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
    EndStruct(&item); err != nil {
    return c.JSON(resp.StatusCode, myShared.CreateResponse(err.Error()))
  }
  item.Links = ItemLinks(item)
  //item.Embedded = itemEmbedded(item)

  return c.JSON(http.StatusOK, item)
}
func Delete(c echo.Context) error {
  if resp, err := mySharedRest.New().DeleteItem("/sessions_tutors").
    SetQuery("session_id=eq." + c.Param("id")).
    SetQuery("tutor_id=eq." + fmt.Sprint(myJwt.CurrentAuth.Id)).
    End(); err != nil {
    return c.JSON(resp.StatusCode, myShared.CreateResponse(err.Error()))
  }
  if _, err := mySharedRest.New().DeleteItem(myShared.PathSessions).
    SetQuery("id=eq." + c.Param("id")).
    End(); err != nil {
    //return c.JSON(resp.StatusCode, myShared.CreateResponse(err.Error()))
  }
  return c.JSON(http.StatusOK, myShared.CreateResponse(""))
}

func DeleteByStudentId(c echo.Context) error {
  if _, err := mySharedRest.New().DeleteItem("/sessions_students").
    SetQuery("session_id=eq." + c.Param("id")).
    SetQuery("student_id=eq." + c.Param("sid")).
    SetQuery("tutor_id=eq." + fmt.Sprint(myJwt.CurrentAuth.Id)).
    End(); err != nil {
    return c.JSON(400, myShared.CreateResponse(err.Error()))
  }
  return c.JSON(http.StatusOK, myShared.CreateResponse(""))
}
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
    "status":             fmt.Sprint(data["status"]),
  }).EndStruct(&item); err != nil {
    return c.JSON(resp.StatusCode, myShared.CreateResponse(err.Error()))
  }
  return c.JSON(http.StatusOK, item)
}

func CreateByClassId(c echo.Context) error {
  var itemSessionTutor, itemSessions _SessionTutor

  _, err := mySharedRest.New().GetItem("/_sessions_tutors").
    SetQuery("class_id=" + "eq." + c.Param("id")).
    SetQuery("tutor_id=" + "eq." + fmt.Sprint(myJwt.CurrentAuth.Id)).
    SetQuery("created_at=gte." + time.Now().Format("2006-01-02")).
    EndStruct(&itemSessions)
  if err != nil {
    if _, err := mySharedRest.New().GetItem(myShared.PathSessions).
      SetQuery("class_id=" + "eq." + c.Param("id")).
      SetQuery("created_at=gte." + time.Now().Format("2006-01-02")).
      EndStruct(&itemSessions); err != nil {

      if resp, err := mySharedRest.New().PostItem(myShared.PathSessions).
        Send("class_id=" + c.Param("id")).
        EndStruct(&itemSessions); err != nil {
        return c.JSON(resp.StatusCode, myShared.CreateResponse(err.Error()))
      }
    }
    if resp, err := mySharedRest.New().PostItem("/sessions_tutors").
      Send("session_id=" + fmt.Sprint(itemSessions.Id)).
      Send("tutor_id=" + fmt.Sprint(myJwt.CurrentAuth.Id)).
      EndStruct(&itemSessionTutor); err != nil {
      return c.JSON(resp.StatusCode, myShared.CreateResponse(err.Error()))
    }
  }

  return c.JSON(http.StatusOK, itemSessions)
}

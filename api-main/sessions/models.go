package sessions

import (
  myShared "../shared"
  myUsers "../users"
  "time"
  "strconv"
)

type (
  SessionStudentRating struct {
    Interaction uint8 `json:"interaction"`
    Cognition   uint8 `json:"cognition"`
    Creativity  uint8 `json:"creativity"`
  }
  SessionStudent struct {
    myShared.Hal
    Id        uint64               `json:"id"`
    CreatedAt time.Time            `json:"created_at"`
    Status    bool                 `json:"status"`
    Feedback  string               `json:"feedback"`
    Rating    SessionStudentRating `json:"rating"`
    StudentId uint64               `json:"student_id" `
  }

  SessionLinks struct {
    myShared.LinksSelf
    Class            myShared.Href   `json:"class"`
    Tutors           myShared.Href   `json:"tutors"`
    Students         []myShared.Href `json:"students,omitempty"`
    StudentsSessions []myShared.Href `json:"students_sessions,omitempty"`
  }
  SessionEmbedded struct {
    Class    myShared.Class_ `json:"class"`
    Students []myShared.User `json:"students,omitempty"`
  }
)

func itemLinksStudentsSessions(id uint64) []myShared.Href {
  var list []myShared.ClassStudents
  myShared.GetItems(map[string]interface{}{
    "data": &list,
    "path": "/sessions_students",
    "query": map[string]string{
      "session_id": "eq." + strconv.FormatUint(id, 10),
      "select":     "id",
    },
  })
  var data []myShared.Href
  for _, v := range list {
    data = append(data, myShared.CreateHref(myShared.PathSessions+"/"+strconv.FormatUint(id, 10)+myShared.PathStudents+"/"+strconv.FormatUint(v.Id, 10)))
  }
  return data
}
func itemLinksStudents(id uint64) []myShared.Href {
  var list []myShared.ClassStudents
  myShared.GetItems(map[string]interface{}{
    "data": &list,
    "path": "/class_students",
    "query": map[string]string{
      "class_id": "eq." + strconv.FormatUint(id, 10),
      "select":   "student_id",
    },
  })
  var data []myShared.Href
  for _, v := range list {
    data = append(data, myShared.CreateHref(myShared.PathStudents+"/"+strconv.FormatUint(v.StudentId, 10)))
  }
  return data
}
func ItemLinks(v myShared.Session) SessionLinks {
  return SessionLinks{
    LinksSelf:        myShared.LinksSelf{Self: myShared.CreateHref(myShared.PathSessions + "/" + strconv.FormatUint(v.Id, 10))},
    Class:            myShared.Href{Href: myShared.PathClasses + "/" + strconv.FormatUint(v.ClassId, 10)},
    Tutors:           myShared.Href{Href: myShared.PathTutors + "/" + strconv.FormatUint(v.TutorId, 10)},
    Students:         itemLinksStudents(v.ClassId),
    StudentsSessions: itemLinksStudentsSessions(v.Id),
  }
}

func itemEmbeddedClassEmbeddedStudents(id uint64) []myShared.UsersProfile {
  var list []myShared.ClassStudents
  myShared.GetItems(map[string]interface{}{
    "data": &list,
    "path": "/class_students",
    "query": map[string]string{
      "class_id": "eq." + strconv.FormatUint(id, 10),
      "select":   "student_id",
    },
  })

  var data []myShared.UsersProfile
  for _, v := range list {
    var item myShared.UsersProfile
    //params := map[string]string{
    //  "sessions_students.student_id": "eq." + strconv.FormatUint(v.StudentId, 10),
    //  "select": "*,sessions_students(*)",
    //}
    //myShared.GetItem(map[string]interface{}{
    //  "data":  &item,
    //  "path":  "/sessions",
    //  "query": params,
    //})
    if true {
      params := map[string]string{
        "user_id": "eq." + strconv.FormatUint(v.StudentId, 10),
      }
      var item myShared.UsersProfile
      myShared.GetItem(map[string]interface{}{
        "data":  &item,
        "path":  "/users_profile",
        "query": params,
      })
      //item.Id = v.StudentId
    }

    data = append(data, item)
  }
  //params["user_id"] = "in.(" + strings.Join(in, ",") + ")"
  //params["select"] = "*,users(id)"
  //myShared.GetItems(map[string]interface{}{
  //  "data":  &data,
  //  "path":  "/users_profile",
  //  "query": params,
  //})
  //log.Println(data,strings.Join(in, ","))
  return data
}
func itemEmbeddedClassEmbedded(data myShared.Class_) myShared.ClassEmbedded {
  var embedded myShared.ClassEmbedded
  lastSession := myShared.ClassItemEmbeddedLastSessions(data.Id)
  //  links.LastSession = &lastSession
  if lastSession.Id > 0 {
    embedded.LastSession = &lastSession
  }
  module := myShared.ClassItemEmbeddedModule(data.ModuleId)
  embedded.Module = &module
  branch := myShared.ClassItemEmbeddedBranch(data.BranchId)
  embedded.Branch = &branch
  tutor := myShared.ClassItemEmbeddedTutor(data.TutorId)
  embedded.Tutor = &tutor

  embedded.Students = myShared.ClassItemEmbeddedStudents(data.Id)

  return embedded
}
func itemEmbeddedClass(id uint64) myShared.Class_ {
  var item myShared.Class_
  myShared.GetItem(map[string]interface{}{
    "data": &item,
    "path": "/classes_ts",
    "query": map[string]string{
      "id": "eq." + strconv.FormatUint(id, 10),
    },
  })
  item.Links = myShared.ClassItemLinks(item)
  item.Embedded = itemEmbeddedClassEmbedded(item)
  return item
}
func itemEmbeddedStudents(id uint64) []myShared.User {
  var list []myShared.ClassStudents
  myShared.GetItems(map[string]interface{}{
    "data": &list,
    "path": "/class_students",
    "query": map[string]string{
      "class_id": "eq." + strconv.FormatUint(id, 10),
      "select":   "student_id",
    },
  })
  var data []myShared.User
  for _, v := range list {

    var item myShared.User
    myShared.GetItem(map[string]interface{}{
      "data": &item,
      "path": myShared.PathUsers,
      "query": map[string]string{
        "id":   "eq." + strconv.FormatUint(v.StudentId, 10),
        "role": "eq.student",
      },
    })
    item.Links = myUsers.ItemLinks(item, "student")
    data = append(data, item)
  }
  return data
}
func itemEmbedded(v myShared.Session) SessionEmbedded {
  return SessionEmbedded{
    Class: itemEmbeddedClass(v.ClassId),
    //Students: itemEmbeddedStudents(v.ClassId),
  }
}

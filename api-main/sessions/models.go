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
    // SessionsTutorsId string `json:"-" db:"sessions_tutors_id"`
    StudentId uint64 `json:"-" db:"student_id"`
  }

  SessionLinks struct {
    myShared.LinksSelf
    Class            myShared.Href   `json:"class"`
    Tutors myShared.Href `json:"tutors"`
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
    LinksSelf: myShared.LinksSelf{Self: myShared.CreateHref(myShared.PathSessions + "/" + strconv.FormatUint(v.Id, 10))},
    Class:            myShared.Href{Href: myShared.PathClasses + "/" + strconv.FormatUint(v.ClassId, 10)},
    Tutors: myShared.Href{Href: myShared.PathTutors + "/" + strconv.FormatUint(v.TutorId, 10)},
    Students:         itemLinksStudents(v.ClassId),
    StudentsSessions: itemLinksStudentsSessions(v.Id),
  }
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
  item.Embedded = myShared.ClassItemEmbedded(item)
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

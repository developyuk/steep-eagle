package sessions

import (
  myShared "../shared"
  myUser "../users"
  myClass "../classes"
  "time"
  "fmt"
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

  StudentAbsence struct {
    Session
    StudentId uint64 `json:"student_id" `
  }
  ClassStudentAbsence struct {
    Class     myClass.Class_ `json:"class"`
    Id uint64         `json:"id"`
    Students  []myUser.User  `json:"students,omitempty"`
  }

  SessionLinks struct {
    myShared.LinksSelf
    Class  myShared.Href `json:"class"`
    Tutors myShared.Href `json:"tutors"`
  }
  //SessionEmbedded struct {
  //  //Class    myShared.Class_ `json:"class"`
  //  Tutor myUser.User `json:"tutor,omitempty"`
  //}

  Session struct {
    myShared.Hal
    Id        uint64    `json:"id"`
    CreatedAt time.Time `json:"created_at"`
    ClassId   uint64    `json:"class_id"`
    TutorId   uint64    `json:"tutor_id"`
    //Users     myUser.User `json:"users"`
  }
)

func ItemLinks(v Session) SessionLinks {
  return SessionLinks{
    LinksSelf: myShared.CreateHrefSelf(myShared.PathSessions + "/" + fmt.Sprint(v.Id)),
    Class:     myShared.Href{Href: myClass.Path + "/" + fmt.Sprint(v.ClassId)},
    Tutors:    myShared.Href{Href: myUser.PathTutors + "/" + fmt.Sprint(v.TutorId)},
  }
}

//func itemEmbeddedTutor(id uint64) myUser.User {
//  var item myUser.User
//  myUser.ItemRest(&mySharedRest.Request{}, "tutor", fmt.Sprint(id), &item)
//  return item
//}
//func itemEmbedded(v Session) SessionEmbedded {
//  return SessionEmbedded{
//    Tutor: itemEmbeddedTutor(v.TutorId),
//  }
//}

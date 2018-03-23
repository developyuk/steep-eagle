package classes

import (
  myShared "../../shared"
  myUser "../../users"
  myClass "../../classes"
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

  Session struct {
    myShared.Hal
    Id        uint64    `json:"id"`
    CreatedAt time.Time `json:"created_at"`
    ClassId   uint64    `json:"class_id"`
    //Users     myUser.User `json:"users"`
  }
  SessionLinks struct {
    myShared.LinksSelf
    Class *myShared.Href `json:"class"`
  }
  //SessionEmbedded struct {
  //  //Class    myShared.Class_ `json:"class"`
  //  Tutor myUser.User `json:"tutor,omitempty"`
  //}

  _SessionTutor struct {
    Session
    TutorId uint64 `json:"tutor_id"`
  }
  _SessionTutorLinks struct {
    SessionLinks
    Tutor *myShared.Href `json:"tutor"`
  }

  _SessionsStudentsStatusNull struct {
    Session
    StudentId uint64 `json:"student_id" `
  }
  _SessionsStudentsStatusNullLinks struct {
    SessionLinks
    Student *myShared.Href `json:"student"`
  }
)

func ItemStudentsStatusNullLinks(v _SessionsStudentsStatusNull) _SessionsStudentsStatusNullLinks {
  return _SessionsStudentsStatusNullLinks{
    SessionLinks: SessionLinks{
      LinksSelf: myShared.CreateHrefSelf(myShared.PathSessions + "/" + fmt.Sprint(v.Id)),
      Class:     myShared.CreateHref(myClass.Path + "/" + fmt.Sprint(v.ClassId)),
    },
    Student: myShared.CreateHref(myUser.PathStudents + "/" + fmt.Sprint(v.StudentId)),
  }
}
func ItemTutorLinks(v _SessionTutor) _SessionTutorLinks {
  return _SessionTutorLinks{
    SessionLinks: SessionLinks{
      LinksSelf: myShared.CreateHrefSelf(myShared.PathSessions + "/" + fmt.Sprint(v.Id)),
      Class:     myShared.CreateHref(myClass.Path + "/" + fmt.Sprint(v.ClassId)),
    },
    Tutor: myShared.CreateHref(myUser.PathUsers + "/" + fmt.Sprint(v.TutorId)),
  }
}
func ItemLinks(v Session) SessionLinks {
  return SessionLinks{
    LinksSelf: myShared.CreateHrefSelf(myShared.PathSessions + "/" + fmt.Sprint(v.Id)),
    Class:     myShared.CreateHref(myClass.Path + "/" + fmt.Sprint(v.ClassId)),
  }
}

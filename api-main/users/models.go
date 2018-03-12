package users

import (
  myShared "../shared"
  myRest "../shared/rest"
  "fmt"
)

type (
  TutorLinks struct {
    myShared.LinksSelf
    Sessions []myShared.Href `json:"sessions,omitempty"`
  }

  StudentLinks struct {
    myShared.LinksSelf
    Classes []myShared.Href `json:"classes,omitempty"`
  }
)
const (
  PathUsers    string = "/users"
  PathStudents string = "/students"
  PathTutors   string = "/tutors"
)

func itemLinks(v myShared.User, role string) interface{} {
  path := getPath(role)
  //log.Println(role)
  if v.Role == "student" {
    var studentLinks StudentLinks
    studentLinks.Self = myShared.CreateHref(PathStudents + "/" + fmt.Sprint(v.Id))
    return studentLinks
  }
  if v.Role == "tutor" {
    var tutorLinks TutorLinks
    tutorLinks.Self = myShared.CreateHref(PathTutors + "/" + fmt.Sprint(v.Id))
    return tutorLinks
  }

  return myShared.LinksSelf{Self: myShared.CreateHref(path + "/" + fmt.Sprint(v.Id))}
}

func itemByUsername(param *UserLoginRequest) (myShared.User, error) {

  var item myShared.User

  if _, err := myRest.New().GetItem("/users_full").
    SetQuery(myShared.RequestRest{Select: "id,email,role,name,photo"}).
    SetQuery("username=eq." + param.Username).
    End(&item); err != nil {
    return item, err
  }

  return item, nil

}

//func itemByEmailPass(param *UserLoginRequest) (myShared.User, error) {
//
//  var item myShared.User
//  _, err := myShared.PostItem(map[string]interface{}{
//    "data": &item,
//    "path": "/rpc/user_by_email_pass",
//    "query": map[string]string{
//      "_email": param.Email,
//      "_pass":  param.Pwd,
//    },
//  })
//
//  if err != nil {
//    return item, err
//  }
//
//  return item, err
//
//}

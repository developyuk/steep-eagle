package users

import (
  myShared "../shared"
  "strconv"
  "log"
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

func itemLinksSessions(id uint64) []myShared.Href {

  var list []myShared.Session
  myShared.GetItems(map[string]interface{}{
    "data": &list,
    "path": "/sessions",
    "query": map[string]string{
      "tutor_id": "eq." + strconv.FormatUint(id, 10),
      "select":   "id",
    },
  })

  var data []myShared.Href
  for _, v := range list {
    data = append(data, myShared.CreateHref(myShared.PathSessions+"/"+strconv.FormatUint(v.Id, 10)))
  }

  return data
}

func itemLinksClass(id uint64) []myShared.Href {
  var list []myShared.ClassStudents
  myShared.GetItems(map[string]interface{}{
    "data": &list,
    "path": "/class_students",
    "query": map[string]string{
      "student_id": "eq." + strconv.FormatUint(id, 10),
      "select":     "class_id",
    },
  })

  var data []myShared.Href
  for _, v := range list {
    data = append(data, myShared.CreateHref(myShared.PathClasses+"/"+strconv.FormatUint(v.ClassId, 10)))
  }

  return data
}

func ItemLinks(v myShared.User, role string) interface{} {
  path := getPath(role)
  //log.Println(role)
  if role == "student" {
    var studentLinks StudentLinks
    studentLinks.Self = myShared.CreateHref(path + "/" + strconv.FormatUint(v.Id, 10))
    studentLinks.Classes = itemLinksClass(v.Id)
    return studentLinks
  }
  if role == "tutor" {
    var tutorLinks TutorLinks
    tutorLinks.Self = myShared.CreateHref(path + "/" + strconv.FormatUint(v.Id, 10))
    tutorLinks.Sessions = itemLinksSessions(v.Id)
    return tutorLinks
  }

  return myShared.LinksSelf{Self: myShared.CreateHref(path + "/" + strconv.FormatUint(v.Id, 10))}
}

func itemByUsername(param *UserLoginRequest) (myShared.User, error) {

  var item myShared.User
  _, err := myShared.GetItem(map[string]interface{}{
    "data": &item,
    "path": myShared.PathUsers,
    "query": map[string]string{
      "username": "eq." + param.Username,
      "select": "id,email,role,users_profile(name,photo)",
    },
  })
  log.Println(item,param)

  if err != nil {
    return item, err
  }

  return item, err

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

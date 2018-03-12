package shared

import (
  //"strconv"
  "time"
  //"strings"
  //myRest "./rest"
)

const (
  PathProgramsTypes string = "/programs/types"
  PathClasses       string = "/classes"
  PathSessions      string = "/sessions"
)

const (
  TimeZone = "Asia/Jakarta"
)

type (
  ClassStudents struct {
    Id        uint64 `json:"id"`
    ClassId   uint64 `json:"class_id"`
    StudentId uint64 `json:"student_id"`
  }
  Class_ struct {
    Hal
    Id         uint64    `json:"id"`
    Name       string    `json:"name"`
    Image      string    `json:"image"`
    Day        string    `json:"day"`
    StartAt    string    `json:"start_at"`
    FinishAt   string    `json:"finish_at"`
    StartAtTs  time.Time `json:"start_at_ts"`
    FinishAtTs time.Time `json:"finish_at_ts"`
    //Time     string `json:"time"`
    //Ts       string `json:"ts"`
    ModuleId uint64 `json:"module_id"`
    BranchId uint64 `json:"branch_id"`
    TutorId  uint64 `json:"tutor_id"`
    //Module   myModules.Module  `json:"module"`
    //Branch   myBranches.Branch `json:"branch"`
  }
  ClassLinks struct {
    LinksSelf
    Module   *Href  `json:"module,omitempty"`
    Branch   *Href  `json:"branch,omitempty"`
    Tutor    *Href  `json:"tutor,omitempty"`
    Students []Href `json:"students,omitempty"`
    Sessions []Href `json:"sessions,omitempty"`
  }
)

type (
  Session struct {
    Hal
    Id        uint64    `json:"id"`
    CreatedAt time.Time `json:"created_at"`
    ClassId   uint64    `json:"class_id"`
    TutorId   uint64    `json:"tutor_id"`
    Users     User      `json:"users"`
  }
)

type (
  User struct {
    Hal
    Id       uint64    `json:"id"`
    Username string    `json:"username,omitempty"`
    Email    string    `json:"email,omitempty"`
    Pass     string    `json:"pass,omitempty"`
    Role     string    `json:"role,omitempty"`
    Name     string    `json:"name,omitempty"`
    Dob      time.Time `json:"dob,omitempty"`
    Photo    string    `json:"photo,omitempty"`
    UserId   uint64    `json:"user_id,omitempty"`
    //UsersProfile []UsersProfile `json:"users_profile"`
  }
)

//func ClassLinksSessions(id uint64) []Href {
//  var list []Session
//
//  myRest.GetItems(map[string]interface{}{
//    "data": &list,
//    "path": "/sessions",
//    "query": map[string]string{
//      "class_id": "eq." + strconv.FormatUint(id, 10),
//      "order":    "created_at.asc",
//      "select":   "id",
//    },
//  })
//
//  var data []Href
//  for _, v := range list {
//    data = append(data, CreateHref(PathSessions+"/"+strconv.FormatUint(v.Id, 10)))
//  }
//
//  return data
//}
//func ClassLinksStudents(id uint64) []Href {
//  var list []ClassStudents
//  myRest.GetItems(map[string]interface{}{
//    "data": &list,
//    "path": "/class_students",
//    "query": map[string]string{
//      "class_id": "eq." + strconv.FormatUint(id, 10),
//      "select":   "student_id",
//    },
//  })
//
//  var data []Href
//  for _, v := range list {
//    data = append(data, CreateHref(PathStudents+"/"+strconv.FormatUint(v.StudentId, 10)))
//  }
//
//  return data
//}
//
//func ClassItemEmbeddedModule(id uint64) Module {
//  var module Module
//
//  _, err := myRest.GetItem(map[string]interface{}{
//    "data": &module,
//    "path": PathModules,
//    "query": map[string]string{
//      "id":    "eq." + strconv.FormatUint(id, 10),
//      "limit": "1",
//      //"select":   "id,created_at",
//    },
//  })
//  if err != nil {
//    return Module{}
//  }
//  //module.Links = myModules.ItemLinks(module)
//
//  return module
//}
//func ClassItemEmbeddedTutor(id uint64) User {
//  var tutor User
//
//  _, err := myRest.GetItem(map[string]interface{}{
//    "data": &tutor,
//    "path": PathUsers,
//    "query": map[string]string{
//      "id":     "eq." + strconv.FormatUint(id, 10),
//      "limit":  "1",
//      "select": "*,users_profile(*)",
//    },
//  })
//  if err != nil {
//    return User{}
//  }
//  //tutor.Links = myBranches.ItemLinks(tutor)
//
//  return tutor
//}
//
//func ClassItemEmbeddedLastSessions(id uint64) Session {
//  var session Session
//
//  _, err := myRest.GetItem(map[string]interface{}{
//    "data": &session,
//    "path": PathSessions,
//    "query": map[string]string{
//      "class_id": "eq." + strconv.FormatUint(id, 10),
//      //"tutor_id": "eq." + strconv.FormatFloat(CurrentAuth.Id, 'f', 0, 64),
//      "order":  "created_at.desc",
//      "limit":  "1",
//      "select": "*,users(*)",
//    },
//  })
//  if err != nil {
//    return Session{}
//  }
//  //session.Links = mySessions.ItemLinks(session)
//
//  return session
//}
//func ClassItemEmbeddedStudents(id uint64) []UsersProfile {
//  var list []ClassStudents
//  myRest.GetItems(map[string]interface{}{
//    "data": &list,
//    "path": "/class_students",
//    "query": map[string]string{
//      "class_id": "eq." + strconv.FormatUint(id, 10),
//      "select":   "student_id",
//    },
//  })
//
//  var data []UsersProfile
//  var in []string
//  params := make(map[string]string)
//  for _, v := range list {
//    in = append(in, strconv.FormatUint(v.StudentId, 10))
//    //  params["user_id"] = "eq." + strconv.FormatUint(v.StudentId, 10)
//    //  //params["role"] = "eq.student"
//    //  var item User
//    //  GetItem(map[string]interface{}{
//    //    "data":  &item,
//    //    "path":  "/users_profile",
//    //    "query": params,
//    //  })
//    //  item.Id = v.StudentId
//    //  //log.Println(strconv.FormatUint(v.StudentId, 10),item)
//    //
//    //  //item.Links = ItemLinks(item, role)
//    //  data = append(data, item)
//  }
//  params["user_id"] = "in.(" + strings.Join(in, ",") + ")"
//  params["select"] = "*,users(id)"
//  myRest.GetItems(map[string]interface{}{
//    "data":  &data,
//    "path":  "/users_profile",
//    "query": params,
//  })
//  //log.Println(data,strings.Join(in, ","))
//  return data
//}

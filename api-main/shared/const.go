package shared

import (
  "os"
  "strconv"
  "time"
  "strings"
)

const (
  PathPrograms      string = "/programs"
  PathProgramsTypes string = "/programs/types"
  PathModules       string = "/modules"
  PathClasses       string = "/classes"
  PathUsers         string = "/users"
  PathStudents      string = "/students"
  PathTutors        string = "/tutors"
  PathSessions      string = "/sessions"
  PathBranches      string = "/branches"
)

var JwtKey = os.Getenv("JWT_KEY")

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
    Id         uint64 `json:"id"`
    Name       string `json:"name"`
    Image      string `json:"image"`
    Day        string `json:"day"`
    StartAt    string `json:"start_at"`
    FinishAt   string `json:"finish_at"`
    StartAtTs  string `json:"start_at_ts"`
    FinishAtTs string `json:"finish_at_ts"`
    //Time     string `json:"time"`
    //Ts       string `json:"ts"`
    ModuleId uint64 `json:"module_id"`
    BranchId uint64 `json:"branch_id"`
    //Module   myModules.Module  `json:"module"`
    //Branch   myBranches.Branch `json:"branch"`
  }
  ClassLinks struct {
    LinksSelf
    Module   *Href  `json:"module,omitempty"`
    Branch   *Href  `json:"branch,omitempty"`
    Students []Href `json:"students,omitempty"`
    Sessions []Href `json:"sessions,omitempty"`
  }
  ClassEmbedded struct {
    LastSession *Session `json:"last_session,omitempty"`
    Module      *Module  `json:"module,omitempty"`
    Branch      *Branch  `json:"branch,omitempty"`
    Students    []User   `json:"students,omitempty"`
  }
)

type (
  Branch struct {
    Hal
    Id      uint64 `json:"id"`
    Name    string `json:"name"`
    Image   string `json:"image"`
    Address string `json:"address"`
  }
)

type (
  Module struct {
    Hal
    Id           uint64 `json:"id"`
    Name         string `json:"name"`
    Image        string `json:"image"`
    SessionTotal uint   `json:"session_total" db:"session_total"`
  }
)

type (
  Session struct {
    Hal
    Id        uint64    `json:"id"`
    CreatedAt time.Time `json:"created_at"`
    ClassId   uint64    `json:"class_id"`
    TutorId   uint64    `json:"tutor_id"`
  }
)

type (
  User struct {
    Hal
    Id        uint64      `json:"id"`
    Name      string      `json:"name"`
    FirstName string      `json:"first_name"`
    LastName  string      `json:"last_name"`
    Email     string      `json:"email"`
    Dob       time.Time   `json:"dob"`
    Photo     string      `json:"photo"`
    Role      string      `json:"role"`
    Users     interface{} `json:"users"`
  }
)

func ClassLinksSessions(id uint64) []Href {
  var list []Session

  GetItems(map[string]interface{}{
    "data": &list,
    "path": "/sessions",
    "query": map[string]string{
      "class_id": "eq." + strconv.FormatUint(id, 10),
      "order":    "created_at.asc",
      "select":   "id",
    },
  })

  var data []Href
  for _, v := range list {
    data = append(data, CreateHref(PathSessions+"/"+strconv.FormatUint(v.Id, 10)))
  }

  return data
}
func ClassLinksStudents(id uint64) []Href {
  var list []ClassStudents
  GetItems(map[string]interface{}{
    "data": &list,
    "path": "/class_students",
    "query": map[string]string{
      "class_id": "eq." + strconv.FormatUint(id, 10),
      "select":   "student_id",
    },
  })

  var data []Href
  for _, v := range list {
    data = append(data, CreateHref(PathStudents+"/"+strconv.FormatUint(v.StudentId, 10)))
  }

  return data
}
func ClassItemLinks(data Class_) ClassLinks {
  var links ClassLinks
  links.Self = CreateHref(PathClasses + "/" + strconv.FormatUint(data.Id, 10))
  linksModule := CreateHref(PathModules + "/" + strconv.FormatUint(data.ModuleId, 10))
  links.Module = &linksModule
  linksBranch := CreateHref(PathBranches + "/" + strconv.FormatUint(data.BranchId, 10))
  links.Branch = &linksBranch
  links.Students = ClassLinksStudents(data.Id)
  links.Sessions = ClassLinksSessions(data.Id)
  //if len(links.Sessions) > 0 {
  //  lastSession := itemLinksLastSessions(data.Id)
  //  links.LastSession = &lastSession
  //}
  return links

}

func ClassItemEmbeddedModule(id uint64) Module {
  var module Module

  _, err := GetItem(map[string]interface{}{
    "data": &module,
    "path": PathModules,
    "query": map[string]string{
      "id":    "eq." + strconv.FormatUint(id, 10),
      "limit": "1",
      //"select":   "id,created_at",
    },
  })
  if err != nil {
    return Module{}
  }
  //module.Links = myModules.ItemLinks(module)

  return module
}
func ClassItemEmbeddedBranch(id uint64) Branch {
  var branch Branch

  _, err := GetItem(map[string]interface{}{
    "data": &branch,
    "path": PathBranches,
    "query": map[string]string{
      "id":    "eq." + strconv.FormatUint(id, 10),
      "limit": "1",
      //"select":   "id,created_at",
    },
  })
  if err != nil {
    return Branch{}
  }
  //branch.Links = myBranches.ItemLinks(branch)

  return branch
}

func ClassItemEmbeddedLastSessions(id uint64) Session {
  var session Session

  _, err := GetItem(map[string]interface{}{
    "data": &session,
    "path": PathSessions,
    "query": map[string]string{
      "class_id": "eq." + strconv.FormatUint(id, 10),
      "tutor_id": "eq." + strconv.FormatFloat(CurrentAuth.Id, 'f', 0, 64),
      "order":    "created_at.desc",
      "limit":    "1",
      //"select":   "id,created_at",
    },
  })
  if err != nil {
    return Session{}
  }
  //session.Links = mySessions.ItemLinks(session)

  return session
}
func ClassItemEmbeddedStudents(id uint64) []User {
  var list []ClassStudents
  GetItems(map[string]interface{}{
    "data": &list,
    "path": "/class_students",
    "query": map[string]string{
      "class_id": "eq." + strconv.FormatUint(id, 10),
      "select":   "student_id",
    },
  })

  var data []User
  var in []string
  params := make(map[string]string)
  for _, v := range list {
    in = append(in, strconv.FormatUint(v.StudentId, 10))
    //  params["user_id"] = "eq." + strconv.FormatUint(v.StudentId, 10)
    //  //params["role"] = "eq.student"
    //  var item User
    //  GetItem(map[string]interface{}{
    //    "data":  &item,
    //    "path":  "/users_profile",
    //    "query": params,
    //  })
    //  item.Id = v.StudentId
    //  //log.Println(strconv.FormatUint(v.StudentId, 10),item)
    //
    //  //item.Links = ItemLinks(item, role)
    //  data = append(data, item)
  }
  params["user_id"] = "in.(" + strings.Join(in, ",") + ")"
  params["select"] = "*,users(id)"
  GetItems(map[string]interface{}{
    "data":  &data,
    "path":  "/users_profile",
    "query": params,
  })
  return data
}
func ClassItemEmbedded(data Class_) ClassEmbedded {
  var embedded ClassEmbedded
  lastSession := ClassItemEmbeddedLastSessions(data.Id)
  //  links.LastSession = &lastSession
  if lastSession.Id > 0 {
    embedded.LastSession = &lastSession
  }
  module := ClassItemEmbeddedModule(data.ModuleId)
  embedded.Module = &module
  branch := ClassItemEmbeddedBranch(data.BranchId)
  embedded.Branch = &branch

  embedded.Students = ClassItemEmbeddedStudents(data.Id)

  return embedded
}

package shared

import (
  "os"
  "strconv"
  "time"
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
    Id       uint64 `json:"id"`
    Name     string `json:"name"`
    Image    string `json:"image"`
    Day      string `json:"day"`
    Time     string `json:"time"`
    Ts       string `json:"ts"`
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
    Id        uint64    `json:"id"`
    Name      string    `json:"name"`
    FirstName string    `json:"first_name"  db:"first_name"`
    LastName  string    `json:"last_name"  db:"last_name"`
    Email     string    `json:"email"`
    Dob       time.Time `json:"dob"`
    Photo     string    `json:"photo"`
    Role      string    `json:"role"`
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
  var list []User
  GetItems(map[string]interface{}{
    "data": &list,
    "path": "/class_students",
    "query": map[string]string{
      "class_id": "eq." + strconv.FormatUint(id, 10),
      "select":   "id",
    },
  })

  var data []Href
  for _, v := range list {
    data = append(data, CreateHref(PathStudents+"/"+strconv.FormatUint(v.Id, 10)))
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

package classes

import (
  myShared "../shared"
  mySessions "../sessions"
  myUsers "../users"
  "strconv"
)

type (
  Class_ struct {
    myShared.Hal
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
    myShared.LinksSelf
    Module      *myShared.Href  `json:"module,omitempty"`
    Branch      *myShared.Href  `json:"branch,omitempty"`
    Students    []myShared.Href `json:"students,omitempty"`
    LastSession *myShared.Href  `json:"last_session,omitempty"`
    Sessions    []myShared.Href `json:"sessions,omitempty"`
  }
)

func itemLinksLastSessions(id uint64) myShared.Href {
  var session mySessions.Session

  myShared.GetItem(map[string]interface{}{
    "path": "/sessions",
    "query": map[string]string{
      "class_id": "eq." + strconv.FormatUint(id, 10),
      "order":    "created_at.desc",
      "limit":    "1",
      "select":   "id",
    },
  })

  return myShared.CreateHref(myShared.PathSessions + "/" + strconv.FormatUint(session.Id, 10))
}
func itemLinksSessions(id uint64) []myShared.Href {
  var list []mySessions.Session

  myShared.GetItems(map[string]interface{}{
    "data": &list,
    "path": "/sessions",
    "query": map[string]string{
      "class_id": "eq." + strconv.FormatUint(id, 10),
      "order":    "created_at.desc",
      "select":   "id",
    },
  })

  var data []myShared.Href
  for _, v := range list {
    data = append(data, myShared.CreateHref(myShared.PathSessions+"/"+strconv.FormatUint(v.Id, 10)))
  }

  return data
}
func itemLinksStudents(id uint64) []myShared.Href {
  var list []myUsers.User
  myShared.GetItems(map[string]interface{}{
    "data": &list,
    "path": "/class_students",
    "query": map[string]string{
      "class_id": "eq." + strconv.FormatUint(id, 10),
      "select":   "id",
    },
  })

  var data []myShared.Href
  for _, v := range list {
    data = append(data, myShared.CreateHref(myShared.PathStudents+"/"+strconv.FormatUint(v.Id, 10)))
  }

  return data
}

func itemLinks(data Class_) ClassLinks {
  var links ClassLinks
  links.Self = myShared.CreateHref(myShared.PathClasses + "/" + strconv.FormatUint(data.Id, 10))
  linksModule := myShared.CreateHref(myShared.PathModules + "/" + strconv.FormatUint(data.ModuleId, 10))
  links.Module = &linksModule
  linksBranch := myShared.CreateHref(myShared.PathBranches + "/" + strconv.FormatUint(data.BranchId, 10))
  links.Branch = &linksBranch
  links.Students = itemLinksStudents(data.Id)
  links.Sessions = itemLinksSessions(data.Id)
  if len(links.Sessions) > 0 {
    lastSession := itemLinksLastSessions(data.Id)
    links.LastSession = &lastSession
  }
  return links

}

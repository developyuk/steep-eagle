package classes

import (
  myShared "../shared"
  myBranch "../branches"
  myRest "../shared/rest"
  "strconv"
)

type (
  Embedded struct {
    LastSession *myShared.Session       `json:"last_session,omitempty"`
    Module      *myShared.Module        `json:"module,omitempty"`
    Branch      *myBranch.Branch        `json:"branch,omitempty"`
    Tutor       *myShared.User          `json:"tutor,omitempty"`
    Students    []myShared.UsersProfile `json:"students,omitempty"`
  }
)

func ItemEmbeddedBranch(id uint64) myBranch.Branch {
  var branch myBranch.Branch

  _, err := myRest.GetItem(map[string]interface{}{
    "data": &branch,
    "path": myBranch.Path,
    "query": map[string]string{
      "id":    "eq." + strconv.FormatUint(id, 10),
      "limit": "1",
      //"select":   "id,created_at",
    },
  })
  if err != nil {
    return myBranch.Branch{}
  }
  //branch.Links = myBranches.ItemLinks(branch)

  return branch
}

func ItemEmbedded(data myShared.Class_) Embedded {
  var embedded Embedded
  lastSession := myShared.ClassItemEmbeddedLastSessions(data.Id)
  //  links.LastSession = &lastSession
  if lastSession.Id > 0 {
    embedded.LastSession = &lastSession
  }
  module := myShared.ClassItemEmbeddedModule(data.ModuleId)
  embedded.Module = &module
  branch := ItemEmbeddedBranch(data.BranchId)
  embedded.Branch = &branch
  tutor := myShared.ClassItemEmbeddedTutor(data.TutorId)
  embedded.Tutor = &tutor

  embedded.Students = myShared.ClassItemEmbeddedStudents(data.Id)

  return embedded
}

func ItemLinks(data myShared.Class_) myShared.ClassLinks {
  var links myShared.ClassLinks
  links.Self = myShared.CreateHref(myShared.PathClasses + "/" + strconv.FormatUint(data.Id, 10))
  linksModule := myShared.CreateHref(myShared.PathModules + "/" + strconv.FormatUint(data.ModuleId, 10))
  links.Module = &linksModule
  linksBranch := myShared.CreateHref(myBranch.Path + "/" + strconv.FormatUint(data.BranchId, 10))
  links.Branch = &linksBranch
  linksTutor := myShared.CreateHref(myShared.PathTutors + "/" + strconv.FormatUint(data.TutorId, 10))
  links.Tutor = &linksTutor
  links.Students = myShared.ClassLinksStudents(data.Id)
  links.Sessions = myShared.ClassLinksSessions(data.Id)
  //if len(links.Sessions) > 0 {
  //  lastSession := itemLinksLastSessions(data.Id)
  //  links.LastSession = &lastSession
  //}
  return links
}

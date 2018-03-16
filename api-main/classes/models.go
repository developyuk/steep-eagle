package classes

import (
  myShared "../shared"
  myModule "../modules"
  myBranch "../branches"
  myUser "../users"
  mySharedRest "../shared/rest"
  "fmt"
)

type (
  embedded struct {
    //LastSession *myShared.Session `json:"last_session,omitempty"`
    Module      myModule.Module  `json:"module,omitempty"`
    Branch      myBranch.Branch  `json:"branch,omitempty"`
    Tutor       myUser.User    `json:"tutor,omitempty"`
    //Students    []myUser.User   `json:"students,omitempty"`
  }
)

func ItemEmbeddedBranch(id uint64) myBranch.Branch {
  var branch myBranch.Branch

  if _, err := mySharedRest.New().GetItem(myBranch.Path).
    SetQuery("id=eq." + fmt.Sprint(id)).
    SetQuery("limit=1").
    End(&branch); err != nil {
    return myBranch.Branch{}
  }
  //branch.Links = myBranches.ItemLinks(branch)

  return branch
}

//func ItemEmbedded(data myShared.Class_) Embedded {
//  var embedded Embedded
//  lastSession := myShared.ClassItemEmbeddedLastSessions(data.Id)
//  //  links.LastSession = &lastSession
//  if lastSession.Id > 0 {
//    embedded.LastSession = &lastSession
//  }
//  module := myShared.ClassItemEmbeddedModule(data.ModuleId)
//  embedded.Module = &module
//  branch := ItemEmbeddedBranch(data.BranchId)
//  embedded.Branch = &branch
//  tutor := myShared.ClassItemEmbeddedTutor(data.TutorId)
//  embedded.Tutor = &tutor
//
//  embedded.Students = myShared.ClassItemEmbeddedStudents(data.Id)
//
//  return embedded
//}

func ItemLinks(data myShared.Class_) myShared.ClassLinks {
  var links myShared.ClassLinks
  links.Self = myShared.CreateHref(myShared.PathClasses + "/" + fmt.Sprint(data.Id))
  links.Module = myShared.CreateHref(myModule.Path + "/" + fmt.Sprint(data.ModuleId))
  links.Branch = myShared.CreateHref(myBranch.Path + "/" + fmt.Sprint(data.BranchId))
  links.Tutor = myShared.CreateHref(myUser.PathTutors + "/" + fmt.Sprint(data.TutorId))
  //links.Students = myShared.ClassLinksStudents(data.Id)
  //links.Sessions = myShared.ClassLinksSessions(data.Id)
  //if len(links.Sessions) > 0 {
  //  lastSession := itemLinksLastSessions(data.Id)
  //  links.LastSession = &lastSession
  //}
  return links
}

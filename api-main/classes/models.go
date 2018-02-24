package classes

import (
  myShared "../shared"
  mySessions "../sessions"
  myModules "../modules"
  myBranches "../branches"
  "strconv"
)

type (
  ClassEmbedded struct {
    LastSession *myShared.Session  `json:"last_session,omitempty"`
    Module      *myModules.Module  `json:"module,omitempty"`
    Branch      *myBranches.Branch `json:"branch,omitempty"`
  }
)

func itemEmbeddedLastSessions(id uint64) myShared.Session {
  var session myShared.Session

  _, err := myShared.GetItem(map[string]interface{}{
    "data": &session,
    "path": myShared.PathSessions,
    "query": map[string]string{
      "class_id": "eq." + strconv.FormatUint(id, 10),
      "order":    "created_at.desc",
      "limit":    "1",
      //"select":   "id,created_at",
    },
  })
  if err != nil {
    return myShared.Session{}
  }
  session.Links = mySessions.ItemLinks(session)

  return session
}

func itemEmbeddedModule(id uint64) myModules.Module {
  var module myModules.Module

  _, err := myShared.GetItem(map[string]interface{}{
    "data": &module,
    "path": myShared.PathModules,
    "query": map[string]string{
      "id":    "eq." + strconv.FormatUint(id, 10),
      "limit": "1",
      //"select":   "id,created_at",
    },
  })
  if err != nil {
    return myModules.Module{}
  }
  module.Links = myModules.ItemLinks(module)

  return module
}
func itemEmbeddedBranch(id uint64) myBranches.Branch {
  var branch myBranches.Branch

  _, err := myShared.GetItem(map[string]interface{}{
    "data": &branch,
    "path": myShared.PathBranches,
    "query": map[string]string{
      "id":    "eq." + strconv.FormatUint(id, 10),
      "limit": "1",
      //"select":   "id,created_at",
    },
  })
  if err != nil {
    return myBranches.Branch{}
  }
  branch.Links = myBranches.ItemLinks(branch)

  return branch
}
func itemEmbedded(data myShared.Class_) ClassEmbedded {
  var embedded ClassEmbedded
  lastSession := itemEmbeddedLastSessions(data.Id)
  //  links.LastSession = &lastSession
  if lastSession.Id > 0 {
    embedded.LastSession = &lastSession
  }
  module := itemEmbeddedModule(data.ModuleId)
  embedded.Module = &module
  branch := itemEmbeddedBranch(data.BranchId)
  embedded.Branch = &branch

  return embedded
}

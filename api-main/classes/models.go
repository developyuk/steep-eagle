package classes

import (
  myShared "../shared"
  myModule "../modules"
  myBranch "../branches"
  myUser "../users"
  "fmt"
  "time"
)

type (
  classLinks struct {
    myShared.LinksSelf
    Module *myShared.Href `json:"module,omitempty"`
    Branch *myShared.Href `json:"branch,omitempty"`
    Tutor  *myShared.Href `json:"tutor,omitempty"`
  }
  embedded struct {
    //LastSession *myShared.Session `json:"last_session,omitempty"`
    Module myModule.Module `json:"module,omitempty"`
    Branch myBranch.Branch `json:"branch,omitempty"`
    Tutor  myUser.User     `json:"tutor,omitempty"`
    //Students    []myUser.User   `json:"students,omitempty"`
  }

  Class_ struct {
    myShared.Hal
    Id         uint64    `json:"id"`
    Day        string    `json:"day"`
    StartAt    string    `json:"start_at"`
    FinishAt   string    `json:"finish_at"`
    StartAtTs  time.Time `json:"start_at_ts"`
    FinishAtTs time.Time `json:"finish_at_ts"`
    ModuleId   uint64    `json:"module_id"`
    BranchId   uint64    `json:"branch_id"`
    TutorId    uint64    `json:"tutor_id"`
  }
)

const Path string = "/classes"

func itemLinks(data Class_) classLinks {
  var links classLinks
  links.Self = myShared.CreateHref(Path + "/" + fmt.Sprint(data.Id))
  links.Module = myShared.CreateHref(myModule.Path + "/" + fmt.Sprint(data.ModuleId))
  links.Branch = myShared.CreateHref(myBranch.Path + "/" + fmt.Sprint(data.BranchId))
  links.Tutor = myShared.CreateHref(myUser.PathTutors + "/" + fmt.Sprint(data.TutorId))
  return links
}

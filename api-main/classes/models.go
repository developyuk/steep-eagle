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
  Links struct {
    myShared.LinksSelf
    Module      *myShared.Href `json:"module,omitempty"`
    Branch      *myShared.Href `json:"branch,omitempty"`
    Tutor       *myShared.Href `json:"tutor,omitempty"`
    LastSession *myShared.Href `json:"last_session,omitempty"`
  }
  Embedded struct {
    //LastSession *myShared.Session `json:"last_session,omitempty"`
    Module myModule.Module `json:"module,omitempty"`
    Branch myBranch.Branch `json:"branch,omitempty"`
    Tutor  myUser.User     `json:"tutor,omitempty"`
    //Students []myUser.User   `json:"students,omitempty"`
  }

  Class_ struct {
    myShared.Hal
    Id              uint64    `json:"id"`
    Day             string    `json:"day"`
    StartAt         string    `json:"start_at"`
    FinishAt        string    `json:"finish_at"`
    StartAtTs       time.Time `json:"start_at_ts"`
    FinishAtTs      time.Time `json:"finish_at_ts"`
    ProgramModuleID uint64    `json:"program_module_id"`
    BranchId        uint64    `json:"branch_id"`
    TutorId         uint64    `json:"tutor_id"`
    Q               string    `json:"q"`
  }
)

const Path string = "/classes"

func ItemLinks(data *Class_) Links {
  var links Links
  links.Self = myShared.CreateHref(Path + "/" + fmt.Sprint(data.Id))
  moduleId, _ := myModule.GetModuleFromProgramModuleId(fmt.Sprint(data.ProgramModuleID))
  links.Module = myShared.CreateHref(myModule.Path + "/" + moduleId)
  links.Branch = myShared.CreateHref(myBranch.Path + "/" + fmt.Sprint(data.BranchId))
  links.Tutor = myShared.CreateHref(myUser.PathTutors + "/" + fmt.Sprint(data.TutorId))
  now := time.Now()
  today := now.Add(-2 * time.Hour).Format("2006-01-02T15:04:05")
  links.LastSession = myShared.CreateHref(Path + "/" + fmt.Sprint(data.Id) + "/sessions?sort=created_at.desc&created_at=gte." + today)
  return links
}

package programs

import (
  myShared "../shared"
  myProgramTypes "./types"
  "fmt"
)

const Path string = "/programs"

type (
  Program struct {
    myShared.Hal
    Id     uint64 `json:"id"`
    Name   string `json:"name,omitempty"`
    TypeId uint64 `json:"type_id"`
  }

  ProgramModules struct {
    Id        uint64 `json:"id"`
    ProgramId uint64 `json:"program_id"`
    ModuleId  uint64 `json:"module_id"`
  }

  ProgramLinks struct {
    myShared.LinksSelf
    Type *myShared.Href `json:"type"`
  }
)

func itemLinks(v Program) ProgramLinks {
  var links ProgramLinks
  links.Self = myShared.CreateHref(Path + "/" + fmt.Sprint(v.Id))
  links.Type = myShared.CreateHref(myProgramTypes.Path + "/" + fmt.Sprint(v.TypeId))
  return links
}

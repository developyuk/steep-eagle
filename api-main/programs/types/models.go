package types

import (
  myShared "../../shared"
  "strconv"
)

type (
  ProgramType struct {
    myShared.Hal
    Id   uint64 `json:"id"`
    Name string `json:"name,omitempty"`
  }
  ProgramLinks struct {
    myShared.LinksSelf
  }
)

func itemLinks(v ProgramType) ProgramLinks {
  var links ProgramLinks
  links.Self = myShared.CreateHref(myShared.PathProgramsTypes + "/" + strconv.FormatUint(v.Id, 10))
  return links
}

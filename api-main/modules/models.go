package modules

import (
  myShared "../shared"
  //myRest "../shared/rest"
  //myPrograms "../programs"
  "fmt"
)

type (
  Module struct {
    myShared.Hal
    Id    uint64 `json:"id"`
    Name  string `json:"name,omitempty"`
    Image string `json:"image,omitempty"`
  }
  ModuleLinks struct {
    myShared.LinksSelf
    //Programs []myShared.Href `json:"programs,omitempty"`
    //Classes  []myShared.Href `json:"classes,omitempty"`
  }
)

var Path string = "/modules"

func itemLinks(v Module) ModuleLinks {
  var links ModuleLinks
  links.Self = myShared.CreateHref(Path+ "/" + fmt.Sprint(v.Id))
  //links.Programs = itemLinksPrograms(v.Id)
  return links
}

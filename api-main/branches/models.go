package branches

import (
  myShared "../shared"
  "fmt"
)

const (
  Path string = "/branches"
)

type (
  Branch struct {
    myShared.Hal
    Id      uint64 `json:"id"`
    Name    string `json:"name,omitempty"`
    Image   string `json:"image,omitempty"`
    Address string `json:"address,omitempty"`
  }
)

func itemLinks(v Branch) myShared.LinksSelf {
  links := myShared.CreateHrefSelf(Path + "/" + fmt.Sprint(v.Id))

  return links
}

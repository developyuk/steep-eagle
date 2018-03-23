package branches

import (
  myShared "../shared"
  mySharedRest "../shared/rest"
  "fmt"
  "net/http"
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

func ItemRest(req *mySharedRest.Request, id string, item *Branch) (*http.Response, error) {
  resp, err := mySharedRest.New().GetItem(Path).ParseRequest(req).
    SetQuery(myShared.RequestRest{
    Id:     "eq." + id,
    Select: req.Select,
  }).End(item)
  if err != nil {
    return resp, err
  }
  item.setItemLinks()

  return nil, nil

}
func (v *Branch) setItemLinks() *Branch {
  v.Links = myShared.LinksSelf{
    Self:myShared.CreateHref(Path + "/" + fmt.Sprint(v.Id)),
  }
  return v

}

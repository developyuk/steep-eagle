package modules

import (
  myShared "../shared"
  mySharedRest "../shared/rest"
  "fmt"
  "net/http"
)

type (
  Module struct {
    myShared.Hal
    Id    uint64 `json:"id"`
    Name  string `json:"name,omitempty"`
    Image string `json:"image,omitempty"`
  }
)

var Path string = "/modules"

func ItemRest(req *mySharedRest.Request, id string, item *Module) (*http.Response, error) {
  resp, err := mySharedRest.New().GetItem(Path).ParseRequest(req).
    SetQuery(myShared.RequestRest{
    Id:     "eq." + id,
    Select: req.Select,
  }).End(&item)
  if err != nil {
    return resp, err
  }
  item.setItemLinks()
  return resp, nil
}

func (v *Module) setItemLinks() *Module {
  v.Links = myShared.LinksSelf{
    Self: myShared.CreateHref(Path + "/" + fmt.Sprint(v.Id)),
  }
  return v

}

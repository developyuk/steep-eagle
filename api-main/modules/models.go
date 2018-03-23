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

func GetModuleFromProgramModuleId(programModuleId string) (string, error) {
  var itemProgramModule myShared.ProgramModule
  _, err := mySharedRest.New().GetItem("/programs_modules").
    SetQuery(myShared.RequestRest{
    Id: "eq." + programModuleId,
  }).EndStruct(&itemProgramModule)

  if err != nil {
    return fmt.Sprint(0), err
  }

  //moduleId, _ := GetModuleFromProgramModuleId(id)
  return fmt.Sprint(itemProgramModule.ModuleId), nil
}

func ItemRest(req *mySharedRest.Request, id string, item *Module) (*http.Response, error) {

  resp, err := mySharedRest.New().GetItem(Path).ParseRequest(req).
    SetQuery(myShared.RequestRest{
    Id:     "eq." + id,
    Select: req.Select,
  }).EndStruct(&item)

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

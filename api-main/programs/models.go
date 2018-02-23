package programs

import (
  myShared "../shared"
  "strconv"
)

type (
  Program struct {
    myShared.Hal
    Id     uint64 `json:"id"`
    Name   string `json:"name"`
    TypeId uint64 `json:"type_id"`
  }

  ProgramModules struct {
    Id        uint64 `json:"id"`
    ProgramId uint64 `json:"program_id"`
    ModuleId  uint64 `json:"module_id"`
  }

  ProgramLinks struct {
    myShared.LinksSelf
    Type    myShared.Href   `json:"type"`
    Modules []myShared.Href `json:"modules,omitempty"`
  }
)

func itemLinksModules(id uint64) []myShared.Href {
  var list []ProgramModules
  myShared.GetItems(map[string]interface{}{
    "data": &list,
    "path": "/program_modules",
    "query": map[string]string{
      "program_id": "eq." + strconv.FormatUint(id, 10),
      "select":     "module_id",
    },
  })

  var data []myShared.Href
  for _, v := range list {
    data = append(data, myShared.CreateHref(myShared.PathModules+"/"+strconv.FormatUint(v.ModuleId, 10)))
  }

  return data
}

func itemLinks(v Program) ProgramLinks {
  var links ProgramLinks
  links.Self = myShared.CreateHref(myShared.PathPrograms + "/" + strconv.FormatUint(v.Id, 10))
  links.Type = myShared.CreateHref(myShared.PathProgramsTypes + "/" + strconv.FormatUint(v.TypeId, 10))
  links.Modules = itemLinksModules(v.Id)
  return links
}

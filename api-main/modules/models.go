package modules

import (
  myShared "../shared"
  myPrograms "../programs"
  "strconv"
)

type (

  ModuleLinks struct {
    myShared.LinksSelf
    Programs []myShared.Href `json:"programs,omitempty"`
    Classes  []myShared.Href `json:"classes,omitempty"`
  }
)

func itemLinksClasses(id uint64) []myShared.Href {
  var list []myShared.Class_
  myShared.GetItems(map[string]interface{}{
    "data": &list,
    "path": "/classes",
    "query": map[string]string{
      "module_id": "eq." + strconv.FormatUint(id, 10),
      "select":    "id",
    },
  })
  //log.Println(resp,list,id)

  var data []myShared.Href
  for _, v := range list {
    data = append(data, myShared.CreateHref(myShared.PathClasses+"/"+strconv.FormatUint(v.Id, 10)))
  }
  return data
}
func itemLinksPrograms(id uint64) []myShared.Href {
  var list []myPrograms.ProgramModules
  myShared.GetItems(map[string]interface{}{
    "data": &list,
    "path": "/program_modules",
    "query": map[string]string{
      "module_id": "eq." + strconv.FormatUint(id, 10),
      "select":    "program_id",
    },
  })
  //log.Println(resp,list,id)

  var data []myShared.Href
  for _, v := range list {
    data = append(data, myShared.CreateHref(myShared.PathPrograms+"/"+strconv.FormatUint(v.ProgramId, 10)))
  }
  return data
}

func ItemLinks(v myShared.Module) ModuleLinks {
  var links ModuleLinks
  links.Self = myShared.CreateHref(myShared.PathModules + "/" + strconv.FormatUint(v.Id, 10))
  links.Programs = itemLinksPrograms(v.Id)
  //links.Classes = itemLinksClasses(v.Id)
  return links
}

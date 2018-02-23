package types

import (
  myShared "../../shared"
  myPrograms "../../programs"
  "strconv"
)

type (
  ProgramType struct {
    myShared.Hal
    Id   uint64 `json:"id"`
    Name string `json:"name"`
  }
  ProgramLinks struct {
    myShared.LinksSelf
    Programs []myShared.Href `json:"programs,omitempty"`
  }
)

func itemLinksPrograms(id uint64) []myShared.Href {
  var list []myPrograms.ProgramModules
  myShared.GetItems(map[string]interface{}{
    "data": &list,
    "path": "/programs",
    "query": map[string]string{
      "type_id": "eq." + strconv.FormatUint(id, 10),
      "select":  "id",
    },
  })

  var data []myShared.Href
  for _, v := range list {
    data = append(data, myShared.CreateHref(myShared.PathPrograms+"/"+strconv.FormatUint(v.Id, 10)))
  }
  return data
}

func itemLinks(v ProgramType) ProgramLinks {
  var links ProgramLinks
  links.Self = myShared.CreateHref(myShared.PathProgramsTypes + "/" + strconv.FormatUint(v.Id, 10))
  links.Programs = itemLinksPrograms(v.Id)
  return links
}

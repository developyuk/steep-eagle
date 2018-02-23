package branches

import (
  myShared "../shared"
  myClasses "../classes"
  "strconv"
)

type (
  Branch struct {
    myShared.Hal
    Id      uint64 `json:"id"`
    Name    string `json:"name"`
    Image   string `json:"image"`
    Address string `json:"address"`
  }

  BranchLinks struct {
    myShared.LinksSelf
    Classes []myShared.Href `json:"classes,omitempty"`
  }
)

func itemLinksClasses(id uint64) []myShared.Href {
  //list := getClassesById(db, id)
  var list []myClasses.Class_

  myShared.GetItems(map[string]interface{}{
    "data": &list,
    "path": myShared.PathClasses,
    "query": map[string]string{
      "branch_id": "eq." + strconv.FormatUint(id, 10),
      "select":    "id",
    },
  })

  var data []myShared.Href
  for _, v := range list {
    data = append(data, myShared.CreateHref(myShared.PathClasses+"/"+strconv.FormatUint(v.Id, 10)))
  }

  return data
}

func itemLinks(v Branch) BranchLinks {
  var links BranchLinks
  links.Self = myShared.CreateHref(myShared.PathBranches + "/" + strconv.FormatUint(v.Id, 10))
  links.Classes = itemLinksClasses(v.Id)
  return links
}

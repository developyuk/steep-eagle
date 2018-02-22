package modules

import (
  myShared "../shared"
  myPrograms "../programs"
  "fmt"
  "github.com/labstack/echo"
  "net/http"
  "encoding/json"
  "strconv"
  "log"
)

type ModuleLinks struct {
  myShared.LinksSelf
  Programs []myShared.Href `json:"programs,omitempty"`
}

func itemLinksPrograms(id int) []myShared.Href {
  var list []myPrograms.ProgramModules
  // program_modules?module_id=eq.1&select=program_id
  resp := myShared.RestItems(map[string]interface{}{
    "path": "program_modules",
    "query": map[string]string{
      "module_id": "eq." + strconv.Itoa(id),
      "select":    "program_id",
    },
  })
  _ = json.Unmarshal(resp.Body(), &list)
  log.Println(resp)

  var data []myShared.Href
  for _, v := range list {
    data = append(data, myShared.CreateHref(fmt.Sprintf("%v/%v", myShared.PathPrograms, v.ProgramId)))
  }
  return data
}

func itemLinks(v Module) ModuleLinks {
  var links ModuleLinks
  links.Self = myShared.CreateHref(fmt.Sprintf("%v/%v", myShared.PathModules, v.Id))
  links.Programs = itemLinksPrograms(v.Id)
  return links
}

func List(c echo.Context) error {
  var list []Module
  resp := myShared.RestItems(map[string]interface{}{
    "path": myShared.PathModules,
  })
  _ = json.Unmarshal(resp.Body(), &list)

  for i, v := range list {
    list[i].Links = itemLinks(v)
  }

  response := myShared.Hal{
    Links:    myShared.LinksSelf{Self: myShared.CreateHref(myShared.PathModules)},
    Embedded: list,
    Count:    len(list),
    Total:    len(list),
  }

  return c.JSON(http.StatusOK, response)
}

func Item(c echo.Context) error {
  // var data Program = GetProgramTypesData(c.Param("id"))
  var item Module
  resp := myShared.RestItem(map[string]interface{}{
    "path": myShared.PathModules,
    "query": map[string]string{
      "id": "eq." + c.Param("id"),
    },
  })
  _ = json.Unmarshal(resp.Body(), &item)

  item.Links = itemLinks(item)
  return c.JSON(http.StatusOK, item)
}

// func UpdateModule(c echo.Context) error {
// 	var data myModels.ProgramType
// 	return c.JSON(http.StatusOK, data)
// }
//
// func DeleteModule(c echo.Context) error {
// 	// id := c.Param("id")
// 	var data myModels.ProgramType
// 	return c.JSON(http.StatusOK, data)
// }

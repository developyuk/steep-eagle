package branches

import (
  myShared "../shared"
  myClasses "../classes"
  "fmt"
  "github.com/labstack/echo"
  "net/http"
  "encoding/json"
  "strconv"
)

type BranchLinks struct {
  myShared.LinksSelf
  Classes []myShared.Href `json:"classes,omitempty"`
}

func itemClassesHref(id int) []myShared.Href {
  //classes := getClassesById(db, id)
  var classes []myClasses.Class_

  resp := myShared.RestItems(map[string]interface{}{
    "path": myShared.PathClasses,
    "query": map[string]string{
      "branch_id": "eq." + strconv.Itoa(id),
      "select":    "id",
    },
  })
  _ = json.Unmarshal(resp.Body(), &classes)

  var data []myShared.Href
  for _, v := range classes {
    data = append(data, myShared.CreateHref(fmt.Sprintf("%v/%v", myShared.PathClasses, v.Id)))
  }

  return data
}

func itemLinks(v Branch) BranchLinks {
  var links BranchLinks
  links.Self = myShared.CreateHref(fmt.Sprintf("%v/%v", myShared.PathBranches, v.Id))
  //links.Classes = itemClassesHref(v.Id)
  return links
}
func List(c echo.Context) error {
  var data []Branch
  resp := myShared.RestItems(map[string]interface{}{
    "path": myShared.PathBranches,
  })
  _ = json.Unmarshal(resp.Body(), &data)

  for i, v := range data {
    data[i].Links = itemLinks(v)
  }
  response := myShared.Hal{
    Links:    myShared.LinksSelf{Self: myShared.CreateHref(myShared.PathBranches)},
    Embedded: data,
    Count:    len(data),
    Total:    len(data),
  }
  return c.JSON(http.StatusOK, response)
}

func Item(c echo.Context) error {
  var item Branch
  resp := myShared.RestItem(map[string]interface{}{
    "path": myShared.PathBranches,
    "query": map[string]string{
      "id": "eq." + c.Param("id"),
    },
  })
  _ = json.Unmarshal(resp.Body(), &item)

  item.Links = itemLinks(item)
  return c.JSON(http.StatusOK, item)
}

// func UpdateBranch(c echo.Context) error {
// 	var data Branch
// 	return c.JSON(http.StatusOK, data)
// }
//
// func DeleteBranch(c echo.Context) error {
// 	// id := c.Param("id")
// 	var data Branch
// 	return c.JSON(http.StatusOK, data)
// }

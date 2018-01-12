package branches

import (
	myShared "../shared"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"net/http"
)

type BranchLinks struct {
	myShared.LinksSelf
	Classes []myShared.Href `json:"classes,omitempty"`
}

var (
	db   *sqlx.DB
	list []Branch
	item Branch
)

func itemClassesHref(id int) []myShared.Href {
	classes := getClassesById(db, id)
	var data []myShared.Href
	for _, v := range classes {
		data = append(data, myShared.CreateHref(fmt.Sprintf("%v/%v", myShared.PathClasses, v)))
	}

	return data
}

func itemLinks(v Branch) BranchLinks {
	return BranchLinks{
		LinksSelf: myShared.LinksSelf{Self: myShared.CreateHref(fmt.Sprintf("%v/%v", myShared.PathBranches, v.Id))},
		Classes:   itemClassesHref(v.Id),
	}
}
func List(c echo.Context) error {
	// User ID from path `users/:id`
	list, db = ListData()

	for i, v := range list {
		list[i].Links = itemLinks(v)
	}
	response := myShared.Hal{
		Links:    myShared.LinksSelf{Self: myShared.Href{Href: myShared.PathBranches}},
		Embedded: list,
		Count:    len(list),
		Total:    len(list),
	}
	return c.JSON(http.StatusOK, response)
}

func Item(c echo.Context) error {
	// User ID from path `users/:id`
	// id := c.Param("id")
	item, db = ItemData(c.Param("id"))
	item.Links = itemLinks(item)
	return c.JSON(http.StatusOK, item)
}

// func UpdateBranch(c echo.Context) error {
// 	// User ID from path `users/:id`
// 	var data Branch
// 	return c.JSON(http.StatusOK, data)
// }
//
// func DeleteBranch(c echo.Context) error {
// 	// User ID from path `users/:id`
// 	// id := c.Param("id")
// 	var data Branch
// 	return c.JSON(http.StatusOK, data)
// }

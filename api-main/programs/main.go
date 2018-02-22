package programs

import (
	myShared "../shared"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"net/http"
)

type ProgramLinks struct {
	myShared.LinksSelf
	Type    myShared.Href   `json:"type"`
	Modules []myShared.Href `json:"modules,omitempty"`
}

var (
	db   *sqlx.DB
	list []Program
	item Program
)

func itemModulesHref(id int) []myShared.Href {
	modules := ItemModulesHrefData(db, id)
	var data []myShared.Href
	for _, v := range modules {
		data = append(data, myShared.Href{Href: fmt.Sprintf("%v/%v", myShared.PathModules, v)})
	}

	return data
}

func itemLinks(v Program) ProgramLinks {
	return ProgramLinks{
		LinksSelf: myShared.LinksSelf{Self: myShared.Href{Href: fmt.Sprintf("%v/%v", myShared.PathPrograms, v.Id)}},
		Type:      myShared.Href{Href: fmt.Sprintf("%v/%v", myShared.PathProgramsTypes, v.TypeId)},
		Modules:   itemModulesHref(v.Id),
	}
}

func List(c echo.Context) error {
	list, db = ListData()

	for i, v := range list {
		list[i].Links = itemLinks(v)
	}
	response := myShared.Hal{
		Links:    myShared.LinksSelf{Self: myShared.Href{Href: myShared.PathPrograms}},
		Embedded: list,
		Count:    len(list),
		Total:    len(list),
	}

	return c.JSON(http.StatusOK, response)
}

func Item(c echo.Context) error {
	// id := c.Param("id")
	item, db = ItemData(c.Param("id"))

	item.Links = itemLinks(item)
	return c.JSON(http.StatusOK, item)
}

// func UpdateProgram(c echo.Context) error {
// 	var data myModels.Program
// 	return c.JSON(http.StatusOK, data)
// }
//
// func DeleteProgram(c echo.Context) error {
// 	// id := c.Param("id")
// 	var data myModels.Program
// 	return c.JSON(http.StatusOK, data)
// }

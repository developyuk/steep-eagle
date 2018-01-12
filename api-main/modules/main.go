package modules

import (
	myShared "../shared"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func itemLinks(v Module) myShared.LinksSelf {
	return myShared.LinksSelf{Self: myShared.Href{
		Href: fmt.Sprintf("%v/%v", myShared.PathModules, v.Id),
	}}
}

func List(c echo.Context) error {
	// User ID from path `users/:id`
	list := ListData()

	for i, v := range list {
		list[i].Links = itemLinks(v)
	}

	response := myShared.Hal{
		Links:    myShared.LinksSelf{Self: myShared.Href{Href: myShared.PathModules}},
		Embedded: list,
		Count:    len(list),
		Total:    len(list),
	}

	return c.JSON(http.StatusOK, response)
}

func Item(c echo.Context) error {
	// User ID from path `users/:id`
	// var data Program = GetProgramTypesData(c.Param("id"))
	item := ItemData(c.Param("id"))
	item.Links = itemLinks(item)
	return c.JSON(http.StatusOK, item)
}

// func UpdateModule(c echo.Context) error {
// 	// User ID from path `users/:id`
// 	var data myModels.ProgramType
// 	return c.JSON(http.StatusOK, data)
// }
//
// func DeleteModule(c echo.Context) error {
// 	// User ID from path `users/:id`
// 	// id := c.Param("id")
// 	var data myModels.ProgramType
// 	return c.JSON(http.StatusOK, data)
// }

package types

import (
	myShared "../../shared"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func itemLinks(v ProgramType) myShared.LinksSelf {
	return myShared.LinksSelf{Self: myShared.Href{
		Href: fmt.Sprintf("%v/%v", myShared.PathProgramsTypes, v.Id),
	}}
}

func List(c echo.Context) error {
	// User ID from path `users/:id`
	list := ListData()

	for i, v := range list {
		list[i].Links = itemLinks(v)
	}

	response := myShared.Hal{
		Links:    myShared.LinksSelf{Self: myShared.Href{Href: myShared.PathProgramsTypes}},
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

// func CreateProgramType(c echo.Context) error {
// 	// User ID from path `users/:id`
// 	// var data Program = GetProgramTypesData(c.Param("id"))
// 	var data myModels.Response = myModels.CreateProgramType(c.FormValue("name"))
// 	return c.JSON(http.StatusOK, data)
// }
//
// func UpdateProgramType(c echo.Context) error {
// 	// User ID from path `users/:id`
// 	var data myModels.ProgramType
// 	return c.JSON(http.StatusOK, data)
// }
//
// func DeleteProgramType(c echo.Context) error {
// 	// User ID from path `users/:id`
// 	// id := c.Param("id")
// 	var data myModels.ProgramType
// 	return c.JSON(http.StatusOK, data)
// }

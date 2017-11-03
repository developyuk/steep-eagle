package controllers

import (
	myModels "../models"
	"github.com/labstack/echo"
	"net/http"
	"fmt"
)
var pathProgramsTypes string = "/programs/types"

func GetProgramTypes(c echo.Context) error {
	// User ID from path `users/:id`
	data := myModels.GetProgramTypes()

	for i, v := range data {
		data[i].Links = LinksSelf{Self: Href{
			Href: fmt.Sprintf("%v/%v", pathProgramsTypes, v.Id),
		}}
	}

	response := myModels.Hal{
		Links:    LinksSelf{Self: Href{Href: pathProgramsTypes}},
		Embedded: data,
	}
	return c.JSON(http.StatusOK, response)
}

func GetProgramType(c echo.Context) error {
	// User ID from path `users/:id`
	// var data Program = GetProgramTypesData(c.Param("id"))
	data := myModels.GetProgramType(c.Param("id"))
  response := data
	response.Links = LinksSelf{Self: Href{
		Href: fmt.Sprintf("%v/%v", pathProgramsTypes, data.Id),
	}}
	return c.JSON(http.StatusOK, response)
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

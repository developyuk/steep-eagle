package controllers

import (
	myModels "../models"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

var pathModules string = "/modules"

func GetModules(c echo.Context) error {
	// User ID from path `users/:id`
	data := myModels.GetModules()

	for i, v := range data {
		data[i].Links = LinksSelf{Self: Href{
			Href: fmt.Sprintf("%v/%v", pathModules, v.Id),
		}}
	}
	response := myModels.Hal{
		Links:    LinksSelf{Self: Href{Href: pathModules}},
		Embedded: data,
	}

	return c.JSON(http.StatusOK, response)
}

func GetModule(c echo.Context) error {
	// User ID from path `users/:id`
	// var data Program = GetProgramTypesData(c.Param("id"))
	response := myModels.GetModule(c.Param("id"))
	response.Links = LinksSelf{Self: Href{
		Href: fmt.Sprintf("%v/%v", pathModules, response.Id),
	}}
	return c.JSON(http.StatusOK, response)
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

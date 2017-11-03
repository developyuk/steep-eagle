package controllers

import (
	myModels "../models"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"net/http"
)

type ProgramLinks struct {
	LinksSelf
	Type    Href   `json:"type"`
	Modules []Href `json:"modules,omitempty"`
}

var pathPrograms string = "/programs"

func GetProgramModules(db *sqlx.DB, id int) []Href {
	modules := myModels.GetModulesFromId(db, id)
	var data []Href
	for _, v := range modules {
		data = append(data, Href{Href: fmt.Sprintf("%v/%v", pathModules, v)})
	}

	return data
}

func GetPrograms(c echo.Context) error {
	// User ID from path `users/:id`
	// User ID from path `users/:id`
	data, db := myModels.GetPrograms()

	for i, v := range data {
		data[i].Links = ProgramLinks{
			LinksSelf: LinksSelf{Self: Href{Href: fmt.Sprintf("%v/%v", pathPrograms, v.Id)}},
			Type:      Href{Href: fmt.Sprintf("%v/%v", pathProgramsTypes, v.TypeId)},
			Modules:   GetProgramModules(db, v.Id),
		}
	}
	response := myModels.Hal{
		Links:    LinksSelf{Self: Href{Href: pathPrograms}},
		Embedded: data,
	}

	return c.JSON(http.StatusOK, response)
}

func GetProgram(c echo.Context) error {
	// User ID from path `users/:id`
	// id := c.Param("id")
	response, db := myModels.GetProgram(c.Param("id"))

	response.Links = ProgramLinks{
		LinksSelf: LinksSelf{Self: Href{Href: fmt.Sprintf("%v/%v", pathPrograms, response.Id)}},
		Type:      Href{Href: fmt.Sprintf("%v/%v", pathProgramsTypes, response.TypeId)},
		Modules:   GetProgramModules(db, response.Id),
	}
	return c.JSON(http.StatusOK, response)
}

// func UpdateProgram(c echo.Context) error {
// 	// User ID from path `users/:id`
// 	var data myModels.Program
// 	return c.JSON(http.StatusOK, data)
// }
//
// func DeleteProgram(c echo.Context) error {
// 	// User ID from path `users/:id`
// 	// id := c.Param("id")
// 	var data myModels.Program
// 	return c.JSON(http.StatusOK, data)
// }

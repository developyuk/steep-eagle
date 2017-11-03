package controllers

import (
	myModels "../models"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"net/http"
	"time"
	"strings"
)

type ClassLinks struct {
	LinksSelf
	Module   Href   `json:"module"`
	Branch   Href   `json:"branch"`
	Students []Href `json:"students,omitempty"`
	Sessions []Href `json:"sessions,omitempty"`
}

var pathClasses string = "/classes"

func GetClassStudents(db *sqlx.DB, id int) []Href {
	students := myModels.GetStudentFromId(db, id)
	var data []Href
	for _, v := range students {
		data = append(data, Href{Href: fmt.Sprintf("%v/%v", pathStudents, v)})
	}

	return data
}

func GetClassSessions(db *sqlx.DB, id int) []Href {
	sessions := myModels.GetSessionFromId(db, id)
	var data []Href
	for _, v := range sessions {
		data = append(data, Href{Href: fmt.Sprintf("%v/%v", pathSessions, v)})
	}

	return data
}

func GenerateLinks(db *sqlx.DB, data myModels.Class_) ClassLinks {
	values := ClassLinks{
		LinksSelf: LinksSelf{Self: Href{Href: fmt.Sprintf("%v/%v", pathClasses, data.Id)}},
		Module:    Href{Href: fmt.Sprintf("%v/%v", pathModules, data.ModuleId)},
		Branch:    Href{Href: fmt.Sprintf("%v/%v", pathBranches, data.BranchId)},
		Students:  GetClassStudents(db, data.Id),
		Sessions:  GetClassSessions(db, data.Id),
	}
	return values
}

func GetClasses(c echo.Context) error {
	params := make(map[string]interface{})

	day := c.QueryParam("day")
	// fmt.Printf("%#v", day)
	if day != "" {
		if day == "today" {
			params["day"] = strings.ToLower(time.Now().Weekday().String())
		}
	}
	data, db := myModels.GetClasses(params)

	for i, v := range data {
		data[i].Links = GenerateLinks(db, v)
	}

	response := myModels.Hal{
		Links:    LinksSelf{Self: Href{Href: pathClasses}},
		Embedded: data,
		Count:    len(data),
	}
	return c.JSON(http.StatusOK, response)
}

func GetClass(c echo.Context) error {
	// User ID from path `users/:id`
	// var data Program = GetProgramTypesData(c.Param("id"))
	response, db := myModels.GetClass(c.Param("id"))
	response.Links = GenerateLinks(db, response)
	return c.JSON(http.StatusOK, response)
}

// func UpdateClass(c echo.Context) error {
// 	// User ID from path `users/:id`
// 	var data myModels.ProgramType
// 	return c.JSON(http.StatusOK, data)
// }
//
// func DeleteClass(c echo.Context) error {
// 	// User ID from path `users/:id`
// 	// id := c.Param("id")
// 	var data myModels.ProgramType
// 	return c.JSON(http.StatusOK, data)
// }

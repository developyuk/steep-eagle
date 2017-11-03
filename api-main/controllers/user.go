package controllers

import (
	myModels "../models"
	// "fmt"
	// "github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"net/http"
)
var pathStudents string = "/students"
var pathTutors string = "/tutors"

func GetUsers(c echo.Context) error {
	// User ID from path `users/:id`
	data, _ := myModels.GetUsers()

    // log.Fatal()
	// for i, v := range data {
	// 	data[i].Links = ClassLinks{
	// 		LinksSelf: LinksSelf{Self: Href{Href: fmt.Sprintf("%v/%v", pathClasses, v.Id)}},
	// 		// Module:    Href{Href: fmt.Sprintf("%v/%v", pathModules, v.ModuleId)},
	// 		// Branch:    Href{Href: fmt.Sprintf("%v/%v", pathBranches, v.BranchId)},
	// 		// Students:  GetClassStudents(db, v.Id),
	// 		// Sessions:  GetClassSessions(db, v.Id),
	// 	}
	// }

	response := myModels.Hal{
		Links:    LinksSelf{Self: Href{Href: pathClasses}},
		Embedded: data,
	}
	return c.JSON(http.StatusOK, response)
}

func GetUser(c echo.Context) error {
	// User ID from path `users/:id`
	// var data Program = GetProgramTypesData(c.Param("id"))
	response, _ := myModels.GetUser(c.Param("id"))
	// response.Links = ClassLinks{
	// 	LinksSelf: LinksSelf{Self: Href{Href: fmt.Sprintf("%v/%v", pathClasses, response.Id)}},
	// 	// Module:    Href{Href: fmt.Sprintf("%v/%v", pathModules, response.ModuleId)},
	// 	// Branch:    Href{Href: fmt.Sprintf("%v/%v", pathBranches, response.BranchId)},
	// 	// Students:  GetClassStudents(db, response.Id),
  //   // Sessions:  GetClassSessions(db, response.Id),
	// }
	return c.JSON(http.StatusOK, response)
}

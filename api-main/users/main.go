package users

import (
	myShared "../shared"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"gopkg.in/guregu/null.v3"
	"net/http"
	"strings"
)

type StudentLinks struct {
	myShared.LinksSelf
	//Class myShared.Href `json:"class,omitempty"`
}

var (
	db   *sqlx.DB
	list []User
	item User
)

func getRole(path string) string {
	role := strings.TrimRight(strings.TrimPrefix(path, "/"), "s")
	role = strings.TrimRight(strings.TrimPrefix(path, "/"), "s/:id")
	return role
}

func getPath(role string) string {
	path := myShared.PathUsers
	if role == "student" {
		path = myShared.PathStudents
	}
	if role == "tutor" {
		path = myShared.PathTutors
	}
	return path
}

func itemClassHref(id null.Int) myShared.Href {
	class := getClassById(db, id)
	data := myShared.CreateHref(fmt.Sprintf("%v/%v", myShared.PathClasses, class))

	return data
}

func itemLinks(v User, role string) StudentLinks {
	path := getPath(role)
	var studentLinks StudentLinks
	studentLinks.LinksSelf = myShared.LinksSelf{Self: myShared.CreateHref(fmt.Sprintf("%v/%v", path, v.Id.Int64))}
	//if role == "student" {
	//	studentLinks.Class = itemClassHref(v.Id)
	//}

	return studentLinks
}

func List(c echo.Context) error {
	params := make(map[string]interface{})
	role := getRole(c.Path())
	if role != "user" {
		params["role"] = role
	}
	list, db = ListData(params)

	// log.Fatal()
	for i, v := range list {
		list[i].Links = itemLinks(v, role)
	}

	path := getPath(role)
	response := myShared.Hal{
		Links:    myShared.LinksSelf{Self: myShared.Href{Href: path}},
		Embedded: list,
		Count:    len(list),
		Total:    len(list),
	}
	return c.JSON(http.StatusOK, response)
}

func Item(c echo.Context) error {
	// var data Program = GetProgramTypesData(c.Param("id"))
	params := make(map[string]interface{})
	params["id"] = c.Param("id")
	role := getRole(c.Path())
	if role != "user" {
		params["role"] = role
	}

	item, db = ItemData(params)
	item.Links = itemLinks(item, role)
	return c.JSON(http.StatusOK, item)
}

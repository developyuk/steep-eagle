package controllers

import (
	myModels "../models"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

var pathBranches string = "/branches"

func GetBranches(c echo.Context) error {
	// User ID from path `users/:id`
	// User ID from path `users/:id`
	data := myModels.GetBranches()

	for i, v := range data {
		data[i].Links = LinksSelf{Self: Href{
			Href: fmt.Sprintf("%v/%v", pathBranches, v.Id),
		}}
	}
	response := myModels.Hal{
		Links:    LinksSelf{Self: Href{Href: pathBranches}},
		Embedded: data,
	}
	return c.JSON(http.StatusOK, response)
}

func GetBranch(c echo.Context) error {
	// User ID from path `users/:id`
	// id := c.Param("id")
	response := myModels.GetBranch(c.Param("id"))
	response.Links = LinksSelf{Self: Href{
		Href: fmt.Sprintf("%v/%v", pathBranches, response.Id),
	}}
	return c.JSON(http.StatusOK, response)
}

func UpdateBranch(c echo.Context) error {
	// User ID from path `users/:id`
	var data myModels.Program
	return c.JSON(http.StatusOK, data)
}

func DeleteBranch(c echo.Context) error {
	// User ID from path `users/:id`
	// id := c.Param("id")
	var data myModels.Program
	return c.JSON(http.StatusOK, data)
}

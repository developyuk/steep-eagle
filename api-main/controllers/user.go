package controllers

import (
	myModels "../models"
	// "fmt"
	// "github.com/jmoiron/sqlx"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"net/http"
	"time"
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

func Sign(c echo.Context) error {
	p := new(myModels.UserLoginRequest)
	if err := c.Bind(p); err != nil {
		return err
	}
	// email := c.FormValue("email")
	// pwd := c.FormValue("pwd")

	response, _ := myModels.GetUserEmailPwd(p)
	if !response.Id.IsZero() {
		// Create token
		token := jwt.New(jwt.SigningMethodHS256)

		// Set claims
		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = response.Email
		claims["role"] = response.Role
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("2ZPFtPJ5@kDe8m2ud%@eaH?ERaEw?c3"))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, map[string]string{
			"token": t,
		})
	}

	return echo.ErrUnauthorized

	// return c.JSON(http.StatusOK, response)
}
func Auth(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	role := claims["role"].(string)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"name": name,
		"role": role,
	})
}

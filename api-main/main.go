package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	// "github.com/dgrijalva/jwt-go"

	myControllers "./controllers"
)

const (
	jwtKey = "2ZPFtPJ5@kDe8m2ud%@eaH?ERaEw?c3"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://192.168.99.100:81", "http://localhost:81"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/sign", myControllers.Sign)

	// e.Use(middleware.JWT([]byte(jwtKey)))

	auth := e.Group("/auth")
	auth.Use(middleware.JWT([]byte(jwtKey)))
	auth.GET("", myControllers.Auth)

	e.GET("/programs/types", myControllers.GetProgramTypes)
	e.GET("/programs/types/:id", myControllers.GetProgramType)
	// e.POST("/programs/types", myControllers.CreateProgramType)
	// e.PUT("/programs/types/:id", myControllers.UpdateProgramType)
	// e.DELETE("/programs/types/:id", myControllers.DeleteProgramType)

	e.GET("/programs", myControllers.GetPrograms)
	e.GET("/programs/:id", myControllers.GetProgram)
	// e.PUT("/programs/:id", myControllers.UpdateProgram)
	// e.DELETE("/programs/:id", myControllers.DeleteProgram)

	e.GET("/modules", myControllers.GetModules)
	e.GET("/modules/:id", myControllers.GetModule)
	// e.PUT("/modules/:id", myControllers.updateModule)
	// e.DELETE("/modules/:id", myControllers.deleteModule)

	classes := e.Group("/classes")
	classes.Use(middleware.JWT([]byte(jwtKey)))
	classes.GET("", myControllers.GetClasses)
	classes.GET(":id", myControllers.GetClass)
	classes.GET(":id/sessions", myControllers.GetClassSessionsFromId)
	classes.POST(":id/sessions", myControllers.CreateClassSessions)
	// e.GET("/classes/:id/students", myControllers.GetClassIdStudents)
	// e.PUT("/classes/:id", updateClass)
	// e.DELETE("/classes/:id", deleteClass)

	e.GET("/sessions/:id/tutors/:tid/students/:sid", myControllers.GetSessionTutorsFromIdTidSid)
	e.GET("/sessions/:id/tutors/:tid", myControllers.GetSessionTutorsFromIdTid)
	e.GET("/sessions", myControllers.GetSessions)
	e.GET("/sessions/:id", myControllers.GetSession)

	e.GET("/branches", myControllers.GetBranches)
	e.GET("/branches/:id", myControllers.GetBranch)
	// e.PUT("/branches/:id", updateBranch)
	// e.DELETE("/branches/:id", deleteBranch)
	//
	e.GET("/students", myControllers.GetUsers)
	e.GET("/students/:id", myControllers.GetUser)
	// e.PUT("/users/students/:id", updateStudent)
	// e.DELETE("/users/students/:id", deleteStudent)

	e.Logger.Fatal(e.Start(":1323"))
}

type Response struct {
	Message string `json:"id"`
}

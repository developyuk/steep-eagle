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

	a := e.Group("/")
	a.Use(middleware.JWT([]byte(jwtKey)))
	a.GET("auth", myControllers.Auth)

	a.GET("programs/types", myControllers.GetProgramTypes)
	a.GET("programs/types/:id", myControllers.GetProgramType)
	// e.POST("/programs/types", myControllers.CreateProgramType)
	// e.PUT("/programs/types/:id", myControllers.UpdateProgramType)
	// e.DELETE("/programs/types/:id", myControllers.DeleteProgramType)

	a.GET("programs", myControllers.GetPrograms)
	a.GET("programs/:id", myControllers.GetProgram)
	// e.PUT("/programs/:id", myControllers.UpdateProgram)
	// e.DELETE("/programs/:id", myControllers.DeleteProgram)

	a.GET("modules", myControllers.GetModules)
	a.GET("modules/:id", myControllers.GetModule)
	// e.PUT("/modules/:id", myControllers.updateModule)
	// e.DELETE("/modules/:id", myControllers.deleteModule)


	a.GET("classes", myControllers.GetClasses)
	a.GET("classes/:id", myControllers.GetClass)
	a.GET("classes/:id/sessions", myControllers.GetClassSessionsFromId)
	a.POST("classes/:id/sessions", myControllers.CreateClassSessions)
	// e.GET("/classes/:id/students", myControllers.GetClassIdStudents)
	// e.PUT("/classes/:id", updateClass)
	// e.DELETE("/classes/:id", deleteClass)

	a.GET("sessions/:id/tutors/:tid/students/:sid", myControllers.GetSessionTutorsFromIdTidSid)
	a.GET("sessions/:id/tutors/:tid", myControllers.GetSessionTutorsFromIdTid)
	a.GET("sessions", myControllers.GetSessions)
	a.GET("sessions/:id", myControllers.GetSession)

	a.GET("branches", myControllers.GetBranches)
	a.GET("branches/:id", myControllers.GetBranch)
	// e.PUT("/branches/:id", updateBranch)
	// e.DELETE("/branches/:id", deleteBranch)
	//
	a.GET("students", myControllers.GetUsers)
	a.GET("students/:id", myControllers.GetUser)
	// e.PUT("/users/students/:id", updateStudent)
	// e.DELETE("/users/students/:id", deleteStudent)

	e.Logger.Fatal(e.Start(":1323"))
}

type Response struct {
	Message string `json:"id"`
}

package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	// "github.com/dgrijalva/jwt-go"

	myBranches "./branches"
	myClasses "./classes"
	myModules "./modules"
	myPrograms "./programs"
	myProgramsTypes "./programs/types"
	mySessions "./sessions"
	myShared "./shared"
	myUsers "./users"
)

func main() {
	e := echo.New()
	//e.Use(middleware.Logger())
	//e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost", "http://localhost:81", "http://localhost:82", "http://35.187.234.36","http://35.187.234.36:81","http://35.187.234.36:82"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	e.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/sign", myUsers.Sign)

	// e.Use(middleware.JWT([]byte(jwtKey)))

	a := e.Group("/")
	a.Use(middleware.JWT([]byte(myShared.JwtKey)))
	a.Use(myShared.GetAuthMiddleware)

	a.GET("auth", myUsers.Auth)

	a.GET("programs/types", myProgramsTypes.List)
	a.GET("programs/types/:id", myProgramsTypes.Item)
	// e.POST("/programs/types", myControllers.CreateProgramType)
	// e.PUT("/programs/types/:id", myControllers.UpdateProgramType)
	// e.DELETE("/programs/types/:id", myControllers.DeleteProgramType)

	a.GET("programs", myPrograms.List)
	a.GET("programs/:id", myPrograms.Item)
	// e.PUT("/programs/:id", myControllers.UpdateProgram)
	// e.DELETE("/programs/:id", myControllers.DeleteProgram)

	a.GET("modules", myModules.List)
	a.GET("modules/:id", myModules.Item)
	// e.PUT("/modules/:id", myControllers.updateModule)
	// e.DELETE("/modules/:id", myControllers.deleteModule)

	a.GET("classes", myClasses.List)
	a.GET("classes/:id", myClasses.Item)
	a.GET("classes/:id/sessions", mySessions.ListByClassId)
	a.POST("classes/:id/sessions", mySessions.CreateByClassId)
	// e.GET("/classes/:id/students", myControllers.GetClassIdStudents)
	// e.PUT("/classes/:id", updateClass)
	// e.DELETE("/classes/:id", deleteClass)

	a.GET("sessions", mySessions.List)
	a.GET("sessions/:id", mySessions.Item)
  a.GET("sessions/:id/students/:sid", mySessions.ItemTutorStudents)

	a.GET("branches", myBranches.List)
	a.GET("branches/:id", myBranches.Item)
	// e.PUT("/branches/:id", updateBranch)
	// e.DELETE("/branches/:id", deleteBranch)

	a.GET("users", myUsers.List)
	a.GET("users/:id", myUsers.Item)

	a.GET("tutors", myUsers.List)
	a.GET("tutors/:id", myUsers.Item)
  a.GET("tutors/:id/sessions", mySessions.ListByTutorId)

	a.GET("students", myUsers.List)
	a.GET("students/:id", myUsers.Item)
	// e.PUT("/users/students/:id", updateStudent)
	// e.DELETE("/users/students/:id", deleteStudent)

	e.Logger.Fatal(e.Start(":1323"))
}

type Response struct {
	Message string `json:"id"`
}

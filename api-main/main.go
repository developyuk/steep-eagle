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
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{}))

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

	a.GET("programs", myPrograms.List)
	a.GET("programs/:id", myPrograms.Item)

	a.GET("modules", myModules.List)
	a.GET("modules/:id", myModules.Item)

	a.GET("classes", myClasses.List)
	a.GET("classes/:id", myClasses.Item)
	a.GET("classes/:id/sessions", mySessions.ListByClassId)
	a.POST("classes/:id/sessions", mySessions.CreateByClassId)

	a.GET("sessions", mySessions.List)
	a.GET("sessions/:id", mySessions.Item)
  a.GET("sessions/:id/students/:sid", mySessions.ItemStudentsBySessionId)

	a.GET("branches", myBranches.List)
	a.GET("branches/:id", myBranches.Item)

	a.GET("users", myUsers.List)
	a.GET("users/:id", myUsers.Item)

	a.GET("tutors", myUsers.List)
	a.GET("tutors/:id", myUsers.Item)
  a.GET("tutors/:id/sessions", mySessions.ListByTutorId)

	a.GET("students", myUsers.List)
	a.GET("students/:id", myUsers.Item)

	e.Logger.Fatal(e.Start(":1323"))
}

type Response struct {
	Message string `json:"id"`
}

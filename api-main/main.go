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
  mySharedJwt "./shared/jwt"
  mySessionsClasses "./sessions/classes"
  myUsers "./users"
)

func main() {
  e := echo.New()
  //e.Use(middleware.Logger())
  //e.Use(middleware.Recover())
  e.Use(middleware.CORSWithConfig(middleware.CORSConfig{}))

  e.GET("/", func(c echo.Context) error {
    return c.String(http.StatusOK, "Hello, World!")
  })
  e.POST("/sign", myUsers.Sign)

  a := e.Group("/")
  a.Use(middleware.JWT([]byte(mySharedJwt.Key)))
  a.Use(mySharedJwt.GetAuthMiddleware)

  a.GET("auth", myUsers.Auth)

  a.GET("programs/types", myProgramsTypes.List)
  a.GET("programs/types/:id", myProgramsTypes.Item)

  a.GET("programs", myPrograms.List)
  a.GET("programs/:id", myPrograms.Item)

  a.GET("modules", myModules.List)
  a.GET("modules/:id", myModules.Item)

  a.GET("classes", myClasses.List)
  a.GET("classes/group/:by", myClasses.ListGroup)
  a.GET("classes/:id", myClasses.Item)
  a.GET("classes/:id/sessions", mySessionsClasses.ListByClassId)
  a.POST("classes/:id/sessions", mySessionsClasses.CreateByClassId)

  a.GET("sessions", mySessionsClasses.List)
  a.GET("sessions/classes", mySessionsClasses.List)
  a.GET("sessions/:id", mySessionsClasses.Item)
  a.DELETE("sessions/:id", mySessionsClasses.Delete)
  a.GET("sessions/classes/:id", mySessionsClasses.Item)
  //a.GET("sessions/:id/students/:sid", mySessionsClasses.ItemStudentsBySessionId)
  a.POST("sessions/:id/students/:sid", mySessionsClasses.CreateByStudentId)
  a.DELETE("sessions/:id/students/:sid", mySessionsClasses.DeleteByStudentId)
  a.POST("sessions/classes/:id/students/:sid", mySessionsClasses.CreateByStudentId)

  a.GET("branches", myBranches.List)
  a.GET("branches/:id", myBranches.Item)

  a.GET("users", myUsers.List)
  a.GET("users/:id", myUsers.Item)

  a.GET("tutors", myUsers.List)
  a.GET("tutors/:id", myUsers.Item)
  a.GET("tutors/:id/sessions", mySessionsClasses.ListByTutorId)

  a.GET("students", myUsers.List)
  a.GET("students/:id", myUsers.Item)

  e.Logger.Fatal(e.Start(":1323"))
}

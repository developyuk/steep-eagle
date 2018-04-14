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
  mySharedWs "./shared/ws"
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

  websocket := e.Group("/ws")
  hubHome := mySharedWs.NewHub()
  go hubHome.Run()
  websocket.GET("/", func(c echo.Context) error {
    mySharedWs.ClientHome, _ = mySharedWs.ServeWs(hubHome, c.Response(), c.Request())

    return c.String(http.StatusOK, "Hello, World!")
  })

  hubStudents := mySharedWs.NewHub()
  go hubStudents.Run()
  websocket.GET("/students", func(c echo.Context) error {
    mySharedWs.ClientStudents, _ = mySharedWs.ServeWs(hubStudents, c.Response(), c.Request())

    return c.String(http.StatusOK, "Hello, World!")
  })

  authJWT := e.Group("/")
  authJWT.Use(middleware.JWT([]byte(mySharedJwt.Key)))
  authJWT.Use(mySharedJwt.GetAuthMiddleware)

  authJWT.GET("auth", myUsers.Auth)

  authJWT.GET("programs/types", myProgramsTypes.List)
  authJWT.GET("programs/types/:id", myProgramsTypes.Item)

  authJWT.GET("programs", myPrograms.List)
  authJWT.GET("programs/:id", myPrograms.Item)

  authJWT.GET("modules", myModules.List)
  authJWT.GET("modules/:id", myModules.Item)

  authJWT.GET("classes", myClasses.List)
  authJWT.GET("classes/group/:by", myClasses.ListGroup)
  authJWT.GET("classes/:id", myClasses.Item)
  authJWT.GET("classes/:id/sessions", mySessionsClasses.ListByClassId)
  authJWT.POST("classes/:id/sessions", mySessionsClasses.CreateByClassId)

  authJWT.GET("sessions", mySessionsClasses.List)
  authJWT.GET("sessions/classes", mySessionsClasses.List)
  authJWT.GET("sessions/:id", mySessionsClasses.Item)
  authJWT.DELETE("sessions/:id", mySessionsClasses.Delete)
  authJWT.GET("sessions/classes/:id", mySessionsClasses.Item)
  //authJWT.GET("sessions/:id/students/:sid", mySessionsClasses.ItemStudentsBySessionId)
  authJWT.POST("sessions/:id/students/:sid", mySessionsClasses.CreateByStudentId)
  authJWT.DELETE("sessions/:id/students/:sid", mySessionsClasses.DeleteByStudentId)
  authJWT.POST("sessions/classes/:id/students/:sid", mySessionsClasses.CreateByStudentId)

  authJWT.GET("branches", myBranches.List)
  authJWT.GET("branches/:id", myBranches.Item)

  authJWT.GET("users", myUsers.List)
  authJWT.GET("users/:id", myUsers.Item)

  authJWT.GET("tutors", myUsers.List)
  authJWT.GET("tutors/:id", myUsers.Item)
  authJWT.GET("tutors/:id/sessions", mySessionsClasses.ListByTutorId)

  authJWT.GET("students", myUsers.List)
  authJWT.GET("students/:id", myUsers.Item)

  e.Logger.Fatal(e.StartTLS(":1323","letsencrypt/live/mtutor.codingcamp.id/fullchain.pem","letsencrypt/live/mtutor.codingcamp.id/privkey.pem"))
}

package main

import (
  "net/http"

  "github.com/labstack/echo"
  "github.com/labstack/echo/middleware"

  myBranches "./branches"
  myClasses "./classes"
  myModules "./modules"
  myPrograms "./programs"
  myProgramsTypes "./programs/types"
  mySharedJwt "./shared/jwt"
  mySharedMqtt "./shared/mqtt"
  mqtt "github.com/eclipse/paho.mqtt.golang"
  mySessionsClasses "./sessions/classes"
  myUsers "./users"
  "os"
  "log"
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

  opts := mqtt.NewClientOptions().AddBroker("tcp://" + os.Getenv("HOST_MQTT") + ":1883")

  c := mqtt.NewClient(opts)
  if token := c.Connect(); token.Wait() && token.Error() != nil {
    log.Println(token.Error())
  }
  if token := c.Subscribe("schedules", 0, mySharedMqtt.MqttSchedules); token.Wait() && token.Error() != nil {
    log.Println(token.Error())
  }
  if token := c.Subscribe("students", 0, mySharedMqtt.MqttStudents); token.Wait() && token.Error() != nil {
    log.Println(token.Error())
  }

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

  e.Logger.Fatal(e.Start(":1323"))
}

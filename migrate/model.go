package main

import (
  postgrest "../api-main/shared/rest"
  "fmt"
  "github.com/davecgh/go-spew/spew"
)

type (
  Response struct {
    Id uint64 `json:"id"`
  }
  Class struct {
    Day      string `json:"day,omitempty"`
    StartAt  string `json:"start_at,omitempty"`
    FinishAt string `json:"finish_at,omitempty"`
  }
  User struct {
    Email    string `json:"email,omitempty"`
    Name     string `json:"name,omitempty"`
    UserName string `json:"username,omitempty"`
  }
)

func createProgramType(v string) Response {
  var resp Response
  _, err := postgrest.New().GetItem("/program_types").
    SetQuery("name=eq." + v).
    EndStruct(&resp)
  if err != nil {
    postgrest.New().PostItem("/program_types").
      Send(map[string]string{
      "name": v,
    }).
      EndStruct(&resp)
  }
  return resp
}

func createProgram(v string, typeId uint64) Response {
  var resp Response
  _, err := postgrest.New().GetItem("/programs").
    SetQuery("name=eq." + v).
    SetQuery("type_id=eq." + fmt.Sprint(typeId)).
    EndStruct(&resp)

  if err != nil {
    postgrest.New().PostItem("/programs").
      Send(map[string]string{
      "name":    v,
      "type_id": fmt.Sprint(typeId),
    }).
      EndStruct(&resp)
  }
  return resp
}

func createModule(v string) Response {
  var resp Response

  _, err := postgrest.New().GetItem("/modules").
    SetQuery("name=eq." + v).
    EndStruct(&resp)
  if err != nil {
    postgrest.New().PostItem("/modules").
      Send(map[string]string{
      "name": v,
    }).
      EndStruct(&resp)
  }
  return resp
}

func createProgramModule(moduleId uint64, programId uint64) Response {
  var resp Response
  _, err := postgrest.New().GetItem("/programs_modules").
    SetQuery("module_id=eq." + fmt.Sprint(moduleId)).
    SetQuery("program_id=eq." + fmt.Sprint(programId)).
    EndStruct(&resp)
  if err != nil {
    postgrest.New().PostItem("/programs_modules").
      Send(map[string]string{
      "module_id":  fmt.Sprint(moduleId),
      "program_id": fmt.Sprint(programId),
    }).
      EndStruct(&resp)
  }
  return resp
}

func createBranch(v string) Response {
  var resp Response
  _, err := postgrest.New().GetItem("/branches").
    SetQuery("name=eq." + v).
    EndStruct(&resp)
  if err != nil {
    postgrest.New().PostItem("/branches").
      Send(map[string]string{
      "name": v,
    }).
      EndStruct(&resp)
  }
  return resp
}

func createTutor(v *User) Response {
  var tutor Response

  _, err := postgrest.New().GetItem("/users").
    SetQuery(map[string]string{
    "username": "eq." + v.UserName,
    "email":    "eq." + v.Email,
    "role":     "eq." + "tutor",
  }).
    EndStruct(&tutor)
  if err != nil {
    postgrest.New().PostItem("/users").
      Send(map[string]string{
      "username": v.UserName,
      "email":    v.Email,
      "role":     "tutor",
    }).
      EndStruct(&tutor)
  }

  createUserProfile(v.UserName, tutor.Id)

  return tutor
}

func createUserProfile(username string, userId uint64) {
  var resp Response

  _, err := postgrest.New().GetItem("/users_profile").
    SetQuery(map[string]string{
    "name":    "eq." + username,
    "user_id": "eq." + fmt.Sprint(userId),
  }).
    EndStruct(&resp)
  if err != nil {
    postgrest.New().PostItem("/users_profile").
      Send(map[string]string{
      "name":    username,
      "user_id": fmt.Sprint(userId),
    }).
      EndStruct(&resp)
  }
}
func createClass(v *Class, programModuleId uint64, branchId uint64, tutorId uint64) Response {
  var resp Response
  _, err := postgrest.New().GetItem("/classes").
    SetQuery(map[string]string{
    "day":               "eq." + v.Day,
    "start_at":          "eq." + v.StartAt,
    "finish_at":         "eq." + v.FinishAt,
    "program_module_id": "eq." + fmt.Sprint(programModuleId),
    "branch_id":         "eq." + fmt.Sprint(branchId),
    "tutor_id":          "eq." + fmt.Sprint(tutorId),
  }).
    EndStruct(&resp)
  if err != nil {
    postgrest.New().PostItem("/classes").
      Send(map[string]string{
      "day":               v.Day,
      "start_at":          v.StartAt,
      "finish_at":         v.FinishAt,
      "program_module_id": fmt.Sprint(programModuleId),
      "branch_id":         fmt.Sprint(branchId),
      "tutor_id":          fmt.Sprint(tutorId),
    }).
      EndStruct(&resp)

  }
  return resp
}

func createStudent(v *User) Response {
  var student Response

  _, err := postgrest.New().GetItem("/users").
    SetQuery(map[string]string{
    "username": "eq." + v.Name,
    "role":     "eq." + "student",
  }).
    EndStruct(&student)
  if err != nil {
    postgrest.New().PostItem("/users").
      Send(map[string]string{
      "username": v.Name,
      "role":     "student",
    }).
      EndStruct(&student)
  }
  createUserProfile(v.Name, student.Id)
  spew.Dump(v.Name)
  return student
}

func createClassStudent(classId uint64, studentId uint64) Response {
  var resp Response
  _, err := postgrest.New().GetItem("/class_students").
    SetQuery(map[string]string{
    "class_id":   "eq." + fmt.Sprint(classId),
    "student_id": "eq." + fmt.Sprint(studentId),
  }).
    EndStruct(&resp)
  if err != nil {
    postgrest.New().PostItem("/class_students").
      Send(map[string]string{
      "class_id":   fmt.Sprint(classId),
      "student_id": fmt.Sprint(studentId),
    }).
      EndStruct(&resp)
  }
  return resp
}

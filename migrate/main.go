//DB_API=localhost:3000 go run main.go
package main

import (
  "encoding/csv"
  //"encoding/json"
  //"fmt"
  "os"
  "log"

  myRest "../api-main/shared/rest"
  mySharedJwt "../api-main/shared/jwt"
  "strings"
  "fmt"
  "github.com/davecgh/go-spew/spew"
)

type (
  DataCsv struct {
    ProgramType string `json:"program_type"`
    Program     string `json:"program"`
    Module      string `json:"module,omitempty"`
    Branch      string `json:"branch,omitempty"`
    Student     *User  `json:"student,omitempty"`
    Class       *Class `json:"class,omitempty"`
    Tutor       *User  `json:"tutor,omitempty"`
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

  Response struct {
    Id uint64 `json:"id"`
  }
)

const authHeader = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6IiIsImV4cCI6MTUyMTk2NjE0NCwiaWQiOjI0LCJuYW1lIjoidGVzdGVyIiwicGhvdG8iOiIiLCJyb2xlIjoib3BlcmF0aW9uIn0.En__5CcOdU2QiempvmFiwRC5aNE34WoYGiKZ2e2LGXw"

func main() {
  datacsv := parseCSV()
  //dataJson, _ := json.Marshal(datacsv)
  mySharedJwt.AuthHeader = authHeader
  var res Response
  for _, v := range datacsv {
    var programType Response

    _, err := myRest.New().GetItem("/program_types").
      SetQuery(map[string]string{
      "name": "eq." + v.ProgramType,
    }).
      EndStruct(&programType)
    if err != nil {
      myRest.New().PostItem("/program_types").
        Send(map[string]string{
        "name": v.ProgramType,
      }).
        EndStruct(&programType)
    }

    var program Response
    _, err = myRest.New().GetItem("/programs").
      SetQuery(map[string]string{
      "name":    "eq." + v.Program,
      "type_id": "eq." + fmt.Sprint(programType.Id),
    }).
      EndStruct(&program)
    if err != nil {
      myRest.New().PostItem("/programs").
        Send(map[string]string{
        "name":    v.Program,
        "type_id": fmt.Sprint(programType.Id),
      }).
        EndStruct(&program)
    }

    var module Response
    _, err = myRest.New().GetItem("/modules").
      SetQuery(map[string]string{
      "name": "eq." + v.Module,
    }).
      EndStruct(&module)
    if err != nil {
      myRest.New().PostItem("/modules").
        Send(map[string]string{
        "name": v.Module,
      }).
        EndStruct(&module)
    }

    var programModule Response
    _, err = myRest.New().GetItem("/programs_modules").
      SetQuery(map[string]string{
      "module_id":  "eq." + fmt.Sprint(module.Id),
      "program_id": "eq." + fmt.Sprint(program.Id),
    }).
      EndStruct(&programModule)
    if err != nil {
      myRest.New().PostItem("/programs_modules").
        Send(map[string]string{
        "module_id":  fmt.Sprint(module.Id),
        "program_id": fmt.Sprint(program.Id),
      }).
        EndStruct(&programModule)
    }

    var branch Response
    _, err = myRest.New().GetItem("/branches").
      SetQuery(map[string]string{
      "name": "eq." + v.Branch,
    }).
      EndStruct(&branch)
    if err != nil {
      myRest.New().PostItem("/branches").
        Send(map[string]string{
        "name": v.Branch,
      }).
        EndStruct(&branch)
    }

    var tutor Response
    _, err = myRest.New().GetItem("/users").
      SetQuery(map[string]string{
      "username": "eq." + v.Tutor.UserName,
      "email":    "eq." + v.Tutor.Email,
      "role":     "eq." + "tutor",
    }).
      EndStruct(&tutor)
    if err != nil {
      myRest.New().PostItem("/users").
        Send(map[string]string{
        "username": v.Tutor.UserName,
        "email":    v.Tutor.Email,
        "role":     "tutor",
      }).
        EndStruct(&tutor)
    }

    myRest.New().GetItem("/users_profile").
      SetQuery(map[string]string{
      "name":    "eq." + v.Tutor.UserName,
      "user_id": "eq." + fmt.Sprint(tutor.Id),
    }).
      EndStruct(&res)
    if err != nil {
      myRest.New().PostItem("/users_profile").
        Send(map[string]string{
        "name":    v.Tutor.UserName,
        "user_id": fmt.Sprint(tutor.Id),
      }).
        EndStruct(&res)
    }

    var class Response
    _, err = myRest.New().GetItem("/classes").
      SetQuery(map[string]string{
      "day":       "eq." + v.Class.Day,
      "start_at":  "eq." + v.Class.StartAt,
      "finish_at": "eq." + v.Class.FinishAt,
      "program_module_id": "eq." + fmt.Sprint(programModule.Id),
      "branch_id": "eq." + fmt.Sprint(branch.Id),
      "tutor_id":  "eq." + fmt.Sprint(tutor.Id),
    }).
      EndStruct(&class)
    if err != nil {
      myRest.New().PostItem("/classes").
        Send(map[string]string{
        "day":       v.Class.Day,
        "start_at":  v.Class.StartAt,
        "finish_at": v.Class.FinishAt,
        "program_module_id": fmt.Sprint(programModule.Id),
        "branch_id": fmt.Sprint(branch.Id),
        "tutor_id":  fmt.Sprint(tutor.Id),
      }).
        EndStruct(&class)
    }

    var student Response
    _, err = myRest.New().GetItem("/users").
      SetQuery(map[string]string{
      "username": "eq." + v.Student.Name,
      "role":     "eq." + "student",
    }).
      EndStruct(&student)
    if err != nil {
      spew.Dump(v.Student.Name)
      myRest.New().PostItem("/users").
        Send(map[string]string{
        "username": v.Student.Name,
        "role":     "student",
      }).
        EndStruct(&student)
    }

    _, err = myRest.New().GetItem("/class_students").
      SetQuery(map[string]string{
      "class_id":   "eq." + fmt.Sprint(class.Id),
      "student_id": "eq." + fmt.Sprint(student.Id),
    }).
      EndStruct(&res)
    if err != nil {
      myRest.New().PostItem("/class_students").
        Send(map[string]string{
        "class_id":   fmt.Sprint(class.Id),
        "student_id": fmt.Sprint(student.Id),
      }).
        EndStruct(&res)
    }

    _, err = myRest.New().GetItem("/users_profile").
      SetQuery(map[string]string{
      "name":    "eq." + v.Student.Name,
      "user_id": "eq." + fmt.Sprint(student.Id),
    }).
      EndStruct(&res)
    if err != nil {
      _, err = myRest.New().PostItem("/users_profile").
        Send(map[string]string{
        "name":    v.Student.Name,
        "user_id": fmt.Sprint(student.Id),
      }).
        EndStruct(&res)
    }
  }
}

func parseCSV() []DataCsv {
  f, err := os.Open("migration.csv")
  if err != nil {
    log.Println(err)
  }
  defer f.Close()

  lines, err := csv.NewReader(f).ReadAll()
  if err != nil {
    log.Println(err)
  }
  var datacsv []DataCsv
  for i, line := range lines {
    if i == 0 {
      continue
    }
    startAtStr := strings.ToLower(strings.TrimSpace(line[5]))
    startAtArr := strings.Split(startAtStr,":")
    startAt := fmt.Sprintf("%02s:%02s", startAtArr[0],startAtArr[1])
    finishAtStr := strings.ToLower(strings.TrimSpace(line[6]))
    finishAtArr := strings.Split(finishAtStr,":")
    finishAt := fmt.Sprintf("%02s:%02s", finishAtArr[0],finishAtArr[1])
    datacsv = append(datacsv, DataCsv{
      ProgramType: strings.ToLower(strings.TrimSpace(line[0])),
      Program:     strings.ToLower(strings.TrimSpace(line[1])),
      Module:      strings.ToLower(strings.TrimSpace(line[2])),
      Branch:      strings.ToLower(strings.TrimSpace(line[3])),
      Student: &User{
        Name: strings.ToLower(strings.TrimSpace(line[7])),
      },
      Class: &Class{
        Day:      strings.ToLower(strings.TrimSpace(line[4])),
        StartAt:  startAt,
        FinishAt: finishAt,
      },
      Tutor: &User{
        Email:    strings.ToLower(strings.TrimSpace(line[8])),
        Name:     strings.ToLower(strings.TrimSpace(line[10])),
        UserName: strings.ToLower(strings.TrimSpace(line[9])),
      },
    })
  }
  return datacsv
}

//DB_API=localhost:3000 go run *.go
package main

import (
  mySharedJwt "../api-main/shared/jwt"
)


const authHeader = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3RlcjNAZ21haWwuY29tIiwiZXhwIjoxNTI2MTQ2NjQzLCJpZCI6MjQsIm5hbWUiOiJ0ZXN0ZXIzIiwicGhvdG8iOiIiLCJyb2xlIjoib3BlcmF0aW9uIiwidXNlcm5hbWUiOiJ0ZXN0ZXIzIn0.Wg7TNyfhsULZOG13GenJS9lxycCKzIdKLuvNXJ71D9k"

func main() {
  datacsv := ParseCSV()
  //dataJson, _ := json.Marshal(datacsv)
  mySharedJwt.AuthHeader = authHeader

  for _, v := range datacsv {
    programType := createProgramType(v.ProgramType)
    program := createProgram(v.Program, programType.Id)
    module := createModule(v.Module)
    programModule := createProgramModule(module.Id, program.Id)
    branch := createBranch(v.Branch)
    tutor := createTutor(v.Tutor)
    class := createClass(v.Class, programModule.Id, branch.Id, tutor.Id)
    student := createStudent(v.Student)
    _ = createClassStudent(class.Id, student.Id)
  }
}

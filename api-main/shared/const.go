package shared

import (
  "os"
)

const (
  PathPrograms      string = "/programs"
  PathProgramsTypes string = "/programs/types"
  PathModules       string = "/modules"
  PathClasses       string = "/classes"
  PathUsers         string = "/users"
  PathStudents      string = "/students"
  PathTutors        string = "/tutors"
  PathSessions      string = "/sessions"
  PathBranches      string = "/branches"
)

var JwtKey = os.Getenv("JWT_KEY")

const (
  TimeZone = "Asia/Jakarta"
)

type (
  ClassStudents struct {
    Id        uint64 `json:"id"`
    ClassId   uint64 `json:"class_id"`
    StudentId uint64 `json:"student_id"`
  }
)

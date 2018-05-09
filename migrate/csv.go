package main

import (
  "os"
  "encoding/csv"
  "fmt"
  "log"
  "strings"
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
)

func ParseCSV() []DataCsv {
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
    startAtStr := validateString(line[5])
    startAtArr := strings.Split(startAtStr, ":")
    startAt := fmt.Sprintf("%02s:%02s", startAtArr[0], startAtArr[1])
    finishAtStr := validateString(line[6])
    finishAtArr := strings.Split(finishAtStr, ":")
    finishAt := fmt.Sprintf("%02s:%02s", finishAtArr[0], finishAtArr[1])
    datacsv = append(datacsv, DataCsv{
      ProgramType: validateString(line[0]),
      Program:     validateString(line[1]),
      Module:      validateString(line[2]),
      Branch:      validateString(line[3]),
      Student: &User{
        Name: validateString(line[7]),
      },
      Class: &Class{
        Day:      validateString(line[4]),
        StartAt:  startAt,
        FinishAt: finishAt,
      },
      Tutor: &User{
        Email:    validateString(line[8]),
        Name:     validateString(line[10]),
        UserName: validateString(line[9]),
      },
    })
  }
  return datacsv
}

func validateString(v string) string {
  return strings.ToLower(strings.TrimSpace(v))
}

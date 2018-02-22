package classes

import (
  myShared "../shared"
  "github.com/jmoiron/sqlx"
  "log"
)

type Class_ struct {
  myShared.Hal
  Id       int         `json:"id"`
  Name     string      `json:"name"`
  Image    string      `json:"image"`
  Day      string      `json:"day"`
  Time     string      `json:"time"`
  Ts       string      `json:"ts"`
  ModuleId interface{} `json:"module_id" db:"module_id"`
  BranchId interface{} `json:"branch_id" db:"branch_id"`
}

func getSessionsById(db *sqlx.DB, id int) []int {

  var data []int
  err := db.Select(&data, `SELECT id
			          FROM sessions
			          where class_id = $1`, id)
  if err != nil {
    log.Fatal(err)
  }

  return data
}

// func GetClassIdStudents(id string) []Href {
// 	db := Connect()
// 	i, _ := strconv.ParseInt(id, 0, 0)
// 	return GetStudentFromId(db, int(i))
// }

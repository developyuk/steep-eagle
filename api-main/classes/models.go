package classes

import (
  myShared "../shared"
  "github.com/jmoiron/sqlx"
  "log"
  "fmt"
  "strings"
  // "time"
)

type Class_ struct {
  myShared.Hal
  Id       int         `json:"id"`
  Name     string      `json:"name"`
  Image    *string     `json:"image"`
  Day      string      `json:"day"`
  Time     string      `json:"time"`
  ModuleId interface{} `json:"-" db:"module_id"`
  BranchId interface{} `json:"-" db:"branch_id"`
}

func getStudentsById(db *sqlx.DB, id int) []int {

  var data []int
  err := db.Select(&data, `SELECT student_id
			          FROM class_students
			          where class_id = $1`, id)
  if err != nil {
    log.Fatal(err)
  }

  return data
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

func ListData(params map[string]interface{}) ([]Class_, *sqlx.DB) {
  db := myShared.Connect()

  var data []Class_
  sql := []string{`SELECT d.id,
       d.name,
       d.image,
       d.day,
       d.time,
       d.module_id,
       d.branch_id
FROM
  (SELECT c.*,
          to_timestamp(cast(cast(a.series as date) as text) || ' ' ||c.time, 'YYYY-MM-DD HH24.MI') ts
   FROM
     ( SELECT g.series,
              to_char(g.series,'FMday') dow
      FROM
        ( SELECT GENERATE_SERIES( NOW() AT TIME ZONE 'Asia/Jakarta', NOW() AT TIME ZONE 'Asia/Jakarta'+'7 DAYS','1 DAY' ) ) g(series) ) a
   JOIN classes c ON c.day = a.dow) d
  JOIN branches b ON b.id = d.branch_id
WHERE ts > (now() - interval '1 hour')
  AND ts < (now() - interval '1 hour' + interval '7 days') `}

  if _, ok := params["q"]; ok {
    params["q"] = fmt.Sprintf("%%%v%%", params["q"])
    sql = append(sql, `AND b.name LIKE :q`)
  }

  if _, ok := params["sort"]; ok {
    sql = append(sql, "ORDER BY")
    if params["sort"] == "datetime" {
      sql = append(sql, "d.ts ASC")
    }
    if params["sort"] == "-datetime" {
      sql = append(sql, "d.ts DESC")
    }
  }
  log.Println(strings.Join(sql, " "), params)
  if len(params) > 0 {
    stmt, _ := db.PrepareNamed(strings.Join(sql, " "))
    _ = stmt.Select(&data, params)
    //_ = db.Select(&data, strings.Join(sql, " "), params[:])
  } else {
    _ = db.Select(&data, strings.Join(sql, " "))
  }

  return data, db
}

func ItemData(id string) (Class_, *sqlx.DB) {
  db := myShared.Connect()

  var data Class_
  _ = db.Get(&data, `SELECT id, name, image, day, time, module_id, branch_id
    FROM classes
    WHERE id = $1`, id)

  return data, db
}

// func GetClassIdStudents(id string) []Href {
// 	db := Connect()
// 	i, _ := strconv.ParseInt(id, 0, 0)
// 	return GetStudentFromId(db, int(i))
// }

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
  Ts       string      `json:"ts"`
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
  var sqlParams []interface{}
  sqlParams = append(sqlParams, params["authId"])

  sql := []string{fmt.Sprintf(`SELECT d.id
	,d.day
	,d.TIME
	,d.ts
	,d.module_id
	,d.branch_id
FROM (
	SELECT c.*
		,(a.dw::DATE::TEXT || ' ' || c.TIME || ':00')::TIMESTAMP AT TIME ZONE 'Asia/Jakarta' ts
	FROM (
		SELECT ts AT TIME ZONE 'Asia/Jakarta' dw
		FROM GENERATE_SERIES(NOW(), NOW() AT TIME ZONE 'Asia/Jakarta' + '7 DAYS', '1 DAY') g(ts)
		) a
	INNER JOIN classes c ON c.day = to_char(dw, 'FMday')
	) d
WHERE d.ts > (NOW() - INTERVAL '1 hour')
	AND d.ts < (NOW() - INTERVAL '1 hour' + INTERVAL '7 days')
	AND d.id NOT IN (
		SELECT s.class_id
		FROM sessions s
		WHERE s.created_at::DATE >= CURRENT_DATE
			AND tutor_id = $%v
		) `,len(sqlParams))}

  if _, ok := params["q"]; ok {
    //params["q"] = fmt.Sprintf("%%%v%%", params["q"])
    sqlParams = append(sqlParams, fmt.Sprintf("%%%v%%", params["q"]))
    sql = append(sql, fmt.Sprintf(`AND b.name LIKE $%v`, len(sqlParams)))
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
  //log.Println(strings.Join(sql, " "), params, sqlParams)

  _ = db.Select(&data, strings.Join(sql, " "), sqlParams...)

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

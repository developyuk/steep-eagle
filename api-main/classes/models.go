package classes

import (
	myShared "../shared"
	"github.com/jmoiron/sqlx"
	"log"
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
	sql := []string{`SELECT id, name, image, day, time,
      module_id, branch_id
	  FROM (
      SELECT g.series, to_char(g.series,'FMday') dow
      FROM (
	      SELECT GENERATE_SERIES(
          NOW() AT TIME ZONE 'Asia/Jakarta', NOW() AT TIME ZONE 'Asia/Jakarta'+'6 DAYS','1 DAY'
        )
	    ) g(series)
	  ) a
    JOIN classes c ON c.day = a.dow
    ORDER BY a.series ASC`}

	//if len(params) > 0 {
	//	sql = append(sql, "WHERE")
	//
	//	if _, ok := params["day"]; ok {
	//		sql = append(sql, "day = :day")
	//	}
	//}

	// log.Println(strings.Join(sql, " "), params)
	//stmt, _ := db.PrepareNamed(strings.Join(sql, " "))
	//_ = stmt.Select(&data, params)
	_ = db.Select(&data,strings.Join(sql, " "))

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

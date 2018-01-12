package models

import (
	"github.com/jmoiron/sqlx"
	"log"
	"strings"
	// "time"
)

type Class_ struct {
	Hal
	Id       int         `json:"id"`
	Name     string      `json:"name"`
	Image    *string     `json:"image"`
	Day      string      `json:"day"`
	Time     string      `json:"time"`
	ModuleId interface{} `json:"-" db:"module_id"`
	BranchId interface{} `json:"-" db:"branch_id"`
}

func GetStudentFromId(db *sqlx.DB, id int) []int {

	var data []int
	err := db.Select(&data, `SELECT student_id
			          FROM class_students
			          where class_id = $1`, id)
	if err != nil {
		log.Fatal(err)
	}

	return data
}

func GetSessionFromId(db *sqlx.DB, id int) []int {

	var data []int
	err := db.Select(&data, `SELECT id
			          FROM sessions
			          where class_id = $1`, id)
	if err != nil {
		log.Fatal(err)
	}

	return data
}

func GetClasses(params map[string]interface{}) ([]Class_, *sqlx.DB) {
	db := Connect()

	var data []Class_
	sql := []string{`SELECT id, name, image, day, time,
      module_id, branch_id
    FROM classes`}

	if len(params) > 0 {
		sql = append(sql, "WHERE")

		if _, ok := params["day"]; ok {
			sql = append(sql, "day = :day")
		}
	}

	stmt, _ := db.PrepareNamed(strings.Join(sql, " "))
	_ = stmt.Select(&data, params)

	return data, db
}

func GetClass(id string) (Class_, *sqlx.DB) {
	db := Connect()

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

package users

import (
	myShared "../shared"
	"github.com/jmoiron/sqlx"
	"gopkg.in/guregu/null.v3"
	"log"
	"strings"
)

type User struct {
	myShared.Hal
	Id        null.Int    `json:"id"`
	Name      null.String `json:"name"`
	FirstName null.String `json:"first_name"  db:"first_name"`
	LastName  null.String `json:"last_name"  db:"last_name"`
	Email     null.String `json:"email"`
	Dob       null.Time   `json:"dob"`
	Photo     null.String `json:"photo"`
	Role      null.String `json:"role"`
}

func getClassById(db *sqlx.DB, id null.Int) int64 {
	var data int64
	err := db.Get(&data, `SELECT class_id
  			          FROM class_students
  			          where id = $1`, id)
	if err != nil {
		log.Fatal(err)
	}

	return data
}
func ListData(params map[string]interface{}) ([]User, *sqlx.DB) {
	db := myShared.Connect()
	sql := []string{`SELECT id, name, first_name, last_name, email,
      photo, role
    FROM users`}

	if len(params) > 0 {
		sql = append(sql, "WHERE")

		if _, ok := params["role"]; ok {
			sql = append(sql, "role = :role")
		}
	}
	var data []User
	// log.Println(strings.Join(sql, " "))
	stmt, _ := db.PrepareNamed(strings.Join(sql, " "))
	err := stmt.Select(&data, params)
	if err != nil {
		log.Fatal(err)
	}

	return data, db
}

func ItemData(params map[string]interface{}) (User, *sqlx.DB) {
	db := myShared.Connect()
	sql := []string{`SELECT id, name, first_name, last_name, email, dob,
     photo, role
    FROM users`}
	if len(params) > 0 {
		sql = append(sql, "WHERE")

		if _, ok := params["id"]; ok {
			sql = append(sql, "id = :id")
		}

		if _, ok := params["role"]; ok {
			sql = append(sql, "AND")
			sql = append(sql, "role = :role")
		}
	}
	var data User
	stmt, _ := db.PrepareNamed(strings.Join(sql, " "))
	err := stmt.Get(&data, params)
	if err != nil {
		log.Fatal(err)
	}

	return data, db
}

func UserByEmailPwdData(param *UserLoginRequest) (User, *sqlx.DB) {
	db := myShared.Connect()

	var data User
	err := db.Get(&data, `SELECT id, name, first_name, last_name, email, dob,
     photo, role
    FROM users
    WHERE email = $1 AND pass = $2`, param.Email, param.Pwd)
	if err != nil {
		// log.Fatal(err)
		return User{}, db
	}

	return data, db

}

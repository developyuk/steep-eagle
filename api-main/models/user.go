package models

import (
	"github.com/jmoiron/sqlx"
	"gopkg.in/guregu/null.v3"
	"log"
	// "time"
)

type User struct {
	Hal
	Id        null.Int    `json:"id"`
	Name      null.String `json:"name"`
	FirstName null.String `json:"first_name"  db:"first_name"`
	LastName  null.String `json:"last_name"  db:"last_name"`
	Email     null.String `json:"email"`
	Dob       null.Time   `json:"dob"`
	Photo     null.String `json:"photo"`
	Role      null.String `json:"role"`
}

func GetUsers() ([]User, *sqlx.DB) {
	db, err := Connect()

	var data []User
	err = db.Select(&data, `SELECT id, name, first_name,
    last_name, email, photo, role
    FROM users`)
	if err != nil {
		log.Fatal(err)
	}

	return data, db
}

func GetUser(id string) (User, *sqlx.DB) {
	db, err := Connect()

	var data User
	err = db.Get(&data, `SELECT id, name, first_name, last_name, email, dob,
     photo, role
    FROM users
    WHERE id = $1`, id)
	if err != nil {
		log.Fatal(err)
	}

	return data, db
}

func GetUserEmailPwd(param *UserLoginRequest) (User, *sqlx.DB) {
	db, err := Connect()

	var data User
	err = db.Get(&data, `SELECT id, name, first_name, last_name, email, dob,
     photo, role
    FROM users
    WHERE email = $1 AND pwd = $2`, param.Email, param.Pwd)
	if err != nil {
		// log.Fatal(err)
		return User{}, db
	}

	return data, db

}

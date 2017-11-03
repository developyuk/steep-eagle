package models

import (
	"github.com/jmoiron/sqlx"
	"log"
	// "time"
)

type User struct {
  Hal
	Id        int `json:"id"`
	Name      string `json:"name"`
	FirstName string `json:"first_name"  db:"first_name"`
	LastName  string `json:"last_name"  db:"last_name"`
	Email     string `json:"email"`
	// Dob       time.Time `json:"dob"`
	Photo     string `json:"photo"`
	Role      string `json:"role"`
}

func GetUsers() ([]User, *sqlx.DB) {
	db, err := Connect()

	var data []User
	err = db.Select(&data, `SELECT id, name, coalesce(first_name,'') first_name,
    coalesce(last_name,'') last_name, coalesce(email,'') email,
     coalesce(photo,'') photo, role
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

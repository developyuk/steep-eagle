package models

import (
	"github.com/jmoiron/sqlx"
	"log"
)

type Program struct {
	Hal
	Id     int         `json:"id"`
	Name   string      `json:"name"`
	TypeId interface{} `json:"-" db:"type_id"`
}

func GetModulesFromId(db *sqlx.DB, id int) []int {

	var data []int
	err := db.Select(&data, `SELECT module_id
					FROM program_modules
					where program_id = $1`, id)

	if err != nil {
		log.Fatal(err)
	}

	return data
}

func GetPrograms() ([]Program, *sqlx.DB) {
	db, err := Connect()

	var data []Program
	err = db.Select(&data, `SELECT id,name,type_id
    FROM programs`)
	if err != nil {
		log.Fatal(err)
	}

	return data, db
}

func GetProgram(id string) (Program, *sqlx.DB) {
	db, err := Connect()

	var data Program
	err = db.Get(&data, `SELECT id,name,type_id
	  FROM programs
	  WHERE id = $1`, id)
	if err != nil {
		log.Fatal(err)
	}

	return data, db
}

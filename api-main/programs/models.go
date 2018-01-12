package programs

import (
	myShared "../shared"
	"github.com/jmoiron/sqlx"
	"log"
)

type Program struct {
	myShared.Hal
	Id     int         `json:"id"`
	Name   string      `json:"name"`
	TypeId interface{} `json:"-" db:"type_id"`
}

func ItemModulesHrefData(db *sqlx.DB, id int) []int {

	var data []int
	err := db.Select(&data, `SELECT module_id
					FROM program_modules
					WHERE program_id = $1`, id)

	if err != nil {
		log.Fatal(err)
	}

	return data
}

func ListData() ([]Program, *sqlx.DB) {
	db := myShared.Connect()

	var data []Program
	err := db.Select(&data, `SELECT id, name, type_id
    FROM programs`)
	if err != nil {
		log.Fatal(err)
	}

	return data, db
}

func ItemData(id string) (Program, *sqlx.DB) {
	db := myShared.Connect()

	var data Program
	err := db.Get(&data, `SELECT id, name, type_id
	  FROM programs
	  WHERE id = $1`, id)
	if err != nil {
		log.Fatal(err)
	}

	return data, db
}

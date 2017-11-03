package models

import (
	"log"
)

type ProgramType struct {
  Hal
	Id   int    `json:"id"`
	Name string `json:"name"`
}


func GetProgramTypes() []ProgramType {
	db, err := Connect()

	var data []ProgramType
	err = db.Select(&data, "SELECT id,name FROM program_types")
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func GetProgramType(id string) ProgramType {
	db, err := Connect()

	var data ProgramType
	err = db.Get(&data, `SELECT id,name
    FROM program_types
    WHERE id = $1`, id)

	if err != nil {
		log.Fatal(err)
	}

	return data
}

func CreateProgramType(name string) Response {
	db, err := Connect()

	db.MustExec(`INSERT INTO program_types(name)
		VALUES ($1)`, name)

	if err != nil {
		log.Fatal(err)
	}
	return Response{Message: ""}
}

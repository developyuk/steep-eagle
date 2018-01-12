package types

import (
	myShared "../../shared"
	"log"
)

type ProgramType struct {
	myShared.Hal
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func ListData() []ProgramType {
	db := myShared.Connect()

	var data []ProgramType
	err := db.Select(&data, "SELECT id, name FROM program_types")
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func ItemData(id string) ProgramType {
	db := myShared.Connect()

	var data ProgramType
	err := db.Get(&data, `SELECT id, name
    FROM program_types
    WHERE id = $1`, id)

	if err != nil {
		log.Fatal(err)
	}

	return data
}

func CreateProgramType(name string) myShared.Response {
	db := myShared.Connect()

	db.MustExec(`INSERT INTO program_types(name)
		VALUES ($1)`, name)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	return myShared.Response{Message: ""}
}

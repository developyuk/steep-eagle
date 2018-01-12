package modules

import (
	myShared "../shared"
	"gopkg.in/guregu/null.v3"
	"log"
)

type Module struct {
	myShared.Hal
	Id           int         `json:"id"`
	Name         string      `json:"name"`
	Image        null.String `json:"image"`
	SessionTotal int         `json:"session_total" db:"session_total"`
}

func ListData() []Module {
	db := myShared.Connect()

	var data []Module
	err := db.Select(&data, `SELECT id, name, image, session_total
    FROM modules`)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func ItemData(id string) Module {
	db := myShared.Connect()

	var data Module
	err := db.Get(&data, `SELECT id, name, image, session_total
    FROM modules
    WHERE id = $1`, id)
	if err != nil {
		log.Fatal(err)
	}

	return data
}

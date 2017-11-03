package models

import (
	"log"
)

type Module struct {
  Hal
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	Image        *string `json:"image"`
	SessionTotal int     `json:"session_total" db:"session_total"`
}


func GetModules() []Module {
	db, err := Connect()

	var data []Module
	err = db.Select(&data, `SELECT id,name,image,session_total
    FROM modules`)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func GetModule(id string) Module {
	db, err := Connect()

	var data Module
	err = db.Get(&data, `SELECT id,name,image,session_total
    FROM modules
    WHERE id = $1`, id)
	if err != nil {
		log.Fatal(err)
	}

	return data
}

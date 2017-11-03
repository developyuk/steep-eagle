package models

import (
	"log"
)

type Branch struct {
  Hal
	Id      int     `json:"id"`
	Name    string  `json:"name"`
	Image   *string `json:"image"`
	Address *string `json:"address"`
}


func GetBranches() []Branch {
	db, err := Connect()

	var data []Branch
	err = db.Select(&data, "SELECT id,name,image,address FROM branches")

	if err != nil {
		log.Fatal(err)
	}

	return data
}

func GetBranch(id string) Branch {
	db, err := Connect()

	var data Branch
	err = db.Get(`SELECT id,name,image,address
		FROM branches
		WHERE id = $1`, id)
	if err != nil {
		log.Fatal(err)
	}

	return data
}

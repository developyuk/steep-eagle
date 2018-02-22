package branches

import (
	myShared "../shared"
	"github.com/jmoiron/sqlx"
	"gopkg.in/guregu/null.v3"
	"log"
)

type Branch struct {
	myShared.Hal
	Id      int         `json:"id"`
	Name    string      `json:"name"`
	Image   null.String `json:"image"`
	Address null.String `json:"address"`
}

func getClassesById(db *sqlx.DB, id int) []int {
	var data []int
	err := db.Select(&data, `SELECT id
  			          FROM classes
  			          where branch_id = $1`, id)
	if err != nil {
		log.Fatal(err)
	}

	return data
}

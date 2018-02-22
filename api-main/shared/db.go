package shared

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

func Connect() *sqlx.DB {
	//log.Println(CurrentAuth)
	conn := "postgres://postgres:postgres@db-main/postgres?sslmode=disable"

	db, err := sqlx.Connect("postgres", conn)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

package shared

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

const (
	OPERATION_CONN string = "postgres://operation:mysecretpassword@db-main/postgres?sslmode=disable"
	TUTOR_CONN     string = "postgres://tutor:mysecretpassword@db-main/postgres?sslmode=disable"
	STUDENT_CONN   string = "postgres://student:mysecretpassword@db-main/postgres?sslmode=disable"
	PARENT_CONN    string = "postgres://parent:mysecretpassword@db-main/postgres?sslmode=disable"
	PARTNER_CONN   string = "postgres://parner:mysecretpassword@db-main/postgres?sslmode=disable"
)

func Connect() *sqlx.DB {
	db, err := sqlx.Connect("postgres", OPERATION_CONN)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

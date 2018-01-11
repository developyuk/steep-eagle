package models

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

const (
	OPERATION_CONN = "postgres://operation:mysecretpassword@db-main/postgres?sslmode=disable"
	TUTOR_CONN     = "postgres://tutor:mysecretpassword@db-main/postgres?sslmode=disable"
	STUDENT_CONN   = "postgres://student:mysecretpassword@db-main/postgres?sslmode=disable"
	PARENT_CONN    = "postgres://parent:mysecretpassword@db-main/postgres?sslmode=disable"
	PARTNER_CONN   = "postgres://parner:mysecretpassword@db-main/postgres?sslmode=disable"
)

type Hal struct {
	Links    interface{} `json:"_links,omitempty"`
	Embedded interface{} `json:"_embedded,omitempty"`
	Total    int         `json:"total,omitempty"`
	Count    int         `json:"count,omitempty"`
}

//
// type Href struct {
// 	Href string `json:"href"`
// }
//
// type LinksSelf struct {
// 	Self Href `json:"self"`
// }

type Response struct {
	Message string `json:"message"`
	Id      int64 `json:"id,omitempty"`
}

type UserLoginRequest struct {
	Email string `json:"email" xml:"email" form:"email" query:"email"`
	Pwd   string `json:"pwd" xml:"pwd" form:"pwd" query:"email"`
}

func Connect() (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", OPERATION_CONN)
	if err != nil {
		log.Fatal(err)
	}
	return db, err
}

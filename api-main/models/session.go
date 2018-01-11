package models

import (
	"github.com/jmoiron/sqlx"
	"gopkg.in/guregu/null.v3"
	"log"
	"time"
)

type Session struct {
	Hal
	Id        int       `json:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	ClassId   int       `json:"-" db:"class_id"`
}

type SessionTutor struct {
	Hal
	Id      int `json:"-"`
	TutorId int `json:"-" db:"tutor_id"`
}

type SessionStudentRating struct {
	Interaction null.Int `json:"interaction"`
	Cognition   null.Int `json:"cognition"`
	Creativity  null.Int `json:"creativity"`
}
type SessionStudent struct {
	Hal
	Id        null.Int             `json:"id"`
	CreatedAt null.Time            `json:"created_at" db:"created_at"`
	Status    null.Bool            `json:"status"`
	Feedback  null.String          `json:"feedback"`
	Rating    SessionStudentRating `json:"rating"`
	// SessionsTutorsId string `json:"-" db:"sessions_tutors_id"`
	StudentId int `json:"-" db:"student_id"`
}

func GetSessionTutorsFromId(db *sqlx.DB, id int) []SessionTutor {

	var data []SessionTutor
	err := db.Select(&data, `SELECT id,tutor_id
			          FROM sessions_tutors
			          WHERE session_id = $1`, id)
	if err != nil {
		log.Fatal(err)
	}

	return data
}

func GetSessionStudentsFromId(db *sqlx.DB, id int, tid int) []SessionStudent {

	var data []SessionStudent
	err := db.Select(&data, ` SELECT ss.id,
      ss.created_at, ss.status, ss.feedback,
      cs.student_id, ss.rating_interaction "rating.interaction",
      ss.rating_cognition "rating.cognition",
      ss.rating_creativity "rating.creativity"
    FROM class_students cs
    LEFT JOIN sessions s ON s.class_id = cs.class_id AND s.id = $1
    LEFT JOIN sessions_tutors t ON s.id = t.session_id  AND t.tutor_id = $2
    LEFT JOIN sessions_students ss ON ss.sessions_tutors_id = t.id AND cs.student_id = ss.student_id
    WHERE tutor_id IS NOT NULL`, id, tid)
	if err != nil {
		log.Fatal(err)
	}

	return data
}

func GetSessionStudentFromId(db *sqlx.DB, id int, tid int, sid int) SessionStudent {

	var data SessionStudent
	err := db.Get(&data, ` SELECT ss.id id, ss.created_at, ss.status, ss.feedback,
      cs.student_id, ss.rating_interaction "rating.interaction",
      ss.rating_cognition "rating.cognition",
      ss.rating_creativity "rating.creativity"
FROM class_students cs
LEFT JOIN sessions s ON s.class_id = cs.class_id AND s.id = $1
LEFT JOIN sessions_tutors t ON s.id = t.session_id  AND t.tutor_id = $2
LEFT JOIN sessions_students ss ON ss.sessions_tutors_id = t.id AND cs.student_id = ss.student_id
WHERE tutor_id IS NOT NULL AND cs.student_id = $3`, id, tid, sid)
	if err != nil {
		log.Fatal(err)
	}

	return data
}
func GetSessions() ([]Session, *sqlx.DB) {
	db, err := Connect()

	var data []Session
	err = db.Select(&data, `SELECT id, created_at, class_id
    FROM sessions`)
	if err != nil {
		log.Fatal(err)
	}

	return data, db
}
func GetSession(id string) (Session, *sqlx.DB) {
	db, err := Connect()

	var data Session
	err = db.Get(&data, `SELECT id, created_at, class_id
    FROM sessions
    WHERE id = $1`, id)
	if err != nil {
		log.Fatal(err)
	}

	return data, db
}
func CreateSessionClass(id string) Response {
	db, err := Connect()
	if err != nil {
		log.Fatal(err)
	}
	var lastInsertId int64
	db.QueryRowx(`INSERT INTO sessions(class_id)
  		VALUES ($1) RETURNING id`, id).Scan(&lastInsertId)

	return Response{Message: "", Id: lastInsertId}
}

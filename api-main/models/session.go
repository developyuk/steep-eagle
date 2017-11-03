package models

import (
	"github.com/jmoiron/sqlx"
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
	Interaction int `json:"interaction"`
	Cognition   int `json:"cognition"`
	Creativity  int `json:"creativity"`
}
type SessionStudent struct {
	Hal
	Id        int                  `json:"id"`
	CreatedAt time.Time            `json:"created_at" db:"created_at"`
	Status    bool                 `json:"status"`
	Feedback  string               `json:"feedback"`
	Rating    SessionStudentRating `json:"rating"`
	// SessionsTutorsId string `json:"-" db:"sessions_tutors_id"`
	StudentId int `json:"-" db:"student_id"`
}

func GetSessionTutorsFromId(db *sqlx.DB, id int) []SessionTutor {

	var data []SessionTutor
	err := db.Select(&data, `SELECT id,tutor_id
			          FROM sessions_tutors
			          where session_id = $1`, id)
	if err != nil {
		log.Fatal(err)
	}

	return data
}

func GetSessionStudentsFromId(db *sqlx.DB, id int, tid int) []SessionStudent {

	var data []SessionStudent
	err := db.Select(&data, ` select coalesce(ss.id,0) id,
      coalesce(ss.created_at,'0001-01-01') created_at,
      coalesce(ss.status,false) status, coalesce(ss.feedback,'') feedback,
      cs.student_id, coalesce(ss.rating_interaction,0) "rating.interaction",
      coalesce(ss.rating_cognition,0) "rating.cognition",
      coalesce(ss.rating_creativity,0) "rating.creativity"
    from class_students cs
    left join sessions s on s.class_id = cs.class_id and s.id = $1
    left JOIN sessions_tutors t ON s.id = t.session_id  and t.tutor_id = $2
    left JOIN sessions_students ss on ss.sessions_tutors_id = t.id and cs.student_id = ss.student_id
    where tutor_id is not null`, id, tid)
	if err != nil {
		log.Fatal(err)
	}

	return data
}

func GetSessionStudentFromId(db *sqlx.DB, id int, tid int, sid int) SessionStudent {

	var data SessionStudent
	err := db.Get(&data, `
select coalesce(ss.id,0) id, coalesce(ss.created_at,'0001-01-01') created_at, coalesce(ss.status,false) status, coalesce(ss.feedback,'') feedback,
      cs.student_id, coalesce(ss.rating_interaction,0) "rating.interaction",
      coalesce(ss.rating_cognition,0) "rating.cognition",
      coalesce(ss.rating_creativity,0) "rating.creativity"
from class_students cs
left join sessions s on s.class_id = cs.class_id and s.id = $1
left JOIN sessions_tutors t ON s.id = t.session_id  and t.tutor_id = $2
left JOIN sessions_students ss on ss.sessions_tutors_id = t.id and cs.student_id = ss.student_id
where tutor_id is not null and cs.student_id = $3`, id, tid, sid)
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

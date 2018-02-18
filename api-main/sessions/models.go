package sessions

import (
  myShared "../shared"
  "github.com/jmoiron/sqlx"
  "gopkg.in/guregu/null.v3"
  "log"
  "strings"
  "time"
)

type Session struct {
  myShared.Hal
  Id        int       `json:"id"`
  CreatedAt time.Time `json:"created_at" db:"created_at"`
  ClassId   int       `json:"-" db:"class_id"`
  TutorId   int       `json:"-" db:"tutor_id"`
}

// type SessionTutor struct {
// 	myShared.Hal
// 	Id      int `json:"-"`
// 	TutorId int `json:"-" db:"tutor_id"`
// }

type SessionStudentRating struct {
  Interaction null.Int `json:"interaction"`
  Cognition   null.Int `json:"cognition"`
  Creativity  null.Int `json:"creativity"`
}
type SessionStudent struct {
  myShared.Hal
  Id        null.Int             `json:"id"`
  CreatedAt null.Time            `json:"created_at" db:"created_at"`
  Status    null.Bool            `json:"status"`
  Feedback  null.String          `json:"feedback"`
  Rating    SessionStudentRating `json:"rating"`
  // SessionsTutorsId string `json:"-" db:"sessions_tutors_id"`
  StudentId int `json:"-" db:"student_id"`
}

func getTutorsById(db *sqlx.DB, id int) []int64 {
  var data []int64
  err := db.Select(&data, `SELECT tutor_id
			          FROM sessions_tutors
			          WHERE session_id = $1`, id)
  if err != nil {
    log.Fatal(err)
  }

  return data
}

func getStudentsById(db *sqlx.DB, params map[string]interface{}) []int64 {
  var sqlParams []interface{}
  sqlParams = append(sqlParams,params["id"])

  sql := []string{`SELECT student_id
     FROM sessions s
     JOIN class_students c ON s.class_id = c.class_id
     WHERE s.id = $1 `}


  //if _, ok := params["students"]; ok {
  //  //params["q"] = fmt.Sprintf("%%%v%%", params["q"])
  //  sqlParams = append(sqlParams, params["created_at"])
  //  sql = append(sql, fmt.Sprintf(` $%v`, len(sqlParams)))
  //}

  var data []int64
  err := db.Select(&data, strings.Join(sql, " "), sqlParams...)
  if err != nil {
    log.Fatal(err)
  }

  return data
}

func getTutorStudentsByIdTid(db *sqlx.DB, id string, tid string) []SessionStudent {

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

func getTutorStudentByIdTidSid(db *sqlx.DB, params map[string]interface{}) SessionStudent {
  var sqlParams []interface{}
  for _, value := range params {
    sqlParams = append(sqlParams, value)
  }

  log.Println(params, sqlParams)
  var data SessionStudent
  err := db.Get(&data, ` SELECT ss.id id, ss.created_at, ss.status, ss.feedback,
      cs.student_id, ss.rating_interaction "rating.interaction",
      ss.rating_cognition "rating.cognition",
      ss.rating_creativity "rating.creativity"
FROM class_students cs
LEFT JOIN sessions s ON s.class_id = cs.class_id AND s.id = $1
LEFT JOIN sessions_students ss ON ss.session_id = s.id AND cs.student_id = ss.student_id
WHERE tutor_id IS NOT NULL AND cs.student_id = $2`, sqlParams...)
  if err != nil {
    log.Fatal(err)
  }

  return data
}
func ListData(params map[string]interface{}) ([]Session, *sqlx.DB) {
  db := myShared.Connect()

  var data []Session
  sql := []string{`SELECT id, created_at, class_id, tutor_id
    FROM sessions`}
  var sqlParams []interface{}
  if len(params) > 0 {
    sql = append(sql, "WHERE")

    if _, ok := params["cid"]; ok {
      sqlParams = append(sqlParams, params["cid"])
      sql = append(sql, "class_id = $1")
    }

    if _, ok := params["tid"]; ok {
      sqlParams = append(sqlParams, params["tid"])
      sql = append(sql, "tutor_id = $1 and created_at >= now() - interval '1 day'")
    }
  }
  //log.Println(sqlParams,strings.Join(sql, " "))
  err := db.Select(&data, strings.Join(sql, " "), sqlParams...)
  if err != nil {
    log.Fatal(err)
  }

  return data, db
}

func ItemData(id string) (Session, *sqlx.DB) {
  db := myShared.Connect()

  var data Session
  err := db.Get(&data, `SELECT id, created_at, class_id, tutor_id
    FROM sessions
    WHERE id = $1`, id)
  if err != nil {
    log.Fatal(err)
  }

  return data, db
}
func CreateByClassIdData(id string) myShared.Response {
  db := myShared.Connect()
  var lastSessionId int64
  db.QueryRowx(`INSERT INTO sessions(class_id, tutor_id)
  		VALUES ($1,$2) RETURNING id`, id, myShared.CurrentAuth.Id).Scan(&lastSessionId)

  return myShared.Response{Message: "", Id: lastSessionId}
}

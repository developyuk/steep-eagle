package controllers

import (
	myModels "../models"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

type SessionLinks struct {
	LinksSelf
	Class Href `json:"class"`
	// Students Href   `json:"students"`
	Tutors []Href `json:"tutors"`
}
type SessionEmbedded struct {
	Tutors []myModels.SessionTutor `json:"tutors"`
}
type SessionTutorLinks struct {
	LinksSelf
	Tutor    Href   `json:"tutor"`
	Students []Href `json:"students"`
}

type SessionTutorEmbedded struct {
	myModels.Hal
	Students []myModels.SessionStudent `json:"students,omitempty"`
}

type SessionStudentsLinks struct {
	LinksSelf
	Student Href `json:"student"`
}

var pathSessions string = "/sessions"

func GetSessionTutors(db *sqlx.DB, id int) []Href {
	sessionTutors := myModels.GetSessionTutorsFromId(db, id)
	var data []Href
	for _, v := range sessionTutors {
		data = append(data, Href{Href: fmt.Sprintf("%v/%v%v/%v", pathSessions, id, pathTutors, v.TutorId)})
	}

	return data
}

func GetSessionTutorsEmbedded(db *sqlx.DB, id int) []myModels.SessionTutor {
	sessionTutors := myModels.GetSessionTutorsFromId(db, id)
	data := sessionTutors
	for i, v := range data {
		data[i] = v
		data[i].Links = SessionTutorLinks{
			LinksSelf: LinksSelf{
				Self: Href{Href: fmt.Sprintf("%v/%v%v/%v", pathSessions, id, pathTutors, v.TutorId)},
			},
			Tutor: Href{Href: fmt.Sprintf("%v/%v", pathTutors, v.TutorId)},
			// Students: Href{Href: fmt.Sprintf("%v/%v%v/%v%v", pathSessions, id, pathTutors, v.Id, pathStudents)},
			Students: GetSessionStudents(db, id, v.TutorId),
		}
		data[i].Embedded = SessionTutorEmbedded{
			Students: GetSessionStudentsEmbedded(db, id, v.TutorId),
		}
	}

	return data
}

func GetSessionStudents(db *sqlx.DB, id int, tid int) []Href {
	sessionStudents := myModels.GetSessionStudentsFromId(db, id, tid)
	var data []Href
	for _, v := range sessionStudents {
		data = append(data, Href{Href: fmt.Sprintf("%v/%v%v/%v%v/%v", pathSessions, id, pathTutors, tid, pathStudents, v.StudentId)})
	}

	return data
}

func GetSessionStudentsEmbedded(db *sqlx.DB, id int, tid int) []myModels.SessionStudent {
	sessionStudents := myModels.GetSessionStudentsFromId(db, id, tid)
	data := sessionStudents
	for i, v := range data {
		data[i] = v
		data[i].Links = SessionStudentsLinks{
			LinksSelf: LinksSelf{
				Self: Href{Href: fmt.Sprintf("%v/%v%v/%v%v/%v", pathSessions, id, pathTutors, tid, pathStudents, v.StudentId)},
			},
			Student: Href{Href: fmt.Sprintf("%v/%v", pathStudents, v.StudentId)},
		}
	}

	return data
}

func GetSessions(c echo.Context) error {
	data, db := myModels.GetSessions()

	for i, v := range data {
		data[i].Links = SessionLinks{
			LinksSelf: LinksSelf{Self: Href{Href: fmt.Sprintf("%v/%v", pathSessions, v.Id)}},
			Class:     Href{Href: fmt.Sprintf("%v/%v", pathClasses, v.ClassId)},
			// Tutors:    Href{Href: fmt.Sprintf("%v/%v/%v", pathSessions, v.Id, "tutors")},
			Tutors: GetSessionTutors(db, v.Id),
		}
		data[i].Embedded = SessionEmbedded{
			Tutors: GetSessionTutorsEmbedded(db, v.Id),
		}
	}

	response := myModels.Hal{
		Links:    LinksSelf{Self: Href{Href: pathSessions}},
		Embedded: data,
	}
	return c.JSON(http.StatusOK, response)
}

func GetSession(c echo.Context) error {
	response, db := myModels.GetSession(c.Param("id"))
	response.Links = SessionLinks{
		LinksSelf: LinksSelf{Self: Href{Href: fmt.Sprintf("%v/%v", pathSessions, response.Id)}},
		Class:     Href{Href: fmt.Sprintf("%v/%v", pathClasses, response.ClassId)},
		// Tutors:    Href{Href: fmt.Sprintf("%v/%v/%v", pathSessions, response.Id, "tutors")},
		Tutors: GetSessionTutors(db, response.Id),
	}
	response.Embedded = SessionEmbedded{
		Tutors: GetSessionTutorsEmbedded(db, response.Id),
	}
	return c.JSON(http.StatusOK, response)
}

func GetSessionTutorsFromIdTid(c echo.Context) error {
	db, _ := myModels.Connect()
	id, _ := strconv.Atoi(c.Param("id"))
	tid, _ := strconv.Atoi(c.Param("tid"))
	data := GetSessionStudentsEmbedded(db, id, tid)

	response := myModels.Hal{
		Links: SessionTutorLinks{
			LinksSelf: LinksSelf{Self: Href{Href: fmt.Sprintf("%v/%v%v/%v", pathSessions, id, pathTutors, tid)}},
			Tutor:     Href{Href: fmt.Sprintf("%v/%v", pathTutors, tid)},
			// Students: Href{Href: fmt.Sprintf("%v/%v%v/%v%v", pathSessions, id, pathTutors, v.Id, pathStudents)},
			Students: GetSessionStudents(db, id, tid),
		},
		Embedded: data,
	}

	// response.Links = LinksSelf{Self: Href{Href: fmt.Sprintf("%v/%v%v/%v", pathSessions, id, pathTutors, tid)}}

	return c.JSON(http.StatusOK, response)
}

func GetSessionTutorsFromIdTidSid(c echo.Context) error {
	db, _ := myModels.Connect()
	id, _ := strconv.Atoi(c.Param("id"))
	tid, _ := strconv.Atoi(c.Param("tid"))
	sid, _ := strconv.Atoi(c.Param("sid"))
	response := myModels.GetSessionStudentFromId(db, id, tid, sid)
	response.Links = LinksSelf{Self: Href{Href: fmt.Sprintf("%v/%v%v/%v%v/%v", pathSessions, id, pathTutors, tid, pathStudents, sid)}}

	return c.JSON(http.StatusOK, response)
}

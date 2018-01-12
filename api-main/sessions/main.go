package sessions

import (
	myShared "../shared"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"net/http"
)

type SessionLinks struct {
	myShared.LinksSelf
	Class  myShared.Href   `json:"class"`
	Tutors []myShared.Href `json:"tutors"`
	// Students []myShared.Href `json:"students"`
}

// type SessionEmbedded struct {
// 	Tutors []SessionTutor `json:"tutors"`
// }
type SessionTutorLinks struct {
	myShared.LinksSelf
	Tutor    myShared.Href   `json:"tutor"`
	// Students []myShared.Href `json:"students,omitempty"`
}

type SessionTutorEmbedded struct {
	myShared.Hal
	Students []SessionStudent `json:"students,omitempty"`
}

type SessionStudentsLinks struct {
	myShared.LinksSelf
	Student myShared.Href `json:"student"`
}

var (
	db   *sqlx.DB
	list []Session
	item Session
)

// func GetSessionTutorsEmbedded(db *sqlx.DB, id int) []SessionTutor {
// 	sessionTutors := GetSessionTutorsFromId(db, id)
// 	data := sessionTutors
// 	for i, v := range data {
// 		data[i] = v
// 		data[i].Links = SessionTutorLinks{
// 			LinksSelf: myShared.LinksSelf{
// 				Self: myShared.Href{Href: fmt.Sprintf("%v/%v%v/%v", myShared.PathSessions, id, myShared.PathTutors, v.TutorId)},
// 			},
// 			Tutor: myShared.Href{Href: fmt.Sprintf("%v/%v", myShared.PathTutors, v.TutorId)},
// 			// Students: Href{Href: fmt.Sprintf("%v/%v%v/%v%v", pathSessions, id, pathTutors, v.Id, pathStudents)},
// 			Students: GetSessionStudents(db, id, v.TutorId),
// 		}
// 		data[i].Embedded = SessionTutorEmbedded{
// 			Students: itemTutorStudentsEmbedded(db, id, v.TutorId),
// 		}
// 	}
//
// 	return data
// }

func itemLinksTutors(id int) []myShared.Href {
	tutors := getTutorsById(db, id)
	var data []myShared.Href
	for _, v := range tutors {
		data = append(data, myShared.Href{Href: fmt.Sprintf("%v/%v%v/%v", myShared.PathSessions, id, myShared.PathTutors, v)})
	}

	return data
}

// func itemLinksStudents(id int) []myShared.Href {
// 	students := getStudentsById(db, id)
// 	var data []myShared.Href
// 	for _, v := range students {
// 		data = append(data, myShared.Href{Href: fmt.Sprintf("%v/%v%v/%v", myShared.PathSessions, id, myShared.PathStudents, v)})
// 	}
//
// 	return data
// }

func itemLinks(v Session) SessionLinks {
	return SessionLinks{
		LinksSelf: myShared.LinksSelf{Self: myShared.Href{Href: fmt.Sprintf("%v/%v", myShared.PathSessions, v.Id)}},
		Class:     myShared.Href{Href: fmt.Sprintf("%v/%v", myShared.PathClasses, v.ClassId)},
		Tutors:    itemLinksTutors(v.Id),
		// Students:  itemLinksStudents(v.Id),
	}
}

func List(c echo.Context) error {
	list, db = ListData()

	for i, v := range list {
		list[i].Links = itemLinks(v)
		// list[i].Embedded = SessionEmbedded{
		// 	Tutors: GetSessionTutorsEmbedded(db, v.Id),
		// }
	}

	response := myShared.Hal{
		Links:    myShared.LinksSelf{Self: myShared.Href{Href: myShared.PathSessions}},
		Embedded: list,
		Count:    len(list),
		Total:    len(list),
	}
	return c.JSON(http.StatusOK, response)
}

func Item(c echo.Context) error {
	item, db = ItemData(c.Param("id"))
	item.Links = itemLinks(item)
	// item.Embedded = SessionEmbedded{
	// 	Tutors: GetSessionTutorsEmbedded(db, item.Id),
	// }
	return c.JSON(http.StatusOK, item)
}

func itemTutorStudentsEmbedded(db *sqlx.DB, id string, tid string) []SessionStudent {
	tutorStudents := getTutorStudentsByIdTid(db, id, tid)
	data := tutorStudents
	for i, v := range data {
		data[i].Links = SessionStudentsLinks{
			LinksSelf: myShared.LinksSelf{
				Self: myShared.Href{Href: fmt.Sprintf("%v/%v%v/%v%v/%v", myShared.PathSessions, id, myShared.PathTutors, tid, myShared.PathStudents, v.StudentId)},
			},
			Student: myShared.Href{Href: fmt.Sprintf("%v/%v", myShared.PathStudents, v.StudentId)},
		}
	}

	return data
}

// func getTutorStudentsLinks(id string, tid string) []myShared.Href {
// 	tutorStudents := getTutorStudentsByIdTid(db, id, tid)
// 	var data []myShared.Href
// 	for _, v := range tutorStudents {
// 		data = append(data, myShared.Href{Href: fmt.Sprintf("%v/%v%v/%v%v/%v", myShared.PathSessions, id, myShared.PathTutors, tid, myShared.PathStudents, v.StudentId)})
// 	}
//
// 	return data
// }

func ItemTutor(c echo.Context) error {
	db = myShared.Connect()
	id := c.Param("id")
	tid := c.Param("tid")
	data := itemTutorStudentsEmbedded(db, id, tid)

	response := myShared.Hal{
		Links: SessionTutorLinks{
			LinksSelf: myShared.LinksSelf{Self: myShared.Href{Href: fmt.Sprintf("%v/%v%v/%v", myShared.PathSessions, id, myShared.PathTutors, tid)}},
			Tutor:     myShared.Href{Href: fmt.Sprintf("%v/%v", myShared.PathTutors, tid)},
			// Students:  getTutorStudentsLinks(id, tid),
		},
		Embedded: SessionTutorEmbedded{
			Students: data,
		},
	}

	// response.Links = LinksSelf{Self: Href{Href: fmt.Sprintf("%v/%v%v/%v", pathSessions, id, pathTutors, tid)}}

	return c.JSON(http.StatusOK, response)
}

func ItemTutorStudents(c echo.Context) error {
	db := myShared.Connect()
	id:= c.Param("id")
	tid := c.Param("tid")
	sid := c.Param("sid")
	response := getTutorStudentByIdTidSid(db, id, tid, sid)
	response.Links = myShared.LinksSelf{Self: myShared.Href{Href: fmt.Sprintf("%v/%v%v/%v%v/%v", myShared.PathSessions, id, myShared.PathTutors, tid, myShared.PathStudents, sid)}}

	return c.JSON(http.StatusOK, response)
}

func CreateClassSessions(c echo.Context) error {
	id := c.Param("id")
	response := CreateSessionClass(id)
	return c.JSON(http.StatusOK, response)
}

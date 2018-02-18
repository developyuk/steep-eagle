package sessions

import (
  myShared "../shared"
  "fmt"
  "github.com/jmoiron/sqlx"
  "github.com/labstack/echo"
  "net/http"
  "log"
)
type SessionStudentLinks struct {
  myShared.Href
  Session myShared.Href `json:"session"`
}
type SessionLinks struct {
  myShared.LinksSelf
  Class  myShared.Href   `json:"class"`
  Tutors   myShared.Href   `json:"tutors"`
  Students []SessionStudentLinks `json:"students"`
}

// type SessionEmbedded struct {
// 	Tutors []SessionTutor `json:"tutors"`
// }
type SessionTutorLinks struct {
  myShared.LinksSelf
  Tutor myShared.Href `json:"tutor"`
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

func itemLinksStudents(params map[string]interface{}) []SessionStudentLinks {
  students := getStudentsById(db, params)
  var data []SessionStudentLinks
  for _, v := range students {
    data = append(data, SessionStudentLinks{
      Href:myShared.Href{Href: fmt.Sprintf("%v/%v", myShared.PathStudents, v)},
      Session:myShared.Href{Href: fmt.Sprintf("%v/%v%v/%v", myShared.PathSessions, params["id"], myShared.PathStudents, v)}})
    //data = append(data, myShared.Href{Href: fmt.Sprintf("%v/%v%v/%v", myShared.PathSessions, params["id"], myShared.PathStudents, v)})
  }


  return data
}

func itemLinks(v Session, params map[string]interface{}) SessionLinks {
  params["id"] = v.Id
  return SessionLinks{
    LinksSelf: myShared.LinksSelf{Self: myShared.Href{Href: fmt.Sprintf("%v/%v", myShared.PathSessions, v.Id)}},
    Class:     myShared.Href{Href: fmt.Sprintf("%v/%v", myShared.PathClasses, v.ClassId)},
    Tutors:   myShared.Href{Href: fmt.Sprintf("%v/%v", myShared.PathTutors, v.TutorId)},
    Students: itemLinksStudents(params),
  }
}
func listParams(params map[string]interface{}) []Session {
  list, db = ListData(params)

  for i, v := range list {
    list[i].Links = itemLinks(v, params)
    // list[i].Embedded = SessionEmbedded{
    // 	Tutors: GetSessionTutorsEmbedded(db, v.Id),
    // }
  }

  return list
}
func List(c echo.Context) error {
  params := make(map[string]interface{})
  list := listParams(params)
  response := myShared.Hal{
    Links:    myShared.LinksSelf{Self: myShared.Href{Href: myShared.PathSessions}},
    Embedded: list,
    Count:    len(list),
    Total:    len(list),
  }
  return c.JSON(http.StatusOK, response)
}

func ListByTutorId(c echo.Context) error {
  tid := c.Param("id")

  params := make(map[string]interface{})
  params["tid"] = tid

  if students := c.QueryParam("students"); students != "" {
    params["students"] = students
  }

  list := listParams(params)
  response := myShared.Hal{
    Links:    myShared.LinksSelf{Self: myShared.Href{Href: fmt.Sprintf("%v/%v%v", myShared.PathTutors, tid, myShared.PathSessions)}},
    Embedded: list,
    Count:    len(list),
    Total:    len(list),
  }
  return c.JSON(http.StatusOK, response)
}

func ListByClassId(c echo.Context) error {
  cid := c.Param("id")

  params := make(map[string]interface{})
  params["cid"] = cid
  list := listParams(params)
  response := myShared.Hal{
    Links:    myShared.LinksSelf{Self: myShared.Href{Href: fmt.Sprintf("%v/%v%v", myShared.PathClasses, cid, myShared.PathSessions)}},
    Embedded: list,
    Count:    len(list),
    Total:    len(list),
  }
  return c.JSON(http.StatusOK, response)
}

func Item(c echo.Context) error {
  params := make(map[string]interface{})
  params["id"] = c.Param("id")
  item, db = ItemData(params["id"].(string))
  item.Links = itemLinks(item, params)
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

  params := make(map[string]interface{})
  if id := c.Param("id"); id != "" {
    params["id"] = id
  }
  if sid := c.Param("sid"); sid != "" {
    params["sid"] = sid
  }
  log.Println(params)
  response := getTutorStudentByIdTidSid(db, params)
  response.Links = myShared.LinksSelf{Self: myShared.Href{Href: fmt.Sprintf("%v/%v%v/%v", myShared.PathSessions, params["id"], myShared.PathStudents, params["sid"])}}

  return c.JSON(http.StatusOK, response)
}

func CreateByClassId(c echo.Context) error {
  id := c.Param("id")
  response := CreateByClassIdData(id)
  return c.JSON(http.StatusOK, response)
}

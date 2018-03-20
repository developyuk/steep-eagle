package users

import (
  myShared "../shared"
  mySharedRest "../shared/rest"
  "fmt"
  "strings"
  "net/http"
  "time"
)

type (
  User struct {
    myShared.Hal
    Id       uint64    `json:"id"`
    Username string    `json:"username,omitempty"`
    Email    string    `json:"email,omitempty"`
    Pass     string    `json:"pass,omitempty"`
    Role     string    `json:"role,omitempty"`
    Name     string    `json:"name,omitempty"`
    Dob      time.Time `json:"dob,omitempty"`
    Photo    string    `json:"photo,omitempty"`
    UserId   uint64    `json:"user_id,omitempty"`
  }
)

const (
  PathUsers    string = "/users"
  PathStudents string = "/students"
  PathTutors   string = "/tutors"
)

func getRole(path string) string {
  role := strings.TrimRight(strings.TrimPrefix(path, "/"), "s")
  role = strings.TrimRight(strings.TrimPrefix(path, "/"), "s/:id")
  return role
}

func getPath(role string) string {
  path := PathUsers
  if role == "student" {
    path = PathStudents
  }
  if role == "tutor" {
    path = PathTutors
  }
  return path
}

func ItemRest(req *mySharedRest.Request, role string, id string, item *User) (*http.Response, error) {
  rest := mySharedRest.New().GetItem("/_users_full").ParseRequest(req)

  if role != "user" {
    rest.SetQuery("role=eq." + role)
  }

  resp, err := rest.SetQuery(myShared.RequestRest{
    Id: "eq." + id,
  }).End(&item)
  if err != nil {
    return resp, err
  }

  item.setItemLinks(getPath(role))
  return resp, nil
}

func (v *User) setItemLinks(path string) *User {
  v.Links = myShared.CreateHrefSelf(path + "/" + fmt.Sprint(v.Id))
  return v

}

func itemByUsername(param *UserLoginRequest) (User, error) {

  var item User

  if _, err := mySharedRest.New().GetItem("/_users_full").
    SetQuery(myShared.RequestRest{Select: "id,email,role,name,photo"}).
    SetQuery("username=eq." + param.Username).
    End(&item); err != nil {
    return item, err
  }

  return item, nil

}

//func itemByEmailPass(param *UserLoginRequest) (myShared.User, error) {
//
//  var item myShared.User
//  _, err := myShared.PostItem(map[string]interface{}{
//    "data": &item,
//    "path": "/rpc/user_by_email_pass",
//    "query": map[string]string{
//      "_email": param.Email,
//      "_pass":  param.Pwd,
//    },
//  })
//
//  if err != nil {
//    return item, err
//  }
//
//  return item, err
//
//}

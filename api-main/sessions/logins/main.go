package logins

import (
  mySharedRest "../../shared/rest"
  "fmt"
)

type (
  SessionLogin struct {
    UserId uint64 `json:"user_id"`
    Status bool   `json:"status"`
  }
)

const PATH = "/sessions_logins"

func Create(id uint64) {
  var item SessionLogin
  if _, err := mySharedRest.New().GetItem(PATH).
    SetQuery("user_id=eq." + fmt.Sprint(id)).
    SetQuery("status=eq.true").
    SetQuery("order=created_at.desc").
    EndStruct(&item); err != nil {

    mySharedRest.New().PostItem(PATH).
      Send(map[string]string{
      "user_id": fmt.Sprint(id),
      "status":  fmt.Sprint(true),
    })
  }

}

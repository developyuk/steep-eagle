package modules

import (
  myShared "../shared"
  "gopkg.in/guregu/null.v3"
)

type Module struct {
  myShared.Hal
  Id           int         `json:"id"`
  Name         string      `json:"name"`
  Image        null.String `json:"image"`
  SessionTotal int         `json:"session_total" db:"session_total"`
}

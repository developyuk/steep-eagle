package modules

import (
  myShared "../shared"
)

type Module struct {
  myShared.Hal
  Id           int    `json:"id"`
  Name         string `json:"name"`
  Image        string `json:"image"`
  SessionTotal int    `json:"session_total" db:"session_total"`
}

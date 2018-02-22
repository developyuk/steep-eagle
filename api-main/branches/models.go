package branches

import (
  myShared "../shared"
)

type Branch struct {
  myShared.Hal
  Id      int    `json:"id"`
  Name    string `json:"name"`
  Image   string `json:"image"`
  Address string `json:"address"`
}

package rest

import "os"

var (
  DbApiUrl = "http://" + os.Getenv("HOST") + ":3000"
  //DbApiUrl = "http://varnish-db-api:6081"
)

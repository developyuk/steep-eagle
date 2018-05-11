package rest

import "os"

var (
  DbApiUrl = "http://" + os.Getenv("HOST_DB_API") + ":3000"
  //DbApiUrl = "http://varnish-db-api:6081"
)

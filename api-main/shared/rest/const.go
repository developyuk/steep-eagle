package rest

import "os"

var (
  DbApiUrl = "http://"+os.Getenv("DB_API")
  //DbApiUrl = "http://varnish-db-api:6081"
)

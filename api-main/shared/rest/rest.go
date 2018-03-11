package rest

import (
  myShared "../../shared"
  myJwt "../jwt"
  "github.com/parnurzeal/gorequest"
  "net/http"
  "errors"
  "strings"
  "fmt"
  "strconv"
  "log"
  "reflect"
)

type myRest struct {
  limit    uint64
  offset   uint64
  Total    uint64
  Count    uint64
  value    interface{}
  response *http.Response
  request  *gorequest.SuperAgent
  errors   []error
}

func (r *myRest) Clear() *myRest {
  r.limit = 10
  r.errors = nil
  r.request = nil
  r.response = nil
  r.value = nil

  return r
}
func New() *myRest {
  r := &myRest{}
  r.Clear()
  r.request = gorequest.New()
  return r
}
func (r *myRest) End(v interface{}) (*http.Response, error) {

  r.response, _, r.errors = r.request.EndStruct(v)
  //log.Println(errs)
  if r.response.StatusCode < 200 || r.response.StatusCode >= 300 {
    r.errors = append(r.errors, errors.New(r.response.Status))
  }
  hrange := strings.Split(r.response.Header.Get("Content-Range"), "/")
  r.Total, _ = strconv.ParseUint(hrange[len(hrange)-1], 10, 64)

  switch reflect.TypeOf(reflect.ValueOf(v).Elem().Interface()).Kind() {
  //switch reflect.TypeOf(v).Kind() {
  case reflect.Slice:
    s := reflect.ValueOf(v).Elem()

    r.Count = uint64(s.Len())
  }
  log.Println(r.Total, r.Count)
  if len(r.errors) > 0 && r.errors[0] != nil {
    return r.response, r.errors[0]
  }
  defer r.Clear()
  return r.response, nil
}

func (r *myRest) ParseRequest(v *myShared.Request) *myRest {
  if v, err := strconv.ParseUint(v.Limit, 10, 64); err != nil {
    r.errors = append(r.errors, err)
  } else {
    r.SetLimit(v)
  }
  if v, err := strconv.ParseUint(v.Page, 10, 64); err != nil {
    r.errors = append(r.errors, err)
  } else {
    r.SetPage(v)
  }
  return r
}
func (r *myRest) SetLimit(v uint64) *myRest {
  r.request.Query("limit="+ fmt.Sprint(v))
  r.limit = v
  return r
}
func (r *myRest) SetPage(v uint64) *myRest {
  offset := (v - 1) * r.limit
  r.request.Query("offset="+ fmt.Sprint(offset))
  r.offset = offset
  return r
}

func (r *myRest) SetQuery(q interface{}) *myRest {
  r.request.Query(q)
  return r
}

func (r *myRest) GetItems(path string) *myRest {
  r.request.Get(DbApiUrl + path).
    Set("Accept", "application/json").
    Set("Authorization", myJwt.AuthHeader).
    Set("Prefer", "count=exact")

  return r
}

func (r *myRest) GetItem(path string) *myRest {
  r.request.Get(DbApiUrl + path).
    Set("Accept", "application/vnd.pgrst.object+json").
    Set("Authorization", myJwt.AuthHeader)

  return r
}

//func PostItem(params map[string]interface{}) (*http.Response, error) {
//
//  resp, _, errs := gorequest.New().
//    Post(DbApiUrl + params["path"].(string)).
//    Send(params["query"].(map[string]string)).
//    Set("Accept", "application/vnd.pgrst.object+json").
//    Set("Authorization", myJwt.AuthHeader).
//    Set("Prefer", "return=representation").
//    End(params["data"])
//  if errs != nil {
//    return resp, errs[0]
//  }
//  if strings.HasPrefix(params["path"].(string), "/rpc") {
//    if resp.StatusCode != 200 {
//      return resp, errors.New(resp.Status)
//    }
//  }
//
//  if resp.StatusCode != 201 {
//    return resp, errors.New(resp.Status)
//  }
//
//  close()
//  return resp, nil
//}

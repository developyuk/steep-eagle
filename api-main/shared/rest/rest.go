package rest

import (
  mySharedJwt "../../shared/jwt"
  "github.com/parnurzeal/gorequest"
  "net/http"
  "errors"
  "strings"
  "fmt"
  "strconv"
  "reflect"
)

type (
  MyRest struct {
    Limit    uint64
    Offset   uint64
    Total    uint64
    Count    uint64
    value    interface{}
    response *http.Response
    request  *gorequest.SuperAgent
    errors   []error
  }

  Request struct {
    Page   string `json:"page" form:"page" query:"page"`
    Limit  string `json:"limit" form:"limit" query:"limit"`
    Sort   string `json:"sort" form:"sort" query:"sort"`
    Embed  string `json:"embed" form:"embed" query:"embed"`
    Select string `json:"select" form:"select" query:"select"`
  }
)

func (r *MyRest) Clear() *MyRest {
  //r.Limit = 10
  r.errors = nil
  r.request = nil
  r.response = nil
  r.value = nil

  return r
}
func New() *MyRest {
  r := &MyRest{}
  r.Clear()
  r.request = gorequest.New()

  return r
}
func (r *MyRest) End() (*http.Response, error) {
  r.response, _, r.errors = r.request.End()
  if len(r.errors) > 0 && r.errors[0] != nil {
    return r.response, r.errors[0]
  }
  if r.response.StatusCode != 204 {
    r.errors = append(r.errors, errors.New(r.response.Status))
  }
  defer r.Clear()
  return r.response, nil
}
func (r *MyRest) EndStruct(v interface{}) (*http.Response, error) {

  r.response, _, r.errors = r.request.EndStruct(v)
  //spew.Dump(r.response,r.request.Url)

  if r.request.Method == gorequest.GET {

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
  }
  if r.request.Method == gorequest.POST {
    if strings.HasPrefix(r.request.Url, "/rpc") {
      if r.response.StatusCode != 200 {
        r.errors = append(r.errors, errors.New(r.response.Status))
      }
    }

    if r.response.StatusCode != 201 {
      r.errors = append(r.errors, errors.New(r.response.Status))
    }
  }

  if len(r.errors) > 0 && r.errors[0] != nil {
    return r.response, r.errors[0]
  }
  defer r.Clear()
  return r.response, nil
}

func (r *MyRest) ParseRequest(v *Request) *MyRest {
  if v, err := strconv.ParseUint(v.Limit, 10, 64); err != nil {
    //r.errors = append(r.errors, err)
    r.SetLimit(10)
  } else {
    r.SetLimit(v)
  }
  if v, err := strconv.ParseUint(v.Page, 10, 64); err != nil {
    //r.errors = append(r.errors, err)
    r.SetPage(1)
  } else {
    r.SetPage(v)
  }
  if len(v.Select) > 0 {
    r.SetQuery("select=" + v.Select)
  }
  return r
}
func (r *MyRest) SetLimit(v uint64) *MyRest {
  r.request.Query("limit=" + fmt.Sprint(v))
  r.Limit = v
  return r
}
func (r *MyRest) SetPage(v uint64) *MyRest {
  offset := (v - 1) * r.Limit
  r.request.Query("offset=" + fmt.Sprint(offset))
  r.Offset = offset
  return r
}

func (r *MyRest) SetQuery(q interface{}) *MyRest {
  r.request.Query(q)
  return r
}

func (r *MyRest) Send(d interface{}) *MyRest {
  r.request.Send(d)
  return r
}

func (r *MyRest) GetItems(path string) *MyRest {
  r.request.Get(DbApiUrl + path).
    Set("Accept", "application/json").
    Set("Prefer", "count=exact").
    Set("Authorization", mySharedJwt.AuthHeader)

  return r
}

func (r *MyRest) GetItem(path string) *MyRest {
  r.request.Get(DbApiUrl + path).
    Set("Accept", "application/vnd.pgrst.object+json").
    Set("Authorization", mySharedJwt.AuthHeader)

  return r
}

func (r *MyRest) PostItem(path string) *MyRest {
  r.request.
    Post(DbApiUrl + path).
    Set("Accept", "application/vnd.pgrst.object+json").
    Set("Prefer", "return=representation").
    Set("Authorization", mySharedJwt.AuthHeader)

  return r
}
func (r *MyRest) DeleteItem(path string) *MyRest {
  r.request.
    Delete(DbApiUrl + path).
    Set("Accept", "application/json").
    Set("Authorization", mySharedJwt.AuthHeader)

  return r
}

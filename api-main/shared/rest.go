package shared

import (
  "github.com/parnurzeal/gorequest"
  "net/http"
  "errors"
)

const (
  DbApiUrl = "http://db-api:3000"
)

func GetItems(params map[string]interface{}) (*http.Response, error) {
  req := gorequest.New()
  req.Get(DbApiUrl + params["path"].(string)).
    Set("Accept", "application/json").
    Set("Authorization", AuthHeader)

  if val, ok := params["query"]; ok {
    req.Query(val.(map[string]string))
  }
  resp, _, errs := req.EndStruct(params["data"])
  //log.Println(errs)
  if errs != nil {
    return resp, errs[0]
  }
  if resp.StatusCode != 200 {
    return resp, errors.New(resp.Status)
  }

  if errs != nil {
    return resp, errs[0]
  }
  return resp, nil

}

func GetItem(params map[string]interface{}) (*http.Response, error) {
  resp, _, errs := gorequest.New().
    Get(DbApiUrl + params["path"].(string)).
    Query(params["query"].(map[string]string)).
    Set("Accept", "application/vnd.pgrst.object+json").
    Set("Authorization", AuthHeader).
    EndStruct(params["data"])
  if errs != nil {
    return resp, errs[0]
  }
  if resp.StatusCode != 200 {
    return resp, errors.New(resp.Status)
  }
  
  if errs != nil {
    return resp, errs[0]
  }

  return resp, nil
}

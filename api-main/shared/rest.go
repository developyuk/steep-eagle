package shared

import (
  "gopkg.in/resty.v1"
)

const (
  DbApiUrl = "http://db-api:3000"
)

func RestItems(params map[string]interface{}) *resty.Response {
  req := resty.R().
    SetHeader("Accept", "application/json").
    SetHeader("Authorization", AuthHeader)

  if val, ok := params["query"].(map[string]string); ok {
    req.SetQueryParams(val)
  }
  resp, _ := req.Get(DbApiUrl + params["path"].(string))
  return resp

}

func RestItem(params map[string]interface{}) *resty.Response {
  resp, _ := resty.R().
    SetQueryParams(params["query"].(map[string]string)).
    SetHeader("Accept", "application/vnd.pgrst.object+json").
    SetHeader("Authorization", AuthHeader).
    Get(DbApiUrl + params["path"].(string))
  return resp
}

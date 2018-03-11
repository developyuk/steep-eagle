package shared

import (
  "fmt"
  "strconv"
)

type (
  Href struct {
    Href string `json:"href,omitempty"`
  }

  LinksPagination struct {
    LinksSelf
    First *Href `json:"first,omitempty"`
    Prev  *Href `json:"prev,omitempty"`
    Next  *Href `json:"next,omitempty"`
    Last  *Href `json:"last,omitempty"`
  }
  LinksSelf struct {
    Self *Href `json:"self"`
  }
  EmbeddedItems struct {
    Items interface{} `json:"items"`
  }

  Hal struct {
    Links    interface{} `json:"_links,omitempty"`
    Embedded interface{} `json:"_embedded,omitempty"`
    Total    uint64      `json:"total,omitempty"`
    Count    uint64      `json:"count,omitempty"`
  }

  response struct {
    Message string `json:"message"`
  }
)

func CreateEmbeddedItems(vv interface{}) EmbeddedItems {
  return EmbeddedItems{Items: vv}
}
func CreateResponse(message string) response {
  return response{Message: message}
}

func CreateHrefSelf(v string) LinksSelf {
  return LinksSelf{Self: CreateHref(v)}
}
func CreateHalLinks(uri string, path string, total uint64, req *Request) LinksPagination {
  var links LinksPagination
  links.Self = CreateHref(uri)

  page, _ := strconv.ParseUint(req.Page, 10, 64)
  limit, _ := strconv.ParseUint(req.Limit, 10, 64)
  if page > 0 {
    if limit == 0 {
      limit = 10
    }

    prev := page - 1
    if prev < 1 {
      prev = 1
    }
    last := total / limit
    if last < 1 {
      last = 1
    }
    next := page + 1
    if next > last {
      next = last
    }

    links.First = CreateHref(path + "?page=1")
    links.Prev = CreateHref(path + "?page=" + fmt.Sprint(prev))
    links.Next = CreateHref(path + "?page=" + fmt.Sprint(next))
    links.Last = CreateHref(path + "?page=" + fmt.Sprint(last))
  }
  return links
}

func CreateHref(v string) *Href {
  return &Href{Href: v}
}

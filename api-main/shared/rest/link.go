package rest
//
//import (
//  "strconv"
//  "fmt"
//)
//
//func CreateHalLinks(uri string, path string, rest mySharedRest) LinksPagination {
//  var links LinksPagination
//  links.Self = CreateHref(uri)
//
//  page, _ := strconv.ParseUint(rest.Offset, 10, 64)
//  Limit, _ := strconv.ParseUint(rest.Limit, 10, 64)
//  if page > 0 {
//    if Limit == 0 {
//      Limit = 10
//    }
//
//    prev := page - 1
//    if prev < 1 {
//      prev = 1
//    }
//    last := rest.Total / Limit
//    if last < 1 {
//      last = 1
//    }
//    next := page + 1
//    if next > last {
//      next = last
//    }
//
//    links.First = CreateHref(path + "?page=1")
//    links.Prev = CreateHref(path + "?page=" + fmt.Sprint(prev))
//    links.Next = CreateHref(path + "?page=" + fmt.Sprint(next))
//    links.Last = CreateHref(path + "?page=" + fmt.Sprint(last))
//  }
//  return links
//}

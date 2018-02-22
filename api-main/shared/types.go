package shared

// type Hal struct {
// 	Links    interface{} `json:"_links"`
// 	Embedded interface{} `json:"_embedded,omitempty"`
// }

type Href struct {
  Href string `json:"href"`
}

type LinksSelf struct {
  Self Href `json:"self"`
}

type Hal struct {
  Links    interface{} `json:"_links,omitempty"`
  Embedded interface{} `json:"_embedded,omitempty"`
  Total    int         `json:"total,omitempty"`
  Count    int         `json:"count,omitempty"`
}

//
// type Href struct {
// 	Href string `json:"href"`
// }
//
// type LinksSelf struct {
// 	Self Href `json:"self"`
// }

type Response struct {
  Message string `json:"message"`
  Id      int64  `json:"id,omitempty"`
}

func CreateSelfHref(href string) LinksSelf {
  return LinksSelf{Self: CreateHref(href)}
}

func CreateHref(href string) Href {
  return Href{Href: href}
}

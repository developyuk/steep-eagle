package controllers

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

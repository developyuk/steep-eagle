package shared

type (
  Request struct {
    Page  string `json:"page" form:"page" query:"page"`
    Limit string `json:"limit" form:"limit" query:"limit"`
    Sort  string `json:"sort" form:"sort" query:"sort"`
    Embed string `json:"embed" form:"embed" query:"embed"`
    Select string `json:"select" form:"select" query:"select"`
  }

  RequestRest struct {
    Id string `json:"id,omitempty"`
    Select string `json:"select,omitempty"`
  }
)

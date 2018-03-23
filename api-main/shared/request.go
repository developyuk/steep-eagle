package shared

type (
  RequestRest struct {
    Id string `json:"id,omitempty"`
    Select string `json:"select,omitempty"`
  }
)

package shared

const (
  PathSessions      string = "/sessions"
  TimeZone = "Asia/Jakarta"
)

type (
  ProgramModule struct {
    Id        uint64 `json:"id"`
    ProgramId uint64 `json:"program_id"`
    ModuleId  uint64 `json:"module_id"`
  }
)

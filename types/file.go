package types

type File struct {
	Name     string `json:"name"`
	Path     string `json:"path"`
	Preview  string `json:"preview"`
	ThreadId int64  `json:"threadId"`
}

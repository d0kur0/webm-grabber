package types

type Thread struct {
	ID    int64 `json:"id"`
	Board Board `json:"board"`
}

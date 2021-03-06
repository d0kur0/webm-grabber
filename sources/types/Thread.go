package types

import (
	"fmt"
)

type Thread struct {
	ID    int64
	Board Board
}

func (thread *Thread) StringId() string {
	return fmt.Sprint(thread.ID)
}
